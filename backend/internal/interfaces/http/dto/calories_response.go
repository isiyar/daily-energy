package dto

type CaloriesAPIResponse struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}

type CaloriesResponse struct {
	Calories *int `json:"calories"`
}
