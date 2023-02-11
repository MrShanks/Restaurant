package restaurant

import "testing"

func TestTable_ReserveTable(t *testing.T) {
	tests := []struct {
		Description string
		Want        Status
	}{
		{"Free table is reserved when reserve table method is called", 2},
	}
	for _, test := range tests {
		t.Run(test.Description, func(t *testing.T) {
			table := &Table{
				Number:   1,
				Capacity: 10,
			}
			table.ReserveTable(10, "Simone")
			if table.Status != test.Want {
				t.Errorf("Expected table status is: %v, got: %v", test.Want, table.Status)
			}
		})
	}
}
