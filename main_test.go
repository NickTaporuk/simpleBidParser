package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"simpleBidParser/routes"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"bytes"
)

func TestIndexHandler(t *testing.T) {

	var TestsDataJson []string

	TestsDataJson = append(
		TestsDataJson,
		"./testdata/example-request-app-android-1.json",
		"./testdata/example-request-app-android-2.json",
		"./testdata/example-request-web-iphone.json",
		"./testdata/example-request-web-safari.json",
		"./testdata/example-request-web-ie8.json")

	for _, data := range TestsDataJson {

		// download test bid request json file
		raw, err := ioutil.ReadFile(data)

		if err != nil {
			t.Error("Bid request test file not found")
		}

		var BidRequestTestDataAndroid = []byte(raw)

		req, err := http.NewRequest("POST", "/", bytes.NewBuffer(BidRequestTestDataAndroid))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(routes.IndexHandler)

		handler.ServeHTTP(rr, req)

		assert.Equal(t,200, rr.Code,"Ok response is expected")
	}
}
