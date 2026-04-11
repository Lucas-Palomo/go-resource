# v2 Architecture Notes

[English](./architecture.md) | [Português (Brasil)](./architecture.pt_br.md)

## Architectural goal

v2 treats `go-resource` as a real reusable library, not only as a file-reading helper.

The design goals are direct:

- separate responsibilities
- reduce coupling
- remove `panic` from the main error path
- leave room for future evolution without deforming the public API

## Responsibility split

### `bundle.go`

Public runtime state, locale state, lookup behavior, fallback resolution, and convenience helpers.

### `loader.go`

Filesystem walking, `fs.FS` support, resource file discovery, file-name parsing, and flattening of nested structures.

### `decoder.go`

Decoder contract and built-in implementations for JSON, YAML, TOML, and Java-style `.properties`.

### `options.go`

Declarative configuration for bundle behavior.

### `errors.go`

Public error values and structured error types.

## Why `LoadDir` and `LoadFS` both exist

v1 was tied to the operating system filesystem. That was too narrow.

v2 explicitly supports `fs.FS`, which opens room for:

- `embed.FS`
- cleaner tests
- easier composition with modern Go libraries

## Why nested objects are flattened

Real i18n catalogs grow. Deep structures are easier to organize than massive flat maps, but runtime lookup still benefits from predictable flat keys.

Example:

```json
{
  "checkout": {
    "button": {
      "confirm": "Confirm"
    }
  }
}
```

becomes:

```text
checkout.button.confirm
```

## Why namespacing is hybrid

v2 accepts namespace information from:

- directory structure
- additional segments in the file name

That supports both common organization styles without forcing only one of them.

Example:

```text
resources/errors/en.json
```

with:

```json
{
  "validation": {
    "required": "Required"
  }
}
```

becomes:

```text
errors.validation.required
```

## Why failure policies are explicit

A reusable library should not silently choose critical behavior that users may want to control.

v2 makes two policies explicit:

- how to handle missing keys
- how to handle duplicate keys

## Why `Get` still exists

`Get` remains for ergonomics and continuity with v1.

`Lookup` is the more correct API for non-trivial scenarios because it returns `error` and makes failure handling explicit.

## Concurrency model

The bundle uses `sync.RWMutex` to protect internal state. The intent is safe concurrent reads with controlled runtime mutation.

## Extra runtime helpers

The following helpers were added to reduce friction in tests, diagnostics, and controlled reload flows:

- `Has`
- `HasFor`
- `Reset`

## Future evolution already prepared

The current structure leaves room for future additions such as:

- pluralization
- named placeholders
- more sophisticated locale matching
- incremental loading
- optional hot reload
- missing-key metrics
- integration with external backends
