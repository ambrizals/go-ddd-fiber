package hello_world

import (
	"context"
	get_hello_name "github.com/ambrizals/go-ddd-fiber/modules/hello-world/usecases/get-hello-name"
	"github.com/ambrizals/go-ddd-fiber/modules/hello-world/usecases/get-hello-world"
	"github.com/danielgtaylor/huma/v2"
	"net/http"
)

func HelloWorldHandler(ctx context.Context, input *struct{}) (*get_hello_world.Response, error) {
	resp, err := GetHelloWorldUseCase()
	if err != nil {
		return nil, err
	}
	// Return a pointer to the response
	return &resp, nil
}

func HelloNameHandler(ctx context.Context, input *get_hello_name.Request) (*get_hello_name.Response, error) {
	resp := get_hello_name.Response{}
	resp.Body.Message = "Hello, " + input.Name
	return &resp, nil
}

func MakeHelloWorldHandler(humaAPI *huma.API) *huma.Group {
	group := huma.NewGroup(*humaAPI, "/v1/hello-world")

	huma.Register(group, huma.Operation{
		OperationID: "hello-world",
		Method:      http.MethodGet,
		Path:        "/hello",
		Summary:     "Hello World",
		Description: "Returns a hello world message",
		Tags:        []string{"hello-world"},
	}, HelloWorldHandler)

	huma.Register(group, huma.Operation{
		OperationID: "hello-name",
		Method:      http.MethodGet,
		Path:        "/hello/{name}",
		Summary:     "Hello Name",
		Description: "Returns a hello name message",
		Tags:        []string{"hello-world"},
	}, HelloNameHandler)

	return group
}
