package dto

type ExperienceRequest struct {
	Title       string `json:"title" binding:"required"`
	Company     string `json:"company"`
	Location    string `json:"location"`
	From        string `json:"from"`
	To          string `json:"to"`
	Description string `json:"description"`
	IsCurrent   bool   `json:"is_current"`
}
