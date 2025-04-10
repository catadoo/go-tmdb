package tmdb

import (
	"fmt"
)

// Movie struct
type Movie struct {
	Adult                bool                   `json:"adult"`
	BackdropPath         string                 `json:"backdropPath"`
	BelongsToCollection  CollectionShort        `json:"belongsToCollection"`
	Budget               uint32                 `json:"budget"`
	Genres               []struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"genres"`
	Homepage            string                 `json:"homepage"`
	ID                  int                    `json:"id"`
	ImdbID              string                 `json:"imdbId"`
	OriginalLanguage    string                 `json:"originalLanguage"`
	OriginalTitle       string                 `json:"originalTitle"`
	Overview            string                 `json:"overview"`
	Popularity          float32                `json:"popularity"`
	PosterPath          string                 `json:"posterPath"`
	ProductionCompanies []struct {
		ID        int    `json:"id"`
		Name      string `json:"name"`
		LogoPath  string `json:"logoPath"`
		Iso3166_1 string `json:"iso31661"`
	} `json:"productionCompanies"`
	ProductionCountries []struct {
		Iso3166_1 string `json:"iso31661"`
		Name      string `json:"name"`
	} `json:"productionCountries"`
	ReleaseDate     string `json:"releaseDate"`
	Revenue         uint32 `json:"revenue"`
	Runtime         uint32 `json:"runtime"`
	SpokenLanguages []struct {
		Iso639_1 string `json:"iso6391"`
		Name     string `json:"name"`
	} `json:"spokenLanguages"`
	Status            string                `json:"status"`
	Tagline           string                `json:"tagline"`
	Title             string                `json:"title"`
	Video             bool                  `json:"video"`
	VoteAverage       float32               `json:"voteAverage"`
	VoteCount         uint32                `json:"voteCount"`
	AlternativeTitles *MovieAlternativeTitles `json:"alternativeTitles,omitempty"`
	Credits           *MovieCredits           `json:"credits,omitempty"`
	Images            *MovieImages            `json:"images,omitempty"`
	Keywords          *MovieKeywords          `json:"keywords,omitempty"`
	Releases          *MovieReleases          `json:"releases,omitempty"`
	Videos            *MovieVideos            `json:"videos,omitempty"`
	Translations      *MovieTranslations      `json:"translations,omitempty"`
	Similar           *MoviePagedResults      `json:"similar,omitempty"`
	Reviews           *MovieReviews           `json:"reviews,omitempty"`
	Lists             *MovieLists             `json:"lists,omitempty"`
	Changes           *MovieChanges           `json:"changes,omitempty"`
	Rating            *MovieRating            `json:"rating,omitempty"`
	ExternalIDs       *MovieExternalIds       `json:"externalIds,omitempty"`
}

// MovieShort struct
type MovieShort struct {
	Adult         bool    `json:"adult"`
	BackdropPath  string  `json:"backdropPath"`
	ID            int     `json:"id"`
	OriginalTitle string  `json:"originalTitle"`
	GenreIDs      []int32 `json:"genreIds"`
	Popularity    float32 `json:"popularity"`
	PosterPath    string  `json:"posterPath"`
	ReleaseDate   string  `json:"releaseDate"`
	Title         string  `json:"title"`
	Overview      string  `json:"overview"`
	Video         bool    `json:"video"`
	VoteAverage   float32 `json:"voteAverage"`
	VoteCount     uint32  `json:"voteCount"`
}

// MovieDatedResults struct
type MovieDatedResults struct {
	Dates struct {
		Minimum string `json:"minimum"`
		Maximum string `json:"maximum"`
	} `json:"dates"`
	Page         int          `json:"page"`
	Results      []MovieShort `json:"results"`
	TotalPages   int          `json:"totalPages"`
	TotalResults int          `json:"totalResults"`
}

// MoviePagedResults struct
type MoviePagedResults struct {
	ID                int                    `json:"id"`
	Page              int                    `json:"page"`
	Results           []MovieShort           `json:"results"`
	TotalPages        int                    `json:"totalPages"`
	TotalResults      int                    `json:"totalResults"`
	AlternativeTitles *MovieAlternativeTitles `json:"alternativeTitles,omitempty"`
	Credits           *MovieCredits           `json:"credits,omitempty"`
	Images            *MovieImages            `json:"images,omitempty"`
	Keywords          *MovieKeywords          `json:"keywords,omitempty"`
	Releases          *MovieReleases          `json:"releases,omitempty"`
	Videos            *MovieVideos            `json:"videos,omitempty"`
	Translations      *MovieTranslations      `json:"translations,omitempty"`
	Similar           *MoviePagedResults      `json:"similar,omitempty"`
	Reviews           *MovieReviews           `json:"reviews,omitempty"`
	Lists             *MovieLists             `json:"lists,omitempty"`
	Changes           *MovieChanges           `json:"changes,omitempty"`
	Rating            *MovieRating            `json:"rating,omitempty"`
}

// MovieAccountState struct
type MovieAccountState struct {
	ID        int `json:"id"`
	Favorite  bool `json:"favorite"`
	Watchlist bool `json:"watchlist"`
	Rated struct {
		Value float32 `json:"value"`
	} `json:"rated"`
}

// MovieAlternativeTitles struct
type MovieAlternativeTitles struct {
	ID     int `json:"id"`
	Titles []struct {
		Iso3166_1 string `json:"iso31661"`
		Title     string `json:"title"`
	} `json:"titles"`
	AccountStates *MovieAccountState `json:"accountStates,omitempty"`
	Credits       *MovieCredits      `json:"credits,omitempty"`
	Images        *MovieImages       `json:"images,omitempty"`
	Keywords      *MovieKeywords     `json:"keywords,omitempty"`
	Releases      *MovieReleases     `json:"releases,omitempty"`
	Videos        *MovieVideos       `json:"videos,omitempty"`
	Translations  *MovieTranslations `json:"translations,omitempty"`
	Similar       *MoviePagedResults `json:"similar,omitempty"`
	Reviews       *MovieReviews      `json:"reviews,omitempty"`
	Lists         *MovieLists        `json:"lists,omitempty"`
	Changes       *MovieChanges      `json:"changes,omitempty"`
	Rating        *MovieRating       `json:"rating,omitempty"`
}

// MovieChanges struct
type MovieChanges struct {
	Changes []struct {
		Key   string `json:"key"`
		Items []struct {
			ID     string `json:"id"`
			Action string `json:"action"`
			Time   string `json:"time"`
		} `json:"items"`
	} `json:"changes"`
}

// MovieCredits struct
type MovieCredits struct {
	ID   int `json:"id"`
	Cast []struct {
		CastID      int    `json:"castId"`
		Character   string `json:"character"`
		CreditID    string `json:"creditId"`
		ID          int    `json:"id"`
		Name        string `json:"name"`
		Gender      int    `json:"gender"`
		Order       int    `json:"order"`
		ProfilePath string `json:"profilePath"`
	} `json:"cast"`
	Crew []struct {
		CreditID    string `json:"creditId"`
		Department  string `json:"department"`
		Gender      int    `json:"gender"`
		ID          int    `json:"id"`
		Job         string `json:"job"`
		Name        string `json:"name"`
		ProfilePath string `json:"profilePath"`
	} `json:"crew"`
	AlternativeTitles *MovieAlternativeTitles `json:"alternativeTitles,omitempty"`
	Images            *MovieImages            `json:"images,omitempty"`
	Keywords          *MovieKeywords          `json:"keywords,omitempty"`
	Releases          *MovieReleases          `json:"releases,omitempty"`
	Videos            *MovieVideos            `json:"videos,omitempty"`
	Translations      *MovieTranslations      `json:"translations,omitempty"`
	Similar           *MoviePagedResults      `json:"similar,omitempty"`
	Reviews           *MovieReviews           `json:"reviews,omitempty"`
	Lists             *MovieLists             `json:"lists,omitempty"`
	Changes           *MovieChanges           `json:"changes,omitempty"`
	Rating            *MovieRating            `json:"rating,omitempty"`
}

// MovieExternalIds struct
type MovieExternalIds struct {
	ID          int    `json:"id"`
	ImdbID      string `json:"imdbId"`
	FacebookID  string `json:"facebookId"`
	InstagramID string `json:"instagramId"`
	TwitterID   string `json:"twitterId"`
}

// MovieImage struct
type MovieImage struct {
	FilePath    string  `json:"filePath"`
	Width       int     `json:"width"`
	Height      int     `json:"height"`
	Iso639_1    string  `json:"iso6391"`
	AspectRatio float32 `json:"aspectRatio"`
	VoteAverage float32 `json:"voteAverage"`
	VoteCount   uint32  `json:"voteCount"`
}

// MovieImages struct
type MovieImages struct {
	ID                int             `json:"id"`
	Backdrops         []MovieImage    `json:"backdrops"`
	Posters           []MovieImage    `json:"posters"`
	Logos             []MovieImage    `json:"logos"`
	AlternativeTitles *MovieAlternativeTitles `json:"alternativeTitles,omitempty"`
	Credits           *MovieCredits           `json:"credits,omitempty"`
	Keywords          *MovieKeywords          `json:"keywords,omitempty"`
	Releases          *MovieReleases          `json:"releases,omitempty"`
	Videos            *MovieVideos            `json:"videos,omitempty"`
	Translations      *MovieTranslations      `json:"translations,omitempty"`
	Similar           *MoviePagedResults      `json:"similar,omitempty"`
	Reviews           *MovieReviews           `json:"reviews,omitempty"`
	Lists             *MovieLists             `json:"lists,omitempty"`
	Changes           *MovieChanges           `json:"changes,omitempty"`
	Rating            *MovieRating            `json:"rating,omitempty"`
}

// MovieKeywords struct
type MovieKeywords struct {
	ID       int `json:"id"`
	Keywords []struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"keywords"`
	AlternativeTitles *MovieAlternativeTitles `json:"alternativeTitles,omitempty"`
	Credits           *MovieCredits           `json:"credits,omitempty"`
	Images            *MovieImages            `json:"images,omitempty"`
	Releases          *MovieReleases          `json:"releases,omitempty"`
	Videos            *MovieVideos            `json:"videos,omitempty"`
	Translations      *MovieTranslations      `json:"translations,omitempty"`
	Similar           *MoviePagedResults      `json:"similar,omitempty"`
	Reviews           *MovieReviews           `json:"reviews,omitempty"`
	Lists             *MovieLists             `json:"lists,omitempty"`
	Changes           *MovieChanges           `json:"changes,omitempty"`
	Rating            *MovieRating            `json:"rating,omitempty"`
}

