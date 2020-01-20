package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func performRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}
func TestRoot(t *testing.T) {
	body := gin.H{
		"name": "cards",
	}
	// Grab our router
	router := SetupRouter()
	w := performRequest(router, "GET", "/")
	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]string
	err := json.Unmarshal([]byte(w.Body.String()), &response)

	value, exists := response["name"]

	assert.Nil(t, err)
	assert.True(t, exists)
	assert.Equal(t, body["name"], value)
}

func TestCreate(t *testing.T) {
	body := gin.H{
		"remaining": 52,
	}
	// Grab our router
	router := SetupRouter()

	w := performRequest(router, "GET", "/cards/create")
	assert.Equal(t, http.StatusOK, w.Code)

	var response Response
	err := json.Unmarshal([]byte(w.Body.String()), &response)

	value := response.Remaining

	assert.Nil(t, err)
	assert.Equal(t, body["remaining"], value)
}

func TestDraw(t *testing.T) {
	body := gin.H{
		"remaining": 50,
	}
	// Grab our router
	router := SetupRouter()

	w := performRequest(router, "GET", "/cards/create")
	assert.Equal(t, http.StatusOK, w.Code)

	var response Response
	err := json.Unmarshal([]byte(w.Body.String()), &response)

	value := response.DeckId

	url := "/cards/draw?uuid=" + value
	url += "&count=2"
	w = performRequest(router, "GET", url)
	assert.Equal(t, http.StatusOK, w.Code)

	err = json.Unmarshal([]byte(w.Body.String()), &response)
	value1 := response.Remaining

	assert.Nil(t, err)
	assert.Equal(t, body["remaining"], value1)
}

func TestTeardownParallel(t *testing.T) {
	t.Run("group", func(t *testing.T) {
		t.Run("Test1", TestRoot)
		t.Run("Test2", TestCreate)
		t.Run("Test3", TestDraw)
	})
}

func TestMain(m *testing.M) {
	// call flag.Parse() here if TestMain uses flags
	os.Exit(m.Run())
}
