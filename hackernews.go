package hackernews

import (
	"encoding/json"
	"strconv"

	"golang.org/x/sync/errgroup"
)

const (
	// HackerNews API paths. Story paths
	// can be found in the stories.go file.
	itemPath    = "item"
	maxItemPath = "maxitem"
	userPath    = "user"
	updatesPath = "updates"
)

// HydrateItems concurrently hydrates a list of item ids into
// Item structs. GetStores and children of items are returned as
// purely a list of item ids. With this list we need to pass each item
// id into GetItem in order to get the full details. This HydrateItems
// method will fetch that list of item ids concurrently to greatly
// improve execution time.
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

// GetItem - Get a single item and return as an Item struct. Items can be
// on of "job", "story", "comment", "poll", or "pollopt". They're identified
// by their ids, which are unique integers.
//
// API DOC: https://github.com/HackerNews/API#items
//
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

// GetUser - Get a single user and return as a User struct.
// Users are identified by case-sensitive ids.
//
// API DOC: https://github.com/HackerNews/API#users
//
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

// GetStories - Get a list of all story ids of the type passed through
// the storyType argument.
//
// StoryType can be one of the following:
//   - StoriesTop
//   - StoriesNew
//   - StoriesBest
//   - StoriesAsk
//   - StoriesShow
//   - StoriesJob
//
// API DOCS:
//   - https://github.com/HackerNews/API#new-top-and-best-stories
//   - https://github.com/HackerNews/API#ask-show-and-job-stories
//
func GetStories(storyType StoryType) (Stories, error) {
	var stories Stories

	// Build url
	url := buildRequestURL(storyType.Path())

	// Call endpoint
	response, err := get(url)
	if err != nil {
		return stories, err
	}

	// Unmarshal to output type
	err = json.Unmarshal(response, &stories)
	if err != nil {
		return stories, err
	}
	return stories, nil
}

// GetMaxItem - The current largest item id.
//
// API DOC: https://github.com/HackerNews/API#max-item-id
//
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

// GetUpates - Item and profile changes.
//
// API DOC: https://github.com/HackerNews/API#changed-items-and-profiles
//
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
