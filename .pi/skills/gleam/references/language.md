# Gleam language reference

## Contents
1. Basic types & literals
2. Bindings, blocks, expressions
3. Functions, labels, captures, pipes
4. Custom types, records, opaque types, aliases, generics
5. Pattern matching (`case`, guards, list/string/bit-array patterns)
6. `use`
7. Result/Option handling idioms
8. Assertions, crashes, debugging (`assert`, `let assert`, `panic`, `todo`, `echo`)
9. Modules & imports
10. Constants, attributes, documentation comments
11. Bit arrays
12. Things Gleam does not have

---

## 1. Basic types & literals

| Type | Literals / notes |
|---|---|
| `Int` | `42`, `-7`, `1_000_000`, `0b1010`, `0o17`, `0xFF`. Arbitrary precision on BEAM; JS numbers (safe to ±2^53) on JS. |
| `Float` | `1.0`, `-0.5`, `1.5e3`, `1.0e-10`. Always needs `.` |
| `String` | `"text"`, UTF-8. Escapes: `\"` `\\` `\n` `\r` `\t` `\f` `\u{1F600}`. Multiline literals allowed. Length is in graphemes. |
| `Bool` | `True`, `False` (it is a custom type) |
| `Nil` | `Nil` — the unit value, returned by side-effecting functions |
| `List(a)` | `[1, 2, 3]`, `[0, ..rest]` (prepend), homogeneous |
| Tuple | `#(1, "a", 2.0)`, access `t.0`, `t.1` |
| `BitArray` | `<<1, 2, 3>>`, `<<"text":utf8>>` |
| `Result(a, e)` | `Ok(value)`, `Error(reason)` — in the prelude |
| Function | `fn(Int, String) -> Bool` |
| `UtfCodepoint` | via `string.utf_codepoint` |

Operators:
- Int: `+ - * / %` (`%` is remainder, sign follows dividend), comparisons `< > <= >=`.
- Float: `+. -. *. /.`, comparisons `<. >. <=. >=.`.
- Both division forms return 0 on division by zero.
- `==` / `!=` structural equality on any two values of the same type (no custom equality).
- `&&`, `||` (short-circuit), `!` (not), `-` (negate), `<>` (string concat).
- Precedence: use `{ }` for grouping, e.g. `{ a + b } * c`. Parentheses are not grouping syntax.

Ordering of non-Int values: `int.compare`, `float.compare`, `string.compare` return `order.Order` (`Lt | Eq | Gt`); `list.sort(xs, by: string.compare)`.

## 2. Bindings, blocks, expressions

```gleam
let x = 1
let x = x + 1            // shadowing, not mutation
let _unused = compute()  // leading underscore: intentionally unused
let total: Int = 10      // optional annotation
let y = {                // block: value of last expression
  let a = 2
  a * a
}
let #(a, b) = #(1, 2)    // destructuring must be exhaustive…
let assert [first, ..] = items  // …or use `let assert` (crashes if no match)
```

Every construct is an expression; `case` and blocks return values. There is no `return`; a function returns its last expression.

## 3. Functions, labels, captures, pipes

```gleam
pub fn add(a: Int, b: Int) -> Int { a + b }   // public
fn helper(x: Int) -> Int { x * 2 }            // private to module

let double = fn(x) { x * 2 }                  // anonymous (closure)
let add_one = add(1, _)                       // function capture: fn(b) { add(1, b) }
```

**Labelled arguments** (external name, internal name):
```gleam
pub fn replace(in string: String, each pattern: String, with replacement: String) -> String
replace(in: "a-b", each: "-", with: "+")
replace("a-b", each: "-", with: "+")   // positional first, then labelled in any order
```
Labels give call-site readability; use them for functions with several same-typed parameters. Shorthand when a variable has the label's name: `replace(in:, each:, with:)`.

**Pipes** pass the left value as the first argument:
```gleam
"  Hello  "
|> string.trim
|> string.lowercase
|> string.append("prefix: ", _)   // capture to target another argument
|> echo                           // debug-print mid-pipe
```

Higher-order functions are the main abstraction mechanism: pass `fn` values or records of functions instead of interfaces/type classes.

