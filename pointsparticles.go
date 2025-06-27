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

	"github.com/hajimehoshi/ebiten/v2"
)

type pointsParticle struct {
	x, y  int
	value int
	frame int
}

type particleSet struct {
	content   []pointsParticle
	firstDead int
}

func (p *particleSet) reset() {
	p.firstDead = 0
}

func (p *particleSet) addParticle(x, y int, value int) {
	newParticle := pointsParticle{x: x, y: y, value: value, frame: 0}
	if p.firstDead < len(p.content) {
		p.content[p.firstDead] = newParticle
	} else {
		p.content = append(p.content, newParticle)
	}
	p.firstDead++
}

func (p *particleSet) update() {
	for pos := 0; pos < p.firstDead; pos++ {
		if !p.content[pos].update() {
			p.firstDead--
			if p.firstDead > pos {
				p.content[pos], p.content[p.firstDead] = p.content[p.firstDead], p.content[pos]
				pos--
			}
		}
	}
}

func (p *pointsParticle) update() bool {
	p.frame++
	p.y--
	return p.frame < 60
}

func (p particleSet) draw(screen *ebiten.Image) {
	for pos := 0; pos < p.firstDead; pos++ {
		p.content[pos].draw(screen)
	}
}

func (p pointsParticle) draw(screen *ebiten.Image) {
	prefix := ""
	if p.value > 0 {
		prefix = "+"
	}
	drawSmallTextCenteredAt(fmt.Sprintf("%s%d", prefix, p.value), float64(p.x), float64(p.y), screen)
}
