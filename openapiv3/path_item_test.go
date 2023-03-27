package openapiv3

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPathItem_GetOperation(t *testing.T) {
	pi := PathItem{
		Get: &Operation{
			OperationID: "listPets",
		},
		Post: &Operation{
			OperationID: "createPet",
		},
	}

	operation, err := pi.GetOperation(http.MethodGet)
	require.NoError(t, err)
	assert.Equal(t, pi.Get, operation)
}

func TestPathItem_GetOperation_UnknownMethod(t *testing.T) {
	pi := PathItem{
		Get: &Operation{
			OperationID: "listPets",
		},
		Post: &Operation{
			OperationID: "createPet",
		},
	}

	operation, err := pi.GetOperation("unknown")
	require.Error(t, err)
	assert.Nil(t, operation)
}

func TestPathItem_GetOperation_NotFound(t *testing.T) {
	pi := PathItem{
		Post: &Operation{
			OperationID: "createPet",
		},
	}

	operation, err := pi.GetOperation(http.MethodGet)
	require.NoError(t, err)
	assert.Nil(t, operation)
}
