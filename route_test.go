package gonatra

import(
    "net/http"
    "testing"
)

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
