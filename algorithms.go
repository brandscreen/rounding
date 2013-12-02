package rounding

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
