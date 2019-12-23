package client

import (
	"encoding/json"
	"net/http"
	"time"
)

func (c *Client) FriendList(showOffline bool) ([]*FriendListItem, error) {
	req, err := http.NewRequest("GET", buildGetFriendsURL(c.baseURL), nil)
	if err != nil {
		return nil, err
	}
	if showOffline {
		q := req.URL.Query()
		q.Add("offline", "true")
		req.URL.RawQuery = q.Encode()
	}
	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Close()
	result := []*FriendListItem{}
	err = json.NewDecoder(resp).Decode(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type FriendListItem struct {
	ID                             string    `json:"id"`
	Username                       string    `json:"username"`
	DisplayName                    string    `json:"displayName"`
	Bio                            string    `json:"bio"`
	CurrentAvatarImageURL          string    `json:"currentAvatarImageUrl"`
	CurrentAvatarThumbnailImageURL string    `json:"currentAvatarThumbnailImageUrl"`
	LastPlatform                   string    `json:"last_platform"`
	Tags                           []string  `json:"tags"`
	DeveloperType                  string    `json:"developerType"`
	Status                         string    `json:"status"`
	StatusDescription              string    `json:"statusDescription"`
	FriendKey                      string    `json:"friendKey"`
	LastLogin                      time.Time `json:"last_login"`
	IsFriend                       bool      `json:"isFriend"`
	Location                       string    `json:"location"`
}
