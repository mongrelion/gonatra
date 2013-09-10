package gonatra

type Route struct {
    Path string
    Verb string
    Callback func()
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

func RegisterRoute(verb, path string, callback func()) bool {
    if ValidVerb(verb) {
        route := Route{path, verb, callback}
        routes = append(routes, route)
        return true
    } else {
        return false
    }
}

func Get(path string, callback func()) bool {
    return RegisterRoute(HTTP_GET, path, callback)
}

func Post(path string, callback func()) bool {
    return RegisterRoute(HTTP_POST, path, callback)
}
