# Changelog

[English](./CHANGELOG.md) | [Português (Brasil)](./CHANGELOG.pt_br.md)

## v2.0.0

### Added

- nova API baseada em `resource.New(...Option)`
- `LoadDir` e `LoadFS`
- suporte a `fs.FS` e `embed.FS`
- flatten de objetos aninhados
- namespace por pastas e por segmentos do nome do arquivo
- `Lookup` e `LookupFor`
- estratégias configuráveis para chave ausente e chave duplicada
- `Decoder` extensível
- documentação de migração e arquitetura
- recursos de exemplo executáveis
- exemplos em nível de pacote e testes extras
- `Has`, `HasFor` e `Reset`
- suporte nativo a decoder de arquivos `.properties` no estilo Java

### Changed

- o módulo agora segue semantic import versioning com `/v2`
- erros deixam de usar `panic` como caminho principal
- o comportamento de fallback de locale foi corrigido
- a documentação markdown foi normalizada para inglês com arquivos companheiros em pt-BR

### Fixed

- bug estrutural de fallback quando a locale atual não existia
- acoplamento excessivo entre parsing, walking e lookup
- comentário de pacote malformado em `doc.go` que podia quebrar a documentação do pacote
