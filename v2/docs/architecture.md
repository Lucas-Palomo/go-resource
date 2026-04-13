# v2 Architecture Notes

[English](./architecture.md) | [Português (Brasil)](./architecture.pt_br.md)

## Design goal

The architecture of v2 is deliberately modest: keep file-based i18n simple, but make the runtime behavior explicit enough for reuse in real applications.

The library is built around one central idea: load resource files into a normalized in-memory catalog, then resolve keys through a locale fallback chain.

## Main runtime pieces

- `Bundle` owns catalogs, decoder registry, locale settings, and lookup strategies
- decoders transform raw files into `map[string]any`
- the loader normalizes decoded content into `map[string]string`
- lookup walks locale fallbacks built from `golang.org/x/text/language`

## Loading pipeline

For each resource file, the loader performs these steps:

1. select the decoder from the file extension
2. parse the locale from the filename
3. derive the namespace from directories and filename segments
4. decode the file into a generic object graph
5. flatten nested objects into dot-notated keys
6. normalize scalar values into strings
7. validate duplicate keys according to the configured strategy
8. merge the pending catalogs into bundle memory

## Key model

A final lookup key may come from three layers:

- folders
- filename namespace segments
- nested objects in the file body

Examples:

```text
resources/errors/en.json            -> errors.*
resources/en.messages.checkout.toml -> messages.checkout.*
```

That model keeps keys predictable even when projects mix directory-based organization and document-based nesting.

## Scalar normalization

The runtime catalog is always `map[string]string`.

This is intentional:

- lookups remain deterministic
- formatting with `fmt.Sprintf` is straightforward
- callers do not need to reason about mixed scalar types at read time

Numbers and booleans are stringified. `nil` values and unsupported composite leaf values are rejected.

## Locale resolution

Lookup follows a locale chain instead of a single exact match.

In practice, resolution uses:

1. the requested locale
2. its parent locales
3. the configured fallback locale
4. the fallback parents
5. `language.Und`

This makes it possible to keep a precise locale for overrides while still falling back to broader catalogs.

## Concurrency model

`Bundle` uses a read-write mutex.

- read paths such as `Lookup`, `Get`, `Has`, `Catalog`, and `Locales` use read locks
- mutation paths such as `LoadFS`, `SetLocale`, `Reset`, and `RegisterDecoder` use write locks

The public API is safe for concurrent reads after loading.

## Merge semantics

`LoadDir` and `LoadFS` are incremental operations.

A successful second load does not replace the current in-memory bundle. It merges additional catalogs into what is already loaded.

That behavior is useful for layered sources, but it changes how reload flows should be designed:

- call `Reset()` before loading again when you want full replacement
- under `ErrorOnDuplicate`, duplicate validation against already-loaded catalogs completes before any mutation happens in the new load

## Decoder extension points

`RegisterDecoder` allows support for additional file formats.

Contract details:

- extensions are normalized to dotted lowercase form
- empty extensions are ignored
- nil decoders are ignored
- unsupported extensions fail loading with `ErrUnsupportedFormat`

## Lookup contract

The API intentionally exposes two lookup styles:

- `Lookup` and `LookupFor` are explicit and error-aware
- `Get` and `GetFor` are lenient helpers that return the original key on lookup errors

That split gives ergonomic access for UI text while preserving a strict path for services, validation, and tests.
