package gonatra

import (
    "net/http"
    "testing"
)

func TestBuildRequest(t *testing.T) {
    route        := Route{"/foo/:id/bar/:bar_id", HTTP_GET, nil, nil}
    url          := "http://example.com/foo/123/bar/456"
    request, err := http.NewRequest(HTTP_GET, url, nil)
    if err != nil {
        t.Errorf("Something went wrong while creating fake request to %s", url)
    }
    gonatraRequest := buildRequest(request, &route)
    // Test it has set HttpRequest
    if gonatraRequest.HttpRequest == nil {
        t.Errorf("expected HttpRequest to be set but got nil")
    }

    // Test it has set Params.
    if gonatraRequest.Params == nil {
        t.Errorf("expected Params to be set but got nil")
    }  else {
        // Test it sets the params properly.
        fooParam := gonatraRequest.Params["id"][0]
        barParam := gonatraRequest.Params["bar_id"][0]
        if fooParam != "123" {
            t.Errorf(`expected param "id" to be "123" but got %s`, fooParam)
        }
        if barParam != "456" {
            t.Errorf(`expected param "bar_id" to be "456" but got %s`, barParam)
        }
    }
}

func TestGetParams(t *testing.T) {
    url            := "http://example.com/users/123/articles/456/comments/789?foo=bar&lolz=katz"
    request, err   := http.NewRequest(HTTP_GET, url, nil)
    if err != nil {
        t.Errorf("Something went wrong while creating fake request to %s", url)
    }
    path           := "/users/:id/articles/:article_id/comments/:comment_id"
    route          := Route{path, HTTP_GET, nil, nil}
    params         := getParams(&route, request)
    expectedParams := map[string]string{
        "id":         "123",
        "article_id": "456",
        "comment_id": "789",
        "foo":        "bar",
        "lolz":       "katz",
    }
    for key, expectedValue := range expectedParams {
        val, keyIsPresent := params[key]
        // Test that the key is present.
        if !keyIsPresent {
            t.Errorf("expected key %s to be present in params map.", key)
        } else {
            // Test that the key holds the proper value.
            if expectedValue != val[0] {
                t.Errorf("expected key %s to have value %s but got %s", key, expectedValue, val)
            }
        }
    }
}
