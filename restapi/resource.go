package restapi

type Resource interface {
	Uri() string
	Get(rp RestParam) Response
	GetRequired(rp RestParam) bool
	Post(rp RestParam) Response
	PostRequired(rp RestParam) bool
	Put(rp RestParam) Response
	PutRequired(rp RestParam) bool
	Delete(rp RestParam) Response
	DeleteRequired(rp RestParam) bool
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

func (GetNotSupported) Get(rp RestParam) Response {
	return Response{405, "", nil}
}

func (GetNotSupported) GetRequired(rp RestParam) bool {
	return true
}

func (PostNotSupported) Post(rp RestParam) Response {
	return Response{405, "", nil}
}

func (PostNotSupported) PostRequired(rp RestParam) bool {
	return true
}

func (PutNotSupported) Put(rp RestParam) Response {
	return Response{405, "", nil}
}

func (PutNotSupported) PutRequired(rp RestParam) bool {
	return true
}

func (DeleteNotSupported) Delete(rp RestParam) Response {
	return Response{405, "", nil}
}

func (DeleteNotSupported) DeleteRequired(rp RestParam) bool {
	return true
}
