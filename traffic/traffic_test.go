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