Recursion: self tail calls are optimised on both targets. Write accumulator-style helpers:
```gleam
fn sum_loop(xs: List(Int), acc: Int) -> Int {
  case xs {
    [] -> acc
    [first, ..rest] -> sum_loop(rest, acc + first)
  }
}
```

## 4. Custom types

```gleam
pub type Shape {
  Circle(radius: Float)
  Rectangle(width: Float, height: Float)
}

pub type Direction { North East South West }     // enum-like

pub type Tree(a) {                               // generic, recursive
  Leaf
  Node(left: Tree(a), value: a, right: Tree(a))
}
```

- Construct: `Circle(1.0)`, `Circle(radius: 1.0)`, `Rectangle(width:, height:)` (shorthand).
- Field access `shape.width` only when the field is shared by all variants with the same label, position and type; otherwise `case`.
- Record update: `Rectangle(..rect, width: 2.0)`. Works in constants too (`const dev = Config(..base, port: 4000)`).
- **Opaque types** hide constructors outside the module — use for invariants / smart constructors:
  ```gleam
  pub opaque type Email { Email(String) }
  pub fn parse(input: String) -> Result(Email, Nil) { … }
  pub fn to_string(email: Email) -> String { let Email(s) = email  s }
  ```
  `opaque` only makes sense on `pub` types (private opaque is an error).
- **Type aliases**: `pub type Headers = List(#(String, String))`.
- **External / phantom types**: `pub type Buffer` (no constructors) for FFI; unused type params act as phantom types for extra safety.
- Type variables are lowercase names (`a`, `value`, `msg`). No higher-kinded types, no type classes, no overloading, no subtyping.

## 5. Pattern matching

```gleam
case value {
  0 -> "zero"
  1 | 2 | 3 -> "small"                    // alternatives
  n if n < 0 -> "negative"               // guard
  _ -> "big"
}
```

- **Multiple subjects**: `case x % 3, x % 5 { 0, 0 -> "FizzBuzz"  0, _ -> "Fizz"  _, 0 -> "Buzz"  _, _ -> int.to_string(x) }`.
- **Lists**: `[]`, `[x]`, `[x, y]`, `[first, ..rest]`, `[1, 2, ..]`.
- **Strings**: one literal prefix then a variable: `"Bearer " <> token -> …`; bind the prefix with `"http" as scheme <> rest`. Suffix/middle patterns are not supported.
- **Records**: `User(name:, ..)`, `Circle(radius: r)`, nested patterns fine.
- **Aliases**: `[_, ..] as non_empty`, `Some(x) as opt`.
- **Tuples**: `#(Ok(a), Ok(b)) -> …` — handy for matching several Results together.
- **Guards** allow comparisons, boolean ops, arithmetic, literal values, record field and tuple index access, and constants — but not function calls. Wrap in `{ }` to clarify precedence.
- Exhaustiveness is checked; unreachable branches warn. Avoid catch-all `_` over custom types so adding a variant forces updates.
- `let` patterns must be irrefutable; otherwise use `case` or `let assert`.

## 6. `use`

`use` turns the rest of the block into an anonymous function passed as the **last** argument:

```gleam
use a, b <- f(x)
rest
// ≡ f(x, fn(a, b) { rest })
```

Common forms:
```gleam
use user <- result.try(load_user(id))                 // short-circuit on Error
use <- bool.guard(when: amount <= 0, return: Error(InvalidAmount))
use <- bool.lazy_guard(when: cached, return: fn() { cache_hit() })
use item <- list.each(items)                          // callback style
use name <- decode.field("name", decode.string)       // decoders
use value <- result.map(parse(input))                 // final step with map
```

Guidelines: use it to flatten callback nesting, especially `result.try` chains. Don't overuse it for non-callback APIs; a pipe or `case` is often clearer. It only affects the enclosing block, so wrap in `{ }` to limit scope.

## 7. Result / Option idioms

```gleam
// Transform & chain
parse(input) |> result.map(normalise) |> result.try(validate)
// Convert error types to your domain error
int.parse(s) |> result.replace_error(InvalidNumber(s))
simplifile.read(path) |> result.map_error(CouldNotRead(path, _))
// Collect
list.try_map(inputs, parse)          // first Error wins → Result(List(b), e)
result.all([a, b, c])                // Result(List(a), e)
result.partition(results)            // #(oks, errors)
list.filter_map(inputs, parse)       // keep only Ok values
// Defaults
result.unwrap(r, 0)   option.unwrap(o, "none")   result.lazy_unwrap(r, fn() { expensive() })
// Option ↔ Result
option.to_result(o, MissingField)    option.from_result(r)
```

