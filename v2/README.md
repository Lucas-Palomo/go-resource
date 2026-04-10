# go-resource v2

**ResourceBundle-style i18n for Go** — a modernized, safer and more extensible rewrite of `go-resource`, inspired by the idea of **Java ResourceBundles** and tailored for **Golang internationalization**.

This is the **recommended line for new projects**.

## Install

```bash
go get github.com/Lucas-Palomo/go-resource/v2@latest
```

## Import

```go
import (
    resource "github.com/Lucas-Palomo/go-resource/v2"
    "golang.org/x/text/language"
)
```

## Why v2 exists

The original v1 worked well for small projects, but it mixed directory walking, parsing, loading and lookup into a single flow, used `panic` in the main path, and had fallback behavior that could become inconsistent when a locale was missing.

v2 introduces:

- explicit error handling
- `LoadDir` and `LoadFS`
- support for `embed.FS`
- JSON, YAML and TOML decoders
- flattening of nested structures with dot notation
- predictable fallback resolution
- configurable behavior for missing keys and duplicate keys
- thread-safe reads

## Example

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

## Docs

- [API reference](./docs/api-reference.md)
- [Architecture notes](./docs/architecture.md)
- [Migration guide](./docs/migration-v1-to-v2.md)

## Repository note

The repository root keeps the legacy v1 module for compatibility. This module lives in `/v2` to preserve Go major-version compatibility.
