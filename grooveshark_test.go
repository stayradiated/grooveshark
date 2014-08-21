package grooveshark

import (
	"encoding/json"
	"fmt"
	"testing"
)

// func TestNewClient(t *testing.T) {
// 	client := NewClient()
// 	client.Connect()
// 	tracks := client.Search("99 Red Balloons Nena", "Songs")
// 	track := tracks[0]
//
// 	songId, _ := strconv.Atoi(track.SongId)
// 	sk := client.GetStreamKey(songId)
//
// 	fileName := track.SongName + ".mp3"
// 	fmt.Println("Downloading to", fileName)
//
// 	// TODO: check file existence first with io.IsExist
// 	output, err := os.Create(fileName)
// 	if err != nil {
// 		fmt.Println("Error while creating", fileName, "-", err)
// 		return
// 	}
// 	defer output.Close()
//
// 	response, err := sk.Download()
// 	if err != nil {
// 		fmt.Println("Error while downloading -", err)
// 		return
// 	}
// 	defer response.Body.Close()
//
// 	n, err := io.Copy(output, response.Body)
// 	if err != nil {
// 		fmt.Println("Error while downloading -", err)
// 		return
// 	}
//
// 	fmt.Println(n, "bytes downloaded.")
// }

// func TestPlaylist(t *testing.T) {
// 	client := NewClient()
// 	client.Connect()
// 	playlist := client.Playlist(89939262)
//
// 	trackIds := make([]int, len(playlist.Songs))
// 	for i, playlistTrack := range playlist.Songs {
// 		trackIds[i] = playlistTrack.Track().SongId
// 	}
//
// 	tracks := client.LookupTrackIds(trackIds)
//
// 	for _, track := range tracks {
// 		fmt.Println(track)
// 	}
// }

func TestSearch(t *testing.T) {
	client := NewClient()
	client.Connect()
	tracks := client.Search("topgear")
	output, _ := json.Marshal(tracks)
	fmt.Println(string(output))
}
