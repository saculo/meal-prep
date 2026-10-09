# gleam_stdlib 1.x reference

Public signatures extracted from `gleam_stdlib` v1.0.5 source (deprecated functions omitted). Labels shown are the call-site labels (`fn get(from: Dict(k, v), get: k)` → `dict.get(d, key)`). For other versions, check `build/packages/gleam_stdlib/src/gleam/*.gleam` or https://hexdocs.pm/gleam_stdlib/.

## Contents
- Notes & key types
- Decoding dynamic data (`gleam/dynamic/decode`)
- Signatures: bit_array, bool, bytes_tree, dict, dynamic/decode, dynamic, float, function, int, io, list, option, order, pair, result, set, string, string_tree, uri

## Notes & key types

- Prelude (no import): `Int Float String Bool Nil List Result(Ok, Error) BitArray UtfCodepoint`.
- `pub type Option(a) { Some(a) None }` · `pub type Order { Lt Eq Gt }` · `pub type ContinueOrStop(a) { Continue(a) Stop(a) }` (for `list.fold_until`) · `pub type DecodeError { DecodeError(expected: String, found: String, path: List(String)) }` · `pub type Uri { Uri(scheme: Option(String), userinfo: Option(String), host: Option(String), port: Option(Int), path: String, query: Option(String), fragment: Option(String)) }`.
- `Dict` is an unordered immutable hash map (Erlang map / JS HAMT); `Set` is built on it. Neither preserves insertion order.
- `string.length`, `string.slice`, `string.reverse` work on grapheme clusters (O(n)). `string.byte_size` is bytes.
- Use `string_tree` / `bytes_tree` for efficient repeated concatenation (e.g. building output), then `to_string` / `to_bit_array` once.
- `list.sort` is stable merge sort; `list.unique` is O(n²)-ish — use `set` for large inputs.
- `int.random(max)` → 0..max-1; `float.random()` → [0, 1).
- `io.println` writes to stdout, `io.println_error` to stderr. For debugging use `echo`.
- Modules no longer in stdlib: iterator (`gleam_yielder`), regex (`gleam_regexp`), queue (`gleam_deque`), string_builder (→ `string_tree`), bytes_builder (→ `bytes_tree`), `dynamic.field`/`dynamic.decode*` (→ `gleam/dynamic/decode`), `io.debug` (→ `echo`).

## Decoding dynamic data

`Dynamic` values come from JSON (`json.parse`), FFI, Erlang messages, databases. Build a `Decoder(t)` and run it:

```gleam
import gleam/dynamic/decode
import gleam/option.{type Option, None}

pub type Pet {
  Pet(name: String, age: Int, tags: List(String), owner: Option(String))
}

pub fn pet_decoder() -> decode.Decoder(Pet) {
  use name <- decode.field("name", decode.string)
  use age <- decode.field("age", decode.int)
  use tags <- decode.optional_field("tags", [], decode.list(decode.string))
  use owner <- decode.optional_field("owner", None, decode.optional(decode.string))
  decode.success(Pet(name:, age:, tags:, owner:))
}

// json.parse(text, pet_decoder())  or  decode.run(dynamic_value, pet_decoder())
```
- `decode.field` requires the key; `decode.optional_field(key, default, d)` allows it to be missing; `decode.optional(d)` accepts `null`/`nil` → `Option`.
- Nested paths: `decode.subfield(["a", "b"], decode.int, next)` or `decode.at(["a", "b"], decode.int)`. Tuple/list indexes: use ints as path segments.
- Custom types from tags: `use tag <- decode.field("type", decode.string)` then `case tag { "cat" -> cat_decoder() … _ -> decode.failure(default, "Pet") }`; or `decode.then(decode.string, fn(tag) { … })`.
- Alternatives: `decode.one_of(first, or: [second, third])`. Transform: `decode.map(d, f)`. Recursive types: `decode.recursive(fn() { tree_decoder() })`.
- Primitive decoders are constants: `decode.string`, `decode.int`, `decode.float`, `decode.bool`, `decode.bit_array`, `decode.dynamic`. Collections: `decode.list(of:)`, `decode.dict(key, value)`.
- The LSP code action "Generate dynamic decoder" writes decoders for a custom type.

## Signatures

