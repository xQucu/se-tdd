package main

import (
	"fmt"
	"time"

	"tdd/traffic"
)

// colorLight formats the light string with nice ANSI background colors
func colorLight(l traffic.Light) string {
	switch l {
	case traffic.Green:
		return "\033[42m\033[30m  GREEN   \033[0m"
	case traffic.Yellow:
		return "\033[43m\033[30m  YELLOW  \033[0m"
	case traffic.Red:
		return "\033[41m\033[37m   RED    \033[0m"
	default:
		return string(l)
	}
}

func main() {
	cfg := traffic.Config{
		GreenTicks:  5,
		YellowTicks: 2,
		RedOverlap:  1,
	}

	c := traffic.NewController(cfg)

	var tick int
	for {
		state := c.State()

		// Clear screen and reset cursor to home (0,0)
		fmt.Print("\033[H\033[2J")

		fmt.Println("==================================================")
		fmt.Println("    TRAFFIC LIGHT SIMULATION (Ctrl+C to stop)")
		fmt.Println("==================================================")
		fmt.Printf("  North-South Lights:  [%s]\n", colorLight(state.NS))
		fmt.Printf("  East-West Lights:    [%s]\n", colorLight(state.EW))
		fmt.Println("==================================================")
		fmt.Printf("  Total Ticks Elapsed: %d\n", tick)
		fmt.Println("==================================================")

		c.Tick()
		tick++

		time.Sleep(800 * time.Millisecond)
	}
}
