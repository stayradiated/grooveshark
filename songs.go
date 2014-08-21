package grooveshark

import (
	"strconv"

	"github.com/stayradiated/grooveshark/requests"
	"github.com/stayradiated/grooveshark/responses"
)

func (c *Client) LookupTrackIds(trackIds []int) []responses.Track {
	var resp responses.GetQueueSongListFromSongIDs

	tracks := make([]string, len(trackIds))
	for i := range trackIds {
		tracks[i] = strconv.Itoa(trackIds[i])
	}

	c.CallMethod("getQueueSongListFromSongIDs", requests.GetQueueSongListFromSongIDs{
		SongIds: tracks,
	}, &resp)

	return resp.Result.Tracks()
}
