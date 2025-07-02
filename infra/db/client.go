package db

import (
	"fmt"
	"github.com/ambrizals/go-ddd-fiber/config"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	instance *Client
	once     sync.Once
)

type Client struct {
	DB *gorm.DB
}

// GetClient returns the singleton instance of the database client
func GetClient() (*Client, error) {
	var err error
	once.Do(func() {
		var envConfig *config.EnvConfig
		envConfig, err = config.LoadEnv()

		if err != nil {
			return
		}

		// Build DSN with optional password
		dsn := fmt.Sprintf("host=%s user=%s", envConfig.DatabaseHost, envConfig.DatabaseUser)
		if envConfig.DatabasePass != "" {
			dsn += fmt.Sprintf(" password=%s", envConfig.DatabasePass)
		}
		dsn += fmt.Sprintf(" dbname=%s port=%s sslmode=disable",
			envConfig.DatabaseName,
			envConfig.DatabasePort,
		)
		var db *gorm.DB
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			return
		}
		instance = &Client{DB: db}
	})

	if err != nil {
		return nil, err
	}

	return instance, nil
}

func (c *Client) Database() *gorm.DB {
	return c.DB
}
