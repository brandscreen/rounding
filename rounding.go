/*
Package rounding provides a variety of rounding routines.

See also:
* http://en.wikipedia.org/wiki/Rounding
* http://www.mathsisfun.com/numbers/rounding-methods.html
*/
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

// RandomGenerator is used for generating a new random floating point number.
type RandomGenerator func() float64

// Rounder is an interface implemented by rounding algorithms
type Rounder interface {
	// Rounds the given value using the rounding algorithm
	Round(float64) float64
}

// rounderImpl is an implementation of Rounder that wraps a rounding function.
type rounderImpl struct {
	round RoundingAlgorithm
}

// Rounds the given value using the rounding algorithm
func (r *rounderImpl) Round(value float64) float64 {
	return r.round(value)
}

// alternateRounderImpl is a rounding algorithm that alternates between an Up and a Down rounding.  It has no bias for sequential calls
type alternateRounderImpl struct {
	Up        Rounder
	Down      Rounder
	Direction RoundingDirection
}

// Rounds the given value using the rounding algorithm
func (r *alternateRounderImpl) Round(value float64) float64 {
	if r.Direction.IsUp() {
		return r.Up.Round(value)
	}

	return r.Down.Round(value)
}

// NewRounder returns a Rounder that wraps the given rounding function.
func NewRounder(rounder RoundingAlgorithm) Rounder {
	return &rounderImpl{
		round: rounder,
	}
}

// NewDown returns a new Down rounder.
func NewDownRounder() Rounder {
	return NewRounder(RoundDown)
}

// NewUp returns a new Up rounder.
func NewUpRounder() Rounder {
	return NewRounder(RoundUp)
}

// NewHalfDown creates a new HalfDown Rounder
func NewHalfDownRounder() Rounder {
	return NewRounder(RoundHalfDown)
}

// NewHalfUp creates a new HalfUp Rounder
func NewHalfUpRounder() Rounder {
	return NewRounder(RoundHalfUp)
}

// NewHalfEven creates a new HalfEven Rounder
func NewHalfEvenRounder() Rounder {
	return NewRounder(RoundHalfEven)
}

// NewSymmetricDown creates a new SymmetricDown Rounder
func NewSymmetricDownRounder() Rounder {
	return NewRounder(RoundSymmetricDown)
}

// NewSymmetricUp creates a new SymmetricUp Rounder
func NewSymmetricUpRounder() Rounder {
	return NewRounder(RoundSymmetricUp)
}

// NewSymmetricHalfDown creates a new SymmetricHalfDown Rounder
func NewSymmetricHalfDownRounder() Rounder {
	return NewRounder(RoundSymmetricHalfDown)
}

// NewSymmetricHalfUp creates a new SymmetricHalfUp Rounder
func NewSymmetricHalfUpRounder() Rounder {
	return NewRounder(RoundSymmetricHalfUp)
}

// Creates a new alternateRounderImpl Rounder using Up and Down Rounders
func NewAlternateRounder() Rounder {
	return &alternateRounderImpl{
		Up:        NewUpRounder(),
		Down:      NewDownRounder(),
		Direction: NewAlternatingRoundingDirection(),
	}
}

// Creates a new alternateRounderImpl Rounder using HalfUp and HalfDown Rounders
func NewHalfAlternateRounder() Rounder {
	return &alternateRounderImpl{
		Up:        NewHalfUpRounder(),
		Down:      NewHalfDownRounder(),
		Direction: NewAlternatingRoundingDirection(),
	}
}

// Creates a new alternateRounderImpl Rounder using Symmetric Up and Down Rounders
func NewSymmetricAlternateRounder() Rounder {
	return &alternateRounderImpl{
		Up:        NewSymmetricUpRounder(),
		Down:      NewSymmetricDownRounder(),
		Direction: NewAlternatingRoundingDirection(),
	}
}

// Creates a new alternateRounderImpl Rounder using Symmetric HalfUp and HalfDown Rounders
func NewSymmetricHalfAlternateRounder() Rounder {
	return &alternateRounderImpl{
		Up:        NewSymmetricHalfUpRounder(),
		Down:      NewSymmetricHalfDownRounder(),
		Direction: NewAlternatingRoundingDirection(),
	}
}

// NewRandom returns a new RandomGenerator using the given midpoint and Up/Down rounders for rounding.
func NewRandomRounder(mid float64, r RandomGenerator) Rounder {
	return &alternateRounderImpl{
		Up:        NewUpRounder(),
		Down:      NewDownRounder(),
		Direction: NewRandomRoundingDirection(mid, r),
	}
}

// NewHalfRandom returns a new RandomGenerator using the given midpoint and HalfUp/HalfDown rounders for rounding.
func NewHalfRandomRounder(mid float64, r RandomGenerator) Rounder {
	return &alternateRounderImpl{
		Up:        NewHalfUpRounder(),
		Down:      NewHalfDownRounder(),
		Direction: NewRandomRoundingDirection(mid, r),
	}
}

// NewSymmetricRandom returns a new RandomGenerator using the given midpoint and SymmetricUp/SymmetricDown rounders for rounding.
func NewSymmetricRandomRounder(mid float64, r RandomGenerator) Rounder {
	return &alternateRounderImpl{
		Up:        NewSymmetricUpRounder(),
		Down:      NewSymmetricDownRounder(),
		Direction: NewRandomRoundingDirection(mid, r),
	}
}

// NewSymmetricHalfRandom returns a new RandomGenerator using the given midpoint and SymmetricHalfUp/SymmetricHalfDown rounders for rounding.
func NewSymmetricHalfRandomRounder(mid float64, r RandomGenerator) Rounder {
	return &alternateRounderImpl{
		Up:        NewSymmetricHalfUpRounder(),
		Down:      NewSymmetricHalfDownRounder(),
		Direction: NewRandomRoundingDirection(mid, r),
	}
}
