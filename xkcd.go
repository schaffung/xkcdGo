package xkcdGo

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const baseUrlPrefix = "https://xkcd.com/"
const baseUrlSuffix = "/info.0.json"

type date struct {
	Month string
	Year  string
	Day   string
}

type XKCDQuery struct {
	date
	Num        int
	Link       string
	News       string
	Transcript string
	Safe_title string
	Alt        string
	Img        string
	Title      string
}

func SearchComic(comicNum int) (*XKCDQuery, error) {
	query := fmt.Sprintf("%s%d%s", baseUrlPrefix, comicNum, baseUrlSuffix)
	resp, err := http.Get(query)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("Search query failed with : %s", resp.Status)
	}

	var result XKCDQuery
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}

	resp.Body.Close()
	return &result, nil
}
