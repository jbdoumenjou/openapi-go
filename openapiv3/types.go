// Package openapiv3 provides an implementation of
// [OpenAPI V3 specification]: https://spec.openapis.org/oas/latest.html.
package openapiv3

import "errors"

// OpenAPI is the root object of the OpenAPI document,
// [ref]: https://spec.openapis.org/oas/latest.html#openapi-object.
//
// An Open API document is a self-contained or composite resource which defines or describes an API or elements of an API.
// The OpenAPI document MUST contain at least one path field, a components field or a webhooks field.
// An OpenAPI document uses and conforms to the OpenAPI Specification.
type OpenAPI struct {
	// REQUIRED. This string MUST be the version number of the OpenAPI Specification that the OpenAPI document uses.
	// The openapiv3 field SHOULD be used by tooling to interpret the OpenAPI document.
	// This is not related to the API info.version string.
	Openapi string `json:"openapi"`
}

func (o OpenAPI) Validate() error {
	if o.Openapi == "" {
		return errors.New("openapi is required")
	}

	return nil
}
