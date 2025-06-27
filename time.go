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
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type timeHandler struct {
	currentTime int
}

func (t *timeHandler) reset() {
	t.currentTime = globalAllowedTime
}

func (t *timeHandler) update() {
	if t.currentTime > 0 {
		t.currentTime--
	}
}

func (t timeHandler) draw(screen *ebiten.Image) {

	points := getTimePoints(t.currentTime)
	prefix := "Bonus :"
	if language == englishLanguage {
		prefix = "Bonus:"
	}
	text := fmt.Sprintf("%s %d", prefix, points)

	drawTextAt(text, float64(globalTimeX), float64(globalTimeY+5), screen)

	timeWidth := (globalTimeWidth * t.currentTime) / globalAllowedTime
	vector.DrawFilledRect(screen, float32(globalTimeBarX), float32(globalTimeY), float32(timeWidth), float32(globalTimeHeight), color.White, false)
}
