package dto

type PostCreateRequest struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
	TagIDs  []uint `json:"tag_ids"`
}
