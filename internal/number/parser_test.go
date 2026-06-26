package number

import "testing"

func TestParseInteger(t *testing.T) {
	tests := map[string]int64{
		"0":    0,
		"-17":  -17,
		"0x2A": 42,
	}

	for input, expected := range tests {
		actual, ok := ParseInteger(input)
		if !ok {
			t.Fatalf("ParseInteger(%q) failed", input)
		}
		if actual != expected {
			t.Fatalf("ParseInteger(%q) = %d, want %d", input, actual, expected)
		}
	}

	if _, ok := ParseInteger("12.5"); ok {
		t.Fatal("ParseInteger(\"12.5\") succeeded, want failure")
	}
}

func TestParseFloat(t *testing.T) {
	tests := map[string]float64{
		"3.5":   3.5,
		"1e2":   100,
		"0x1p2": 4,
	}

	for input, expected := range tests {
		actual, ok := ParseFloat(input)
		if !ok {
			t.Fatalf("ParseFloat(%q) failed", input)
		}
		if actual != expected {
			t.Fatalf("ParseFloat(%q) = %v, want %v", input, actual, expected)
		}
	}

	if _, ok := ParseFloat("not-a-number"); ok {
		t.Fatal("ParseFloat(\"not-a-number\") succeeded, want failure")
	}
}
