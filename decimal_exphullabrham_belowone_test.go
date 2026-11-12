package decimal

import "testing"

func TestDecimal_ExpHullAbrham_BelowOneSignificantFigures(t *testing.T) {
	t.Run("[QP-QTGO-31-4] TestDecimal_ExpHullAbrham passes with significant-figure expectations", func(t *testing.T) {
		for _, testCase := range []struct {
			Dec              string
			OverallPrecision uint32
			ExpectedDec      string
		}{
			{"-1.0", 1, "0.4"},
			{"-1.0", 5, "0.36788"},
			{"-0.5", 5, "0.60653"},
			{"-0.569297", 10, "0.5659231429"},
			{"-0.569297", 16, "0.565923142859576"},
			{"-0.569297", 20, "0.56592314285957604443"},
			{"-3.0", 1, "0.05"},
			{"-3.0", 2, "0.05"},
			{"-3.0", 10, "0.04978706837"},
			{"-5.26", 2, "0.0052"},
			{"-5.26", 10, "0.005195304719"},
			{"-5.2663117716", 2, "0.0052"},
			{"-5.2663117716", 10, "0.005162616411"},
			{"-26.1", 2, "0.0000000000046"},
			{"-26.1", 15, "0.00000000000462289492466867"},
			{"-50.1591", 10, "0.0000000000000000000001645052084"},
			{"-50.1591", 30, "0.000000000000000000000164505208424152908216198963722"},
		} {
			d, _ := NewFromString(testCase.Dec)
			expected, _ := NewFromString(testCase.ExpectedDec)

			exp, err := d.ExpHullAbrham(testCase.OverallPrecision)
			if err != nil {
				t.Fatal(err)
			}

			if exp.Cmp(expected) != 0 {
				t.Errorf("expected %s, got %s, for decimal %s at precision %d", testCase.ExpectedDec, exp.String(), testCase.Dec, testCase.OverallPrecision)
			}
		}
	})
}
