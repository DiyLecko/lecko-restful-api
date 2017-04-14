package restapi

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

const (
	GET    = "GET"
	POST   = "POST"
	PUT    = "PUT"
	DELETE = "DELETE"
)

type Resource interface {
	Uri() string
	Get(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) Response
	GetRequired(r *http.Request, ps httprouter.Params) bool
	Post(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) Response
	PostRequired(r *http.Request, ps httprouter.Params) bool
	Put(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) Response
	PutRequired(r *http.Request, ps httprouter.Params) bool
	Delete(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) Response
	DeleteRequired(r *http.Request, ps httprouter.Params) bool
}

type (
	UriNotRegistered   struct{}
	GetNotSupported    struct{}
	PostNotSupported   struct{}
	PutNotSupported    struct{}
	DeleteNotSupported struct{}
)

func (UriNotRegistered) Uri() string {
	return ""
}

func (GetNotSupported) Get(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) Response {
	return Response{405, "", nil}
}

func (GetNotSupported) GetRequired(r *http.Request, ps httprouter.Params) bool {
	return true
}

func (PostNotSupported) Post(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) Response {
	return Response{405, "", nil}
}

func (PostNotSupported) PostRequired(r *http.Request, ps httprouter.Params) bool {
	return true
}

func (PutNotSupported) Put(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) Response {
	return Response{405, "", nil}
}

func (PutNotSupported) PutRequired(r *http.Request, ps httprouter.Params) bool {
	return true
}

func (DeleteNotSupported) Delete(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) Response {
	return Response{405, "", nil}
}

func (DeleteNotSupported) DeleteRequired(r *http.Request, ps httprouter.Params) bool {
	return true
}