### gleam/bit_array
```gleam
fn from_string(x: String) -> BitArray
fn bit_size(x: BitArray) -> Int
fn byte_size(x: BitArray) -> Int
fn pad_to_bytes(data: BitArray) -> BitArray
fn append(to first: BitArray, suffix second: BitArray) -> BitArray
fn slice(from string: BitArray, at position: Int, take length: Int) -> Result(BitArray, Nil)
fn is_utf8(bits: BitArray) -> Bool
fn to_string(bits: BitArray) -> Result(String, Nil)
fn concat(bit_arrays: List(BitArray)) -> BitArray
fn base64_encode(input: BitArray, padding: Bool) -> String
fn base64_decode(encoded: String) -> Result(BitArray, Nil)
fn base64_url_encode(input: BitArray, padding: Bool) -> String
fn base64_url_decode(encoded: String) -> Result(BitArray, Nil)
fn base16_encode(input: BitArray) -> String
fn base16_decode(input: String) -> Result(BitArray, Nil)
fn inspect(input: BitArray) -> String
fn compare(a: BitArray, with b: BitArray) -> order.Order
fn starts_with(bits: BitArray, prefix: BitArray) -> Bool
```

### gleam/bool
```gleam
fn and(a: Bool, b: Bool) -> Bool
fn or(a: Bool, b: Bool) -> Bool
fn negate(bool: Bool) -> Bool
fn nor(a: Bool, b: Bool) -> Bool
fn nand(a: Bool, b: Bool) -> Bool
fn exclusive_or(a: Bool, b: Bool) -> Bool
fn exclusive_nor(a: Bool, b: Bool) -> Bool
fn to_string(bool: Bool) -> String
fn guard(when requirement: Bool, return consequence: a, otherwise alternative: fn() -> a) -> a
fn lazy_guard(when requirement: Bool, return consequence: fn() -> a, otherwise alternative: fn() -> a) -> a
```

### gleam/bytes_tree
```gleam
opaque type BytesTree
fn new() -> BytesTree
fn prepend(to second: BytesTree, prefix first: BitArray) -> BytesTree
fn append(to first: BytesTree, suffix second: BitArray) -> BytesTree
fn prepend_tree(to second: BytesTree, prefix first: BytesTree) -> BytesTree
fn append_tree(to first: BytesTree, suffix second: BytesTree) -> BytesTree
fn prepend_string(to second: BytesTree, prefix first: String) -> BytesTree
fn append_string(to first: BytesTree, suffix second: String) -> BytesTree
fn concat(trees: List(BytesTree)) -> BytesTree
fn concat_bit_arrays(bits: List(BitArray)) -> BytesTree
fn from_string(string: String) -> BytesTree
fn from_string_tree(tree: string_tree.StringTree) -> BytesTree
fn from_bit_array(bits: BitArray) -> BytesTree
fn to_bit_array(tree: BytesTree) -> BitArray
fn byte_size(tree: BytesTree) -> Int
```

### gleam/dict
```gleam
type Dict(key, value)
fn size(dict: Dict(k, v)) -> Int
fn is_empty(dict: Dict(k, v)) -> Bool
fn to_list(dict: Dict(k, v)) -> List(#(k, v))
fn from_list(list: List(#(k, v))) -> Dict(k, v)
fn has_key(dict: Dict(k, v), key: k) -> Bool
fn new() -> Dict(k, v)
fn get(from: Dict(k, v), get: k) -> Result(v, Nil)
fn insert(into dict: Dict(k, v), for key: k, insert value: v) -> Dict(k, v)
fn map_values(in dict: Dict(k, v), with fun: fn(k, v) -> a) -> Dict(k, a)
fn keys(dict: Dict(k, v)) -> List(k)
fn values(dict: Dict(k, v)) -> List(v)
fn filter(in dict: Dict(k, v), keeping predicate: fn(k, v) -> Bool) -> Dict(k, v)
fn take(from dict: Dict(k, v), keeping desired_keys: List(k)) -> Dict(k, v)
fn merge(into dict: Dict(k, v), from new_entries: Dict(k, v)) -> Dict(k, v)
fn delete(from dict: Dict(k, v), delete key: k) -> Dict(k, v)
fn drop(from dict: Dict(k, v), drop disallowed_keys: List(k)) -> Dict(k, v)
fn upsert(in dict: Dict(k, v), update key: k, with fun: fn(Option(v)) -> v) -> Dict(k, v)
fn fold(over dict: Dict(k, v), from initial: acc, with fun: fn(acc, k, v) -> acc) -> acc
fn each(dict: Dict(k, v), fun: fn(k, v) -> a) -> Nil
fn combine(dict: Dict(k, v), other: Dict(k, v), with fun: fn(v, v) -> v) -> Dict(k, v)
fn group(key: fn(v) -> k, list: List(v)) -> Dict(k, List(v))
```

