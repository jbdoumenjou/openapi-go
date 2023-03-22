package openapiv3

import "fmt"

// Schema allows the definition of input and output data types ([ref]).
// These types can be objects, but also primitives and arrays.
// This object is a superset of the [JSON Schema Specification Draft 2020-12].
//
// For more information about the properties, see [JSON Schema Core] and [JSON Schema Validation].
//
// Unless stated otherwise, the property definitions follow those of JSON Schema and do not add any additional semantics.
// Where JSON Schema indicates that behavior is defined by the application (e.g. for annotations),
// OAS also defers the definition of semantics to the application consuming the OpenAPI document.
//
// This object MAY be extended with [Specification Extensions], though as noted,
// additional properties MAY omit the x- prefix within this object.
//
// [ref]: https://spec.openapis.org/oas/latest.html#schema-object
// [JSON Schema Specification Draft 2020-12]: https://tools.ietf.org/html/draft-bhutton-json-schema-00
// [JSON Schema Core]: https://tools.ietf.org/html/draft-bhutton-json-schema-00
// [JSON Schema Validation]: https://tools.ietf.org/html/draft-bhutton-json-schema-validation-00
// [Specification Extensions]: https://spec.openapis.org/oas/latest.html#specificationExtensions
type Schema struct {
	// Adds support for polymorphism.
	// The discriminator is an object name that is used to differentiate between other schemas which may satisfy the payload description.
	// See Composition and Inheritance (https://spec.openapis.org/oas/latest.html#schemaComposition) for more details.
	Discriminator *Discriminator `json:"discriminator,omitempty"`
	// This MAY be used only on properties schemas. It has no effect on root schemas.
	// Adds additional metadata to describe the XML representation of this property.
	XML *XML `json:"xml,omitempty"`
	// Additional external documentation for this schema.
	ExternalDocs *ExternalDocumentation `json:"externalDocs,omitempty"`
	// Deprecated: A free-form property to include an example of an instance for this schema.
	// To represent examples that cannot be naturally represented in JSON or YAML,
	// a string value can be used to contain the example with escaping where necessary.
	//
	// The example property has been deprecated in favor of the JSON Schema examples keyword.
	// Use of example is discouraged, and later versions of this specification may remove it.
	Example any `json:"example,omitempty"`
}

// Validate validates a Schema.
func (sc Schema) Validate() error {
	if sc.ExternalDocs != nil {
		if err := sc.ExternalDocs.Validate(); err != nil {
			return fmt.Errorf("invalid externalDocs: %w", err)
		}
	}

	return nil
}
