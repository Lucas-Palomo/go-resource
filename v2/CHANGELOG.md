# Changelog

[English](./CHANGELOG.md) | [Português (Brasil)](./CHANGELOG.pt_br.md)

This changelog tracks the **v2 module** (`github.com/Lucas-Palomo/go-resource/v2`).

## [v2.0.0] - 2026-04-11

### Added

- new API centered on `resource.New(...Option)`
- `LoadDir` and `LoadFS`
- support for `fs.FS` and `embed.FS`
- built-in decoders for JSON, YAML, TOML, and Java-style `.properties`
- nested-object flattening into dot notation
- namespace composition from folders and filename segments
- `Lookup`, `LookupFor`, `Get`, and `GetFor`
- configurable missing-key and duplicate-key strategies
- runtime helpers such as `Has`, `HasFor`, `Locales`, `Catalog`, and `Reset`
- decoder registration through `RegisterDecoder`
- architecture, API, and migration documents
- runnable examples for the flat and `.properties` workflows
