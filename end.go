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
)

type goal struct {
	frenchText  string
	englishText string
	frenchInfo  string
	englishInfo string
	minPoints   int
	content     tileContent
}

var goals []goal = []goal{
	{
		frenchText:  "Dictature",
		englishText: "Dictatorship",
		frenchInfo:  "La manifestation a été un échec total.\nLe gouvernement en a profité pour mettre en prison la plupart des leaders\net installer un régime autoritaire sous prétexte de protéger les citoyens.",
		englishInfo: "The protest has been a huge failure.\nThe government used it as a lever to setup an authoritarian regime.",
		minPoints:   0,
		content:     contentManyCops,
	},
	{
		frenchText:  "Statu quo",
		englishText: "Status quo",
		frenchInfo:  "Une petite manifestation sans effet. Rien n'a changé.",
		englishInfo: "A small protest with almost no effect.",
		minPoints:   1,
		content:     contentCop,
	},
	{
		frenchText:  "Petite victoire",
		englishText: "A small victory",
		frenchInfo:  "La manifestation était un peu plus grosse que d'habitude.\nLe gouvernement a été contraint de changer sa politique.\nLa qualité de vie s'améliore progressivement.",
		englishInfo: "The protest was a bit larger than expected.\nThe government had to accept a political change.\nThe quality of life of the citizens improves step by step.",
		minPoints:   5000,
		content:     contentPeople,
	},
	{
		frenchText:  "Victoire !",
		englishText: "Victory!",
		frenchInfo:  "La manifestation était gigantesque.\nLe gouvernement a démissionné, un conseil de citoyen a pris sa place\net drastiquement changé de ligne politique. Devant les progrès manifestes\npour le bien être de tous, les pays voisins ont rapidement suivi.\nQuelques mois plus tard le capitalisme est définitivement de l'histoire ancienne.",
		englishInfo: "This was the biggest protest ever.\nThe government had to quit. A board of citizens was put in charge of the country.\nThey significantly changed the law.\nSeeing the astonishing progress of the country well being,\nneighboring countries adopted the same policies.\nA few month later capitalism was ancient history.",
		minPoints:   15000,
		content:     contentDemonstration,
	},
}

func (g *game) setupEnd() {
	g.endFrames = 0
	g.endAchievementPosition = 0
	sort.Slice(g.newAchievementPositions, func(i, j int) bool {
		return g.achievements[g.newAchievementPositions[i]].number < g.achievements[g.newAchievementPositions[j]].number
	})

	g.endGoalReached = goals[0]

	for _, theGoal := range goals {
		if theGoal.minPoints > g.endGoalReached.minPoints && theGoal.minPoints <= g.score.current {
			g.endGoalReached = theGoal
		}
	}

	generateCharacters(g.endGoalReached.content)
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

	drawCharacters(screen, 0.3)

	// display result
	height := g.endGoalReached.draw(screen, globalWidth/2, globalEndGoalY)

	// display score
	g.score.drawCurrentAt(screen, globalEndScoreX, globalEndGoalY+height+10, true)

	// display max score
	g.score.drawMaxAt(screen, globalEndMaxScoreX, globalEndGoalY+height+40)

	// display the new achievements
	if len(g.newAchievementPositions) > 0 {
		text := "Nouveaux trophées"
		if language == englishLanguage {
			text = "New achievements"
		}
		drawTextCenteredAt(text, float64(globalWidth/2), float64(globalEndDisplayAchievementY-30), screen)
		g.achievements[g.newAchievementPositions[g.endAchievementPosition]].draw(globalEndDisplayAchievementX, globalEndDisplayAchievementY, screen)
	}
}

func (g goal) draw(screen *ebiten.Image, x, y int) (size int) {

	text := g.frenchText
	info := g.frenchInfo

	if language == englishLanguage {
		text = g.englishText
		info = g.englishInfo
	}

	height := drawTextCenteredAt(text, float64(x), float64(y), screen)

	y += int(height) + 5

	size += int(height)

	height = drawSmallTextCenteredAt(info, float64(x), float64(y), screen)

	size += int(height)

	return
}
