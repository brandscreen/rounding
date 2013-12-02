package rounding

/*
Copyright (c) 2013 Brandscreen Pty Ltd

Permission is hereby granted, free of charge, to any person obtaining a copy of
this software and associated documentation files (the "Software"), to deal in
the Software without restriction, including without limitation the rights to
use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
the Software, and to permit persons to whom the Software is furnished to do so,
subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
*/

import (
	"testing"
)

type MockRandomGenerator struct {
	Next float64
}

func (m MockRandomGenerator) Float64() float64 {
	return m.Next
}

func NewMockRandomGenerator() *MockRandomGenerator {
	return &MockRandomGenerator{}
}

func runTests(t *testing.T, unit Rounder, tests [][]float64) {
	for _, test := range tests {
		rounded := unit.Round(test[0])
		if rounded != test[1] {
			t.Errorf("Rounding did not work on %f: %f != %f", test[0], test[1], rounded)
		}
	}
}

func TestRoundDown(t *testing.T) {
	runTests(t, NewDownRounder(), [][]float64{
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
	runTests(t, NewSymmetricDownRounder(), [][]float64{
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
	runTests(t, NewUpRounder(), [][]float64{
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
	runTests(t, NewSymmetricUpRounder(), [][]float64{
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
	runTests(t, NewHalfUpRounder(), [][]float64{
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
	runTests(t, NewSymmetricHalfUpRounder(), [][]float64{
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
	runTests(t, NewHalfDownRounder(), [][]float64{
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
	runTests(t, NewSymmetricHalfDownRounder(), [][]float64{
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
	runTests(t, NewHalfEvenRounder(), [][]float64{
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
	runTests(t, NewAlternateRounder(), [][]float64{
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
	runTests(t, NewSymmetricAlternateRounder(), [][]float64{
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
	runTests(t, NewHalfAlternateRounder(), [][]float64{
		{0.0, 0.0},
	})
}

func TestRoundSymmetricHalfAlternate(t *testing.T) {
	runTests(t, NewSymmetricHalfAlternateRounder(), [][]float64{
		{0.0, 0.0},
	})
}

func TestRoundRandom(t *testing.T) {
	m := NewMockRandomGenerator()
	runTests(t, NewRandomRounder(0.5, m.Float64), [][]float64{
		{0.0, 0.0},
	})

	m.Next = 0.6
	runTests(t, NewRandomRounder(0.5, m.Float64), [][]float64{
		{0.0, 0.0},
	})
}

func TestRoundHalfRandom(t *testing.T) {
	m := NewMockRandomGenerator()
	runTests(t, NewHalfRandomRounder(0.5, m.Float64), [][]float64{
		{0.0, 0.0},
	})
}

func TestRoundSymmetricRandom(t *testing.T) {
	m := NewMockRandomGenerator()
	runTests(t, NewSymmetricRandomRounder(0.5, m.Float64), [][]float64{
		{0.0, 0.0},
	})
}

func TestRoundSymmetricHalfRandom(t *testing.T) {
	m := NewMockRandomGenerator()
	runTests(t, NewSymmetricHalfRandomRounder(0.5, m.Float64), [][]float64{
		{0.0, 0.0},
	})
}
