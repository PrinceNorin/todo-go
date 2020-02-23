package types

type Todo struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Completed bool   `json:"completed"`
}
