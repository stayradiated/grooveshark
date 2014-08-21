package grooveshark

import (
	"fmt"

	"github.com/stayradiated/grooveshark/session"
)

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
	fmt.Println("We are online")
}

func (c *Client) CallMethod(method string, parameters interface{}, resp interface{}) {
	request := session.NewRequest(c.session, method, parameters)
	request.Sign()
	request.Send(resp)
}
