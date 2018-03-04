package util

import (
	"testing"
	"math/rand"
)

func TestFixture(t *testing.T) {
	fixture := []int{8,0,3,0,1,0,0,3,7,0}
	r := rand.New(rand.NewSource(0))
	z := NewZipf(r, 10, 1.0)
	for i := 0; i < 10; i++ {
		if fixture[i] != z.next() {
			t.Errorf("fixture mismatch")
		}
	}
}
