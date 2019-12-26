package client_test

import (
	"github.com/nii236/vrchat-go/client"
	"testing"
)

func TestClient_FriendList(t *testing.T) {
	c, err := client.NewClient("https://api.vrchat.cloud/api/1", "", "")
	if err != nil {
		t.Errorf("unexpected err: %v", err)
	}
	_, err = c.FriendList(true)
	if err != nil {
		t.Errorf("unexpected err: %v", err)
	}
	// fmt.Printf("%+v", client.MustEncodeJSON(friends))
}
