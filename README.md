# go-resource

[English](./README.md) | [Português (Brasil)](./README.pt_br.md)

**ResourceBundle-inspired internationalization for Go.**

`go-resource` is a file-based i18n library for Go, inspired by Java ResourceBundles and adapted to a Go-first workflow. It allows applications to load localized labels, messages, and errors from resource files instead of scattering strings across handlers, services, and domain code.

This repository currently carries **two lines**:

- **v1** at the repository root  
  Stable maintenance line for existing consumers. Current stable release: **`v1.0.1`**.
- **v2** in [`/v2`](./v2)  
  Next major line of the project. It contains the architectural rewrite and the new API.

> `v1.0.0` was retracted. Use `v1.0.1+` for the legacy line.
>
> If you only consume tagged releases, stay on **v1.0.1** until **`v2.0.0`** is published.

## Which line should you use?

### Use v1 when

- you already import `github.com/Lucas-Palomo/go-resource/pkg/resource`
- you want the lowest migration cost
- you only need the original API plus the `v1.0.1` stability fixes

### Use v2 when

- you are planning new work or a structured migration
- you need explicit error handling
- you want `fs.FS` / `embed.FS` support
- you want nested objects flattened into dot notation
- you need configurable policies for missing keys and duplicate keys
- you want native support for Java-style `.properties` files

## Repository layout

```text
.
├── go.mod                    # v1 module: github.com/Lucas-Palomo/go-resource
├── pkg/resource              # v1 package
├── docs                      # root/v1 documentation
└── v2                        # v2 module: github.com/Lucas-Palomo/go-resource/v2
```

## Install

### v1

```bash
go get github.com/Lucas-Palomo/go-resource@v1.0.1
```

Import path:

```go
import "github.com/Lucas-Palomo/go-resource/pkg/resource"
```

### v2

After the first v2 tag is published:

```bash
go get github.com/Lucas-Palomo/go-resource/v2@latest
```

Import path:

```go
import resource "github.com/Lucas-Palomo/go-resource/v2"
```

## v1 quick start

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

## v2 quick start

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

## Resource organization

Both lines accept semantic directory structures such as:

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

The project also supports filename-based namespacing:

```text
/resources
  en.errors.json
  pt_BR.errors.json
  en.messages.checkout.toml
```

In v2, nested objects are flattened automatically. Example:

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

## Supported formats

### v1

- JSON
- YAML / YML
- TOML

### v2

- JSON
- YAML / YML
- TOML
- Java-style `.properties`

## Documentation

### v1

- [Versioning and release strategy](./docs/versioning-strategy.md)
- [v1 API Reference](./docs/v1-reference.md)

### v2

- [v2 README](./v2/README.md)
- [v2 API reference](./v2/docs/api-reference.md)
- [v2 architecture notes](./v2/docs/architecture.md)
- [Migration from v1 to v2](./v2/docs/migration-v1-to-v2.md)

## License

This repository is distributed under **The Unlicense**.