// MovieLists struct
type MovieLists struct {
	ID      int `json:"id"`
	Page    int `json:"page"`
	Results []struct {
		Description   string `json:"description"`
		FavoriteCount int    `json:"favoriteCount"`
		ID            string `json:"id"`
		ItemCount     int    `json:"itemCount"`
		Iso639_1      string `json:"iso6391"`
		Name          string `json:"name"`
		PosterPath    string `json:"posterPath"`
	} `json:"results"`
	TotalPages        int `json:"totalPages"`
	TotalResults      int `json:"totalResults"`
	AlternativeTitles *MovieAlternativeTitles `json:"alternativeTitles,omitempty"`
	Credits           *MovieCredits           `json:"credits,omitempty"`
	Images            *MovieImages            `json:"images,omitempty"`
	Keywords          *MovieKeywords          `json:"keywords,omitempty"`
	Releases          *MovieReleases          `json:"releases,omitempty"`
	Videos            *MovieVideos            `json:"videos,omitempty"`
	Translations      *MovieTranslations      `json:"translations,omitempty"`
	Similar           *MoviePagedResults      `json:"similar,omitempty"`
	Reviews           *MovieReviews           `json:"reviews,omitempty"`
	Changes           *MovieChanges           `json:"changes,omitempty"`
	Rating            *MovieRating            `json:"rating,omitempty"`
}

