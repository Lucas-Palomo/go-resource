# Versioning and Release Strategy

[English](./versioning-strategy.md) | [Português (Brasil)](./versioning-strategy.pt_br.md)

This repository follows a **multi-major layout**: the original v1 remains at the repository root, while the breaking rewrite lives in `/v2`.

## Published tags

| Tag | Module | Import path | Date       | Meaning |
|---|---|---|------------|---|
| `v1.0.0` | v1 | `github.com/Lucas-Palomo/go-resource` | 2024-08-20 | First public release, now retracted |
| `v1.0.1` | v1 | `github.com/Lucas-Palomo/go-resource` | 2026-04-11 | Stable maintenance release for existing users |
| `v2.0.0` | v2 | `github.com/Lucas-Palomo/go-resource/v2` | 2026-04-11 | First stable v2 tag, later retracted because its documentation was semantically incorrect |
| `v2.0.1` | v2 | `github.com/Lucas-Palomo/go-resource/v2` | 2026-04-13 | Stable documentation-correction release for v2 |

## Module paths

- root module: `github.com/Lucas-Palomo/go-resource`
- v1 package path: `github.com/Lucas-Palomo/go-resource/pkg/resource`
- v2 module: `github.com/Lucas-Palomo/go-resource/v2`

## Why the repository is structured this way

Go major versions require **semantic import versioning**.

That means:

- the original public line keeps the root import path
- the breaking rewrite must live under `/v2`
- both lines can coexist without breaking existing consumers

## Current support policy

### Root / v1

- purpose: compatibility-first maintenance line
- recommended version: `v1.0.1`
- expected change profile: low-risk fixes, documentation, and stability work

### `/v2`

- purpose: active major line
- recommended version for new projects: `v2.0.1`
- expected change profile: forward-looking feature evolution under the v2 module path

## Why `v1.0.0` was retracted

The first public v1 release exposed normal loading failures through `panic`.

`v1.0.1` corrected that contract without forcing a breaking rewrite:

- `Load()` no longer panics by default
- `LoadWithError()` returns explicit errors
- `Err()` preserves compatibility for older calling styles
- fallback behavior in `Get()` was corrected

The root `go.mod` retracts `v1.0.0` so Go tooling signals that it is not the recommended version.

## Why `v2.0.0` was retracted

`v2.0.0` was functionally valid, but its published documentation described the module state and usage semantics incorrectly enough to justify a correction release.

`v2.0.1` exists to become the stable v2 entry point with corrected documentation and release metadata.

## Practical rules

- keep using v1 when preserving the existing import path matters most
- choose v2 for new code or for teams that need explicit error handling and richer resource modeling
- do not publish breaking API changes in the root module path
- publish future breaking work under the versioned module path that Go expects

## Release rule of thumb

- **v1 releases**: maintenance and compatibility
- **v2 releases**: primary product evolution
