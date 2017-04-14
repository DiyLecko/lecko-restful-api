package example

import (
	"github.com/DiyLecko/lecko-restful-api/restapi"
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
func (resource examResource) Get(rp restapi.RestParam) restapi.Response {
	result := map[string]interface{}{
		"a": 1,
		"b": 2,
	}

	// restapi.Response{statusCode, message, data}
	return restapi.Response{200, "message", result}
}

func (resource examResource) Post(rp restapi.RestParam) restapi.Response {
	result := "\"requiredField\" field value is " + rp.Request.FormValue("requiredField")
	return restapi.Response{200, "message", result}
}

// "requiredField" field is required.
var examResourcePostRequired = []string{
	"requiredField",
}

func (examResource) PostRequired(rp restapi.RestParam) bool {
	if rp.Request != nil {
		for _, arg := range examResourcePostRequired {
			if rp.Request.FormValue(arg) == "" {
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

func (resource examIdResource) Get(rp restapi.RestParam) restapi.Response {
	result := map[string]interface{}{
		"id": rp.Params.ByName("id"),
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

	// Now you can access "[Get,Post]localhost:3000/api/v1/exam" and "[Get]localhost:3000/api/v1/exam/123"
	// If you access "[Put,Delete]/api/v1/exam" and "[Post,Put,Delete]/api/v1/exam/123", then restapi will call restapi.___NotSupported
}
