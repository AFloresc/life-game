package main

import (
	"fmt"
	"time"
)

func PrintGrid(g *Grid) {
	for y := 0; y < g.Height; y++ {
		for x := 0; x < g.Width; x++ {
			if g.IsAlive(x, y) {
				fmt.Print("⬛")
			} else {
				fmt.Print("⬜")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func main() {
	grid := NewGrid(10, 10)

	// Glider
	grid.SetAlive(1, 0)
	grid.SetAlive(2, 1)
	grid.SetAlive(0, 2)
	grid.SetAlive(1, 2)
	grid.SetAlive(2, 2)

	var frames []*Grid

	for i := 0; i < 20; i++ {
		PrintGrid(grid)
		grid = NextGeneration(grid)
		time.Sleep(300 * time.Millisecond)
	}

	err := ExportToHTML(frames, "vida.html")
	if err != nil {
		fmt.Println("Error al exportar:", err)
	} else {
		fmt.Println("Exportado a vida.html")
	}
}
