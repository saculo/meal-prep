# Conventions, patterns and anti-patterns

Source: https://gleam.run/documentation/conventions-patterns-and-anti-patterns/ (official). Use this as the review checklist for Gleam code.

## Contents
- Conventions (C1–C10)
- Patterns (P1–P6)
- Anti-patterns (A1–A10)
- Review checklist

---

## Conventions

**C1. Qualify imported functions and constants.** `string.to_graphemes(x)`, not `import gleam/string.{to_graphemes}`. Types and constructors may be imported unqualified when it reads better (`import gleam/option.{type Option, None, Some}`).

**C2. Annotate every module-level function** — all arguments and the return type, including private functions. (LSP code action "Annotate all top level definitions" does it.)

**C3. Fallible functions return `Result`.** Not `Option`, never panic. Use `Result(a, Nil)` if there's no useful error detail.
```gleam
pub fn first(list: List(a)) -> Result(a, Nil) {
  case list {
    [item, ..] -> Ok(item)
    _ -> Error(Nil)
  }
}
```

**C4. Singular module names**, every path segment: `app/user`, `app/payment/invoice` — not `app/users`.

**C5. Acronyms are words**: `Json`, `build_json`, `HttpRequest`, `user_id`. (`JSON` would become `j_s_o_n` on the BEAM.)

**C6. Conversion functions are `x_to_y`.** Drop the prefix when the module is named after the type: in `my_app/identifier`, `pub fn to_string(id: Identifier) -> String` (called as `identifier.to_string`). Name the format when there's a specific one (`date_to_rfc3339`), or a descriptive verb (`float.round`, not `float_to_int`). Never `into_`, `as_`, `of_`.

**C7. Fallible function names** describe the domain (`parse_json`, `enqueue`). Use the `try_` prefix only for a Result-handling variant of an existing function (`list.try_map`, `list.try_fold`). No FP jargon names (`monadic_bind`).

**C8. Use the core libraries** rather than re-implementing: `gleam_stdlib`, `gleam_time` (`Timestamp`, `Duration`), `gleam_json`, `gleam_http` (`Request`, `Response`), `gleam_erlang`, `gleam_otp`, `gleam_javascript`.

**C9. Tool configuration goes in `gleam.toml` under `[tools.<tool_name>]`**, not separate config files. Dynamic config uses env vars / CLI args.

