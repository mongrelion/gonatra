package gonatra

import (
    "net/http"
    "log"
    "regexp"
)

const (
    CONTENT_TYPE_JSON       = "application/json"
    CONTENT_TYPE_PLAIN_TEXT = "text/plain"
    HTTP_GET                = "GET"
    HTTP_POST               = "POST"
    HTTP_PUT                = "PUT"
    HTTP_DELETE             = "DELETE"
)

var (
    paramRegexp     = regexp.MustCompile(":[a-zA-Z0-9_]+")
    pathRegexp      = regexp.MustCompile(":?[a-zA-Z0-9_]+")
    paramNameRegexp = regexp.MustCompile("[a-zA-Z0-9_]+")
    validVerbs      = []string{HTTP_GET, HTTP_POST, HTTP_PUT, HTTP_DELETE}
    routes          = make([]Route, 0, 0)
    Session         = session{m: make(map[string]string)}
)

func init() {
    http.HandleFunc("/", dispatcher)
}

// TODO: When calling the route's callback, instead of passing a gonatra.Request
//       object, send the same *http.Request object and a third argument, the params.
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

func Get(path string, callback func(res http.ResponseWriter, req *Request)) bool {
    return registerRoute(HTTP_GET, path, callback)
}

func Post(path string, callback func(res http.ResponseWriter, req *Request)) bool {
    return registerRoute(HTTP_POST, path, callback)
}

func Start(port string) {
    log.Fatal(http.ListenAndServe(port, nil))
}
