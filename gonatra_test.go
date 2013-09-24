package gonatra

import(
    "net/http"
    "net/http/httptest"
    "testing"
)

func TestDispatcher(t *testing.T) {
    rootShowed, albumsShowed, songCreated := false, false, false
    // Register some verbs:
    Get("/", func(http.ResponseWriter, *Request) {
        rootShowed = true
    })

    Get("/albums/:id", func(http.ResponseWriter, *Request) {
        albumsShowed = true
    })
    Post("/albums/:id/songs", func(http.ResponseWriter, *Request) {
        songCreated = true
    })

    url           := "http://example.com/foo"
    response      := httptest.NewRecorder()
    request, err  := http.NewRequest(HTTP_GET, url, nil)
    if err != nil {
        t.Errorf("Something went wrong while creating fake request to %s", url)
        return
    }
    // Test it returns NotFound for no matching route.
    dispatcher(response, request)
    if response.Code != 404 {
        t.Errorf("expected status code to be 404 but got %d", response.Code)
    }

    url           = "http://example.com/albums/123"
    response      = httptest.NewRecorder()
    request, err  = http.NewRequest(HTTP_POST, url, nil) // The route is actually registered for GET verb.
    if err != nil {
        t.Errorf("Something went wrong while creating fake request to %s", url)
        return
    }

    // Test route found but verb doesn't match.
    dispatcher(response, request)
    if response.Code != 404 {
        t.Errorf("expected status code to be 404 but got %d", response.Code)
    }

    // Test it calls the callback for the matched route.
    url           = "http://example.com/albums/123"
    response      = httptest.NewRecorder()
    request, err  = http.NewRequest(HTTP_GET, url, nil)
    if err != nil {
        t.Errorf("Something went wrong while creating fake request to %s", url)
        return
    }

    // Test route found but verb doesn't match.
    dispatcher(response, request)
    if response.Code != 200 {
        t.Errorf("expected status code to be 200 but got %d", response.Code)
    }
    if !albumsShowed {
        t.Errorf("callback for /albums/:id was never called")
    }

    // Test it calls the callback for the matched route.
    url           = "http://example.com/albums/123/songs"
    response      = httptest.NewRecorder()
    request, err  = http.NewRequest(HTTP_POST, url, nil)
    if err != nil {
        t.Errorf("Something went wrong while creating fake request to %s", url)
        return
    }

    // Test route found but verb doesn't match.
    dispatcher(response, request)
    if response.Code != 200 {
        t.Errorf("expected status code to be 200 but got %d", response.Code)
    }
    if !songCreated {
        t.Errorf("callback for /albums/:id/songs was never called")
    }
}

func TestGet(t *testing.T) {
    path      := "/testget"
    result    := Get(path, func(http.ResponseWriter, *Request) {})
    lastRoute := routes[len(routes) -1]

    // Test that the Get method returns true
    if !result {
        t.Errorf("expected Get() to return %t but got %t", !result, result)
    }

    // Test that the given path was set properly
    if lastRoute.Path != path {
        t.Errorf("expected path '%s' but got '%s'", path, lastRoute.Path)
    }

    // Test that the registered verb for the route is GET.
    if lastRoute.Verb != HTTP_GET {
        t.Errorf("expected HTTP verb to be %s but got %s", HTTP_GET, lastRoute.Verb)
    }
}

func TestPost(t *testing.T) {
    path      := "/testpost"
    result    := Post(path, func(http.ResponseWriter, *Request) {})
    lastRoute := routes[len(routes) -1]

    // Test that the Post method returns true
    if !result {
        t.Errorf("expected Post() to return %t but got %t", !result, result)
    }

    // Test that the given path was set properly
    if lastRoute.Path != path {
        t.Errorf("expected path '%s' but got '%s'", path, lastRoute.Path)
    }

    // Test that the registered verb for the route is GET.
    if lastRoute.Verb != HTTP_POST {
        t.Errorf("expected HTTP verb to be %s but got %s", HTTP_POST, lastRoute.Verb)
    }
}
