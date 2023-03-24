package openapiv3

// JSONSchema represents a very naive implementation of [JSON Schema Specification].
//
// [JSON Schema Specification]: https://json-schema.org/specification.html
type JSONSchema struct {
	Format     string                `json:"format,omitempty"`
	Items      map[string]string     `json:"items,omitempty"`
	Maximum    int                   `json:"maximum,omitempty"`
	MaxItems   int                   `json:"maxItems,omitempty"`
	Properties map[string]JSONSchema `json:"properties,omitempty"`
	Ref        string                `json:"$ref,omitempty"`
	Required   []string              `json:"required,omitempty"`
	Type       string                `json:"type,omitempty"`
}
