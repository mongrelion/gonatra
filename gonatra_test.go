package gonatra

import "testing"

func TestValidVerb(t *testing.T) {
    validVerbs   := []string{"GET", "POST", "PUT", "DELETE"}
    invalidVerbs := []string{"W00T", "PATCH", "INVALID", "JOE"}

    for _, verb := range validVerbs {
        if (!ValidVerb(verb)) {
            t.Errorf("expected %d verb to be valid but was invalid.")
        }
    }

    for _, verb := range invalidVerbs {
        if (ValidVerb(verb)) {
            t.Errorf("expected %d verb to be invalid but was valid.")
        }
    }
}

func TestRegisterRoute(t *testing.T) {
    route         := Route{"/foo", HTTP_GET, func() {}}
    result        := RegisterRoute(HTTP_GET, route.Path, route.Callback)

    // Test that it returns true given a valid path, verb and callback.
    if (!result) {
        t.Errorf("expected RegisterRoute() to return %t but got %t", !result, result)
    }

    // Test that it appends a new route to the "routes" array.
    if (len(routes) == 0) {
        t.Errorf("expected 1 route to be registered but got 0.")
    }

    // Test that the registered route matches everything
    lastRoute := routes[len(routes) -1]
    if (lastRoute.Path != route.Path) {
        t.Errorf("expected route path to be %s but got %s", route.Path, lastRoute.Path)
    }
    if (lastRoute.Verb != route.Verb) {
        t.Errorf("expected route verb to be %s but got %s", route.Verb, lastRoute.Verb)
    }
}

func TestGet(t *testing.T) {
    path      := "/foo"
    result    := Get(path, func() {})
    lastRoute := routes[len(routes) -1]

    // Test that the Get method returns true
    if (!result) {
        t.Errorf("expected Get() to return %t but got %t", !result, result)
    }

    // Test that the given path was set properly
    if (lastRoute.Path != path) {
        t.Errorf("expected path '%s' but got '%s'", path, lastRoute.Path)
    }

    // Test that the registered verb for the route is GET.
    if (lastRoute.Verb != HTTP_GET) {
        t.Errorf("expected HTTP verb to be %s but got %s", HTTP_GET, lastRoute.Verb)
    }
}

func TestPost(t *testing.T) {
    path      := "/foo"
    result    := Post(path, func() {})
    lastRoute := routes[len(routes) -1]

    // Test that the Post method returns true
    if (!result) {
        t.Errorf("expected Post() to return %t but got %t", !result, result)
    }

    // Test that the given path was set properly
    if (lastRoute.Path != path) {
        t.Errorf("expected path '%s' but got '%s'", path, lastRoute.Path)
    }

    // Test that the registered verb for the route is GET.
    if (lastRoute.Verb != HTTP_POST) {
        t.Errorf("expected HTTP verb to be %s but got %s", HTTP_POST, lastRoute.Verb)
    }
}
