package tax

import "testing"

type sample struct {
	revenue float64
	tax     float64
}

func TestTax(t *testing.T) {
	testcases := []sample{
		{revenue: 100000, tax: 0},
		{revenue: 200000, tax: (200000 - 150000) * 0.05},
		{revenue: 400000, tax: (300000-150000)*0.05 + (400000-300000)*0.10},
		{revenue: 600000, tax: (300000-150000)*0.05 + (500000-300000)*0.10 + (600000-500000)*0.15},
		{revenue: 800000, tax: (300000-150000)*0.05 + (500000-300000)*0.10 + (750000-500000)*0.15 + (800000-750000)*0.20},
		{revenue: 1500000, tax: (300000-150000)*0.05 + (500000-300000)*0.10 + (750000-500000)*0.15 + (1000000-750000)*0.20 + (1500000-1000000)*0.25},
		{revenue: 3000000, tax: (300000-150000)*0.05 + (500000-300000)*0.10 + (750000-500000)*0.15 + (1000000-750000)*0.20 + (2000000-1000000)*0.25 + (3000000-2000000)*0.30},
		{revenue: 6000000, tax: (300000-150000)*0.05 + (500000-300000)*0.10 + (750000-500000)*0.15 + (1000000-750000)*0.20 + (2000000-1000000)*0.25 + (5000000-2000000)*0.30 + (6000000-5000000)*0.35},
	}

	for _, test := range testcases {
		tax, err := CalculateThaiTax(test.revenue)
		if err != nil {
			t.Error("error", err.Error())
		}

		if tax != test.tax {
			t.Errorf("Error: revenue %f, expected tax %f, got %f", test.revenue, test.tax, tax)
		}
	}

	_, err := CalculateThaiTax(-1)
	if err == nil {
		t.Error("negative revenue supposed to return err")
	}

}
