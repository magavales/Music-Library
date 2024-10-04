package database

import (
	"github.com/gin-gonic/gin"
	"github.com/magavales/Music-Library/todo/internal/models"
	"github.com/magavales/Music-Library/todo/internal/models/configs"
)

type Postgres interface {
	Create(context *gin.Context, song models.Song) (int64, error)
	Get(context *gin.Context, sql string) ([]models.Song, error)
	GetByID(context *gin.Context, id int) (models.Song, error)
	Delete(context *gin.Context, id int) error
	Update(context *gin.Context, songNew models.Song, id int) error
}

type Database struct {
	Postgres
}

func NewDatabase(config configs.DatabaseConfig) *Database {
	return &Database{
		Postgres: NewPostgresDB(config),
	}
}
