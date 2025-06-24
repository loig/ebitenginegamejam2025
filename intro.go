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

type intro struct {
	frame    int
	position int
	content  []string
}

var frenchIntroContent []string = []string{
	"2087, tous les syndicats sans exceptions\narrivent enfin à s'unir derrière un but commun :\nla fin du capitalisme.",
	"Affaiblis par leurs prises de positions\ntrop timorées depuis de nombreuses années\nil ne leur sera pas facile de réunir assez\nde citoyens derrières eux pour atteindre\nl'objectif qu'ils se sont fixé.",
	"Le jour de la grande manifestation unitaire est arrivé !\n D'ici quelques heures tous les malheurs du monde\nne seront peut-être plus qu'un lointain souvenir.",
	"Clique pour jouer !",
}

var englishIntroContent []string = []string{
	"2087, all the workers unions around the world\nare finally united towards one goal.\nThey will end capitalism.",
	"However, they are weakened by years and years of\ntimid fights. They may not be able to unit sufficiently\nmany citizens for the accomplishment of their goal.",
	"Today is the day: the biggest demonstration ever may happen.\nIn a few hours the world may become a better place.",
	"Clic to play!",
}

// set up the intro
func (i *intro) reset() {
	i.position = 0
	i.frame = 0
	switch language {
	case frenchLanguage:
		i.content = frenchIntroContent
	case englishLanguage:
		i.content = englishIntroContent
	}
}

// update the intro
func (i *intro) update() (finished bool) {
	if inputSelect() {
		i.position++
		i.frame = 0
	}

	i.frame++
	if i.frame >= globalIntroTime {
		i.position++
		i.frame = 0
	}

	return i.position >= len(i.content)
}

// draw the intro
func (i intro) draw(screen *ebiten.Image) {
	position := i.position
	if position >= len(i.content) {
		position = len(i.content) - 1
	}

	shift := globalIntroY
	for pos := 0; pos <= position; pos++ {
		height := drawTextCenteredAt(i.content[pos], globalWidth/2, shift, screen)
		shift += height + globalIntroSep
	}
}
