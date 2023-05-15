package data

type Bill struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Quantity  int    `json:"quantity"`
	Price     int    `json:"price"`
	CreatedAt string `json:"created_at"`
}