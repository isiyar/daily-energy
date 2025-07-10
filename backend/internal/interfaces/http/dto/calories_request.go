package dto

type CaloriesRequest struct {
	Title string `json:"title" validate:"required"`
}
