package requests

import "github.com/stayradiated/grooveshark/session"

type GetResultsFromSearch struct {
	Guts       int    `json:"guts"`
	PPOverride bool   `json:"ppOverride"`
	Query      string `json:"query"`
	Type       string `json:"type"`
}

type markSong struct {
	SongId         string `json:"songID"`
	StreamKey      string `json:"streamKey"`
	StreamServerId string `json:"streamServerID"`
}

type MarkSongDownloadedEx markSong
type MarkStreamKeyOver30Seconds markSong

type GetStreamKeyFromSongIDEx struct {
	SongId   int              `json:"songID"`
	Type     int              `json:"type"`
	Prefetch bool             `json:"prefetch"`
	Mobile   bool             `json:"mobile"`
	ReturnTS bool             `json:"returnTS"`
	Country  *session.Country `json:"country"`
}
