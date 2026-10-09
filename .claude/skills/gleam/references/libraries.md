# Core libraries & ecosystem

Versions noted are the latest at time of writing (Sept 2026); check `manifest.toml` / `build/packages/<pkg>/src` for the version actually in use, and search https://packages.gleam.run for others. Prefer core-team packages (C8 in `conventions.md`).

## Contents
1. gleam_json
2. gleam_http (Request/Response types)
3. gleam_time
4. Web servers: wisp + mist (complete example)
5. Ecosystem map

---

## 1. gleam_json (3.x, both targets)

Encoding:
```gleam
import gleam/json

pub fn user_to_json(user: User) -> json.Json {
  let User(name:, age:, email:) = user
  json.object([
    #("name", json.string(name)),
    #("age", json.int(age)),
    #("email", json.nullable(email, json.string)),      // Option(a) → value | null
    #("tags", json.array(user.tags, json.string)),
  ])
}
json.to_string(user_to_json(user))       // or json.to_string_tree for large output
```
Also: `json.bool`, `json.float`, `json.null()`, `json.dict(d, key_to_string, value_to_json)`, `json.preprocessed_array(List(Json))`.

Decoding uses `gleam/dynamic/decode` (see `stdlib.md`):
```gleam
json.parse(from: text, using: user_decoder())      // Result(User, json.DecodeError)
json.parse_bits(from: bits, using: user_decoder())
```
`json.DecodeError`: `UnexpectedEndOfInput | UnexpectedByte(String) | UnexpectedSequence(String) | UnableToDecode(List(decode.DecodeError))`.
The LSP can generate both the decoder and the to-JSON function for a custom type.

## 2. gleam_http (4.x, both targets)

Shared `Request(body)` / `Response(body)` types used by all servers and clients.
```gleam
import gleam/http
import gleam/http/request
import gleam/http/response

let assert Ok(req) = request.to("https://api.example.com/users?page=2")
let req =
  req
  |> request.set_method(http.Post)                  // Get Post Put Patch Delete Head Options Connect Trace Other(String)
  |> request.set_header("content-type", "application/json")
  |> request.set_body(json.to_string(payload))
request.get_query(req)   request.path_segments(req)   request.get_header(req, "authorization")

response.new(201) |> response.set_body("created") |> response.set_header("location", "/users/1")
```
Headers are lowercase `#(String, String)` lists (the `http.Header` alias is deprecated). Cookies: `request.get_cookies`, `response.set_cookie(resp, name, value, cookie.defaults(http.Https))`.
Clients: `gleam_httpc` (Erlang: `httpc.send(req)`), `gleam_fetch` (JS, returns a promise). Use the sans-IO pattern for SDKs.

## 3. gleam_time (1.x, both targets)

```gleam
import gleam/time/duration
import gleam/time/timestamp
import gleam/time/calendar

let now = timestamp.system_time()
let later = timestamp.add(now, duration.minutes(30))
timestamp.to_rfc3339(now, calendar.utc_offset)        // "2026-09-27T12:00:00Z"
timestamp.parse_rfc3339("2026-09-27T12:00:00Z")      // Result(Timestamp, Nil)
timestamp.to_unix_seconds(now)  timestamp.from_unix_seconds(1_700_000_000)
timestamp.difference(start, end) |> duration.to_seconds
let #(date, time) = timestamp.to_calendar(now, calendar.local_offset())
// calendar.Date(year:, month: calendar.September, day:), calendar.TimeOfDay(hours:, minutes:, seconds:, nanoseconds:)
```
`Timestamp` is a UTC instant; `Duration` a signed span; calendar types are only for human-facing conversion. Store/transmit timestamps, not calendar dates. Use `duration.*` constructors instead of raw millisecond ints in APIs.

## 4. Web servers: wisp + mist (Erlang)

`mist` is the HTTP server; `wisp` is the framework on top (middleware, bodies, static files, cookies, logging). Verified with wisp 2.2 / mist 6.0:

