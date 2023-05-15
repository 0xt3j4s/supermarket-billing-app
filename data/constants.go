package data

type Item struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Quantity  int    `json:"quantity"`
	Price     int    `json:"price"`
	AddedAt string `json:"added_at"`
}

type Bill struct {
	ID        int64    `json:"id"`
	UserName      string `json:"user_name"`
	Items     []Item `json:"items"`
	CreatedAt string `json:"created_at"`
}