### gleam/dynamic/decode
```gleam
type DecodeError
opaque type Decoder(t)
fn subfield(field_path: List(name), field_decoder: Decoder(t), next: fn(t) -> Decoder(final)) -> Decoder(final)
fn run(data: Dynamic, decoder: Decoder(t)) -> Result(t, List(DecodeError))
fn at(path: List(segment), inner: Decoder(a)) -> Decoder(a)
fn success(data: t) -> Decoder(t)
fn decode_error(expected expected: String, found found: Dynamic) -> List(DecodeError)
fn field(field_name: name, field_decoder: Decoder(t), next: fn(t) -> Decoder(final)) -> Decoder(final)
fn optional_field(key: name, default: t, field_decoder: Decoder(t), next: fn(t) -> Decoder(final)) -> Decoder(final)
fn optionally_at(path: List(segment), default: a, inner: Decoder(a)) -> Decoder(a)
const string: Decoder(String)
const bool: Decoder(Bool)
const int: Decoder(Int)
const float: Decoder(Float)
const dynamic: Decoder(Dynamic)
const bit_array: Decoder(BitArray)
fn list(of inner: Decoder(a)) -> Decoder(List(a))
fn dict(key: Decoder(key), value: Decoder(value)) -> Decoder(Dict(key, value))
fn optional(inner: Decoder(a)) -> Decoder(Option(a))
fn map(decoder: Decoder(a), transformer: fn(a) -> b) -> Decoder(b)
fn map_errors(decoder: Decoder(a), transformer: fn(List(DecodeError)) -> List(DecodeError)) -> Decoder(a)
fn collapse_errors(decoder: Decoder(a), name: String) -> Decoder(a)
fn then(decoder: Decoder(a), next: fn(a) -> Decoder(b)) -> Decoder(b)
fn one_of(first: Decoder(a), or alternatives: List(Decoder(a))) -> Decoder(a)
fn failure(placeholder: a, expected name: String) -> Decoder(a)
fn new_primitive_decoder(name: String, decoding_function: fn(Dynamic) -> Result(t, t)) -> Decoder(t)
fn recursive(inner: fn() -> Decoder(a)) -> Decoder(a)
```

### gleam/dynamic
```gleam
type Dynamic
fn classify(data: Dynamic) -> String
fn bool(a: Bool) -> Dynamic
fn string(a: String) -> Dynamic
fn float(a: Float) -> Dynamic
fn int(a: Int) -> Dynamic
fn bit_array(a: BitArray) -> Dynamic
fn list(a: List(Dynamic)) -> Dynamic
fn array(a: List(Dynamic)) -> Dynamic
fn properties(entries: List(#(Dynamic, Dynamic))) -> Dynamic
fn nil() -> Dynamic
```

### gleam/float
```gleam
fn parse(string: String) -> Result(Float, Nil)
fn to_string(x: Float) -> String
fn clamp(x: Float, min min_bound: Float, max max_bound: Float) -> Float
fn compare(a: Float, with b: Float) -> Order
fn loosely_compare(a: Float, with b: Float, tolerating tolerance: Float) -> Order
fn loosely_equals(a: Float, with b: Float, tolerating tolerance: Float) -> Bool
fn min(a: Float, b: Float) -> Float
fn max(a: Float, b: Float) -> Float
fn ceiling(x: Float) -> Float
fn floor(x: Float) -> Float
fn round(x: Float) -> Int
fn truncate(x: Float) -> Int
fn to_precision(x: Float, precision: Int) -> Float
fn absolute_value(x: Float) -> Float
fn power(base: Float, of exponent: Float) -> Result(Float, Nil)
fn square_root(x: Float) -> Result(Float, Nil)
fn negate(x: Float) -> Float
fn sum(numbers: List(Float)) -> Float
fn product(numbers: List(Float)) -> Float
fn random() -> Float
fn modulo(dividend: Float, by divisor: Float) -> Result(Float, Nil)
fn divide(a: Float, by b: Float) -> Result(Float, Nil)
fn add(a: Float, b: Float) -> Float
fn multiply(a: Float, b: Float) -> Float
fn subtract(a: Float, b: Float) -> Float
fn logarithm(x: Float) -> Result(Float, Nil)
fn exponential(x: Float) -> Float
```

