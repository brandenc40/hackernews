package hackernews

// Stories - GetStories() response struct. An array of item ids.
type Stories []int

// StoryType - Enum value for types of stories available
type StoryType int

// Enum for types of stories
const (
	StoriesTop StoryType = iota
	StoriesNew
	StoriesBest
	StoriesAsk
	StoriesShow
	StoriesJob
)

func (s StoryType) String() string {
	switch s {
	case StoriesTop:
		return "top"
	case StoriesNew:
		return "new"
	case StoriesBest:
		return "best"
	case StoriesAsk:
		return "ask"
	case StoriesShow:
		return "show"
	case StoriesJob:
		return "job"
	}
	return ""
}

// Path returned the hackernews API path
// for the given StoryType
func (s StoryType) Path() string {
	return s.String() + "stories"
}
