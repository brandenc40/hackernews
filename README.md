# HackerNews API

[![codecov](https://codecov.io/gh/brandenc40/hackernews/branch/master/graph/badge.svg?token=YDCD39G1C1)](undefined)

#### Easy to use gateway for the HackerNews API that leverages goroutines for highly efficient requests.

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

// Get the user who posted the top story
user, err := hackernews.GetUser(hydratedStories[0].By)
if err != nil {
    // handle error
}
```
