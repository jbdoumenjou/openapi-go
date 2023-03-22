package openapiv3

import "errors"

// License information for the exposed API ([ref]).
//
// Example:
//
//	{
//	 "name": "Apache 2.0",
//	 "identifier": "Apache-2.0"
//	}
//
// This object MAY be extended with [Specification Extensions].
//
// [ref]: https://spec.openapis.org/oas/latest.html#license-object
// [Specification Extensions]: https://spec.openapis.org/oas/latest.html#specificationExtensions
type License struct {
	// REQUIRED. The license name used for the API.
	Name string `json:"name"`
	// An SPDX (https://spdx.dev/spdx-specification-21-web-version/#h.jxpfx0ykyb60) license expression for the API.
	// The identifier field is mutually exclusive of the url field.
	Identifier string `json:"identifier,omitempty"`
	// A URL to the license used for the API. This MUST be in the form of a URL.
	// The url field is mutually exclusive of the identifier field.
	URL string `json:"url,omitempty"`
}

// Validate validates a License.
func (l *License) Validate() error {
	if l.Name == "" {
		return errors.New("name is required")
	}

	if l.Identifier != "" && l.URL != "" {
		return errors.New("identifier and URL are mutually exclusive")
	}
	// TODO: identifier validation.

	return nil
}
