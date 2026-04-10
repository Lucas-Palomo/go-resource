# Versioning and Release Strategy

This repository uses a **multi-major layout**:

- root module: `github.com/Lucas-Palomo/go-resource`
- v2 module: `github.com/Lucas-Palomo/go-resource/v2`

## Why

The project had already been published publicly as a v1 module. To avoid breaking existing consumers, the new major version lives in `/v2`.

## Retract policy

The root `go.mod` retracts `v1.0.0` to signal that the initial public release should not be used as the recommended version.

Suggested release flow:

1. keep root as the legacy v1 line
2. publish a corrected root tag such as `v1.0.1`
3. publish the new major under `/v2` as `v2.0.0`

## Practical consequences

- existing v1 users keep importing `github.com/Lucas-Palomo/go-resource/pkg/resource`
- new projects should prefer `github.com/Lucas-Palomo/go-resource/v2`
- v1 can receive minimal compatibility maintenance while v2 evolves independently
