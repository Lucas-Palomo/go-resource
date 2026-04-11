# Changelog

[English](./CHANGELOG.md) | [Português (Brasil)](./CHANGELOG.pt_br.md)

Este changelog acompanha o **módulo v2** (`github.com/Lucas-Palomo/go-resource/v2`).

## [v2.0.0] - Não lançado

### Adicionado

- nova API baseada em `resource.New(...Option)`
- `LoadDir` e `LoadFS`
- suporte a `fs.FS` e `embed.FS`
- flatten de objetos aninhados
- namespacing por pastas e por segmentos do nome do arquivo
- `Lookup` e `LookupFor`
- estratégias configuráveis para chaves ausentes e duplicadas
- `Decoder` extensível
- documentação de migração e arquitetura
- recursos de exemplo executáveis
- exemplos em nível de pacote e testes extras
- `Has`, `HasFor` e `Reset`
- suporte nativo a decoder de arquivos `.properties` no estilo Java

### Alterado

- o módulo segue semantic import versioning por meio de `/v2`
- o fluxo de erro deixa de depender de `panic` como caminho principal
- o comportamento de fallback de locale passou a ser explícito e previsível
- a documentação markdown foi alinhada em inglês com arquivos companheiros em pt-BR

### Corrigido

- bug estrutural de fallback visto na linha antiga quando a locale atual não existia
- acoplamento excessivo entre parsing, walking e responsabilidades de lookup
- problemas de documentação de pacote em `doc.go`
