package decimal

import "testing"

func TestDecimal_ExpHullAbrham_SignificantDigitsBelowOne(t *testing.T) {
	for _, testCase := range []struct {
		name             string
		dec              string
		overallPrecision uint32
		expected         string
	}{
		{"[QP-QTGO-31-1] Returns 0.05 for exp(-3) at precision 1", "-3", 1, "0.05"},
		{"[QP-QTGO-31-2] Returns 0.0000000000046 for exp(-26.1) at precision 2", "-26.1", 2, "0.0000000000046"},
		{"[QP-QTGO-31-3] Returns three significant digits for exp(-1) at precision 3", "-1", 3, "0.368"},
	} {
		t.Run(testCase.name, func(t *testing.T) {
			d, _ := NewFromString(testCase.dec)
			expected, _ := NewFromString(testCase.expected)

			exp, err := d.ExpHullAbrham(testCase.overallPrecision)
			if err != nil {
				t.Fatal(err)
			}

			if exp.Cmp(expected) != 0 {
				t.Errorf("expected %s, got %s, for decimal %s", testCase.expected, exp.String(), testCase.dec)
			}
		})
	}
}
