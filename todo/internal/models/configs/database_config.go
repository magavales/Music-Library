package configs

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

type DatabaseConfig struct {
	user        string
	password    string
	host        string
	port        string
	name        string
	connections string
}

func (dc *DatabaseConfig) Parse() error {
	err := godotenv.Load("todo.env")

	dc.user = os.Getenv("DATABASE_USER")
	dc.password = os.Getenv("DATABASE_PASSWORD")
	dc.host = os.Getenv("DATABASE_HOST")
	dc.port = os.Getenv("DATABASE_PORT")
	dc.name = os.Getenv("DATABASE_NAME")
	dc.connections = os.Getenv("DATABASE_CONN")

	return err
}

func (dc *DatabaseConfig) String() string {
	return fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s pool_max_conns=%s", dc.user, dc.password, dc.host, dc.port, dc.name, dc.connections)
}
