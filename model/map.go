package model

import "fmt"

type Map struct {
	Surface   [][]string
	dimension int
}

func NewMap(m int) *Map {
	nMap := Map{}
	nMap.dimension = m

	for i := 0; i < m; i++ {
		row := make([]string, m)
		for j := 0; j < m; j++ {
			row[j] = "-"
		}
		nMap.Surface = append(nMap.Surface, row)
	}

	return &nMap
}

func (m *Map) SetShip(i, j int) {
	m.Surface[i][j] = "B"

}

func (m *Map) SetDeadShip(i, j int) {
	m.Surface[i][j] = "X"

}

func (m *Map) SetMissile(i, j int) {
	m.Surface[i][j] = "O"
}

func (m Map) IsShip(i, j int) bool {
	return m.Surface[i][j] == "B"
}

func (m Map) ShowMap() {
	for i := 0; i < m.dimension; i++ {
		fmt.Println(m.Surface[i])
	}
}
