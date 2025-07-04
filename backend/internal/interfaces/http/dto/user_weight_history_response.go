package dto

type UserWeightHistoryResponse struct {
	Utgid  int64 `json:"utgid"`
	Date   int64 `json:"date" `
	Weight int   `json:"weight" validate:"required,min=10,max=300"`
	Height int   `json:"height" validate:"required,min=10,max=300"`
}
