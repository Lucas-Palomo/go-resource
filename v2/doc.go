// Package resource provides a ResourceBundle-inspired internationalization layer for Go.
//
// The package loads JSON, YAML, TOML, and Java-style .properties resource files from a directory or any fs.FS,
// merges them into locale catalogs, resolves locale fallbacks, and formats dynamic labels.
//
// Supported resource layouts:
//
//	/resources
//	  /errors
//	    en.json
//	    pt_BR.json
//	  /messages
//	    /checkout
//	      en.yaml
//	      pt_BR.yaml
//
// And also the dotted filename layout:
//
//	/resources
//	  en.errors.json
//	  pt_BR.errors.json
//	  en.messages.checkout.toml
//
// Nested objects are flattened using dot notation. For example:
//
//	{
//	  "checkout": {
//	    "title": "Checkout"
//	  }
//	}
//
// becomes:
//
//	checkout.title
//
// When the file is already scoped by namespace, both prefixes are combined.
// For example, resources/errors/en.json with the JSON above becomes the key
// errors.checkout.title.
package resource
