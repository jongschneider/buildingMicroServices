package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/matryer/is"
)

func TestHandlerReturnsBadRequestWhenNoSearchCriteriaIsSent(t *testing.T) {
	req, resp, handler, is := setupTest(nil, t)

	handler.ServeHTTP(resp, req)

	is.True(resp.Code == http.StatusBadRequest) // 	"Error message should be BadRequest"

}

func TestHandlerReturnsBadRequestWhenBlankSearchCriteriaIsSent(t *testing.T) {
	req, resp, handler, is := setupTest(searchRequest{}, t)

	handler.ServeHTTP(resp, req)

	is.True(resp.Code == http.StatusBadRequest) // 	"Error message should be BadRequest"

}

func setupTest(d interface{}, t *testing.T) (*http.Request, *httptest.ResponseRecorder, Search, *is.I) {
	handler := Search{}
	resp := httptest.NewRecorder()
	is := is.New(t)

	if d == nil {
		return httptest.NewRequest("POST", "/search", nil), resp, handler, is
	}

	body, _ := json.Marshal(d)

	return httptest.NewRequest("POST", "/search", bytes.NewReader(body)), resp, handler, is
}
