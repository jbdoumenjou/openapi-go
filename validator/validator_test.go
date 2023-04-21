package validator

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type Pet struct {
	ID   int    `json:"id"`   // required
	Name string `json:"name"` // required
	Tag  string `json:"tag,omitempty"`
}

func TestValidator_ValidateCreatePet(t *testing.T) {
	handler := http.HandlerFunc(func(rw http.ResponseWriter, request *http.Request) {
		rw.Header().Set("Content-Type", "application/json")
		rw.WriteHeader(http.StatusCreated)

		var p Pet
		err := json.NewDecoder(request.Body).Decode(&p)
		require.NoError(t, err)

		err = json.NewEncoder(rw).Encode(p)
		require.NoError(t, err)
	})

	validator, err := FromFile("testdata/petstore.yaml")
	require.NoError(t, err)

	server := httptest.NewServer(validator.Validate(t, handler))

	pet := Pet{
		ID:   0,
		Name: "tori",
		Tag:  "bird",
	}
	content, err := json.Marshal(pet)

	require.NoError(t, err)
	req, err := http.NewRequestWithContext(context.Background(), http.MethodPost, server.URL+"/pets", bytes.NewReader(content))
	require.NoError(t, err)

	resp, err := http.DefaultClient.Do(req)
	require.NoError(t, err)

	defer func() { _ = resp.Body.Close() }()

	got, err := io.ReadAll(resp.Body)
	require.NoError(t, err)

	expected, err := json.Marshal(pet)
	require.NoError(t, err)

	assert.JSONEq(t, string(expected), string(got))
}

func TestValidator_ValidateListPets(t *testing.T) {
	pets := []Pet{
		{
			ID:   0,
			Name: "Foo",
			Tag:  "bird",
		},
		{
			ID:   1,
			Name: "Bar",
			Tag:  "dog",
		},
	}

	handler := http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(writer).Encode(pets)
		require.NoError(t, err)
	})

	validator, err := FromFile("testdata/petstore.yaml")
	require.NoError(t, err)

	server := httptest.NewServer(validator.Validate(t, handler))

	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, server.URL+"/pets", http.NoBody)
	require.NoError(t, err)

	resp, err := http.DefaultClient.Do(req)
	require.NoError(t, err)

	defer func() { _ = resp.Body.Close() }()

	got, err := io.ReadAll(resp.Body)
	require.NoError(t, err)

	expected, err := json.Marshal(pets)
	require.NoError(t, err)

	assert.JSONEq(t, string(expected), string(got))
}

func TestValidator_ValidateGetPet(t *testing.T) {
	pet := Pet{
		ID:   0,
		Name: "Foo",
		Tag:  "bird",
	}

	handler := http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")

		err := json.NewEncoder(writer).Encode(pet)
		require.NoError(t, err)
	})

	validator, err := FromFile("testdata/petstore.yaml")
	require.NoError(t, err)

	server := httptest.NewServer(validator.Validate(t, handler))

	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, server.URL+"/pets/0", http.NoBody)
	require.NoError(t, err)

	resp, err := http.DefaultClient.Do(req)
	require.NoError(t, err)

	defer func() { _ = resp.Body.Close() }()

	got, err := io.ReadAll(resp.Body)
	require.NoError(t, err)

	expected, err := json.Marshal(pet)
	require.NoError(t, err)

	assert.JSONEq(t, string(expected), string(got))
}

func TestName(t *testing.T) {
	spec := "/pets/{petId}"
	path := "/pets/0"

	specIdx := strings.LastIndex(spec, "/")
	pathIdx := strings.LastIndex(path, "/")
	require.Equal(t, specIdx, pathIdx)
	expected := spec[:specIdx]
	require.Equal(t, expected, path[:pathIdx])
}
