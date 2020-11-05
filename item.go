package hackernews

// ItemType describes what type of object
// the Item struct represents
type ItemType string

const (
	// ItemTypeJob is a job item
	ItemTypeJob ItemType = "job"
	// ItemTypeStory is a story item
	ItemTypeStory ItemType = "story"
	// ItemTypeComment is a comment item
	ItemTypeComment ItemType = "comment"
	// ItemTypePoll is a poll item
	ItemTypePoll ItemType = "poll"
	// ItemTypePollOpt is a poll option item
	ItemTypePollOpt ItemType = "pollopt"
)

// Item - GetItem() response struct. Items can
// be any of the ItemType options.
type Item struct {
	ID          int      `json:"id"`
	Deleted     string   `json:"deleted"`
	Type        ItemType `json:"type"`
	By          string   `json:"by"`
	Time        int64    `json:"time"`
	Text        string   `json:"text"`
	Dead        bool     `json:"dead"`
	Parent      int      `json:"parent"`
	Poll        int      `json:"poll"`
	Kids        []int    `json:"kids"`
	URL         string   `json:"url"`
	Score       int      `json:"score"`
	Title       string   `json:"title"`
	Parts       []int    `json:"parts"`
	Descendants int      `json:"descendants"`
}

// MaxItem response struct. A single item ID.
type MaxItem int
