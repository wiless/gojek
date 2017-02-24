package main

import (
	"errors"
	"fmt"
	"strings"
)

const NPlayers int = 2

type ShipStatus int

const (
	Empty ShipStatus = iota
	ActiveShip
	DeadShip
	Missed
)

var shipstatus = []string{"_", "B", "X", "O"}

func (s ShipStatus) String() string {
	return shipstatus[s]
}

type matrix [][]ShipStatus
type Board struct {
	GridSize      int
	TotalShips    int
	ShipPositions [][]Pos // Ship Positions for Each Player
	Grid          [NPlayers]matrix
	Hits          [NPlayers]int // Stores the number of Hits of each Player
}

func (b *Board) LoadBoard(bstr []byte) error {
	inputstr := string(bstr)
	lines := strings.Split(inputstr, "\n")
	if len(lines) < 5 {
		return errors.New("Not enough lines")
	}

	for i := 0; i < 5; i++ {
		line := lines[i]
		switch i {
		case 0:
			b.GridSize = ParseInt(line)
			b.Grid[0] = make(matrix, b.GridSize)
			b.Grid[1] = make(matrix, b.GridSize) // b.Grid[0] //make(matrix, b.GridSize)
			for r := 0; r < b.GridSize; r++ {
				b.Grid[0][r] = make([]ShipStatus, b.GridSize) // for both players
				b.Grid[1][r] = make([]ShipStatus, b.GridSize)
			}

		case 1:
			b.TotalShips = ParseInt(line)
			b.ShipPositions = make([][]Pos, 2)

			for p := 0; p < NPlayers; p++ {
				b.ShipPositions[p] = make([]Pos, b.TotalShips)
			}
			// if b.TotalShips > b.GridSize/2 {
			// 	return fmt.Errorf("No of Ships %d cannot be > %d", b.TotalShips, b.GridSize/2)
			// }

		case 2, 3:
			playerindx := i - 2
			pos, err := ParseXYs(line)
			if err != nil {
				return err
			}
			if len(pos) != b.TotalShips {
				return fmt.Errorf("Mismatch # of positions (%d) passed for %d Ships", len(pos), b.TotalShips)
			}
			// Store Ship Postion of player and Tag in the grid
			for i, p := range pos {
				if p.X >= b.GridSize || p.Y >= b.GridSize {
					return fmt.Errorf("Position beyond Grid size !! for Player %d", 2)
				}
				b.ShipPositions[playerindx][i] = p
				if b.Grid[playerindx][p.X][p.Y] != 0 {
					return fmt.Errorf("Something Wrong !! Cant have two Ships in same place @(%v) !! ", p)
				}
				// if playerindx == 0 {
				b.Grid[playerindx][p.X][p.Y] = ActiveShip // Tag with Player 1
			}

		}
	}

	return nil
}

func (b Board) Result() string {
	diff := b.Hits[0] - b.Hits[1]
	switch {
	case diff > 0:
		return "Player 1 Won"
	case diff < 0:
		return "Player 2 Won"
	}

	return "Its a draw"
}

func (b *Board) LaunchMissile(playerindx int, p Pos) (hit bool) {
	var targetGrid int
	if playerindx == 0 {
		targetGrid = 1
	} // Note other case is taken care

	if b.Grid[targetGrid][p.X][p.Y] == ActiveShip {
		b.Grid[targetGrid][p.X][p.Y] = DeadShip
		b.Hits[playerindx]++

		return true
	}

	/// Rehitting a DeadShip or Empty ship is Missed !!
	b.Grid[targetGrid][p.X][p.Y] = Missed
	// May be hit twice !! ??
	return false
}

func ParseMissileActions(b []byte) (missiles [][]Pos, e error) {
	str := string(b)
	tmp := strings.SplitAfterN(str, "\n", 5)
	actiontxt := tmp[4] // the initial 4 rows are for board

	lines := strings.Split(actiontxt, "\n")
	totalMissiles := ParseInt(lines[0])
	missiles = make([][]Pos, 2) //
	// Load same number of missiles for both player
	missiles[0] = make([]Pos, TotalMissiles)
	missiles[1] = make([]Pos, TotalMissiles)
	var err error

	missiles[0], err = ParseXYs(lines[1])

	if err != nil {
		return missiles, err
	}
	if totalMissiles != len(missiles[0]) {
		return missiles, fmt.Errorf("#Missiles != Positions")
	}

	missiles[1], err = ParseXYs(lines[2])
	if err != nil {
		return missiles, err
	}
	if totalMissiles != len(missiles[1]) {
		return nil, fmt.Errorf("#Missiles != Positions")
	}

	return missiles, nil
}
