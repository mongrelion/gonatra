package gonatra

import(
    "net/http"
    "net/http/httptest"
    "testing"
)

func TestValidVerb(t *testing.T) {
    invalidVerbs := []string{"W00T", "PATCH", "INVALID", "JOE"}

    for _, verb := range validVerbs {
        if !validVerb(verb) {
            t.Errorf("expected %d verb to be valid but was invalid.")
        }
    }

    for _, verb := range invalidVerbs {
        if validVerb(verb) {
            t.Errorf("expected %d verb to be invalid but was valid.")
        }
    }
}

func TestRegisterRoute(t *testing.T) {
    route  := Route{"/testregisterroute", HTTP_GET, func(http.ResponseWriter, *Request) {}, nil}
    result := registerRoute(HTTP_GET, route.Path, route.Callback)

    // Test that it returns true given a valid path, verb and callback.
    if !result {
        t.Errorf("expected RegisterRoute() to return %t but got %t", !result, result)
    }

    // Test that the registered route matches everything
    lastRoute := routes[len(routes) -1]
    if lastRoute.Path != route.Path {
        t.Errorf("expected route path to be %s but got %s", route.Path, lastRoute.Path)
    }
    if lastRoute.Verb != route.Verb {
        t.Errorf("expected route verb to be %s but got %s", route.Verb, lastRoute.Verb)
    }

    // Test it generates the regexp for the route
    if lastRoute.Rgxp == nil {
        t.Errorf("expected route regular expression to be generated but got nil")
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
        t.Errorf("expected Content-Type to be \"%s\" but got %s", CONTENT_TYPE_PLAIN_TEXT, contentType)
    }
}

func TestRenderJSON(t *testing.T) {
    expectedJSON := "{\"name\":\"apple\",\"color\":\"red\"}"
    response     := httptest.NewRecorder()
    RenderJSON(response, expectedJSON)
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

func TestSetSessionKey(t *testing.T) {
    key, value := "lolcat", "icanhazburger"
    SetSessionKey(key, value)

    // Test the variable is set in session var.
    val, ok := session.m[key]
    if !ok {
        t.Errorf("expected key \"%s\" to be defined in session map.", key)
    }

    // Test the key value is properly set.
    if val != value {
        t.Errorf("expected key \"%s\" key to have value \"%s\" but got %s", key, value, val)
    }
}

func TestGenRouteRegexp(t *testing.T) {
    rgxp := genRouteRegexp("/fruits/:id")
    if rgxp.String() != "/fruits/.+" {
        t.Errorf("expected regular expression body to equal \"/fruits/.+\" but got \"%s\"", rgxp.String())
    }
}

func TestMatchRoute(t *testing.T) {
    carsUrl    := "/cars/123"
    fruitsUrl  := "/fruits/123"
    route      := Route{"/cars/:id", HTTP_GET, nil, nil}
    route.Rgxp = genRouteRegexp(route.Path)

    if !matchRoute(&route, carsUrl) {
        t.Errorf("expected route %s to match %s", route.Path, carsUrl)
    }
    if matchRoute(&route, fruitsUrl) {
        t.Errorf("expected route %s not to match %s", route.Path, fruitsUrl)
    }
}

func TestDispatcher(t *testing.T) {
    albumsShowed, songCreated := false, false
    // Register some verbs:
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

func TestSetHeader(t *testing.T) {
    response := httptest.NewRecorder()
    expectedContentType := "text/plain"
    setHeader(response, "Content-Type", expectedContentType)
    contentType := response.Header().Get("Content-Type")
    if contentType != expectedContentType {
        t.Errorf("expected Content-Type header to be %s but got %s", expectedContentType, contentType)
    }
}
