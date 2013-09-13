package gonatra

import (
    "net/http/httptest"
    "testing"
)

func TestSetHeader(t *testing.T) {
    response := httptest.NewRecorder()
    expectedContentType := "text/plain"
    setHeader(response, "Content-Type", expectedContentType)
    contentType := response.Header().Get("Content-Type")
    if contentType != expectedContentType {
        t.Errorf("expected Content-Type header to be %s but got %s", expectedContentType, contentType)
    }
}

func TestRenderText(t *testing.T) {
    text     := "Go, gonatra!"
    response := httptest.NewRecorder()
    RenderText(response, text)

    // Test that the response body matches the sent string.
    if response.Body.String() != text {
        t.Errorf("expected response body to be %s but got %s", text, response.Body.String())
    }

    // Test that the header is properly set.
    contentType := response.Header().Get("Content-Type")
    if contentType != "text/plain" {
        t.Errorf(`expected Content-Type to be "text/plain" but got %s`, contentType)
    }
}

func TestRenderJSON(t *testing.T) {
    type Fruit struct {
        Name  string `json:"name"`
        Color string `json:"color"`
    }
    fruit        := Fruit{"apple", "red"}
    expectedJSON := `{"name":"apple","color":"red"}`
    response     := httptest.NewRecorder()
    RenderJSON(response, fruit)
    responseBody := response.Body.String()
    // Test that the response body is a JSON representation of the object.
    if responseBody != expectedJSON {
        t.Errorf("expected response body to be %s but got %s", expectedJSON, responseBody)
    }

    // Test that the header is properly set.
    contentType := response.Header().Get("Content-Type")
    if contentType != CONTENT_TYPE_JSON {
        t.Errorf("expected Content-Type to be \"%s\" but got %s", CONTENT_TYPE_JSON, contentType)
    }
}
