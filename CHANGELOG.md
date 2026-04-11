# Changelog

[English](./CHANGELOG.md) | [Português (Brasil)](./CHANGELOG.pt_br.md)

This changelog tracks the **root module** (`github.com/Lucas-Palomo/go-resource`), which is the v1 maintenance line.

## [v1.0.1] - 2026-04-11

### Changed

- removed `panic` from the v1 loading path
- added `LoadWithError()` for explicit error handling without breaking the legacy `Load()` API
- added `Err()` so callers using `Load()` can inspect the last loading failure
- corrected fallback resolution in `Get()` so the default locale is actually consulted when the current locale does not contain the key
- initialized `currentLocale` with the default locale in `NewBundle()` for safer out-of-the-box behavior

### Documentation

- rewrote the root README to make the v1 maintenance status explicit
- added dedicated v1 reference docs in English and Portuguese
- aligned the root documentation with the repository multi-major strategy

### Release management

- retracted `v1.0.0` in the root `go.mod`

## [v1.0.0] - 2024-08-20

### Initial public release

- first public release of the v1 API
