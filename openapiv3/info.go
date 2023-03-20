package openapiv3

import (
	"errors"
	"fmt"
)

// Info provides metadata about the API ([ref]).
// The metadata MAY be used by the clients if needed,
// and MAY be presented in editing or documentation generation tools for convenience.
//
// example:
//
//		{
//		  "title": "Sample Pet Store App",
//		  "summary": "A pet store manager.",
//		  "description": "This is a sample server for a pet store.",
//		  "termsOfService": "https://example.com/terms/",
//		  "contact": {
//		    "name": "API Support",
//		    "url": "https://www.example.com/support",
//		    "email": "support@example.com"
//		  },
//		  "license": {
//		    "name": "Apache 2.0",
//		    "url": "https://www.apache.org/licenses/LICENSE-2.0.html"
//		  },
//	      "version": "1.0.1"
//	   }
//
// This object MAY be extended with [Specification Extensions].
//
// [ref]: https://spec.openapis.org/oas/latest.html#infoObject
// [Specification Extensions]: https://spec.openapis.org/oas/latest.html#specificationExtensions
type Info struct {
	// REQUIRED. The title of the API.
	Title string `json:"title"`
	// A short summary of the API.
	Summary string `json:"summary,omitempty"`
	// A description of the API.
	// CommonMark (https://spec.commonmark.org/) syntax MAY be used for rich text representation.
	Description string `json:"description,omitempty"`
	// A URL to the Terms of Service for the API. This MUST be in the form of a URL.
	TermsOfService string `json:"termsOfservice,omitempty"`
	// The contact information for the exposed API.
	Contact *Contact `json:"contact,omitempty"`
	// The license information for the exposed API.
	License *License `json:"license,omitempty"`
	// REQUIRED. The version of the OpenAPI document (which is distinct from the OpenAPI Specification version or the API implementation version).
	Version string `json:"version"`
}

// Validate validates an Info.
func (i Info) Validate() error {
	if i.Title == "" {
		return errors.New("title is required")
	}
	if i.Version == "" {
		return errors.New("version is required")
	}

	if i.License != nil {
		if err := i.License.Validate(); err != nil {
			return fmt.Errorf("invalid License: %w", err)
		}
	}

	return nil
}
