package models

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

type DBModel struct{
	DB *sql.DB
}

// Get は、もしidが一致する映画が１つあれば、その映画とエラーを返します
func (m *DBModel) Get(id int) (*Movie, error){
	ctx, cancel := context.WithTimeout(context.Background(), 3 * time.Second)
	defer cancel()

	query := `
		select
			id, title, description, year, release_date, runtime, rating, mpaa_rating, created_at, updated_at
		from
			movies where id = $1
	`
	row := m.DB.QueryRowContext(ctx, query, id)

	var movie Movie
	err := row.Scan(
		&movie.ID,
		&movie.Title,
		&movie.Description,
		&movie.Year,
		&movie.ReleaseDate,
		&movie.Runtime,
		&movie.Rating,
		&movie.MPAARating,
		&movie.CreatedAd,
		&movie.UpdatedAt,
	)
	if err != nil{
		return nil, err
	}

	// ジャンルを取得する
	query = `
		select
			mg.id, mg.movie_id, mg.genre_id, g.genre_name
		from
			movies_genres mg
			left join genres g on (g.id = mg.genre_id)
		where
			mg.movie_id = $1 
	`
	rows, _ := m.DB.QueryContext(ctx, query, id)
	defer rows.Close()

	genres := make(map[int]string)
	for rows.Next(){
		var mg MovieGenre
		err := rows.Scan(
			&mg.ID,
			&mg.MovieID,
			&mg.GenreID,
			&mg.Genre.GenreName,
		)
		if err != nil{
			return nil, err
		}
		genres[mg.ID] = mg.Genre.GenreName
	}

	movie.MovieGenre = genres

	return &movie, nil
}

// All は、すべての映画とエラー（ある場合）を返します
func (m *DBModel) All(genre ...int) ([]*Movie, error){
	ctx, cancel := context.WithTimeout(context.Background(), 3 * time.Second)
	defer cancel();

	// ジャンル指定の場合
	where := ""
	if len(genre) > 0{
		where = fmt.Sprintf("where id in (select movie_id from movies_genres where genre_id = %d)", genre[0])
	}

	query := fmt.Sprintf(`
		select 
			id, title, description, year, release_date, runtime, rating, mpaa_rating, created_at, updated_at 
		from 
			movies %s order by title`, where)

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var movies []*Movie

	for rows.Next(){
		var movie Movie
		err := rows.Scan(
			&movie.ID,
			&movie.Title,
			&movie.Description,
			&movie.Year,
			&movie.ReleaseDate,
			&movie.Runtime,
			&movie.Rating,
			&movie.MPAARating,
			&movie.CreatedAd,
			&movie.UpdatedAt,	
		)
		if err != nil{
			return nil, err
		}

		// ジャンルを取得する
		genreQuery := `
			select
				mg.id, mg.movie_id, mg.genre_id, g.genre_name
			from
				movies_genres mg
				left join genres g on (g.id = mg.genre_id)
			where
				mg.movie_id = $1 
		`
		genreRows, _ := m.DB.QueryContext(ctx, genreQuery, movie.ID)
	
		genres := make(map[int]string)
		for genreRows.Next(){
			var mg MovieGenre
			err := genreRows.Scan(
				&mg.ID,
				&mg.MovieID,
				&mg.GenreID,
				&mg.Genre.GenreName,
			)
			if err != nil{
				return nil, err
			}
			genres[mg.ID] = mg.Genre.GenreName
		}

		genreRows.Close()

		movie.MovieGenre = genres
		movies = append(movies, &movie)
	
	}
	return movies, nil
}

func (m *DBModel) GenresAll() ([]*Genre, error){
	ctx, cancel := context.WithTimeout(context.Background(), 3 * time.Second)
	defer cancel()

	query := `select 
				id, genre_name, created_at, updated_at 
			from
				genres order by genre_name
	`

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil{
		return nil, err
	}
	defer rows.Close()

	var genres []*Genre

	for rows.Next(){
		var g Genre
		err := rows.Scan(
			&g.ID,
			&g.GenreName,
			&g.CreatedAd,
			&g.UpdatedAt,
		)
		if err != nil{
			return nil, err
		}
		genres = append(genres, &g)
	}

	return genres, nil
}

func (m *DBModel) InsertMovie(movie Movie) error{
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	stmt := `
		insert into movies (title, description, year, release_date, runtime, rating, mpaa_rating, created_at, updated_at)
		values ($1, $2, $3, $4, $5, $6, $7, $8, $9)
	`

	_, err := m.DB.ExecContext(ctx, stmt,
		movie.Title,
		movie.Description,
		movie.Year,
		movie.ReleaseDate, 
		movie.Runtime,
		movie.Rating,
		movie.MPAARating,
		movie.CreatedAd,
		movie.UpdatedAt,
	)
	if err != nil {
		return err
	}
	return nil
}

func (m *DBModel) UpdateMovie(movie Movie) error{
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	stmt := `
		update movies set 
			title = $1, description = $2, year = $3, release_date = $4, runtime = $5, rating = $6, mpaa_rating = $7, updated_at = $8
		where id = $9
	`

	_, err := m.DB.ExecContext(ctx, stmt,
		movie.Title,
		movie.Description,
		movie.Year,
		movie.ReleaseDate, 
		movie.Runtime,
		movie.Rating,
		movie.MPAARating,
		movie.UpdatedAt,
		movie.ID,
	)
	if err != nil {
		return err
	}
	return nil
}