package rounding

import (
	"testing"
)

func runTests(t *testing.T, unit Rounder, tests [][]float64) {
	for _, test := range tests {
		rounded := unit.Round(test[0])
		if rounded != test[1] {
			t.Errorf("Rounding did not work on %f: %f != %f", test[0], test[1], rounded)
		}
	}
}

func TestRoundDown(t *testing.T) {
	runTests(t, NewDown(), [][]float64{
		{0.0, 0.0},
		{1.0, 1.0},
		{-1.0, -1.0},
		{0.9, 0.0},
		{-0.9, -1.0},
		{1.1, 1.0},
		{-1.1, -2.0},
		{0.1, 0.0},
		{0.2, 0.0},
		{0.3, 0.0},
		{0.4, 0.0},
		{0.5, 0.0},
		{0.6, 0.0},
		{0.7, 0.0},
		{0.8, 0.0},
		{0.9, 0.0},
		{-0.1, -1.0},
		{-0.2, -1.0},
		{-0.3, -1.0},
		{-0.4, -1.0},
		{-0.5, -1.0},
		{-0.6, -1.0},
		{-0.7, -1.0},
		{-0.8, -1.0},
		{-0.9, -1.0},
	})
}

func TestRoundSymmetricDown(t *testing.T) {
	runTests(t, NewSymmetricDown(), [][]float64{
		{0.0, 0.0},
		{1.0, 1.0},
		{-1.0, -1.0},
		{0.9, 0.0},
		{-0.9, 0.0},
		{1.1, 1.0},
		{-1.1, -1.0},
		{0.1, 0.0},
		{0.2, 0.0},
		{0.3, 0.0},
		{0.4, 0.0},
		{0.5, 0.0},
		{0.6, 0.0},
		{0.7, 0.0},
		{0.8, 0.0},
		{0.9, 0.0},
		{-0.1, -0.0},
		{-0.2, -0.0},
		{-0.3, -0.0},
		{-0.4, -0.0},
		{-0.5, -0.0},
		{-0.6, -0.0},
		{-0.7, -0.0},
		{-0.8, -0.0},
		{-0.9, -0.0},
	})
}

func TestRoundUp(t *testing.T) {
	runTests(t, NewUp(), [][]float64{
		{0.0, 0.0},
		{1.0, 1.0},
		{-1.0, -1.0},
		{0.9, 1.0},
		{-0.9, 0.0},
		{1.1, 2.0},
		{-1.1, -1.0},
		{0.1, 1.0},
		{0.2, 1.0},
		{0.3, 1.0},
		{0.4, 1.0},
		{0.5, 1.0},
		{0.6, 1.0},
		{0.7, 1.0},
		{0.8, 1.0},
		{0.9, 1.0},
		{-0.1, 0.0},
		{-0.2, 0.0},
		{-0.3, 0.0},
		{-0.4, 0.0},
		{-0.5, 0.0},
		{-0.6, 0.0},
		{-0.7, 0.0},
		{-0.8, 0.0},
		{-0.9, 0.0},
	})
}

func TestRoundSymmetricUp(t *testing.T) {
	runTests(t, NewSymmetricUp(), [][]float64{
		{0.0, 0.0},
		{1.0, 1.0},
		{-1.0, -1.0},
		{0.9, 1.0},
		{-0.9, -1.0},
		{1.1, 2.0},
		{-1.1, -2.0},
		{0.1, 1.0},
		{0.2, 1.0},
		{0.3, 1.0},
		{0.4, 1.0},
		{0.5, 1.0},
		{0.6, 1.0},
		{0.7, 1.0},
		{0.8, 1.0},
		{0.9, 1.0},
		{-0.1, -1.0},
		{-0.2, -1.0},
		{-0.3, -1.0},
		{-0.4, -1.0},
		{-0.5, -1.0},
		{-0.6, -1.0},
		{-0.7, -1.0},
		{-0.8, -1.0},
		{-0.9, -1.0},
	})
}

func TestRoundHalfUp(t *testing.T) {
	runTests(t, NewHalfUp(), [][]float64{
		{0.0, 0.0},
		{1.0, 1.0},
		{-1.0, -1.0},
		{0.9, 1.0},
		{-0.9, -1.0},
		{1.1, 1.0},
		{-1.1, -1.0},
		{0.1, 0.0},
		{0.2, 0.0},
		{0.3, 0.0},
		{0.4, 0.0},
		{0.5, 1.0},
		{0.6, 1.0},
		{0.7, 1.0},
		{0.8, 1.0},
		{0.9, 1.0},
		{-0.1, 0.0},
		{-0.2, 0.0},
		{-0.3, 0.0},
		{-0.4, 0.0},
		{-0.5, 0.0},
		{-0.6, -1.0},
		{-0.7, -1.0},
		{-0.8, -1.0},
		{-0.9, -1.0},
	})
}

func TestRoundSymmetricHalfUp(t *testing.T) {
	runTests(t, NewSymmetricHalfUp(), [][]float64{
		{0.0, 0.0},
		{1.0, 1.0},
		{-1.0, -1.0},
		{0.9, 1.0},
		{-0.9, -1.0},
		{1.1, 1.0},
		{-1.1, -1.0},
		{0.1, 0.0},
		{0.2, 0.0},
		{0.3, 0.0},
		{0.4, 0.0},
		{0.5, 1.0},
		{0.6, 1.0},
		{0.7, 1.0},
		{0.8, 1.0},
		{0.9, 1.0},
		{-0.1, 0.0},
		{-0.2, 0.0},
		{-0.3, 0.0},
		{-0.4, 0.0},
		{-0.5, -1.0},
		{-0.6, -1.0},
		{-0.7, -1.0},
		{-0.8, -1.0},
		{-0.9, -1.0},
	})
}

