package hello_world

import (
	"github.com/ambrizals/go-ddd-fiber/modules/hello-world/repository"
	"github.com/ambrizals/go-ddd-fiber/modules/hello-world/usecases/get-hello-world"
)

var (
	repo = repository.NewHelloWorldRepository()
)

func GetHelloWorldUseCase() (get_hello_world.GetHelloWorldResponse, error) {
	return get_hello_world.Execute(repo)
}