### gleam/function
```gleam
fn identity(x: a) -> a
```

### gleam/int
```gleam
fn absolute_value(x: Int) -> Int
fn power(base: Int, of exponent: Float) -> Result(Float, Nil)
fn square_root(x: Int) -> Result(Float, Nil)
fn parse(string: String) -> Result(Int, Nil)
fn base_parse(string: String, base: Int) -> Result(Int, Nil)
fn to_string(x: Int) -> String
fn to_base_string(x: Int, base: Int) -> Result(String, Nil)
fn to_base2(x: Int) -> String
fn to_base8(x: Int) -> String
fn to_base16(x: Int) -> String
fn to_base36(x: Int) -> String
fn to_float(x: Int) -> Float
fn clamp(x: Int, min min_bound: Int, max max_bound: Int) -> Int
fn compare(a: Int, with b: Int) -> Order
fn min(a: Int, b: Int) -> Int
fn max(a: Int, b: Int) -> Int
fn is_even(x: Int) -> Bool
fn is_odd(x: Int) -> Bool
fn negate(x: Int) -> Int
fn sum(numbers: List(Int)) -> Int
fn product(numbers: List(Int)) -> Int
fn random(max: Int) -> Int
fn divide(dividend: Int, by divisor: Int) -> Result(Int, Nil)
fn remainder(dividend: Int, by divisor: Int) -> Result(Int, Nil)
fn modulo(dividend: Int, by divisor: Int) -> Result(Int, Nil)
fn floor_divide(dividend: Int, by divisor: Int) -> Result(Int, Nil)
fn add(a: Int, b: Int) -> Int
fn multiply(a: Int, b: Int) -> Int
fn subtract(a: Int, b: Int) -> Int
fn bitwise_and(x: Int, y: Int) -> Int
fn bitwise_not(x: Int) -> Int
fn bitwise_or(x: Int, y: Int) -> Int
fn bitwise_exclusive_or(x: Int, y: Int) -> Int
fn bitwise_shift_left(x: Int, y: Int) -> Int
fn bitwise_shift_right(x: Int, y: Int) -> Int
fn range(from start: Int, to stop: Int, with acc: acc, run reducer: fn(acc, Int) -> acc) -> acc
```

### gleam/io
```gleam
fn print(string: String) -> Nil
fn print_error(string: String) -> Nil
fn println(string: String) -> Nil
fn println_error(string: String) -> Nil
```