// MovieRating struct
type MovieRating struct {
	StatusCode    int    `json:"statusCode"`
	StatusMessage string `json:"statusMessage"`
}

// MovieReleases struct
type MovieReleases struct {
	ID        int `json:"id"`
	Countries []struct {
		Iso3166_1     string `json:"iso31661"`
		Certification string `json:"certification"`
		ReleaseDate   string `json:"releaseDate"`
	} `json:"countries"`
	AlternativeTitles *MovieAlternativeTitles `json:"alternativeTitles,omitempty"`
	Credits           *MovieCredits           `json:"credits,omitempty"`
	Images            *MovieImages            `json:"images,omitempty"`
	Keywords          *MovieKeywords          `json:"keywords,omitempty"`
	Videos            *MovieVideos            `json:"videos,omitempty"`
	Translations      *MovieTranslations      `json:"translations,omitempty"`
	Similar           *MoviePagedResults      `json:"similar,omitempty"`
	Reviews           *MovieReviews           `json:"reviews,omitempty"`
	Lists             *MovieLists             `json:"lists,omitempty"`
	Changes           *MovieChanges           `json:"changes,omitempty"`
	Rating            *MovieRating            `json:"rating,omitempty"`
}

// MovieReviews struct
type MovieReviews struct {
	ID      int `json:"id"`
	Page    int `json:"page"`
	Results []struct {
		ID      string `json:"id"`
		Author  string `json:"author"`
		Content string `json:"content"`
		URL     string `json:"url"`
	} `json:"results"`
	TotalPages        int `json:"totalPages"`
	TotalResults      int `json:"totalResults"`
	AlternativeTitles *MovieAlternativeTitles `json:"alternativeTitles,omitempty"`
	Credits           *MovieCredits           `json:"credits,omitempty"`
	Images            *MovieImages            `json:"images,omitempty"`
	Keywords          *MovieKeywords          `json:"keywords,omitempty"`
	Releases          *MovieReleases          `json:"releases,omitempty"`
	Videos            *MovieVideos            `json:"videos,omitempty"`
	Translations      *MovieTranslations      `json:"translations,omitempty"`
	Similar           *MoviePagedResults      `json:"similar,omitempty"`
	Lists             *MovieLists             `json:"lists,omitempty"`
	Changes           *MovieChanges           `json:"changes,omitempty"`
	Rating            *MovieRating            `json:"rating,omitempty"`
}

