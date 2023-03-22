package openapiv3

import "errors"

// Reference is a simple object to allow referencing other components in the OpenAPI document ([ref]),
// internally and externally.
// The $ref string value contains a URI [[RFC3986]], which identifies the location of the value being referenced.
//
// Examples:
//
//	{
//		"$ref": "#/components/schemas/Pet"
//	}
//
//	{
//	 "$ref": "Pet.json"
//	}
//
//	{
//	 "$ref": "definitions.json#/Pet"
//	}
//
// See the rules for resolving [Relative References].
//
// This object cannot be extended with additional properties and any properties added SHALL be ignored.
//
// Note that this restriction on additional properties is a difference between Reference Objects and Schema Objects that contain a $ref keyword.
//
// [ref]: https://spec.openapis.org/oas/latest.html#reference-object
// [RFC3986]: https://spec.openapis.org/oas/latest.html#bib-RFC3986
// [Relative References]: https://spec.openapis.org/oas/latest.html#relativeReferencesURI
type Reference struct {
	// Reference a Parameter
	// REQUIRED. The reference identifier. This MUST be in the form of a URI.
	Ref string `json:"$ref"`
	// A short summary which by default SHOULD override that of the referenced component.
	// If the referenced object-type does not allow a summary field, then this field has no effect.
	Summary string `json:"summary,omitempty"`
	// A description which by default SHOULD override that of the referenced component.
	// CommonMark syntax (https://spec.commonmark.org/) MAY be used for rich text representation.
	// If the referenced object-type does not allow a description field, then this field has no effect.
	Description string `json:"description,omitempty"`
}

// Validate validates a Reference.
func (r Reference) Validate() error {
	if r.Ref == "" {
		return errors.New("ref is required")
	}
	// TODO: validate ref value.

	return nil
}
