package web

type PostCreateRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}
