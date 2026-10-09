---
name: gleam
description: Expert Gleam programming — writing, reviewing, refactoring, debugging and explaining Gleam code and projects on the BEAM (Erlang/OTP) and JavaScript targets. Covers the language (custom types, pattern matching, use, pipes, labels, bit arrays, echo/assert/todo/panic), official conventions, patterns and anti-patterns, gleam_stdlib 1.x API, core libraries (gleam_json, gleam_http, gleam_time, gleam_erlang, gleam_otp actors/supervisors), FFI/externals to Erlang, Elixir and JavaScript, the gleam CLI and gleam.toml, testing with gleeunit, the Gleam language server (LSP) and its code actions, and deployment (escript, erlang-shipment). Use whenever a task involves .gleam files, gleam.toml, manifest.toml, the `gleam` command, Gleam packages from Hex, or questions about Gleam.
---

# Gleam

Gleam is a statically typed, immutable, functional language that compiles to Erlang (BEAM, default) or JavaScript. It has no classes, no exceptions, no null, no loops, no `if`, no type classes, no macros, and no implicit conversions. Everything is an expression.

## Workflow

1. **Orient.** Run `gleam --version` and read `gleam.toml` (name, `target`, deps). Check `manifest.toml` for resolved versions before using a dependency's API. Dependency sources live in `build/packages/<pkg>/src/` — read them when unsure of an API instead of guessing.
2. **Write code** following the rules below and `references/conventions.md`.
3. **Verify** after every change, in this order:
   - `gleam format` (never hand-format; formatter is canonical, `gleam format --check` in CI)
   - `gleam check` (fast type check, no codegen) → `gleam build --warnings-as-errors` when finishing
   - `gleam test` (add `--target javascript` for JS or multi-target code)
   - Treat warnings as bugs: unused values, unused imports, inexhaustive patterns, deprecated APIs, `echo`/`todo` left behind.
4. Compiler errors are precise and usually contain the fix — read the full message and hint.

Erlang target needs Erlang/OTP (`erl`, `escript`) and sometimes `rebar3`; JS target needs Node (default), Deno, or Bun. If Erlang is missing, `gleam check`/`gleam build --target javascript` still works for pure Gleam code.

## Rules that matter most

These are the mistakes most often made when writing Gleam from memory of other languages:

