# Migration from v1 to v2

[English](./migration-v1-to-v2.md) | [Português (Brasil)](./migration-v1-to-v2.pt_br.md)

## Executive summary

v2 changes the project from a small package useful for simple scenarios into a stronger base for real use in libraries, frameworks, and larger applications. The main change is not cosmetic. It corrects the API contract.

## Core problems in v1

### 1. `panic` as error flow

In v1, file-reading failures, parsing problems, and invalid formats crash the application. That is a bad contract for a library. The consumer must decide how to handle the error.

### 2. Fallback bug

In v1, `Get()` tries to fall back to the default locale when the current locale does not resolve the message. The problem is that `GetWithLocale()` returns the key itself when the locale does not exist, so the result is not empty and the fallback does not happen as expected.

### 3. Coupled responsibilities

The same file mixes:

- directory walking
- format parsing
- message merging
- locale selection
- string formatting

That makes evolution, testing, and maintenance harder.

### 4. No `fs.FS`

v1 ties consumption to the traditional filesystem. That limits usage with `embed.FS`.

### 5. No formal policy for collisions and missing keys

When a key is repeated, the behavior is not an explicit API decision. When a key does not exist, the return behavior is also not configurable.

### 6. No support for nested objects

v1 works in practice with `map[string]string`. That limits organization for larger catalogs.

## Breaking changes

### 1. Module path

v2 follows Go semantic major versioning:

```go
import resource "github.com/Lucas-Palomo/go-resource/v2"
```

### 2. Initialization

#### v1

```go
bundle := resource.NewBundle("./resources", language.English)
bundle.Load()
```

#### v2

```go
bundle := resource.New(
	resource.WithFallbackLocale(language.English),
)

if err := bundle.LoadDir("./resources"); err != nil {
	panic(err)
}
```

### 3. Loading

v2 separates loading by source:

- `LoadDir(root string)`
- `LoadFS(fsys fs.FS, root string)`

### 4. Lookup

#### v1

```go
bundle.Get("hello", "Lucas")
bundle.GetWithLocale(language.English, "hello", "Lucas")
```

#### v2

```go
bundle.Get("hello", "Lucas")
bundle.GetFor(language.English, "hello", "Lucas")

value, err := bundle.Lookup("hello", "Lucas")
value, err := bundle.LookupFor(language.English, "hello", "Lucas")
```

`Lookup` is the preferred API when you want explicit error control.

### 5. Nested structures

#### v1

```json
{
  "hello": "Hello"
}
```

#### v2

Besides the flat format, v2 also accepts:

```json
{
  "checkout": {
    "title": "Checkout",
    "button": {
      "confirm": "Confirm"
    }
  }
}
```

That becomes:

- `checkout.title`
- `checkout.button.confirm`

### 6. Namespacing by folder and filename

Example:

```text
resources/errors/en.json
```

with content:

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

## New capabilities

### Missing-key strategy

```go
resource.WithMissingKeyStrategy(resource.ReturnKeyOnMissing)
resource.WithMissingKeyStrategy(resource.ReturnEmptyOnMissing)
resource.WithMissingKeyStrategy(resource.ErrorOnMissing)
```

### Duplicate-key strategy

```go
resource.WithDuplicateKeyStrategy(resource.OverwriteOnDuplicate)
resource.WithDuplicateKeyStrategy(resource.ErrorOnDuplicate)
```

### Custom decoder registration

```go
bundle.RegisterDecoder(".ini", myDecoder)
```

### Inspection and reset helpers

```go
bundle.Has("checkout.title")
bundle.HasFor(language.English, "checkout.title")
bundle.Reset()
```

## Recommended adoption strategy

### Conservative migration

- keep the current file structure
- change the import to `/v2`
- initialize with `resource.New(...)`
- replace `Load()` with `LoadDir()`
- keep using `Get()` at first
- later evolve to `Lookup()` where you need explicit control

### Structural migration

- reorganize messages by namespace
- convert flat files into nested objects
- enable `ErrorOnDuplicate`
- enable `ErrorOnMissing` in test environments

## Full example

```go
bundle := resource.New(
	resource.WithFallbackLocale(language.English),
	resource.WithLocale(language.MustParse("pt-BR")),
	resource.WithMissingKeyStrategy(resource.ErrorOnMissing),
	resource.WithDuplicateKeyStrategy(resource.ErrorOnDuplicate),
)

if err := bundle.LoadDir("./resources"); err != nil {
	return err
}

msg, err := bundle.Lookup("checkout.hello", "Lucas")
if err != nil {
	return err
}

fmt.Println(msg)
```

## Final recommendation

Use v2 as the long-term foundation. v1 can continue to exist for compatibility, but it should not be the main line of evolution.
