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

	mouseX, mouseY := ebiten.CursorPosition()

	g.liveDisplayUpdateAchievements()

	switch g.state {
	case stateLanguageSelect:
		if languageSelectUpdate(mouseX) {
			g.state = stateIntro
			g.intro.reset()
		}
	case stateIntro:
		if g.intro.update() {
			g.state = stateTitle
		}
	case stateTitle:
		g.updateTitle(mouseX, mouseY)
		if g.state == statePlay {
			g.playArea = buildPlayArea()
			g.score.reset()
			g.timeHandler.reset()
			g.newAchievementPositions = nil
		}
	case stateCredits, stateHowTo, stateAchievements:
		if inputSelect() {
			g.state = stateTitle
			resetTitle()
		}
	case statePlay:
		if g.updatePlay(mouseX, mouseY) {
			g.state = stateEnd
			g.score.setMax()
			g.setupEnd()
		}
	case stateEnd:
		g.updateEnd()
		if inputSelect() {
			g.state = stateTitle
			resetTitle()
		}
	}

	return nil
}

// Define the keys for choice validation
func inputSelect() bool {
	return inpututil.IsKeyJustPressed(ebiten.KeyEnter) ||
		inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft)
}

func (g *game) updatePlay(mouseX, mouseY int) (finished bool) {
	// find which tile is below the mouse
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
			g.score.update(getTimePoints(g.timeHandler.currentTime), getDemonstrationPoints(g.playArea.demonstrationSize), getCopsPoints(g.playArea.maxCopsSize))
			g.checkAchievements()
			g.playArea.drawNewTile(g.playArea.heldHandTile)
			g.timeHandler.reset()
		}
		g.playArea.holdTile = false
	}

	// update time
	g.timeHandler.update()

	// check end of game
	finished = g.playArea.checkEndOfPlay()
	return
}
