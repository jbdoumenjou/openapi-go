package openapiv3

import "errors"

// ServerVariable is an object representing a Server Variable for server URL template substitution ([ref]).
//
// This object MAY be extended with [Specification Extensions].
//
// [ref]: https://spec.openapis.org/oas/latest.html#server-variable-object
// [Specification Extensions]: https://spec.openapis.org/oas/latest.html#specificationExtensions
type ServerVariable struct {
	// An enumeration of string values to be used if the substitution options are from a limited set.
	// The array MUST NOT be empty.
	Enum []string `json:"enum,omitempty"`
	// REQUIRED. The default value to use for substitution, which SHALL be sent if an alternate value is not supplied.
	// Note this behavior is different from the Schema Object’s treatment of default values,
	// because in those cases parameter values are optional.
	// If the enum is defined, the value MUST exist in the enum’s values.
	Default string `json:"default"`
	// An optional description for the server variable.
	// CommonMark syntax (https://spec.commonmark.org/) MAY be used for rich text representation.
	Description string `json:"description,omitempty"`
}

// Validate validates a ServerVariable.
func (s ServerVariable) Validate() error {
	if s.Default == "" {
		return errors.New("default is required")
	}
	if s.Enum != nil && len(s.Enum) == 0 {
		return errors.New("non empty enum is required")
	}

	return nil
}
