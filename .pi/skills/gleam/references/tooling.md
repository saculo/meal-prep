# Tooling: CLI, gleam.toml, testing, publishing, language server

## Contents
1. Project layout
2. CLI reference
3. gleam.toml
4. Dependencies & manifest
5. Testing
6. Documentation & publishing
7. Language server (LSP): features, code actions, editor setup
8. CI

---

## 1. Project layout

`gleam new my_app` (flags: `--template javascript`, `--skip-git`, `--skip-github`, `--name`) creates:
```
my_app/
├── gleam.toml
├── manifest.toml            # generated lock file — commit it for applications
├── src/my_app.gleam         # entry module: `pub fn main()` is what `gleam run` executes
├── src/my_app/…             # all other modules namespaced under the package name
├── test/my_app_test.gleam   # gleeunit entry: `pub fn main() { gleeunit.main() }`
├── dev/                     # optional: dev scripts; `gleam dev` runs dev/my_app_dev.gleam main
├── priv/                    # optional: static files shipped with the app (Erlang)
└── build/                   # compiled output + downloaded deps (build/packages) — gitignored
```

## 2. CLI reference

| Command | Purpose |
|---|---|
| `gleam run [-m module] [--target js] [--runtime node\|deno\|bun] [-- args]` | Compile and run `main` |
| `gleam check [--target]` | Type-check only (fastest feedback) |
| `gleam build [--target] [--warnings-as-errors] [--no-print-progress]` | Compile |
| `gleam test [--target] [--runtime]` | Run tests (the `test/<pkg>_test.gleam` main) |
| `gleam format [paths] [--check] [--stdin]` | Format (canonical style) |
| `gleam fix` | Apply automatic migrations for deprecated syntax |
| `gleam dev` | Run `dev/<pkg>_dev.gleam` main |
| `gleam shell` | Erlang REPL with project loaded |
| `gleam add pkg [pkg@2] [--dev]` / `gleam remove pkg` | Add/remove deps (updates gleam.toml + manifest) |
| `gleam update [pkg]` / `gleam deps update` | Upgrade within constraints |
| `gleam deps list` / `deps tree [--package p \| --invert p]` / `deps outdated` / `deps download` | Inspect deps |
| `gleam clean` | Delete `build/` |
| `gleam docs build [--open]` / `docs publish` / `docs remove` | HTML docs |
| `gleam publish [--replace]` | Publish to Hex |
| `gleam hex authenticate` / `hex retire` / `hex unretire` / `hex revert` / `hex owner add\|transfer` | Hex management |
| `gleam export escript` | Single-file BEAM executable |
| `gleam export erlang-shipment` | Deployable precompiled Erlang release dir |
| `gleam export hex-tarball` / `package-interface --out f.json` / `package-information` / `javascript-prelude` / `typescript-prelude` | Misc exports; `package-interface` gives a JSON of all public types & signatures (useful for tooling/API inspection) |
| `gleam lsp` | Start language server (stdio) |

Arguments after the module go to the program: `gleam run -- get USER`; read them with the `argv` package.

## 3. gleam.toml

```toml
name = "my_app"                       # required, snake_case, = Hex name
version = "1.0.0"                     # required
description = "…"                     # required for publishing
licences = ["Apache-2.0"]             # SPDX; required for publishing
repository = { type = "github", user = "me", repo = "my_app" }  # also gitlab, codeberg, sourcehut, bitbucket, tangled, forgejo/gitea (+host), custom (+url)
links = [{ title = "Website", href = "https://example.com" }]
gleam = ">= 1.14.0"                   # required compiler version
target = "erlang"                     # or "javascript"
internal_modules = ["my_app/internal", "my_app/internal/*"]  # default shown

[dependencies]
gleam_stdlib = ">= 1.0.0 and < 2.0.0"
local_lib = { path = "../local_lib" }
git_lib = { git = "https://github.com/me/repo", ref = "a8b3c5d82", path = "packages/lib" }  # pin a commit SHA

[dev_dependencies]
gleeunit = ">= 1.0.0 and < 2.0.0"

[documentation]
pages = [{ title = "Guide", path = "guide.html", source = "./docs/guide.md" }]

[erlang]
application_start_module = "my_app@application"   # OTP application callback module (apps only)
extra_applications = ["inets", "ssl"]

[javascript]
runtime = "node"                      # node | deno | bun
typescript_declarations = true
source_maps = true
[javascript.deno]
allow_net = ["example.com:443"]      # also allow_all/env/ffi/hrtime/read/run/sys/write

[tools.my_tool]                       # config for third-party tools lives here
```

## 4. Dependencies

