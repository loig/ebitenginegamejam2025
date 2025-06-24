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
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func (g *game) Draw(screen *ebiten.Image) {

	g.liveDisplayDrawAchievements(screen)

	switch g.state {
	case stateLanguageSelect:
		languageSelectDraw(screen)
	case stateIntro:
		g.intro.draw(screen)
	case stateTitle:
		g.drawTitle(screen)
	case stateCredits:
		if language == frenchLanguage {
			ebitenutil.DebugPrint(screen, "Crédits")
		} else {
			ebitenutil.DebugPrint(screen, "Credits")
		}
	case stateHowTo:
		if language == frenchLanguage {
			ebitenutil.DebugPrint(screen, "Comment jouer")
		} else {
			ebitenutil.DebugPrint(screen, "How to play")
		}
	case stateAchievements:
		g.drawAchievementsScreen(screen)
	case statePlay:
		g.drawPlay(screen)
	case stateEnd:
		g.drawEnd(screen)
	}

}

func (g game) drawPlay(screen *ebiten.Image) {

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
	g.score.drawCurrentAt(screen, globalPlayScoreX, globalPlayScoreY)

	// Draw the time
	g.timeHandler.draw(screen)
}
