package gonatra

import (
    "net/http"
    "log"
    "regexp"
    "sync"
)

type Route struct {
    Path     string
    Verb     string
    Callback func(response http.ResponseWriter, request *http.Request)
    Rgxp     *regexp.Regexp
}

const (
    HTTP_GET    = "GET"
    HTTP_POST   = "POST"
    HTTP_PUT    = "PUT"
    HTTP_DELETE = "DELETE"
)

var (
    paramRegexp         = regexp.MustCompile(":[a-zA-Z0-9_]+")
    pathRegexp          = regexp.MustCompile(":?[a-zA-Z0-9_]+")
    validVerbs []string = []string{HTTP_GET, HTTP_POST, HTTP_PUT, HTTP_DELETE}
    routes     []Route  = make([]Route, 0, 0)
    session             = struct{
        sync.RWMutex
        m map[string]string
    }{m: make(map[string]string)}
)

func GenRouteRegexp(route string) *regexp.Regexp {
    return regexp.MustCompile(paramRegexp.ReplaceAllString(route, ".+"))
}

func MatchRoute(route *Route, path string) bool {
    return route.Rgxp.MatchString(path)
}

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
        rgxp  := GenRouteRegexp(path)
        route := Route{path, verb, callback, rgxp}
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

/* Retrieve a value from session */
func GetSessionKey(k string) (val string) {
    session.RLock()
    val, _ = session.m[k]
    session.RUnlock()
    return
}

/* Set a value in session */
func SetSessionKey(k, v string) string {
    session.Lock()
    session.m[k] = v
    session.Unlock()
    return v
}
