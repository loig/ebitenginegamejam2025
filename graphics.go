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
	"image"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

//go:embed background.png
var backgroundBytes []byte
var backgroundImage *ebiten.Image

//go:embed tile.png
var tileBytes []byte
var tileImage *ebiten.Image

//go:embed peopleback.png
var peoplebackBytes []byte
var peoplebackImage *ebiten.Image

//go:embed copsback.png
var copsbackBytes []byte
var copsbackImage *ebiten.Image

//go:embed protestback.png
var protestbackBytes []byte
var protestbackImage *ebiten.Image

//go:embed manycopsback.png
var manycopsbackBytes []byte
var manycopsbackImage *ebiten.Image

//go:embed flags.png
var flagsBytes []byte
var flagsImage *ebiten.Image

//go:embed people.png
var peopleBytes []byte
var peopleImage *ebiten.Image

func loadGraphics() {
	decoded, _, err := image.Decode(bytes.NewReader(backgroundBytes))
	if err != nil {
		log.Fatal(err)
	}
	backgroundImage = ebiten.NewImageFromImage(decoded)

	decoded, _, err = image.Decode(bytes.NewReader(tileBytes))
	if err != nil {
		log.Fatal(err)
	}
	tileImage = ebiten.NewImageFromImage(decoded)

	decoded, _, err = image.Decode(bytes.NewReader(peoplebackBytes))
	if err != nil {
		log.Fatal(err)
	}
	peoplebackImage = ebiten.NewImageFromImage(decoded)

	decoded, _, err = image.Decode(bytes.NewReader(copsbackBytes))
	if err != nil {
		log.Fatal(err)
	}
	copsbackImage = ebiten.NewImageFromImage(decoded)

	decoded, _, err = image.Decode(bytes.NewReader(protestbackBytes))
	if err != nil {
		log.Fatal(err)
	}
	protestbackImage = ebiten.NewImageFromImage(decoded)

	decoded, _, err = image.Decode(bytes.NewReader(manycopsbackBytes))
	if err != nil {
		log.Fatal(err)
	}
	manycopsbackImage = ebiten.NewImageFromImage(decoded)

	decoded, _, err = image.Decode(bytes.NewReader(flagsBytes))
	if err != nil {
		log.Fatal(err)
	}
	flagsImage = ebiten.NewImageFromImage(decoded)

	decoded, _, err = image.Decode(bytes.NewReader(peopleBytes))
	if err != nil {
		log.Fatal(err)
	}
	peopleImage = ebiten.NewImageFromImage(decoded)
}
