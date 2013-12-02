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

// RoundingDirection is an interface used to determine whether an Up or Down rounding is used.
type RoundingDirection interface {
	// IsUp determines if Up (true) or Down (false) direction is used.
	IsUp() bool
}

// AlternatingRoundingDirection is a RoundingDirection that toggles between up and down.
type AlternatingRoundingDirection struct {
	toggle bool
}

// IsUp determines if Up (true) or Down (false) direction is used.
func (r *AlternatingRoundingDirection) IsUp() bool {
	r.toggle = !r.toggle
	return r.toggle
}

// NewAlternatingRoundingDirection returns a new RoundingDirection that alternates between Up and Down
func NewAlternatingRoundingDirection() RoundingDirection {
	return &AlternatingRoundingDirection{}
}

// RandomRoundingDirection is a RoundingDirection that returns a random Up or Down signal
// based on a RandomGenerator and a Midpoint.
type RandomRoundingDirection struct {
	Mid    float64
	Random RandomGenerator
}

// IsUp determines if Up (true) or Down (false) direction is used.
func (r *RandomRoundingDirection) IsUp() bool {
	return r.Random() < r.Mid
}

// NewRandomRoundingDirection returns a RoundingDirection using the given midpoint and RandomGenerator
func NewRandomRoundingDirection(mid float64, random RandomGenerator) RoundingDirection {
	return &RandomRoundingDirection{
		Mid:    mid,
		Random: random,
	}
}
