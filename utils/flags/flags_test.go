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
	if Has(bitfield, 1<<37) && Has(bitfield, 1<<13, 1<<37) {
		t.Log("Successfully checked if the flag is contained in the bitfield")
		return
	}
	t.Errorf("Failed checking if the flag is contained in the bitfield (bitfield: %d)", bitfield)
}

func TestFlagsHasNot(t *testing.T) {
	bitfield := Compute(1<<13, 1<<37)
	if HasNot(bitfield, 1<<8) && HasNot(bitfield, 1<<13, 1<<8) {
		t.Log("Successfully checked if the flag is not contained in the bitfield")
		return
	}
	t.Errorf("Failed checking if the flag is not contained in the bitfield (bitfield: %d)", bitfield)
}
