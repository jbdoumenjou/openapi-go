package openapiv3

import "errors"

// ExternalDocumentation allows referencing an external resource for extended documentation([ref]).
//
// Example:
//
//	{
//	  "description": "Find more info here",
//	  "url": "https://example.com"
//	}
//
// This object MAY be extended with [Specification Extensions].
//
// [ref]: https://spec.openapis.org/oas/latest.html#external-documentation-object
// [Specification Extensions]: https://spec.openapis.org/oas/latest.html#specificationExtensions
type ExternalDocumentation struct {
	// A description of the target documentation.
	// CommonMark syntax (https://spec.commonmark.org/) MAY be used for rich text representation.
	Description string `json:"description,omitempty"`
	// REQUIRED.
	// The URL for the target documentation. This MUST be in the form of a URL.
	URL string `json:"url"`
}

// Validate validates ExternalDocumentation.
func (ed ExternalDocumentation) Validate() error {
	if ed.URL == "" {
		return errors.New("url is required")
	}
	// TODO: URL validation.

	return nil
}
