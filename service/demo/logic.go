package demo

type Interface interface {
	Greeter() *GreeterResponse
}

type Logic struct {
}

type GreeterResponse struct {
	Message string
}

func (l *Logic) Greeter() *GreeterResponse {
	return &GreeterResponse{
		Message: "Hello World!!",
	}
}
