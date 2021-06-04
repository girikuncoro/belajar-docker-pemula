package schema

type Todo struct {
	ID   int    `json:"id"`
	Note string `json:"note"`
	Done bool   `json:"done"`
}
