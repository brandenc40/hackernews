package hackernews

// User - GetUser() response struct
type User struct {
	ID        string `json:"id"`
	Delay     int    `json:"delay"`
	Created   int64  `json:"created"`
	Karma     int    `json:"karma"`
	About     string `json:"about"`
	Submitted []int  `json:"submitted"`
}
