# go-resource v2

[English](./README.md) | [Português (Brasil)](./README.pt_br.md)

`go-resource/v2` is a file-driven i18n library for Go.

It keeps the original ResourceBundle-inspired idea simple: load localized text from files and resolve keys by locale. In v2, that idea is packaged with a clearer runtime contract, explicit loading, better extension points, and support for modern Go filesystem abstractions such as `fs.FS` and `embed.FS`.

## Current release

- stable v2 release: `v2.0.1`
- module path: `github.com/Lucas-Palomo/go-resource/v2`
- recommended for: new projects and migrations from v1

## Installation

```bash
go get github.com/Lucas-Palomo/go-resource/v2@v2.0.1
```

```go
import resource "github.com/Lucas-Palomo/go-resource/v2"
```

## What v2 gives you

- explicit loading through `LoadDir` and `LoadFS`
- support for `fs.FS` and `embed.FS`
- JSON, YAML, TOML, and Java-style `.properties`
- nested object flattening into dot notation
- namespace composition from folders and filename segments
- locale fallback based on `golang.org/x/text/language`
- configurable strategies for missing keys and duplicate keys
- runtime helpers such as `Has`, `Locales`, `Catalog`, and `Reset`
- decoder extensibility with `RegisterDecoder`
- thread-safe reads after loading

## Quick start

```go
package main

import (
	"fmt"
	"log"

	resource "github.com/Lucas-Palomo/go-resource/v2"
	"golang.org/x/text/language"
)

func main() {
	bundle := resource.New(
		resource.WithFallbackLocale(language.English),
		resource.WithLocale(language.BrazilianPortuguese),
	)

	if err := bundle.LoadDir("./resources"); err != nil {
		log.Fatal(err)
	}

	fmt.Println(bundle.Get("title"))
	fmt.Println(bundle.Get("checkout.hello", "Lucas"))
	fmt.Println(bundle.Get("errors.validation.required"))
}
```

## Supported resource formats

Built-in decoders are available for:

- `.json`
- `.yaml`
- `.yml`
- `.toml`
- `.properties`

### Java-style `.properties`

Supported behaviors include:

- separators with `=`, `:`, or whitespace
- comments with `#` or `!`
- line continuation with trailing backslash
- escape sequences such as `\t`, `\n`, `\r`, `\f`, and `\uXXXX`

Example:

```properties
title = Checkout
checkout.hello = Hello, %s
errors.validation.required: Required field
```

## How resource paths become keys

v2 combines three sources to build the final lookup key:

1. folder names under the resource root
2. filename segments after the locale and before the extension
3. nested objects inside the decoded file

### Folder-based namespace

```text
resources/
  errors/
    en.json
  messages/
    checkout/
      pt_BR.yaml
```

This produces keys under `errors.*` and `messages.checkout.*`.

### Filename-based namespace

```text
resources/
  en.messages.checkout.toml
  pt_BR.messages.checkout.toml
```

This produces keys under `messages.checkout.*`.

### Nested objects

```json
{
  "checkout": {
    "button": {
      "confirm": "Confirm"
    }
  }
}
```

This becomes `checkout.button.confirm`.

When folder namespace and nested objects are combined, the prefixes are joined. For example, `resources/errors/en.json` with the JSON above becomes `errors.checkout.button.confirm`.

## Runtime behavior that matters

### Strict lookups vs lenient helpers

Use `Lookup` and `LookupFor` when errors matter.

Use `Get` and `GetFor` when returning the key itself is an acceptable fallback.

That distinction is important: `Get` and `GetFor` return the original key on lookup errors, including missing keys and `ErrBundleNotLoaded`.

### Repeated loads merge into memory

`LoadDir` and `LoadFS` are incremental. A successful second load merges more catalogs into the bundle already in memory.

When you want replacement semantics, call `Reset()` before loading again.

### Unsupported files fail loudly

Files with unsupported extensions are not ignored. Loading stops with `ErrUnsupportedFormat`.

Keep helper files outside the resource tree, or register a decoder for the extension you want to support.

### Scalar normalization

The runtime catalog is normalized to `map[string]string`.

That means:

- strings stay strings
- numbers are converted with `fmt.Sprint`
- booleans are converted with `fmt.Sprint`
- `nil` values are rejected
- unsupported composite leaf values return `ErrInvalidResourceValue`

## Default behavior

`resource.New()` starts with:

- fallback locale: `language.English`
- current locale: fallback locale when not explicitly set
- missing-key strategy: `ReturnKeyOnMissing`
- duplicate-key strategy: `ErrorOnDuplicate`

## Documentation

- [API reference](./docs/api-reference.md)
- [Architecture notes](./docs/architecture.md)
- [Migration from v1](./docs/migration-v1-to-v2.md)
- [v2 changelog](./CHANGELOG.md)
