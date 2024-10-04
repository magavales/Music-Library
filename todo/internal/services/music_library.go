package services

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/magavales/Music-Library/todo/internal/database"
	"github.com/magavales/Music-Library/todo/internal/models"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type MusicLibrary struct {
	db database.Postgres
}

func NewMusicLibrary(db database.Postgres) *MusicLibrary {
	return &MusicLibrary{db: db}
}

func (ml *MusicLibrary) Create(context *gin.Context) {
	var (
		err  error
		song models.Song
		id   int64
	)
	err = song.DecodeJSON(context.Request.Body)
	if err != nil {
		logrus.Errorf("Error parsing data from the query: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	id, err = ml.db.Create(context, song)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	} else {
		context.JSON(http.StatusOK, gin.H{"id": id})
		return
	}
}

func (ml *MusicLibrary) Get(context *gin.Context) {
	var (
		err   error
		songs []models.Song
		query string
	)
	url := context.Request.URL
	params := models.NewParameter()
	params.Parse(url)
	query, err = params.GetSQL()
	if err != nil {
		if err.Error() == "query doesn't contain limit and offset" || err.Error() == "no params for filtering" {
			logrus.Infof("Error with query's parameters: %v", errors.New("query is bad or query doesn't have any parameters"))
			context.JSON(http.StatusBadRequest, gin.H{"error": errors.New("query is bad or query doesn't have any parameters").Error()})
			return
		} else {
			logrus.Infof("Error creating query: %v", err)
			context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}
	songs, err = ml.db.Get(context, query)
	if err != nil {
		logrus.Infof("Error getting songs: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	} else {
		body, err := json.Marshal(songs)
		if err != nil {
			logrus.Infof("Error marshalling songs: %v", err)
			context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		} else {
			context.Writer.WriteHeader(http.StatusOK)
			context.Writer.Write(body)
		}
	}

}

func (ml *MusicLibrary) GetByID(context *gin.Context) {
	var (
		err  error
		song models.Song
		id   int
	)
	id, _ = strconv.Atoi(context.Param("id"))
	song, err = ml.db.GetByID(context, id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	body, err := json.Marshal(song)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	} else {
		context.Writer.WriteHeader(http.StatusOK)
		context.Writer.Write(body)
		return
	}
}

func (ml *MusicLibrary) GetText(context *gin.Context) {
	var (
		err  error
		id   int
		song models.Song
	)
	id, _ = strconv.Atoi(context.Param("id"))
	params := models.NewParameter()
	params.Parse(context.Request.URL)
	page, err := params.GetPage()
	if err != nil {
		logrus.Infof("Error with query's parameters: %v", err)
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	song, err = ml.db.GetByID(context, id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	text := song.PaginationForText(page)
	if text == nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": errors.New("page not found").Error()})
		return
	} else {
		context.JSON(http.StatusOK, gin.H{"text": text})
	}
}

func (ml *MusicLibrary) Delete(context *gin.Context) {
	var (
		err error
		id  int
	)
	id, _ = strconv.Atoi(context.Param("id"))
	err = ml.db.Delete(context, id)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else {
		context.JSON(http.StatusOK, gin.H{"status": "deleted"}) //???????????
		return
	}
}

func (ml *MusicLibrary) Update(context *gin.Context) {
	var (
		err     error
		id      int
		songNew models.Song
		songOld models.Song
	)
	id, _ = strconv.Atoi(context.Param("id"))
	err = songNew.DecodeJSON(context.Request.Body)
	if err != nil {
		logrus.Errorf("Error parsing data from the query: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	songOld, err = ml.db.GetByID(context, id)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	songNew.Merge(songOld)

	err = ml.db.Update(context, songNew, id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	} else {
		context.JSON(http.StatusOK, gin.H{"status": "updated"})
		return
	}
}
