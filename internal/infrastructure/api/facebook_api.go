package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const (
	FacebookBaseURL = "https://graph.facebook.com/v20.0"
)

// FacebookAPIService interacts with the Facebook API for a page
type FacebookAPIService struct {
}

// NewFacebookAPIService creates a new service with the access token and page ID
func NewFacebookAPIService() *FacebookAPIService {
	return &FacebookAPIService{}
}

// CreatePost creates a post on a Facebook page
func (service *FacebookAPIService) CreatePost(message string, pageID, accessToken string) (string, error) {
	url := fmt.Sprintf("%s/%s/feed", FacebookBaseURL, pageID)

	postData := map[string]string{
		"message":      message,
		"access_token": accessToken,
	}

	jsonData, err := json.Marshal(postData)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to create post, status code: %d", resp.StatusCode)
	}

	// retrieve post ID from response
	var result struct {
		ID string `json:"id"`
	}

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return "", err
	}

	return result.ID, nil
}

// GetPostsComments retrieves comments for all posts of a page in the last 24 hours
func (service *FacebookAPIService) GetComments(pageID, accessToken string) ([]map[string]interface{}, error) {
	yesterday := time.Now().AddDate(0, 0, -1).Unix()
	url := fmt.Sprintf("%s/%s/posts?fields=comments.since(%d)&access_token=%s", FacebookBaseURL, pageID, yesterday, accessToken)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get posts comments, status code: %d", resp.StatusCode)
	}

	var result struct {
		Data []map[string]interface{} `json:"data"`
	}

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result.Data, nil
}

// ReplyToComment replies to comments for posts of a page in the last 24 hours
func (service *FacebookAPIService) ReplyToComment(commentID, replyText string, accessToken string) error {
	url := fmt.Sprintf("%s/%s/comments", FacebookBaseURL, commentID)

	replyData := map[string]string{
		"message":      replyText,
		"access_token": accessToken,
	}

	jsonData, err := json.Marshal(replyData)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to reply to comment, status code: %d", resp.StatusCode)
	}

	return nil
}

// Get Page Post Insights
func (service *FacebookAPIService) GetPagePostInsights(postID, accessToken string) ([]map[string]interface{}, error) {
	url := fmt.Sprintf("%s/%s/insights?metric=post_impressions,post_engagements&access_token=%s", FacebookBaseURL, postID, accessToken)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get post insights, status code: %d", resp.StatusCode)
	}

	var result struct {
		Data []map[string]interface{} `json:"data"`
	}

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result.Data, nil
}