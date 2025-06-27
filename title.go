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
	"fmt"
	"image"
	"math/rand/v2"

	"github.com/hajimehoshi/ebiten/v2"
)

type titleButton struct {
	x, y           int
	width, height  int
	frenchContent  string
	englishContent string
	frenchImage    *ebiten.Image
	englishImage   *ebiten.Image
	imageScale     float64
	hover          bool
	effect         func(*game)
}

// Definition of the buttons available at title screen
var titleButtons []titleButton
var characters []int

// load images for title
func loadTitleButtons() {
	titleButtons = []titleButton{
		titleButton{
			x: (globalWidth - globalTitleButtonWidth) / 2, y: 700, width: globalTitleButtonWidth, height: globalTitleButtonHeight,
			frenchContent:  "Crédits",
			englishContent: "Credits",
			effect:         func(g *game) { g.state = stateCredits },
		},
		titleButton{
			x: (globalWidth - globalTitleButtonWidth) / 2, y: 475,
			width: globalTitleButtonWidth, height: globalTitleButtonHeight,
			frenchContent:  "Comment jouer",
			englishContent: "How to play",
			effect:         func(g *game) { g.state = stateHowTo },
		},
		titleButton{
			x: (globalWidth - globalTitleButtonWidth) / 2, y: 550,
			width: globalTitleButtonWidth, height: globalTitleButtonHeight,
			frenchContent:  "Trophées",
			englishContent: "Achievements",
			effect:         func(g *game) { g.state = stateAchievements },
		},
		titleButton{
			x: (globalWidth - globalTitleButtonWidth) / 2, y: 400,
			width: globalTitleButtonWidth, height: globalTitleButtonHeight,
			frenchContent:  "Jouer",
			englishContent: "Play",
			effect:         func(g *game) { g.state = statePlay },
		},
		titleButton{
			x: (globalWidth - globalTitleButtonWidth) / 2, y: 625,
			width: globalTitleButtonWidth, height: globalTitleButtonHeight,
			frenchContent:  "Revoir l'introduction",
			englishContent: "See the introduction",
			effect:         func(g *game) { g.state = stateIntro },
		},
		titleButton{
			x: globalWidth - globalFlagWidth/4 - 10, y: 10,
			width: globalFlagWidth / 4, height: globalFlagHeight / 4,
			englishImage: flagsImage.SubImage(image.Rect(0, 0, globalFlagWidth, globalFlagHeight)).(*ebiten.Image),
			frenchImage:  flagsImage.SubImage(image.Rect(globalFlagWidth, 0, 2*globalFlagWidth, globalFlagHeight)).(*ebiten.Image),
			imageScale:   0.25,
			effect:       func(g *game) { toggleLanguage() },
		},
	}
}

// reset the title screen
func resetTitle() {
	for buttonPosition := range titleButtons {
		titleButtons[buttonPosition].hover = false
	}
	generateCharacters(contentDemonstration)
}

// check button selection
func (g *game) updateTitle(mouseX, mouseY int) {
	for buttonPosition, button := range titleButtons {
		if mouseX >= button.x && mouseX < button.x+button.width &&
			mouseY >= button.y && mouseY < button.y+button.height {
			if !titleButtons[buttonPosition].hover {
				g.SoundManager.NextSounds[soundMvtID] = true
			}
			titleButtons[buttonPosition].hover = true
			if inputSelect() {
				button.effect(g)
			}
		} else {
			titleButtons[buttonPosition].hover = false
		}
	}
}

