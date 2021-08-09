package basic

import "testing"

func TestTypeAssertion(t *testing.T) {

	var expected = []struct {
		in  interface{}
		out string
	}{
		{123, "int"},
		{"hello", "string"},
		{[]byte{}, "unknown"},
	}

	for _, val := range expected {
		got := typeAssertion(val.in)
		if val.out != got {
			t.Errorf("expected: %v, got: %v", val.out, got)
		}
	}
}
