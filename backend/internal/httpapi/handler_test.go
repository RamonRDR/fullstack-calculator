package httpapi

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type responseBody struct {
	Result *float64 `json:"result"`
	Error  string   `json:"error"`
}

func TestCalculateEndpointReturnsResult(t *testing.T) {
	response := performRequest(t, http.MethodPost, `{"operation":"add","a":10,"b":5}`)

	if response.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", response.Code, http.StatusOK)
	}

	body := decodeResponse(t, response)
	if body.Result == nil || *body.Result != 15 {
		t.Fatalf("result = %v, want 15", body.Result)
	}
}

func TestCalculateEndpointValidationErrors(t *testing.T) {
	tests := []struct {
		name      string
		body      string
		wantError string
	}{
		{name: "division by zero", body: `{"operation":"divide","a":10,"b":0}`, wantError: "division by zero"},
		{name: "unsupported operation", body: `{"operation":"modulo","a":10,"b":3}`, wantError: "unsupported operation"},
		{name: "missing operand", body: `{"operation":"add","a":10}`, wantError: "both operands are required"},
		{name: "missing operation", body: `{"a":10,"b":5}`, wantError: "operation is required"},
		{name: "malformed JSON", body: `{"operation":"add",`, wantError: "invalid request body"},
		{name: "unknown field", body: `{"operation":"add","a":10,"b":5,"extra":true}`, wantError: "invalid request body"},
		{name: "multiple JSON values", body: `{"operation":"add","a":10,"b":5} {}`, wantError: "invalid request body"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			response := performRequest(t, http.MethodPost, tt.body)
			if response.Code != http.StatusBadRequest {
				t.Fatalf("status = %d, want %d", response.Code, http.StatusBadRequest)
			}
			body := decodeResponse(t, response)
			if body.Error != tt.wantError {
				t.Fatalf("error = %q, want %q", body.Error, tt.wantError)
			}
		})
	}
}

func TestCalculateEndpointRejectsOtherMethods(t *testing.T) {
	response := performRequest(t, http.MethodGet, "")

	if response.Code != http.StatusMethodNotAllowed {
		t.Fatalf("status = %d, want %d", response.Code, http.StatusMethodNotAllowed)
	}
	if response.Header().Get("Allow") != http.MethodPost {
		t.Fatalf("Allow header = %q, want %q", response.Header().Get("Allow"), http.MethodPost)
	}
}

func performRequest(t *testing.T, method, body string) *httptest.ResponseRecorder {
	t.Helper()

	request := httptest.NewRequest(method, "/api/calculate", strings.NewReader(body))
	response := httptest.NewRecorder()
	NewHandler().ServeHTTP(response, request)
	return response
}

func decodeResponse(t *testing.T, response *httptest.ResponseRecorder) responseBody {
	t.Helper()

	var body responseBody
	if err := json.NewDecoder(response.Body).Decode(&body); err != nil {
		t.Fatalf("decode response: %v", err)
	}
	return body
}
