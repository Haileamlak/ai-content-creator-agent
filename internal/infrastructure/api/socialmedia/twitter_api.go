package api

import (
	"fmt"
	"io"
	"net/http"
)

type TwitterAPI struct {
	BearerToken string
}

func NewTwitterAPI(bearerToken string) *TwitterAPI {
	return &TwitterAPI{BearerToken: bearerToken}
}

// Example function to post a tweet
func (api *TwitterAPI) PostTweet(tweet string) error {
	// Simulated call to Twitter API
	fmt.Printf("Posting tweet: %s\n", tweet)
	return nil
}

// Example function to get user timeline
func (api *TwitterAPI) GetUserTimeline(username string) (string, error) {
	url := fmt.Sprintf("https://api.twitter.com/2/tweets?username=%s", username)

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", "Bearer "+ api.BearerToken)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}

	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)
	return string(body), nil
}

// Function to get trending hashtags for an industry
func (api *TwitterAPI) GetTrendingHashtags(industry string) ([]string, error) {
	url := fmt.Sprintf("https://api.twitter.com/2/trends?industry=%s", industry)

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", "Bearer "+ api.BearerToken)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)
	return []string{string(body)}, nil
}