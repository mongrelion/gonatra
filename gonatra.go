package gonatra

import (
    "fmt"
    "net/http"
    "log"
    "regexp"
    "sync"
)

const (
    HTTP_GET                = "GET"
    HTTP_POST               = "POST"
    HTTP_PUT                = "PUT"
    HTTP_DELETE             = "DELETE"
    CONTENT_TYPE_JSON       = "application/json"
    CONTENT_TYPE_PLAIN_TEXT = "text/plain"
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

func setHeader(response http.ResponseWriter, key, value string) {
    response.Header().Set(key, value)
}

func getParams(route *Route, req *http.Request) map[string][]string {
    params := make(map[string][]string)
    // Params from query string and form.
    req.ParseForm()
    for param, values := range req.Form {
        params[param] = values
    }

    // Named params, specified in the route declaration
    pathMatches := pathRegexp.FindAllString(route.Path, -1)
    urlMatches  := pathRegexp.FindAllString(req.URL.Path, -1)
    for i, paramName := range pathMatches {
        if paramRegexp.MatchString(paramName) {
            param         := paramNameRegexp.FindString(paramName)
            params[param]  = []string{urlMatches[i]}
        }
    }

    return params
}

func buildRequest(httpReq *http.Request, route *Route) Request {
    params := getParams(route, httpReq)
    return Request{httpReq, params}
}

func dispatcher(response http.ResponseWriter, request *http.Request) {
    for _, route := range routes {
        if matchRoute(&route, request.URL.Path) {
            if route.Verb == request.Method {
                request.ParseForm()
                gonatraRequest := buildRequest(request, &route)
                route.Callback(response, &gonatraRequest)
                return
            }
        }
    }
    http.NotFound(response, request)
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
        routes = append(routes, route)
        return true
    } else {
        return false
    }
}

func Get(path string, callback func(res http.ResponseWriter, req *Request)) bool {
    return registerRoute(HTTP_GET, path, callback)
}

func Post(path string, callback func(res http.ResponseWriter, req *Request)) bool {
    return registerRoute(HTTP_POST, path, callback)
}

func RenderText(response http.ResponseWriter, str string) {
    setHeader(response, "Content-Type", CONTENT_TYPE_PLAIN_TEXT)
    fmt.Fprint(response, str)
}

// From http://golang.org/doc/articles/json_and_go.html:
// "The json package only accesses the exported fields of struct types
// (those that begin with an uppercase letter). Therefore only the exported
// fields of a struct will be present in the JSON output."
// Using Go structs, your JSON output should look something like this:
// {"Id":123,"Name":"John Doe","Email":"john@doe.com"}
// (notice the key names starting with an uppercase letter).
// So, instead of receiving a struct and calling json.Marshal() on it, it's up
// to you to send your already built JSON object as a string to this method.
func RenderJSON(response http.ResponseWriter, str string) {
    response.Header().Set("Content-Type", CONTENT_TYPE_JSON)
    fmt.Fprint(response, str)
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
