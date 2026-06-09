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

func TestFullTransitions(t *testing.T) {
	cfg := Config{
		GreenTicks:  5,
		YellowTicks: 2,
		RedOverlap:  1,
	}
	c := NewController(cfg)

	// Step 1: NS Green / EW Red (ticks 0 to 4)
	for i := 0; i < cfg.GreenTicks; i++ {
		state := c.State()
		if state.NS != Green || state.EW != Red {
			t.Fatalf("Tick %d: expected NS Green / EW Red, got NS %s / EW %s", i, state.NS, state.EW)
		}
		c.Tick()
	}

	// Step 2: NS Yellow / EW Red (ticks 5 to 6)
	for i := 0; i < cfg.YellowTicks; i++ {
		state := c.State()
		if state.NS != Yellow || state.EW != Red {
			t.Fatalf("Tick %d: expected NS Yellow / EW Red, got NS %s / EW %s", cfg.GreenTicks+i, state.NS, state.EW)
		}
		c.Tick()
	}

	// Step 3: NS Red / EW Red (overlap) (tick 7)
	for i := 0; i < cfg.RedOverlap; i++ {
		state := c.State()
		if state.NS != Red || state.EW != Red {
			t.Fatalf("Tick %d: expected NS Red / EW Red, got NS %s / EW %s", cfg.GreenTicks+cfg.YellowTicks+i, state.NS, state.EW)
		}
		c.Tick()
	}

	// Step 4: NS Red / EW Green (ticks 8 to 12)
	for i := 0; i < cfg.GreenTicks; i++ {
		state := c.State()
		if state.NS != Red || state.EW != Green {
			t.Fatalf("Tick %d: expected NS Red / EW Green, got NS %s / EW %s", cfg.GreenTicks+cfg.YellowTicks+cfg.RedOverlap+i, state.NS, state.EW)
		}
		c.Tick()
	}

	// Step 5: NS Red / EW Yellow (ticks 13 to 14)
	for i := 0; i < cfg.YellowTicks; i++ {
		state := c.State()
		if state.NS != Red || state.EW != Yellow {
			t.Fatalf("Tick %d: expected NS Red / EW Yellow, got NS %s / EW %s", cfg.GreenTicks+cfg.YellowTicks+cfg.RedOverlap+cfg.GreenTicks+i, state.NS, state.EW)
		}
		c.Tick()
	}

	// Step 6: NS Red / EW Red (overlap) (tick 15)
	for i := 0; i < cfg.RedOverlap; i++ {
		state := c.State()
		if state.NS != Red || state.EW != Red {
			t.Fatalf("Tick %d: expected NS Red / EW Red, got NS %s / EW %s", cfg.GreenTicks+cfg.YellowTicks+cfg.RedOverlap+cfg.GreenTicks+cfg.YellowTicks+i, state.NS, state.EW)
		}
		c.Tick()
	}

	// Step 7: Back to NS Green / EW Red (tick 16)
	state := c.State()
	if state.NS != Green || state.EW != Red {
		t.Fatalf("expected cycle to restart with NS Green / EW Red, got NS %s / EW %s", state.NS, state.EW)
	}
}




