package openapiv3

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPaths_Get(t *testing.T) {
	paths := Paths{
		"/pets": PathItem{
			Get: &Operation{
				OperationID: "listPets",
			},
		},
		"/pets/{petId}": PathItem{
			Get: &Operation{
				OperationID: "showPetById",
			},
		},
	}

	pathItem, err := paths.Get("/pets/0")
	require.NoError(t, err)
	assert.Equal(t, paths["/pets/{petId}"], pathItem)
}
