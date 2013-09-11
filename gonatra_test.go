package gonatra

import(
    "net/http"
    "testing"
)

func TestValidVerb(t *testing.T) {
    validVerbs   := []string{"GET", "POST", "PUT", "DELETE"}
    invalidVerbs := []string{"W00T", "PATCH", "INVALID", "JOE"}

    for _, verb := range validVerbs {
        if !ValidVerb(verb) {
            t.Errorf("expected %d verb to be valid but was invalid.")
        }
    }

    for _, verb := range invalidVerbs {
        if ValidVerb(verb) {
            t.Errorf("expected %d verb to be invalid but was valid.")
        }
    }
}

func TestRegisterRoute(t *testing.T) {
    route  := Route{"/testregisterroute", HTTP_GET, func(http.ResponseWriter, *Request) {}, nil}
    result := RegisterRoute(HTTP_GET, route.Path, route.Callback)

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
    /* Skipped until I figure out how to create a new ResponseWrite from the scratch.
    str       := "test"
    sLen      := len(str)
    response  := new(http.ResponseWriter)
    bLen, err := RenderText(response, "test")
    if bLen != sLen {
        t.Errorf("%d bytes expected to be written by Render but got %d", sLen, bLen)
    }
    */
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
    // Skipped until I figure out how to create a new request and response from the scratch.
}

func TestGetParams(t *testing.T) {
    /* Skipped until I figure out how to create a request and a response from the scratch.
    path           := "/users/:id/articles/:article_id/comments/:comment_id"
    url            := "/users/123/articles/456/comments/789"
    route          := Route{path, HTTP_GET, nil, nil}
    params         := GetParams(&route, url)
    expectedParams := map[string]string{
        "id": "123",
        "article_id": "456",
        "comment_id": "789",
    }
    for key, expectedValue := range expectedParams {
        val, keyIsPresent := params[key]
        // Test that the key is present.
        if !keyIsPresent {
            t.Errorf("expected key %s to be present in params map.", key)
        } else {
            // Test that the key holds the proper value.
            if expectedValue != val {
                t.Errorf("expected key %s to have value %s but got %s", key, expectedValue, val)
            }
        }
    }
    */
}

func TestBuildRequest(t *testing.T) {
    // Test that it returns a gonatra.Reques object holding a
    // htt.Request pointer and a map of params.
    // Skipped until I figure out how to create a request and a response from the scratch.
}

func TestRenderJson(t *testing.T) {
    // Skipped until I figure out how to create a request and a response from the scratch.
}
