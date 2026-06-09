package traffic

type Light string

const (
	Red    Light = "RED"
	Yellow Light = "YELLOW"
	Green  Light = "GREEN"
)

type phase int

const (
	nsGreen phase = iota
	nsYellow
	nsAllRed
	ewGreen
	ewYellow
	ewAllRed
	phaseCount
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
	config              Config
	currentPhase        phase
	ticksInCurrentPhase int
}

func NewController(cfg Config) *Controller {
	return &Controller{
		config:              cfg,
		currentPhase:        nsGreen,
		ticksInCurrentPhase: 0,
	}
}

func (c *Controller) Tick() {
	c.ticksInCurrentPhase++

	var limit int
	switch c.currentPhase {
	case nsGreen, ewGreen:
		limit = c.config.GreenTicks
	case nsYellow, ewYellow:
		limit = c.config.YellowTicks
	case nsAllRed, ewAllRed:
		limit = c.config.RedOverlap
	}

	if c.ticksInCurrentPhase >= limit {
		c.ticksInCurrentPhase = 0
		c.currentPhase = (c.currentPhase + 1) % phaseCount
	}
}

func (c *Controller) State() IntersectionState {
	switch c.currentPhase {
	case nsGreen:
		return IntersectionState{NS: Green, EW: Red}
	case nsYellow:
		return IntersectionState{NS: Yellow, EW: Red}
	case nsAllRed, ewAllRed:
		return IntersectionState{NS: Red, EW: Red}
	case ewGreen:
		return IntersectionState{NS: Red, EW: Green}
	case ewYellow:
		return IntersectionState{NS: Red, EW: Yellow}
	default:
		return IntersectionState{NS: Red, EW: Red}
	}
}