Avoid check-then-assert (`result.is_ok(r)` then `let assert Ok(v) = r`) — pattern-match or combine instead.

## 8. Assertions, crashes, debugging

| Construct | Use |
|---|---|
| `assert expr` | Boolean assertion; crashes with a detailed report showing operands. Primary test assertion: `assert add(1, 2) == 3`. Optional `assert x as "message"`. |
| `let assert Pattern = value as "msg"` | Crash unless it matches. Fine in tests, app startup, supervised processes; not in libraries. |
| `panic` / `panic as "reason"` | Unreachable/impossible states in applications. |
| `todo` / `todo as "note"` | Placeholder that type-checks as any type (also in constants); compiler warns. |
| `echo value` / `echo value as "label"` | Debug print to stderr with location, returns value. Remove before committing; don't publish with it. |

## 9. Modules & imports

- One file = one module; path under `src/` is the module name: `src/my_app/user.gleam` → `my_app/user`.
- Only `pub` items are exported. Modules matching `internal_modules` globs (default `<pkg>/internal`, `<pkg>/internal/*`) are hidden from docs and not public API.
- ```gleam
  import gleam/list                                  // use as list.map
  import gleam/option.{type Option, None, Some}     // unqualified types/constructors OK
  import my_app/user.{type User, User} as user_mod  // alias
  import gleam/dynamic/decode                        // last segment is the qualifier
  ```
- No circular imports. Module names are global on the BEAM → put all modules under a directory named after the package.
- `src/` cannot import from `test/`, `dev/` or dev-dependencies; `test/` and `dev/` can import everything.

## 10. Constants, attributes, doc comments

```gleam
pub const max_retries = 3
pub const default_config = Config(host: "0.0.0.0", port: 8080)
const greeting = "Hello, " <> "world"
```
Constants take literals, constructors, other constants, record updates, lists/tuples and string concat — no function calls. `const` is module-level only.

Attributes:
- `@external(erlang, "module", "function")` / `@external(javascript, "./file.mjs", "fn")` — FFI (see `ffi-targets.md`).
- `@deprecated("Use x instead")` — warns at call sites.
- `@internal` — public but hidden from docs/autocomplete for other packages.

Comments: `//` regular, `///` documents the next definition, `////` documents the module (top of file). Doc comments are Markdown and appear in `gleam docs build` output and LSP hovers. Don't put a regular comment between a doc comment and its definition (compiler warns).

## 11. Bit arrays

```gleam
<<3:size(8), 1024:size(16)-big, "hi":utf8, payload:bits>>
case packet {
  <<version:4, length:12, body:bytes-size(length), rest:bits>> -> …
  _ -> …
}
```
Segment options: `int`, `float`, `bits`, `bytes`, `utf8`, `utf16`, `utf32`, `utf8_codepoint`, `signed`, `unsigned`, `big`, `little`, `native`, `size(n)`, `unit(n)`. Default segment is an 8-bit int. `bytes` = byte-aligned; `bits` = any bit size. The JavaScript target supports most options but bit arrays must end up byte-aligned in some operations — test on JS if you target it. Use `gleam/bit_array` for conversion (`to_string`, `base64_encode`, `slice`).

## 12. Not in Gleam (and what to use)

| Missing | Use instead |
|---|---|
| `if/else` | `case` on `Bool`, `bool.guard` |
| loops | recursion, `list.fold/map/each`, `int.range` |
| exceptions / try-catch | `Result` + `result.try`/`use`; supervisors on BEAM |
| null | `Option(a)` or specific variants |
| mutation | shadowing, returning new values, actors for shared state |
| type classes / interfaces | pass functions, records of functions |
| string interpolation | `<>`, `string.concat`, `string.join` |
| macros / reflection / metaprogramming | code generation in `dev/`, LSP code actions |
| implicit number conversion | `int.to_float`, `float.round`, `float.truncate` |
| index access on lists | pattern matching, `dict`, or `list.drop` + `list.first` |
| method call syntax | module functions + pipes |
