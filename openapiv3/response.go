package openapiv3

// Response describes a single response from an API Operation ([ref]),
// including design-time, static links to operations based on the response.
//
// Examples:
//
//	{
//	 "description": "A complex object array response",
//	 "content": {
//	   "application/json": {
//	     "schema": {
//	       "type": "array",
//	       "items": {
//	         "$ref": "#/components/schemas/VeryComplexType"
//	       }
//	     }
//	   }
//	 }
//	}
//
//	{
//	 "description": "A simple string response",
//	 "content": {
//	   "text/plain": {
//	     "schema": {
//	       "type": "string"
//	     }
//	   }
//	 }
//
// }
//
//	{
//	 "description": "A simple string response",
//	 "content": {
//	   "text/plain": {
//	     "schema": {
//	       "type": "string",
//	       "example": "whoa!"
//	     }
//	   }
//	 },
//	 "headers": {
//	   "X-Rate-Limit-Limit": {
//	     "description": "The number of allowed requests in the current period",
//	     "schema": {
//	       "type": "integer"
//	     }
//	   },
//	   "X-Rate-Limit-Remaining": {
//	     "description": "The number of remaining requests in the current period",
//	     "schema": {
//	       "type": "integer"
//	     }
//	   },
//	   "X-Rate-Limit-Reset": {
//	     "description": "The number of seconds left in the current period",
//	     "schema": {
//	       "type": "integer"
//	     }
//	   }
//	 }
//	}
//
// This object MAY be extended with [Specification Extensions].
//
// [ref]: https://spec.openapis.org/oas/latest.html#response-object
// [Specification Extensions]: https://spec.openapis.org/oas/latest.html#specificationExtensions
type Response struct {
	Reference

	//  REQUIRED. A description of the response.
	// 	CommonMark syntax (https://spec.commonmark.org/) MAY be used for rich text representation.
	Description string `json:"description"`
	// Maps a header name to its definition.
	// [RFC7230] (https://spec.openapis.org/oas/latest.html#bib-RFC7230) states header names are case insensitive.
	// If a response header is defined with the name "Content-Type", it SHALL be ignored.
	Headers map[string]Header
	// A map containing descriptions of potential response payloads.
	// The key is a media type or media type range (https://tools.ietf.org/html/rfc7231#appendix-D) and the value describes it.
	// For responses that match multiple keys, only the most specific key is applicable.
	// e.g. text/plain overrides text/*
	Content map[string]MediaType `json:"content,omitempty"`
	// A map of operations links that can be followed from the response.
	// The key of the map is a short name for the link,
	// following the naming constraints of the names for Component Objects (https://spec.openapis.org/oas/latest.html#componentsObject).
	Links map[string]Link `json:"links,omitempty"`
}
