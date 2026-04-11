# Migration from v1 to v2

[English](./migration-v1-to-v2.md) | [Português (Brasil)](./migration-v1-to-v2.pt_br.md)

## Executive summary

v2 is not a cosmetic rewrite. It corrects the library contract and gives the project a stronger base for long-term evolution.

The practical rule is simple:

- stay on v1.0.1 when compatibility is the main goal
- migrate to v2 when you need a stronger API and room to grow

## Core problems in v1

### 1. `panic` as error flow

Historically, v1 surfaced loading failures through `panic`. That is not a good default for a reusable library.

### 2. Fallback bug

In the old line, `Get()` attempted to fall back to the default locale, but the interaction with `GetWithLocale()` could prevent the fallback from happening correctly when the current locale did not exist.

### 3. Coupled responsibilities

Directory walking, parsing, message merging, locale selection, and formatting lived too close together.

### 4. No `fs.FS`

The old line was tied to the traditional filesystem, which limited usage with `embed.FS`.

### 5. No explicit policy for missing keys and duplicates

Behavior existed, but it was not a clearly configurable API decision.

### 6. No first-class nested object support

Larger catalogs become harder to organize when everything must effectively behave like a flat `map[string]string`.

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
	return err
}
```

### 3. Loading API

v2 separates loading by source:

- `LoadDir(root string)`
- `LoadFS(fsys fs.FS, root string)`

### 4. Lookup API

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

Rule of thumb:

- use `Get` / `GetFor` for convenience
- use `Lookup` / `LookupFor` when explicit error handling matters

### 5. Nested structures

#### v1

```json
{
  "hello": "Hello"
}
```

#### v2

v2 still accepts flat files, but it also accepts nested objects:

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

which become:

- `checkout.title`
- `checkout.button.confirm`

### 6. Namespace by directory and filename

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

### 7. `.properties` support

v2 adds native support for Java-style `.properties` files. That matters if your existing translation assets already use that format.

## New capabilities in v2

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

## Recommended migration strategy

### Conservative migration

Use this path when you want the smallest operational change.

- change the import to `/v2`
- initialize with `resource.New(...)`
- replace `Load()` with `LoadDir()`
- keep the current file layout initially
- keep using `Get()` at first
- migrate specific flows to `Lookup()` where explicit control is valuable

### Structural migration

Use this path when you want to extract more value from v2.

- reorganize resources by namespace
- convert flat files into nested objects where it helps readability
- enable `ErrorOnDuplicate`
- enable `ErrorOnMissing` in tests or validation pipelines
- adopt `.properties` where Java compatibility matters

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

Keep v1 for compatibility.

Choose v2 for the main line of evolution.
