package main

import (
	"fmt"
	"io"
	"os"

	"github.com/MrShanks/Restaurant/internal/order"
	"github.com/MrShanks/Restaurant/internal/restaurant"
	"github.com/MrShanks/Restaurant/internal/waiter"
)

func createMenu() order.Menu {
	menu := order.NewMenu()

	item1 := order.NewItem("Spring rolls", 11.0)
	item2 := order.NewItem("Steamed Gioza", 8.0)
	item3 := order.NewItem("Ravioli", 16.0)
	item4 := order.NewItem("Ramen", 26.0)
	item5 := order.NewItem("Pad Thai", 13.0)
	item6 := order.NewItem("Pork Steak", 32.0)
	item7 := order.NewItem("Beef Steak", 39.0)
	item8 := order.NewItem("Wagiou Steak", 120.0)
	item9 := order.NewItem("Tiramisu", 6.0)
	item10 := order.NewItem("Creme Brule", 5.0)
	item11 := order.NewItem("Cheese Cake", 7.0)
	item12 := order.NewItem("Scones", 3.0)

	menu.AddItem(item1)
	menu.AddItem(item2)
	menu.AddItem(item3)
	menu.AddItem(item4)
	menu.AddItem(item5)
	menu.AddItem(item6)
	menu.AddItem(item7)
	menu.AddItem(item8)
	menu.AddItem(item9)
	menu.AddItem(item10)
	menu.AddItem(item11)
	menu.AddItem(item12)

	return menu
}

func Hello(out io.Writer) {
	fmt.Fprint(out, "Welcome to our restaurant\n")
}

func main() {
	Hello(os.Stdout)

	morbido := restaurant.NewRestaurant("morbido")

	morbido.AddTable(10)
	morbido.AddTable(10)
	morbido.AddTable(8)
	morbido.AddTable(8)
	morbido.AddTable(8)
	morbido.AddTable(4)
	morbido.AddTable(2)
	morbido.AddTable(2)

	simone := waiter.NewWaiter("Simone")
	menu := createMenu()

	itemsTable1 := []order.Item{
		menu.GetItem(1),
		menu.GetItem(2),
		menu.GetItem(4),
		menu.GetItem(7),
		menu.GetItem(12),
	}
	itemsTable2 := []order.Item{
		menu.GetItem(1),
		menu.GetItem(2),
		menu.GetItem(3),
		menu.GetItem(4),
		menu.GetItem(5),
		menu.GetItem(6),
		menu.GetItem(7),
		menu.GetItem(8),
		menu.GetItem(9),
		menu.GetItem(10),
		menu.GetItem(11),
		menu.GetItem(12),
	}

	if err := morbido.Tables[0].ReserveTable(10, "simone"); err != nil {
		fmt.Println(err)
	}

	order1 := simone.NewOrder(morbido.Tables[0], itemsTable1)
	order2 := simone.NewOrder(morbido.Tables[1], itemsTable2)
	order1.PrettyPrint()
	order2.PrettyPrint()
	menu.PrettyPrint()
}
