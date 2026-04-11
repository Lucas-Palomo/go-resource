// Package resource provides a small ResourceBundle-inspired API for loading
// localized messages from JSON, YAML and TOML files.
//
// This package is the v1 maintenance line of go-resource. Starting with
// v1.0.1, the loading path no longer relies on panic as the default error
// strategy. New code should prefer LoadWithError, while Load remains available
// for source compatibility with older call sites.
package resource
