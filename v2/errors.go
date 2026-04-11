package resource

import (
	"errors"
	"fmt"

	"golang.org/x/text/language"
)

var (
	// ErrBundleNotLoaded indicates that lookups were attempted before any resources were loaded.
	ErrBundleNotLoaded = errors.New("resource bundle not loaded")

	// ErrUnsupportedFormat indicates that no decoder exists for the resource file extension.
	ErrUnsupportedFormat = errors.New("unsupported resource format")

	// ErrInvalidResourceFileName indicates that the file naming convention could not be parsed.
	ErrInvalidResourceFileName = errors.New("invalid resource file name")

	// ErrDuplicateKey indicates that the same fully qualified key was declared more than once.
	ErrDuplicateKey = errors.New("duplicate resource key")

	// ErrMissingKey indicates that a lookup key was not found in the locale chain.
	ErrMissingKey = errors.New("missing resource key")

	// ErrInvalidResourceValue indicates that the resource decoder produced an unsupported value.
	ErrInvalidResourceValue = errors.New("invalid resource value")
)

// KeyNotFoundError describes a lookup miss after locale fallback resolution.
type KeyNotFoundError struct {
	Key      string
	Locale   language.Tag
	Fallback language.Tag
}

func (e KeyNotFoundError) Error() string {
	return fmt.Sprintf("%s: key=%q locale=%q fallback=%q", ErrMissingKey, e.Key, e.Locale, e.Fallback)
}

func (e KeyNotFoundError) Unwrap() error {
	return ErrMissingKey
}

// DuplicateKeyError describes a collision during catalog loading.
type DuplicateKeyError struct {
	Locale language.Tag
	Key    string
	File   string
}

func (e DuplicateKeyError) Error() string {
	return fmt.Sprintf("%s: locale=%q key=%q file=%q", ErrDuplicateKey, e.Locale, e.Key, e.File)
}

func (e DuplicateKeyError) Unwrap() error {
	return ErrDuplicateKey
}

// ResourceFileNameError describes a resource file that does not match the expected convention.
type ResourceFileNameError struct {
	Path   string
	Reason string
}

func (e ResourceFileNameError) Error() string {
	return fmt.Sprintf("%s: path=%q reason=%s", ErrInvalidResourceFileName, e.Path, e.Reason)
}

func (e ResourceFileNameError) Unwrap() error {
	return ErrInvalidResourceFileName
}
