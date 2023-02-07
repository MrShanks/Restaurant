package order

import "fmt"

type Order struct {
	ID    int
	Table int
	Items []Item
	Notes string
	Bill  float64
}

type Menu struct {
	Items map[int]Item
}

type Item struct {
	ID    int
	Name  string
	Price float64
}

var itemID int
var menuItemID int

func NewItem(name string, price float64) Item {
	itemID++
	return Item{
		ID:    itemID,
		Name:  name,
		Price: price,
	}
}

func NewMenu() Menu {
	items := make(map[int]Item)
	return Menu{Items: items}
}

func (m *Menu) AddItem(item Item) {
	menuItemID++
	m.Items[menuItemID] = item
}

func (m *Menu) GetItem(id int) Item {
	return m.Items[id]
}

func (o *Order) PrettyPrint() {
	fmt.Println("*-----------------------*")
	fmt.Printf("| Order: %-15v|\n", o.ID)
	fmt.Printf("*-----------------------*\n")
	fmt.Printf("| Items:                |\n")
	for _, item := range o.Items {
		fmt.Printf("|   %-20v|\n", item.Name)
	}
	fmt.Printf("*-----------------------*\n")
	fmt.Printf("| Cost: %-16v|\n", o.Bill)
	fmt.Println("*-----------------------*")
}

func (m *Menu) PrettyPrint() {
	fmt.Println("*---------------------------------*")
	fmt.Println("|              MENU               |")
	fmt.Println("*---------------------------------*")
	for i := 1; i <= len(m.Items); i++ {
		item := m.GetItem(i)
		fmt.Printf("| %-3v  %-20v   %-4v|\n", item.ID, item.Name, item.Price)
	}
	fmt.Println("*---------------------------------*")
}