- **Operators are type-specific.** Int: `+ - * / %  < > <= >=`. Float: `+. -. *. /.  <. >. <=. >=.`. String concat: `<>`. No implicit Int↔Float; use `int.to_float`, `float.round`/`truncate`. Integer and float division by zero return `0`/`0.0` (no crash) — use `int.divide`/`float.divide` for a `Result`.
- **Float literals need a decimal point**: `1.0`, not `1`. Digits may use `_` separators.
- **No string interpolation**: `"Hello, " <> name <> "!"`. Convert with `int.to_string`, `float.to_string`, `string.inspect` (any value, debug form).
- **Tuples are `#(a, b)`**, accessed with `.0`/`.1` or destructured `let #(a, b) = pair`.
- **Lists are immutable linked lists**: prepend with `[x, ..rest]`; `list.append` and `list.length` are O(n). There is no index access — use a different structure (`dict`) or pattern matching. Compare with `[]` rather than `list.length(xs) == 0` (compiler warns).
- **Importing types needs `type`**: `import gleam/option.{type Option, None, Some}`. Qualify functions (`list.map`), never import functions unqualified.
- **Fallible functions return `Result(value, error)`**; use `Result(a, Nil)` if no detail. Not `Option`. `Option` is for optional *data* (e.g. record fields).
- **Chain Results with `use`**: `use user <- result.try(fetch_user(id))`. `use` makes the rest of the block a callback — it is not async/await and not early return syntax in general.
- **No `if`**: `case cond { True -> … False -> … }`, or `use <- bool.guard(when: cond, return: value)` for early exits.
- **Recursion instead of loops**; tail-recursive helpers with accumulators are idiomatic and optimised on both targets. Prefer `list.fold`, `list.map`, `list.filter_map`, `int.range`, etc. when they fit.
- **Record field access** `user.name` works only for fields present in every variant (same position/label). Otherwise pattern match. Update with `User(..user, age: 31)`. Label shorthand: `User(name:, age:)` in constructors and patterns.
- **Exhaustive `case`**: match every variant explicitly instead of `_` so the compiler flags new variants.
- **Debugging**: `echo value` (prints with file:line, returns the value, works in pipes: `|> echo`). `io.debug` no longer exists. Remove `echo` before finishing.
- **Tests & assertions**: `assert expression` (e.g. `assert add(1, 2) == 3`) and `let assert Ok(x) = …` (optionally `as "message"`). Test functions are `pub fn something_test()` in `test/`.
- **Libraries never panic** (`panic`, `let assert`, `assert`, `todo`) — return `Result`. Applications may `let assert` at boundaries (e.g. startup config) and OTP code may crash under a supervisor.
- **Removed from the stdlib** (use packages): `gleam/iterator` → `gleam_yielder`, `gleam/regex` → `gleam_regexp`, `gleam/queue` → `gleam_deque`, `gleam/string_builder` → `gleam/string_tree`, `gleam/bytes_builder` → `gleam/bytes_tree`, old `gleam/dynamic` decoders → `gleam/dynamic/decode`, `io.debug` → `echo`. `gleam/erlang` process/OTP APIs are in `gleam_erlang`/`gleam_otp` (1.x APIs differ significantly from pre-1.0 tutorials).
- **Pipe into non-first argument** with a capture: `x |> string.append("prefix", _)`. Writing `a |> f(b)` means `f(a, b)`; if `f(b)` returns a function write `a |> f(b)()` (the implicit form is deprecated since 1.18).
- **Names**: `snake_case` for values/functions/modules, `PascalCase` for types/constructors, acronyms as words (`Json`, `HttpClient`), no abbreviations, singular module names, modules namespaced under the package name (`src/my_pkg/user.gleam`).

## Minimal idiomatic module

```gleam
import gleam/int
import gleam/result

pub type Account {
  Account(id: Int, balance: Int)
}

pub type WithdrawError {
  InvalidAmount(input: String)
  InsufficientFunds(balance: Int, requested: Int)
}

/// Withdraw an amount given as text, e.g. from a form field.
pub fn withdraw(
  account: Account,
  amount input: String,
) -> Result(Account, WithdrawError) {
  use amount <- result.try(
    int.parse(input) |> result.replace_error(InvalidAmount(input)),
  )
  case amount <= account.balance {
    True -> Ok(Account(..account, balance: account.balance - amount))
    False ->
      Error(InsufficientFunds(balance: account.balance, requested: amount))
  }
}
```

## Reference files

Load only what the task needs:

| Task | File |
|---|---|
| Syntax & semantics: types, patterns, `use`, labels, bit arrays, constants, attributes, strings/numbers | `references/language.md` |
| Official conventions, patterns (errors, builders, sans-IO, invalid states) and anti-patterns — **read before writing or reviewing non-trivial code** | `references/conventions.md` |
| `gleam_stdlib` 1.x function signatures (list, string, dict, result, option, int, float, decode, …) | `references/stdlib.md` |
| BEAM & OTP: processes, subjects, selectors, actors, supervisors, Erlang data, deployment | `references/beam-otp.md` |
| FFI/externals (Erlang, Elixir, JavaScript), JS target specifics, multi-target code | `references/ffi-targets.md` |
| CLI commands, `gleam.toml`, dependencies, testing, publishing, LSP features & code actions, editor setup | `references/tooling.md` |
| JSON, HTTP, time and popular ecosystem packages (wisp, mist, lustre, simplifile, envoy, pog, …) | `references/libraries.md` |

Authoritative sources: https://gleam.run/documentation/, https://tour.gleam.run, https://hexdocs.pm/gleam_stdlib/, https://packages.gleam.run. For a package's exact API, prefer its source in `build/packages/` over memory.
