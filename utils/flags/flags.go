// Package flags is a package for helper functions related to bitwise flags
package flags

// Integer represents a collection of possible integers for generic functions
type Integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

// Compute creates a new bitfield for the given flags
func Compute[T Integer](flags ...T) T {
	var bitfield T
	for _, flag := range flags {
		bitfield |= flag
	}
	return bitfield
}

// Add adds the given flags to the bitfield
func Add[T Integer](bitfield T, flags ...T) T {
	for _, flag := range flags {
		bitfield |= flag
	}
	return bitfield
}

// Remove removes the given flags from the bitfield
func Remove[T Integer](bitfield T, flags ...T) T {
	var toRemove T
	for _, flag := range flags {
		toRemove |= flag
	}
	return bitfield &^ toRemove
}

// Toggle toggles the given flags from the bitfield
func Toggle[T Integer](bitfield T, flags ...T) T {
	for _, flag := range flags {
		bitfield ^= flag
	}
	return bitfield
}

// Has checks if all the given flags are in the bitfield
func Has[T Integer](bitfield T, flags ...T) bool {
	for _, flag := range flags {
		if bitfield&flag != flag {
			return false
		}
	}
	return true
}

// HasAny checks if any of the given flags is in the bitfield
func HasAny[T Integer](bitfield T, flags ...T) bool {
	for _, flag := range flags {
		if bitfield&flag == flag {
			return true
		}
	}
	return false
}

// HasNot checks if all the given flags are not in the bitfield
func HasNot[T Integer](bitfield T, flags ...T) bool {
	return !HasAny(bitfield, flags...)
}

// HasNotAny checks if any of the given flags is not in the bitfield
func HasNotAny[T Integer](bitfield T, flags ...T) bool {
	return !Has(bitfield, flags...)
}
