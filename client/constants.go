package client

import "fmt"

const ReleaseAPIURL = "https://api.vrchat.cloud/api/1"

func buildGetFriendsURL(baseURL string) string {
	return fmt.Sprintf("%s/auth/user/friends", baseURL)
}

func buildAuthURL(baseURL string) string {
	return fmt.Sprintf("%s/auth/user", baseURL)
}
func buildConfigURL(baseURL string) string {
	return fmt.Sprintf("%s/config", baseURL)
}

func buildListWorldsURL(baseURL string) string {
	return fmt.Sprintf("%s/worlds", baseURL)
}

func buildGetWorldURL(baseURL string, worldID string) string {
	return fmt.Sprintf("%s/worlds/%s", baseURL, worldID)
}

func buildGetWorldInstanceURL(baseURL, worldID, instanceID string) string {
	return fmt.Sprintf("%s/worlds/%s/%s", baseURL, worldID, instanceID)
}
