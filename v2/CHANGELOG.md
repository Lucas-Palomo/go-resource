# Changelog

[English](./CHANGELOG.md) | [Português (Brasil)](./CHANGELOG.pt_br.md)

This changelog tracks the **v2 module** (`github.com/Lucas-Palomo/go-resource/v2`).

## [v2.0.0] - Unreleased

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

- the module follows semantic import versioning through `/v2`
- error flows no longer rely on `panic` as the primary path
- locale fallback behavior was made explicit and predictable
- markdown documentation was aligned into English and pt-BR companion files

### Fixed

- structural fallback bug seen in the old line when the current locale did not exist
- excessive coupling between parsing, walking, and lookup responsibilities
- package documentation issues in `doc.go`
