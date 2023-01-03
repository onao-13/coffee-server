package coffee

type Coffee struct {
	Id    int64   `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func New(id int64, name string, price float64) Coffee {
	return Coffee{id, name, price}
}
