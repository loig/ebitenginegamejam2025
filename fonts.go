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
	"bytes"
	_ "embed"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

//go:embed WorkSans-Regular.ttf
var workSansRegular_ttf []byte
var workSansFaceSource *text.GoTextFaceSource
var workSansFace *text.GoTextFace

func loadFonts() {
	s, err := text.NewGoTextFaceSource(bytes.NewReader(workSansRegular_ttf))
	if err != nil {
		log.Fatal(err)
	}
	workSansFaceSource = s

	workSansFace = &text.GoTextFace{
		Source: workSansFaceSource,
		Size:   24,
	}
}

func drawTextAt(theText string, x, y int, screen *ebiten.Image) {
	op := &text.DrawOptions{}
	op.GeoM.Translate(float64(x), float64(y))
	op.LineSpacing = workSansFace.Size * 1.5
	text.Draw(screen, theText, workSansFace, op)
}

func drawTextCenteredAt(theText string, x, y float64, screen *ebiten.Image) (height float64) {
	op := &text.DrawOptions{}
	op.GeoM.Translate(x, y)
	op.LineSpacing = workSansFace.Size * 1.5
	op.PrimaryAlign = text.AlignCenter
	text.Draw(screen, theText, workSansFace, op)

	_, height = text.Measure(theText, workSansFace, workSansFace.Size*1.5)
	return
}
