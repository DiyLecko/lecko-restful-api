package restapi

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type API struct {
	Router *httprouter.Router
	IsCORS bool
}

func Init() *API {
	instance := new(API)
	instance.Router = httprouter.New()
	instance.IsCORS = true
	return instance
}

func (api *API) Start(port string) {
	portString := ":" + port

	fmt.Println("Lecko RESTful API server is launched, PORT is " + port)

	http.ListenAndServe(portString, api.Router)
}

func (api *API) AddResource(resource Resource) {
	fmt.Println("\"" + resource.Uri() + "\" api is registerd")

	api.Router.GET(resource.Uri(), func(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		res := resource.Get(rw, r, ps)
		api.Response(rw, r, res)
	})
	api.Router.POST(resource.Uri(), func(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		res := resource.Post(rw, r, ps)
		api.Response(rw, r, res)
	})
	api.Router.PUT(resource.Uri(), func(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		res := resource.Put(rw, r, ps)
		api.Response(rw, r, res)
	})
	api.Router.DELETE(resource.Uri(), func(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		res := resource.Delete(rw, r, ps)
		api.Response(rw, r, res)
	})
	api.Router.OPTIONS(resource.Uri(), func(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		api.Response(rw, r, Response{200, "", nil})
	})
}

func (api *API) Response(rw http.ResponseWriter, req *http.Request, res Response) {
	content, err := json.Marshal(res)

	if err != nil {
		abort(rw, 500)
	}

	if api.IsCORS == true {
		if origin := req.Header.Get("Origin"); origin != "" {
			rw.Header().Set("Access-Control-Allow-Origin", origin)
			rw.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			rw.Header().Set("Access-Control-Allow-Headers",
				"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		} else {
			rw.Header().Set("Access-Control-Allow-Origin", "*")
			rw.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			rw.Header().Set("Access-Control-Allow-Headers",
				"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		}

		if req.Method == "OPTIONS" {
			return
		}
	}

	rw.WriteHeader(res.Code)
	rw.Write(content)
}

func abort(rw http.ResponseWriter, statusCode int) {
	rw.WriteHeader(statusCode)
}
