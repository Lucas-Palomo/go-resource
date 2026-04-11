# Versioning and Release Strategy

This repository follows the standard Go major-version strategy.

## v1 line

- module: `github.com/Lucas-Palomo/go-resource`
- package: `github.com/Lucas-Palomo/go-resource/pkg/resource`
- status: maintenance / compatibility line

## Why `v1.0.1`

The initial public release exposed library failures through `panic` in the loading path.
That is not a good default contract for a reusable package.

`v1.0.1` exists to correct the root line without forcing a disruptive rewrite.

## Retract policy

The root `go.mod` retracts `v1.0.0` so consumers see that the initial release should not be used as the recommended version.

## Suggested release flow

1. publish the root maintenance fix as `v1.0.1`
2. keep future root changes minimal and compatibility-focused
3. publish the new major line separately as `v2.0.0` under `/v2`

## Practical consequences

- existing v1 users keep importing `github.com/Lucas-Palomo/go-resource/pkg/resource`
- v1 remains appropriate for low-risk compatibility maintenance
- any architectural evolution should happen in v2
