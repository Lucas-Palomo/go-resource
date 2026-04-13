# Notas de arquitetura da v2

[English](./architecture.md) | [Português (Brasil)](./architecture.pt_br.md)

## Objetivo arquitetural

A v2 trata `go-resource` primeiro como uma biblioteca reutilizável.

Objetivos de design:

- separar responsabilidades
- reduzir comportamento oculto
- remover `panic` do caminho normal de erro
- tornar explícitos o fallback e o comportamento de lookup
- preservar espaço para extensão

## Separação de responsabilidades

- `bundle.go`: estado de runtime, estado de locale, comportamento de lookup, sincronização
- `loader.go`: walking de filesystem, parsing de path, flattening e tratamento de duplicidade
- `decoder.go`: contrato de decoder e decoders JSON, YAML e TOML
- `properties_decoder.go`: parser de `.properties` no estilo Java
- `options.go`: configuração de políticas em runtime
- `errors.go`: erros sentinela públicos e tipos estruturados de erro

## Por que `fs.FS` importa

A superfície de carregamento é construída sobre `fs.FS`, o que habilita:

- diretórios do SO
- `embed.FS`
- filesystems em memória para testes
- filesystems wrapper de outras bibliotecas

## Modelo de normalização em runtime

A v2 sempre constrói um catálogo plano `map[string]string` por locale.

Pipeline:

1. descobrir o arquivo
2. fazer parsing da locale e do namespace a partir do path
3. decodificar o conteúdo em `map[string]any`
4. achatar objetos aninhados em chaves com notação por ponto
5. fazer merge no catálogo da locale

## Modelo de fallback

A resolução usa uma cadeia de locales baseada em `golang.org/x/text/language`:

- locale solicitada
- seus pais
- locale de fallback
- pais da locale de fallback
- `language.Und` ao final

## Políticas explícitas

A v2 torna duas políticas de primeira classe:

- como tratar chaves ausentes
- como tratar chaves duplicadas

## Modelo de concorrência

O bundle usa `sync.RWMutex`.

Intenção:

- leituras concorrentes seguras
- mutação controlada para troca de locale, registro de decoder, reset e operações de carga