```gleam
// src/web.gleam
import envoy
import gleam/erlang/process
import mist
import web/router
import wisp
import wisp/wisp_mist

pub fn main() -> Nil {
  wisp.configure_logger()
  let assert Ok(secret_key_base) = envoy.get("SECRET_KEY_BASE")
    as "SECRET_KEY_BASE must be set"
  let assert Ok(_) =
    wisp_mist.handler(router.handle_request, secret_key_base)
    |> mist.new
    |> mist.bind("0.0.0.0")
    |> mist.port(8000)
    |> mist.start
  process.sleep_forever()
}
```
```gleam
// src/web/router.gleam
import gleam/dynamic/decode
import gleam/http.{Get, Post}
import gleam/json
import wisp.{type Request, type Response}

pub fn handle_request(request: Request) -> Response {
  use request <- middleware(request)
  case wisp.path_segments(request) {
    [] -> wisp.html_response("<h1>Hello</h1>", 200)
    ["notes"] -> notes(request)
    ["notes", id] -> show_note(request, id)
    _ -> wisp.not_found()
  }
}

fn middleware(request: Request, next: fn(Request) -> Response) -> Response {
  let request = wisp.method_override(request)
  use <- wisp.log_request(request)
  use <- wisp.rescue_crashes
  use request <- wisp.handle_head(request)
  next(request)
}

fn notes(request: Request) -> Response {
  case request.method {
    Get -> wisp.json_response(json.to_string(json.array([], json.string)), 200)
    Post -> create_note(request)
    _ -> wisp.method_not_allowed([Get, Post])
  }
}

fn create_note(request: Request) -> Response {
  use body <- wisp.require_json(request)
  case decode.run(body, decode.field("text", decode.string, decode.success)) {
    Ok(text) ->
      json.object([#("text", json.string(text))])
      |> json.to_string
      |> wisp.json_response(201)
    Error(_) -> wisp.unprocessable_content()
  }
}

fn show_note(request: Request, id: String) -> Response {
  use <- wisp.require_method(request, Get)
  wisp.json_response(json.to_string(json.object([#("id", json.string(id))])), 200)
}
```
Other wisp helpers: `wisp.ok()`, `created()`, `no_content()`, `bad_request(detail)`, `require_string_body`, `require_form`, `use <- wisp.serve_static(req, under: "/static", from: priv_dir)`, `wisp.set_cookie/get_cookie`, `wisp.string_body`, `wisp.json_body`. Testing: `wisp/simulate` builds requests for handler unit tests without a server. Put the server in a supervision tree with `mist.supervised(builder)` for production.
Pass shared dependencies (DB pool, config) with a `Context` record argument: `handle_request(request, ctx)`, partially applied: `wisp_mist.handler(router.handle_request(_, ctx), secret)`.

## 5. Ecosystem map

| Need | Package | Notes |
|---|---|---|
| Env vars | `envoy` | `envoy.get("X")` → `Result(String, Nil)`, both targets |
| CLI args | `argv` (raw), `glint` (flags, help) | `argv.load().arguments` |
| Files & dirs | `simplifile` | `simplifile.read(path)` → `Result(String, FileError)`; both targets (Node/Deno/Bun) |
| Paths | `filepath` | join/split/extension |
| Regex | `gleam_regexp` | `regexp.from_string`, `regexp.check`, `regexp.scan` |
| Lazy sequences | `gleam_yielder` | replaces `gleam/iterator` |
| Deque | `gleam_deque` | replaces `gleam/queue` |
| Crypto / hashing / HMAC / random bytes | `gleam_crypto` | both targets |
| TOML | `tom` | |
| Terminal colours | `gleam_community_ansi` | |
| HTTP server | `mist` (+ `wisp`) | Erlang |
| HTTP client | `gleam_httpc` (Erlang), `gleam_fetch` (JS) | |
| Frontend / SPA / components / SSR | `lustre` (+ `lustre_dev_tools`) | Elm-architecture; both targets |
| PostgreSQL | `pog` | typed SQL generation: `squirrel` |
| SQLite | `sqlight` | |
| Logging | `logging` | Erlang logger config; wisp uses it |
| JS promises / interop | `gleam_javascript` | |
| Snapshot tests | `birdie` | |
| Property tests | `qcheck` | |
| Test runner alternative | `startest` | |

Curated lists: https://github.com/gleam-lang/awesome-gleam, https://packages.gleam.run.
