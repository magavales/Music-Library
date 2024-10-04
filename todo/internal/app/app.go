package app

import (
	"github.com/magavales/Music-Library/todo/internal/database"
	"github.com/magavales/Music-Library/todo/internal/models/configs"
	"github.com/magavales/Music-Library/todo/internal/services"
	"github.com/magavales/Music-Library/todo/internal/transport/handler"
	"github.com/magavales/Music-Library/todo/internal/transport/httpserver"
)

func Run() {
	var (
		err    error
		config configs.DatabaseConfig
	)
	err = config.Parse()
	if err != nil {
		panic(err)
	}
	db := database.NewDatabase(config)
	service := services.NewService(db)
	h := handler.NewHandler(service)
	server := httpserver.New()
	err = server.Run(h.InitRouter())
	if err != nil {
		panic(err)
	}
}
