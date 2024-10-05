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
	logrus.Infof("The requset has been recieved!")
	logrus.Infof("Parsing the requst body is running!")
	err = song.DecodeJSON(context.Request.Body)
	if err != nil {
		logrus.Debugf("Error parsing data from the query: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	logrus.Infof("Parsing the requst body has been finished successfully!")

	logrus.Infof("The song is passed on to the database!")
	id, err = ml.db.Create(context, song)
	if err != nil {
		logrus.Infof("Creating the song failed with error: %v ", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	} else {
		logrus.Infof("Creating the song has been finished successfully!")
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
	logrus.Infof("The requset has been recieved!")
	url := context.Request.URL
	params := models.NewParameter()
	logrus.Infof("Parsing the query parameters is running!")
	params.Parse(url)
	logrus.Infof("Parsing the query parameters has been finished successfully!")
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
	logrus.Infof("Getting the songs from the database!")
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
			logrus.Infof("Getting the songs from the database has been finished successfully!")
			context.Writer.WriteHeader(http.StatusOK)
			context.Writer.Write(body)
		}
	}

}

func (ml *MusicLibrary) GetByID(context *gin.Context) {
	var (
		err error
		id  int
	)
	logrus.Infof("The requset has been recieved!")
	song := models.NewSong()
	id, _ = strconv.Atoi(context.Param("id"))
	logrus.Infof("Getting the song from the database!")
	song, err = ml.db.GetByID(context, id)
	if err != nil {
		logrus.Infof("Error getting songs: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	logrus.Infof("Getting the song from the database has been finished successfully!")
	body, err := json.Marshal(song)
	if err != nil {
		logrus.Infof("Error marshalling songs: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	} else {
		logrus.Infof("Getting the song from the database has been finished successfully!")
		context.Writer.WriteHeader(http.StatusOK)
		context.Writer.Write(body)
		return
	}
}

func (ml *MusicLibrary) GetText(context *gin.Context) {
	var (
		err error
		id  int
	)
	logrus.Infof("The requset has been recieved!")
	song := models.NewSong()
	id, _ = strconv.Atoi(context.Param("id"))
	params := models.NewParameter()
	logrus.Infof("Parsing the query parameters is running!")
	params.Parse(context.Request.URL)
	logrus.Infof("Parsing the query parameters has been finished successfully!")
	page, err := params.GetPage()
	if err != nil {
		logrus.Infof("Error with query's parameters: %v", err)
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	logrus.Infof("Getting the song from the database!")
	song, err = ml.db.GetByID(context, id)
	if err != nil {
		logrus.Infof("Error getting songs: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	logrus.Infof("Getting the song from the database has been finished successfully!")

	logrus.Infof("Getting the text of the song from the database!")
	text := song.PaginationForText(page)
	if text == nil {
		logrus.Infof("Error getting text: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": errors.New("page not found").Error()})
		return
	} else {
		logrus.Infof("Getting the text of the song from the database has been finished successfully!")
		context.JSON(http.StatusOK, gin.H{"text": text})
		return
	}
}

func (ml *MusicLibrary) Delete(context *gin.Context) {
	var (
		err error
		id  int
	)
	logrus.Infof("The requset has been recieved!")
	id, _ = strconv.Atoi(context.Param("id"))

	logrus.Infof("Deleting the song from the database!")
	err = ml.db.Delete(context, id)
	if err != nil {
		logrus.Infof("Error deleting song: %v", err)
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else {
		logrus.Infof("Deleting the song from the database has been finished successfully!")
		context.JSON(http.StatusOK, gin.H{"status": "deleted"}) //???????????
		return
	}
}

func (ml *MusicLibrary) Update(context *gin.Context) {
	var (
		err error
		id  int
	)
	logrus.Infof("The requset has been recieved!")
	songNew := models.NewSong()
	songOld := models.NewSong()
	id, _ = strconv.Atoi(context.Param("id"))
	logrus.Infof("Parsing the requst body is running!")
	err = songNew.DecodeJSON(context.Request.Body)
	if err != nil {
		logrus.Debugf("Error parsing data from the query: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	logrus.Infof("Parsing the requst body has been finished successfully!")

	logrus.Infof("Getting the song from the database!")
	songOld, err = ml.db.GetByID(context, id)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	logrus.Infof("Getting the song from the database has been finished successfully!")

	logrus.Infof("Merging the song from the database!")
	songNew.Merge(songOld)
	logrus.Infof("Merging the song from the database has been finished successfully!")

	logrus.Infof("Updating the song from the database!")
	err = ml.db.Update(context, songNew, id)
	if err != nil {
		logrus.Infof("Error updating song: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	} else {
		logrus.Infof("Updating the song from the database has been finished successfully!")
		context.JSON(http.StatusOK, gin.H{"status": "updated"})
		return
	}
}
