package openapiv3

// Components holds a set of reusable objects for different aspects of the OAS ([ref]).
// All objects defined within the components object will have no effect on the API
// unless they are explicitly referenced from properties outside the components object.
//
// All the fixed fields declared above are objects that MUST use keys that match the regular expression: ^[a-zA-Z0-9\.\-_]+$.
//
// Field Name Examples:
//
// User
// User_1
// User_Name
// user-name
// my.org.User
//
// Example:
//
//	"components": {
//	 "schemas": {
//	   "GeneralError": {
//	     "type": "object",
//	     "properties": {
//	       "code": {
//	         "type": "integer",
//	         "format": "int32"
//	       },
//	       "message": {
//	         "type": "string"
//	       }
//	     }
//	   },
//	   "Category": {
//	     "type": "object",
//	     "properties": {
//	       "id": {
//	         "type": "integer",
//	         "format": "int64"
//	       },
//	       "name": {
//	         "type": "string"
//	       }
//	     }
//	   },
//	   "Tag": {
//	     "type": "object",
//	     "properties": {
//	       "id": {
//	         "type": "integer",
//	         "format": "int64"
//	       },
//	       "name": {
//	         "type": "string"
//	       }
//	     }
//	   }
//	 },
//	 "parameters": {
//	   "skipParam": {
//	     "name": "skip",
//	     "in": "query",
//	     "description": "number of items to skip",
//	     "required": true,
//	     "schema": {
//	       "type": "integer",
//	       "format": "int32"
//	     }
//	   },
//	   "limitParam": {
//	     "name": "limit",
//	     "in": "query",
//	     "description": "max records to return",
//	     "required": true,
//	     "schema" : {
//	       "type": "integer",
//	       "format": "int32"
//	     }
//	   }
//	 },
//	 "responses": {
//	   "NotFound": {
//	     "description": "Entity not found."
//	   },
//	   "IllegalInput": {
//	     "description": "Illegal input for operation."
//	   },
//	   "GeneralError": {
//	     "description": "General Error",
//	     "content": {
//	       "application/json": {
//	         "schema": {
//	           "$ref": "#/components/schemas/GeneralError"
//	         }
//	       }
//	     }
//	   }
//	 },
//	 "securitySchemes": {
//	   "api_key": {
//	     "type": "apiKey",
//	     "name": "api_key",
//	     "in": "header"
//	   },
//	   "petstore_auth": {
//	     "type": "oauth2",
//	     "flows": {
//	       "implicit": {
//	         "authorizationUrl": "https://example.org/api/oauth/dialog",
//	         "scopes": {
//	           "write:pets": "modify pets in your account",
//	           "read:pets": "read your pets"
//	         }
//	       }
//	     }
//	   }
//	 }
//	}
//
// This object MAY be extended with [Specification Extensions].
//
// [ref]: https://spec.openapis.org/oas/latest.html#components-object
// [Specification Extensions]: https://spec.openapis.org/oas/latest.html#specificationExtensions
type Components struct {
	// An object to hold reusable Schema Objects.
	Schemas map[string]Schema `json:"schemas,omitempty"`
	// An object to hold reusable Response Objects.
	Responses map[string]Response `json:"responses,omitempty"`
	// An object to hold reusable Parameter Objects.
	Parameters map[string]Parameter `json:"parameters,omitempty"`
	// An object to hold reusable Example Objects.
	Examples map[string]Example `json:"examples,omitempty"`
	// An object to hold reusable Request Body Objects.
	RequestBodies map[string]RequestBody `json:"requestBodies,omitempty"`
	// An object to hold reusable Header Objects.
	Headers map[string]Header `json:"headers,omitempty"`
	// An object to hold reusable Security Scheme Objects.
	SecuritySchemes map[string]SecurityScheme `json:"securitySchemes,omitempty"`
	// An object to hold reusable Link Objects.
	Links map[string]Link `json:"links,omitempty"`
	// An object to hold reusable Callback Objects.
	Callbacks map[string]Callback `json:"callbacks,omitempty"`
	// An object to hold reusable Path Item Object.
	PathItems map[string]PathItem `json:"pathItems,omitempty"`
}
