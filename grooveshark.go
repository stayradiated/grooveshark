package grooveshark

import "github.com/stayradiated/grooveshark/session"

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

func (c *Client) CallMethod(method string, parameters interface{}, resp interface{}) error {
	request := session.NewRequest(c.session, method, parameters)
	request.Sign()
	return request.Send(resp)
}