- Packages come from Hex (https://packages.gleam.run to search, https://hex.pm). Gleam, Erlang (rebar3) and Elixir (mix) packages all work on the Erlang target.
- Version requirement syntax is Hex's: `">= 1.2.0 and < 2.0.0"`, `"~> 1.2"`. `gleam add` writes a sensible range.
- `manifest.toml` pins exact versions; commit it for applications.
- Read a dependency's real API in `build/packages/<name>/src/` or `gleam export package-interface`, and its docs at `https://hexdocs.pm/<name>/`.

## 5. Testing

gleeunit (default): every `pub fn *_test()` in `test/` runs; a panic = failure.
```gleam
import gleeunit
import my_app/cart

pub fn main() -> Nil {
  gleeunit.main()
}

pub fn total_of_empty_cart_is_zero_test() {
  assert cart.total(cart.new()) == 0
}

pub fn parse_rejects_negative_quantity_test() {
  assert cart.parse_quantity("-1") == Error(cart.NegativeQuantity(-1))
}

pub fn add_item_test() {
  let assert Ok(cart) = cart.new() |> cart.add("apple", quantity: 2)
  assert cart.item_count(cart) == 2
}
```
- Use `assert` (shows both sides on failure) rather than the older `gleeunit/should` helpers (`should.equal`), which still work but are no longer the recommended style.
- Test internal helpers by putting them in `my_app/internal/*` modules (public but not public API).
- Test both targets for cross-target libraries.
- Other tools: `birdie` (snapshot tests, `birdie.snap(content, title: "…")`, review with `gleam run -m birdie`), `qcheck` (property-based), `startest` (alternative runner).

## 6. Documentation & publishing

- `///` on public definitions and `////` at module top become HexDocs; include examples in fenced ```gleam blocks.
- `gleam docs build --open` to preview.
- Before `gleam publish`: fill `description`, `licences`, `repository`; version bump; no `echo`/`todo`; modules namespaced; semver respected (breaking change → major). `gleam publish` refuses modules with no public definitions.
- `gleam hex authenticate` once; publishing uploads source + docs.

## 7. Language server

Bundled in the `gleam` binary: `gleam lsp` (stdio). It uses the `target` from `gleam.toml` (default Erlang). Works only inside Gleam projects (single files get formatting only). It type-checks unsaved buffers but never runs code generation or compiles Erlang/Elixir.

Editor setup:
- **VS Code**: official "Gleam" extension; `gleam` must be on PATH.
- **Neovim**: `nvim-lspconfig`, then `vim.lsp.enable('gleam')` (0.11+) or `require('lspconfig').gleam.setup({})`; `:TSInstall gleam` for highlighting.
- **Helix, Zed, Gram**: built in / auto-installed extension.
- **Others**: run `gleam lsp` from the project root.
- **Coding agents**: if an LSP tool is available in the environment, configure it to launch `gleam lsp` for `*.gleam` files; otherwise rely on `gleam check`, which reports the same diagnostics.

Features: diagnostics, formatting, hover (types + docs), go to definition / type definition, find references, rename (variables, functions, types, constructors, record fields, labels, type variables, module aliases; renaming a module file updates imports), completion with auto-import and type-directed ranking, signature help, document symbols, document highlight, folding.

Code actions (use these instead of hand-writing boilerplate when working in an editor):

| Group | Actions |
|---|---|
| Types & imports | Add annotation(s) / annotate all top-level definitions · Add missing import · Remove unused imports · Qualify / unqualify · Add missing type parameter · Replace `_` with inferred type · Remove `opaque` from private type |
| Patterns | Add missing patterns · Pattern match on variable/argument/call result/discarded value · Collapse nested `case` · Merge case branches with identical bodies · Convert inexhaustive `let` to `case` · Fill unused fields · Remove redundant tuples · Remove unreachable clauses · Use label shorthand |
| Generation | Generate function (also across modules) · Generate type · **Generate dynamic decoder** for a custom type · **Generate to-JSON function** (gleam_json) · Create unknown module · Fill labels · Add omitted labels |
| Refactor | Extract variable · Extract constant · Extract function · Inline variable · Convert to/from pipe · Convert to/from `use` · Wrap in / remove block · Add/remove anonymous function wrapper · Expand function capture · Remove redundant record update · Discard unused result · Interpolate string · Convert int base · Toggle doc comment · Fix identifier case · Correct operator in guard · Remove `echo` |

## 8. CI

The generated `.github/workflows/test.yml` uses `erlef/setup-beam` with `otp-version`, `gleam-version` (and optionally `rebar3-version`), then runs:
```sh
gleam deps download
gleam test
gleam format --check src test
```
Add `gleam build --warnings-as-errors` and `gleam test --target javascript` (with `actions/setup-node`) for multi-target libraries.
