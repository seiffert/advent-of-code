package main

import "testing"

func TestIsTriangle(t *testing.T) {
	triangles := [][]int{
		{2, 2, 3},
		{3, 4, 5},
		{1, 2, 2},
	}
	for _, triangle := range triangles {
		if !IsTriangle(triangle[0], triangle[1], triangle[2]) {
			t.Errorf("%d/%d/%d should be a triangle", triangle[0], triangle[1], triangle[2])
		}
	}
}

func TestIsNoTriangle(t *testing.T) {
	noTriangles := [][]int{
		{1, 1, 3},
		{50000, 1, 2},
		{7, 3, 1},
	}
	for _, triangle := range noTriangles {
		if IsTriangle(triangle[0], triangle[1], triangle[2]) {
			t.Errorf("%d/%d/%d should not be be a triangle", triangle[0], triangle[1], triangle[2])
		}
	}
}
