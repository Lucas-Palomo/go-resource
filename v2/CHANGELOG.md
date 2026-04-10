# Changelog

## v2.0.0

### Added

- nova API baseada em `resource.New(...Option)`
- `LoadDir` e `LoadFS`
- suporte a `fs.FS` e `embed.FS`
- flatten de objetos aninhados
- namespace por pastas e por nome de arquivo
- `Lookup` e `LookupFor`
- estratégias configuráveis para chave ausente e chave duplicada
- `Decoder` extensível
- documentação de migração e arquitetura

### Changed

- módulo agora segue semantic import versioning com `/v2`
- erros deixam de usar `panic` como caminho principal
- fallback de locale foi corrigido

### Fixed

- bug estrutural de fallback quando a locale atual não existia
- acoplamento excessivo entre parsing, walking e lookup
