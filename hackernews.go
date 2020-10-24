package hackernews

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"strconv"

	"golang.org/x/sync/errgroup"
)

const (
	// HackerNews API URL
	urlScheme   = "https"
	urlHost     = "hacker-news.firebaseio.com"
	urlBasePath = "v0"

	// HackerNews API paths
	itemPath        = "item"
	maxItemPath     = "maxitem"
	userPath        = "user"
	topStoriesPath  = "topstories"
	bestStoriesPath = "beststories"
	newStoriesPath  = "newstories"
	askStoriesPath  = "askstories"
	showStoriesPath = "showstories"
	jobStoriesPath  = "jobstories"
)

// HydrateItems - Hydrates a list of item ids in concurrently
func HydrateItems(itemIDs []int) ([]Item, error) {
	var g errgroup.Group
	items := make([]Item, len(itemIDs))

	for i, itemID := range itemIDs {
		i, itemID := i, itemID
		g.Go(func() error {
			item, err := GetItem(itemID)
			if err == nil {
				items[i] = item
			}
			return err
		})
	}
	if err := g.Wait(); err != nil {
		return nil, err
	}

	return items, nil
}

// GetItem -
func GetItem(itemID int) (Item, error) {
	// Build url
	url := buildRequestURL(itemPath, strconv.Itoa(itemID))

	// Call endpoint
	response, err := get(url)
	if err != nil {
		return Item{}, err
	}

	// Unmarshal to output type
	var item Item
	err = json.Unmarshal(response, &item)
	if err != nil {
		return Item{}, err
	}
	return item, nil
}

// GetUser -
func GetUser(userID string) (User, error) {
	// Build url
	url := buildRequestURL(userPath, userID)

	// Call endpoint
	response, err := get(url)
	if err != nil {
		return User{}, err
	}

	// Unmarshal to output type
	var user User
	err = json.Unmarshal(response, &user)
	if err != nil {
		return User{}, err
	}
	return user, nil
}

// GetTopStories -
func GetTopStories() (TopStories, error) {
	// Build url
	url := buildRequestURL(topStoriesPath)

	// Call endpoint
	response, err := get(url)
	if err != nil {
		return TopStories{}, err
	}

	// Unmarshal to output type
	var topStories TopStories
	err = json.Unmarshal(response, &topStories)
	if err != nil {
		return TopStories{}, err
	}
	return topStories, nil
}

// GetNewStories -
func GetNewStories() (NewStories, error) {
	// Build url
	url := buildRequestURL(newStoriesPath)

	// Call endpoint
	response, err := get(url)
	if err != nil {
		return NewStories{}, err
	}

	// Unmarshal to output type
	var newStories NewStories
	err = json.Unmarshal(response, &newStories)
	if err != nil {
		return NewStories{}, err
	}
	return newStories, nil
}

// GetBestStories -
func GetBestStories() (BestStories, error) {
	// Build url
	url := buildRequestURL(newStoriesPath)

	// Call endpoint
	response, err := get(url)
	if err != nil {
		return BestStories{}, err
	}

	// Unmarshal to output type
	var bestStories BestStories
	err = json.Unmarshal(response, &bestStories)
	if err != nil {
		return BestStories{}, err
	}
	return bestStories, nil
}

// GetAskStories -
func GetAskStories() (AskStories, error) {
	// Build url
	url := buildRequestURL(askStoriesPath)

	// Call endpoint
	response, err := get(url)
	if err != nil {
		return AskStories{}, err
	}

	// Unmarshal to output type
	var askStories AskStories
	err = json.Unmarshal(response, &askStories)
	if err != nil {
		return AskStories{}, err
	}
	return askStories, nil
}

// GetShowStories -
func GetShowStories() (ShowStories, error) {
	// Build url
	url := buildRequestURL(showStoriesPath)

	// Call endpoint
	response, err := get(url)
	if err != nil {
		return ShowStories{}, err
	}

	// Unmarshal to output type
	var showStories ShowStories
	err = json.Unmarshal(response, &showStories)
	if err != nil {
		return ShowStories{}, err
	}
	return showStories, nil
}

// GetJobStories -
func GetJobStories() (JobStories, error) {
	// Build url
	url := buildRequestURL(jobStoriesPath)

	// Call endpoint
	response, err := get(url)
	if err != nil {
		return JobStories{}, err
	}

	// Unmarshal to output type
	var jobStories JobStories
	err = json.Unmarshal(response, &jobStories)
	if err != nil {
		return JobStories{}, err
	}
	return jobStories, nil
}

func get(url string) ([]byte, error) {
	// Call endpoint
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Extract response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func buildRequestURL(paths ...string) string {
	finalPaths := []string{urlBasePath}
	for _, path := range paths {
		finalPaths = append(finalPaths, path)
	}
	url := url.URL{
		Scheme: urlScheme,
		Host:   urlHost,
		Path:   path.Join(finalPaths...),
	}
	return url.String() + ".json"
}
