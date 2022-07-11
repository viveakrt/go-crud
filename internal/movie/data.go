package movie

type Title struct {
	// ID                  string  `json:"id" sql:"id"`
	Title string `json:"title" sql:"title"`
	// Type                string  `json:"type" sql:"type"`
	Description      string `json:"description" sql:"description"`
	ReleaseYear      int64  `json:"release_year" sql:"release_year"`
	AgeCertification string `json:"age_certification" sql:"age_certification"`
	Runtime          int64  `json:"runtime" sql:"runtime"`
	// Genres              string  `json:"genres" sql:"genres"`
	// ProductionCountries string  `json:"production_countries" sql:"production_countries"`
	// Seasons             string  `json:"seasons" sql:"seasons"`
	// ImdbID              float64 `json:"imdb_id" sql:"imdb_id"`
	ImdbScore      float64 `json:"imdb_score" sql:"imdb_score"`
	ImdbVotes      float64 `json:"imdb_votes" sql:"imdb_votes"`
	TmdbPopularity string  `json:"tmdb_popularity" sql:"tmdb_popularity"`
	TmdbScore      float64 `json:"tmdb_score" sql:"tmdb_score"`
}
type Colour struct {
	ID    string `json:"id"`
	Color string `json:"color"`
}
type Titles []Title

type Credit struct {
	PersonID  string `json:"person_id" sql:"person_id"`
	ID        string `json:"id" sql:"id"`
	Name      string `json:"name" sql:"name"`
	Character string `json:"character" sql:"character"`
	Role      string `json:"role" sql:"role"`
}

type Credits []Credit
