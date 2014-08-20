package grooveshark

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

type Client struct {
	session *Session
	request *Request
}

func NewClient() (client *Client) {
	client = &Client{}
	client.session = NewSession()
	client.request = NewRequest(client.session)
	fmt.Println(client)
	return client
}

func (c *Client) Connect() {
	c.session.Connect()
	c.session.Check()
	fmt.Println("We are online")
}

func (c *Client) signMethod(method string) string {
	token := c.session.GetToken()
	nonce := make([]byte, 3)
	rand.Read(nonce)
	fmt.Println(token, hex.EncodeToString(nonce))
	return ""
}
