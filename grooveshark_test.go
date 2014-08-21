package grooveshark

import (
	"fmt"
	"testing"
)

func TestNewClient(t *testing.T) {
	client := NewClient()
	client.Connect()
	tracks := client.Search("Nothing's Gonna Stop Us Now", "Songs")
	for _, track := range tracks {
		fmt.Println(track.SongName, track.ArtistName)
	}
}
