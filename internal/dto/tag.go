package dto

type TagRequest struct {
	TagName string `json:"tag_name" binding:"required"`
}
