package client_test

import (
	"fmt"
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
	err = c.Authenticate(testUser, testPass)
	if err != nil {
		t.Errorf("unexpected err: %v", err)
	}
	worlds, err := c.WorldList()
	if err != nil {
		t.Errorf("unexpected err: %v", err)
	}
	fmt.Printf("%+v", client.MustEncodeJSON(worlds))
}
func TestClient_WorldInstance(t *testing.T) {
	c, err := client.NewClient(client.ReleaseAPIURL)
	if err != nil {
		t.Errorf("unexpected err: %v", err)
	}
	err = c.Authenticate(testUser, testPass)
	if err != nil {
		t.Errorf("unexpected err: %v", err)
	}

	instance, err := c.WorldInstance(testWorldID, testWorldInstanceID)
	if err != nil {
		t.Errorf("unexpected err: %v", err)
	}
	fmt.Printf("%+v", instance)
}
func TestClient_WorldGet(t *testing.T) {
	c, err := client.NewClient(client.ReleaseAPIURL)
	if err != nil {
		t.Errorf("unexpected err: %v", err)
	}
	err = c.Authenticate(testUser, testPass)
	if err != nil {
		t.Errorf("unexpected err: %v", err)
	}
	world, err := c.WorldGet(testWorldID)
	if err != nil {
		t.Errorf("unexpected err: %v", err)
	}
	fmt.Printf("%+v\n", client.MustEncodeJSON(world))
}
