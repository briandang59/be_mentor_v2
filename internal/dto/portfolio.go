package dto

type PortfolioRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Url         string `json:"url" binding:"required,url"`
	JobTitle    string `json:"job_title"`
	TimePeriod  string `json:"time_period"`
	IsDraft     bool   `json:"is_draft"`
	ThumbnailID uint   `json:"thumbnail_id"`
}
