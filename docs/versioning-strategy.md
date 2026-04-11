# Versioning and Release Strategy

[English](./versioning-strategy.md) | [Português (Brasil)](./versioning-strategy.pt_br.md)

This repository uses a **multi-major layout** so the original public v1 can remain compatible while the next major evolves independently.

## Modules

- root module: `github.com/Lucas-Palomo/go-resource`
- root package: `github.com/Lucas-Palomo/go-resource/pkg/resource`
- v2 module: `github.com/Lucas-Palomo/go-resource/v2`

## Current status

### Root / v1

- status: **stable maintenance line**
- recommended release: **`v1.0.1+`**
- goal: preserve compatibility for existing consumers

### `/v2`

- status: **next major line**
- goal: carry the architectural rewrite and future evolution
- release model: publish separately as **`v2.0.0`** under the `/v2` module path

## Why this layout exists

The project had already been published publicly as a v1 module. Because Go major versions require semantic import versioning, the correct path for the next breaking line is `/v2`.

That keeps existing consumers on the root import path while allowing the new line to evolve without breaking v1 users.

## Why `v1.0.1` exists

The initial public release exposed loading failures through `panic`, which is not a solid default contract for a reusable library.

`v1.0.1` exists to stabilize the root line without forcing a disruptive rewrite:

- keep the existing import path
- stop crashing the process on normal loading failures
- correct fallback behavior
- document the supported v1 contract clearly

## Retract policy

The root `go.mod` retracts `v1.0.0` so Go tooling signals that the first public release should not be considered the recommended version.

## Practical consequences

- existing v1 users keep importing `github.com/Lucas-Palomo/go-resource/pkg/resource`
- the root line should receive only low-risk, compatibility-focused changes
- architectural evolution should happen in `/v2`
- once `v2.0.0` is published, new projects should prefer `github.com/Lucas-Palomo/go-resource/v2`
