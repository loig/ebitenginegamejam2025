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

type intro struct {
	position int
	content  []string
}

var frenchIntroContent []string = []string{
	"Bonjour",
	"C'est l'intro en français",
	"Et ouais!",
}

var englishIntroContent []string = []string{
	"Hello",
	"This is the english version of the intro",
}

// set up the intro
func (i *intro) reset() {
	i.position = 0
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
	}

	return i.position >= len(i.content)
}

// draw the intro
func (i intro) draw(screen *ebiten.Image) {
	position := i.position
	if position >= len(i.content) {
		position = len(i.content) - 1
	}
	ebitenutil.DebugPrint(screen, i.content[i.position])
}
