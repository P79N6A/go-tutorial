package search

import (
	"os"
	"encoding/json"
)

const dataFile = "goinaction/data/data.json"

type Feed struct {
	Name string `json:"site"`    // `json:"site"`是go中的tag,描述了一些元数据
	URI  string `json:"link"`
	Type string `json:"type"`
}

func RetrieveFeeds() ([]*Feed, error){
	file, err := os.Open(dataFile)
	if err != nil {
		return nil,err
	}
	defer file.Close()

	var feeds []*Feed
	err = json.NewDecoder(file).Decode(&feeds)

	return feeds, err
}