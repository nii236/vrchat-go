package client_test

import (
	"github.com/nii236/vrchat-go/client"
	"os"
	"testing"
)

func TestClient_Authenticate(t *testing.T) {
	c, err := client.NewClient("https://api.vrchat.cloud/api/1", "", "")
	if err != nil {
		t.Errorf("unexpected err: %v", err)
	}
	user := os.Getenv("VRCHAT_USER")
	pass := os.Getenv("VRCHAT_PASS")
	_, err = c.Authenticate(user, pass)
	if err != nil {
		t.Errorf("unexpected err: %v", err)
	}
}
