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
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type playArea struct {
	grid                   [globalGridHeight][globalGridWidth]tile
	gridHasTile            [globalGridHeight][globalGridWidth]bool
	gridHoverX, gridHoverY int
	gridHover              bool
	hand                   [globalHandSize]tile
	handHasTile            [globalHandSize]bool
	handHoverPos           int
	handHover              bool
	heldHandTile           int
	holdTile               bool
	holdX, holdY           int
	deck                   [globalDeckSize]tile
	deckPos                int
}

// build a fresh area
func buildPlayArea() playArea {
	var res playArea
	res.gridHasTile[globalDemonstrationStartY][globalDemonstrationStartX] = true
	res.grid[globalDemonstrationStartY][globalDemonstrationStartX] = tile{
		content:          [4]tileContent{contentDemonstration, contentDemonstration, contentDemonstration, contentDemonstration},
		contentGroupSize: [4]int{4, 4, 4, 4},
	}
	res.deck = getDeck()
	for pos := 0; pos < globalHandSize; pos++ {
		res.drawNewTile(pos)
	}
	return res
}

// update information about what is below the mouse
func (p *playArea) updateMousePosition(mouseX, mouseY int) {
	p.gridHoverX = (mouseX - globalGridX) / globalTileSize
	p.gridHoverY = (mouseY - globalGridY) / globalTileSize
	p.gridHover = (mouseX-globalGridX) >= 0 && (mouseY-globalGridY) >= 0 &&
		p.gridHoverX >= 0 && p.gridHoverX < globalGridWidth &&
		p.gridHoverY >= 0 && p.gridHoverY < globalGridHeight
	p.handHoverPos = (mouseY - globalHandY) / (globalTileSize + globalHandSep)
	p.handHover = p.handHoverPos >= 0 && p.handHoverPos < globalHandSize &&
		(mouseX-globalHandX) >= 0 && (mouseX-globalHandX)/globalTileSize == 0 &&
		(mouseY-globalHandY) >= 0 && (mouseY-globalHandY) < p.handHoverPos*(globalTileSize+globalHandSep)+globalTileSize

}

// check if a tile can be dropped at current mouse position
func (p playArea) canDropTile() bool {
	return p.holdTile &&
		p.gridHover && !p.gridHasTile[p.gridHoverY][p.gridHoverX] &&
		((p.gridHoverX-1 >= 0 && p.gridHasTile[p.gridHoverY][p.gridHoverX-1]) ||
			(p.gridHoverX+1 < globalGridWidth && p.gridHasTile[p.gridHoverY][p.gridHoverX+1]) ||
			(p.gridHoverY-1 >= 0 && p.gridHasTile[p.gridHoverY-1][p.gridHoverX]) ||
			(p.gridHoverY+1 < globalGridHeight && p.gridHasTile[p.gridHoverY+1][p.gridHoverX]))
}

func (p *playArea) drawNewTile(handPos int) {
	if p.deckPos < len(p.deck) {
		p.hand[handPos] = p.deck[p.deckPos]
		p.deckPos++
		p.handHasTile[handPos] = true
		return
	}
	p.handHasTile[handPos] = false
}

// drop a tile at the current mouse position
func (p *playArea) dropTile() {
	p.grid[p.gridHoverY][p.gridHoverX] = p.hand[p.heldHandTile]
	p.gridHasTile[p.gridHoverY][p.gridHoverX] = true

	// update the demonstration if needed
	demonstrationUpdateNeeded := false
	for contentNumber := 0; contentNumber < 4 && !demonstrationUpdateNeeded; contentNumber++ {
		position := contentPosition{tileX: p.gridHoverX, tileY: p.gridHoverY, contentNumber: contentNumber}
		neighbors := position.getNeighbors()
		for _, n := range neighbors {
			if p.gridHasTile[n.tileY][n.tileX] {
				if p.grid[n.tileY][n.tileX].content[n.contentNumber] == contentDemonstration {
					demonstrationUpdateNeeded = true
					break
				}
			}
		}
	}

	if demonstrationUpdateNeeded {
		p.updateDemonstration()
	}

	// update the other contents
	for contentNumber := 0; contentNumber < 4; contentNumber++ {
		if p.grid[p.gridHoverY][p.gridHoverX].content[contentNumber] != contentDemonstration {
			position := contentPosition{tileX: p.gridHoverX, tileY: p.gridHoverY, contentNumber: contentNumber}
			p.updateContentFromPosition(position)
		}
	}
}

