package main

type Grid struct {
	Width, Height int
	Cells         [][]bool
}

// NewGrid crea una cuadrícula vacía
func NewGrid(width, height int) *Grid {
	cells := make([][]bool, height)
	for y := range cells {
		cells[y] = make([]bool, width)
	}
	return &Grid{Width: width, Height: height, Cells: cells}
}

// SetAlive marca una célula como viva
func (g *Grid) SetAlive(x, y int) {
	if x >= 0 && x < g.Width && y >= 0 && y < g.Height {
		g.Cells[y][x] = true
	}
}

// IsAlive verifica si una célula está viva
func (g *Grid) IsAlive(x, y int) bool {
	if x < 0 || x >= g.Width || y < 0 || y >= g.Height {
		return false
	}
	return g.Cells[y][x]
}