// MovieTranslations struct
type MovieTranslations struct {
	ID           int `json:"id"`
	Translations []struct {
		Iso3166_1   string `json:"iso31661"`
		Iso639_1    string `json:"iso6391"`
		Name        string `json:"name"`
		EnglishName string `json:"englishName"`
		Data        struct {
			Title    string `json:"title,omitempty"`
			Overview string `json:"overview,omitempty"`
			Homepage string `json:"homepage,omitempty"`
		} `json:"data"`
	} `json:"translations"`
	AlternativeTitles *MovieAlternativeTitles `json:"alternativeTitles,omitempty"`
	Credits           *MovieCredits           `json:"credits,omitempty"`
	Images            *MovieImages            `json:"images,omitempty"`
	Keywords          *MovieKeywords          `json:"keywords,omitempty"`
	Releases          *MovieReleases          `json:"releases,omitempty"`
	Videos            *MovieVideos            `json:"videos,omitempty"`
	Similar           *MoviePagedResults      `json:"similar,omitempty"`
	Reviews           *MovieReviews           `json:"reviews,omitempty"`
	Lists             *MovieLists             `json:"lists,omitempty"`
	Changes           *MovieChanges           `json:"changes,omitempty"`
	Rating            *MovieRating            `json:"rating,omitempty"`
}

// MovieRecommendations struct
type MovieRecommendations struct {
	Page    int `json:"page"`
	Results []struct {
		Adult            bool    `json:"adult"`
		BackdropPath     string  `json:"backdropPath"`
		GenreIDs         []int   `json:"genreIds"`
		ID               int     `json:"id"`
		OriginalLanguage string  `json:"originalLanguage"`
		OriginalTitle    string  `json:"originalTitle"`
		Overview         string  `json:"overview"`
		ReleaseDate      string  `json:"releaseDate"`
		PosterPath       string  `json:"posterPath"`
		Popularity       float32 `json:"popularity"`
		Title            string  `json:"title"`
		Video            bool    `json:"video"`
		VoteAverage      float32 `json:"voteAverage"`
		VoteCount        uint32  `json:"voteCount"`
	} `json:"results"`
	TotalPages   int `json:"totalPages"`
	TotalResults int `json:"totalResults"`
}

// MovieVideos struct
type MovieVideos struct {
	ID      int `json:"id"`
	Results []struct {
		ID       string `json:"id"`
		Iso639_1 string `json:"iso6391"`
		Key      string `json:"key"`
		Name     string `json:"name"`
		Site     string `json:"site"`
		Size     int    `json:"size"`
		Type     string `json:"type"`
	} `json:"results"`
	AlternativeTitles *MovieAlternativeTitles `json:"alternativeTitles,omitempty"`
	Credits           *MovieCredits           `json:"credits,omitempty"`
	Images            *MovieImages            `json:"images,omitempty"`
	Keywords          *MovieKeywords          `json:"keywords,omitempty"`
	Releases          *MovieReleases          `json:"releases,omitempty"`
	Translations      *MovieTranslations      `json:"translations,omitempty"`
	Similar           *MoviePagedResults      `json:"similar,omitempty"`
	Reviews           *MovieReviews           `json:"reviews,omitempty"`
	Lists             *MovieLists             `json:"lists,omitempty"`
	Changes           *MovieChanges           `json:"changes,omitempty"`
	Rating            *MovieRating            `json:"rating,omitempty"`
}


