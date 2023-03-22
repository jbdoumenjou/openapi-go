// Package openapiv3 provides an implementation of [OpenAPI V3 specification]
//
// [OpenAPI V3 specification]: https://spec.openapis.org/oas/latest.html..
package openapiv3

import (
	"errors"
	"fmt"
)

// OpenAPI is the root object of the OpenAPI document ([ref]),
//
// An OpenAPI document is a self-contained or composite resource which defines or describes an API or elements of an API.
// The OpenAPI document MUST contain at least one path field, a components field or a webhooks field.
// An OpenAPI document uses and conforms to the OpenAPI Specification.
//
// This object MAY be extended with [Specification Extensions].
//
// [ref]: https://spec.openapis.org/oas/latest.html#openapi-object.
// [Specification Extensions]: https://spec.openapis.org/oas/latest.html#specificationExtensions
type OpenAPI struct {
	// REQUIRED. This string MUST be the version number (https://spec.openapis.org/oas/latest.html#versions)
	// of the OpenAPI Specification that the OpenAPI document uses.
	// The openapiv3 field SHOULD be used by tooling to interpret the OpenAPI document.
	// This is not related to the API info.version (https://spec.openapis.org/oas/latest.html#infoVersion) string.
	Openapi string `json:"openapi"`
	// REQUIRED. Provides metadata about the API. The metadata MAY be used by tooling as required.
	Info Info `json:"info"`
	// The default value for the $schema keyword within
	// Schema Objects (https://spec.openapis.org/oas/latest.html#schemaObject) contained within this OAS document.
	// This MUST be in the form of a URI.
	// TODO: support JSON Schema Specifications (https://datatracker.ietf.org/doc/html/draft-bhutton-json-schema-00)
	JSONSchemaDialect string `json:"jsonSchemaDialect,omitempty"`
	// An array of Server Objects, which provide connectivity information to a target server.
	// If the servers property is not provided, or is an empty array,
	// the default value would be a Server Object with a url value of /.
	Servers []Server `json:"servers,omitempty"`
	// The available paths and operations for the API.
	Paths Paths `json:"paths,omitempty"`
	// The incoming webhooks that MAY be received as part of this API and that the API consumer MAY choose to implement.
	// Closely related to the callbacks feature, this section describes requests initiated other than by an API call,
	// for example by an out-of-band registration.
	// The key name is a unique string to refer to each webhook,
	// while the (optionally referenced) Path Item Object describes a request that may be initiated by the API provider and the expected responses.
	// An example is available.
	Webhooks map[string]PathItem `json:"webhooks,omitempty"`
	// An element to hold various schemas for the document.
	Components *Components `json:"components,omitempty"`
	// A declaration of which security mechanisms can be used across the API.
	// The list of values includes alternative security requirement objects that can be used.
	// Only one of the security requirement objects need to be satisfied to authorize a request.
	// Individual operations can override this definition.
	// To make security optional, an empty security requirement ({}) can be included in the array.
	Security []SecurityRequirement `json:"security,omitempty"`
	// A list of tags used by the document with additional metadata.
	// The order of the tags can be used to reflect on their order by the parsing tools.
	// Not all tags that are used by the Operation Object must be declared.
	// The tags that are not declared MAY be organized randomly or based on the tools’ logic.
	// Each tag name in the list MUST be unique.
	Tags []Tag `json:"tags,omitempty"`
	// Additional external documentation.
	ExternalDocs *ExternalDocumentation `json:"externalDocs,omitempty"`
}

// Validate validates an OpenAPI.
func (o *OpenAPI) Validate() error {
	if o.Openapi == "" {
		return errors.New("openapi is required")
	}

	if err := o.Info.Validate(); err != nil {
		return fmt.Errorf("invalid Info: %w", err)
	}

	for _, server := range o.Servers {
		if err := server.Validate(); err != nil {
			return fmt.Errorf("invalid server: %w", err)
		}
	}

	if o.Paths != nil {
		if err := o.Paths.Validate(); err != nil {
			return fmt.Errorf("invalid paths: %w", err)
		}
	}

	return nil
}
