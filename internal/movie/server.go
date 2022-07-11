package movie

import (
	"database/sql"
	"fmt"
	"myapp/pkg/db/mysql"
)

func fetchMovie() (titles Titles) {
	row, err := mysql.SelectMovieStmt.Query()

	if err != nil {
		return
	}
	defer row.Close()

	var title Title
	var description, ageCertification, tmdbPopularity sql.NullString
	var imdbScore, imdbVotes, tmdbScore sql.NullFloat64
	for row.Next() {
		err = row.Scan(&title.Title, &description,
			&title.ReleaseYear, &ageCertification,
			&title.Runtime, &imdbScore, &imdbVotes,
			&tmdbPopularity, &tmdbScore)

		if description.Valid {
			title.Description = description.String
		}

		if ageCertification.Valid {
			title.AgeCertification = ageCertification.String
		}
		if tmdbPopularity.Valid {
			title.TmdbPopularity = tmdbPopularity.String
		}
		if imdbScore.Valid {
			title.ImdbScore = imdbScore.Float64
		}
		if imdbVotes.Valid {
			title.ImdbVotes = imdbVotes.Float64
		}
		if tmdbScore.Valid {
			title.TmdbScore = tmdbScore.Float64
		}

		if err != nil {
			fmt.Println(err)
		} else {
			titles = append(titles, title)
		}
	}
	return
}
