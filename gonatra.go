package gonatra

import (
    "net/http"
    "log"
)

type Route struct {
    Path string
    Verb string
    Callback func(response http.ResponseWriter, request *http.Request)
}

const (
    HTTP_GET    = "GET"
    HTTP_POST   = "POST"
    HTTP_PUT    = "PUT"
    HTTP_DELETE = "DELETE"
)

var (
    validVerbs []string = []string{HTTP_GET, HTTP_POST, HTTP_PUT, HTTP_DELETE}
    routes     []Route  = make([]Route, 0, 0)
)

func ValidVerb(verb string) bool {
    for _, validVerb := range validVerbs {
        if (verb == validVerb) {
            return true
        }
    }
    return false
}

func RegisterRoute(verb, path string, callback func(res http.ResponseWriter, req *http.Request)) bool {
    if ValidVerb(verb) {
        route := Route{path, verb, callback}
        routes = append(routes, route)
        http.HandleFunc(route.Path, func(response http.ResponseWriter, request *http.Request) {
            if (route.Verb == request.Method) {
                route.Callback(response, request)
            } else {
                http.NotFound(response, request)
            }
        })
        return true
    } else {
        return false
    }
}

func Get(path string, callback func(res http.ResponseWriter, req *http.Request)) bool {
    return RegisterRoute(HTTP_GET, path, callback)
}

func Post(path string, callback func(res http.ResponseWriter, req *http.Request)) bool {
    return RegisterRoute(HTTP_POST, path, callback)
}

func RenderText(res http.ResponseWriter, str string) (int, error) {
    return res.Write([]byte(str))
}

func Start(port string) {
    log.Fatal(http.ListenAndServe(port, nil))
}
