package configs

import (
	"github.com/joho/godotenv"
	"os"
)

type ServerConfig struct {
	port string
}

func (sc *ServerConfig) Parse() error {
	err := godotenv.Load("todo.env")
	sc.port = os.Getenv("SERVER_PORT")

	return err
}

func (sc *ServerConfig) Get() string {
	return sc.port
}
