package main

import (
	"errors"
	"fmt"
	"math/rand"
)

type Universe struct {
	size        int
	data        [][]int
	newGenerate [][]int
}

func (u *Universe) Set(col, row, value int) error {
	if col < 0 || col >= u.size || row < 0 || row >= u.size {
		return errors.New("index out of bounds")
	}

	if value < 0 || value > 1 {
		return errors.New("value must be 0 or 1")
	}

	u.data[col][row] = value
	return nil
}

func (u *Universe) Get(col, row int) (int, error) {
	if col < 0 || col >= u.size || row < 0 || row >= u.size {
		return -1, errors.New("index out of bounds")
	}

	return u.data[col][row], nil
}

func (u *Universe) NextGeneration() {
	newGenerate := make([][]int, u.size)
	for i := range newGenerate {
		newGenerate[i] = make([]int, u.size)
	}

	for i := range u.data {
		for j := range u.data[i] {
			cell := NewCell(i, j, u.size)
			neighbours := cell.GetNeighbours()
			alive := 0
			for _, neighbour := range neighbours {
				if u.data[neighbour.x][neighbour.y] == 1 {
					alive++
					if alive > 3 {
						break
					}
				}
			}

			if u.data[i][j] == 1 {
				if alive < 2 || alive > 3 {
					newGenerate[i][j] = 0
				}
			} else {
				if alive == 3 {
					newGenerate[i][j] = 1
				}
			}
		}
	}

	for i := range u.data {
		copy(u.data[i], newGenerate[i])
	}
}

type Cell struct {
	x, y  int
	limit int
}

func (c *Cell) GetNeighbours() []Coordinate {

	offsets := []struct{ dx, dy int }{
		{-1, 0},  // Izquierda
		{-1, -1}, // Izquierda-arriba
		{0, -1},  // Arriba
		{1, -1},  // Derecha-arriba
		{1, 0},   // Derecha
		{1, 1},   // Derecha-abajo
		{0, 1},   // Abajo
		{-1, 1},  // Izquierda-abajo
	}

	// Slice para almacenar los vecinos
	neighbours := make([]Coordinate, 0, len(offsets))

	// Calcular los vecinos
	for _, offset := range offsets {
		neighbour := NewCoordinate(c.limit)
		neighbour.Set(c.x+offset.dx, c.y+offset.dy)
		neighbours = append(neighbours, *neighbour)
	}

	return neighbours
}

type Coordinate struct {
	x, y  int
	limit int
}

func NewCoordinate(limit int) *Coordinate {
	return &Coordinate{limit: limit}
}

func (c *Coordinate) Set(x, y int) {
	c.x = normalizar(x, c.limit)
	c.y = normalizar(y, c.limit)
}

func normalizar(value, limit int) int {
	if value < 0 {
		value = limit - 1
	} else if value == limit {
		value = 0
	}

	return value
}

func NewCell(x, y, limit int) *Cell {
	return &Cell{x: x, y: y, limit: limit}
}

func main() {
	var size, numbeSeed, generation int
	fmt.Scan(&size, &numbeSeed, &generation)

	rand.Seed(int64(numbeSeed))

	universe, err := newUniverse(size)
	if err != nil {
		fmt.Println(err)
	}

	sparkLife(universe)

	// universe.generated = append(universe.generated, *universe)
	universe.NextGeneration()
	// universe.NextGeneration()
	// universe.NextGeneration()

	fmt.Println(universe.data)

	displayUniverse(universe)

}

func newUniverse(size int) (*Universe, error) {
	if size <= 0 {
		return nil, errors.New("size must be greater than 0")
	}

	data := make([][]int, size)
	newGenerate := make([][]int, size)

	for i := range data {
		data[i] = make([]int, size)
		newGenerate[i] = make([]int, size)
	}

	return &Universe{size: size, data: data, newGenerate: newGenerate}, nil
}

func sparkLife(universe *Universe) error {
	for i := range universe.data {
		for j := range universe.data[i] {
			life := rand.Intn(2)
			if err := universe.Set(i, j, life); err != nil {
				fmt.Printf("Error in col %d, row %d", i, j)
				return err
			}
		}
	}
	return nil
}

func displayUniverse(universe *Universe) error {
	for i := range universe.data {
		for j := range universe.data[i] {
			life, err := universe.Get(i, j)
			if err != nil {
				return err
			}

			if life == 0 {
				fmt.Print(" ")
			} else {
				fmt.Print("O")
			}
		}
		fmt.Println()
	}

	return nil
}
