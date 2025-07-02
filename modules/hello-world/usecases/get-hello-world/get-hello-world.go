package get_hello_world

import (
	repository "github.com/ambrizals/go-ddd-fiber/modules/hello-world/repository"
)

func Execute(authRepository repository.HelloWorldRepository) (Response, error) {
	res := authRepository.GetHelloWorld()

	resp := Response{}
	resp.Body.Message = res
	return resp, nil
}
