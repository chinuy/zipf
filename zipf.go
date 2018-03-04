package util

import (
	"math"
	"math/rand"
)

func NewZipf(r *rand.Rand, n int, alpha float64) *Zipf {
	z := &Zipf{}
	z.rng = r

	tmp := make([]float64, n)
	for i := 1; i < n+1; i++ {
		tmp[i-1] = 1.0 / math.Pow(float64(i), alpha)
	}

	zeta := make([]float64, n+1)
	z.distMap = make([]float64, n+1)
	z.prob = make([]float64, n)

	zeta[0] = 0
	for i := 1; i < n+1; i++ {
		zeta[i] += zeta[i-1] + tmp[i-1]
	}

	for i := 0; i < n+1; i++ {
		z.distMap[i] = zeta[i] / zeta[n]
	}

	for i := 1; i < n+1; i++ {
		z.prob[i-1] = z.distMap[i] - z.distMap[i-1]
	}

	s := 0.0
	for i := 0; i < n; i++ {
		s += z.prob[i]
	}

	return z
}

type Zipf struct {
	alpha float64
	distMap []float64
	prob []float64
	rng *rand.Rand
}

func (z *Zipf) next() int {
	u := z.rng.Float64()
	for i := 0; i < len(z.prob); i++ {
		if u < z.prob[i] {
			return i
		}
		u -= z.prob[i]
	}
	// should not come here
	// return last index for safty
	return len(z.prob)-1
}