### gleam/list
```gleam
fn length(of list: List(a)) -> Int
fn count(list: List(a), where predicate: fn(a) -> Bool) -> Int
fn reverse(list: List(a)) -> List(a)
fn is_empty(list: List(a)) -> Bool
fn contains(list: List(a), any elem: a) -> Bool
fn first(list: List(a)) -> Result(a, Nil)
fn rest(list: List(a)) -> Result(List(a), Nil)
fn group(list: List(v), by key: fn(v) -> k) -> Dict(k, List(v))
fn filter(list: List(a), keeping predicate: fn(a) -> Bool) -> List(a)
fn filter_map(list: List(a), with fun: fn(a) -> Result(b, e)) -> List(b)
fn map(list: List(a), with fun: fn(a) -> b) -> List(b)
fn map2(list1: List(a), list2: List(b), with fun: fn(a, b) -> c) -> List(c)
fn map_fold(over list: List(a), from initial: acc, with fun: fn(acc, a) -> #(acc, b)) -> #(acc, List(b))
fn index_map(list: List(a), with fun: fn(a, Int) -> b) -> List(b)
fn try_map(over list: List(a), with fun: fn(a) -> Result(b, e)) -> Result(List(b), e)
fn drop(from list: List(a), up_to n: Int) -> List(a)
fn take(from list: List(a), up_to n: Int) -> List(a)
fn new() -> List(a)
fn wrap(item: a) -> List(a)
fn append(first: List(a), second: List(a)) -> List(a)
fn prepend(to list: List(a), this item: a) -> List(a)
fn flatten(lists: List(List(a))) -> List(a)
fn flat_map(over list: List(a), with fun: fn(a) -> List(b)) -> List(b)
fn fold(over list: List(a), from initial: acc, with fun: fn(acc, a) -> acc) -> acc
fn fold_right(over list: List(a), from initial: acc, with fun: fn(acc, a) -> acc) -> acc
fn index_fold(over list: List(a), from initial: acc, with fun: fn(acc, a, Int) -> acc) -> acc
fn try_fold(over list: List(a), from initial: acc, with fun: fn(acc, a) -> Result(acc, e)) -> Result(acc, e)
type ContinueOrStop(a)
fn fold_until(over list: List(a), from initial: acc, with fun: fn(acc, a) -> ContinueOrStop(acc)) -> acc
fn find(in list: List(a), one_that is_desired: fn(a) -> Bool) -> Result(a, Nil)
fn find_map(in list: List(a), with fun: fn(a) -> Result(b, c)) -> Result(b, Nil)
fn all(in list: List(a), satisfying predicate: fn(a) -> Bool) -> Bool
fn any(in list: List(a), satisfying predicate: fn(a) -> Bool) -> Bool
fn zip(list: List(a), with other: List(b)) -> List(#(a, b))
fn strict_zip(list: List(a), with other: List(b)) -> Result(List(#(a, b)), Nil)
fn unzip(input: List(#(a, b))) -> #(List(a), List(b))
fn intersperse(list: List(a), with elem: a) -> List(a)
fn unique(list: List(a)) -> List(a)
fn sort(list: List(a), by compare: fn(a, a) -> Order) -> List(a)
fn repeat(item a: a, times times: Int) -> List(a)
fn split(list list: List(a), at index: Int) -> #(List(a), List(a))
fn split_while(list list: List(a), satisfying predicate: fn(a) -> Bool) -> #(List(a), List(a))
fn key_find(in keyword_list: List(#(k, v)), find desired_key: k) -> Result(v, Nil)
fn key_filter(in keyword_list: List(#(k, v)), find desired_key: k) -> List(v)
fn key_pop(list: List(#(k, v)), key: k) -> Result(#(v, List(#(k, v))), Nil)
fn key_set(list: List(#(k, v)), key: k, value: v) -> List(#(k, v))
fn each(list: List(a), f: fn(a) -> b) -> Nil
fn try_each(over list: List(a), with fun: fn(a) -> Result(b, e)) -> Result(Nil, e)
fn partition(list: List(a), with categorise: fn(a) -> Bool) -> #(List(a), List(a))
fn permutations(list: List(a)) -> List(List(a))
fn window(list: List(a), by n: Int) -> List(List(a))
fn window_by_2(list: List(a)) -> List(#(a, a))
fn drop_while(in list: List(a), satisfying predicate: fn(a) -> Bool) -> List(a)
fn take_while(in list: List(a), satisfying predicate: fn(a) -> Bool) -> List(a)
fn chunk(in list: List(a), by f: fn(a) -> k) -> List(List(a))
fn sized_chunk(in list: List(a), into count: Int) -> List(List(a))
fn reduce(over list: List(a), with fun: fn(a, a) -> a) -> Result(a, Nil)
fn scan(over list: List(a), from initial: acc, with fun: fn(acc, a) -> acc) -> List(acc)
fn last(list: List(a)) -> Result(a, Nil)
fn combinations(items: List(a), by n: Int) -> List(List(a))
fn combination_pairs(items: List(a)) -> List(#(a, a))
fn interleave(list: List(List(a))) -> List(a)
fn transpose(list_of_lists: List(List(a))) -> List(List(a))
fn shuffle(list: List(a)) -> List(a)
fn max(over list: List(a), with compare: fn(a, a) -> Order) -> Result(a, Nil)
fn sample(from list: List(a), up_to n: Int) -> List(a)
```

