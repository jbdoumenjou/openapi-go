package openapiv3

// Responses is a container for the expected responses of an operation ([ref]).
// The container maps an HTTP response code to the expected response.
//
// The documentation is not necessarily expected to cover all possible HTTP response codes because they may not be known in advance.
// However, documentation is expected to cover a successful operation response and any known errors.
//
// The default MAY be used as a default response object for all HTTP codes that are not covered individually by the Responses Object.
//
// The Responses Object MUST contain at least one response code,
// and if only one response code is provided it SHOULD be the response for a successful operation call.
//
// Example:
//
//	{
//	 "200": {
//	   "description": "a pet to be returned",
//	   "content": {
//	     "application/json": {
//	       "schema": {
//	         "$ref": "#/components/schemas/Pet"
//	       }
//	     }
//	   }
//	 },
//	 "default": {
//	   "description": "Unexpected error",
//	   "content": {
//	     "application/json": {
//	       "schema": {
//	         "$ref": "#/components/schemas/ErrorModel"
//	       }
//	     }
//	   }
//	 }
//	}
//
// This object MAY be extended with [Specification Extensions].
//
// [ref]: https://spec.openapis.org/oas/latest.html#responses-object
// [Specification Extensions]: https://spec.openapis.org/oas/latest.html#specificationExtensions
type Responses map[string]Response
