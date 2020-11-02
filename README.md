# HackerNews API

[![codecov](https://codecov.io/gh/brandenc40/hackernews/branch/master/graph/badge.svg?token=YDCD39G1C1)](undefined)

### Easy to use client for the HackerNews API that leverages goroutines for highly efficient requests.

##### Full API documentation can be found here: https://github.com/HackerNews/API

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
