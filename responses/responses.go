package responses

type Header struct {
	Session         string `json:"session"`
	ServiceVersion  string `json:"serviceVersion"`
	PrefetchEnabled bool   `json:"prefetchedEnabled"`
}

type GetResultsFromSearch struct {
	Header Header `json:"header"`
	Result struct {
		Result           TrackList `json:"result"`
		Version          string    `json:"version"`
		AssignedVersion  string    `json:"assignedVersion"`
		AskForSuggestion bool      `json:"AskForSuggestion"`
	} `json:"result"`
}

type TrackList []Track

type Track struct {
	SongId                 string  `json:"SongID"`
	AlbumId                string  `json:"AlbumID"`
	ArtistId               string  `json:"ArtistID"`
	SongName               string  `json:"SongName"`
	AlbumName              string  `json:"AlbumName"`
	ArtistName             string  `json:"ArtistName"`
	Year                   string  `json:"Year"`
	TrackNum               string  `json:"TrackNum"`
	CoverArtFilename       string  `json:"CoverArtFilename"`
	ArtistCoverArtFilename string  `json:"ArtistCoverArtFilename"`
	TSAdded                string  `json:"TSAdded"`
	AvgDuration            string  `json:"AvgDuration"`
	EstimateDuration       string  `json:"EstimateDuration"`
	Flags                  int     `json:"Flags"`
	IsLowBitrateAvailable  string  `json:"IsLowBitrateAvailable"`
	IsVerified             string  `json:"IsVerified"`
	Popularity             int     `json:"Popularity"`
	Score                  float64 `json:"Score"`
	RawScore               int     `json:"RawScore"`
	PopularityIndex        int     `json:"PopularityIndex"`
}
