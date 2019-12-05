package config

import (
	"fmt"
	"os"
	"sync"

	"github.com/BurntSushi/toml"
)

// Database defines database config
type Database struct {
	Name *string `toml:"name"`
	URL  *string `toml:"url"`
}

type Server struct {
	// TODO: pass context from config files
	Path *string `toml:"context-path"`
}

type Config struct {
	Db     *Database `toml:"database"`
	Server *Server   `toml:"server"`
}

var (
	confInstance *Config
	once         sync.Once
	envFlag      *string
)

var (
	// ProjectPath is the root dir of project
	ProjectPath string
)

func init() {
	gopath := os.Getenv("GOPATH")

	if len(gopath) == 0 {
		panic("GOPATH should be set!")
	}

	ProjectPath = gopath + "/src/astroshot/gin-demo"
}

// GetConfig initiates and returns global service configuration
func GetConfig(env *string) *Config {

	if env == nil {
		panic("env cannot be nil")
	}

	configFilePath := fmt.Sprintf("%s/conf/%s.toml", ProjectPath, *env)
	once.Do(func() {
		envFlag = env
		if _, err := toml.DecodeFile(configFilePath, &confInstance); err != nil {
			panic(err)
		}
	})

	return confInstance
}

// GetEnv returns running environment of server specified by command
func GetEnv() *string {
	return envFlag
}
