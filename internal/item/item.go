package item

type Item struct {
	ID    int
	Name  string
	Price float64
}

var itemID int

func NewItem(name string, price float64) Item {
	itemID++
	return Item{
		ID:    itemID,
		Name:  name,
		Price: price,
	}
}