// find which tile contents are part of the demonstration (and update its size)
// this is basically a search from the initial position of the demonstration
func (p *playArea) updateDemonstration() {

	nexts := []contentPosition{
		contentPosition{tileX: globalDemonstrationStartX, tileY: globalDemonstrationStartY, contentNumber: 0},
		contentPosition{tileX: globalDemonstrationStartX, tileY: globalDemonstrationStartY, contentNumber: 1},
		contentPosition{tileX: globalDemonstrationStartX, tileY: globalDemonstrationStartY, contentNumber: 2},
		contentPosition{tileX: globalDemonstrationStartX, tileY: globalDemonstrationStartY, contentNumber: 3},
	}

	seen := map[contentPosition]bool{
		contentPosition{tileX: globalDemonstrationStartX, tileY: globalDemonstrationStartY, contentNumber: 0}: true,
		contentPosition{tileX: globalDemonstrationStartX, tileY: globalDemonstrationStartY, contentNumber: 1}: true,
		contentPosition{tileX: globalDemonstrationStartX, tileY: globalDemonstrationStartY, contentNumber: 2}: true,
		contentPosition{tileX: globalDemonstrationStartX, tileY: globalDemonstrationStartY, contentNumber: 3}: true,
	}

	for len(nexts) > 0 {
		current := nexts[0]
		nexts = nexts[1:]

		neighbors := current.getNeighbors()
		for _, n := range neighbors {
			if !seen[n] && p.gridHasTile[n.tileY][n.tileX] {
				if p.grid[n.tileY][n.tileX].content[n.contentNumber] == contentDemonstration ||
					p.grid[n.tileY][n.tileX].content[n.contentNumber] == contentPeople {
					nexts = append(nexts, n)
					seen[n] = true
				}
			}
		}
	}

	demonstrationSize := len(seen)
	for position := range seen {
		p.grid[position.tileY][position.tileX].content[position.contentNumber] = contentDemonstration
		p.grid[position.tileY][position.tileX].contentGroupSize[position.contentNumber] = demonstrationSize
	}
}

// count the size of a type of content
func (p *playArea) updateContentFromPosition(position contentPosition) {

	content := p.grid[position.tileY][position.tileX].content[position.contentNumber]

	nexts := []contentPosition{position}
	seen := map[contentPosition]bool{position: true}

	for len(nexts) > 0 {
		current := nexts[0]
		nexts = nexts[1:]

		neighbors := current.getNeighbors()
		for _, n := range neighbors {
			if !seen[n] && p.gridHasTile[n.tileY][n.tileX] {
				if p.grid[n.tileY][n.tileX].content[n.contentNumber] == content {
					nexts = append(nexts, n)
					seen[n] = true
				}
			}
		}
	}

	areaSize := len(seen)
	for position := range seen {
		p.grid[position.tileY][position.tileX].contentGroupSize[position.contentNumber] = areaSize
	}
}

/*
Content layout :
0 1
3 2
*/
type tile struct {
	content          [4]tileContent
	contentGroupSize [4]int
}

type contentPosition struct {
	tileX, tileY  int
	contentNumber int
}

