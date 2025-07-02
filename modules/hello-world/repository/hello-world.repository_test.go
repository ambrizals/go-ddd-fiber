package repository_test

import (
	"testing"

	"github.com/ambrizals/go-ddd-fiber/modules/hello-world/repository"
	"github.com/ambrizals/go-ddd-fiber/testutils"
	"github.com/stretchr/testify/suite"
)

type HelloWorldRepositoryTestSuite struct {
	testutils.TestSuite
	authRepo repository.HelloWorldRepository
}

func (suite *HelloWorldRepositoryTestSuite) SetupTest() {
	suite.authRepo = repository.NewHelloWorldRepository()
}

func TestHelloWorldRepository(t *testing.T) {
	suite.Run(t, new(HelloWorldRepositoryTestSuite))
}

// Add your test cases here
func (suite *HelloWorldRepositoryTestSuite) TestGetHelloWorld() {
	suite.T().Run("Test GetHelloWorld", func(t *testing.T) {
		result := suite.authRepo.GetHelloWorld()
		suite.Equal("Hello, world!", result)
	})
}
