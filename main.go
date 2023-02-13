package main

import (
	"fmt"
	"io"
	"os"

	"github.com/MrShanks/Restaurant/internal/item"
	"github.com/MrShanks/Restaurant/internal/menu"
	"github.com/MrShanks/Restaurant/internal/restaurant"
	"github.com/MrShanks/Restaurant/internal/waiter"
)

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
	var menuMorbido = menu.LoadMenu()
	menuMorbido.PrettyPrint()

	itemsTable1 := []item.Item{
		menuMorbido.GetItem(1),
		menuMorbido.GetItem(2),
		menuMorbido.GetItem(4),
		menuMorbido.GetItem(7),
		menuMorbido.GetItem(12),
	}
	itemsTable2 := []item.Item{
		menuMorbido.GetItem(1),
		menuMorbido.GetItem(2),
		menuMorbido.GetItem(3),
		menuMorbido.GetItem(4),
		menuMorbido.GetItem(5),
		menuMorbido.GetItem(6),
		menuMorbido.GetItem(7),
		menuMorbido.GetItem(8),
		menuMorbido.GetItem(9),
		menuMorbido.GetItem(10),
		menuMorbido.GetItem(11),
		menuMorbido.GetItem(12),
	}

	if err := morbido.Tables[0].ReserveTable(10, "simone"); err != nil {
		fmt.Println(err)
	}

	order1 := simone.NewOrder(morbido.Tables[0], itemsTable1)
	order2 := simone.NewOrder(morbido.Tables[1], itemsTable2)
	order1.PrettyPrint()
	order2.PrettyPrint()

}
