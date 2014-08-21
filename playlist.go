package grooveshark

import (
	"strconv"

	"github.com/stayradiated/grooveshark/requests"
	"github.com/stayradiated/grooveshark/responses"
)

func (c *Client) Playlist(playlistId int) responses.Playlist {
	var resp responses.GetPlaylistByID

	c.CallMethod("getPlaylistByID", &requests.GetPlaylistByID{
		PlaylistId: strconv.Itoa(playlistId),
	}, &resp)

	return resp.Result
}
