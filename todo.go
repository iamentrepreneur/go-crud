package go_crud

type TodoList struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type UserList struct {
	Id     int `json:"id"`
	USerId int `json:"user_id"`
	ListId int `json:"list_id"`
}

type TodoItem struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

type ListsItem struct {
	Id     int `json:"id"`
	ListId int `json:"list_id"`
	ItemId int `json:"item_id"`
}
