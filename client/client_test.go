package client_test

import (
	"testing"
	"vrchat-go/client"
)

const testUser = "TESTUSER"
const testPass = "TESTPASS"

func TestClient_Authenticate(t *testing.T) {
	c, err := client.NewClient("https://api.vrchat.cloud/api/1")
	if err != nil {
		t.Errorf("unexpected err: %v", err)
	}
	err = c.Authenticate(testUser, testPass)
	if err != nil {
		t.Errorf("unexpected err: %v", err)
	}
}
