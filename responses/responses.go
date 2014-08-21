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

type returnedResult struct {
	Return bool `json:"Return"`
}

type MarkSongDownloadedEx struct {
	Header Header         `json:"header"`
	Result returnedResult `json:"result"`
}

type GetStreamKeyFromSongIDEx struct {
	Header Header    `json:"header"`
	Result StreamKey `json:"result"`
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

type StreamKey struct {
	Expires        int    `json:"Expires"`
	FileId         string `json:"FileID"`
	FileToken      string `json:"FileToken"`
	SongId         int    `json:"SongID"`
	Ip             string `json:"ip"`
	IsMobile       bool   `json:"isMobile"`
	StreamKey      string `json:"streamKey"`
	StreamServerId int    `json:"streamServerID"`
	Timestamp      int    `json:"ts"`
	USecs          string `json:"uSecs"`
}
