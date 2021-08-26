package improve

import (
	"strings"
	"testing"
)
var data = `jack
rose
lucy
tom
mickey`

var expected = 5
func TestCountLines1(t *testing.T) {
	got, err := CountLines1(strings.NewReader(data))
	if err != nil || got != expected {
		t.Errorf("expected: 5 , but got %v, and error: %v", got ,err )
	}
}

func TestCountLines2(t *testing.T) {
	got, err := CountLines2(strings.NewReader(data))
	if err != nil || got != expected {
		t.Errorf("expected: 5 , but got %v, and error: %v", got ,err )
	}

}
