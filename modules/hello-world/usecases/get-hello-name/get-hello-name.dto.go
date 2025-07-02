package get_hello_name

type Response struct {
	Body struct {
		Message string `json:"message" example:"Hello, world!" doc:"Greeting message"`
	}
}

type Request struct {
	Name string `path:"name" maxLength:"30" example:"world" doc:"Name to greet"`
}
