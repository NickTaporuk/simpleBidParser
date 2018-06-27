package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	_ "github.com/stretchr/testify/assert"
	"simpleBidParser/routes"
	"github.com/stretchr/testify/assert"
)

func TestIndexHandler(t *testing.T) {

	req, err := http.NewRequest("POST", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(routes.IndexHandler)

	handler.ServeHTTP(rr, req)

	assert.Equal(t,200, rr.Code,"Ok response is expected")
}
