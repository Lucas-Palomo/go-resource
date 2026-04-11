# Changelog

## [v1.0.1] - 2026-04-11

### Changed
- removed `panic` from the v1 loading path
- added `LoadWithError()` for explicit error handling without breaking the existing `Load()` call site
- added `Err()` to inspect the last load failure when using legacy `Load()`
- fixed fallback resolution in `Get()` so the default locale is actually consulted when the current locale does not contain the key
- initialized `currentLocale` with the default locale in `NewBundle()` for safer out-of-the-box behavior

### Documentation
- rewrote the root README for the v1 maintenance line
- added dedicated v1 reference docs in English and Portuguese
- added package documentation through `pkg/resource/doc.go`

### Release management
- retracted `v1.0.0` in the root `go.mod`

## [v1.0.0] - 2024-08-20

### Initial public release
- first public release of the v1 API