func TestRoundHalfDown(t *testing.T) {
	runTests(t, NewHalfDown(), [][]float64{
		{0.0, 0.0},
		{1.0, 1.0},
		{-1.0, -1.0},
		{0.9, 1.0},
		{-0.9, -1.0},
		{1.1, 1.0},
		{-1.1, -1.0},
		{0.1, 0.0},
		{0.2, 0.0},
		{0.3, 0.0},
		{0.4, 0.0},
		{0.5, 0.0},
		{0.6, 1.0},
		{0.7, 1.0},
		{0.8, 1.0},
		{0.9, 1.0},
		{-0.1, 0.0},
		{-0.2, 0.0},
		{-0.3, 0.0},
		{-0.4, 0.0},
		{-0.5, -1.0},
		{-0.6, -1.0},
		{-0.7, -1.0},
		{-0.8, -1.0},
		{-0.9, -1.0},
	})
}

func TestRoundSymmetricHalfDown(t *testing.T) {
	runTests(t, NewSymmetricHalfDown(), [][]float64{
		{0.0, 0.0},
		{1.0, 1.0},
		{-1.0, -1.0},
		{0.9, 1.0},
		{-0.9, -1.0},
		{1.1, 1.0},
		{-1.1, -1.0},
		{0.1, 0.0},
		{0.2, 0.0},
		{0.3, 0.0},
		{0.4, 0.0},
		{0.5, 0.0},
		{0.6, 1.0},
		{0.7, 1.0},
		{0.8, 1.0},
		{0.9, 1.0},
		{-0.1, 0.0},
		{-0.2, 0.0},
		{-0.3, 0.0},
		{-0.4, 0.0},
		{-0.5, 0.0},
		{-0.6, -1.0},
		{-0.7, -1.0},
		{-0.8, -1.0},
		{-0.9, -1.0},
	})
}

func TestRoundHalfEven(t *testing.T) {
	runTests(t, NewHalfEven(), [][]float64{
		{8.0, 8.0},
		{7.6, 8.0},
		{7.5, 8.0},
		{7.4, 7.0},
		{7.0, 7.0},
		{1.0, 1.0},
		{0.6, 1.0},
		{0.5, 0.0},
		{0.4, 0.0},
		{0.0, 0.0},
		{0.4, 0.0},
		{-0.5, 0.0},
		{-0.6, -1.0},
		{-1.0, -1.0},
		{-7.0, -7.0},
		{-7.4, -7.0},
		{-7.5, -8.0},
		{-7.6, -8.0},
		{-8.0, -8.0},
	})
}

func TestRoundAlternate(t *testing.T) {
	runTests(t, NewAlternate(), [][]float64{
		{0.0, 0.0},
		{0.0, 0.0},
		{1.0, 1.0},
		{1.0, 1.0},
		{-1.0, -1.0},
		{-1.0, -1.0},
		{0.9, 1.0},
		{0.9, 0.0},
		{-0.9, -0.0},
		{-0.9, -1.0},
		{1.1, 2.0},
		{1.1, 1.0},
		{-1.1, -1.0},
		{-1.1, -2.0},
	})
}

func TestRoundSymmetricAlternate(t *testing.T) {
	runTests(t, NewSymmetricAlternate(), [][]float64{
		{0.0, 0.0},
		{0.0, 0.0},
		{1.0, 1.0},
		{1.0, 1.0},
		{-1.0, -1.0},
		{-1.0, -1.0},
		{0.9, 1.0},
		{0.9, 0.0},
		{-0.9, -1.0},
		{-0.9, -0.0},
		{1.1, 2.0},
		{1.1, 1.0},
		{-1.1, -2.0},
		{-1.1, -1.0},
	})
}

func TestRoundHalfAlternate(t *testing.T) {
	runTests(t, NewHalfAlternate(), [][]float64{
		{0.0, 0.0},
	})
}

func TestRoundSymmetricHalfAlternate(t *testing.T) {
	runTests(t, NewSymmetricHalfAlternate(), [][]float64{
		{0.0, 0.0},
	})
}

//RoundHalfEven			 87.50% (7/8)

type MockRandomGenerator struct {
	Next float64
}

func (m MockRandomGenerator) Float64() float64 {
	return m.Next
}

func TestRoundRandom(t *testing.T) {
	m := MockRandomGenerator{}
	runTests(t, NewRandom(0.5, m), [][]float64{
		{0.0, 0.0},
	})

	m.Next = 0.6
	runTests(t, NewRandom(0.5, m), [][]float64{
		{0.0, 0.0},
	})
}

func TestRoundHalfRandom(t *testing.T) {
	runTests(t, NewHalfRandom(0.5, MockRandomGenerator{}), [][]float64{
		{0.0, 0.0},
	})
}

func TestRoundSymmetricRandom(t *testing.T) {
	runTests(t, NewSymmetricRandom(0.5, MockRandomGenerator{}), [][]float64{
		{0.0, 0.0},
	})
}

func TestRoundSymmetricHalfRandom(t *testing.T) {
	runTests(t, NewSymmetricHalfRandom(0.5, MockRandomGenerator{}), [][]float64{
		{0.0, 0.0},
	})
}
