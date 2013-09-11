package gonatra

import (
    "net/http"
    "log"
    "regexp"
    "sync"
)

type Request struct {
    HttpRequest *http.Request
    Params      map[string][]string
}

type Route struct {
    Path     string
    Verb     string
    Callback func(response http.ResponseWriter, request *Request)
    Rgxp     *regexp.Regexp
}

const (
    HTTP_GET    = "GET"
    HTTP_POST   = "POST"
    HTTP_PUT    = "PUT"
    HTTP_DELETE = "DELETE"
)

var (
    paramRegexp     = regexp.MustCompile(":[a-zA-Z0-9_]+")
    pathRegexp      = regexp.MustCompile(":?[a-zA-Z0-9_]+")
    paramNameRegexp = regexp.MustCompile("[a-zA-Z0-9_]+")
    validVerbs      = []string{HTTP_GET, HTTP_POST, HTTP_PUT, HTTP_DELETE}
    routes          = make([]Route, 0, 0)
    session         = struct{
        sync.RWMutex
        m map[string]string
    }{m: make(map[string]string)}
)

func init() {
    http.HandleFunc("/", dispatcher)
}

func getParams(route *Route, req *http.Request) map[string][]string {
    params      := make(map[string][]string)
    // Named params, specified in the route declaration
    pathMatches := pathRegexp.FindAllString(route.Path, -1)
    urlMatches  := pathRegexp.FindAllString(req.URL.Path, -1)
    for i, paramName := range pathMatches {
        if (paramRegexp.MatchString(paramName)) {
            param         := paramNameRegexp.FindString(paramName)
            params[param]  = []string{urlMatches[i]}
        }
    }

    // Params from query string and form.
    req.ParseForm()
    for param, values := range req.Form {
        params[param] = values
    }
    return params
}

func buildRequest(httpReq *http.Request, route *Route) Request {
    params := getParams(route, httpReq)
    return Request{httpReq, params}
}

func dispatcher(res http.ResponseWriter, req *http.Request) {
    for _, route := range routes {
        if (matchRoute(&route, req.URL.Path)) {
            if (route.Verb == req.Method) {
                req.ParseForm()
                request := buildRequest(req, &route)
                route.Callback(res, &request)
                return
            }
        }
    }
    http.NotFound(res, req)
}

func genRouteRegexp(route string) *regexp.Regexp {
    return regexp.MustCompile(paramRegexp.ReplaceAllString(route, ".+"))
}

func matchRoute(route *Route, path string) bool {
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

func RegisterRoute(verb, path string, callback func(res http.ResponseWriter, req *Request)) bool {
    if ValidVerb(verb) {
        rgxp  := genRouteRegexp(path)
        route := Route{path, verb, callback, rgxp}
        routes = append(routes, route)
        return true
    } else {
        return false
    }
}

func Get(path string, callback func(res http.ResponseWriter, req *Request)) bool {
    return RegisterRoute(HTTP_GET, path, callback)
}

func Post(path string, callback func(res http.ResponseWriter, req *Request)) bool {
    return RegisterRoute(HTTP_POST, path, callback)
}

func RenderText(res http.ResponseWriter, str string) {
    res.Write([]byte(str))
}

func RenderJson(res http.ResponseWriter, str string) {
    res.Header().Set("Content-Type", "application/json")
    RenderText(res, str)
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
