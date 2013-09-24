package gonatra

import (
    "net/http"
)

type Params map[string]string

func getParams(route *Route, request *http.Request) Params {
    params := make(Params)
    // Params from query string and form.
    request.ParseForm()
    for param, values := range request.Form {
        params[param] = values[0]
    }

    // Named params, specified in the route declaration
    pathMatches := pathRegexp.FindAllString(route.Path, -1)
    urlMatches  := pathRegexp.FindAllString(request.URL.Path, -1)
    for i, paramName := range pathMatches {
        if paramRegexp.MatchString(paramName) {
            param         := paramNameRegexp.FindString(paramName)
            params[param]  = urlMatches[i]
        }
    }

    return params
}
