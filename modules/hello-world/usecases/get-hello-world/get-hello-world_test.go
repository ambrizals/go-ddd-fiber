package get_hello_world

import (
	"github.com/ambrizals/go-ddd-fiber/modules/hello-world/repository"
	"github.com/ambrizals/go-ddd-fiber/testutils"
	"github.com/stretchr/testify/suite"
	"testing"
)

type GetHelloWorldUseCaseTestSuite struct {
	testutils.TestSuite
	authRepo repository.HelloWorldRepository
}

func (suite *GetHelloWorldUseCaseTestSuite) SetupTest() {
	suite.authRepo = repository.NewHelloWorldRepository()
}

func TestGetHelloWorldUseCase(t *testing.T) {
	suite.Run(t, new(GetHelloWorldUseCaseTestSuite))
}

// Add your test cases here
func (suite *GetHelloWorldUseCaseTestSuite) TestGetHelloWorld() {
	suite.T().Run("Test GetHelloWorld", func(t *testing.T) {
		result, err := Execute(suite.authRepo)
		suite.NoError(err)
		suite.Equal("Hello, world!", result.Body.Message)
	})
}
