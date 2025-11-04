package main

import (
	"fmt"
	"os"
)

func ExportToHTML(frames []*Grid, filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	// HTML y estilos
	fmt.Fprintln(f, "<!DOCTYPE html><html><head><meta charset='UTF-8'><title>Juego de la Vida</title>")
	fmt.Fprintln(f, `<style>
  body { font-family: sans-serif; text-align: center; background: white; }
  #grid { display: grid; margin: auto; }
  .cell { width: 10px; height: 10px; }
  .alive { background: black; }
  .dead { background: white; }
</style></head><body>`)

	fmt.Fprintln(f, "<h2>Juego de la Vida</h2><div id='grid'></div>")
	fmt.Fprintln(f, "<script>")
	fmt.Fprintln(f, "const frames = [")

	// Serialización de frames sin coma extra
	for i, frame := range frames {
		fmt.Fprint(f, "[")
		for y := 0; y < frame.Height; y++ {
			fmt.Fprint(f, "[")
			for x := 0; x < frame.Width; x++ {
				val := 0
				if frame.IsAlive(x, y) {
					val = 1
				}
				fmt.Fprintf(f, "%d", val)
				if x < frame.Width-1 {
					fmt.Fprint(f, ",")
				}
			}
			fmt.Fprint(f, "]")
			if y < frame.Height-1 {
				fmt.Fprint(f, ",")
			}
		}
		fmt.Fprint(f, "]")
		if i < len(frames)-1 {
			fmt.Fprintln(f, ",")
		} else {
			fmt.Fprintln(f)
		}
	}
	fmt.Fprintln(f, "];")

	// JavaScript para animación
	fmt.Fprintf(f, `
const width = %d;
const height = %d;
const grid = document.getElementById("grid");
grid.style.gridTemplateColumns = "repeat(" + width + ", 10px)";

for (let i = 0; i < width * height; i++) {
  const cell = document.createElement("div");
  cell.classList.add("cell");
  grid.appendChild(cell);
}

let generation = 0;
function render(frame) {
  const cells = grid.children;
  for (let y = 0; y < height; y++) {
    for (let x = 0; x < width; x++) {
      const idx = y * width + x;
      cells[idx].className = "cell " + (frame[y][x] ? "alive" : "dead");
    }
  }
}

function next() {
  render(frames[generation]);
  generation = (generation + 1) %% frames.length;
}
setInterval(next, 300);
</script></body></html>
`, frames[0].Width, frames[0].Height)

	return nil
}
