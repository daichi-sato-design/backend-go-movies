package models

import (
	"database/sql"
	"time"
)

// Models はデータベースのラッパーです
type Models struct{
	DB DBModel
}

// NewModels はdbプールを持つモデルを返します
func NewModels(db *sql.DB) Models{
	return Models{
		DB: DBModel{DB: db},
	}
}

// Movie は映画のタイプです
type Movie struct{
	ID			int				`json:"id"`
	Title		string			`json:"title"`
	Description	string			`json:"description"`
	Year		int				`json:"year"`
	ReleaseDate	time.Time		`json:"release_date"`
	Runtime		int				`json:"runtime"`
	Rating		int				`json:"rating"`
	MPAARating	string			`json:"mpaa_rating"`
	CreatedAd	time.Time		`json:"created_at"`
	UpdatedAt	time.Time		`json:"updated_at"`
	MovieGenre	map[int]string	`json:"genres"`
	Poster		string			`json:"poster"`
}

// MovieGenre は映画ジャンルのタイプです
type MovieGenre struct{
	ID 			int			`json:"-"`
	MovieID		string		`json:"-"`
	GenreID		string		`json:"-"`
	Genre		Genre		`json:"genre"`
	CreatedAd	time.Time	`json:"-"`
	UpdatedAt	time.Time	`json:"-"`
}

// Genre はジャンルのタイプです
type Genre struct{
	ID 			int			`json:"id"`
	GenreName	string		`json:"genre_name"`
	CreatedAd	time.Time	`json:"-"`
	UpdatedAt	time.Time	`json:"-"`
}

// User はユーザーのタイプです
type User struct{
	ID			int
	Email		string
	Password	string
}