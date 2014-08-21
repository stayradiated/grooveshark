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

type GetPlaylistByID struct {
	Header Header   `json:"header"`
	Result Playlist `json:"result"`
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

type Playlist struct {
	// About: ""
	// AlbumFiles: [1085811.jpg, 2688559.jpg, 5978723.jpg, 977767.jpg]
	// FName: "Jocelyn Ziegler"
	// LName: ""
	// LastModifiedBy: 4893556
	// Name: "Summer 2014"
	// Picture: "977767-1085811-2688559-5978723.jpg"
	// PlaylistID: 100071579
	// SongCount: 16
	// Songs: [{SongID:41177002, Name:Marilyn Monroe, SongNameID:329517, AlbumID:9707883, AlbumName:Girl,…},…]
	// TSAdded: "2014-08-11 21:22:28"
	// TSModified: 1407874857
	// UUID: "53e96c54056d678f5f000000"
	// UserID: 4893556
	// UserName: "Jocelyn Ziegler"
	// UserPicture: "4893556-20140525202840.jpg"
	// Username: "Jocelyn Ziegler"
	// tooBig: false
}

type PlaylistTrack struct {
	// AlbumID: "9707883"
	// AlbumName: "Girl"
	// ArtistID: "22594"
	// ArtistName: "Pharrell Williams"
	// AvgRating: "0"
	// CoverArtFilename: null
	// EstimateDuration: "350"
	// Flags: "0"
	// IsLowBitrateAvailable: "1"
	// IsVerified: "1"
	// Name: "Marilyn Monroe"
	// Popularity: "1423300847"
	// SongID: "41177002"
	// SongNameID: "329517"
	// TrackNum: "1"
	// Year: "2014"
}
