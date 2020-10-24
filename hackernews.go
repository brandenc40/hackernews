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
	updatesPath     = "updates"
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
	var item Item

	// Build url
	url := buildRequestURL(itemPath, strconv.Itoa(itemID))

	// Call endpoint
	response, err := get(url)
	if err != nil {
		return item, err
	}

	// Unmarshal to output type
	err = json.Unmarshal(response, &item)
	if err != nil {
		return item, err
	}
	return item, nil
}

// GetUser -
func GetUser(userID string) (User, error) {
	var user User

	// Build url
	url := buildRequestURL(userPath, userID)

	// Call endpoint
	response, err := get(url)
	if err != nil {
		return user, err
	}

	// Unmarshal to output type
	err = json.Unmarshal(response, &user)
	if err != nil {
		return user, err
	}
	return user, nil
}

// GetMaxItem -
func GetMaxItem() (MaxItem, error) {
	var maxItem MaxItem

	// Build url
	url := buildRequestURL(maxItemPath)

	// Call endpoint
	response, err := get(url)
	if err != nil {
		return maxItem, err
	}

	// Unmarshal to output type
	err = json.Unmarshal(response, &maxItem)
	if err != nil {
		return maxItem, err
	}
	return maxItem, nil
}

// GetUpates -
func GetUpates() (Updates, error) {
	var updates Updates

	// Build url
	url := buildRequestURL(updatesPath)

	// Call endpoint
	response, err := get(url)
	if err != nil {
		return updates, err
	}

	// Unmarshal to output type
	err = json.Unmarshal(response, &updates)
	if err != nil {
		return updates, err
	}
	return updates, nil
}

// GetTopStories -
func GetTopStories() (TopStories, error) {
	var topStories TopStories

	// Build url
	url := buildRequestURL(topStoriesPath)

	// Call endpoint
	response, err := get(url)
	if err != nil {
		return topStories, err
	}

	// Unmarshal to output type
	err = json.Unmarshal(response, &topStories)
	if err != nil {
		return topStories, err
	}
	return topStories, nil
}

// GetNewStories -
func GetNewStories() (NewStories, error) {
	var newStories NewStories

	// Build url
	url := buildRequestURL(newStoriesPath)

	// Call endpoint
	response, err := get(url)
	if err != nil {
		return newStories, err
	}

	// Unmarshal to output type
	err = json.Unmarshal(response, &newStories)
	if err != nil {
		return newStories, err
	}
	return newStories, nil
}

// GetBestStories -
func GetBestStories() (BestStories, error) {
	var bestStories BestStories

	// Build url
	url := buildRequestURL(bestStoriesPath)

	// Call endpoint
	response, err := get(url)
	if err != nil {
		return bestStories, err
	}

	// Unmarshal to output type
	err = json.Unmarshal(response, &bestStories)
	if err != nil {
		return bestStories, err
	}
	return bestStories, nil
}

// GetAskStories -
func GetAskStories() (AskStories, error) {
	var askStories AskStories
	// Build url
	url := buildRequestURL(askStoriesPath)

	// Call endpoint
	response, err := get(url)
	if err != nil {
		return askStories, err
	}

	// Unmarshal to output type
	err = json.Unmarshal(response, &askStories)
	if err != nil {
		return askStories, err
	}
	return askStories, nil
}

// GetShowStories -
func GetShowStories() (ShowStories, error) {
	var showStories ShowStories

	// Build url
	url := buildRequestURL(showStoriesPath)

	// Call endpoint
	response, err := get(url)
	if err != nil {
		return showStories, err
	}

	// Unmarshal to output type
	err = json.Unmarshal(response, &showStories)
	if err != nil {
		return showStories, err
	}
	return showStories, nil
}

// GetJobStories -
func GetJobStories() (JobStories, error) {
	var jobStories JobStories

	// Build url
	url := buildRequestURL(jobStoriesPath)

	// Call endpoint
	response, err := get(url)
	if err != nil {
		return jobStories, err
	}

	// Unmarshal to output type
	err = json.Unmarshal(response, &jobStories)
	if err != nil {
		return jobStories, err
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
