package openapiv3

import (
	"fmt"
	"net/http"
)

// PathItem describes the operations available on a single path ([ref]).
// A Path Item MAY be empty, due to [ACL constraints].
// The path itself is still exposed to the documentation viewer,
// but they will not know which operations and parameters are available.
//
// Example:
//
// {
//
//	 "get": {
//	   "description": "Returns pets based on ID",
//	   "summary": "Find pets by ID",
//	   "operationId": "getPetsById",
//	   "responses": {
//	     "200": {
//	       "description": "pet response",
//	       "content": {
//	         "*/*": {
//	           "schema": {
//	             "type": "array",
//	             "items": {
//	               "$ref": "#/components/schemas/Pet"
//	             }
//	           }
//	         }
//	       }
//	     },
//	     "default": {
//	       "description": "error payload",
//	       "content": {
//	         "text/html": {
//	           "schema": {
//	             "$ref": "#/components/schemas/ErrorModel"
//	           }
//	         }
//	       }
//	     }
//	   }
//	 },
//	 "parameters": [
//	   {
//	     "name": "id",
//	     "in": "path",
//	     "description": "ID of pet to use",
//	     "required": true,
//	     "schema": {
//	       "type": "array",
//	       "items": {
//	         "type": "string"
//	       }
//	     },
//	     "style": "simple"
//	   }
//	 ]
//	}
//
// This object MAY be extended with [Specification Extensions].
//
// [ref]: https://spec.openapis.org/oas/latest.html#path-item-object
// [ACL constraints]: https://spec.openapis.org/oas/latest.html#securityFiltering
// [Specification Extensions]: https://spec.openapis.org/oas/latest.html#specificationExtensions
type PathItem struct {
	// Allows for a referenced definition of this path item.
	// The referenced structure MUST be in the form of a Path Item Object (https://spec.openapis.org/oas/latest.html#pathItemObject).
	// In case a Path Item Object field appears both in the defined object and the referenced object, the behavior is undefined.
	// See the rules for resolving Relative References (https://spec.openapis.org/oas/latest.html#relativeReferencesURI).
	Reference

	// A definition of a GET operation on this path.
	Get *Operation `json:"get,omitempty"`
	// A definition of a PUT operation on this path.
	Put *Operation `json:"put,omitempty"`
	// A definition of a POST operation on this path.
	Post *Operation `json:"post,omitempty"`
	// A definition of a DELETE operation on this path.
	Delete *Operation `json:"delete,omitempty"`
	// A definition of an options operation on this path.
	Options *Operation `json:"options,omitempty"`
	// A definition of a HEAD operation on this path.
	Head *Operation `json:"head,omitempty"`
	// A definition of a PATCH operation on this path.
	Patch *Operation `json:"patch,omitempty"`
	// A definition of a Trace operation on this path.
	Trace *Operation `json:"trace,omitempty"`
	// An alternative server array to service all operations in this path.
	Servers []Server `json:"servers,omitempty"`
	// A list of parameters that are applicable for all the operations described under this path.
	// These parameters can be overridden at the operation level, but cannot be removed there.
	// The list MUST NOT include duplicated parameters.
	// A unique parameter is defined by a combination of a name and location.
	// The list can use the Reference Object to link to parameters that are defined at the OpenAPI Object’s components/parameters.
	Parameters []Parameter `json:"parameters,omitempty"`
}

// Validate validates a PathItem.
func (pi *PathItem) Validate() error {
	if pi.Get != nil {
		if err := pi.Get.Validate(); err != nil {
			return fmt.Errorf("invalid Get: %w", err)
		}
	}
	if pi.Put != nil {
		if err := pi.Put.Validate(); err != nil {
			return fmt.Errorf("invalid Put: %w", err)
		}
	}
	if pi.Post != nil {
		if err := pi.Post.Validate(); err != nil {
			return fmt.Errorf("invalid Post: %w", err)
		}
	}
	if pi.Delete != nil {
		if err := pi.Delete.Validate(); err != nil {
			return fmt.Errorf("invalid Delete: %w", err)
		}
	}
	if pi.Options != nil {
		if err := pi.Options.Validate(); err != nil {
			return fmt.Errorf("invalid Options: %w", err)
		}
	}
	if pi.Head != nil {
		if err := pi.Head.Validate(); err != nil {
			return fmt.Errorf("invalid Head: %w", err)
		}
	}
	if pi.Patch != nil {
		if err := pi.Patch.Validate(); err != nil {
			return fmt.Errorf("invalid Patch: %w", err)
		}
	}
	if pi.Trace != nil {
		if err := pi.Trace.Validate(); err != nil {
			return fmt.Errorf("invalid Trace: %w", err)
		}
	}

	return nil
}

// GetOperation returns the operation matching the http method from the PathItem.
func (pi *PathItem) GetOperation(method string) (*Operation, error) {
	switch method {
	case http.MethodGet:
		return pi.Get, nil
	default:
		return nil, fmt.Errorf("unsupported method %q", method)
	}
}
