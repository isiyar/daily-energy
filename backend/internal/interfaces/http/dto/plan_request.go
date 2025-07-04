package dto

type PlanRequest struct {
	Date int64 `json:"date" validate:"required"`
}
