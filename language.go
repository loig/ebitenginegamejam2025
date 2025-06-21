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

var language int = frenchLanguage

const (
	frenchLanguage int = iota
	englishLanguage
)

func languageSelectUpdate(mouseX int) (finished bool) {

	language = frenchLanguage
	if mouseX >= globalWidth/2 {
		language = englishLanguage
	}

	return inputSelect()
}

func languageSelectDraw(screen *ebiten.Image) {

	switch language {
	case frenchLanguage:
		ebitenutil.DebugPrint(screen, "Français")
	case englishLanguage:
		ebitenutil.DebugPrint(screen, "English")
	}

}

func toggleLanguage() {
	switch language {
	case frenchLanguage:
		language = englishLanguage
	case englishLanguage:
		language = frenchLanguage
	}
}
