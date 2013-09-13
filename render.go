package gonatra

import (
    "encoding/json"
    "fmt"
    "net/http"
)

func setHeader(response http.ResponseWriter, key, value string) {
    response.Header().Set(key, value)
}

func RenderText(response http.ResponseWriter, str string) {
    setHeader(response, "Content-Type", CONTENT_TYPE_PLAIN_TEXT)
    fmt.Fprint(response, str)
}

// Receives a http.ResponseWriter and an object, serialising it with json.Marshal()
// An empty JSON object will be rendered if the object couldn't be serialised.
// TODO: Return something!
func RenderJSON(response http.ResponseWriter, obj interface{}) {
    response.Header().Set("Content-Type", CONTENT_TYPE_JSON)
    jObj, err := json.Marshal(obj)
    if err != nil {
        jObj = []byte("{}")
    }
    response.Write(jObj)
}
