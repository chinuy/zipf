package zipf

import (
	"testing"
	"math/rand"
	"math"
)

const EPLISON = 0.001

func TestStatistic(t *testing.T) {
	alpha := 1.0
	n := 10
	run := 1000000
	m := make(map[int]int, run)
	prob := make(map[int]float64, run)

	s := 0.0
	for i := 0; i < n; i++ {
		s += 1 / math.Pow(float64(i+1), alpha)
	}

	for i := 0; i < n; i++ {
		prob[i] = 1/math.Pow(float64(i+1), alpha)/s
	}

	r := rand.New(rand.NewSource(0))
	z := NewZipf(r, alpha, uint64(n))
	for i := 0; i < run; i++ {
		m[int(z.Uint64())]++
	}

	for i := 0; i < n; i++ {
		if math.Abs(float64(m[i])/float64(run) - prob[i]) > EPLISON {
			t.Errorf("probability mismatch")
		}
	}

}

func TestFixture(t *testing.T) {
	fixture := []uint64{8,0,3,0,1,0,0,3,7,0}
	r := rand.New(rand.NewSource(0))
	z := NewZipf(r, 1.0, 10)
	for i := 0; i < 10; i++ {
		if fixture[i] != z.Uint64() {
			t.Errorf("fixture mismatch")
		}
	}
}
