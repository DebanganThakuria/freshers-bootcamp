package main

import (
	"reflect"
	"testing"
)

func TestMatrix_GetRows(t *testing.T) {
	m := Matrix{3, 3, [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}}

	if m.GetRows() != 3 {
		t.Error("GetRows has error!")
	}
}

func TestMatrix_GetColumns(t *testing.T) {
	m := Matrix{3, 3, [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}}

	if m.GetColumns() != 3 {
		t.Error("GetColumns is not working!")
	}
}

func TestMatrix_Update(t *testing.T) {
	m := Matrix{3, 3, [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}}

	if m.Update(0, 0, 10); !reflect.DeepEqual(m.Data, [][]int{{10, 2, 3}, {4, 5, 6}, {7, 8, 9}}) {
		t.Error("Update failed!")
	}

}

func TestMatrix_Add(t *testing.T) {
	m := Matrix{3, 3, [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}}

	x := Matrix{3, 3, [][]int{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}}}

	m.Add(x)
	if !reflect.DeepEqual(m.Data, [][]int{{2, 3, 4}, {5, 6, 7}, {8, 9, 10}}) {
		t.Error("Add Failed!")
	}
}