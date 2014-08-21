package grooveshark

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/stayradiated/grooveshark/requests"
	"github.com/stayradiated/grooveshark/responses"
)

type StreamKey struct {
	client    *Client
	Ip        string
	SongId    int
	StreamKey string
}

func (c *Client) GetStreamKey(songId int) *StreamKey {
	sk := c.getStreamKeyFromSongIDEx(songId)
	return newStreamKey(sk.Ip, sk.StreamKey, songId, c)
}

func newStreamKey(ip, key string, songId int, client *Client) (streamKey *StreamKey) {
	streamKey = &StreamKey{
		Ip:        ip,
		client:    client,
		SongId:    songId,
		StreamKey: key,
	}
	return streamKey
}

func (sk *StreamKey) Download() (*http.Response, error) {
	sk.client.markSongDownloadedEx(sk)
	url := "http://" + sk.Ip + "/stream.php?streamKey=" + sk.StreamKey
	return http.Get(url)
}

func (c *Client) getStreamKeyFromSongIDEx(songId int) responses.StreamKey {
	var resp responses.GetStreamKeyFromSongIDEx

	c.CallMethod("getStreamKeyFromSongIDEx", requests.GetStreamKeyFromSongIDEx{
		SongId:   songId,
		Type:     0,
		Prefetch: false,
		Mobile:   false,
		ReturnTS: false,
		Country:  c.session.Country,
	}, &resp)

	return resp.Result
}

func (c *Client) markSongDownloadedEx(sk *StreamKey) {
	var resp responses.MarkSongDownloadedEx

	c.CallMethod("markSongDownloadedEx", requests.MarkSongDownloadedEx{
		SongId:         strconv.Itoa(sk.SongId),
		StreamKey:      sk.StreamKey,
		StreamServerId: sk.Ip,
	}, &resp)
}

func (c *Client) markStreamKeyOver30Seconds(sk *StreamKey) {
	var resp interface{}

	c.CallMethod("markStreamKeyOver30Seconds", requests.MarkStreamKeyOver30Seconds{
		SongId:         strconv.Itoa(sk.SongId),
		StreamKey:      sk.StreamKey,
		StreamServerId: sk.Ip,
	}, &resp)

	fmt.Println(resp)
}
