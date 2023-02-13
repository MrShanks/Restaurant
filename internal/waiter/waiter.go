package waiter

import (
	"github.com/MrShanks/Restaurant/internal/item"
	"github.com/MrShanks/Restaurant/internal/order"
	"github.com/MrShanks/Restaurant/internal/restaurant"
)

type Waiter struct {
	Name   string
	Orders []order.Order
}

var orderID int

func NewWaiter(name string) *Waiter {
	return &Waiter{
		Name:   name,
		Orders: nil,
	}
}

func (w *Waiter) NewOrder(table restaurant.Table, items []item.Item) order.Order {
	orderID++
	bill := 0.0
	for _, item := range items {
		bill += item.Price
	}
	newOrder := order.Order{
		ID:    orderID,
		Table: table.Number,
		Items: items, Bill: bill,
	}
	w.Orders = append(w.Orders, newOrder)
	return newOrder
}
