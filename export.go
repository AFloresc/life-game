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

	// HTML básico con estilo en línea
	fmt.Fprintln(f, "<!DOCTYPE html><html><head><meta charset='UTF-8'><title>Juego de la Vida</title>")
	fmt.Fprintln(f, "<style>body{font-family:monospace;} .cell{width:10px;height:10px;display:inline-block;} .alive{background:black;} .dead{background:white;}</style></head><body>")

	for i, frame := range frames {
		fmt.Fprintf(f, "<h3>Generación %d</h3><div>", i)
		for y := 0; y < frame.Height; y++ {
			for x := 0; x < frame.Width; x++ {
				class := "dead"
				if frame.IsAlive(x, y) {
					class = "alive"
				}
				fmt.Fprintf(f, "<div class='cell %s'></div>", class)
			}
			fmt.Fprintln(f, "<br>")
		}
		fmt.Fprintln(f, "</div><hr>")
	}

	fmt.Fprintln(f, "</body></html>")
	return nil
}
