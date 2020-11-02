package hackernews

// Updates response struct
type Updates struct {
	Items    []int    `json:"items"`
	Profiles []string `json:"profiles"`
}
