# Externals (FFI) and compilation targets

Source: https://gleam.run/documentation/externals/. Rule of thumb: prefer pure Gleam or an existing package; use externals only when there's no suitable alternative. The compiler **cannot check** that an external matches its declared type — you are responsible.

## Contents
1. External functions
2. External types
3. Erlang FFI
4. Elixir FFI
5. JavaScript FFI & data API
6. Targets: choosing, multi-target code, JS-specific notes

---

## 1. External functions

```gleam
@external(erlang, "lists", "reverse")
@external(javascript, "./my_app_ffi.mjs", "reverse_list")
pub fn reverse_list(list: List(a)) -> List(a)
```
- Arguments and return type annotations are mandatory.
- Provide one attribute per supported target. With a Gleam body, the body is the fallback for targets without an `@external`:
  ```gleam
  @external(erlang, "lists", "reverse")
  pub fn reverse_list(list: List(a)) -> List(a) { reverse_and_prepend(list, []) }
  ```
- Calling an external that has no implementation for the current target is a compile error ("Unsupported target").
- Keep externals private where possible and wrap them in a typed Gleam API (convert errors to `Result`, validate inputs).

## 2. External types

```gleam
pub type ZipHandle                 // no constructors: values only come from externals

@external(erlang, "my_app_ffi", "open_zip")
pub fn open(path: String) -> Result(ZipHandle, Nil)
```
- Make specific types (`ZipHandle`, `Buffer`) rather than reusing `Pid` or `Dynamic` — prevents misuse. Never type FFI arguments as `Dynamic` (anti-pattern).
- Since 1.14, an external type can map to a concrete foreign type for generated Erlang specs / TypeScript declarations:
  ```gleam
  @external(erlang, "erlang", "map")
  @external(javascript, "../dict.d.mts", "Dict")
  pub type Dict(key, value)
  ```

## 3. Erlang FFI

Place Erlang sources next to Gleam code in `src/` (e.g. `src/my_app_ffi.erl`). Erlang module names are global — prefix with the package name.

```erlang
-module(my_app_ffi).
-export([badge_count/0, read_env/1]).

badge_count() -> 8.

%% Return shapes that match Gleam types:
read_env(Name) ->
    case os:getenv(binary_to_list(Name)) of
        false -> {error, nil};                     % Error(Nil)
        Value -> {ok, unicode:characters_to_binary(Value)}  % Ok(String)
    end.
```
```gleam
@external(erlang, "my_app_ffi", "read_env")
fn read_env(name: String) -> Result(String, Nil)
```
Data mapping: see `beam-otp.md` §7 (strings are binaries, `Nil` is `nil`, records are tagged tuples `{snake_case_tag, Field1, …}`, field-less variants are atoms). Erlang macros can't be used from Gleam; improper lists are not valid Gleam lists; Erlang functions that return charlists need converting (`gleam/erlang/charlist`).

Calling Gleam from Erlang: module `my_app/user` is atom `my_app@user`; functions keep their names.

## 4. Elixir FFI

Target is still `erlang`; module names need the `Elixir.` prefix:
```gleam
@external(erlang, "Elixir.Pokemon", "badge_count")
pub fn badge_count() -> Int
```
Elixir `.ex` files in `src/` and Elixir Hex deps are compiled with Elixir (must be installed). Elixir macros are not callable. Atoms: Gleam `Guest` is `:guest`; `User(id: 10)` is `{:user, 10}`.

## 5. JavaScript FFI

The module path is used verbatim in a JS `import`, relative to the generated `.mjs` for the Gleam module (source files are mirrored into the build dir, so relative to the `.gleam` file works). Use the `.mjs` extension.

```gleam
// src/my_app.gleam
@external(javascript, "./my_app_ffi.mjs", "now_ms")
fn now_ms() -> Int
```
```javascript
// src/my_app_ffi.mjs
export function now_ms() { return Date.now(); }
```
npm packages: `@external(javascript, "has-flag", "hasFlag")` — manage them with your own `package.json`; Gleam does no npm handling.

Gleam values in JS:

| Gleam | JavaScript |
|---|---|
| `Int`, `Float` | `number` |
| `String` | `string` |
| `Bool` | `boolean` |
| `Nil` | `undefined` |
| Tuple | array (never mutate it) |
| `List` | Gleam linked list — use the prelude API, not arrays |
| `Result`, custom types | class instances — use the generated API |

Generated data API (since 1.13). Custom-type helpers are exported from the compiled module that defines the type (e.g. `./school.mjs`); prelude helpers (Result, List, BitArray, `toList(array)`) come from the package's `gleam.mjs`, which re-exports `prelude.mjs` — from `src/my_app_ffi.mjs` import `"./gleam.mjs"`:
- Custom types: constructor `Type$Variant(...)`, check `Type$isVariant(v)`, fields `Type$Variant$0(v)`, `Type$Variant$label(v)`, shared-label accessor `Type$label(v)`.
- Results: `Result$Ok(x)`, `Result$Error(e)`, `Result$isOk(r)`, `Result$Ok$0(r)`, `Result$Error$0(r)`.
- Lists: `List$Empty()`, `List$NonEmpty(head, tail)`, `List$isEmpty`, `List$NonEmpty$first`, `List$NonEmpty$rest`. `toList(array)` converts a JS array. Inspect `build/dev/javascript/prelude.mjs` for the full list.
- Bit arrays: `BitArray$BitArray(new Uint8Array([...]))`, `BitArray$isBitArray`, `BitArray$BitArray$data`.

Promises / JS interop types: the `gleam_javascript` package (`gleam/javascript/promise`: `new`, `resolve`, `map`, `map_try`, `await`, `try_await`, `tap`, `rescue`, `await_list`, `race_list`, `wait`; write `use x <- promise.await(p)`). HTTP on JS: `gleam_fetch`.

Importing Gleam from JS: compiled modules live in `build/dev/javascript/<package>/<module path>.mjs`; e.g. `import * as user from "./build/dev/javascript/my_app/my_app/user.mjs"`.

## 6. Targets

- Default target is `erlang`; set `target = "javascript"` in `gleam.toml` or pass `--target javascript`. JS runtime: `[javascript] runtime = "node" | "deno" | "bun"` or `--runtime`.
- `gleam new my_app --template javascript` scaffolds a JS project.
- Pure Gleam code (and most of `gleam_stdlib`) works on both targets. `gleam_erlang`/`gleam_otp` are Erlang-only; `gleam_javascript` is JS-only. For libraries, test both: `gleam test --target erlang && gleam test --target javascript`.
- **JS-target differences to keep in mind**:
  - Ints are JS numbers: precision is lost beyond ±2^53 − 1; bitwise ops are 32-bit. Division by zero still returns 0.
  - Float `6.0` prints as `6` in JS output (`float.to_string` handles formatting consistently; don't rely on `string.inspect` for user output).
  - No processes/OTP; concurrency is via promises. Tail calls are optimised only for self-recursion.
  - `echo` output labels records differently between targets (JS shows labels; Erlang shows positional fields).
- Optional JS outputs: `[javascript] typescript_declarations = true` (generates `.d.mts`), `source_maps = true`. Deno permissions under `[javascript.deno]`.
- Frontend apps usually use Lustre (`lustre_dev_tools` for bundling/dev server).
