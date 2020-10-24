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

topStory := hydratedStories[0]

// Get the user who posted the top story
user, err := hackernews.GetUser(topStory.By)
if err != nil {
    // handle error
}

// Use the Type arg to know what kind of item it is.
// In this example we are using GetItem to grab a Poll.
if topStory.Type == "poll" {
    poll, err := hackernews.GetItem(topStory.Poll)
    if err != nil {
        // handle error
    }
}
```