// GetMovieInfo for a specific movie id
// https://developers.themoviedb.org/3/movies/get-movie-details
func (tmdb *TMDb) GetMovieInfo(id int, options map[string]string) (*Movie, error) {
	var availableOptions = map[string]struct{}{
		"language":           {},
		"append_to_response": {}}
	var movie Movie
	optionsString := getOptionsString(options, availableOptions)
	uri := fmt.Sprintf("%s/movie/%v?api_key=%s%s", baseURL, id, tmdb.apiKey, optionsString)
	result, err := getTmdb(uri, &movie)
	return result.(*Movie), err
}

// GetMovieAccountStates gets the status of whether or not the movie has been rated or added to their favourite or movie watch list
// https://developers.themoviedb.org/3/movies/get-movie-account-states
func (tmdb *TMDb) GetMovieAccountStates(id int, sessionID string) (*MovieAccountState, error) {
	var state MovieAccountState
	uri := fmt.Sprintf("%s/movie/%v/account_states?api_key=%s&session_id=%s", baseURL, id, tmdb.apiKey, sessionID)
	result, err := getTmdb(uri, &state)
	return result.(*MovieAccountState), err
}

// GetMovieAlternativeTitles for a specific movie id
// https://developers.themoviedb.org/3/movies/get-movie-alternative-titles
func (tmdb *TMDb) GetMovieAlternativeTitles(id int, options map[string]string) (*MovieAlternativeTitles, error) {
	var availableOptions = map[string]struct{}{
		"country":            {},
		"append_to_response": {}}
	var titles MovieAlternativeTitles
	optionsString := getOptionsString(options, availableOptions)
	uri := fmt.Sprintf("%s/movie/%v/alternative_titles?api_key=%s%s", baseURL, id, tmdb.apiKey, optionsString)
	result, err := getTmdb(uri, &titles)
	return result.(*MovieAlternativeTitles), err
}

// GetMovieChanges for a specific movie id
// https://developers.themoviedb.org/3/movies/get-movie-changes
func (tmdb *TMDb) GetMovieChanges(id int, options map[string]string) (*MovieChanges, error) {
	var availableOptions = map[string]struct{}{
		"start_date": {},
		"end_date":   {}}
	var changes MovieChanges
	optionsString := getOptionsString(options, availableOptions)
	uri := fmt.Sprintf("%s/movie/%v/changes?api_key=%s%s", baseURL, id, tmdb.apiKey, optionsString)
	result, err := getTmdb(uri, &changes)
	return result.(*MovieChanges), err
}

// GetMovieCredits for a specific movie id
// https://developers.themoviedb.org/3/movies/get-movie-credits
func (tmdb *TMDb) GetMovieCredits(id int, options map[string]string) (*MovieCredits, error) {
	var availableOptions = map[string]struct{}{
		"append_to_response": {}}
	var credits MovieCredits
	optionsString := getOptionsString(options, availableOptions)
	uri := fmt.Sprintf("%s/movie/%v/credits?api_key=%s%s", baseURL, id, tmdb.apiKey, optionsString)
	result, err := getTmdb(uri, &credits)
	return result.(*MovieCredits), err
}

// GetMovieImages for a specific movie id
// https://developers.themoviedb.org/3/movies/get-movie-images
func (tmdb *TMDb) GetMovieImages(id int, options map[string]string) (*MovieImages, error) {
	var availableOptions = map[string]struct{}{
		"language":               {},
		"append_to_response":     {},
		"include_image_language": {}}
	var images MovieImages
	optionsString := getOptionsString(options, availableOptions)
	uri := fmt.Sprintf("%s/movie/%v/images?api_key=%s%s", baseURL, id, tmdb.apiKey, optionsString)
	result, err := getTmdb(uri, &images)
	return result.(*MovieImages), err
}

