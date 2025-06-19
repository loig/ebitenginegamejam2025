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
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type playArea struct {
	grid                   [globalGridHeight][globalGridWidth]tile
	gridHasTile            [globalGridHeight][globalGridWidth]bool
	gridHoverX, gridHoverY int
	gridHover              bool
	hand                   [globalHandSize]tile
	handHasTile            [globalHandSize]bool
	handHoverPos           int
	handHover              bool
	heldHandTile           int
	holdTile               bool
	holdX, holdY           int
	deck                   [globalDeckSize]tile
	deckPos                int
}

// build a fresh area
func buildPlayArea() playArea {
	var res playArea
	res.gridHasTile[globalGridHeight/2][globalGridWidth/2] = true
	for pos := 0; pos < globalHandSize; pos++ {
		res.drawNewTile(pos)
	}
	return res
}

// update information about what is below the mouse
func (p *playArea) updateMousePosition(mouseX, mouseY int) {
	p.gridHoverX = (mouseX - globalGridX) / globalTileSize
	p.gridHoverY = (mouseY - globalGridY) / globalTileSize
	p.gridHover = (mouseX-globalGridX) >= 0 && (mouseY-globalGridY) >= 0 &&
		p.gridHoverX >= 0 && p.gridHoverX < globalGridWidth &&
		p.gridHoverY >= 0 && p.gridHoverY < globalGridHeight
	p.handHoverPos = (mouseY - globalHandY) / (globalTileSize + globalHandSep)
	p.handHover = p.handHoverPos >= 0 && p.handHoverPos < globalHandSize &&
		(mouseX-globalHandX) >= 0 && (mouseX-globalHandX)/globalTileSize == 0 &&
		(mouseY-globalHandY) >= 0 && (mouseY-globalHandY) < p.handHoverPos*(globalTileSize+globalHandSep)+globalTileSize

}

// check if a tile can be dropped at current mouse position
func (p playArea) canDropTile() bool {
	return p.gridHover && !p.gridHasTile[p.gridHoverY][p.gridHoverX] &&
		((p.gridHoverX-1 >= 0 && p.gridHasTile[p.gridHoverY][p.gridHoverX-1]) ||
			(p.gridHoverX+1 < globalGridWidth && p.gridHasTile[p.gridHoverY][p.gridHoverX+1]) ||
			(p.gridHoverY-1 >= 0 && p.gridHasTile[p.gridHoverY-1][p.gridHoverX]) ||
			(p.gridHoverY+1 < globalGridHeight && p.gridHasTile[p.gridHoverY+1][p.gridHoverX]))
}

func (p *playArea) drawNewTile(handPos int) {
	if p.deckPos < len(p.deck) {
		p.hand[handPos] = p.deck[p.deckPos]
		p.deckPos++
		p.handHasTile[handPos] = true
		return
	}
	p.handHasTile[handPos] = false
}

type tile struct {
	east  tileContent
	north tileContent
	south tileContent
	west  tileContent
}

func (t tile) draw(tileX, tileY int, screen *ebiten.Image) {
	vector.DrawFilledRect(screen, float32(tileX), float32(tileY), float32(globalTileSize), float32(globalTileSize), color.RGBA{R: 150, G: 0, B: 0, A: 255}, false)
}

type tileContent = int

const (
	contentCity int = iota
	contentCop
	contentNature
	contentPeople
)
