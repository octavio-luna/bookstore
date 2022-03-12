package methods

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

func NewBook(title string, author string, amount_available int64) Book {
	return Book{
		Title:            title,
		Author:           author,
		Amount_available: amount_available,
	}
}

func NewClub_Member(name string, last_name string) Club_Member {
	return Club_Member{
		Name:     name,
		LastName: last_name,
	}
}
