package dto

type ContactInformationRequest struct {
	Phone    string `json:"phone" binding:"omitempty,max=20"`
	LinkedIn string `json:"linkedin" binding:"omitempty,url"`
	Facebook string `json:"facebook" binding:"omitempty,url"`
	Zalo     string `json:"zalo" binding:"omitempty"`
	Telegram string `json:"telegram" binding:"omitempty"`
	X        string `json:"x" binding:"omitempty,url"`
	GitHub   string `json:"github" binding:"omitempty,url"`
	Twitter  string `json:"twitter" binding:"omitempty,url"`
	Website  string `json:"website" binding:"omitempty,url"`
	Location string `json:"location" binding:"omitempty,max=255"`
}
