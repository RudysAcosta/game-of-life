package main

import (
	"testing"
)

func TestGetneighbours(t *testing.T) {
	u, _ := NewUniverse(3)
	u.Set(0, 0, 1)
	u.Set(0, 1, 1)
	u.Set(0, 2, 1)
	u.Set(1, 0, 1)
	u.Set(1, 1, 1)
	u.Set(1, 2, 1)
	u.Set(2, 0, 1)
	u.Set(2, 1, 1)
	u.Set(2, 2, 1)

	neighbours := u.countAliveNeighbours(1, 1)
	if neighbours != 8 {
		t.Errorf("Expected 8, got %d", neighbours)
	}
}
