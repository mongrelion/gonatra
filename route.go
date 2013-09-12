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
    return regexp.MustCompile(paramRegexp.ReplaceAllString(route, ".+"))
}

func matchRoute(route *Route, path string) bool {
    return route.Rgxp.MatchString(path)
}

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
        // TODO: Make this snippet threa-safe!
        routes = append(routes, route)
        return true
    } else {
        return false
    }
}
