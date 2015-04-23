package grooveshark

import (
	"github.com/stayradiated/grooveshark/responses"
	"github.com/stayradiated/grooveshark/session"
)

type Interface interface {
	Connect()
	CallMethod(method string, parameters interface{}, resp interface{}) (err error)
	Playlist(playlistId int) (playlist responses.Playlist)
	Search(query string) (tracks []responses.Track)
	LookupTrackIds(trackIds []int) (tracks []responses.Track)
	GetStreamKey(songId int) (*StreamKey, error)
}

type Client struct {
	session *session.Session
}

func NewClient() (client *Client) {
	return &Client{
		session: session.NewSession(session.HtmlSharkSession),
	}
}

func (c *Client) Connect() {
	c.session.Initiate()
}

func (c *Client) CallMethod(method string, parameters interface{}, resp interface{}) (err error) {
	request := session.NewRequest(c.session, method, parameters)
	request.Sign()
	return request.Send(resp)
}
