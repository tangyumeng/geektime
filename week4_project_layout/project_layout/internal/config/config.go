package config

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/google/wire"
)

var Provider = wire.NewSet(New) // 将New方法声明为Provider，表示New方法可以创建一个被别人依赖的对象,也就是Config对象

var ReadConfigError = errors.New("read config file error")

type Config struct {
	Database database `json:"database"`
}

type database struct {
	Dsn string `json:"dsn"`
}

func New() (*Config, error) {
	fp, err := os.Open("../config/app.json")
	if err != nil {
		return nil, err
		// errors.Wrap(ReadConfigError, fmt.Sprintf("error: %v", err))
	}

	defer fp.Close()
	var cfg Config

	if err := json.NewDecoder(fp).Decode(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
