package traffic

type Light string

const (
	Red    Light = "RED"
	Yellow Light = "YELLOW"
	Green  Light = "GREEN"
)

type IntersectionState struct {
	NS Light
	EW Light
}

type Config struct {
	GreenTicks  int
	YellowTicks int
	RedOverlap  int
}

type Controller struct {
	config Config
	ticks  int
}

func NewController(cfg Config) *Controller {
	return &Controller{
		config: cfg,
		ticks:  0,
	}
}

func (c *Controller) Tick() {
	c.ticks++
}

func (c *Controller) State() IntersectionState {
	if c.ticks >= c.config.GreenTicks+c.config.YellowTicks {
		return IntersectionState{
			NS: Red,
			EW: Red,
		}
	}
	if c.ticks >= c.config.GreenTicks {
		return IntersectionState{
			NS: Yellow,
			EW: Red,
		}
	}
	return IntersectionState{
		NS: Green,
		EW: Red,
	}
}


