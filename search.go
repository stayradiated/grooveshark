package grooveshark

import (
	"github.com/stayradiated/grooveshark/requests"
	"github.com/stayradiated/grooveshark/responses"
)

func (c *Client) Search(query string) (tracks []responses.Track) {
	var resp responses.GetResultsFromSearch

	c.CallMethod("getResultsFromSearch", requests.GetResultsFromSearch{
		Guts:       0,
		PPOverride: false,
		Query:      query,
		Type:       "Songs",
	}, &resp)

	return resp.Result.Result.Tracks()
}
