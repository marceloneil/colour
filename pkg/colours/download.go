package colours

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const wikipediaArticlePrefix = "https://en.wikipedia.org/wiki"
const colourListArticle = "List_of_colors_(compact)"

func downloadColourArticle() (article string, err error) {
	query := url.Values{
		"action": {"raw"},
	}
	listURL := fmt.Sprintf("%s/%s?%s",
		wikipediaArticlePrefix, colourListArticle, query.Encode())
	resp, err := http.Get(listURL)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
