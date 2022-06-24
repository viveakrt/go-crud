package migrate

type Title struct {
	ID                  string `json:"id"`
	Title               string `json:"title"`
	Type                string `json:"type"`
	Description         string `json:"description"`
	ReleaseYear         string `json:"release_year"`
	AgeCertification    string `json:"age_certification"`
	Runtime             string `json:"runtime"`
	Genres              string `json:"genres"`
	ProductionCountries string `json:"production_countries"`
	Seasons             string `json:"seasons"`
	ImdbID              string `json:"imdb_id"`
	ImdbScore           string `json:"imdb_score"`
	ImdbVotes           string `json:"imdb_votes"`
	TmdbPopularity      string `json:"tmdb_popularity"`
	TmdbScore           string `json:"tmdb_score"`
}

type Titles []Title

type Credit struct {
	PersonID  string `json:"person_id"`
	ID        string `json:"id"`
	Name      string `json:"name"`
	Character string `json:"character"`
	Role      string `json:"role"`
}

type Credits []Credit
