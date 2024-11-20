package main

import (
	"errors"
	"fmt"
	"math/rand"
)

type Universe struct {
	size int
	data [][]int
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

func main() {
	var size, numbeSeed int
	fmt.Scan(&size, &numbeSeed)

	rand.Seed(int64(numbeSeed))

	universe, err := newUniverse(size)
	if err != nil {
		fmt.Println(err)
	}

	sparkLife(universe)
	displayUniverse(universe)

	// fmt.Println(universe.data)
}

func newUniverse(size int) (*Universe, error) {
	if size <= 0 {
		return nil, errors.New("rand.Intn(2)size must be greater than 0")
	}

	data := make([][]int, size)

	for i := range data {
		data[i] = make([]int, size)
	}

	return &Universe{size: size, data: data}, nil
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
