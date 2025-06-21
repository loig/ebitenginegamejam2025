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
	"sort"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func (g *game) setupEnd() {
	g.endFrames = 0
	g.endAchievementPosition = 0
	sort.Slice(g.newAchievementPositions, func(i, j int) bool {
		return g.achievements[g.newAchievementPositions[i]].number < g.achievements[g.newAchievementPositions[j]].number
	})
}

func (g *game) updateEnd() {
	if len(g.newAchievementPositions) > 0 {
		g.endFrames++
		if g.endFrames >= globalEndDisplayAchievementTime {
			g.endFrames = 0
			g.endAchievementPosition = (g.endAchievementPosition + 1) % len(g.newAchievementPositions)
		}
	}
}

func (g game) drawEnd(screen *ebiten.Image) {

	// display score
	g.score.drawCurrentAt(screen, globalEndScoreX, globalEndScoreY)

	// display max score
	g.score.drawMaxAt(screen, globalEndMaxScoreX, globalEndMaxScoreY)

	// display the new achievements
	if len(g.newAchievementPositions) > 0 {
		achievementText := g.achievements[g.newAchievementPositions[g.endAchievementPosition]].frenchText
		if language == englishLanguage {
			achievementText = g.achievements[g.newAchievementPositions[g.endAchievementPosition]].englishText
		}
		ebitenutil.DebugPrintAt(screen, achievementText, globalEndDisplayAchievementX, globalEndDisplayAchievementY)
	}
}
