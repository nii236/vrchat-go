package client

import (
	"encoding/json"
	"net/http"
	"time"
)

func (c *Client) WorldList() ([]*WorldListItem, error) {
	req, err := http.NewRequest("GET", buildListWorldsURL(c.baseURL), nil)
	if err != nil {
		return nil, err
	}
	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Close()
	result := []*WorldListItem{}
	err = json.NewDecoder(resp).Decode(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
func (c *Client) WorldGet(worldID string) (*World, error) {
	req, err := http.NewRequest("GET", buildGetWorldURL(c.baseURL, worldID), nil)
	if err != nil {
		return nil, err
	}
	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Close()
	result := &World{}
	err = json.NewDecoder(resp).Decode(result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type WorldListItem struct {
	ID                  string    `json:"id"`
	Name                string    `json:"name"`
	AuthorID            string    `json:"authorId"`
	AuthorName          string    `json:"authorName"`
	Capacity            int       `json:"capacity"`
	ImageURL            string    `json:"imageUrl"`
	ThumbnailImageURL   string    `json:"thumbnailImageUrl"`
	ReleaseStatus       string    `json:"releaseStatus"`
	Organization        string    `json:"organization"`
	Tags                []string  `json:"tags"`
	Favorites           int       `json:"favorites"`
	CreatedAt           time.Time `json:"created_at"`
	UpdatedAt           time.Time `json:"updated_at"`
	PublicationDate     time.Time `json:"publicationDate"`
	LabsPublicationDate string    `json:"labsPublicationDate"`
	Visits              int       `json:"visits"`
	UnityPackages       []struct {
		Platform     string `json:"platform"`
		UnityVersion string `json:"unityVersion"`
	} `json:"unityPackages"`
	Popularity int `json:"popularity"`
	Heat       int `json:"heat"`
	Occupants  int `json:"occupants"`
}

type World struct {
	ID                string   `json:"id"`
	Name              string   `json:"name"`
	Description       string   `json:"description"`
	Featured          bool     `json:"featured"`
	AuthorID          string   `json:"authorId"`
	AuthorName        string   `json:"authorName"`
	Capacity          int      `json:"capacity"`
	Tags              []string `json:"tags"`
	ReleaseStatus     string   `json:"releaseStatus"`
	ImageURL          string   `json:"imageUrl"`
	ThumbnailImageURL string   `json:"thumbnailImageUrl"`
	AssetURL          string   `json:"assetUrl"`
	AssetURLObject    struct {
	} `json:"assetUrlObject"`
	PluginURL       string `json:"pluginUrl"`
	PluginURLObject struct {
	} `json:"pluginUrlObject"`
	UnityPackageURL       string `json:"unityPackageUrl"`
	UnityPackageURLObject struct {
	} `json:"unityPackageUrlObject"`
	Namespace           string `json:"namespace"`
	UnityPackageUpdated bool   `json:"unityPackageUpdated"`
	UnityPackages       []struct {
		ID             string `json:"id"`
		AssetURL       string `json:"assetUrl"`
		AssetURLObject struct {
		} `json:"assetUrlObject"`
		PluginURL       string `json:"pluginUrl"`
		PluginURLObject struct {
		} `json:"pluginUrlObject"`
		UnityVersion    string    `json:"unityVersion"`
		UnitySortNumber int64     `json:"unitySortNumber"`
		AssetVersion    int       `json:"assetVersion"`
		Platform        string    `json:"platform"`
		CreatedAt       time.Time `json:"created_at"`
	} `json:"unityPackages"`
	Version             int             `json:"version"`
	Organization        string          `json:"organization"`
	PreviewYoutubeID    interface{}     `json:"previewYoutubeId"`
	Favorites           int             `json:"favorites"`
	CreatedAt           time.Time       `json:"created_at"`
	UpdatedAt           time.Time       `json:"updated_at"`
	PublicationDate     time.Time       `json:"publicationDate"`
	LabsPublicationDate string          `json:"labsPublicationDate"`
	Visits              int             `json:"visits"`
	Popularity          int             `json:"popularity"`
	Heat                int             `json:"heat"`
	PublicOccupants     int             `json:"publicOccupants"`
	PrivateOccupants    int             `json:"privateOccupants"`
	Occupants           int             `json:"occupants"`
	Instances           [][]interface{} `json:"instances"`
}
