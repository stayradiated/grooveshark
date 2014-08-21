package grooveshark

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"testing"
)

func TestNewClient(t *testing.T) {
	client := NewClient()
	client.Connect()
	tracks := client.Search("Time After Time Cyndi Lauper", "Songs")
	track := tracks[0]

	songId, _ := strconv.Atoi(track.SongId)
	sk := client.GetStreamKey(songId)

	fileName := track.SongName + ".mp3"
	fmt.Println("Downloading to", fileName)

	// TODO: check file existence first with io.IsExist
	output, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error while creating", fileName, "-", err)
		return
	}
	defer output.Close()

	response, err := sk.Download()
	if err != nil {
		fmt.Println("Error while downloading -", err)
		return
	}
	defer response.Body.Close()

	n, err := io.Copy(output, response.Body)
	if err != nil {
		fmt.Println("Error while downloading -", err)
		return
	}

	fmt.Println(n, "bytes downloaded.")
}
