/*
Package rounding provides a variety of rounding routines.

See also:
* http://en.wikipedia.org/wiki/Rounding
* http://www.mathsisfun.com/numbers/rounding-methods.html
*/
package rounding

// RandomGenerator is an interface used for generating a new random floating point number.
// The standard math.Rand class implements this.
type RandomGenerator interface {
	Float64() float64
}

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
func NewDown() Rounder {
	return NewRounder(RoundDown)
}

// NewUp returns a new Up rounder.
func NewUp() Rounder {
	return NewRounder(RoundUp)
}

// NewHalfDown creates a new HalfDown Rounder
func NewHalfDown() Rounder {
	return NewRounder(RoundHalfDown)
}

// NewHalfUp creates a new HalfUp Rounder
func NewHalfUp() Rounder {
	return NewRounder(RoundHalfUp)
}

// NewHalfEven creates a new HalfEven Rounder
func NewHalfEven() Rounder {
	return NewRounder(RoundHalfEven)
}

// NewSymmetricDown creates a new SymmetricDown Rounder
func NewSymmetricDown() Rounder {
	return NewRounder(RoundSymmetricDown)
}

// NewSymmetricUp creates a new SymmetricUp Rounder
func NewSymmetricUp() Rounder {
	return NewRounder(RoundSymmetricUp)
}

// NewSymmetricHalfDown creates a new SymmetricHalfDown Rounder
func NewSymmetricHalfDown() Rounder {
	return NewRounder(RoundSymmetricHalfDown)
}

// NewSymmetricHalfUp creates a new SymmetricHalfUp Rounder
func NewSymmetricHalfUp() Rounder {
	return NewRounder(RoundSymmetricHalfUp)
}

// Creates a new alternateRounderImpl Rounder using Up and Down Rounders
func NewAlternate() Rounder {
	return &alternateRounderImpl{
		Up:        NewUp(),
		Down:      NewDown(),
		Direction: NewAlternatingRoundingDirection(),
	}
}

// Creates a new alternateRounderImpl Rounder using HalfUp and HalfDown Rounders
func NewHalfAlternate() Rounder {
	return &alternateRounderImpl{
		Up:        NewHalfUp(),
		Down:      NewHalfDown(),
		Direction: NewAlternatingRoundingDirection(),
	}
}

// Creates a new alternateRounderImpl Rounder using Symmetric Up and Down Rounders
func NewSymmetricAlternate() Rounder {
	return &alternateRounderImpl{
		Up:        NewSymmetricUp(),
		Down:      NewSymmetricDown(),
		Direction: NewAlternatingRoundingDirection(),
	}
}

// Creates a new alternateRounderImpl Rounder using Symmetric HalfUp and HalfDown Rounders
func NewSymmetricHalfAlternate() Rounder {
	return &alternateRounderImpl{
		Up:        NewSymmetricHalfUp(),
		Down:      NewSymmetricHalfDown(),
		Direction: NewAlternatingRoundingDirection(),
	}
}

// NewRandom returns a new RandomGenerator using the given midpoint and Up/Down rounders for rounding.
func NewRandom(mid float64, r RandomGenerator) Rounder {
	return &alternateRounderImpl{
		Up:        NewUp(),
		Down:      NewDown(),
		Direction: NewRandomRoundingDirection(mid, r),
	}
}

// NewHalfRandom returns a new RandomGenerator using the given midpoint and HalfUp/HalfDown rounders for rounding.
func NewHalfRandom(mid float64, r RandomGenerator) Rounder {
	return &alternateRounderImpl{
		Up:        NewHalfUp(),
		Down:      NewHalfDown(),
		Direction: NewRandomRoundingDirection(mid, r),
	}
}

// NewSymmetricRandom returns a new RandomGenerator using the given midpoint and SymmetricUp/SymmetricDown rounders for rounding.
func NewSymmetricRandom(mid float64, r RandomGenerator) Rounder {
	return &alternateRounderImpl{
		Up:        NewSymmetricUp(),
		Down:      NewSymmetricDown(),
		Direction: NewRandomRoundingDirection(mid, r),
	}
}

// NewSymmetricHalfRandom returns a new RandomGenerator using the given midpoint and SymmetricHalfUp/SymmetricHalfDown rounders for rounding.
func NewSymmetricHalfRandom(mid float64, r RandomGenerator) Rounder {
	return &alternateRounderImpl{
		Up:        NewSymmetricHalfUp(),
		Down:      NewSymmetricHalfDown(),
		Direction: NewRandomRoundingDirection(mid, r),
	}
}
