package maze

import (
	"testing"
)

func TestGenerationSizeOfResult(t *testing.T) {

	m := new(Maze)
	h := 10
	w := 10
	grid := m.GenerateMaze(h, w)

	if len(grid) != h {
		t.Fatal("height is not correct")
	}

	for i := 0; i < len(grid); i++ {
		if len(grid[i]) != w {
			t.Fatal("some width of grid not correct")
		}
	}

}

func TestConsoleOutput(t *testing.T) {

	m := new(Maze)
	h := 10
	w := 10
	grid := m.GenerateMaze(h, w)

	m.PrintConsole(grid)

}
