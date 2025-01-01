# GO - RouterUtils

A package that provides some small helper functions, that can be used with go 1.22 `http.ServeMux` definitions.

---

## Route - Definition helper

Function that makes defining routes and their handlers easier.

---

### `AppendListToMux`

```go
func AppendListToMux(*http.ServeMux, *RouteDefinitionList, MiddlewareFunc)
```

this function allow for an easy definition of multiple Routes and Verbs.

In addition all defined routes can be wrapped into a Middware `http.HandlerFunc`.

For an Example see the [`RouteDefinitionList`](#RouteDefinitionList) Type.

---

### `StatCatResponder`

A ServeMux-HandlerFunc, that combines multiple static files into one response.
This function is meant to reduce the amout of calls, the client
needs to make for calling multiple static Javascript or CSS files.

```go
func StatCatResponder(contentType string, filenames ...string) http.HandlerFunc
```

Example:

```go
func DefineRoutes(mux *http.ServeMux) {

	mux.HandleFunc("GET /app.js", go_routerutils.StatCatResponder("text/javascript",
		"./static/js/htmx.js",
		"./static/js/jquery.js",
		"./static/js/main.js",
	))

	mux.HandleFunc("GET /app.css", go_routerutils.StatCatResponder("text/css",
		"./static/css/bootstrap.css",
		"./static/css/main.css",
	))

}
```

Uppon requesting `GET /app.js`, the response will be of `content-type: text/javascript`
with the content of the 3 given files.

---

## Response Generators

Functions that can generate various formated HTTP-Response Messages

---

### `RespondWithCode`

```go
func RespondWithCode(http.ResponseWriter, status int, msg string, args ...any)
```

sets the Response-Status to the given `status` and writes `msg` and `args` into the response body.
the response body is generated via `fmt.Sprintf`.

Great if all that matters is the Response Status-Code.

example:

```go
func handleRoute(w http.ResponseWriter, r *http.Request) {

	go_routerutils.RespondWithCode(w, 200, "response for: %s", r.URL);

}
```

---

### `RespondWithError`

```go
func RespondWithError(w http.ResponseWriter, status int, err error) bool
```

This function does 3 things if an error is given.

1. It generates an HTTP-Response with the given `status` and the error as Response-Body.
2. If the given status is >= 500, a stacktrace is written to stdout.
3. The function responds with a `bool`, which makes it easy to use in Error-Validation.

Example:

```go
func handleRoute(w http.ResponseWriter, r *http.Request) {

	_, err := someFunctionThatGeneratesAnError()

	if go_routerutils.RespondWithError(w, 500, err) {
		// end if an error was generated
		return
	}

	// If err is `nil`, continue here
	// ... do more stuff
}
```

---

### `RespondWithJSON`

```go
func RespondWithJSON(w http.ResponseWriter, status int, data any)
```

Takes any input `data` and converts it into JSON-Format, according to Go's `json.Marshal` function.
The response content-type will be set to `application/json`.

Example:

```go
type Person struct{
	id int,
	name any,
}
func handleRoute(w http.ResponseWriter, r *http.Request) {

	person := Person{
		id: 1,
		data: "Max Mustermann",
	}

	go_routerutils.RespondWithJSON(w, 200, person);

}
```

---

## Body Parsing Helpers
Functions to make reading request bodies easier.

---

### `ReadBodyFromRequest`

```go
func ReadBodyFromRequest(r *http.Request) (body []byte, err error)
```

Parses the the entire RequestBody into a Slice of bytes `[]byte`.

Example:

```go
func handleRoute(w http.ResponseWriter, r *http.Request) {

	body, err := go_routerutils.ReadBodyFromRequest(r)
	if go_routerutils.RespondWithError(w, 500, err) {
		return
	}

	go_routerutils.RespondWithCode(w, 200, "your Body was: %s\n", body)

}
```

---

### `ReadJSONBodyFromRequest`
```go
func ReadJSONBodyFromRequest(r *http.Request) (output map[string]interface{}, err error)
```
Parses the entire Request-Body into a JSON-Object.
This function by itself is kind of useless. Use the next function `AssertJSONFieldType` function
to make use of the response of this function.

---

### `AssertJSONFieldType`
```go
func AssertJSONFieldType[T any](mp map[string]interface{}, key string) (out T, err error) {
```
This takes the result of [`ReadJSONBodyFromRequest`](#readjsonbodyfromrequest) and extracts
a single field from it.

Example:
```go
func handleRoute(w http.ResponseWriter, r *http.Request) {

	var err error
	json, err := go_routerutils.ReadJSONBodyFromRequest(r)
	if go_routerutils.RespondWithError(w, 500 /* Internal Server error */, err) {
		return
	}

	// Extract the field `yourname` from request-json. Expect it to be a string
	var name string
	name, err = go_routerutils.AssertJSONFieldType[string](json, "yourname");
	if go_routerutils.RespondWithError(w, 400 /* Bad Request */, err) {
		return
	}

	go_routerutils.RespondWithCode(w, 200, "Hello, %s°", name)

}
```


---

## Types

### RouteDefinitionList

```go
type RouteDefinitionList map[string]map[string]http.HandlerFunc
```

A List of Route Definitions that is passed to the [`AppendListToMux`](#_AppendListToMux_) function.
The list is a 2 dimensional Array.

- the first layer index is the `HTTP-Method` / `HTTP-Verb`
- the second layer index is the actual route
- the value is a `http.HandlerFunc` that is invoced, if the route is called

Example:

```go
	go_routerutils.AppendListToMux(mux, &go_routerutils.RouteDefinitionList{
		"GET": {
			"/notegroups":             getNoteGroups,
			"/notegroups/{grp}/notes": getNotesForGroup,
		},

		"POST": {
			"/notegroups": postNoteGroup,
		},

		"PATCH": {
			"/notegroups/{grp}/move/{location}/{targetGrp}": patchMoveNoteGroup,
		},

		"DELETE": {
			"/notegroups/{grp}": deleteNoteGroup,
		},
	}, nil)
```

```go
type MiddlewareFunc func(http.HandlerFunc)http.HandlerFunc
```
