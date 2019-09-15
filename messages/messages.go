package messages

type Request interface {
	IsRequest()
	Action() string
	GetResponse() Response
}

type Response interface {
	IsResponse()
}
