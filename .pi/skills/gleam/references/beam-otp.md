# BEAM, Erlang & OTP with Gleam

Packages: `gleam_erlang` (1.x: processes, atoms, charlists, nodes) and `gleam_otp` (1.x: actors, supervisors). Both are Erlang-target only. APIs changed substantially at 1.0 — older blog posts/tutorials (`actor.start(state, handler)`, `process.new_selector() |> process.selecting(...)`, `supervisor.start_spec`) are outdated. When in doubt, read `build/packages/gleam_otp/src/gleam/otp/*.gleam`.

## Contents
1. BEAM mental model
2. Processes, subjects, messages
3. Selectors, monitors, links, timers
4. Actors (gleam/otp/actor)
5. Supervisors (static & factory)
6. OTP applications
7. Erlang data representation & interop notes
8. Running & deploying (shell, escript, erlang-shipment, containers)

---

## 1. BEAM mental model

- Lightweight isolated processes (millions possible), preemptively scheduled across cores, each with its own heap and GC. No shared memory; communication is by copying messages.
- "Let it crash": in OTP applications, handle expected errors with `Result`, but let unexpected failures crash a process and have a supervisor restart it into a known-good state. Libraries still must not panic (unless OTP-oriented and designed for supervision).
- Links propagate crashes between processes; monitors notify without crashing. `process.spawn` links to the parent; `process.spawn_unlinked` does not.
- Timeouts are integers in **milliseconds**.
- Atoms are never garbage collected: don't create them from untrusted input (`atom.create`); use `atom.get` for existing ones.
- Gleam Ints are arbitrary-precision on the BEAM; Strings are UTF-8 binaries (not Erlang charlists).

## 2. Processes, subjects, messages (`gleam/erlang/process`)

A `Subject(message)` is a typed channel owned by the process that created it; only the owner can receive from it.

```gleam
import gleam/erlang/process.{type Subject}

let subject = process.new_subject()
process.spawn(fn() { process.send(subject, "hello") })
let assert Ok(message) = process.receive(subject, within: 1000)  // Result(msg, Nil) on timeout
process.receive_forever(subject)
```

Request/response: put a reply subject in the message; `process.call` creates it, sends, and waits (panics on timeout).
```gleam
pub type Message { Get(reply_to: Subject(Int)) }
let value = process.call(server, waiting: 1000, sending: Get)   // Get is fn(Subject(Int)) -> Message
```

Named processes (for supervisors restarting processes): create the name once, at startup, and pass it down.
```gleam
let name = process.new_name(prefix: "counter")   // Name(Message) — call once, never dynamically in a loop (it creates an atom)
let subject = process.named_subject(name)        // works across restarts
process.register(pid, name)  process.named(name) // Result(Pid, Nil)
```

Other useful functions: `self()`, `sleep(ms)`, `sleep_forever()`, `is_alive(pid)`, `kill(pid)`, `send_exit(pid)`, `send_abnormal_exit(pid, reason)`, `trap_exits(True)`, `subject_owner(subject)`.

## 3. Selectors, monitors, links, timers

Receive from multiple subjects/other messages:
```gleam
let selector =
  process.new_selector()
  |> process.select(for: own_subject)                              // payload type = message type
  |> process.select_map(for: other_subject, mapping: FromOther)    // map to common type
  |> process.select_monitors(ProcessDown)                          // fn(process.Down) -> payload
  |> process.select_trapped_exits(ProcessExited)                   // requires trap_exits(True)
  |> process.select_other(UnexpectedMessage)                       // fn(Dynamic) -> payload
process.selector_receive(from: selector, within: 5000)
```
Monitoring: `let monitor = process.monitor(pid)`, `process.select_specific_monitor(selector, monitor, mapping)`, `process.demonitor_process(monitor)`.
Timers: `let timer = process.send_after(subject, 1000, Tick)`, `process.cancel_timer(timer)`.
Linking: `process.link(pid)`, `process.unlink(pid)`.

## 4. Actors (`gleam/otp/actor`)

Typed equivalent of `gen_server`: holds state, handles one message at a time, integrates with OTP supervision and `sys` debugging.

```gleam
import gleam/erlang/process.{type Subject}
import gleam/otp/actor
import gleam/otp/supervision

pub opaque type Message {
  Increment(by: Int)
  Get(reply_to: Subject(Int))
  Shutdown
}

pub fn start(name: process.Name(Message)) -> actor.StartResult(Subject(Message)) {
  actor.new(0)                       // initial state; data returned = the actor's Subject
  |> actor.named(name)               // optional: register under a name
  |> actor.on_message(handle_message)
  |> actor.start                     // Result(actor.Started(Subject(Message)), actor.StartError)
}

/// Child spec so a supervisor can start/restart this actor.
pub fn supervised(name: process.Name(Message)) -> supervision.ChildSpecification(Subject(Message)) {
  supervision.worker(fn() { start(name) })
}

pub fn increment(counter: Subject(Message), by amount: Int) -> Nil {
  actor.send(counter, Increment(amount))
}

pub fn get(counter: Subject(Message)) -> Int {
  actor.call(counter, waiting: 1000, sending: Get)
}

fn handle_message(state: Int, message: Message) -> actor.Next(Int, Message) {
  case message {
    Increment(by:) -> actor.continue(state + by)
    Get(reply_to:) -> {
      process.send(reply_to, state)
      actor.continue(state)
    }
    Shutdown -> actor.stop()          // or actor.stop_abnormal("reason")
  }
}
```

