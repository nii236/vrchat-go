package client_test

import (
	"testing"
	"vrchat-go/client"
)

func TestClient_FriendList(t *testing.T) {
	c, err := client.NewClient(client.ReleaseAPIURL)
	if err != nil {
		t.Errorf("unexpected err: %v", err)
	}
	_, err = c.Authenticate()
	if err != nil {
		t.Errorf("unexpected err: %v", err)
	}
	_, err = c.FriendList(true)
	if err != nil {
		t.Errorf("unexpected err: %v", err)
	}
	// fmt.Printf("%+v", client.MustEncodeJSON(friends))
}
