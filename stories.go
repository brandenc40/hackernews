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

// Path returned the hackernews API path
// for the given StoryType
func (s StoryType) Path() string {
	switch s {
	case StoriesTop:
		return "topstories"
	case StoriesNew:
		return "newstories"
	case StoriesBest:
		return "beststories"
	case StoriesAsk:
		return "askstories"
	case StoriesShow:
		return "showstories"
	case StoriesJob:
		return "jobstories"
	}
	return ""
}
