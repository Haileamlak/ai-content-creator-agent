package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const NewsBaseURL = "https://newsapi.org/v2/"

type NewsAPI struct {
	APIKey string
}

func NewNewsAPI(apiKey string) *NewsAPI {
	return &NewsAPI{APIKey: apiKey}
}

func (n *NewsAPI) GetTopHeadlines() ([]string, error) {
	url := fmt.Sprintf("%s/top-headlines?country=us&apiKey=%s", NewsBaseURL, n.APIKey)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	// result should a list of the description of the news
	// e.g. ["description1", "description2", ...]
	var descriptions []string
	for _, article := range data["articles"].([]interface{}) {
		descriptions = append(descriptions, article.(map[string]interface{})["description"].(string))
	}

	return descriptions, nil
}

// SearchNews searches for news based on the query from the last 24 hours
func (n *NewsAPI) SearchNews(query string, limit int) ([]string, error) {
	fromTime := time.Now().AddDate(0, 0, -1).Format("2006-01-02")
	url := fmt.Sprintf("%severything?q=%s&from=%s&sortBy=relevance&apiKey=%s&pageSize=%d", NewsBaseURL, query, fromTime, n.APIKey, limit)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	// result should a list of the description of the news
	// e.g. ["description1", "description2", ...]
	var descriptions []string
	for _, article := range data["articles"].([]interface{}) {
		descriptions = append(descriptions, article.(map[string]interface{})["description"].(string))
	}

	return descriptions, nil
}
