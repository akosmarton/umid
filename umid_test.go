package umid

import (
	"reflect"
	"testing"
)

func TestNewAsBytes(t *testing.T) {
	b := NewAsBytes()

	if len(b) != 10 {
		t.Fail()
	}

	if reflect.DeepEqual(b, make([]byte, 10)) {
		t.Fail()
	}
}

func TestNewAsString(t *testing.T) {
	s := NewAsString()

	if len(s) != 20 {
		t.Fail()
	}

	if s == "00000000000000000000" {
		t.Fail()
	}
}

func BenchmarkNewAsBytes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewAsBytes()
	}
}

func BenchmarkNewAsString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewAsString()
	}
}
