package restaurant

import (
	"fmt"
	"testing"
)

func TestTable_ReserveTable(t *testing.T) {
	tests := []struct {
		Description string
		Guests      int
		Want        Status
		WantError   error
	}{
		{"Free table is reserved when reserve table method is called", 10, 2, nil},
		{"Free table is reserved when reserve table method is called", 12, 0, fmt.Errorf("Table excited capacity")},
	}
	for _, test := range tests {
		t.Run(test.Description, func(t *testing.T) {
			table := &Table{
				Number:   1,
				Capacity: 10,
			}
			err := table.ReserveTable(test.Guests, "Simone")

			fmt.Println(err)
			fmt.Println(test.WantError)

			if err != test.WantError {
				t.Errorf("Expected an error, didn't get one")
			}
			if table.Status != test.Want {
				t.Errorf("Expected table status is: %v, got: %v", test.Want, table.Status)
			}
		})
	}
}
