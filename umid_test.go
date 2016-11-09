package umid

import (
	"reflect"
	"testing"
)

func TestNewAsBytes(t *testing.T) {
	z := make([]byte, 10)
	b := NewAsBytes()

	if reflect.DeepEqual(z, b) {
		t.Fail()
	}
}

func TestNewAsString(t *testing.T) {
	s := NewAsString()

	if s == "" {
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
