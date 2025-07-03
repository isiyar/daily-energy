package dto

type PlanResponse struct {
	Id                string `json:"id"`
	Utgid             int    `json:"utgid"`
	Date              int64  `json:"date"`
	CaloriesToConsume int    `json:"calories_to_consume"`
	CaloriesToBurn    int    `json:"calories_to_burn"`
	Recommendation    string `json:"recommendation"`
}