**C10. Source directories**:
- `src/` — package code. Imports deps + `src/` only (not dev-deps, `test/`, `dev/`).
- `test/` — tests; may import anything.
- `dev/` — dev tooling, generators, scripts (`gleam dev` runs `dev/<pkg>_dev.gleam`'s `main`); may import anything.

## Patterns

**P1. Descriptive, domain-level errors.** One variant per business failure, carrying context; wrap lower-level errors as fields.
```gleam
pub type NoteBookError {
  NoteAlreadyExists(path: String)
  NoteCouldNotBeCreated(path: String, reason: simplifile.FileError)
  NoteCouldNotBeRead(path: String, reason: simplifile.FileError)
  NoteInvalidFrontmatter(path: String, reason: tom.ParseError)
}
```
Bad: variants without context (`NoteCouldNotBeRead`) or modelled on dependencies (`FileError(..)`, `TomlError(..)`).

**P2. Comment liberally** — explain what and *why*, even for clear code.

**P3. Make invalid states impossible** with custom types.
```gleam
// Bad: Visitor(id: Some(1), email: None) is constructible
pub type Visitor { Visitor(id: Option(Int), email: Option(String)) }
// Good
pub type Visitor {
  LoggedInUser(id: Int, email: String)
  Guest
}
```
Combine with opaque types + smart constructors (`pub fn new(..) -> Result(T, E)`) for validated values.

**P4. Replace Bools with custom types.** `SchoolPerson(name: String, role: Role)` with `pub type Role { Student Teacher }` instead of `is_student: Bool`. Gives meaning, prevents mixing booleans, easy to add a third state, avoids combinatorial bool fields. Same for function arguments: prefer a variant over a bare `True`.

**P5. Sans-IO for API clients/SDKs.** Provide pairs of pure functions — build a `Request`, parse a `Response` — and let users send it with any HTTP client, on any target, with their own retries/rate limiting.
```gleam
pub fn create_user_request(name: String) -> Request(String) {
  request.new()
  |> request.set_method(http.Post)
  |> request.set_host("example.com")
  |> request.set_path("/users")
  |> request.set_body(json.to_string(json.object([#("name", json.string(name))])))
  |> request.prepend_header("content-type", "application/json")
}

pub fn create_user_response(response: Response(String)) -> Result(User, ApiError) {
  case response.status {
    201 -> Ok(User(name: response.body))
    409 -> Error(UserNameAlreadyInUse)
    429 -> Error(RateLimitWasHit)
    code -> Error(GotUnexpectedResponse(code, response.body))
  }
}
```
Don't take a "send" function as an argument instead — that ties users to one client shape and breaks multi-target support.

**P6. Builder pattern** for records with many optional fields: `new` takes required fields (labelled), chainable setters take the record first and return it.
```gleam
pub fn new(text text: String) -> Button {
  Button(text:, colour: "pink", classes: set.new())
}
pub fn colour(button: Button, value: String) -> Button {
  Button(..button, colour: value)
}
// button.new(text: "Continue") |> button.colour("green") |> button.to_html
```
Make the type opaque when invariants must be enforced.

## Anti-patterns

**A1. Abbreviations.** Write names in full: `capacity`, `offset`, `session`, not `cap`, `off`, `ss`.

**A2. Fragmented modules.** Don't split prematurely into `client`, `config`, `decode`, `error`, `types`… modules. Organise around the business domain and the best API for users; one well-designed module (`import my_library`) is often right. Large modules are fine. *AI-generated code frequently makes this mistake — pay extra attention.*

**A3. Panicking in libraries.** No `panic`, `let assert`, `assert`, `todo` in library code. Exception: OTP libraries designed to crash within supervision trees.

**A4. Global namespace pollution.** Put every module under `src/<package_name>/` (plus `src/<package_name>.gleam`). Top-level `src/inventory.gleam` in a package will collide with other packages.

**A5. Namespace trespassing.** Don't put modules in another package's namespace (e.g. `src/lustre/` in a Lustre add-on) unless you maintain that package.

**A6. Grouping by design pattern / kind.** No `types`, `utils`, `constants`, `functions`, `models`, `controllers`, `services`, `monads` modules. Group by domain: `app/stock`, `app/billing`.

**A7. Check-then-assert.**
```gleam
// Bad
case result.is_ok(data) {
  True -> { let assert Ok(value) = data  process(value) }
  False -> data
}
// Good
case data { Ok(value) -> process(value)  Error(e) -> Error(e) }
data |> result.try(process)
use value <- result.try(data)
```

**A8. `Dynamic` in FFI signatures.** Define a specific external type (`pub type Buffer`) instead of accepting `Dynamic`. `Dynamic` is for untyped data to be decoded.

**A9. Catch-all patterns over custom types.** Match each variant (`Student -> …  Teacher -> …`) so the compiler catches new variants. `_` is fine for open-ended values (ints, strings).

**A10. Category-theory abstractions.** No `Monoid`, `Functor`, `Catamorphism` records. Solve the specific problem: `int.sum(costs)`.

## Review checklist

- [ ] `gleam format --check` clean; no compiler warnings
- [ ] Every top-level function annotated; names unabbreviated, snake_case, acronyms as words
- [ ] Functions/constants imported qualified
- [ ] Fallible functions return `Result` with domain error types carrying context
- [ ] No `panic`/`let assert`/`todo`/`echo` in library or production paths
- [ ] Exhaustive matches on custom types (no lazy `_`)
- [ ] Bools replaced by custom types where they carry domain meaning; invalid states unrepresentable
- [ ] Modules singular, domain-oriented, namespaced under the package, not fragmented
- [ ] No check-then-assert; Results combined with `result.try`/`use`/`case`
- [ ] Core libraries used for time, JSON, HTTP types
- [ ] Tests for public behaviour (`gleam test`, both targets if multi-target)
- [ ] Public items have `///` docs; module has `////` doc if it's public API
