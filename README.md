# HackerNews API

[![PkgGoDev](https://pkg.go.dev/badge/github.com/brandenc40/hackernews)](https://pkg.go.dev/github.com/brandenc40/hackernews)
[![codecov](https://codecov.io/gh/brandenc40/hackernews/branch/master/graph/badge.svg?token=YDCD39G1C1)](undefined)
![Tests](https://github.com/brandenc40/hackernews/workflows/Tests/badge.svg)
![CodeQL](https://github.com/brandenc40/hackernews/workflows/CodeQL/badge.svg)
![golangci-lint](https://github.com/brandenc40/hackernews/workflows/golangci-lint/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/brandenc40/hackernews)](https://goreportcard.com/report/github.com/brandenc40/hackernews)

### Easy to use client for the HackerNews API that leverages goroutines for highly efficient requests.

##### Full HackerNews API documentation can be found here: https://github.com/HackerNews/API

##### Available Functions

```go
GetPaginatedStories(storyType StoryType, storiesPerPage int, pageNumber int) (PaginatedStoriesResponse, error)
GetStories(storyType StoryType) (Stories, error)
GetItem(itemID int) (Item, error)
GetUser(userID string) (User, error)
GetMaxItem() (MaxItem, error)
GetUpates() (Updates, error)
HydrateItems(itemIDs []int) ([]Item, error)
```

##### Example

```go
import "github.com/brandenc40/hackernews"

// Grab the item ids of all top stories
topStories, err := hackernews.GetStories(hackernews.StoriesTop)
if err != nil {
    // handle error
}

// Hydrate that list of item ids into a list of Item structs
hydratedStories, err := hackernews.HydrateItems(topStories)
if err != nil {
    // handle error
}

topStory := hydratedStories[0]

// Get the user who posted the top story
user, err := hackernews.GetUser(topStory.By)
if err != nil {
    // handle error
}
```
