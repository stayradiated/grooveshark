package grooveshark

import (
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

func (c *Client) GetStreamKey(songId int) (*StreamKey, error) {
	sk, err := c.getStreamKeyFromSongIDEx(songId)
	if err != nil {
		return nil, err
	}
	return newStreamKey(sk.Ip, sk.StreamKey, songId, c), nil
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

func (sk *StreamKey) Url() string {
	return "http://" + sk.Ip + "/stream.php?streamKey=" + sk.StreamKey
}

func (sk *StreamKey) Download() (*http.Response, error) {
	// must call markSongDownloadedEx before downloading song
	go sk.client.markSongDownloadedEx(sk)
	return http.Get(sk.Url())
}

func (c *Client) getStreamKeyFromSongIDEx(songId int) (responses.StreamKey, error) {
	var resp responses.GetStreamKeyFromSongIDEx

	err := c.CallMethod("getStreamKeyFromSongIDEx", requests.GetStreamKeyFromSongIDEx{
		SongId:   songId,
		Type:     0,
		Prefetch: false,
		Mobile:   false,
		Country:  c.session.Country,
	}, &resp)

	if err != nil {
		return responses.StreamKey{}, err
	}

	return resp.Result, nil
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
}
