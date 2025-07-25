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
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

var language int = frenchLanguage

const (
	frenchLanguage int = iota
	englishLanguage
)

func languageSelectUpdate(mouseX int) (finished, changed bool) {

	oldLanguage := language

	language = frenchLanguage
	if mouseX >= globalWidth/2 {
		language = englishLanguage
	}

	return inputSelect(), oldLanguage != language
}

func languageSelectDraw(screen *ebiten.Image) {

	drawTextCenteredAt("Choisis ta langue/Choose your language", float64(globalWidth/2), 100, screen)

	flagX := globalWidth/4 - globalFlagWidth/2
	flagY := (globalHeight - globalFlagHeight) / 2
	opt := &ebiten.DrawImageOptions{}
	if language != frenchLanguage {
		opt.GeoM.Scale(0.5, 0.5)
		flagX = globalWidth/4 - globalFlagWidth/4
		flagY = globalHeight/2 - globalFlagHeight/4
	}
	opt.GeoM.Translate(float64(flagX), float64(flagY))
	screen.DrawImage(flagsImage.SubImage(image.Rect(0, 0, globalFlagWidth, globalFlagHeight)).(*ebiten.Image), opt)

	flagX = 3*globalWidth/4 - globalFlagWidth/2
	flagY = (globalHeight - globalFlagHeight) / 2
	opt = &ebiten.DrawImageOptions{}
	if language != englishLanguage {
		opt.GeoM.Scale(0.5, 0.5)
		flagX = 3*globalWidth/4 - globalFlagWidth/4
		flagY = globalHeight/2 - globalFlagHeight/4
	}
	opt.GeoM.Translate(float64(flagX), float64(flagY))
	screen.DrawImage(flagsImage.SubImage(image.Rect(globalFlagWidth, 0, 2*globalFlagWidth, globalFlagHeight)).(*ebiten.Image), opt)

}

func toggleLanguage() {
	switch language {
	case frenchLanguage:
		language = englishLanguage
	case englishLanguage:
		language = frenchLanguage
	}
}