- `actor.Started(pid:, data:)` — `data` is what the actor returns (by default its `Subject`).
- Custom initialisation (runs inside the new process, e.g. to create extra subjects or selectors):
  ```gleam
  actor.new_with_initialiser(1000, fn(subject) {
    let selector = process.new_selector() |> process.select(subject)
    actor.initialised(initial_state)
    |> actor.selecting(selector)
    |> actor.returning(subject)
    |> Ok
  })
  |> actor.on_message(handle_message)
  |> actor.start
  ```
- `actor.with_selector(actor.continue(state), selector)` changes what the actor listens to.
- `StartError`: `InitTimeout | InitFailed(String) | InitExited(ExitReason)`.
- Wrap the actor behind a small module API (`start`, `supervised`, domain functions); keep `Message` opaque when callers shouldn't construct messages directly.
- Don't do long blocking work inside `handle_message` for actors that must stay responsive — spawn a worker or use a separate actor.

## 5. Supervisors

**Static supervisor** (fixed children, started in order):
```gleam
import gleam/otp/static_supervisor as supervisor

pub fn start_tree() -> actor.StartResult(supervisor.Supervisor) {
  let counter_name = process.new_name("counter")
  supervisor.new(supervisor.OneForOne)            // OneForOne | OneForAll | RestForOne
  |> supervisor.restart_tolerance(intensity: 3, period: 5)
  |> supervisor.add(counter.supervised(counter_name))
  |> supervisor.add(web.supervised(counter_name))
  |> supervisor.start
}
```
Nest supervisors with `supervisor.supervised(builder)`.

Child spec tweaks (`gleam/otp/supervision`): `supervision.worker(start)`, `supervision.supervisor(start)`, `|> supervision.restart(supervision.Transient)` (`Permanent | Transient | Temporary`), `|> supervision.timeout(ms: 5000)`, `|> supervision.significant(True)` (with `supervisor.auto_shutdown`), `|> supervision.map_data(f)`.

Because a restarted actor gets a new `Subject`, talk to supervised actors via **named subjects** (`process.named_subject(name)`) created from a `Name` made once at startup.

**Factory supervisor** (`gleam/otp/factory_supervisor`) — dynamic children of one kind (like `simple_one_for_one`):
```gleam
let assert Ok(started) =
  factory_supervisor.worker_child(fn(user_id) { session.start(user_id) })
  |> factory_supervisor.restart_strategy(supervision.Transient)
  |> factory_supervisor.start
factory_supervisor.start_child(started.data, user_id)
```
Use `factory_supervisor.named` + `factory_supervisor.get_by_name` to reach it from elsewhere; `factory_supervisor.supervised(builder)` to put it in a tree.

## 6. OTP applications

For long-running services, start the tree in `main` and block:
```gleam
pub fn main() -> Nil {
  let assert Ok(_) = start_tree()
  process.sleep_forever()
}
```
Alternatively set `[erlang] application_start_module = "my_app@application"` in `gleam.toml` and implement the OTP application behaviour; `extra_applications = ["inets", "ssl"]` starts additional OTP apps. `application.priv_directory("my_app")` locates the `priv/` directory at runtime.

## 7. Erlang data representation

| Gleam | Erlang |
|---|---|
| `Int`, `Float` | integer, float |
| `String` | `<<"utf8 binary">>` |
| `True`/`False`, `Nil` | `true`/`false`, `nil` |
| `List(a)` | proper list |
| `#(a, b)` | `{A, B}` |
| `Ok(x)` / `Error(e)` | `{ok, X}` / `{error, E}` |
| `Guest` (no fields) | atom `guest` |
| `User(id: 10)` | `{user, 10}` (tag = snake_case atom, fields positional) |
| `BitArray` | bitstring |
| `Dict(k, v)` | map |
| module `my_app/user` | module atom `my_app@user` |

- Charlists (`'text'`) ↔ `gleam/erlang/charlist`.
- Atoms: `gleam/erlang/atom` (`atom.create`, `atom.get`, `atom.to_string`, `atom.decoder()`).
- Data from Erlang with unknown shape: treat as `Dynamic` and decode with `gleam/dynamic/decode`.
- `echo`/`string.inspect` print Gleam syntax on both targets.
- Erlang/Elixir Hex packages can be dependencies directly in `gleam.toml` (compiled with rebar3/mix, which must be installed). Call them via `@external` (see `ffi-targets.md`).

## 8. Running & deploying

- `gleam run` / `gleam run -m my_app/tool` — run `main` of the root/other module.
- `gleam shell` — Erlang REPL with the project loaded; call Gleam modules as `'my_app@user':function(Args).`
- `gleam export escript` — single-file executable `./my_app` (needs Erlang installed on target; since 1.17). Good for CLI tools.
- `gleam export erlang-shipment` — `build/erlang-shipment/` with compiled BEAM files + `entrypoint.sh run`; copy to a server/container with a matching Erlang/OTP. Standard for services.
- Docker images: `ghcr.io/gleam-lang/gleam:v<version>-erlang-alpine` for building; runtime can be `erlang:<otp>-alpine`. Typical Dockerfile: build stage runs `gleam export erlang-shipment`, runtime stage copies `build/erlang-shipment` and runs `/app/entrypoint.sh run`.
- Guides: https://gleam.run/deployment/linux-server/, https://gleam.run/deployment/fly/.
- Observability: Erlang's `logger` via the `logging` package; `observer`/`sys` tools work on Gleam actors.
