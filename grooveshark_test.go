package grooveshark

import "testing"

func TestNewClient(t *testing.T) {
	client := NewClient()
	client.Connect()
}
