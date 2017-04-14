lecko-restful-api
===================


Lecko's Simple RESTful API Server with GoLang


What is this?
-------------

I needed RESTful API Server in GoLang, but every framework and libraries are so difficult for me. So i decide that i will make simple RESTful API Server!



How to use?
-------------

1. `go get github.com/DiyLecko/lecko-restful-api`
2. `import "github.com/DiyLecko/lecko-restful-api/example`
3. `example.StartExampleRest()`
4. `Open your browser and go "http://localhost:3000/api/v1/exam"`

Also you can create new apis with following Example.



Example
--------------
```golang
package example

import (
	"net/http"

	"github.com/DiyLecko/lecko-restful-api/restapi"
	"github.com/julienschmidt/httprouter"
)

type examResource struct {
	restapi.GetNotSupported
	restapi.PostNotSupported
	restapi.PutNotSupported
	restapi.DeleteNotSupported
}

// /api/v1/exam
func (examResource) Uri() string {
	return "/api/v1/exam"
}

// You can make Get, Post, Put, Delete
func (resource examResource) Get(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) restapi.Response {
	result := map[string]interface{}{
		"a": 1,
		"b": 2,
	}

	// restapi.Response{statusCode, message, data}
	return restapi.Response{200, "message", result}
}

func (resource examResource) Post(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) restapi.Response {
	result := "\"requiredField\" field value is " + r.FormValue("requiredField")
	return restapi.Response{200, "message", result}
}

// "requiredField" field is required.
var examResourcePostRequired = []string{
	"requiredField",
}

func (examResource) PostRequired(r *http.Request, ps httprouter.Params) bool {
	if r != nil {
		for _, arg := range examResourcePostRequired {
			if r.FormValue(arg) == "" {
				return false
			}
		}
	}

	return true
}

// /api/v1/exam/:id
type examIdResource struct {
	restapi.GetNotSupported
	restapi.PostNotSupported
	restapi.PutNotSupported
	restapi.DeleteNotSupported
}

func (examIdResource) Uri() string {
	return "/api/v1/exam/:id"
}

func (resource examIdResource) Get(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) restapi.Response {
	result := map[string]interface{}{
		"id": ps.ByName("id"),
	}

	return restapi.Response{200, "message", result}
}

func StartExampleRest() {
	// Init restapi
	api := restapi.Init()
	api.IsCORS = true // default is true

	// Add examResource, examIdResource To restapi
	api.AddResource(new(examResource))
	api.AddResource(new(examIdResource))

	// Start restapi with Port 3000
	api.Start("3000")

	// Now you can access "localhost:3000/api/v1/exam" and "localhost:3000/api/v1/exam/123"
}
```



TODO
------------
1. Remove net/http and github.com/julienschmidt/httprouter in import...
2. Add examples
3. Add db support