// find the neighbors of a contentPosition (a tile + the position in the layout of the tile)
func (c contentPosition) getNeighbors() []contentPosition {
	res := make([]contentPosition, 0)

	existsLeft := c.tileX-1 >= 0
	existsAbove := c.tileY-1 >= 0
	existsRight := c.tileX+1 < globalGridWidth
	existsBelow := c.tileY+1 < globalGridHeight

	switch c.contentNumber {
	case 0:
		res = append(res, contentPosition{tileX: c.tileX, tileY: c.tileY, contentNumber: 3})
		res = append(res, contentPosition{tileX: c.tileX, tileY: c.tileY, contentNumber: 1})
		if existsLeft {
			res = append(res, contentPosition{tileX: c.tileX - 1, tileY: c.tileY, contentNumber: 1})
		}
		if existsAbove {
			res = append(res, contentPosition{tileX: c.tileX, tileY: c.tileY - 1, contentNumber: 3})
		}
	case 1:
		res = append(res, contentPosition{tileX: c.tileX, tileY: c.tileY, contentNumber: 2})
		res = append(res, contentPosition{tileX: c.tileX, tileY: c.tileY, contentNumber: 0})
		if existsRight {
			res = append(res, contentPosition{tileX: c.tileX + 1, tileY: c.tileY, contentNumber: 0})
		}
		if existsAbove {
			res = append(res, contentPosition{tileX: c.tileX, tileY: c.tileY - 1, contentNumber: 2})
		}
	case 2:
		res = append(res, contentPosition{tileX: c.tileX, tileY: c.tileY, contentNumber: 1})
		res = append(res, contentPosition{tileX: c.tileX, tileY: c.tileY, contentNumber: 3})
		if existsRight {
			res = append(res, contentPosition{tileX: c.tileX + 1, tileY: c.tileY, contentNumber: 3})
		}
		if existsBelow {
			res = append(res, contentPosition{tileX: c.tileX, tileY: c.tileY + 1, contentNumber: 1})
		}
	case 3:
		res = append(res, contentPosition{tileX: c.tileX, tileY: c.tileY, contentNumber: 0})
		res = append(res, contentPosition{tileX: c.tileX, tileY: c.tileY, contentNumber: 2})
		if existsLeft {
			res = append(res, contentPosition{tileX: c.tileX - 1, tileY: c.tileY, contentNumber: 2})
		}
		if existsBelow {
			res = append(res, contentPosition{tileX: c.tileX, tileY: c.tileY + 1, contentNumber: 0})
		}
	}

	return res
}

func (t tile) draw(tileX, tileY int, screen *ebiten.Image) {
	//vector.DrawFilledRect(screen, float32(tileX), float32(tileY), float32(globalTileSize), float32(globalTileSize), color.RGBA{R: 150, G: 0, B: 0, A: 255}, false)

	contentColors := map[tileContent]color.Color{
		contentCity:          color.RGBA{R: 240, G: 230, B: 140, A: 255},
		contentCop:           color.RGBA{R: 65, G: 105, B: 225, A: 255},
		contentNature:        color.RGBA{R: 167, G: 214, B: 125, A: 255},
		contentPeople:        color.RGBA{R: 252, G: 142, B: 172, A: 255},
		contentDemonstration: color.RGBA{R: 255, G: 196, B: 12, A: 255},
	}

	// top left
	if t.contentGroupSize[0] > 1 {
		vector.DrawFilledRect(screen, float32(tileX), float32(tileY), float32(globalTileSize/2), float32(globalTileSize/2), contentColors[t.content[0]], false)
	} else {
		vector.DrawFilledRect(screen, float32(tileX+globalTileSize/8), float32(tileY+globalTileSize/8), float32(globalTileSize/4), float32(globalTileSize/4), contentColors[t.content[0]], false)
	}
	// top right
	if t.contentGroupSize[1] > 1 {
		vector.DrawFilledRect(screen, float32(tileX+globalTileSize/2), float32(tileY), float32(globalTileSize/2), float32(globalTileSize/2), contentColors[t.content[1]], false)
	} else {
		vector.DrawFilledRect(screen, float32(tileX+5*globalTileSize/8), float32(tileY+globalTileSize/8), float32(globalTileSize/4), float32(globalTileSize/4), contentColors[t.content[1]], false)
	}
	// bottom right
	if t.contentGroupSize[2] > 1 {
		vector.DrawFilledRect(screen, float32(tileX+globalTileSize/2), float32(tileY+globalTileSize/2), float32(globalTileSize/2), float32(globalTileSize/2), contentColors[t.content[2]], false)
	} else {
		vector.DrawFilledRect(screen, float32(tileX+5*globalTileSize/8), float32(tileY+5*globalTileSize/8), float32(globalTileSize/4), float32(globalTileSize/4), contentColors[t.content[2]], false)
	}
	// bottom left
	if t.contentGroupSize[3] > 1 {
		vector.DrawFilledRect(screen, float32(tileX), float32(tileY+globalTileSize/2), float32(globalTileSize/2), float32(globalTileSize/2), contentColors[t.content[3]], false)
	} else {
		vector.DrawFilledRect(screen, float32(tileX+globalTileSize/8), float32(tileY+5*globalTileSize/8), float32(globalTileSize/4), float32(globalTileSize/4), contentColors[t.content[3]], false)
	}
}

type tileContent = int

const (
	contentCity int = iota
	contentCop
	contentNature
	contentPeople
	contentDemonstration
)
