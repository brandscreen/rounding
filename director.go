package rounding

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
	return r.Random.Float64() < r.Mid
}

// NewRandomRoundingDirection returns a RoundingDirection using the given midpoint and RandomGenerator
func NewRandomRoundingDirection(mid float64, random RandomGenerator) RoundingDirection {
	return &RandomRoundingDirection{
		Mid:    mid,
		Random: random,
	}
}
