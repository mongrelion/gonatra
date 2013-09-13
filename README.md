Gonatra
=======
Lightweight web framework on top of [Go]'s [net/http] package.

---

#What does it offer?
It offers a simple routes management API:

```
gonatra.Get(path string, callback func(res http.ResponseWriter, req *gonatra.Request))
gonatra.Post(path string, callback func(res http.ResponseWriter, req *gonatra.Request))

```

For managing session, use:

```
gonatra.Session.Get(key string) string
gonatra.Session.Set(key, value string)

```

---

#Handlers

These are the functions that are going to be called once the Gonatra's dispatcher has found a matching route for the current path and HTTP method.  
Your handlers must offer the following signature:

```
func Handler(res http.ResponseWriter, req *gonatra.Request) {
	// Body of the handler here.
}
```
The [http.ResponseWriter] parameter contains information such as Status Code, Content Length, Headers and much more. If you want more insight on this you can go and check the documentation. It's pretty well explained.  
The gonatra.Request parameter encapsulates two main objects: a [http.Request] (which contains information like Method, URL, Body, Host, etc) and a gonatra.Params map, which has a map containing the parameters extracted from the URL, from the query string and from the POSTed form (if the case applies).

---

#Registering handlers

For registering handlers simply specify which verb you want to respond to, on which path and which function should be called when the route is matched:

```
// Give it the name you want.
func FooHandler(res http.ResponseWriter, req *gonatra.Request) {
	// Do something here.
}

func KatzHandler(res http.ResponseWriter, req *gonatra.Request) {
	// Do something here.
}

// Then register it to some verb.
Get("/foo/bar", FooHandler)
Post("/lolz/:id/katz", KatzHandler)

```

---

#Accessing params
In your handler signature you specified receiving a second argument, a request of type gonatra.Request. We previously mentioned that this object encapsulates both a http.Request and a map with parameters. This is how you can access it:

```
// Say you registered a route /albums/:id/photos
// To access the :id param you should do something like this:
func Handler(res http.ResponseWriter, req *gonatra.Request) {
	albumId := req.Params["id"][0]
	// Whatever.
}

```

You can also, of course, access parameters specified in the query string:

```
// Let's say that the requested path is /albums/123?foo=bar&lol=haha
func Handler(res http.ResponseWriter, req *gonatra.Request) {
	req.Params["id"][0]  == "123"  // => true
	req.Params["foo"][0] == "bar"  // => true
	req.Params["lol"][0] == "haha" // => true
}

```
You should access parameters set on your request BODY (POSTed by a form or whatever) the same way.

---

#Rendering text
For rendering text simply call the gonatra.RenderText() method and pass both the response and the text you want to render:

```
func Handler(res http.ResponseWriter, req *gonatra.Request) {
	gonatra.RenderText(res, "Stop! Hammer time!")
}

```

---

#Rendering JSON
For rendering JSON you simply have to pass both the response and an object that can be serialised by the [json.Marshal()](http://golang.org/pkg/encoding/json/#Marshal) function. Take into account that when using user defined structs, all fields must be publicly accessible, that is, field names must start with a uppercase letter if you want it to be present in its JSON representation (for more information about this check this [article](http://carlosleon.info/articles/generating-json-in-go)). Let's look at an example:

```
type Fruit struct {
	Name  string `json:"name"`
	Color string `json:"color"`
}

func Handler(res http.ResponseWriter, req *gonatra.Request) {
	fruit := Fruit{"apple", "red"}
	gonatra.RenderJSON(res, fruit)
}

```

When you try this out, you will get this output: ```{"name":"apple","colour":"red"}```

---

#Working example

```
package main

import (
	"fmt"
	"github.com/mongrelion/gonatra"
	"net/http"
)

// Create your handlers
func Greeter(res http.ResponseWriter, req *gonatra.Request) {
    name  := req.Params["name"][0]
    greet := fmt.Fprint("Hi there, %s", name)
	gonatra.RenderText(res, greet)
}

func main() {
	// Register that handler
	gonatra.Get("/hello/:name", Greeter)

	// You can also register anonymous functions
	gonatra.Get("/lolz", func(res http.ResponseWriter, req *gonatra.Request) {
		gonatra.RenderText(res, "icanhazcheeseburger")
	})

	// Finally, start listening on some port
	gonatra.Start(":9292")
}

```

Try it out in console

```
$ curl http://localhost:9292/hello/John
Hi there, John
$ curl http://localhost:9292/lolz
icanhazcheeseburger

```

---

###The MIT License (MIT)

Copyright (c) 2013 Carlos León - http://carlosleon.info

Permission is hereby granted, free of charge, to any person obtaining a copy of
this software and associated documentation files (the "Software"), to deal in
the Software without restriction, including without limitation the rights to
use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
the Software, and to permit persons to whom the Software is furnished to do so,
subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
FOR A PARTICULAR PURPOSE AND NON INFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

[Go]: http://golang.org/
[net/http]: http://golang.org/pkg/net/http/
[http.ResponseWriter]: http://golang.org/pkg/net/http/#Response
[http.Request]: http://golang.org/pkg/net/http/#Request
[JSON]: http://json.org/
[encoding/json]: http://golang.org/pkg/encoding/json/