// GetMovieKeywords for a specific movie id
// https://developers.themoviedb.org/3/movies/get-movie-keywords
func (tmdb *TMDb) GetMovieKeywords(id int, options map[string]string) (*MovieKeywords, error) {
	var availableOptions = map[string]struct{}{
		"append_to_response": {}}
	var keywords MovieKeywords
	optionsString := getOptionsString(options, availableOptions)
	uri := fmt.Sprintf("%s/movie/%v/keywords?api_key=%s%s", baseURL, id, tmdb.apiKey, optionsString)
	result, err := getTmdb(uri, &keywords)
	return result.(*MovieKeywords), err
}

// GetMovieLatest gets the latest movie
// https://developers.themoviedb.org/3/movies/get-latest-movie
func (tmdb *TMDb) GetMovieLatest() (*Movie, error) {
	var movie Movie
	uri := fmt.Sprintf("%s/movie/latest?api_key=%s", baseURL, tmdb.apiKey)
	result, err := getTmdb(uri, &movie)
	return result.(*Movie), err
}

// GetMovieLists that the movie belongs to
// https://developers.themoviedb.org/3/movies/get-movie-lists
func (tmdb *TMDb) GetMovieLists(id int, options map[string]string) (*MovieLists, error) {
	var availableOptions = map[string]struct{}{
		"page":               {},
		"language":           {},
		"append_to_response": {}}
	var lists MovieLists
	optionsString := getOptionsString(options, availableOptions)
	uri := fmt.Sprintf("%s/movie/%v/lists?api_key=%s%s", baseURL, id, tmdb.apiKey, optionsString)
	result, err := getTmdb(uri, &lists)
	return result.(*MovieLists), err
}

// GetMovieNowPlaying that have been, or are being released this week
// https://developers.themoviedb.org/3/movies/get-now-playing
func (tmdb *TMDb) GetMovieNowPlaying(options map[string]string) (*MovieDatedResults, error) {
	var availableOptions = map[string]struct{}{
		"page":     {},
		"language": {}}
	var nowPlaying MovieDatedResults
	optionsString := getOptionsString(options, availableOptions)
	uri := fmt.Sprintf("%s/movie/now_playing?api_key=%s%s", baseURL, tmdb.apiKey, optionsString)
	result, err := getTmdb(uri, &nowPlaying)
	return result.(*MovieDatedResults), err
}

// GetMoviePopular gets the list of popular movies on The Movie Database
// https://developers.themoviedb.org/3/movies/get-popular-movies
func (tmdb *TMDb) GetMoviePopular(options map[string]string) (*MoviePagedResults, error) {
	var availableOptions = map[string]struct{}{
		"page":     {},
		"language": {}}
	var popular MoviePagedResults
	optionsString := getOptionsString(options, availableOptions)
	uri := fmt.Sprintf("%s/movie/popular?api_key=%s%s", baseURL, tmdb.apiKey, optionsString)
	result, err := getTmdb(uri, &popular)
	return result.(*MoviePagedResults), err
}

// GetMovieReleases for a specific movie id
// https://developers.themoviedb.org/3/movies/get-movie-release-dates
func (tmdb *TMDb) GetMovieReleases(id int, options map[string]string) (*MovieReleases, error) {
	var availableOptions = map[string]struct{}{
		"append_to_response": {}}
	var releases MovieReleases
	optionsString := getOptionsString(options, availableOptions)
	uri := fmt.Sprintf("%s/movie/%v/releases?api_key=%s%s", baseURL, id, tmdb.apiKey, optionsString)
	result, err := getTmdb(uri, &releases)
	return result.(*MovieReleases), err
}

// GetMovieReviews for a specific movie id
// https://developers.themoviedb.org/3/movies/get-movie-reviews
func (tmdb *TMDb) GetMovieReviews(id int, options map[string]string) (*MovieReviews, error) {
	var availableOptions = map[string]struct{}{
		"page":               {},
		"language":           {},
		"append_to_response": {}}
	var reviews MovieReviews
	optionsString := getOptionsString(options, availableOptions)
	uri := fmt.Sprintf("%s/movie/%v/reviews?api_key=%s%s", baseURL, id, tmdb.apiKey, optionsString)
	result, err := getTmdb(uri, &reviews)
	return result.(*MovieReviews), err
}

// GetMovieSimilar for a specific movie id
// https://developers.themoviedb.org/3/movies/get-similar-movies
func (tmdb *TMDb) GetMovieSimilar(id int, options map[string]string) (*MoviePagedResults, error) {
	var availableOptions = map[string]struct{}{
		"page":               {},
		"language":           {},
		"append_to_response": {}}
	var similar MoviePagedResults
	optionsString := getOptionsString(options, availableOptions)
	uri := fmt.Sprintf("%s/movie/%v/similar?api_key=%s%s", baseURL, id, tmdb.apiKey, optionsString)
	result, err := getTmdb(uri, &similar)
	return result.(*MoviePagedResults), err
}

