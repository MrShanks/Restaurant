package order

import (
	"fmt"

	"github.com/MrShanks/Restaurant/internal/item"
)

type Order struct {
	ID    int
	Table int
	Items []item.Item
	Notes string
	Bill  float64
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
