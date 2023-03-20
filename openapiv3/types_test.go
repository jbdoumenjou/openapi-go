package openapiv3

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOpenAPI_Validate(t *testing.T) {
	tests := []struct {
		desc     string
		OpenAPI  OpenAPI
		expected assert.ErrorAssertionFunc
	}{
		{
			desc:     "valid openAPI document",
			OpenAPI:  OpenAPI{Openapi: "v3"},
			expected: assert.NoError,
		},
		{
			desc:     "expectError openAPI document",
			OpenAPI:  OpenAPI{Openapi: ""},
			expected: assert.Error,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()

			err := test.OpenAPI.Validate()
			test.expected(t, err)
		})
	}
}
