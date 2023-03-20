package openapiv3

// Link represents a possible design-time link for a response ([ref]).
// The presence of a link does not guarantee the caller’s ability to successfully invoke it,
// rather it provides a known relationship and traversal mechanism between responses and other operations.
//
// Unlike dynamic links (i.e. links provided in the response payload),
// the OAS linking mechanism does not require link information in the runtime response.
//
// For computing links, and providing instructions to execute them,
// a [runtime expression] is used for accessing values in an operation
// and using them as parameters while invoking the linked operation.
//
// A linked operation MUST be identified using either an operationRef or operationId.
// In the case of an operationId, it MUST be unique and resolved in the scope of the OAS document.
// Because of the potential for name clashes,
// the operationRef syntax is preferred for O	penAPI documents with external references.
//
// Example:
//
//	{
//	 "paths": {
//	   "/users/{id}": {
//	     "parameters": [
//	       {
//	         "name": "id",
//	         "in": "path",
//	         "required": true,
//	         "description": "the user identifier, as userId",
//	         "schema": {
//	           "type": "string"
//	         }
//	       }
//	     ],
//	     "get": {
//	       "responses": {
//	         "200": {
//	           "description": "the user being returned",
//	           "content": {
//	             "application/json": {
//	               "schema": {
//	                 "type": "object",
//	                 "properties": {
//	                   "uuid": {
//	                     "type": "string",
//	                     "format": "uuid"
//	                   }
//	                 }
//	               }
//	             }
//	           },
//	           "links": {
//	             "address": {
//	               "operationId": "getUserAddress",
//	               "parameters": {
//	                 "userId": "$request.path.id"
//	               }
//	             }
//	           }
//	         }
//	       }
//	     }
//	   },
//	   "/users/{userid}/address": {
//	     "parameters": [
//	       {
//	         "name": "userid",
//	         "in": "path",
//	         "required": true,
//	         "description": "the user identifier, as userId",
//	         "schema": {
//	           "type": "string"
//	         }
//	       }
//	     ],
//	     "get": {
//	       "operationId": "getUserAddress",
//	       "responses": {
//	         "200": {
//	           "description": "the user's address"
//	         }
//	       }
//	     }
//	   }
//	 }
//	}
//
// This object MAY be extended with [Specification Extensions].
//
// [ref]: https://spec.openapis.org/oas/latest.html#link-object
// [runtime expression]: https://spec.openapis.org/oas/latest.html#runtimeExpression
// [Specification Extensions]: https://spec.openapis.org/oas/latest.html#specificationExtensions
type Link struct {
	Reference

	// A relative or absolute URI reference to an OAS operation.
	// This field is mutually exclusive of the operationId field, and MUST point to an Operation Object (https://spec.openapis.org/oas/latest.html#operationObject).
	// Relative operationRef values MAY be used to locate an existing Operation Object in the OpenAPI definition.
	// See the rules for resolving Relative References (https://spec.openapis.org/oas/latest.html#relativeReferencesURI).
	OperationRef string `json:"operationRef,omitempty"`
	// The name of an existing, resolvable OAS operation, as defined with a unique operationId.
	// This field is mutually exclusive of the operationRef field.
	OperationID string `json:"operationId,omitempty"`
	// A map representing parameters to pass to an operation as specified with operationId or identified via operationRef.
	// The key is the parameter name to be used,
	// whereas the value can be a constant or an expression (https://spec.openapis.org/oas/latest.html#runtime-expressions)
	// to be evaluated and passed to the linked operation.
	// The parameter name can be qualified using the parameter location [{in}.]{name} (https://spec.openapis.org/oas/latest.html#parameterIn)
	// for operations that use the same parameter name in different locations (e.g. path.id).
	Parameters map[string]any `json:"parameters,omitempty"`
	// A literal value or {expression} (https://spec.openapis.org/oas/latest.html#runtime-expressions) to use as a request body when calling the target operation.
	RequestBody any `json:"requestBody,omitempty"`
	// A description of the link.
	// CommonMark syntax (https://spec.commonmark.org/) MAY be used for rich text representation.
	Description string `json:"description,omitempty"`
	// A server object to be used by the target operation.
	Server *Server `json:"server,omitempty"`
}
