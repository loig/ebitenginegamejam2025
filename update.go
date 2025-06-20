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
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func (g *game) Update() error {

	// find which tile is below the mouse
	mouseX, mouseY := ebiten.CursorPosition()
	g.playArea.updateMousePosition(mouseX, mouseY)

	// get tile
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		if !g.playArea.holdTile && g.playArea.handHover && g.playArea.handHasTile[g.playArea.handHoverPos] {
			g.playArea.holdTile = true
			g.playArea.heldHandTile = g.playArea.handHoverPos
		}
	}

	// define where the held tile shall be drawn
	g.playArea.holdX = mouseX - globalTileSize/2
	g.playArea.holdY = mouseY - globalTileSize/2
	if g.playArea.gridHover && !g.playArea.gridHasTile[g.playArea.gridHoverY][g.playArea.gridHoverX] {
		g.playArea.holdX = globalGridX + g.playArea.gridHoverX*globalTileSize
		g.playArea.holdY = globalGridY + g.playArea.gridHoverY*globalTileSize
	}

	// release tile
	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		if g.playArea.canDropTile() {
			g.playArea.dropTile()
			g.playArea.drawNewTile(g.playArea.heldHandTile)
		}
		g.playArea.holdTile = false
	}

	return nil
}
