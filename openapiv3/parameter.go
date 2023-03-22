package openapiv3

import (
	"errors"
	"fmt"
)

// Parameter describes a single operation parameter ([ref]).
//
// A unique parameter is defined by a combination of a Name and In (location).
//
// There are four possible parameter locations specified by the In field:
//
//   - path - Used together with [Path Templating], where the parameter value is actually part of the operation’s URL.
//     This does not include the host or base path of the API. For example, in /items/{itemId}, the path parameter is itemId.
//   - query - Parameters that are appended to the URL. For example, in /items?id=###, the query parameter is id.
//   - header - Custom headers that are expected as part of the request.
//     Note that [[RFC7230]] states header names are case-insensitive.
//   - cookie - Used to pass a specific cookie value to the API.
//
// The rules for serialization of the parameter are specified in one of two ways.
// For simpler scenarios, a schema and style can describe the structure and syntax of the parameter.
//
// For more complex scenarios, the content property can define the media type and schema of the parameter.
// A parameter MUST contain either a schema property, or a content property, but not both.
// When example or examples are provided in conjunction with the schema object,
// the example MUST follow the prescribed serialization strategy for the parameter.
//
// [ref]: https://spec.openapis.org/oas/latest.html#parameter-object
// [Path Templating]: https://spec.openapis.org/oas/latest.html#pathTemplating
// [RFC7230]: https://spec.openapis.org/oas/latest.html#bib-RFC7230
type Parameter struct {
	Reference

	// REQUIRED. The name of the parameter. Parameter names are case-sensitive.
	//
	// - If in is "path", the name field MUST correspond to a template expression occurring within the path field in the Paths Object.
	// See Path Templating for further information.
	// - If in is "header" and the name field is "Accept", "Content-Type" or "Authorization", the parameter definition SHALL be ignored.
	// - For all other cases, the name corresponds to the parameter name used by the in property.
	// TODO: validation.
	Name string `json:"name"`
	// REQUIRED. The location of the parameter. Possible values are "query", "header", "path" or "cookie".
	In string `json:"in"`
	// A brief description of the parameter.
	// This could contain examples of use.
	// CommonMark syntax (https://spec.commonmark.org/) MAY be used for rich text representation.
	Description string `json:"description,omitempty"`
	// Determines whether this parameter is mandatory.
	// If the parameter location is "path", this property is REQUIRED and its value MUST be true.
	// Otherwise, the property MAY be included and its default value is false.
	Required *bool `json:"required,omitempty"`
	// Specifies that a parameter is deprecated and SHOULD be transitioned out of usage.
	// Default value is false.
	Deprecated bool `json:"deprecated,omitempty"`
	// Sets the ability to pass empty-valued parameters.
	// This is valid only for query parameters and allows sending a parameter with an empty value.
	// Default value is false.
	// If style is used, and if behavior is n/a (cannot be serialized), the value of allowEmptyValue SHALL be ignored.
	// Use of this property is NOT RECOMMENDED, as it is likely to be removed in a later revision.
	AllowEmptyValue bool `json:"allowEmptyValue,omitempty"`

	// Describes how the parameter value will be serialized depending on the type of the parameter value.
	// Default values (based on value of in): for query - form; for path - simple; for header - simple; for cookie - form.
	// TODO: set default value according to In value.
	Style string `json:"style"`
	// When this is true, parameter values of type array or object generate separate parameters
	// for each value of the array or key-value pair of the map.
	// For other types of parameters this property has no effect.
	// When style is form, the default value is true. For all other styles, the default value is false.
	// TODO: set default value according to Style value.
	Explode bool `json:"explode"`
	// The schema defining the type used for the parameter.
	Schema *Schema `json:"schema"`
	// Example of the parameter’s potential value.
	// The example SHOULD match the specified schema and encoding properties if present.
	// The example field is mutually exclusive of the examples field.
	// Furthermore, if referencing a schema that contains an example,
	// the example value SHALL override the example provided by the schema.
	// To represent examples of media types that cannot naturally be represented in JSON or YAML,
	// a string value can contain the example with escaping where necessary.
	Example any `json:"example,omitempty"`
	// Examples of the parameter’s potential value.
	// Each example SHOULD contain a value in the correct format as specified in the parameter encoding.
	// The examples field is mutually exclusive of the example field.
	// Furthermore, if referencing a schema that contains an example,
	// the examples value SHALL override the example provided by the schema.
	Examples map[string]Example `json:"examples,omitempty"`
	// A map containing the representations for the parameter.
	// The key is the media type and the value describes it.
	// The map MUST only contain one entry.
	Content map[string]MediaType `json:"content,omitempty"`
}

// Validate validates a Parameter.
func (p Parameter) Validate() error {
	isParameterObject := p.Name != "" || p.In != "" || p.Required != nil || p.Deprecated || p.AllowEmptyValue
	isParameterRef := p.Reference.Ref != "" || p.Reference.Summary != ""

	if !isParameterObject && !isParameterRef {
		return errors.New("must be a parameter object or reference")
	}

	if isParameterRef && isParameterObject {
		return errors.New("parameter ref and object are mutually exclusive")
	}

	if isParameterRef {
		return p.Reference.Validate()
	}
	// Possible values are "query", "header", "path" or "cookie"
	switch p.In {
	case "query", "header", "path", "cookie":
	default:
		return errors.New("possible values of in are \"query\", \"header\", \"path\" or \"cookie\"")
	}

	if p.Schema != nil {
		if err := p.Schema.Validate(); err != nil {
			return fmt.Errorf("invalid schema: %w", err)
		}
	}

	for _, example := range p.Examples {
		if err := example.Validate(); err != nil {
			return fmt.Errorf("invalid example: %w", err)
		}
	}

	return nil
}
