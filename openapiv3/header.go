package openapiv3

import "errors"

// Header follows the structure of the Parameter Object with the following changes ([ref]):
//
//   - name MUST NOT be specified, it is given in the corresponding headers map.
//   - in MUST NOT be specified, it is implicitly in header.
//   - All traits that are affected by the location MUST be applicable to a location of header (for example, style).
//
// Example:
//
//	 {
//	 "description": "The number of allowed requests in the current period",
//	 "schema": {
//	   "type": "integer"
//	 }
//	}
//
// This object MAY be extended with [Specification Extensions].
//
// [ref]: https://spec.openapis.org/oas/latest.html#header-object
// [Specification Extensions]: https://spec.openapis.org/oas/latest.html#specificationExtensions
type Header struct {
	Parameter
}

// Validate validates a Header.
func (h Header) Validate() error {
	if h.Name == "" {
		return errors.New("name is required")
	}

	return nil
}
