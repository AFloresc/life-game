package main

func CountNeighbors(g *Grid, x, y int) int {
	count := 0
	for dy := -1; dy <= 1; dy++ {
		for dx := -1; dx <= 1; dx++ {
			if dx == 0 && dy == 0 {
				continue
			}
			if g.IsAlive(x+dx, y+dy) {
				count++
			}
		}
	}
	return count
}

func NextGeneration(current *Grid) *Grid {
	next := NewGrid(current.Width, current.Height)
	for y := 0; y < current.Height; y++ {
		for x := 0; x < current.Width; x++ {
			neighbors := CountNeighbors(current, x, y)
			if current.IsAlive(x, y) {
				if neighbors == 2 || neighbors == 3 {
					next.SetAlive(x, y)
				}
			} else {
				if neighbors == 3 {
					next.SetAlive(x, y)
				}
			}
		}
	}
	return next
}
