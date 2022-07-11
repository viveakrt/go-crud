package mysql

import (
	"database/sql"
	"fmt"
	"os"
)

var SelectMovieStmt *sql.Stmt

func prepareStatements() {
	SelectMovies()
}

func SelectMovies() {
	stmt, err := DB.Prepare(`
	SELECT 
		title,
		description, 
		release_year,
		age_certification,
		runtime,
		imdb_score,
		imdb_votes,
		tmdb_popularity,
		tmdb_score
		FROM 
		content 
		WHERE 
		type = "MOVIE";
	`)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	SelectMovieStmt = stmt

}
