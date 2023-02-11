package restaurant

type Restaurant struct {
	Name   string
	Tables []Table
}

var number int

func NewRestaurant(name string) Restaurant {
	return Restaurant{
		Name: name,
	}
}

func (r *Restaurant) AddTable(capacity int) {
	number++
	r.Tables = append(r.Tables, Table{
		Number:   number,
		Capacity: capacity,
	})
}
