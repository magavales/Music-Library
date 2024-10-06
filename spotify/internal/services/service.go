package services

import (
	"github.com/gin-gonic/gin"
	"github.com/magavales/Music-Library/todo/internal/database"
)

type LibraryRepository interface {
	Create(context *gin.Context)
	Get(context *gin.Context)
	GetByID(context *gin.Context)
	GetText(context *gin.Context)
	Delete(context *gin.Context)
	Update(context *gin.Context)
}

type Service struct {
	LibraryRepository
}

func NewService(db *database.Database) *Service {
	return &Service{
		LibraryRepository: NewMusicLibrary(db.Postgres),
	}
}
