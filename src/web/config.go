package web

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
)

type Config struct {
	ApplicationName string
	DbConf          DBConfig
	LogConf         LogConfig
	SqlTemplate     string
	StaticSource    string
	Property        map[string]interface{}
}

type DBConfig struct {
	DbSource string
	MaxOpen  int
	MaxIdle  int
}

type LogConfig struct {
	LogPath      string
	LogFileName  string
	LogLevel     string
	MaxAge       int64
	RotationTime int64
}

var config *Config
var once sync.Once

func NewConfig(path string) *Config {
	once.Do(func() {
		config = &Config{}
		file, _ := os.Open(path)
		defer file.Close()
		decoder := json.NewDecoder(file)
		err := decoder.Decode(config)
		if err != nil {
			fmt.Println("Error:", err)
		}
	})
	return config
}
