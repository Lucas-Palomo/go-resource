# v2 Architecture Notes

[English](./architecture.md) | [Português (Brasil)](./architecture.pt_br.md)

## Architectural goal

v2 treats `go-resource` as a reusable library first.

Design goals:

- separate responsibilities
- reduce hidden behavior
- remove `panic` from the normal error path
- make fallback and lookup behavior explicit
- preserve room for extension

## Responsibility split

- `bundle.go`: runtime state, locale state, lookup behavior, synchronization
- `loader.go`: filesystem walking, path parsing, flattening, duplicate handling
- `decoder.go`: decoder contract plus JSON, YAML, and TOML decoders
- `properties_decoder.go`: Java-style `.properties` parser
- `options.go`: runtime policy configuration
- `errors.go`: public sentinel errors and structured error types

## Why `fs.FS` matters

The loading surface is built around `fs.FS`, which enables:

- OS directories
- `embed.FS`
- in-memory test filesystems
- wrapper filesystems from other libraries

## Runtime normalization model

v2 always builds a flat `map[string]string` catalog per locale.

Pipeline:

1. discover the file
2. parse locale and namespace from the path
3. decode content into `map[string]any`
4. flatten nested objects into dot-notated keys
5. merge into the locale catalog

## Fallback model

Resolution uses a locale chain based on `golang.org/x/text/language`:

- requested locale
- its parents
- fallback locale
- parents of the fallback locale
- `language.Und` at the end

## Explicit policies

v2 makes two policies first-class:

- how to handle missing keys
- how to handle duplicate keys

## Concurrency model

The bundle uses `sync.RWMutex`.

Intent:

- safe concurrent reads
- controlled mutation for locale changes, decoder registration, reset, and load operations
