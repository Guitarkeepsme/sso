package main

import (
	"fmt"
	"sso/internal/config"
)

func main() {
	cfg := config.MustLoad()

	fmt.Println(cfg)
	// ToDo: инициализировать объект конфигурации

	// ToDo: инициализировать логгер

	// ToDo: запустить приложение

	// ToDo: запустить сервер
}
