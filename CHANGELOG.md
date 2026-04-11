# Changelog

[English](./CHANGELOG.md) | [Português (Brasil)](./CHANGELOG.pt_br.md)

This changelog tracks the **root module** (`github.com/Lucas-Palomo/go-resource`), which corresponds to the v1 line.

## [v1.0.1] - 2026-04-11

### Changed

- removed `panic` from the default v1 loading flow
- added `LoadWithError()` for explicit error handling without breaking the legacy `Load()` API
- added `Err()` so callers that still use `Load()` can inspect the last loading failure
- fixed fallback resolution in `Get()` so the default locale is consulted when the current locale does not contain the key
- initialized `currentLocale` with the default locale in `NewBundle()` for safer default behavior

### Documentation

- rewrote the root README around the real repository layout
- documented the difference between the v1 and v2 resource models
- added dedicated v1 reference and versioning documents in English and pt-BR

### Release management

- retracted `v1.0.0` in the root `go.mod`

## [v1.0.0] - 2024-08-20

### Added

- first public release of the original v1 API
