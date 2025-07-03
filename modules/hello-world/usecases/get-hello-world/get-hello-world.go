package get_hello_world

import "github.com/ambrizals/go-ddd-fiber/modules/hello-world/repository"

func Execute(helloWorldRepository repository.HelloWorldRepository) (GetHelloWorldResponse, error) {
	res := helloWorldRepository.GetHelloWorld()

	resp := GetHelloWorldResponse{}
	resp.Body.Message = res
	return resp, nil
}
