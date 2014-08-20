package grooveshark

import "testing"

func TestSignMethod(t *testing.T) {
	client := NewClient()
	client.signMethod("test")
}
