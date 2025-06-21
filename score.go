/*
		Union, a game for the 2025 Ebitengine game jam.
		Copyright (C) 2025 Loïg Jezequel

	    This program is free software: you can redistribute it and/or modify
	    it under the terms of the GNU General Public License as published by
	    the Free Software Foundation, either version 3 of the License, or
	    (at your option) any later version.

	    This program is distributed in the hope that it will be useful,
	    but WITHOUT ANY WARRANTY; without even the implied warranty of
	    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
	    GNU General Public License for more details.

	    You should have received a copy of the GNU General Public License
	    along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/
package main

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type score struct {
	fromDemonstration int
	fromCops          int
	fromTime          int
	current           int
	max               int
}

// update the current score according to the effects of the last tile placed
func (s *score) update(timePoints, demonstrationPoints, copsPoints int) (demonstrationsIncrement, copsDecrement int) {
	s.fromTime += timePoints

	demonstrationsIncrement = demonstrationPoints - s.fromDemonstration
	copsDecrement = copsPoints - s.fromCops

	s.fromDemonstration = demonstrationPoints
	s.fromCops = copsPoints

	oldScore := s.current
	s.current = s.fromTime + s.fromDemonstration - s.fromCops
	if s.current < 0 {
		s.current = 0
		if oldScore == 0 {
			return 0, 0
		}
	}

	return
}

// reset the score at the start of a game
func (s *score) reset() {
	s.fromDemonstration = 0
	s.fromCops = 0
	s.fromTime = 0
	s.current = 0
}

// update the max score at the end of a game
func (s *score) setMax() {
	if s.current > s.max {
		s.max = s.current
	}
}

// draw the score
func (s score) draw(screen *ebiten.Image) {
	theScore := fmt.Sprintf("Score: %d", s.current)
	ebitenutil.DebugPrintAt(screen, theScore, globalScoreX, globalScoreY)
}

// get demonstration points from the size of the demonstration
func getDemonstrationPoints(demonstrationSize int) int {
	return demonstrationSize*demonstrationSize - 16
}

// get cops points from the sizes of the largest set of cops
func getCopsPoints(copsSize int) int {
	if copsSize > 1 {
		return copsSize * copsSize * copsSize
	}
	return 0
}

// get time points from the remaining time
func getTimePoints(remainingTime int) int {
	return (remainingTime * 20) / globalAllowedTime
}
