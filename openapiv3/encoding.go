package openapiv3

import "fmt"

// Encoding A single encoding definition applied to a single schema property ([ref]).
//
// Example:
//
//	{
//	 "requestBody": {
//	   "content": {
//	     "multipart/form-data": {
//	       "schema": {
//	         "type": "object",
//	         "properties": {
//	           "id": {
//	             "type": "string",
//	             "format": "uuid"
//	           },
//	           "address": {
//	             "type": "object",
//	             "properties": {}
//	           },
//	           "historyMetadata": {
//	             "description": "metadata in XML format",
//	             "type": "object",
//	             "properties": {}
//	           },
//	           "profileImage": {}
//	         }
//	       },
//	       "encoding": {
//	         "historyMetadata": {
//	           "contentType": "application/xml; charset=utf-8"
//	         },
//	         "profileImage": {
//	           "contentType": "image/png, image/jpeg",
//	           "headers": {
//	             "X-Rate-Limit-Limit": {
//	               "description": "The number of allowed requests in the current period",
//	               "schema": {
//	                 "type": "integer"
//	               }
//	             }
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
// [ref]: https://spec.openapis.org/oas/latest.html#encoding-object
// [Specification Extensions]: https://spec.openapis.org/oas/latest.html#specificationExtensions
type Encoding struct {
	// The Content-Type for encoding a specific property.
	// Default value depends on the property
	// type: for object - application/json; for array – the default is defined based on the inner type;
	// for all other cases the default is application/octet-stream.
	// The value can be a specific media type (e.g. application/json),
	// a wildcard media type (e.g. image/*), or a comma-separated list of the two types.
	ContentType string `json:"contentType,omitempty"`
	// A map allowing additional information to be provided as headers, for example Content-Disposition.
	// Content-Type is described separately and SHALL be ignored in this section.
	// This property SHALL be ignored if the request body media type is not a multipart.
	Headers map[string]Header `json:"headers,omitempty"`
	// Describes how a specific property value will be serialized depending on its type.
	// See Parameter Object for details on the style property.
	// The behavior follows the same values as query parameters, including default values.
	// This property SHALL be ignored if the request body media type is not application/x-www-form-urlencoded or multipart/form-data.
	// If a value is explicitly defined, then the value of contentType (implicit or explicit) SHALL be ignored.
	Style string `json:"style,omitempty"`
	// When this is true, property values of type array or object generate separate parameters for each value of the array,
	// or key-value-pair of the map. For other types of properties this property has no effect.
	// When Style is form, the default value is true.
	// For all other styles, the default value is false.
	// This property SHALL be ignored if the request body media type is not application/x-www-form-urlencoded or multipart/form-data.
	// If a value is explicitly defined, then the value of contentType (implicit or explicit) SHALL be ignored.
	Explode bool `json:"explode,omitempty"`
	// Determines whether the parameter value SHOULD allow reserved characters,
	// as defined by [RFC3986] :/?#[]@!$&'()*+,;= to be included without percent-encoding.
	// The default value is false.
	// This property SHALL be ignored if the request body media type is not application/x-www-form-urlencoded or multipart/form-data.
	// If a value is explicitly defined, then the value of contentType (implicit or explicit) SHALL be ignored.
	AllowReserved bool `json:"allowReserved,omitempty"`
}

// Validate validates an Encoding.
func (e Encoding) Validate() error {
	for _, header := range e.Headers {
		if err := header.Validate(); err != nil {
			return fmt.Errorf("invalid header: %w", err)
		}
	}
	return nil
}
