//Package wiki provides primitives for retrieving WikiPedia pages
//as a WikiPage struct
package wiki

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"regexp"
)

const wikiURL = "https://en.wikipedia.org/w/api.php"

type apiResponse struct {
	Query struct {
		Pages map[string]struct {
			PageID  int    `json:"pageid"`
			Title   string `json:"title"`
			Extract string `json:"extract"`
		} `json:"pages"`
	} `json:"query"`
}

func getFromWiki(pageID string) (*http.Response, error) {
	url, err := url.Parse(wikiURL)
	if err != nil {
		return nil, err
	}
	query := url.Query()
	query.Set("action", "query")
	query.Set("prop", "extracts")
	query.Set("explaintext", "true")
	query.Set("format", "json")
	query.Set("pageids", pageID)
	url.RawQuery = query.Encode()

	return http.Get(url.String())
}

func decodeWikiJSON(wikiIO *io.ReadCloser) (*apiResponse, error) {
	decoder := json.NewDecoder(*wikiIO)
	var apiR apiResponse
	err := decoder.Decode(&apiR)
	if err != nil {
		return nil, err
	}
	return &apiR, nil
}

// GetPage - retrieves and parses a WikiPage
func GetPage(pageID string) (*Page, error) {
	resp, err := getFromWiki(pageID)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	apiR, err := decodeWikiJSON(&resp.Body)
	if err != nil {
		return nil, err
	}

	// If Page is missing
	if apiR.Query.Pages[pageID].Extract == "" {
		return nil, errors.New("Page " + pageID + " does not exist")
	}

	wordRegexp := regexp.MustCompile("\\w{4,}")
	words := wordRegexp.FindAllString(apiR.Query.Pages[pageID].Extract, -2)
	return &Page{apiR.Query.Pages[pageID].Title, words}, nil
}

// Page a structure for a wikipage
type Page struct {
	Title string
	Words []string
}
