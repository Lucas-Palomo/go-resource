# go-resource v2

[English](./README.md) | [Português (Brasil)](./README.pt_br.md)

**ResourceBundle-style i18n for Go** — a modernized, safer, and more extensible rewrite of `go-resource`.

This module is the **next major line** of the project. It is the long-term direction of the library and the recommended migration target for users that outgrow v1.

> Until `v2.0.0` is tagged, consumers that only use stable releases should remain on `v1.0.1`.

## Module path

```go
import resource "github.com/Lucas-Palomo/go-resource/v2"
```

## Installation

After the first v2 release is published:

```bash
go get github.com/Lucas-Palomo/go-resource/v2@latest
```

## Why v2 exists

v1 was useful for small projects, but it had structural limits:

- loading failures were historically surfaced through `panic`
- directory walking, parsing, loading, fallback, and lookup were tightly coupled
- there was no `fs.FS` / `embed.FS` support
- missing keys and duplicate keys had no explicit policy
- nested objects were not a first-class resource model

v2 fixes the contract instead of only changing names.

## What v2 adds

- explicit error handling
- `LoadDir(root string)` and `LoadFS(fsys fs.FS, root string)`
- support for `fs.FS` and `embed.FS`
- JSON, YAML, TOML, and Java-style `.properties` decoders
- flattening of nested structures into dot notation
- namespacing by directory and filename segments
- predictable fallback resolution
- configurable policies for missing keys and duplicate keys
- thread-safe reads
- helpers such as `Has`, `HasFor`, and `Reset`

## Resource file conventions

### Folder-based namespaces

```text
/resources
  /errors
    en.json
    pt_BR.json
  /messages
    /checkout
      en.yaml
      pt_BR.yaml
```

### Filename-based namespaces

```text
/resources
  en.errors.json
  pt_BR.errors.json
  en.messages.checkout.toml
```

### Java-style `.properties`

```properties
title = Checkout
checkout.hello = Hello, %s
errors.validation.required: Required
checkout.description = Confirm your order  before leaving
welcome = Olá
```

Supported `.properties` features:

- separators: `=`, `:`, or whitespace
- comments with `#` or `!`
- line continuation with trailing backslash
- escape sequences such as `\t`, `\n`, and `\uXXXX`

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

becomes:

```text
checkout.button.confirm
```

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

	fmt.Println(bundle.Get("checkout.title"))
	fmt.Println(bundle.Get("checkout.hello", "Lucas"))
	fmt.Println(bundle.Get("errors.validation.required"))
}
```

Runnable examples are available in:

- [`./examples/basic`](./examples/basic)
- [`./examples/properties`](./examples/properties)

## Docs

- [API reference](./docs/api-reference.md)
- [Architecture notes](./docs/architecture.md)
- [Migration guide](./docs/migration-v1-to-v2.md)

## Relationship with the repository root

The repository root keeps the v1 maintenance line for compatibility. The new major lives under `/v2` to preserve Go semantic import versioning.
