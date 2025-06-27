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
	//screenX, screenY int
}

type achievementSet = []achievement

// build the actual achievements
func makeAchievementSet() achievementSet {
	theSet := achievementSet{
		achievement{
			frenchText:  "Quelques policiers",
			englishText: "A few cops",
			condition: func(p playArea) bool {
				return p.maxCopsSize >= 2
			},
		},
		achievement{
			frenchText:  "Un peu chaud",
			englishText: "Dangerous protest",
			condition: func(p playArea) bool {
				return p.maxCopsSize >= globalNumCops/2
			},
		},
		achievement{
			frenchText:  "C'est foutu",
			englishText: "This is the end",
			condition: func(p playArea) bool {
				return p.maxCopsSize >= 3*globalNumCops/4
			},
		},
		achievement{
			frenchText:  "Un début de mobilisation",
			englishText: "Birth of a demonstration",
			condition: func(p playArea) bool {
				return p.demonstrationSize > 4
			},
		},
		achievement{
			frenchText:  "On va y arriver",
			englishText: "Almost there!",
			condition: func(p playArea) bool {
				return p.demonstrationSize >= globalNumPeople/2
			},
		},
		achievement{
			frenchText:  "Cette fois c'est la bonne",
			englishText: "We will do it!",
			condition: func(p playArea) bool {
				return p.demonstrationSize >= globalNumPeople+4
			},
		},
		achievement{
			frenchText:  "Assignés à résidence",
			englishText: "No protest here!",
			condition: func(p playArea) bool {
				return p.demonstrationSize <= 4 && !p.handHasTile[0] && !p.handHasTile[1] && !p.handHasTile[2] && !p.holdTile
			},
		},
		achievement{
			frenchText:  "La police avec nous",
			englishText: "No cops here!",
			condition: func(p playArea) bool {
				return p.maxCopsSize <= 4 && !p.handHasTile[0] && !p.handHasTile[1] && !p.handHasTile[2] && !p.holdTile
			},
		},
		achievement{
			frenchText:  "Trop de policiers",
			englishText: "Too many cops",
			condition: func(p playArea) bool {
				return p.maxCopsSize > p.demonstrationSize
			},
		},
		achievement{
			frenchText:  "Beaucoup prop de policiers",
			englishText: "Far too many cops",
			condition: func(p playArea) bool {
				return p.maxCopsSize > p.demonstrationSize && !p.handHasTile[0] && !p.handHasTile[1] && !p.handHasTile[2] && !p.holdTile
			},
		},
	}

	for pos := range theSet {
		theSet[pos].number = pos
	}

	return theSet
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

// display achievements screen
func (g game) drawAchievementsScreen(screen *ebiten.Image) {

	x := 20
	y := 20
	yStep := 20
	for _, achievement := range g.achievements {
		text := achievement.frenchText
		if language == englishLanguage {
			text = achievement.englishText
		}
		if achievement.obtained {
			text += " (Ok)"
		}
		ebitenutil.DebugPrintAt(screen, text, x, y)
		y += yStep
	}

}
