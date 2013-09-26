package gonatra

import (
    "strings"
    "net/http"
    "regexp"
)

type Route struct {
    Path     string
    Verb     string
    Callback func(response http.ResponseWriter, request *http.Request, params Params)
    Rgxp     *regexp.Regexp
}

func genRouteRegexp(route string) *regexp.Regexp {
    if "/" == route {
        return regexp.MustCompile("^/$")
    } else {
        return regexp.MustCompile(paramRegexp.ReplaceAllString(route, ".+"))
    }
}

func (r *Route) matchesRoute(path string) bool {
    return r.Rgxp.MatchString(path)
}

func (r *Route) matchesVerb(req *http.Request) (matches bool) {
    matches = r.Verb == req.Method;
    if !matches && req.Method == "POST" && (r.Verb == "PUT" || r.Verb == "DELETE") {
        req.ParseForm()
        if method, present := req.Form["_method"]; present {
            m := strings.ToUpper(method[0])
            matches = r.Verb == m
        }
    }
    return
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

func registerRoute(verb, path string, callback func(res http.ResponseWriter, req *http.Request, params Params)) bool {
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
