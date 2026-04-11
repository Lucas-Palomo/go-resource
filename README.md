# go-resource

[English](./README.md) | [Português (Brasil)](./README.pt_br.md)

File-based internationalization for Go, inspired by Java `ResourceBundle`, but adapted to idiomatic Go usage.

This repository contains **two module lines** and **three published tags**:

| Tag | Line | Module path | Date | Status |
|---|---|---|---|---|
| `v1.0.0` | v1 | `github.com/Lucas-Palomo/go-resource` | 2024-08-20 | Historical release, retracted |
| `v1.0.1` | v1 | `github.com/Lucas-Palomo/go-resource` | 2026-04-11 | Stable maintenance release |
| `v2.0.0` | v2 | `github.com/Lucas-Palomo/go-resource/v2` | 2026-04-11 | First stable v2 release |

## What this project does

`go-resource` loads localized messages from resource files instead of hardcoding text across handlers, services, and domain logic.

Supported source formats across the repository:

- JSON
- YAML / YML
- TOML
- Java-style `.properties` in v2

## Which line should you use?

### Use v1 when

- you already import `github.com/Lucas-Palomo/go-resource/pkg/resource`
- you want the lowest migration cost
- you only need the original API with the `v1.0.1` stability fixes

### Use v2 when

- you are starting new work
- you want explicit error handling
- you need `fs.FS` or `embed.FS`
- you want nested objects flattened into dot notation
- you want namespace composition from directories and filename segments
- you need configurable policies for missing keys or duplicate keys
- you want native `.properties` support

## Repository layout

```text
.
├── go.mod                    # v1 module: github.com/Lucas-Palomo/go-resource
├── pkg/resource              # v1 package
├── docs                      # v1-focused documentation
├── examples                  # v1 sample resources
└── v2                        # v2 module: github.com/Lucas-Palomo/go-resource/v2
```

## Install

### v1

```bash
go get github.com/Lucas-Palomo/go-resource@v1.0.1
```

```go
import "github.com/Lucas-Palomo/go-resource/pkg/resource"
```

### v2

```bash
go get github.com/Lucas-Palomo/go-resource/v2@v2.0.0
```

```go
import resource "github.com/Lucas-Palomo/go-resource/v2"
```

## Resource model by line

### v1

v1 loads **flat `map[string]string` catalogs** from JSON, YAML, and TOML files.

Important consequences:

- values must be strings
- nested objects are not a supported resource model
- directories are only for physical organization
- extra filename segments such as `en.errors.json` do **not** become key namespaces

Example:

```text
resources/
  errors/
    en.json
  messages/
    pt_BR.yaml
```

This is valid in v1, but keys stay exactly as declared inside the files.

### v2

v2 loads resource files into **flat runtime catalogs** with richer input rules:

- nested objects are flattened into dot notation
- directory names become namespace segments
- extra filename segments become namespace segments
- `.properties` is supported natively
- scalar values such as strings, numbers, and booleans are converted to strings in the final catalog

Example:

```text
resources/
  errors/
    en.json
  en.messages.checkout.toml
```

Combined with:

```json
{
  "validation": {
    "required": "Required field"
  }
}
```

Produces:

```text
errors.validation.required
```

## Quick start: v1

```go
package main

import (
	"fmt"
	"log"

	"github.com/Lucas-Palomo/go-resource/pkg/resource"
	"golang.org/x/text/language"
)

func main() {
	bundle := resource.NewBundle("./resources", language.English)

	if err := bundle.LoadWithError(); err != nil {
		log.Fatal(err)
	}

	bundle.SetLocale(language.BrazilianPortuguese)

	fmt.Println(bundle.Get("title"))
	fmt.Println(bundle.Get("hello", "Lucas"))
}
```

## Quick start: v2

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

## Documentation

### Repository and v1

- [Versioning and release strategy](./docs/versioning-strategy.md)
- [v1 reference](./docs/v1-reference.md)
- [Root changelog](./CHANGELOG.md)

### v2

- [v2 README](./v2/README.md)
- [v2 API reference](./v2/docs/api-reference.md)
- [v2 architecture notes](./v2/docs/architecture.md)
- [Migration from v1 to v2](./v2/docs/migration-v1-to-v2.md)
- [v2 changelog](./v2/CHANGELOG.md)

## License

This repository is distributed under **The Unlicense**.
