package client_test

import (
	"testing"
	"vrchat-go/client"
)

func TestClient_Authenticate(t *testing.T) {
	c, err := client.NewClient("https://api.vrchat.cloud/api/1")
	if err != nil {
		t.Errorf("unexpected err: %v", err)
	}
	err = c.Authenticate()
	if err != nil {
		t.Errorf("unexpected err: %v", err)
	}
}
