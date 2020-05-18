package feed

import (
	"encoding/json"
	"os"
)

type Feed struct {
	Name string `json:"name"`
	URI  string `json:"link"`
	Type string `json:"type"`
}

func RetrieveFeeds(feedDefintionFile string) ([]*Feed, error) {
	file, err := os.Open(feedDefintionFile)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	var feeds []*Feed
	err = json.NewDecoder(file).Decode(&feeds)

	return feeds, err
}
