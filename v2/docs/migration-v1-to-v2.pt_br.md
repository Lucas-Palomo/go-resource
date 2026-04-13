# Migração da v1 para a v2

[English](./migration-v1-to-v2.md) | [Português (Brasil)](./migration-v1-to-v2.pt_br.md)

## Em uma frase

A v2 mantém a mesma ideia de i18n orientado a arquivos da v1, mas substitui comportamentos implícitos por carregamento explícito, semântica de lookup mais clara e um modelo de runtime mais previsível.

## Path de import

```go
import resource "github.com/Lucas-Palomo/go-resource/v2"
```

## O que mudou conceitualmente

### 1. Construção e carregamento ficaram separados

Na v1, a localização dos recursos e a construção ficavam mais acopladas.

Na v2, você cria o bundle primeiro e carrega os recursos de forma explícita depois:

- `New(...Option)`
- `LoadDir(root string) error`
- `LoadFS(fsys fs.FS, root string) error`

Isso melhora validação no startup, testes e recursos embarcados.

### 2. O suporte a arquivos ficou mais amplo

A v2 suporta:

- JSON
- YAML / YML
- TOML
- `.properties` no estilo Java

Ela também suporta composição de namespace a partir de pastas e segmentos do nome do arquivo.

### 3. O lookup agora tem caminho estrito e caminho leniente

Use o caminho estrito quando a falha precisa ser visível:

- `Lookup`
- `LookupFor`

Use o caminho leniente quando retornar a própria chave for um fallback intencional:

- `Get`
- `GetFor`

Importante: `Get` e `GetFor` retornam a chave original em erros de lookup. Para código novo de serviço, prefira o caminho estrito.

### 4. O carregamento é incremental

Um segundo carregamento bem-sucedido faz merge de novos catálogos no bundle atual em memória.

Se o seu fluxo antigo assumia substituição completa, chame `Reset()` antes de carregar novamente.

### 5. Arquivos não suportados falham de forma explícita

Quando o loader encontra uma extensão sem decoder registrado, o carregamento falha com `ErrUnsupportedFormat`.

Isso é mais estrito do que ignorar arquivos silenciosamente, mas torna o comportamento de startup mais previsível.

## Exemplo mínimo de migração

```go
bundle := resource.New(
    resource.WithFallbackLocale(language.English),
    resource.WithLocale(language.BrazilianPortuguese),
)

if err := bundle.LoadDir("./resources"); err != nil {
    log.Fatal(err)
}

label, err := bundle.Lookup("checkout.button.confirm")
if err != nil {
    log.Fatal(err)
}
```

## Checklist de migração

- mova código novo para `Lookup` ou `LookupFor`
- mantenha `Get` e `GetFor` apenas onde retornar a chave for intencional
- valide que a árvore de recursos contém apenas formatos de arquivo suportados
- decida se o seu fluxo de reload deve fazer merge ou chamar `Reset()` antes
- prefira `LoadFS` quando os recursos vierem de `embed.FS` ou fixtures de teste
