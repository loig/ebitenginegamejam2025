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

var frenchCredits []string = []string{
	"On est là !",
	" ",
	"Un jeu réalisé (code, graphismes, musique) par Loïg Jezequel\nen deux semaines du 15 au 29 juin 2025 pour la jam annuelle\ndu moteur Ebitengine.",
	" ",
	"Le code source est disponible sous licence GPL-3.0 :\nhttps://github.com/loig/ebitenginegamejam2025",
	"Les musiques sont toutes plus ou moins librement inspirées\ndu chant de révolte «On est là».",
	"La police «Work Sans» utilisée est sous licence OFL-1.1 :\nhttps://github.com/weiweihuanghuang/Work-Sans",
	" ",
	" ",
	" ",
	" ",
	"Clique pour revenir au menu !",
}

var englishCredits []string = []string{
	"On est là !",
	" ",
	"A game made Ebitengine Game Jam 2025. Source code, gra-\nphics and music by Loïg Jezequel.",
	" ",
	"The source code is available under GPL-3.0 licence:\nhttps://github.com/loig/ebitenginegamejam2025",
	"All the music pieces are variations around the french protest\nsong \"On est là\".",
	"The \"Work Sans\" font is available under OFL-1.1 licence:\nhttps://github.com/weiweihuanghuang/Work-Sans",
	" ",
	" ",
	" ",
	" ",
	" ",
	"Click to go back to title screen!",
}

func drawCredits(screen *ebiten.Image) {

	credits := frenchCredits
	if language == englishLanguage {
		credits = englishCredits
	}

	y := 20.0
	for _, text := range credits {
		height := drawTextAt(text, 20.0, y, screen)
		y += height + 20
	}

}
