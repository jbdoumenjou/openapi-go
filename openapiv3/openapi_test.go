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
			desc: "valid full openAPI document",
			OpenAPI: OpenAPI{
				Openapi: "3.1.0",
				Info: Info{
					Title:          "Sample Pet Store App",
					Summary:        "A pet store manager.",
					Description:    "This is a sample server for a pet store.",
					TermsOfService: "https://example.com/terms/",
					Contact: &Contact{
						Name:  "API Support",
						URL:   "https://www.example.com/support",
						Email: "support@example.com",
					},
					License: &License{
						Name:       "Apache 2.0",
						Identifier: "Apache-2.0",
					},
					Version: "1.0.1",
				},
				JSONSchemaDialect: "https://json-schema.org/draft/2020-12/schema",
				Servers: []Server{
					{
						URL:         "https://development.gigantic-server.com/v1",
						Description: "Development server",
						Variables: map[string]ServerVariable{
							"username": {
								Default:     "demo",
								Description: "this value is assigned by the service provider, in this example `gigantic-server.com`",
							},
							"port": {
								Enum: []string{
									"8443",
									"443",
								},
								Default: "8443",
							},
							"basePath": {
								Default: "v2",
							},
						},
					},
				},
			},
			expected: assert.NoError,
		},
		{
			desc: "valid minimal openAPI document",
			OpenAPI: OpenAPI{
				Openapi: "3.1.0",
				Info: Info{
					Title:   "Sample Pet Store App",
					Version: "1.0.1",
				},
			},
			expected: assert.NoError,
		},
		{
			desc:     "Invalid OpenAPI openapi",
			OpenAPI:  OpenAPI{Openapi: ""},
			expected: assert.Error,
		},
		{
			desc: "Invalid OpenAPI Info.title",
			OpenAPI: OpenAPI{
				Openapi: "3.1.0",
				Info: Info{
					Title:   "",
					Version: "1.0.1",
				},
			},
			expected: assert.Error,
		},
		{
			desc: "Invalid OpenAPI Info.version",
			OpenAPI: OpenAPI{
				Openapi: "3.1.0",
				Info: Info{
					Title:   "Sample Pet Store App",
					Version: "",
				},
			},
			expected: assert.Error,
		},
		{
			desc: "Invalid OpenAPI Info.License empty",
			OpenAPI: OpenAPI{
				Openapi: "3.1.0",
				Info: Info{
					Title:   "Sample Pet Store App",
					Version: "1.0.1",
					License: &License{},
				},
			},
			expected: assert.Error,
		},
		{
			desc: "Invalid OpenAPI Info.License invalid exclusion",
			OpenAPI: OpenAPI{
				Openapi: "3.1.0",
				Info: Info{
					Title:   "Sample Pet Store App",
					Version: "1.0.1",
					License: &License{
						Name:       "Apache 2.0",
						Identifier: "Apache-2.0",
						URL:        "https://example.org",
					},
				},
			},
			expected: assert.Error,
		},
		{
			desc: "Invalid OpenAPI Server.ServerVariable empty default",
			OpenAPI: OpenAPI{
				Openapi: "3.1.0",
				Info: Info{
					Title:   "Sample Pet Store App",
					Version: "1.0.1",
				},
				Servers: []Server{
					{
						URL:         "https://development.gigantic-server.com/v1",
						Description: "Development server",
						Variables: map[string]ServerVariable{
							"username": {
								Default: "demo",
								Enum:    []string{},
							},
						},
					},
				},
			},
			expected: assert.Error,
		},
		{
			desc: "Invalid OpenAPI Server.ServerVariable empty enum",
			OpenAPI: OpenAPI{
				Openapi: "3.1.0",
				Info: Info{
					Title:   "Sample Pet Store App",
					Version: "1.0.1",
				},
				Servers: []Server{
					{
						URL:         "https://development.gigantic-server.com/v1",
						Description: "Development server",
						Variables: map[string]ServerVariable{
							"username": {
								Default:     "",
								Description: "this value is assigned by the service provider, in this example `gigantic-server.com`",
							},
							"port": {
								Enum: []string{
									"8443",
									"443",
								},
								Default: "8443",
							},
							"basePath": {
								Default: "v2",
							},
						},
					},
				},
			},
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
