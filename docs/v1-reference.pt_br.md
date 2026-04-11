# Referência da v1

[English](./v1-reference.md) | [Português (Brasil)](./v1-reference.pt_br.md)

## Pacote

```go
import "github.com/Lucas-Palomo/go-resource/pkg/resource"
```

## O que é a v1

A v1 é a linha de `go-resource` com foco em compatibilidade.

Ela é intencionalmente pequena:

- carregamento baseado em filesystem
- catálogos planos de string
- troca simples de locale
- formatação via `fmt.Sprintf`

A partir da `v1.0.1`, a linha é estável e não usa mais `panic` como comportamento padrão de carregamento.

## Formatos suportados

- JSON
- YAML / YML
- TOML

A v1 **não** suporta:

- arquivos `.properties` no estilo Java
- objetos de recurso aninhados
- composição de namespace a partir de diretórios ou de segmentos pontuados do nome do arquivo

## Modelo de recursos

O caminho de decode da v1 carrega arquivos em `map[string]string`.

Consequências práticas:

- cada valor deve ser string
- as chaves permanecem exatamente como foram declaradas no arquivo
- a estrutura de pastas apenas organiza arquivos em disco
- um nome como `en.errors.json` é aceito porque o primeiro segmento é a locale, mas `errors` não vira parte da chave em runtime

## Construção e carregamento

### `func NewBundle(resourcesFolder string, defaultLocale language.Tag) *Bundle`

Cria um bundle para uma pasta raiz e uma locale default.

### `func (b *Bundle) Load()`

Método de compatibilidade. Na `v1.0.1`, não entra mais em `panic` por padrão. Ele armazena internamente o último erro de carregamento.

### `func (b *Bundle) LoadWithError() error`

Método preferível para código novo na v1. Retorna erros de carregamento diretamente.

### `func (b *Bundle) Err() error`

Retorna o último erro de carregamento capturado por `Load()` ou `LoadWithError()`.

## API de lookup

### `func (b *Bundle) SetLocale(locale language.Tag)`

Altera a locale atual.

### `func (b *Bundle) Get(id string, replacers ...any) string`

Ordem de resolução na `v1.0.1`:

1. locale atual
2. locale default
3. retornar a própria chave se continuar sem resolução

### `func (b *Bundle) GetWithLocale(locale language.Tag, id string, replacers ...any) string`

Resolve uma chave para uma locale específica sem alterar o estado do bundle.

## Valores de erro

- `ErrUnknownBundleEngine`
- `ErrUnsupportedFileExt`
- `ErrInvalidResourceFileName`

## Notas de comportamento

- chaves duplicadas não são expostas como política pública configurável
- arquivos carregados mais tarde podem sobrescrever chaves anteriores na mesma locale
- o pacote é mais adequado para cenários menores e mais simples
- migre para a v2 quando precisar de objetos aninhados, `.properties`, `fs.FS` ou políticas explícitas
