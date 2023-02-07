package table

var number int

type Table struct {
	Number   int
	Capacity int
	Status
}

func NewTable(capacity int) Table {
	number++
	return Table{
		Number:   number,
		Capacity: capacity,
	}
}

type Status int

const (
	Free Status = iota
	Occupied
	Reserved
)
