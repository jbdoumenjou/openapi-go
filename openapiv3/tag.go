package openapiv3

// Tag adds metadata to a single tag that is used by the Operation Object ([ref]).
// It is not mandatory to have a Tag Object per tag defined in the Operation Object instances.
//
// Example:
//
//	{
//		"name": "pet",
//		"description": "Pets operations"
//	}
//
// This object MAY be extended with [Specification Extensions].
//
// [ref]: https://spec.openapis.org/oas/latest.html#tag-object
// [Specification Extensions]: https://spec.openapis.org/oas/latest.html#specificationExtensions
type Tag struct {
	// REQUIRED. The name of the tag.
	Name string `json:"name"`
	// A description for the tag.
	// CommonMark syntax (https://spec.commonmark.org/) MAY be used for rich text representation.
	Description string `json:"description,omitempty"`
	// Additional external documentation for this tag.
	ExternalDocs ExternalDocumentation `json:"externalDocs,omitempty"`
}
