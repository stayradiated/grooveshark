package grooveshark

import (
	"github.com/stayradiated/grooveshark/requests"
	"github.com/stayradiated/grooveshark/responses"
)

func (c *Client) Search(query, searchType string) responses.TrackList {
	var resp responses.GetResultsFromSearch

	c.CallMethod("getResultsFromSearch", requests.GetResultsFromSearch{
		Guts:       0,
		PPOverride: false,
		Query:      query,
		Type:       searchType,
	}, &resp)

	return resp.Result.Result
}
