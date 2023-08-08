package entity

type ApplyPlan struct {
	Id           int    `json:"id"`
	CompanyName  string `json:"companyName"`
	Position     string `json:"position"`
	PostedDate   string `json:"postedDate"`
	DeadlineDate string `json:"deadlineDate"`
	Source       string `json:"source"`
}

type AppData struct {
	Version int         `json:"version"`
	LastId  int         `json:"lastId"`
	Data    []ApplyPlan `json:"data"`
}
