//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"project-layout/internal/config"
	"project-layout/internal/db"
)

func InitApp() (*App, error) {
	// 写法2（参考wire官方文档写法）
	wire.Build(config.Provider, db.Provider, NewApp)
	return &App{}, nil // 这里返回值没有实际意义，只需符合函数签名即可，生成的 wire_gen.go 会帮你包装该值
}
