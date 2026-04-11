# Migration from v1 to v2

[English](./migration-v1-to-v2.md) | [Português (Brasil)](./migration-v1-to-v2.pt_br.md)

## Executive summary

Stay on `v1.0.1` when compatibility is the primary goal.
Move to `v2.0.0` when you want a stronger contract and a more capable resource model.

## Conceptual changes

| Concern | v1 | v2 |
|---|---|---|
| Module path | root | `/v2` |
| Constructor | `NewBundle` | `New(...Option)` |
| Loading | `Load()` / `LoadWithError()` | `LoadDir()` / `LoadFS()` |
| Filesystem | OS only | any `fs.FS` |
| Resource model | flat string maps only | nested objects flattened |
| `.properties` | no | yes |
| Folder / filename namespace | no | yes |
| Missing-key policy | implicit | configurable |
| Duplicate-key policy | implicit | configurable |

## Key behavioral difference

In v1, folder structure and dotted filename segments organize files but do **not** become runtime key prefixes.

In v2, they do.

## Migration checklist

- update imports to `/v2`
- replace `NewBundle` with `New(...Option)`
- replace `GetWithLocale` with `GetFor` or `LookupFor`
- use `LoadDir` or `LoadFS`
- review catalogs that relied on flat keys only
- review whether folder names or dotted filenames should now become namespace prefixes
- choose missing-key and duplicate-key strategies explicitly
