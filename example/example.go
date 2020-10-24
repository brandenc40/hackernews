package example

import (
	"fmt"

	"github.com/brandenc40/hackernews"
)

func main() {
	// Grab the item ids of all top stories
	topStories, err := hackernews.GetStories(hackernews.StoriesTop)
	if err != nil {
		panic(err)
	}

	// Hydrate that list of item ids into a list of Item structs
	hydratedStories, err := hackernews.HydrateItems(topStories)
	if err != nil {
		panic(err)
	}

	// Get the user posted the top story
	user, err := hackernews.GetUser(hydratedStories[0].By)
	if err != nil {
		panic(err)
	}

	fmt.Println(user)
}