### gleam/option
```gleam
type Option(a)
fn all(list: List(Option(a))) -> Option(List(a))
fn is_some(option: Option(a)) -> Bool
fn is_none(option: Option(a)) -> Bool
fn to_result(option: Option(a), e) -> Result(a, e)
fn from_result(result: Result(a, e)) -> Option(a)
fn unwrap(option: Option(a), or default: a) -> a
fn lazy_unwrap(option: Option(a), or default: fn() -> a) -> a
fn map(over option: Option(a), with fun: fn(a) -> b) -> Option(b)
fn flatten(option: Option(Option(a))) -> Option(a)
fn then(option: Option(a), apply fun: fn(a) -> Option(b)) -> Option(b)
fn or(first: Option(a), second: Option(a)) -> Option(a)
fn lazy_or(first: Option(a), second: fn() -> Option(a)) -> Option(a)
fn values(options: List(Option(a))) -> List(a)
```

### gleam/order
```gleam
type Order
fn negate(order: Order) -> Order
fn to_int(order: Order) -> Int
fn compare(a: Order, with b: Order) -> Order
fn reverse(orderer: fn(a, a) -> Order) -> fn(a, a) -> Order
fn break_tie(in order: Order, with other: Order) -> Order
fn lazy_break_tie(in order: Order, with comparison: fn() -> Order) -> Order
```

### gleam/pair
```gleam
fn first(pair: #(a, b)) -> a
fn second(pair: #(a, b)) -> b
fn swap(pair: #(a, b)) -> #(b, a)
fn map_first(of pair: #(a, b), with fun: fn(a) -> c) -> #(c, b)
fn map_second(of pair: #(a, b), with fun: fn(b) -> c) -> #(a, c)
fn new(first: a, second: b) -> #(a, b)
```

### gleam/result
```gleam
fn is_ok(result: Result(a, e)) -> Bool
fn is_error(result: Result(a, e)) -> Bool
fn map(over result: Result(a, e), with fun: fn(a) -> b) -> Result(b, e)
fn map_error(over result: Result(a, e), with fun: fn(e) -> f) -> Result(a, f)
fn flatten(result: Result(Result(a, e), e)) -> Result(a, e)
fn try(result: Result(a, e), apply fun: fn(a) -> Result(b, e)) -> Result(b, e)
fn unwrap(result: Result(a, e), or default: a) -> a
fn lazy_unwrap(result: Result(a, e), or default: fn() -> a) -> a
fn unwrap_error(result: Result(a, e), or default: e) -> e
fn or(first: Result(a, e), second: Result(a, e)) -> Result(a, e)
fn lazy_or(first: Result(a, e), second: fn() -> Result(a, e)) -> Result(a, e)
fn all(results: List(Result(a, e))) -> Result(List(a), e)
fn partition(results: List(Result(a, e))) -> #(List(a), List(e))
fn replace(result: Result(a, e), value: b) -> Result(b, e)
fn replace_error(result: Result(a, e), error: f) -> Result(a, f)
fn values(results: List(Result(a, e))) -> List(a)
fn try_recover(result: Result(a, e), with fun: fn(e) -> Result(a, f)) -> Result(a, f)
```

### gleam/set
```gleam
opaque type Set(member)
fn new() -> Set(member)
fn size(set: Set(member)) -> Int
fn is_empty(set: Set(member)) -> Bool
fn insert(into set: Set(member), this member: member) -> Set(member)
fn contains(in set: Set(member), this member: member) -> Bool
fn delete(from set: Set(member), this member: member) -> Set(member)
fn to_list(set: Set(member)) -> List(member)
fn from_list(members: List(member)) -> Set(member)
fn fold(over set: Set(member), from initial: acc, with reducer: fn(acc, member) -> acc) -> acc
fn filter(in set: Set(member), keeping predicate: fn(member) -> Bool) -> Set(member)
fn map(set: Set(member), with fun: fn(member) -> mapped) -> Set(mapped)
fn drop(from set: Set(member), drop disallowed: List(member)) -> Set(member)
fn take(from set: Set(member), keeping desired: List(member)) -> Set(member)
fn union(of first: Set(member), and second: Set(member)) -> Set(member)
fn intersection(of first: Set(member), and second: Set(member)) -> Set(member)
fn difference(from first: Set(member), minus second: Set(member)) -> Set(member)
fn is_subset(first: Set(member), of second: Set(member)) -> Bool
fn is_disjoint(first: Set(member), from second: Set(member)) -> Bool
fn symmetric_difference(of first: Set(member), and second: Set(member)) -> Set(member)
fn each(set: Set(member), fun: fn(member) -> a) -> Nil
```

