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

type achievement struct {
	obtained    bool
	displayed   bool
	displayTime int
	condition   func(playArea) bool
	frenchText  string
	englishText string
	frenchInfo  string
	englishInfo string
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
			frenchInfo:  "Forme un groupe de policiers",
			englishInfo: "Make a group of cops",
			condition: func(p playArea) bool {
				return p.maxCopsSize >= 2
			},
		},
		achievement{
			frenchText:  "Un peu chaud",
			englishText: "Dangerous protest",
			frenchInfo:  "Forme un grand groupe de policiers",
			englishInfo: "Make a large group of cops",
			condition: func(p playArea) bool {
				return p.maxCopsSize >= globalNumCops/2
			},
		},
		achievement{
			frenchText:  "C'est foutu",
			englishText: "This is the end",
			frenchInfo:  "Forme un très grand groupe de policiers",
			englishInfo: "Make a very large group of cops",
			condition: func(p playArea) bool {
				return p.maxCopsSize >= 3*globalNumCops/4
			},
		},
		achievement{
			frenchText:  "Un début de mobilisation",
			englishText: "Birth of a demonstration",
			frenchInfo:  "Fais grandir la manifestation",
			englishInfo: "Increase the demonstration size",
			condition: func(p playArea) bool {
				return p.demonstrationSize > 4
			},
		},
		achievement{
			frenchText:  "On va y arriver",
			englishText: "Almost there!",
			frenchInfo:  "Crée une grande manifestation",
			englishInfo: "Make a large demonstration",
			condition: func(p playArea) bool {
				return p.demonstrationSize >= globalNumPeople/2
			},
		},
		achievement{
			frenchText:  "Cette fois c'est la bonne",
			englishText: "We will do it!",
			frenchInfo:  "Crée une très grande manifestation",
			englishInfo: "Make a very large demonstration",
			condition: func(p playArea) bool {
				return p.demonstrationSize >= globalNumPeople+4
			},
		},
		achievement{
			frenchText:  "Assignés à résidence",
			englishText: "No protest here!",
			frenchInfo:  "Ne fais pas du tout grandir la manifestation",
			englishInfo: "Do not increase the demonstration size",
			condition: func(p playArea) bool {
				return p.demonstrationSize <= 4 && !p.handHasTile[0] && !p.handHasTile[1] && !p.handHasTile[2] && !p.holdTile
			},
		},
		achievement{
			frenchText:  "La police avec nous",
			englishText: "No cops here!",
			frenchInfo:  "Ne forme aucun groupe de plus de 4 policiers",
			englishInfo: "Do not make any group of cops of size 4 or more",
			condition: func(p playArea) bool {
				return p.maxCopsSize <= 4 && !p.handHasTile[0] && !p.handHasTile[1] && !p.handHasTile[2] && !p.holdTile
			},
		},
		achievement{
			frenchText:  "Trop de policiers",
			englishText: "Too many cops",
			frenchInfo:  "Forme un groupe de policiers plus grand que la manifestation",
			englishInfo: "Make a group of cops that is larger than the demonstration",
			condition: func(p playArea) bool {
				return p.maxCopsSize > p.demonstrationSize
			},
		},
		achievement{
			frenchText:  "Beaucoup trop de policiers",
			englishText: "Far too many cops",
			frenchInfo:  "Termine la partie avec un groupe de policiers plus grand que la manifestation",
			englishInfo: "End the game with a group of cops that is larger than the demonstration",
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
				g.soundManager.NextSounds[soundAchievementID] = true
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
			g.achievements[achievementPos].drawSmall(globalLiveDisplayAchievementsX, achievementY, screen)
			achievementY -= globalLiveDisplayAchievementsHeight + globalLiveDisplayAchievementsSep
		}
	}
}

// small display of an achievement
func (a achievement) drawSmall(x, y int, screen *ebiten.Image) {
	text := a.frenchText
	if language == englishLanguage {
		text = a.englishText
	}

	opt := &ebiten.DrawImageOptions{}
	opt.GeoM.Translate(float64(x), float64(y))
	opt.ColorScale.ScaleAlpha(0.8)
	screen.DrawImage(achievementsmallImage, opt)

	centerX := x + globalLiveDisplayAchievementsWidth/2
	drawSmallTextCenteredAt(text, float64(centerX), float64(y+5), screen)

}

// display achievements screen
func (g game) drawAchievementsScreen(screen *ebiten.Image) {

	x := (globalWidth - globalAchievementWidth) / 2
	y := 20
	yStep := 5

	text := "Trophées"
	if language == englishLanguage {
		text = "Achievements"
	}

	height := drawTextAt(text, float64(x), float64(y), screen)

	y += int(height)
	y += 2 * yStep

	for _, achievement := range g.achievements {

		achievement.draw(x, y, screen)

		y += yStep
		y += globalAchievementHeight
	}

	x = globalWidth/2 + 50

	text = "Clique pour revenir au menu"
	if language == englishLanguage {
		text = "Click to go back to title"
		x += 50
	}

	y += 20

	drawTextAt(text, float64(x), float64(y), screen)
}

// display one achievement in large size
func (a achievement) draw(x, y int, screen *ebiten.Image) {
	text := a.frenchText
	info := a.frenchInfo
	if language == englishLanguage {
		text = a.englishText
		info = a.englishInfo
	}
	theImage := achievementnokImage
	if a.obtained {
		theImage = achievementokImage
	}

	opt := &ebiten.DrawImageOptions{}
	opt.GeoM.Translate(float64(x), float64(y))
	screen.DrawImage(theImage, opt)

	x += 75
	y += 5

	height := drawTextAt(text, float64(x), float64(y), screen)

	y += int(height)

	drawSmallTextAt(info, float64(x), float64(y), screen)
}
