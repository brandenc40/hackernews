package hackernews

// Updates - GetUpdates() response struct
type Updates struct {
	Items    []int    `json:"items"`
	Profiles []string `json:"profiles"`
}
