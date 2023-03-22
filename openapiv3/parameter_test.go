package openapiv3

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParameter_Validate(t *testing.T) {
	tests := []struct {
		desc     string
		json     string
		expected assert.ErrorAssertionFunc
	}{
		{
			desc: "valid parameter object",
			json: `{
      "name": "petId",
      "in": "path",
      "description": "ID of pet that needs to be updated",
      "required": true,
      "schema": {
        "type": "string",
		"externalDocs": {
            "url": "https://foo"
        }
      },
      "examples": { 
         "foo": {
           "summary": "A foo example",
           "value": {"bar": "baz"}
         }
      }
    }`,
			expected: assert.NoError,
		},
		{
			desc: "valid parameter ref",
			json: `{
      "$ref": "#/components/schemas/Pet",
      "summary": "a Pet"
      
    }`,
			expected: assert.NoError,
		},

		{
			desc: "empty object",
			json: `{
      "$ref": "",
      "name": "",
      "in": ""
    }`,
			expected: assert.Error,
		},
		{
			desc: "cannot use both parameter ref and object",
			json: `{
      "$ref": "#/components/schemas/Pet",
      "name": "petId",
      "in": "path",
      "description": "ID of pet that needs to be updated",
      "required": true,
      "schema": {
        "type": "string"
      }
    }`,
			expected: assert.Error,
		},
		{
			desc: "invalid parameter ref with empty ref",
			json: `{
      "$ref": "",
      "summary": "a Pet"
      
    }`,
			expected: assert.Error,
		},
		{
			desc: "invalid parameter object with invalid in",
			json: `{
      "name": "petId",
      "in": "unknown",
      "description": "ID of pet that needs to be updated",
      "required": true,
      "schema": {
        "type": "string"
      }
    }`,
			expected: assert.Error,
		},
		{
			desc: "invalid parameter/schema/externalDocs",
			json: `{
      "name": "petId",
      "in": "path",
      "description": "ID of pet that needs to be updated",
      "required": true,
      "schema": {
        "externalDocs": {
            "url": ""
        }
      }
    }`,
			expected: assert.Error,
		},
		{
			desc: "invalid parameter/examples/example",
			json: `{
      "name": "petId",
      "in": "path",
      "description": "ID of pet that needs to be updated",
      "required": true,      
      "examples": { 
         "foo": {
           "$ref": "foo",
           "summary": "A foo example",
           "value": {"bar": "baz"}
         }
      }
    }`,
			expected: assert.Error,
		},
		{
			desc: "empty parameter/examples/example",
			json: `{
      "name": "petId",
      "in": "path",
      "description": "ID of pet that needs to be updated",
      "required": true,      
      "examples": { 
         "foo": {}
      }
    }`,
			expected: assert.Error,
		},
		{
			desc: "cannot use both example ref and object ",
			json: `{
		 "name": "petId",
		 "in": "path",
		 "description": "ID of pet that needs to be updated",
		 "required": true,
		 "examples": {
		    "foo": {
		      "$ref": "",
              "summary": "A foo example"
		    }
		 }
		}`,
			expected: assert.Error,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()

			var p Parameter
			err := json.Unmarshal([]byte(test.json), &p)
			require.NoError(t, err)

			err = p.Validate()
			test.expected(t, err)
		})
	}
}
