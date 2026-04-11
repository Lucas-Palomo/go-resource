# v2 Architecture Notes

[English](./architecture.md) | [Português (Brasil)](./architecture.pt_br.md)

## Architecture goal

v2 was designed to treat `go-resource` as a real library, not only as a file-reading helper. The target is simple:

- separate responsibilities
- reduce coupling
- remove `panic` from the main path
- prepare room for future extensions without deforming the API

## Separation of responsibilities

### `bundle.go`

Public state, lookup, fallback, and runtime operations.

### `loader.go`

Loading layer, directory / `fs.FS` walking, file-name parsing, and data flattening.

### `decoder.go`

Decoder contract and built-in implementations for JSON, YAML, and TOML.

### `options.go`

Declarative library configuration.

### `errors.go`

Public errors and structured error types.

## Decision: `LoadDir` and `LoadFS`

v1 was tied to the traditional filesystem. v2 accepts `fs.FS`, which opens space for:

- `embed.FS`
- simpler tests
- better composition with modern Go libraries

## Decision: flattening into dot notation

In real i18n projects, files grow. Flat maps quickly become hard to maintain. v2 converts nested objects into predictable flat keys.

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

## Decision: hybrid namespacing

v2 accepts namespaces from:

- directory structure
- extra segments in the file name

That supports two organization styles without forcing only one.

## Decision: explicit failure policies

A mature library should not hide important choices. v2 makes two behaviors explicit:

- how to handle missing keys
- how to handle duplicate keys

## Decision: `Get` stays, but `Lookup` is the more correct API

`Get` stays for ergonomics and mental continuity with v1. `Lookup` is the preferred API for serious scenarios because it returns `error` and makes the flow explicit.

## Decision: thread-safe reads

The structure uses `sync.RWMutex` to protect internal state, allowing concurrent reads safely.

## Additional runtime helpers

`Has`, `HasFor`, and `Reset` were added as low-friction helpers for tests, diagnostics, and controlled reload flows.

## Future evolutions already prepared

The current architecture leaves room for:

- pluralization
- named placeholders
- more sophisticated fallback matching
- incremental loading
- optional hot reload
- missing-key metrics
- integration with external backends
