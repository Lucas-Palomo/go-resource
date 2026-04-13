# Changelog

[English](./CHANGELOG.md) | [Português (Brasil)](./CHANGELOG.pt_br.md)

This changelog tracks the v2 module (`github.com/Lucas-Palomo/go-resource/v2`).

## [v2.0.1] - 2026-04-12

### Documentation

- rewrote the v2 README to match the actual published state of the module
- clarified installation, supported resource conventions, and lookup semantics
- aligned the API reference, architecture notes, and migration guide with the real v2 runtime model
- synchronized English and pt-BR documentation

### Release management

- retracted `v2.0.0` because its documentation was semantically incorrect
- established `v2.0.1` as the stable v2 entry point

### Runtime

- no intended API or behavioral change relative to `v2.0.0`

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
