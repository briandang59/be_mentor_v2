package dto

type LanguageRequest struct {
	Name   string `json:"name" binding:"required"`
	Object string `json:"object"`
	From   string `json:"from"`
	To     string `json:"to"`
}
