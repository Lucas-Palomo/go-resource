# go-resource v2

[English](./README.md) | [Português (Brasil)](./README.pt_br.md)

`go-resource/v2` is the stable major line for new development.

It keeps the ResourceBundle-inspired idea, but upgrades the contract in the areas that matter for a reusable Go library: error handling, resource modeling, extension points, and filesystem abstraction.

## Release status

- current stable version: `v2.0.1`
- previous stable tag: `v2.0.0`, now retracted because its published documentation was semantically incorrect
- module path: `github.com/Lucas-Palomo/go-resource/v2`
- recommended audience: new projects and structured migrations from v1

## Installation

```bash
go get github.com/Lucas-Palomo/go-resource/v2@v2.0.1
```

```go
import resource "github.com/Lucas-Palomo/go-resource/v2"
```

## Core features

- explicit loading errors
- `LoadDir(root string)` and `LoadFS(fsys fs.FS, root string)`
- support for `fs.FS` and `embed.FS`
- JSON, YAML, TOML, and Java-style `.properties`
- nested objects flattened into dot notation
- namespace composition from folders and filename segments
- locale fallback chain built on `golang.org/x/text/language`
- configurable missing-key and duplicate-key strategies
- runtime helpers such as `Has`, `HasFor`, `Locales`, `Catalog`, and `Reset`
- decoder extension via `RegisterDecoder`
- thread-safe reads

## Resource conventions

### Folder-based namespace

```text
resources/
  errors/
    en.json
  messages/
    checkout/
      pt_BR.yaml
```

### Filename-based namespace

```text
resources/
  en.messages.checkout.toml
  pt_BR.messages.checkout.toml
```

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

Becomes `checkout.button.confirm`.

When the file is already scoped by folder or filename namespace, the prefixes are combined. For example, `resources/errors/en.json` with the JSON above produces `errors.checkout.button.confirm`.

### Scalar normalization

v2 normalizes the decoded resource catalog into `map[string]string`.

That means:

- strings remain strings
- numbers are converted with `fmt.Sprint`
- booleans are converted with `fmt.Sprint`
- `nil` values are rejected
- unsupported composite leaf values return `ErrInvalidResourceValue`

## Java-style `.properties`

Supported behaviors:

- separators: `=`, `:`, or whitespace
- comments with `#` or `!`
- line continuation with trailing backslash
- escape sequences such as `\t`, `\n`, `\r`, `\f`, and `\uXXXX`

Example:

```properties
title = Checkout
checkout.hello = Hello, %s
errors.validation.required: Required field
```

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

## Default behavior

`resource.New()` starts with these defaults:

- fallback locale: `language.English`
- current locale: fallback locale when not explicitly set
- missing-key strategy: `ReturnKeyOnMissing`
- duplicate-key strategy: `ErrorOnDuplicate`

## Documentation

- [API reference](./docs/api-reference.md)
- [Architecture notes](./docs/architecture.md)
- [Migration from v1](./docs/migration-v1-to-v2.md)
- [v2 changelog](./CHANGELOG.md)
