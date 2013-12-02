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
	"math"
)

// RoundingAlgorithm is a function signature used for rounding.
type RoundingAlgorithm func(float64) float64

// RoundDown rounds the given value down using math.Floor
func RoundDown(value float64) float64 {
	return math.Floor(value)
}

// RoundUp rounds the given value up using math.Ceil
func RoundUp(value float64) float64 {
	return math.Ceil(value)
}

// RoundHalfUp rounds towards +Infinity
func RoundHalfUp(value float64) float64 {
	return math.Floor(value + 0.5)
}

// RoundHalfDown rounds towards -Infinity
func RoundHalfDown(value float64) float64 {
	return math.Ceil(value - 0.5)
}

// RoundHalfEven is often known as Banker's Rounding and generally has no bias.
func RoundHalfEven(value float64) float64 {
	if value < 0.0 {
		return -RoundHalfEven(-value)
	}

	integralPart, fractionalPart := math.Modf(value)

	// If 'value' is exactly halfway between two integers
	if fractionalPart == 0.5 {
		// If 'integralPart' is even then return 'integralPart'
		if math.Remainder(integralPart, 2.0) == 0 {
			return integralPart
		}

		// Else return the nearest even integer
		return RoundSymmetricUp(integralPart + 0.5)
	}

	// Otherwise use the usual round to closest
	// (Either symmetric half-up or half-down will do)
	return RoundSymmetricHalfUp(value)
}

// Symmetric performs a symmetric rounding using a specified RoundingAlgorithm
func Symmetric(rounder RoundingAlgorithm, value float64) float64 {
	result := rounder(math.Abs(value))
	if value < 0.0 {
		return -result
	} else {
		return result
	}
}

// RoundSymmetricDown rounds symmetrically down towards zero
func RoundSymmetricDown(value float64) float64 {
	return Symmetric(RoundDown, value)
}

// RoundSymmetricUp rounds symmetrically up away from zero
func RoundSymmetricUp(value float64) float64 {
	return Symmetric(RoundUp, value)
}

// RoundSymmetricHalfDown rounds symmetrically down towards zero
func RoundSymmetricHalfDown(value float64) float64 {
	return Symmetric(RoundHalfDown, value)
}

// RoundSymmetricHalfUp rounds symmetrically up away from zero
func RoundSymmetricHalfUp(value float64) float64 {
	return Symmetric(RoundHalfUp, value)
}
