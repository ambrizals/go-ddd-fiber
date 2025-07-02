package config

import (
	"github.com/caarlos0/env"
	"os"
	"path/filepath"
	"sync"

	"github.com/joho/godotenv"
)

type EnvConfig struct {
	DatabasePort string `env:"DatabasePort" envDefault:"5432"`
	DatabaseHost string `env:"DatabaseHost" envDefault:"localhost"`
	DatabaseUser string `env:"DatabaseUser" envDefault:"postgres"`
	DatabasePass string `env:"DatabasePass" envDefault:""`
	DatabaseName string `env:"DatabaseName" envDefault:""`
}

var (
	instance *EnvConfig
	once     sync.Once
)

func findRootDir() string {
	// Try to find the root directory by looking for go.mod
	dir, _ := os.Getwd()
	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			break
		}
		dir = parent
	}
	return "."
}

func LoadEnv() (*EnvConfig, error) {
	var err error
	once.Do(func() {
		// Initialize instance first
		instance = &EnvConfig{}
		rootDir := findRootDir()
		envFile := filepath.Join(rootDir, ".env")

		// Load .env file
		if err = godotenv.Load(envFile); err != nil {
			return
		}

		// Parse environment variables into the instance
		if err = env.Parse(instance); err != nil {
			return
		}
	})

	if err != nil {
		return nil, err
	}

	return instance, nil
}
