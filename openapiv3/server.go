package openapiv3

import "fmt"

// Server is an object representing a Server ([ref]).
//
// example:
//
//	 {
//		   "url": "https://{username}.gigantic-server.com:{port}/{basePath}",
//		   "description": "The production API server",
//		   "variables": {
//		     "username": {
//		       "default": "demo",
//		       "description": "this value is assigned by the service provider, in this example `gigantic-server.com`"
//		     },
//		     "port": {
//		       "enum": [
//		         "8443",
//		         "443"
//		       ],
//		       "default": "8443"
//		     },
//		     "basePath": {
//		     "default": "v2"
//		     }
//		   }
//		}
//
// This object MAY be extended with [Specification Extensions].
//
// [Specification Extensions]: https://spec.openapis.org/oas/latest.html#specificationExtensions
// [ref]: https://spec.openapis.org/oas/latest.html#server-object
type Server struct {
	// REQUIRED. A URL to the target host.
	// This URL supports Server Variables and MAY be relative,
	// to indicate that the host location is relative to the location where the OpenAPI document is being served.
	// Variable substitutions will be made when a variable is named in {brackets}.
	// TODO: URL validation.
	URL string `json:"url"`
	// An optional string describing the host designated by the URL.
	// CommonMark syntax (https://spec.commonmark.org/) MAY be used for rich text representation.
	Description string `json:"description,omitempty"`
	// A map between a variable name and its value. The value is used for substitution in the server’s URL template.
	Variables map[string]ServerVariable `json:"variables,omitempty"`
}

// Validate validates a Server.
func (s Server) Validate() error {
	for _, variable := range s.Variables {
		if err := variable.Validate(); err != nil {
			return fmt.Errorf("invalid server variable: %w", err)
		}
	}
	return nil
}
