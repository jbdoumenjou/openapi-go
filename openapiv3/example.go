package openapiv3

import "errors"

// Example defines an example ([ref]).
//
// In all cases, the example value is expected to be compatible with the type schema of its associated value.
// Tooling implementations MAY choose to validate compatibility automatically,
// and reject the example value(s) if incompatible.
//
// Examples:
//
// In a request body:
//
//	{
//	 "requestBody": {
//	   "content": {
//	     "application/json": {
//	       "schema": {
//	         "$ref": "#/components/schemas/Address"
//	       },
//	       "examples": {
//	         "foo": {
//	           "summary": "A foo example",
//	           "value": {
//	             "foo": "bar"
//	           }
//	         },
//	         "bar": {
//	           "summary": "A bar example",
//	           "value": {
//	             "bar": "baz"
//	           }
//	         }
//	       }
//	     },
//	     "application/xml": {
//	       "examples": {
//	         "xmlExample": {
//	           "summary": "This is an example in XML",
//	           "externalValue": "https://example.org/examples/address-example.xml"
//	         }
//	       }
//	     },
//	     "text/plain": {
//	       "examples": {
//	         "textExample": {
//	           "summary": "This is a text example",
//	           "externalValue": "https://foo.bar/examples/address-example.txt"
//	         }
//	       }
//	     }
//	   }
//	 }
//	}
//
// In a parameter:
//
//	{
//	 "parameters": [
//	   {
//	     "name": "zipCode",
//	     "in": "query",
//	     "schema": {
//	       "type": "string",
//	       "format": "zip-code"
//	     },
//	     "examples": {
//	       "zip-example": {
//	         "$ref": "#/components/examples/zip-example"
//	       }
//	     }
//	   }
//	 ]
//	}
//
// In a response:
//
//	{
//	 "responses": {
//	   "200": {
//	     "description": "your car appointment has been booked",
//	     "content": {
//	       "application/json": {
//	         "schema": {
//	           "$ref": "#/components/schemas/SuccessResponse"
//	         },
//	         "examples": {
//	           "confirmation-success": {
//	             "$ref": "#/components/examples/confirmation-success"
//	           }
//	         }
//	       }
//	     }
//	   }
//	 }
//	}
//
// This object MAY be extended with [Specification Extensions].
//
// [ref]: https://spec.openapis.org/oas/latest.html#example-object
// [Specification Extensions]: https://spec.openapis.org/oas/latest.html#specificationExtensions
type Example struct {
	Reference

	// Short description for the example.
	Summary string `json:"summary,omitempty"`
	// Long description for the example.
	// CommonMark syntax (https://spec.commonmark.org/) MAY be used for rich text representation.
	Description string `json:"description,omitempty"`
	// Embedded literal example.
	// The value field and externalValue field are mutually exclusive.
	// To represent examples of media types that cannot be naturally represented in JSON or YAML,
	// use a string value to contain the example, escaping where necessary.
	Value any `json:"value,omitempty"`
	// A URI that points to the literal example.
	// This provides the capability to reference examples that cannot easily be included in JSON or YAML documents.
	// The value field and externalValue field are mutually exclusive.
	// See the rules for resolving Relative References (https://spec.openapis.org/oas/latest.html#relativeReferencesURI).
	ExternalValue string `json:"externalValue,omitempty"`
}

// Validate validates an Example.
func (ex Example) Validate() error {
	isExampleObject := ex.Value != nil || ex.ExternalValue != ""
	isExampleRef := ex.Reference.Ref != ""

	if !isExampleObject && !isExampleRef {
		return errors.New("must be an example object or reference")
	}

	if isExampleRef && isExampleObject {
		return errors.New("example ref and object are mutually exclusive")
	}

	if isExampleRef {
		return ex.Reference.Validate()
	}

	return nil
}
