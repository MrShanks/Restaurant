package order

import "fmt"

type Menu struct {
	Items map[int]Item
}

var menuItemID int

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

func (m *Menu) PrettyPrint() {
	fmt.Println("*--------------------------------------*")
	fmt.Println("|                 MENU                 |")
	fmt.Println("*--------------------------------------*")
	for i := 1; i <= len(m.Items); i++ {
		item := m.GetItem(i)
		fmt.Printf("| %-3v| %-20v   %-4v %v |\n", item.ID, item.Name, item.Price, "CHF")
	}
	fmt.Println("*--------------------------------------*")
}
