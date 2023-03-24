package openapiv3

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func boolPtr(b bool) *bool {
	return &b
}

func TestOpenAPI_FromFile(t *testing.T) {
	petStoreOAS := &OpenAPI{
		Openapi: "3.0.0",
		Info: Info{
			Version: "1.0.0",
			Title:   "Swagger Petstore",
			License: &License{
				Name: "MIT",
			},
		},
		Servers: []Server{
			{
				URL: "http://petstore.swagger.io/v1",
			},
		},
		Paths: Paths{
			"/pets": PathItem{
				Get: &Operation{
					Summary:     "List all pets",
					OperationID: "listPets",
					Tags:        []string{"pets"},
					Parameters: []Parameter{
						{
							Name:        "limit",
							In:          "query",
							Description: "How many items to return at one time (max 100)",
							Required:    boolPtr(false),
							Schema: &Schema{
								JSONSchema: JSONSchema{
									Type:    "integer",
									Maximum: 100,
									Format:  "int32",
								},
							},
						},
					},
					Responses: &Responses{
						"200": Response{
							Description: "A paged array of pets",
							Headers: map[string]Header{
								"x-next": {
									Parameter: Parameter{
										Description: "A link to the next page of responses",
										Schema: &Schema{
											JSONSchema: JSONSchema{
												Type: "string",
											},
										},
									},
								},
							},
							Content: map[string]MediaType{
								"application/json": {
									Schema: &Schema{
										JSONSchema: JSONSchema{
											Ref: "#/components/schemas/Pets",
										},
									},
								},
							},
						},
						"default": Response{
							Description: "unexpected error",
							Content: map[string]MediaType{
								"application/json": {
									Schema: &Schema{
										JSONSchema: JSONSchema{
											Ref: "#/components/schemas/Error",
										},
									},
								},
							},
						},
					},
				},
				Post: &Operation{
					Summary:     "Create a pet",
					OperationID: "createPets",
					Tags:        []string{"pets"},
					Responses: &Responses{
						"201": Response{
							Description: "Null response",
						},
						"default": Response{
							Description: "unexpected error",
							Content: map[string]MediaType{
								"application/json": {
									Schema: &Schema{
										JSONSchema: JSONSchema{
											Ref: "#/components/schemas/Error",
										},
									},
								},
							},
						},
					},
				},
			},
			"/pets/{petId}": PathItem{
				Get: &Operation{
					Summary:     "Info for a specific pet",
					OperationID: "showPetById",
					Tags:        []string{"pets"},
					Parameters: []Parameter{
						{
							Name:        "petId",
							In:          "path",
							Required:    boolPtr(true),
							Description: "The id of the pet to retrieve",
							Schema: &Schema{
								JSONSchema: JSONSchema{
									Type: "string",
								},
							},
						},
					},
					Responses: &Responses{
						"200": Response{
							Description: "Expected response to a valid request",
							Content: map[string]MediaType{
								"application/json": {
									Schema: &Schema{
										JSONSchema: JSONSchema{
											Ref: "#/components/schemas/Pet",
										},
									},
								},
							},
						},
						"default": Response{
							Description: "unexpected error",
							Content: map[string]MediaType{
								"application/json": {
									Schema: &Schema{
										JSONSchema: JSONSchema{
											Ref: "#/components/schemas/Error",
										},
									},
								},
							},
						},
					},
				},
			},
		},
		Components: &Components{
			Schemas: map[string]Schema{
				"Pet": {
					JSONSchema: JSONSchema{
						Type:     "object",
						Required: []string{"id", "name"},
						Format:   "",
						Properties: map[string]JSONSchema{
							"id": {
								Type:   "integer",
								Format: "int64",
							},
							"name": {
								Type: "string",
							},
							"tag": {
								Type: "string",
							},
						},
					},
				},
				"Pets": {
					JSONSchema: JSONSchema{
						Type:     "array",
						MaxItems: 100,
						Items: map[string]string{
							"$ref": "#/components/schemas/Pet",
						},
					},
				},
				"Error": {
					JSONSchema: JSONSchema{
						Type:     "object",
						Required: []string{"code", "message"},
						Properties: map[string]JSONSchema{
							"code": {
								Type:   "integer",
								Format: "int32",
							},
							"message": {
								Type: "string",
							},
						},
					},
				},
			},
		},
	}

	tests := []struct {
		desc              string
		filePath          string
		expected          *OpenAPI
		expectedAssertion assert.ErrorAssertionFunc
	}{
		{
			desc:              "valid json openAPI specification file",
			filePath:          "testdata/petstore.json",
			expected:          petStoreOAS,
			expectedAssertion: assert.NoError,
		},
		{
			desc:              "valid yaml openAPI specification file",
			filePath:          "testdata/petstore.yaml",
			expected:          petStoreOAS,
			expectedAssertion: assert.NoError,
		},
		{
			desc:              "valid yaml openAPI specification file",
			filePath:          "testdata/petstore.yml",
			expected:          petStoreOAS,
			expectedAssertion: assert.NoError,
		},
		{
			desc:              "unknown openAPI specification file extension",
			filePath:          "testdata/petstore.unknown",
			expectedAssertion: assert.Error,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()
			openapi, err := FromFile(test.filePath)
			test.expectedAssertion(t, err)
			assert.Equal(t, test.expected, openapi)
		})
	}
}

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
