package posts

type Filter struct {
	Limit  int    `json:"limit"`
	Offset int    `json:"offset"`
	Status string `json:"status"`
}
