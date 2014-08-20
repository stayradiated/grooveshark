package grooveshark

import (
	"fmt"
	netUrl "net/url"
)

type Client struct {
	session *Session
}

func NewClient() (client *Client) {
	return &Client{
		session: NewSession(HtmlSharkSession),
	}
}

func (c *Client) Connect() {
	c.session.Initiate()
	fmt.Println("We are online")
}

func (c *Client) CallMethod(method string, parameters *Parameters) {
	request := NewRequest(c.session, method, parameters)
	request.Sign()
	request.Send()
}

func (c *Client) DownloadTrack(ip, streamKey string) {
	url := "/stream.php?streamKey=" + netUrl.QueryEscape(streamKey)
	fmt.Println(url)
}
