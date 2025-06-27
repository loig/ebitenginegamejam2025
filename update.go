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

	g.soundManager.UpdateMusic(0.5)

	g.soundManager.PlaySounds()

	switch g.state {
	case stateLanguageSelect:
		finished, changed := languageSelectUpdate(mouseX)
		if changed {
			g.soundManager.NextSounds[soundMvtID] = true
		}
		if finished {
			g.state = stateIntro
			g.soundManager.ChangeMusic(introMusicTrack)
			g.intro.reset()
			g.soundManager.NextSounds[soundSelectID] = true
		}
	case stateIntro:
		if g.intro.update() {
			g.state = stateTitle
			g.soundManager.ChangeMusic(titleMusicTrack)
			resetTitle()
			g.soundManager.NextSounds[soundSelectID] = true
		}
	case stateTitle:
		g.updateTitle(mouseX, mouseY)
		if g.state != stateTitle {
			g.soundManager.NextSounds[soundSelectID] = true
		}
		if g.state == statePlay {
			g.soundManager.ChangeMusic(themeMusicTrack)
			g.playArea = buildPlayArea()
			g.score.reset()
			g.timeHandler.reset()
			g.newAchievementPositions = nil
		}
		if g.state == stateIntro {
			g.soundManager.ChangeMusic(introMusicTrack)
			g.intro.reset()
		}
	case stateCredits, stateHowTo, stateAchievements:
		if inputSelect() {
			g.state = stateTitle
			resetTitle()
			g.soundManager.NextSounds[soundSelectID] = true
		}
	case statePlay:
		g.updateTooglePeople(mouseX, mouseY)
		g.particles.update()
		if g.updatePlay(mouseX, mouseY) {
			g.state = stateEndPlay
		}
	case stateEndPlay:
		g.updateTooglePeople(mouseX, mouseY)
		g.particles.update()
		if inputSelect() {
			g.state = stateEnd
			g.soundManager.ChangeMusic(titleMusicTrack)
			g.score.setMax()
			g.setupEnd()
			g.soundManager.NextSounds[soundSelectID] = true
			g.particles.reset()
		}
	case stateEnd:
		g.updateEnd()
		if inputSelect() {
			g.state = stateTitle
			resetTitle()
			g.soundManager.NextSounds[soundSelectID] = true
		}
	}

	return nil
}

// Define the keys for choice validation
func inputSelect() bool {
	return inpututil.IsKeyJustPressed(ebiten.KeyEnter) ||
		inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft)
}

func (g *game) updateTooglePeople(mouseX, mouseY int) {
	if mouseX >= globalWidth-globalTogglePeopleWidth/globalTogglePeopleScale && mouseX < globalWidth &&
		mouseY >= 0 && mouseY < globalTogglePeopleHeight/globalTogglePeopleScale &&
		inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		g.drawPeople = !g.drawPeople
	}
}

func (g *game) updatePlay(mouseX, mouseY int) (finished bool) {
	// find which tile is below the mouse
	g.playArea.updateMousePosition(mouseX, mouseY)

	// get tile
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		if !g.playArea.holdTile && g.playArea.handHover && g.playArea.handHasTile[g.playArea.handHoverPos] {
			g.playArea.holdTile = true
			g.playArea.heldHandTile = g.playArea.handHoverPos
			g.soundManager.NextSounds[soundMvtID] = true
		}
	}

	// define where the held tile shall be drawn
	oldX, oldY := g.playArea.holdX, g.playArea.holdY
	g.playArea.holdX = mouseX - globalTileSize/2
	g.playArea.holdY = mouseY - globalTileSize/2
	if g.playArea.gridHover && !g.playArea.gridHasTile[g.playArea.gridHoverY][g.playArea.gridHoverX] {
		g.playArea.holdX = globalGridX + g.playArea.gridHoverX*globalTileSize
		g.playArea.holdY = globalGridY + g.playArea.gridHoverY*globalTileSize
		if g.playArea.holdTile && (g.playArea.holdX != oldX || g.playArea.holdY != oldY) {
			g.soundManager.NextSounds[soundMvtID] = true
		}
	}

	// release tile
	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		if g.playArea.canDropTile() {
			g.playArea.dropTile()
			timePoints := getTimePoints(g.timeHandler.currentTime)
			if timePoints != 0 {
				g.particles.addParticle(globalTimeX+120, globalTimeY, timePoints)
			}
			demonstrationIncrement, copsIncrement := g.score.update(timePoints, getDemonstrationPoints(g.playArea.demonstrationSize), getCopsPoints(g.playArea.maxCopsSize))
			if demonstrationIncrement != 0 || copsIncrement != 0 {
				g.particles.addParticle(g.playArea.holdX+globalTileSize/2, g.playArea.holdY+globalTileSize/2-20, demonstrationIncrement+copsIncrement)
			}
			g.checkAchievements()
			g.playArea.drawNewTile(g.playArea.heldHandTile)
			g.timeHandler.reset()
			g.soundManager.NextSounds[soundDropID] = true
		}
		g.playArea.holdTile = false
	}

	// update time
	g.timeHandler.update()

	// check end of game
	finished = g.playArea.checkEndOfPlay()
	if finished {
		g.checkAchievements()
	}
	return
}
