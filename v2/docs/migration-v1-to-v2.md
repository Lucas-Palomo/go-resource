# Migration from v1 to v2

[English](./migration-v1-to-v2.md) | [Português (Brasil)](./migration-v1-to-v2.pt_br.md)

## In one sentence

v2 keeps the same file-based i18n idea as v1, but replaces implicit behavior with explicit loading, clearer lookup semantics, and a more predictable runtime model.

## Import path

```go
import resource "github.com/Lucas-Palomo/go-resource/v2"
```

## What changed conceptually

### 1. Construction is separate from loading

In v1, resource location and construction were more tightly coupled.

In v2, you create a bundle first and load resources explicitly afterward:

- `New(...Option)`
- `LoadDir(root string) error`
- `LoadFS(fsys fs.FS, root string) error`

This is better for startup validation, testing, and embedded resources.

### 2. File support is broader

v2 supports:

- JSON
- YAML / YML
- TOML
- Java-style `.properties`

It also supports namespace composition from folders and filename segments.

### 3. Lookup now has a strict path and a lenient path

Use the strict path when failure must be visible:

- `Lookup`
- `LookupFor`

Use the lenient path when returning the key is an intentional fallback:

- `Get`
- `GetFor`

Important: `Get` and `GetFor` return the original key on lookup errors. For new service code, prefer the strict path.

### 4. Loading is incremental

A successful second load merges new catalogs into the current bundle in memory.

If your old flow assumed replacement, call `Reset()` before loading again.

### 5. Unsupported files fail explicitly

When the loader finds a file extension without a registered decoder, loading fails with `ErrUnsupportedFormat`.

That is stricter than silently ignoring files, but it makes startup behavior easier to reason about.

## Minimal migration example

```go
bundle := resource.New(
    resource.WithFallbackLocale(language.English),
    resource.WithLocale(language.BrazilianPortuguese),
)

if err := bundle.LoadDir("./resources"); err != nil {
    log.Fatal(err)
}

label, err := bundle.Lookup("checkout.button.confirm")
if err != nil {
    log.Fatal(err)
}
```

## Migration checklist

- move new code to `Lookup` or `LookupFor`
- keep `Get` and `GetFor` only where returning the key is intentional
- verify that the resource tree contains only supported file formats
- decide whether your reload flow should merge or should call `Reset()` first
- prefer `LoadFS` when resources come from `embed.FS` or test fixtures
