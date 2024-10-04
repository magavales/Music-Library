package database

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/magavales/Music-Library/todo/internal/models"
	"github.com/magavales/Music-Library/todo/internal/models/configs"
	"github.com/sirupsen/logrus"
)

type PostgresDB struct {
	pool *pgxpool.Pool
}

func NewPostgresDB(config configs.DatabaseConfig) *PostgresDB {
	pool := Connect(config)
	return &PostgresDB{
		pool: pool,
	}
}

func Connect(config configs.DatabaseConfig) *pgxpool.Pool {
	var (
		err  error
		conf *pgxpool.Config
		pool *pgxpool.Pool
	)
	err = config.Parse()
	if err != nil {
		return pool
	}

	strConfig := config.String()
	conf, err = pgxpool.ParseConfig(strConfig)
	if err != nil {
		return pool
	}

	pool, err = pgxpool.NewWithConfig(context.Background(), conf)
	if err != nil {
		return pool
	}

	return pool
}

func (pdb *PostgresDB) Create(context *gin.Context, song models.Song) (int64, error) {
	var id int64

	rows, err := pdb.pool.Query(context, "INSERT INTO music_library (group_name, song_name, release_date, text, link) VALUES ($1, $2, $3, $4, $5) RETURNING id", song.GroupName, song.SongName, song.ReleaseDate, song.Text, song.Link)
	if err != nil {
		logrus.Infof("Error inserting row: %v", err)
		return 0, err
	}
	if rows.Next() {
		values, err := rows.Values()
		if err != nil {
			logrus.Infof("Error iterating rows: %v", err)
			return 0, err
		}
		id = values[0].(int64)
	} else {
		logrus.Infof("Error iterating rows: %v", err)
		return 0, pgx.ErrNoRows
	}

	return id, err
}

func (pdb *PostgresDB) Get(context *gin.Context, sql string) ([]models.Song, error) {
	var (
		song  models.Song
		songs []models.Song
	)

	rows, err := pdb.pool.Query(context, sql)
	if err != nil {
		logrus.Infof("Error selecting row: %v", err)
		return []models.Song{}, err
	}
	for rows.Next() {
		values, err := rows.Values()
		if err != nil {
			logrus.Infof("Error iterating rows: %v", err)
			return []models.Song{}, err
		}
		song.ParseRows(values)
		songs = append(songs, song)
	}
	if err = rows.Err(); err != nil {
		logrus.Infof("Error iterating rows: %v", err)
		return songs, pgx.ErrNoRows
	}

	return songs, err
}

func (pdb *PostgresDB) GetByID(context *gin.Context, id int) (models.Song, error) {
	var song models.Song

	rows, err := pdb.pool.Query(context, "SELECT * FROM music_library WHERE id=$1", id)
	if err != nil {
		logrus.Infof("Error selecting row: %v", err)
		return models.Song{}, err
	}
	if rows.Next() {
		values, err := rows.Values()
		if err != nil {
			logrus.Infof("Error iterating rows: %v", err)
			return models.Song{}, err
		}
		song.ParseRows(values)
	} else {
		logrus.Infof("Error iterating rows: %v", err)
		return models.Song{}, pgx.ErrNoRows
	}

	return song, err
}

func (pdb *PostgresDB) Delete(context *gin.Context, id int) error {
	ans, err := pdb.pool.Exec(context, "DELETE FROM music_library WHERE id=$1", id)
	if err != nil {
		logrus.Infof("Error deleting row: %v", err)
		return err
	}

	if ans.String() == "DELETE 0" {
		logrus.Infof("Error deleting row: no rows were deleted")
		return pgx.ErrNoRows
	}

	return err
}

func (pdb *PostgresDB) Update(context *gin.Context, songNew models.Song, id int) error {
	ans, err := pdb.pool.Exec(context, "UPDATE music_library SET group_name=$2, song_name=$3, release_date=$4, text=$5, link=$6 WHERE id=$1", id, songNew.GroupName, songNew.SongName, songNew.ReleaseDate, songNew.Text, songNew.Link)
	if err != nil {
		logrus.Infof("Error updating row: %v", err)
		return err
	}

	if ans.String() == "UPDATE 0" {
		logrus.Infof("Error updating row: no rows were updated")
		return pgx.ErrNoRows
	}

	return err
}
