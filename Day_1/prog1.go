package main

import (
	"encoding/json"
	"fmt"
)

type Matrix struct{
	Rows int `json:"rows"`
	Columns int `json:"columns"`
	Data [][]int `json:"data"`
}

func (m *Matrix) GetRows() int {
	return m.Rows
}

func (m *Matrix) GetColumns() int {
	return m.Columns
}

func (m *Matrix) Update(i, j, x int) {
	m.Data[i][j] = x
}

func (m *Matrix) Add(x Matrix) {
	for i := 0; i < m.Rows; i++ {
		for j := 0 ; j < m.Columns; j++ {
			m.Data[i][j] += x.Data[i][j]
		}
	}
}

func (m *Matrix) JsonDump() {
	b, err := json.MarshalIndent(m, "", " ")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(b))
}

func main() {
	var m1, m2 Matrix

	m1 = Matrix{
		3,
		3,
		[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
	}

	m2 = Matrix{
		3,
		3,
		[][]int{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}},
	}

	fmt.Println(m1.GetRows())
	fmt.Println(m1.GetColumns())
	m1.Update(0, 0, 10)
	fmt.Println(m1.Data)
	m1.Add(m2)
	fmt.Println(m1.Data)
	m1.JsonDump()
}
