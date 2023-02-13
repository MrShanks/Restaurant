package menu

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/MrShanks/Restaurant/internal/item"
)

type Menu struct {
	Items map[int]item.Item
}

var menuItemID int

func NewMenu() Menu {
	items := make(map[int]item.Item)
	return Menu{Items: items}
}

func LoadMenu() Menu {
	f, err := os.Open("./internal/menu/morbido-menu.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	csvLines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	menu := NewMenu()

	for index, line := range csvLines {
		if index == 0 {
			continue
		}
		id, err := strconv.Atoi(line[0])
		if err != nil {
			log.Fatal(err)
		}
		name := line[1]
		price, err := strconv.ParseFloat(line[2], 32)
		if err != nil {
			log.Fatal(err)
		}
		if err != nil {
			log.Fatal(err)
		}
		item := item.Item{
			ID:    id,
			Name:  name,
			Price: price,
		}
		menu.AddItem(item)
	}
	return menu
}

func (m *Menu) AddItem(item item.Item) {
	menuItemID++
	m.Items[menuItemID] = item
}

func (m *Menu) GetItem(id int) item.Item {
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