// generate the characters for display
func generateCharacters(content tileContent) {
	scaleFactor := 4
	lineNum := 0
	characterPosition := 0
	for y := -globalPeopleGraphicsHeight * scaleFactor; y < globalHeight; y += (globalPeopleGraphicsHeight * scaleFactor) / 4 {
		lineNum++
		for x := (lineNum%2)*(globalPeopleGraphicsWidth*scaleFactor)/2 - globalPeopleGraphicsWidth*scaleFactor; x < globalWidth; x += globalPeopleGraphicsWidth * scaleFactor {
			newCharacter := rand.IntN(globalNumPeopleGraphics)
			switch content {
			case contentCop:
				newCharacter = globalNumPeopleGraphics + 2
			case contentManyCops:
				newCharacter = globalNumPeopleGraphics + 1
			case contentPeople:
				newCharacter = globalNumPeopleGraphics
			}
			if len(characters) <= characterPosition {
				characters = append(characters, newCharacter)
			} else {
				characters[characterPosition] = newCharacter
			}
			characterPosition++
		}
	}
}

// draw the characters for display
func drawCharacters(screen *ebiten.Image, scale float32) {
	scaleFactor := 4
	lineNum := 0
	characterPosition := 0
	for y := -globalPeopleGraphicsHeight * scaleFactor; y < globalHeight; y += (globalPeopleGraphicsHeight * scaleFactor) / 4 {
		lineNum++
		for x := (lineNum%2)*(globalPeopleGraphicsWidth*scaleFactor)/2 - globalPeopleGraphicsWidth*scaleFactor; x < globalWidth; x += globalPeopleGraphicsWidth * scaleFactor {
			theImage := peopleImage.SubImage(image.Rect(
				characters[characterPosition]*globalPeopleGraphicsWidth, 0,
				(characters[characterPosition]+1)*globalPeopleGraphicsWidth, globalPeopleGraphicsHeight)).(*ebiten.Image)
			opt := &ebiten.DrawImageOptions{}
			opt.GeoM.Scale(float64(scaleFactor), float64(scaleFactor))
			opt.GeoM.Translate(float64(x), float64(y))
			opt.ColorScale.Scale(scale, scale, scale, 1)
			screen.DrawImage(theImage, opt)
			characterPosition++
		}
	}
}

func (g game) drawTitle(screen *ebiten.Image) {

	// draw the demonstration
	drawCharacters(screen, 1)

	// draw the title
	opt := &ebiten.DrawImageOptions{}
	opt.GeoM.Translate(
		float64(globalWidth-globalTitleWidth)/2,
		float64(globalTitleY),
	)
	screen.DrawImage(titleImage, opt)

	// draw the buttons
	for _, button := range titleButtons {
		if button.frenchImage == nil {
			theText := button.frenchContent
			if language == englishLanguage {
				theText = button.englishContent
			}
			//vector.DrawFilledRect(screen, float32(button.x), float32(button.y), float32(button.width), float32(button.height), theColor, false)
			opt := &ebiten.DrawImageOptions{}
			if !button.hover {
				opt.ColorScale.Scale(0.6, 0.6, 0.6, 1)
			}
			opt.GeoM.Translate(float64(button.x), float64(button.y))
			screen.DrawImage(buttonbackImage, opt)

			drawTextCenteredAt(theText, float64(button.x+button.width/2), float64(button.y+10), screen)
			//ebitenutil.DebugPrintAt(screen, theText, button.x+10, button.y+5)
		} else {
			theImage := button.frenchImage
			if language == englishLanguage {
				theImage = button.englishImage
			}
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Scale(button.imageScale, button.imageScale)
			op.GeoM.Translate(float64(button.x), float64(button.y))
			screen.DrawImage(theImage, op)
		}
	}

	// draw the score
	if g.score.max > 0 {
		x := float64(globalWidth-globalTitleButtonWidth) / 2
		y := 250.0
		opt := &ebiten.DrawImageOptions{}
		opt.GeoM.Translate(x, y)
		screen.DrawImage(buttonbackImage, opt)

		textInfo := "Meilleur score :"
		if language == englishLanguage {
			textInfo = "Best score:"
		}
		text := fmt.Sprintf("%s %d", textInfo, g.score.max)
		drawTextCenteredAt(text, float64(globalWidth/2), y+10.0, screen)
	}

}
