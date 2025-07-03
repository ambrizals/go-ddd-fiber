package get_hello_world

type GetHelloWorldResponse struct {
	Body struct {
		Message string `json:"message" example:"Hello, world!" doc:"Greeting message"`
	}
}
