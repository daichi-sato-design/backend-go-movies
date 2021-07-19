package models

import (
	"database/sql"
	"time"
)

// Models is the wrapper for database
type Models struct{
	DB DBModel
}

// NewModels returns models with db pool
func NewModels(db *sql.DB) Models{
	return Models{
		DB: DBModel{DB: db},
	}
}

// Movie is the type for movie
type Movie struct{
	ID			int				`json:"id"`
	Title		string			`json:"title"`
	Description	string			`json:"description"`
	Year		int				`json:"year"`
	ReleaseDate	time.Time		`json:"release_date"`
	Runtime		int				`json:"runtime"`
	Rating		int				`json:"rating"`
	MPAARating	string			`json:"mpaa_rating"`
	CreatedAd	time.Time		`json:"-"`
	UpdatedAt	time.Time		`json:"-"`
	MovieGenre	[]MovieGenre	`json:"genres"`
}

// Genre is the type for genre
type Genre struct{
	ID 			int			`json:"-"`
	GenreName	string		`json:"genre_name"`
	CreatedAd	time.Time	`json:"-"`
	UpdatedAt	time.Time	`json:"-"`
}

// MovieGenre is the type for movie genre
type MovieGenre struct{
	ID 			int			`json:"-"`
	MovieID		string		`json:"-"`
	GenreID		string		`json:"-"`
	Genre		Genre		`json:"genre"`
	CreatedAd	time.Time	`json:"-"`
	UpdatedAt	time.Time	`json:"-"`
}