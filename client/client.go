package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"time"
)

// ErrNotImplemented is used to stub empty funcs
var ErrNotImplemented = errors.New("not implemented")

type Client struct {
	baseURL string
	*http.Client
	authToken string
	apiKey    string
}

func NewClient(baseURL, authToken, apiKey string) (*Client, error) {
	jar, err := cookiejar.New(nil)
	if err != nil {
		return nil, err
	}
	httpClient := &http.Client{
		Jar: jar,
	}
	if os.Getenv("VRCHAT_AUTH_TOKEN") != "" {
		authToken = os.Getenv("VRCHAT_AUTH_TOKEN")
	}

	if os.Getenv("VRCHAT_API_KEY") != "" {
		apiKey = os.Getenv("VRCHAT_API_KEY")
	}

	c := &Client{baseURL, httpClient, authToken, apiKey}
	return c, nil
}

func Token(baseURL, user, pass string) (*AuthResponse, string, string, error) {
	req, err := http.NewRequest("GET", buildAuthURL(baseURL), nil)
	if err != nil {
		return nil, "", "", err
	}
	jar, err := cookiejar.New(nil)
	if err != nil {
		return nil, "", "", err
	}
	httpClient := &http.Client{
		Jar: jar,
	}
	req.SetBasicAuth(user, pass)
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, "", "", err
	}
	if resp.StatusCode != 200 {
		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, "", "", err
		}
		return nil, "", "", fmt.Errorf("non 200 response: %v %v", resp.StatusCode, string(b))
	}
	u, err := url.Parse(ReleaseAPIURL)
	if err != nil {
		return nil, "", "", err
	}
	apiKey := ""
	authToken := ""
	for _, cookie := range httpClient.Jar.Cookies(u) {
		if cookie.Name == "apiKey" {
			apiKey = cookie.Value
		}
		if cookie.Name == "auth" {
			authToken = cookie.Value
		}
	}
	data := &AuthResponse{}
	err = json.NewDecoder(resp.Body).Decode(data)
	if err != nil {
		return nil, "", "", err
	}
	return data, apiKey, authToken, nil
}
func (c *Client) RemoteConfig() (*RemoteConfigResponse, error) {
	return nil, ErrNotImplemented
}

type RemoteConfigResponse struct {
}
type AuthResponse struct {
	ID                             string        `json:"id"`
	Username                       string        `json:"username"`
	DisplayName                    string        `json:"displayName"`
	Bio                            string        `json:"bio"`
	BioLinks                       []interface{} `json:"bioLinks"`
	PastDisplayNames               []interface{} `json:"pastDisplayNames"`
	HasEmail                       bool          `json:"hasEmail"`
	HasPendingEmail                bool          `json:"hasPendingEmail"`
	Email                          string        `json:"email"`
	ObfuscatedEmail                string        `json:"obfuscatedEmail"`
	ObfuscatedPendingEmail         string        `json:"obfuscatedPendingEmail"`
	EmailVerified                  bool          `json:"emailVerified"`
	HasBirthday                    bool          `json:"hasBirthday"`
	Unsubscribe                    bool          `json:"unsubscribe"`
	Friends                        []string      `json:"friends"`
	FriendGroupNames               []interface{} `json:"friendGroupNames"`
	CurrentAvatarImageURL          string        `json:"currentAvatarImageUrl"`
	CurrentAvatarThumbnailImageURL string        `json:"currentAvatarThumbnailImageUrl"`
	CurrentAvatar                  string        `json:"currentAvatar"`
	CurrentAvatarAssetURL          string        `json:"currentAvatarAssetUrl"`
	AcceptedTOSVersion             int           `json:"acceptedTOSVersion"`
	SteamID                        string        `json:"steamId"`
	SteamDetails                   struct {
	} `json:"steamDetails"`
	OculusID              string `json:"oculusId"`
	HasLoggedInFromClient bool   `json:"hasLoggedInFromClient"`
	HomeLocation          string `json:"homeLocation"`
	TwoFactorAuthEnabled  bool   `json:"twoFactorAuthEnabled"`
	Feature               struct {
		TwoFactorAuth bool `json:"twoFactorAuth"`
	} `json:"feature"`
	Status             string        `json:"status"`
	StatusDescription  string        `json:"statusDescription"`
	State              string        `json:"state"`
	Tags               []interface{} `json:"tags"`
	DeveloperType      string        `json:"developerType"`
	LastLogin          time.Time     `json:"last_login"`
	LastPlatform       string        `json:"last_platform"`
	AllowAvatarCopying bool          `json:"allowAvatarCopying"`
	IsFriend           bool          `json:"isFriend"`
	FriendKey          string        `json:"friendKey"`
	OnlineFriends      []interface{} `json:"onlineFriends"`
	ActiveFriends      []interface{} `json:"activeFriends"`
	OfflineFriends     []string      `json:"offlineFriends"`
}

func (c *Client) Do(req *http.Request) (io.ReadCloser, error) {
	u, err := url.Parse(ReleaseAPIURL)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	if c.apiKey == "" {
		return nil, errors.New("no apikey in client")
	}
	authTokenCookie := &http.Cookie{Name: "auth", Value: c.authToken}
	apiKeyCookie := &http.Cookie{Name: "apiKey", Value: c.apiKey}
	c.Jar.SetCookies(u, []*http.Cookie{authTokenCookie, apiKeyCookie})
	q.Set("apiKey", c.apiKey)
	req.URL.RawQuery = q.Encode()
	resp, err := c.Client.Do(req)
	if resp.StatusCode != 200 {
		result := &ErrorResponse{}
		err = json.NewDecoder(resp.Body).Decode(result)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("non 200 response: %v %w", resp.StatusCode, result)
	}
	return resp.Body, nil
}
