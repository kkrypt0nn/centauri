package flags

import "testing"

func TestFlagsCompute(t *testing.T) {
	expect := 1<<13 | 1<<37
	bitfield := Compute(1<<13, 1<<37)
	if expect != bitfield {
		t.Errorf("Failed computing flags (got: %d, expected: %d)", bitfield, expect)
		return
	}
	t.Log("Successfully computed the flags")
}

func TestFlagsAdd(t *testing.T) {
	expect := 1<<13 | 1<<37
	bitfield := Compute(1 << 13)
	bitfield = Add(bitfield, 1<<37)
	if expect != bitfield {
		t.Errorf("Failed adding the flag (got: %d, expected: %d)", bitfield, expect)
		return
	}
	t.Log("Successfully added the flag")
}

func TestFlagsRemove(t *testing.T) {
	expect := 1<<13 | 1<<37
	bitfield := Compute(1<<13, 1<<37, 1<<50)
	bitfield = Remove(bitfield, 1<<50)
	if expect != bitfield {
		t.Errorf("Failed removing the flag (got: %d, expected: %d)", bitfield, expect)
		return
	}
	t.Log("Successfully removed the flag")
}

func TestFlagsToggle(t *testing.T) {
	expect := 1 << 13
	bitfield := Compute(1<<13, 1<<37)
	bitfield = Toggle(bitfield, 1<<37)
	if expect != bitfield {
		t.Errorf("Failed toggling the flag (got: %d, expected: %d)", bitfield, expect)
		return
	}
	t.Log("Successfully toggled the flag")
}

func TestFlagsHas(t *testing.T) {
	bitfield := Compute(1<<13, 1<<37)
	if Has(bitfield, 1<<37) && Has(bitfield, 1<<13, 1<<37) && !Has(bitfield, 1<<8) && !Has(bitfield, 1<<8, 1<<13) {
		t.Log("Successfully checked if the flags are contained in the bitfield")
		return
	}
	t.Errorf("Failed checking if the flags are contained in the bitfield (bitfield: %d)", bitfield)
}

func TestFlagsHasAny(t *testing.T) {
	bitfield := Compute(1<<13, 1<<37)
	if HasAny(bitfield, 1<<37) && HasAny(bitfield, 1<<8, 1<<37) && !HasAny(bitfield, 1<<8) && !HasAny(bitfield, 1<<7, 1<<8) {
		t.Log("Successfully checked if any of the flag is contained in the bitfield")
		return
	}
	t.Errorf("Failed checking if any of the flag is contained in the bitfield (bitfield: %d)", bitfield)
}

func TestFlagsHasNot(t *testing.T) {
	bitfield := Compute(1<<13, 1<<37)
	if HasNot(bitfield, 1<<8) && HasNot(bitfield, 1<<7, 1<<8) && !HasNot(bitfield, 1<<8, 1<<13) {
		t.Log("Successfully checked if the flags are not contained in the bitfield")
		return
	}
	t.Errorf("Failed checking if the flags are not contained in the bitfield (bitfield: %d)", bitfield)
}

func TestFlagsHasNotAny(t *testing.T) {
	bitfield := Compute(1<<13, 1<<37)
	if HasNotAny(bitfield, 1<<8) && HasNotAny(bitfield, 1<<8, 1<<13) && HasNotAny(bitfield, 1<<13, 1<<37, 1<<8) && !HasNotAny(bitfield, 1<<13, 1<<37) {
		t.Log("Successfully checked if any of the flag is not contained in the bitfield")
		return
	}
	t.Errorf("Failed checking if any of the flag is not contained in the bitfield (bitfield: %d)", bitfield)
}
