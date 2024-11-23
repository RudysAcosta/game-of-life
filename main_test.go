package main

import (
	"testing"
)

func TestGetneighbours(t *testing.T) {
	cell := Cell{x: 7, y: 0, limit: 8}
	neighbours := cell.GetNeighbours()
	if len(neighbours) != 8 {
		t.Errorf("Expected 8 neighbours, got %d", len(neighbours))
	}

	expected := []Coordinate{
		{6, 0, 8},
		{6, 7, 8},
		{7, 7, 8},
		{0, 7, 8},
		{0, 0, 8},
		{0, 1, 8},
		{7, 1, 8},
		{6, 1, 8},
	}

	for i, n := range neighbours {
		if n != expected[i] {
			t.Errorf("Expected %v, got %v", expected[i], n)
		}
	}
}
