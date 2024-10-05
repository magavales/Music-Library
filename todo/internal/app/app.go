package app

import (
	"github.com/magavales/Music-Library/todo/internal/database"
	"github.com/magavales/Music-Library/todo/internal/models/configs"
	"github.com/magavales/Music-Library/todo/internal/services"
	"github.com/magavales/Music-Library/todo/internal/transport/handler"
	"github.com/magavales/Music-Library/todo/internal/transport/httpserver"
	"github.com/sirupsen/logrus"
)

func Run() {
	var (
		err    error
		config configs.DatabaseConfig
	)
	logrus.Infof("The app has been started!")
	logrus.Infof("Parsing the database config has been started!")
	err = config.Parse()
	if err != nil {
		logrus.Debugf("Parsing the database config has been failed! Error: %v", err.Error())
		panic(err)
	}
	db := database.NewDatabase(config)
	logrus.Infof("The database has been started!")
	service := services.NewService(db)
	logrus.Infof("The service has been started!")
	h := handler.NewHandler(service)
	logrus.Infof("The handler's service has been started!")
	server := httpserver.New()
	logrus.Infof("The http server has been started!")
	err = server.Run(h.InitRouter())
	if err != nil {
		logrus.Debugf("The http server has been failed! Error: %v", err.Error())
		panic(err)
	}
}
