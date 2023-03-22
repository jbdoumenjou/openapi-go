package openapiv3

import "fmt"

// MediaType provides schema and examples for the media type identified by its key ([ref]).
//
// Example:
//
//	{
//	 "application/json": {
//	   "schema": {
//	        "$ref": "#/components/schemas/Pet"
//	   },
//	   "examples": {
//	     "cat" : {
//	       "summary": "An example of a cat",
//	       "value":
//	         {
//	           "name": "Fluffy",
//	           "petType": "Cat",
//	           "color": "White",
//	           "gender": "male",
//	           "breed": "Persian"
//	         }
//	     },
//	     "dog": {
//	       "summary": "An example of a dog with a cat's name",
//	       "value" :  {
//	         "name": "Puma",
//	         "petType": "Dog",
//	         "color": "Black",
//	         "gender": "Female",
//	         "breed": "Mixed"
//	       },
//	     "frog": {
//	         "$ref": "#/components/examples/frog-example"
//	       }
//	     }
//	   }
//	 }
//	}
//
// This object MAY be extended with [Specification Extensions].
//
// [Specification Extensions]: https://spec.openapis.org/oas/latest.html#specificationExtensions
// [ref]: https://spec.openapis.org/oas/latest.html#media-type-object
type MediaType struct {
	// The schema defining the content of the request, response, or parameter.
	Schema *Schema `json:"schema,omitempty"`
	// Example of the media type.
	// The example object SHOULD be in the correct format as specified by the media type.
	// The example field is mutually exclusive of the examples field.
	// Furthermore, if referencing a schema which contains an example,
	// the example value SHALL override the example provided by the schema.
	Example any `json:"example,omitempty"`
	// Examples of the media type.
	// Each example object SHOULD match the media type and specified schema if present.
	// The examples field is mutually exclusive of the example field.
	// Furthermore, if referencing a schema which contains an example,
	// the examples value SHALL override the example provided by the schema.
	Examples map[string]Example `json:"examples,omitempty"`
	// A map between a property name and its encoding information.
	// The key, being the property name, MUST exist in the schema as a property.
	// The encoding object SHALL only apply to requestBody objects when the media type is multipart or application/x-www-form-urlencoded.
	Encoding map[string]Encoding `json:"encoding,omitempty"`
}

// Validate validates a MediaType.
func (m MediaType) Validate() error {
	for _, encoding := range m.Encoding {
		if err := encoding.Validate(); err != nil {
			return fmt.Errorf("invalid encoding: %w", err)
		}
	}
	return nil
}
