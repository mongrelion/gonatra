package gonatra

import (
    "net/http"
    "testing"
)

func TestGetParams(t *testing.T) {
    url            := "http://example.com/users/123/articles/456/comments/789?foo=bar&lolz=katz"
    request, err   := http.NewRequest(HTTP_GET, url, nil)
    if err != nil {
        t.Errorf("Something went wrong while creating fake request to %s", url)
    }
    path         := "/users/:id/articles/:article_id/comments/:comment_id"
    route        := Route{path, HTTP_GET, nil, nil}
    params       := getParams(&route, request)
    expectations := map[string]string{
        "id"         : "123",
        "article_id" : "456",
        "comment_id" : "789",
        "foo"        : "bar",
        "lolz"       : "katz",
    }
    for param, expected := range expectations {
        actual, keyIsPresent := params[param]
        // Test that the key is present.
        if !keyIsPresent {
            t.Errorf("expected param %s to be present in params map.", param)
        } else {
            // Test that the key holds the proper value.
            if expected != actual {
                t.Errorf("expected param %s to have value %s but got %s", param, expected, actual)
            }
        }
    }
}
