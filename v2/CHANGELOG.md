# Changelog

[English](./CHANGELOG.md) | [Português (Brasil)](./CHANGELOG.pt_br.md)

## v2.0.0

### Added

- new API based on `resource.New(...Option)`
- `LoadDir` and `LoadFS`
- support for `fs.FS` and `embed.FS`
- flattening for nested objects
- namespacing by folders and filename segments
- `Lookup` and `LookupFor`
- configurable strategies for missing keys and duplicate keys
- extensible `Decoder`
- migration and architecture documentation
- runnable example resources
- package examples and extra tests
- `Has`, `HasFor`, and `Reset`
- native Java-style `.properties` decoder support

### Changed

- the module now follows semantic import versioning with `/v2`
- error flows no longer use `panic` as the primary path
- locale fallback behavior was corrected
- markdown documentation was normalized to English with pt-BR companion files

### Fixed

- structural fallback bug when the current locale did not exist
- excessive coupling between parsing, walking, and lookup
- malformed `doc.go` package comment that could break package documentation
