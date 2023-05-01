package car24_car_service

type Model struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type CreateModel struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type UpdateModel struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type DeleteModel struct {
	ID string `json:"id" swaggerignore:"true"`
}

type ModelQueryParam struct {
	Search string `json:"search"`
	Offset int    `json:"offset" default:"0"`
	Limit  int    `json:"limit" default:"10"`
}

type ModelList struct {
	Models []Model `json:"brands"`
	Count  int     `json:"count"`
}
