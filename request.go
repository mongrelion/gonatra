package gonatra

import (
    "net/http"
)

type Request struct {
    HttpRequest *http.Request
    Params      map[string][]string
}

func buildRequest(httpReq *http.Request, route *Route) Request {
    params := getParams(route, httpReq)
    return Request{httpReq, params}
}

func getParams(route *Route, request *http.Request) map[string][]string {
    params := make(map[string][]string)
    // Params from query string and form.
    request.ParseForm()
    for param, values := range request.Form {
        params[param] = values
    }

    // Named params, specified in the route declaration
    pathMatches := pathRegexp.FindAllString(route.Path, -1)
    urlMatches  := pathRegexp.FindAllString(request.URL.Path, -1)
    for i, paramName := range pathMatches {
        if paramRegexp.MatchString(paramName) {
            param         := paramNameRegexp.FindString(paramName)
            params[param]  = []string{urlMatches[i]}
        }
    }

    return params
}
