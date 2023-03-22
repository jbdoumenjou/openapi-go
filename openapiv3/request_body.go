package openapiv3

import (
	"errors"
	"fmt"
)

// RequestBody describes a single request body ([ref]).
//
// The requestBody is fully supported in HTTP methods where the HTTP 1.1 specification [[RFC7231]] has explicitly defined semantics for request bodies.
// In other cases where the HTTP spec is vague (such as [GET], [HEAD] and [DELETE]),
// requestBody is permitted but does not have well-defined semantics and SHOULD be avoided if possible.
//
// [ref]: https://spec.openapis.org/oas/latest.html#request-body-object
// [GET]: https://tools.ietf.org/html/rfc7231#section-4.3.1
// [HEAD]: https://tools.ietf.org/html/rfc7231#section-4.3.2
// [DELETE]: https://tools.ietf.org/html/rfc7231#section-4.3.5
// [RFC7231]: https://spec.openapis.org/oas/latest.html#bib-RFC7231
type RequestBody struct {
	Reference

	// A brief description of the request body.
	// This could contain examples of use.
	// CommonMark syntax (https://spec.commonmark.org/) MAY be used for rich text representation.
	Description string `json:"description,omitempty"`
	// REQUIRED.
	// The content of the request body.
	// The key is a media type or media type range (https://tools.ietf.org/html/rfc7231#appendix-D) and the value describes it.
	// For requests that match multiple keys, only the most specific key is applicable. e.g. text/plain overrides text/*
	Content map[string]MediaType `json:"content"`
	// Determines if the request body is required in the request. Defaults to false.
	Required bool `json:"required"`
}

// Validate validates a RequestBody.
func (rb RequestBody) Validate() error {
	if len(rb.Content) == 0 {
		return errors.New("content is required")
	}

	for _, mediaType := range rb.Content {
		if err := mediaType.Validate(); err != nil {
			return fmt.Errorf("invalid mediaType: %w", err)
		}
	}

	return nil
}
