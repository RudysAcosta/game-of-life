package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type Universe struct {
	size              int
	currentGeneration [][]int
	nextGeneration    [][]int
}

func NewUniverse(size int) (*Universe, error) {
	if size <= 0 {
		return nil, errors.New("size must be greater than 0")
	}

	currentGeneration := make([][]int, size)
	nextGeneration := make([][]int, size)
	for i := range currentGeneration {
		currentGeneration[i] = make([]int, size)
		nextGeneration[i] = make([]int, size)
	}

	return &Universe{
		size:              size,
		currentGeneration: currentGeneration,
		nextGeneration:    nextGeneration,
	}, nil
}

func (u *Universe) Set(col, row, value int) error {
	if !u.isValidCoordinate(col, row) {
		return errors.New("index out of bounds")
	}
	if value < 0 || value > 1 {
		return errors.New("value must be 0 or 1")
	}
	u.currentGeneration[col][row] = value
	return nil
}

func (u *Universe) Get(col, row int) (int, error) {
	if !u.isValidCoordinate(col, row) {
		return -1, errors.New("index out of bounds")
	}
	return u.currentGeneration[col][row], nil
}

func (u *Universe) isValidCoordinate(col, row int) bool {
	return col >= 0 && col < u.size && row >= 0 && row < u.size
}

func (u *Universe) NextGeneration() {
	for i := range u.currentGeneration {
		for j := range u.currentGeneration[i] {
			alive := u.countAliveNeighbours(i, j)
			if u.currentGeneration[i][j] == 1 {
				u.nextGeneration[i][j] = boolToInt(alive == 2 || alive == 3)
			} else {
				u.nextGeneration[i][j] = boolToInt(alive == 3)
			}
		}
	}
	u.swapGenerations()
}

func (u *Universe) countAliveNeighbours(col, row int) int {
	offsets := []struct{ dx, dy int }{
		{-1, 0}, {-1, -1}, {0, -1}, {1, -1},
		{1, 0}, {1, 1}, {0, 1}, {-1, 1},
	}
	alive := 0
	for _, offset := range offsets {
		nx, ny := normalizar(col+offset.dx, u.size), normalizar(row+offset.dy, u.size)
		if u.currentGeneration[nx][ny] == 1 {
			alive++
		}
	}
	return alive
}

func (u *Universe) swapGenerations() {
	for i := range u.currentGeneration {
		copy(u.currentGeneration[i], u.nextGeneration[i])
	}
}

func (u *Universe) Alive() int {
	alive := 0
	for i := range u.currentGeneration {
		for j := range u.currentGeneration[i] {
			alive += u.currentGeneration[i][j]
		}
	}
	return alive
}

func sparkLife(u *Universe) {
	for i := range u.currentGeneration {
		for j := range u.currentGeneration[i] {
			u.currentGeneration[i][j] = rand.Intn(2)
		}
	}
}

func displayUniverse(u *Universe) {
	for i := range u.currentGeneration {
		for j := range u.currentGeneration[i] {
			if u.currentGeneration[i][j] == 0 {
				fmt.Print(" ")
			} else {
				fmt.Print("O")
			}
		}
		fmt.Println()
	}
}

func normalizar(value, limit int) int {
	return (value + limit) % limit
}

func boolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}

func main() {
	var size, generations int
	fmt.Scan(&size)

	generations = 15

	rand.Seed(time.Now().UnixNano())
	universe, err := NewUniverse(size)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	sparkLife(universe)

	for i := 1; i < generations; i++ {
		fmt.Printf("Generation #%d\n", i)
		fmt.Printf("Alive: %d\n", universe.Alive())
		displayUniverse(universe)
		universe.NextGeneration()
	}

}
