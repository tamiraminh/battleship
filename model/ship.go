package model

import (
	"errors"
	"strconv"
	"strings"
)

type Player struct {
	Name          string
	Map           *Map
	ActiveShips   int
	NumberMissile int
}

func NewPlayer(name string, nship int, shipPositionRaw string, m int, nMissile int) *Player {
	p := Player{}
	p.Name = name
	p.Map = NewMap(m)
	p.ActiveShips = nship
	p.NumberMissile = nMissile

	positions := strings.Split(shipPositionRaw, ":")
	for _, pos := range positions {
		arrIndex := strings.Split(pos, ",")
		i, _ := strconv.Atoi(arrIndex[0])
		j, _ := strconv.Atoi(arrIndex[1])

		p.Map.SetShip(i, j)

	}

	return &p
}

func (p *Player) AttackPlayer(missile string, p2 *Player) (err error) {
	missilePos := strings.Split(missile, ":")
	if len(missilePos) > p.NumberMissile {
		return errors.New("missile not sufficient")
	}

	for _, pos := range missilePos {
		arrIndex := strings.Split(pos, ",")
		i, _ := strconv.Atoi(arrIndex[0])
		j, _ := strconv.Atoi(arrIndex[1])

		if p2.Map.IsShip(i, j) {
			p2.Map.SetDeadShip(i, j)
			p2.ActiveShips -= 1
		} else {
			p2.Map.SetMissile(i, j)
		}
	}

	return
}
