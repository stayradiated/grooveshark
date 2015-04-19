package responses

import (
	"strconv"
	"time"
)

type Header struct {
	Session         string `json:"session"`
	ServiceVersion  string `json:"serviceVersion"`
	PrefetchEnabled bool   `json:"prefetchedEnabled"`
}

type GetResultsFromSearch struct {
	Header Header `json:"header"`
	Result struct {
		Result           SearchTrackSlice `json:"result"`
		Version          string           `json:"version"`
		AssignedVersion  string           `json:"assignedVersion"`
		AskForSuggestion bool             `json:"AskForSuggestion"`
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

type GetQueueSongListFromSongIDs struct {
	Header Header             `json:"header"`
	Result PlaylistTrackSlice `json:"result"`
}

type SearchTrackSlice []SearchTrack

func (s *SearchTrackSlice) Tracks() (tracks []Track) {
	// convert SearchTrack into Track
	tracks = make([]Track, len(*s))
	for i := range *s {
		tracks[i] = (*s)[i].Track()
	}
	return tracks
}

type SearchTrack struct {
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

func (s *SearchTrack) Track() (track Track) {
	track = Track{
		Flags:            s.Flags,
		SongName:         s.SongName,
		AlbumName:        s.AlbumName,
		ArtistName:       s.ArtistName,
		Popularity:       s.Popularity,
		CoverArtFilename: s.CoverArtFilename,
	}

	track.Year, _ = strconv.Atoi(s.Year)
	track.SongId, _ = strconv.Atoi(s.SongId)
	track.AlbumId, _ = strconv.Atoi(s.AlbumId)
	track.TrackNum, _ = strconv.Atoi(s.TrackNum)
	track.ArtistId, _ = strconv.Atoi(s.ArtistId)

	estimateDuration, _ := strconv.ParseFloat(s.EstimateDuration, 64)
	if estimateDuration != 4096 {
		track.EstimateDuration = time.Duration(int64(estimateDuration)) * time.Second
	}

	track.IsVerified, _ = strconv.ParseBool(s.IsVerified)
	track.IsLowBitrateAvailable, _ = strconv.ParseBool(s.IsLowBitrateAvailable)

	return track
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
	About          string             `json:"About"`
	AlbumFiles     []string           `json:"AlbumFiles"`
	FName          string             `json:"FName"`
	LName          string             `json:"LName"`
	LastModifiedBy int                `json:"LastModifiedBy"`
	Name           string             `json:"Name"`
	Picture        string             `json:"Picture"`
	PlaylistId     int                `json:"PlaylistID"`
	SongCount      int                `json:"SongCount"`
	Songs          PlaylistTrackSlice `json:"Songs"`
	TSAdded        string             `json:"TSAdded"`
	TSModified     int                `json:"TSModified"`
	UUID           string             `json:"UUID"`
	UserId         int                `json:"UserID"`
	UserName       string             `json:"UserName"`
	UserPicture    string             `json:"UserPicture"`
	Username       string             `json:"Username"`
	TooBig         bool               `json:"tooBig"`
}

type PlaylistTrackSlice []PlaylistTrack

func (p *PlaylistTrackSlice) Tracks() (tracks []Track) {
	// convert SearchTrack into Track
	tracks = make([]Track, len(*p))
	for i := range *p {
		tracks[i] = (*p)[i].Track()
	}
	return tracks
}

type PlaylistTrack struct {
	SongId                string `json:"SongID"`
	Name                  string `json:"Name"`
	SongNameId            string `json:"SongNameID"`
	AlbumId               string `json:"AlbumID"`
	AlbumName             string `json:"AlbumName"`
	ArtistId              string `json:"ArtistID"`
	ArtistName            string `json:"ArtistName"`
	AvgRating             string `json:"AvgRating"`
	IsVerified            string `json:"IsVerified"`
	CoverArtFilename      string `json:"CoverArtFilename"`
	Year                  string `json:"Year"`
	EstimateDuration      string `json:"EstimateDuration"`
	Popularity            string `json:"Popularity"`
	TrackNum              string `json:"TrackNum"`
	IsLowBitrateAvailable string `json:"IsLowBitrateAvailable"`
	Flags                 string `json:"Flags"`
}

func (p *PlaylistTrack) Track() (track Track) {
	track = Track{
		SongName:         p.Name,
		AlbumName:        p.AlbumName,
		ArtistName:       p.ArtistName,
		CoverArtFilename: p.CoverArtFilename,
	}

	track.Year, _ = strconv.Atoi(p.Year)
	track.Flags, _ = strconv.Atoi(p.Flags)
	track.SongId, _ = strconv.Atoi(p.SongId)
	track.AlbumId, _ = strconv.Atoi(p.AlbumId)
	track.TrackNum, _ = strconv.Atoi(p.TrackNum)
	track.ArtistId, _ = strconv.Atoi(p.ArtistId)
	track.Popularity, _ = strconv.Atoi(p.Popularity)

	estimateDuration, _ := strconv.ParseFloat(p.EstimateDuration, 64)
	if estimateDuration != 4096 {
		track.EstimateDuration = time.Duration(int64(estimateDuration)) * time.Second
	}

	track.IsVerified, _ = strconv.ParseBool(p.IsVerified)
	track.IsLowBitrateAvailable, _ = strconv.ParseBool(p.IsLowBitrateAvailable)

	return track
}

type Track struct {
	SongId                int
	SongName              string
	ArtistId              int
	ArtistName            string
	AlbumId               int
	AlbumName             string
	Year                  int
	CoverArtFilename      string
	EstimateDuration      time.Duration
	Flags                 int
	TrackNum              int
	Popularity            int
	IsLowBitrateAvailable bool
	IsVerified            bool
}
