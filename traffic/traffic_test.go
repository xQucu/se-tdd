package traffic

import "testing"

func TestInitialState(t *testing.T) {
	cfg := Config{
		GreenTicks:  5,
		YellowTicks: 2,
		RedOverlap:  1,
	}
	c := NewController(cfg)
	if c == nil {
		t.Fatal("expected controller to be non-nil")
	}
	state := c.State()
	if state.NS != Green {
		t.Errorf("expected initial NS Green, got %s", state.NS)
	}
	if state.EW != Red {
		t.Errorf("expected initial EW Red, got %s", state.EW)
	}
}

func TestNSGreenToNSYellow(t *testing.T) {
	cfg := Config{
		GreenTicks:  5,
		YellowTicks: 2,
		RedOverlap:  1,
	}
	c := NewController(cfg)

	// Ticks 0 to 4 should be NS Green, EW Red
	for i := 0; i < 5; i++ {
		state := c.State()
		if state.NS != Green || state.EW != Red {
			t.Fatalf("Tick %d: expected NS Green / EW Red, got NS %s / EW %s", i, state.NS, state.EW)
		}
		c.Tick()
	}

	// At tick 5, NS should turn Yellow, EW remains Red
	state := c.State()
	if state.NS != Yellow || state.EW != Red {
		t.Errorf("expected NS Yellow / EW Red, got NS %s / EW %s", state.NS, state.EW)
	}
}

func TestNSYellowToAllRed(t *testing.T) {
	cfg := Config{
		GreenTicks:  5,
		YellowTicks: 2,
		RedOverlap:  1,
	}
	c := NewController(cfg)

	// Tick through Green phase (5 ticks)
	for i := 0; i < 5; i++ {
		c.Tick()
	}
	// Tick through Yellow phase (2 ticks)
	for i := 0; i < 2; i++ {
		state := c.State()
		if state.NS != Yellow || state.EW != Red {
			t.Fatalf("Tick %d: expected NS Yellow / EW Red, got NS %s / EW %s", 5+i, state.NS, state.EW)
		}
		c.Tick()
	}

	// At tick 7, NS should turn Red (both Red)
	state := c.State()
	if state.NS != Red || state.EW != Red {
		t.Errorf("expected both Red, got NS %s / EW %s", state.NS, state.EW)
	}
}

func TestAllRedToEWGreen(t *testing.T) {
	cfg := Config{
		GreenTicks:  5,
		YellowTicks: 2,
		RedOverlap:  1,
	}
	c := NewController(cfg)

	// Tick 5 + 2 + 1 = 8 times
	for i := 0; i < 8; i++ {
		c.Tick()
	}

	// At tick 8, EW should be Green, NS should be Red
	state := c.State()
	if state.NS != Red || state.EW != Green {
		t.Errorf("expected NS Red / EW Green, got NS %s / EW %s", state.NS, state.EW)
	}
}



