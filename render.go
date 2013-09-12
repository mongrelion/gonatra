package gonatra

import (
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
