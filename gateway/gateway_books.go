package gateway

type Book struct {
	ID               int64  `json:"id"`
	Title            string `json:"title"`
	Author           string `json:"author"`
	Amount_available int64  `json:"amount_available"`
}

type Club_Member struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	LastName string `json:"last_name"`
}
