package openapiv3

// XML is a metadata object that allows for more fine-tuned XML model definitions ([ref]).
//
// When using arrays, XML element names are not inferred (for singular/plural forms)
// and the name property SHOULD be used to add that information.
// See examples for expected behavior.
//
// This object MAY be extended with [Specification Extensions].
//
// [ref]: https://spec.openapis.org/oas/latest.html#xml-object
// [Specification Extensions]: https://spec.openapis.org/oas/latest.html#specificationExtensions
type XML struct {
	// Replaces the name of the element/attribute used for the described schema property.
	// When defined within items, it will affect the name of the individual XML elements within the list.
	// When defined alongside type being array (outside the items),
	// it will affect the wrapping element and only if wrapped is true.
	// If wrapped is false, it will be ignored.
	Name string `json:"name,omitempty"`
	// The URI of the namespace definition. This MUST be in the form of an absolute URI.
	Namespace string `json:"namespace,omitempty"`
	// The prefix to be used for the Name.
	Prefix string `json:"prefix,omitempty"`
	// Declares whether the property definition translates to an attribute instead of an element.
	// Default value is false.
	Attribute bool `json:"attribute"`
	// MAY be used only for an array definition.
	// Signifies whether the array is wrapped (for example, <books><book/><book/></books>) or unwrapped (<book/><book/>).
	// The definition takes effect only when defined alongside type being array (outside the items).
	// Default value is false.
	Wrapped bool `json:"wrapped"`
}