// GetMovieTopRated gets the list of top rated movies
// https://developers.themoviedb.org/3/movies/get-top-rated-movies
func (tmdb *TMDb) GetMovieTopRated(options map[string]string) (*MoviePagedResults, error) {
	var availableOptions = map[string]struct{}{
		"page":     {},
		"language": {}}
	var topRated MoviePagedResults
	optionsString := getOptionsString(options, availableOptions)
	uri := fmt.Sprintf("%s/movie/top_rated?api_key=%s%s", baseURL, tmdb.apiKey, optionsString)
	result, err := getTmdb(uri, &topRated)
	return result.(*MoviePagedResults), err
}

// GetMovieTranslations for a specific movie id
// https://developers.themoviedb.org/3/movies/get-movie-translations
func (tmdb *TMDb) GetMovieTranslations(id int, options map[string]string) (*MovieTranslations, error) {
	var availableOptions = map[string]struct{}{
		"append_to_response": {}}
	var translations MovieTranslations
	optionsString := getOptionsString(options, availableOptions)
	uri := fmt.Sprintf("%s/movie/%v/translations?api_key=%s%s", baseURL, id, tmdb.apiKey, optionsString)
	result, err := getTmdb(uri, &translations)
	return result.(*MovieTranslations), err
}

// GetMovieRecommendations gets a list of recommended movies for a movie by id
// https://developers.themoviedb.org/3/movies/get-movie-recommendations
func (tmdb *TMDb) GetMovieRecommendations(id int, options map[string]string) (*MovieRecommendations, error) {
	var availableOptions = map[string]struct{}{
		"language": {},
		"page":     {}}
	var movieRec MovieRecommendations
	optionsString := getOptionsString(options, availableOptions)
	uri := fmt.Sprintf("%s/movie/%v/recommendations?api_key=%s%s", baseURL, id, tmdb.apiKey, optionsString)
	result, err := getTmdb(uri, &movieRec)
	return result.(*MovieRecommendations), err
}

// GetMovieVideos for a specific movie id
// https://developers.themoviedb.org/3/movies/get-movie-recommendations
func (tmdb *TMDb) GetMovieVideos(id int, options map[string]string) (*MovieVideos, error) {
	var availableOptions = map[string]struct{}{
		"language":           {},
		"append_to_response": {}}
	var videos MovieVideos
	optionsString := getOptionsString(options, availableOptions)
	uri := fmt.Sprintf("%s/movie/%v/videos?api_key=%s%s", baseURL, id, tmdb.apiKey, optionsString)
	result, err := getTmdb(uri, &videos)
	return result.(*MovieVideos), err
}

// GetMovieUpcoming by release date
// https://developers.themoviedb.org/3/movies/get-upcoming
func (tmdb *TMDb) GetMovieUpcoming(options map[string]string) (*MovieDatedResults, error) {
	var availableOptions = map[string]struct{}{
		"page":     {},
		"language": {}}
	var upcoming MovieDatedResults
	optionsString := getOptionsString(options, availableOptions)
	uri := fmt.Sprintf("%s/movie/upcoming?api_key=%s%s", baseURL, tmdb.apiKey, optionsString)
	result, err := getTmdb(uri, &upcoming)
	return result.(*MovieDatedResults), err
}

// GetMovieExternalIds gets the external ids for a movie
// https://developers.themoviedb.org/3/movies/get-movie-external-ids
func (tmdb *TMDb) GetMovieExternalIds(movieID int, options map[string]string) (*MovieExternalIds, error) {
	// currently there are not options, left it so it may be updated in the future without breaking existing code
	var ids MovieExternalIds
	uri := fmt.Sprintf("%s/movie/%v/external_ids?api_key=%s", baseURL, movieID, tmdb.apiKey)
	result, err := getTmdb(uri, &ids)
	return result.(*MovieExternalIds), err
}

// // SetMovieRating lets users rate a movie
// // https://developers.themoviedb.org/3/movies/rate-movie
// func (tmdb *TMDb) SetMovieRating(id int) (*MovieRating, error) {
// 	// TODO
// 	var rating MovieRating
// 	var err error
// 	return &rating, err
// }
