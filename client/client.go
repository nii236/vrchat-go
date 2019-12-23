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
)

type Client struct {
	baseURL string
	*http.Client
	// apiKey    string
	// authToken string
}

func NewClient(baseURL string) (*Client, error) {
	jar, err := cookiejar.New(nil)
	if err != nil {
		return nil, err
	}
	httpClient := &http.Client{
		Jar: jar,
	}
	c := &Client{baseURL, httpClient}
	return c, nil
}

func (c *Client) Authenticate(user, pass string) error {
	req, err := http.NewRequest("GET", buildAuthURL(c.baseURL), nil)
	if err != nil {
		return err
	}
	req.SetBasicAuth(user, pass)
	resp, err := c.Client.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		return fmt.Errorf("non 200 response: %v %v", resp.StatusCode, string(b))
	}
	return nil
}

func (c *Client) Do(req *http.Request) (io.ReadCloser, error) {
	u, err := url.Parse(ReleaseAPIURL)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	apiKey := ""
	for _, cookie := range c.Client.Jar.Cookies(u) {
		if cookie.Name == "apiKey" {
			apiKey = cookie.Value
		}
	}
	if apiKey == "" {
		return nil, errors.New("no apikey in cookie")
	}
	q.Set("apiKey", apiKey)
	req.URL.RawQuery = q.Encode()
	fmt.Println(req.URL)
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
