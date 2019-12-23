package client_test

import (
	"testing"
	"vrchat-go/client"
)

const testWorldID = "wrld_6caf5200-70e1-46c2-b043-e3c4abe69e0f"
const testWorldInstanceID = "98573"

func TestClient_WorldList(t *testing.T) {
	c, err := client.NewClient(client.ReleaseAPIURL)
	if err != nil {
		t.Errorf("unexpected err: %v", err)
	}
	_, err = c.Authenticate()
	if err != nil {
		t.Errorf("unexpected err: %v", err)
	}
	_, err = c.WorldList()
	if err != nil {
		t.Errorf("unexpected err: %v", err)
	}
	// fmt.Printf("%+v", client.MustEncodeJSON(worlds))
}

func TestClient_WorldGet(t *testing.T) {
	c, err := client.NewClient(client.ReleaseAPIURL)
	if err != nil {
		t.Errorf("unexpected err: %v", err)
	}
	_, err = c.Authenticate()
	if err != nil {
		t.Errorf("unexpected err: %v", err)
	}
	_, err = c.WorldGet(testWorldID)
	if err != nil {
		t.Errorf("unexpected err: %v", err)
	}
	// fmt.Printf("%+v\n", client.MustEncodeJSON(world))
}