### gleam/string
```gleam
fn is_empty(str: String) -> Bool
fn length(string: String) -> Int
fn reverse(string: String) -> String
fn replace(in string: String, each pattern: String, with substitute: String) -> String
fn lowercase(string: String) -> String
fn uppercase(string: String) -> String
fn compare(a: String, b: String) -> order.Order
fn slice(from string: String, at_index idx: Int, length len: Int) -> String
fn crop(from string: String, before substring: String) -> String
fn drop_start(from string: String, up_to num_graphemes: Int) -> String
fn drop_end(from string: String, up_to num_graphemes: Int) -> String
fn contains(does haystack: String, contain needle: String) -> Bool
fn starts_with(string: String, prefix: String) -> Bool
fn ends_with(string: String, suffix: String) -> Bool
fn split(x: String, on substring: String) -> List(String)
fn split_once(string: String, on substring: String) -> Result(#(String, String), Nil)
fn append(to first: String, suffix second: String) -> String
fn concat(strings: List(String)) -> String
fn repeat(string: String, times times: Int) -> String
fn join(strings: List(String), with separator: String) -> String
fn pad_start(string: String, to desired_length: Int, with pad_string: String) -> String
fn pad_end(string: String, to desired_length: Int, with pad_string: String) -> String
fn trim(string: String) -> String
fn trim_start(string: String) -> String
fn trim_end(string: String) -> String
fn pop_grapheme(string: String) -> Result(#(String, String), Nil)
fn to_graphemes(string: String) -> List(String)
fn to_utf_codepoints(string: String) -> List(UtfCodepoint)
fn from_utf_codepoints(utf_codepoints: List(UtfCodepoint)) -> String
fn utf_codepoint(value: Int) -> Result(UtfCodepoint, Nil)
fn utf_codepoint_to_int(cp: UtfCodepoint) -> Int
fn to_option(string: String) -> Option(String)
fn first(string: String) -> Result(String, Nil)
fn last(string: String) -> Result(String, Nil)
fn capitalise(string: String) -> String
fn inspect(term: anything) -> String
fn byte_size(string: String) -> Int
fn remove_prefix(from string: String, matching prefix: String) -> String
fn remove_suffix(from string: String, matching suffix: String) -> String
```

### gleam/string_tree
```gleam
type StringTree
fn new() -> StringTree
fn prepend(to tree: StringTree, prefix prefix: String) -> StringTree
fn append(to tree: StringTree, suffix second: String) -> StringTree
fn prepend_tree(to tree: StringTree, prefix prefix: StringTree) -> StringTree
fn append_tree(to tree: StringTree, suffix suffix: StringTree) -> StringTree
fn from_strings(strings: List(String)) -> StringTree
fn concat(trees: List(StringTree)) -> StringTree
fn from_string(string: String) -> StringTree
fn to_string(tree: StringTree) -> String
fn byte_size(tree: StringTree) -> Int
fn join(trees: List(StringTree), with sep: String) -> StringTree
fn lowercase(tree: StringTree) -> StringTree
fn uppercase(tree: StringTree) -> StringTree
fn reverse(tree: StringTree) -> StringTree
fn split(tree: StringTree, on pattern: String) -> List(StringTree)
fn replace(in tree: StringTree, each pattern: String, with substitute: String) -> StringTree
fn is_equal(a: StringTree, b: StringTree) -> Bool
fn is_empty(tree: StringTree) -> Bool
```

### gleam/uri
```gleam
type Uri
const empty = Uri(scheme: None, userinfo: None, host: None, port: None, path: "", query: None, fragment: None)
fn parse(uri_string: String) -> Result(Uri, Nil)
fn parse_query(query: String) -> Result(List(#(String, String)), Nil)
fn query_to_string(query: List(#(String, String))) -> String
fn percent_encode(value: String) -> String
fn percent_decode(value: String) -> Result(String, Nil)
fn path_segments(path: String) -> List(String)
fn to_string(uri: Uri) -> String
fn origin(uri: Uri) -> Result(String, Nil)
fn merge(base: Uri, relative: Uri) -> Result(Uri, Nil)
```
