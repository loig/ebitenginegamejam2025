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
)

func (g *game) Draw(screen *ebiten.Image) {

	switch g.state {
	case stateLanguageSelect:
		languageSelectDraw(screen)
	case stateIntro:
		g.intro.draw(screen)
	case stateTitle:
		g.drawTitle(screen)
	case stateCredits:
		drawCredits(screen)
	case stateHowTo:
		drawHowToPlay(screen)
	case stateAchievements:
		g.drawAchievementsScreen(screen)
	case statePlay:
		g.drawPlay(false, screen)
	case stateEndPlay:
		g.drawPlay(true, screen)
	case stateEnd:
		g.drawEnd(screen)
	}

	g.liveDisplayDrawAchievements(screen)

}

func (g game) drawPlay(finished bool, screen *ebiten.Image) {

	// Draw the background
	opt := &ebiten.DrawImageOptions{}
	screen.DrawImage(backgroundImage, opt)

	// Draw the grid
	tileY := globalGridY
	for y, line := range g.playArea.grid {
		tileX := globalGridX
		for x, tile := range line {
			if g.playArea.gridHasTile[y][x] {
				tile.draw(tileX, tileY, screen)
			}
			tileX += globalTileSize
		}
		tileY += globalTileSize
	}

	// Draw the hand
	tileX, tileY := globalHandX, globalHandY
	for pos, tile := range g.playArea.hand {
		//mouseHover := (g.playArea.handHoverPos == pos) && g.playArea.handHover
		if g.playArea.handHasTile[pos] && !(g.playArea.holdTile && g.playArea.heldHandTile == pos) {
			tile.draw(tileX, tileY, screen)
		}
		tileY += globalTileSize + globalHandSep
	}

	// Draw the held tile
	if g.playArea.holdTile {
		g.playArea.hand[g.playArea.heldHandTile].draw(g.playArea.holdX, g.playArea.holdY, screen)
	}

	// Draw the score
	g.score.drawCurrentAt(screen, globalPlayScoreX, globalPlayScoreY, false)

	if !finished {
		// Draw the time
		screen.DrawImage(bonusImage, opt)
		g.timeHandler.draw(screen)
	} else {
		// Draw end game message
		text := "C'est fini, clique pour voir les résultats"
		if language == englishLanguage {
			text = "End of the game, click to see the results"
		}
		x := globalWidth / 2
		drawTextCenteredAt(text, float64(x), float64(globalTimeY), screen)
	}
}
