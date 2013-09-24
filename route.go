package gonatra

import (
    "net/http"
    "regexp"
)

type Route struct {
    Path     string
    Verb     string
    Callback func(response http.ResponseWriter, request *Request)
    Rgxp     *regexp.Regexp
}

func genRouteRegexp(route string) *regexp.Regexp {
    if "/" == route {
        return regexp.MustCompile("^/$")
    } else {
        return regexp.MustCompile(paramRegexp.ReplaceAllString(route, ".+"))
    }
}

func matchRoute(route *Route, path string) bool {
    return route.Rgxp.MatchString(path)
}

// TODO: Instead of looping all over the array, use a function that returns
//       The index of the member in the array and return if the position != -1
//       That should make this method faster.
func validVerb(verb string) bool {
    for _, vVerb := range validVerbs {
        if verb == vVerb {
            return true
        }
    }
    return false
}

func registerRoute(verb, path string, callback func(res http.ResponseWriter, req *Request)) bool {
    if validVerb(verb) {
        rgxp  := genRouteRegexp(path)
        route := Route{path, verb, callback, rgxp}
        // TODO: Make this snippet thread-safe!
        routes = append(routes, route)
        return true
    } else {
        return false
    }
}
