# go-resource

**ResourceBundle-inspired internationalization for Go.**

`go-resource` is a lightweight Go library for managing dynamic labels, messages, errors and localized texts from JSON, YAML and TOML resource files. The project started with an idea similar to Java ResourceBundles, adapted to the Go ecosystem with a small API and a simple file-based workflow.

> This package is the **v1 maintenance line**.
> Version `v1.0.0` was retracted. Use `v1.0.1+` for legacy compatibility.

## Why this release exists

This `v1.0.1` line exists to stabilize the original package before publishing the new major version separately.

Main goals of this patch release:
- stop using `panic` as the default failure mode in a library
- preserve the existing import path: `github.com/Lucas-Palomo/go-resource/pkg/resource`
- keep migration cost small for existing consumers
- document the supported v1 behavior clearly

## Install

```bash
go get github.com/Lucas-Palomo/go-resource@v1.0.1
```

## Quick start

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

## Legacy-compatible loading

The old `Load()` method still exists for compatibility, but it no longer panics.
Instead, it stores the last loading error internally.

```go
bundle.Load()
if err := bundle.Err(); err != nil {
	log.Fatal(err)
}
```

For new code, prefer `LoadWithError()`.

## Supported resource formats

- JSON
- YAML / YML
- TOML

## Resource file contract

v1 expects file names in the form:

```text
<locale>.<ext>
```

Examples:

```text
resources/
  en.json
  pt_BR.yaml
  es.toml
```

Nested directories are allowed. Every valid resource file found under the configured root is merged into the locale catalog.

## Fallback behavior

`Get()` now checks the current locale first and then the configured default locale.
If the key is still not found, the key itself is returned.

## Documentation

- [Versioning and release strategy](./docs/versioning-strategy.md)
- [Estratégia de versionamento e release](./docs/versioning-strategy.pt_br.md)
- [v1 reference](./docs/v1-reference.md)
- [Referência da v1](./docs/v1-reference.pt_br.md)

## Package path

```go
import "github.com/Lucas-Palomo/go-resource/pkg/resource"
```

## License

This repository is distributed under **The Unlicense**.
