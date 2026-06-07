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

type Controller struct{}

func NewController(cfg Config) *Controller {
	return nil
}

func (c *Controller) Tick() {}

func (c *Controller) State() IntersectionState {
	return IntersectionState{}
}
