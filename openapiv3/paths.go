package openapiv3

import (
	"fmt"
	"strings"
)

// Paths holds the relative paths to the individual endpoints and their operations ([ref]).
// The path is appended to the URL from the [Server Object] in order to construct the full URL.
// The Paths MAY be empty, due to [Access Control List (ACL) constraints].
//
// The path field name MUST begin with a forward slash (/).
// The path is appended (no relative URL resolution) to the expanded URL from the [Server Object]’s url field
// in order to construct the full URL. [Path templating] is allowed.
// When matching URLs, concrete (non-templated) paths would be matched before their templated counterparts.
// Templated paths with the same hierarchy but different templated names MUST NOT exist as they are identical.
// In case of ambiguous matching, it’s up to the tooling to decide which one to use.
//
// This object MAY be extended with [Specification Extensions].
//
// [ref]: https://spec.openapis.org/oas/latest.html#paths-object
// [Server Object]: https://spec.openapis.org/oas/latest.html#serverObject
// [Access Control List (ACL) constraints]: https://spec.openapis.org/oas/latest.html#securityFiltering
// [Path templating]: https://spec.openapis.org/oas/latest.html#pathTemplating
// [Specification Extensions]: https://spec.openapis.org/oas/latest.html#specificationExtensions
type Paths map[string]PathItem

// Validate validates Paths.
func (pa Paths) Validate() error {
	for _, pathItem := range pa {
		if err := pathItem.Validate(); err != nil {
			return fmt.Errorf("invalid pathItem: %w", err)
		}
	}

	return nil
}

// Get returns a PathItem that matches a request path.
// ex: /pets/0 will return the PathItem under /pets/{petId}.
func (pa Paths) Get(path string) (PathItem, error) {
	pathItem, exists := pa[path]
	if exists {
		return pathItem, nil
	}

	reqIdx := strings.LastIndex(path, "/")
	var hasPrefix bool
	var p string
	for p, pathItem = range pa {
		pathIdx := strings.LastIndex(p, "/")
		if reqIdx == pathIdx && path[:reqIdx] == p[:pathIdx] {
			hasPrefix = true
			continue
		}
	}

	if !hasPrefix {
		return PathItem{}, fmt.Errorf("undefined path: %q", path)
	}

	return pathItem, nil
}
