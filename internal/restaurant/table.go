package restaurant

import "fmt"

type Table struct {
	Number          int
	Capacity        int
	ReservationName string
	Status
}

type Status int

const (
	Free Status = iota
	Occupied
	Reserved
)

func (t *Table) ReserveTable(guests int, name string) error {
	if guests > t.Capacity {
		return fmt.Errorf("Table excited capacity")
	}
	t.Status = Reserved
	t.ReservationName = name
	return nil
}
