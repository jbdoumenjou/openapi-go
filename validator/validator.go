// Package validator provides a validation of request based on a OpenAPI specification.
package validator

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/jbdoumenjou/openapi-go/openapiv3"
)

// Validator contains a validator.
type Validator struct {
	OAS *openapiv3.OpenAPI
}

// FromFile loads a validator from a OpenAPI specification file.
func FromFile(path string) (*Validator, error) {
	oas, err := openapiv3.FromFile(path)
	if err != nil {
		return nil, fmt.Errorf("openAPI from file: %w", err)
	}

	checker := &Validator{OAS: oas}

	return checker, nil
}

// Validate validates a request against the OpenApi Specifications.
func (v *Validator) Validate(t *testing.T, handler http.Handler) http.Handler {
	t.Helper()

	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		t.Run("validate request against OpenAPI Specification", func(t *testing.T) {
			if err := v.validateRequest(request); err != nil {
				t.Errorf("OAS validation failed: Request invalid: %s", err.Error())
			}
		})

		handler.ServeHTTP(writer, request)
		// TODO: use the recorder to validate the response.
	})
}

func (v *Validator) validateRequest(request *http.Request) error {
	// TODO: validate server
	requestPath := request.URL.Path
	pathItem, err := v.OAS.Paths.Get(requestPath)
	if err != nil {
		return fmt.Errorf("undefined path %q", requestPath)
	}

	method := request.Method
	operation, err := pathItem.GetOperation(method)
	if err != nil {
		return fmt.Errorf("get operation for path %q: %w", requestPath, err)
	}

	if operation == nil {
		return fmt.Errorf("unsupported method %q for path %q", http.MethodGet, requestPath)
	}

	// for _, parameter := range operation.Parameters {
	//	if parameter.Required != nil && *parameter.Required == true {
	//	}
	//}
	// TODO: validate parameters

	return nil
}
