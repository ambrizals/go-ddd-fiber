package testutils

import (
	"testing"

	"github.com/ambrizals/go-ddd-fiber/config"
	"github.com/ambrizals/go-ddd-fiber/infra/db"
	"github.com/stretchr/testify/suite"
)

type TestSuite struct {
	suite.Suite
	DBClient  *db.Client
	EnvConfig *config.EnvConfig
}

func (suite *TestSuite) SetupSuite() {
	// Use the existing config package to load environment variables
	var err error
	suite.EnvConfig, err = config.LoadEnv()
	if err != nil {
		suite.FailNowf("Failed to load environment variables", "error: %v", err)
	}

	// Initialize database client using the existing configuration
	suite.DBClient, err = db.GetClient()
	if err != nil {
		suite.FailNowf("Failed to connect to database", "error: %v", err)
	}
}

func (suite *TestSuite) TearDownSuite() {
	// Clean up resources if needed
	if suite.DBClient != nil {
		sqlDB, err := suite.DBClient.DB.DB()
		if err != nil {
			suite.T().Errorf("Failed to get database instance: %v", err)
			return
		}
		if err := sqlDB.Close(); err != nil {
			suite.T().Errorf("Failed to close database connection: %v", err)
		}
	}
}

func TestRunSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}
