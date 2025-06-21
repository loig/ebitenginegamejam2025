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

type achievement struct {
	obtained    bool
	displayed   bool
	displayTime int
	condition   func(playArea) bool
	frenchText  string
	englishText string
	number      int
}

type achievementSet = []achievement

// build the actual achievements
func makeAchievementSet() achievementSet {
	return achievementSet{
		achievement{
			frenchText:  "Quelques policiers",
			englishText: "A few cops",
			condition: func(p playArea) bool {
				return p.maxCopsSize >= 2
			},
			number: 1,
		},
		achievement{
			frenchText:  "Un début de mobilisation",
			englishText: "The birth of a demonstration",
			condition: func(p playArea) bool {
				return p.demonstrationSize > 4
			},
			number: 2,
		},
	}
}

// check if some achievements are validated by the current playArea
func (g *game) checkAchievements() {
	for achievementPos := range g.achievements {
		if !g.achievements[achievementPos].obtained {
			g.achievements[achievementPos].obtained = g.achievements[achievementPos].condition(g.playArea)
			if g.achievements[achievementPos].obtained {
				g.newAchievementPositions = append(g.newAchievementPositions, achievementPos)
			}
		}
	}
}

// live display achievements when obtained
func (g *game) liveDisplayUpdateAchievements() {
	for _, achievementPos := range g.newAchievementPositions {
		if !g.achievements[achievementPos].displayed {
			g.achievements[achievementPos].displayTime++
			if g.achievements[achievementPos].displayTime > globalLiveDisplayTimeAchievement {
				g.achievements[achievementPos].displayed = true
			}
		}
	}
}

func (g game) liveDisplayDrawAchievements(screen *ebiten.Image) {
	achievementY := globalLiveDisplayAchievementsY
	for pos := len(g.newAchievementPositions) - 1; pos >= 0; pos-- {
		achievementPos := g.newAchievementPositions[pos]
		if !g.achievements[achievementPos].displayed {
			text := g.achievements[achievementPos].frenchText
			if language == englishLanguage {
				text = g.achievements[achievementPos].englishText
			}
			ebitenutil.DebugPrintAt(screen, text, globalLiveDisplayAchievementsX, achievementY)
			achievementY -= globalLiveDisplayAchievementsSize + globalLiveDisplayAchievementsSep
		}
	}
}
