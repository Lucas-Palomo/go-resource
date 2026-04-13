# Changelog

## [v2.0.1] - 2026-04-13

### Documentation

- rewrote the v2 README with a clearer release narrative and a more approachable usage guide
- clarified installation, supported file formats, namespace composition, and lookup behavior
- aligned the API reference, architecture notes, and migration guide with the actual v2 runtime contract
- synchronized English and pt-BR documentation

### Runtime hardening

- ignored invalid custom decoder registrations when the extension is empty or the decoder is nil
- made duplicate-key validation against previously loaded catalogs happen before any mutation during `LoadFS`
- documented the incremental merge semantics of repeated load operations

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
