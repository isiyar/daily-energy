package dto

type UserWeightHistoryRequest struct {
	Date   int64 `json:"date" validate:"required"`
	Weight int   `json:"weight" validate:"required,min=10,max=300"`
	Height int   `json:"height" validate:"required,min=10,max=300"`
}
