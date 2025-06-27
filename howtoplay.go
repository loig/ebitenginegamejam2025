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

import "github.com/hajimehoshi/ebiten/v2"

var frenchHowTo []string = []string{
	"Ceci est une tuile. Avec la souris on peut glisser/déposer\nde nouvelles tuiles sur la surface de jeu. On ne peut po-\nser les tuiles qu'à côté d'autres tuiles. Une fois posée,",
	"une tuile ne peut plus être déplacée. La partie se termine quand\nla surface de jeu est remplie.",
	"Chaque tuile est séparée en quatre quarts qui peuvent chacun\ncontenir un type de personnage (citoyen ou policier) ou être vide.",
	"Ceci est un citoyen.",
	"Ceci est un policier. Si des policiers se touchent ils\nforment un groupe.",
	"Ceci est aussi un policier. Il en représente le plus grand\ngroupe. Ce groupe fait perdre des points selon sa taille.",
	"Ceci est la manifestation. Les citoyens qui la touchent\nla rejoignent. Elle fait gagner des points selon sa taille.",
	"En plus de ça, une tuile posée rapidement rapporte un bonus.",
	"Clique pour revenir au menu !",
}

var englishHowTo []string = []string{
	"This is a tile. Use your mouse to drag and drop new tiles\nto the play area. Tiles must be placed near other tiles.\nOnce it is placed, a tile can no longer be moved. The",
	"game ends as soon as the play area is filled.",
	"Each tile has four quarters. Each quarters may either be empty\nor contain a character: citizen or cop.",
	"This is a citizen.",
	"This is a cop. Side by side cops form a group.",
	"This is also a cop, representing the largest group of cops.\nYou lose points according to the size of this group.",
	"This is the demonstration. Citizen next to it will join the\nprotest. You gain points according to its size.",
	"Moreover, placing tiles quickly will give you a point bonus.",
	"Click to go back to title screen!",
}

func drawHowToPlay(screen *ebiten.Image) {

	howTo := frenchHowTo
	if language == englishLanguage {
		howTo = englishHowTo
	}

	tile1 := tile{content: [4]tileContent{
		contentCop,
		contentPeople,
		contentPeople,
		contentCity,
	}, people: [4][8]int{copSet, peopleSet, peopleSet, copSet}}

	x, y := 20, 20
	tile1.draw(x, y, screen, true)
	height := drawTextAt(howTo[0], float64(x+100), float64(y), screen)

	y += int(height) + 5
	height = drawTextAt(howTo[1], float64(x), float64(y), screen)

	y += int(height)
	y += 20

	height = drawTextAt(howTo[2], float64(x), float64(y), screen)
	y += int(height)
	y += 20

	var tile2 tile
	tile2.content[0] = contentPeople
	tile2.people[0] = peopleSet

	tile2.draw(x, y, screen, true)
	drawTextAt(howTo[3], float64(x+100), float64(y), screen)

	y += 100

	var tile3 tile
	tile3.content[0] = contentCop
	tile3.people[0] = copSet

	tile3.draw(x, y, screen, true)
	drawTextAt(howTo[4], float64(x+100), float64(y), screen)

	y += 100
	var tile4 tile
	tile4.content[0] = contentManyCops
	tile4.people[0] = manyCopsSet

	tile4.draw(x, y, screen, true)
	height = drawTextAt(howTo[5], float64(x+100), float64(y), screen)

	y += 100

	var tile5 tile
	tile5.content = [4]int{contentDemonstration, contentDemonstration, contentDemonstration, contentDemonstration}
	tile5.contentGroupSize = [4]int{4, 4, 4, 4}
	tile5.neighboring = [4]int{6, 12, 9, 3}
	tile5.people = [4][8]int{
		{0, 1, 2, 3, 4, 5, 6, 7},
		{8, 9, 10, 11, 12, 13, 14, 15},
		{0, 1, 2, 3, 4, 5, 6, 7},
		{8, 9, 10, 11, 12, 13, 14, 15},
	}

	tile5.draw(x, y, screen, true)
	height = drawTextAt(howTo[6], float64(x+100), float64(y), screen)

	y += 100

	height = drawTextAt(howTo[7], float64(x), float64(y), screen)

	y += int(height)
	y += 30

	if language == englishLanguage {
		y += 40
	}

	drawTextAt(howTo[8], float64(x+400), float64(y), screen)

}
