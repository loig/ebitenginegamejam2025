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
	"math"
	"math/rand/v2"

	"github.com/hajimehoshi/ebiten/v2"
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
	demonstrationSize      int
	maxCopsSize            int
	maxCopsPosition        contentPosition
	placedTiles            int
}

// build a fresh area
func buildPlayArea() playArea {
	var res playArea
	res.gridHasTile[globalDemonstrationStartY][globalDemonstrationStartX] = true
	res.grid[globalDemonstrationStartY][globalDemonstrationStartX] = tile{
		content:          [4]tileContent{contentDemonstration, contentDemonstration, contentDemonstration, contentDemonstration},
		contentGroupSize: [4]int{4, 4, 4, 4},
		neighboring:      [4]int{6, 12, 9, 3},
	}
	for quarter := 0; quarter < 4; quarter++ {
		for pos := 0; pos < len(res.grid[globalDemonstrationStartY][globalDemonstrationStartX].people[quarter]); pos++ {
			res.grid[globalDemonstrationStartY][globalDemonstrationStartX].people[quarter][pos] = rand.IntN(globalNumPeopleGraphics)
		}
	}
	res.deck = getDeck()
	for pos := 0; pos < globalHandSize; pos++ {
		res.drawNewTile(pos)
	}
	res.placedTiles = 1
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
	p.placedTiles++

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

	maxCopsPosition := p.maxCopsPosition

	// update the other contents
	for contentNumber := 0; contentNumber < 4; contentNumber++ {
		if p.grid[p.gridHoverY][p.gridHoverX].content[contentNumber] != contentDemonstration {
			position := contentPosition{tileX: p.gridHoverX, tileY: p.gridHoverY, contentNumber: contentNumber}
			p.updateContentFromPosition(position)
		}
	}

	// update the maxCops area
	if maxCopsPosition != p.maxCopsPosition {
		p.setAreaContent(maxCopsPosition, contentCop)
		p.setAreaContent(p.maxCopsPosition, contentManyCops)
	}

	// update the info on neighboring
	for contentNumber := 0; contentNumber < 4; contentNumber++ {
		position := contentPosition{tileX: p.gridHoverX, tileY: p.gridHoverY, contentNumber: contentNumber}
		p.updateNeighborsFromPosition(position)
	}

	// update the info on corners
	position := getAboveLeftPosition(contentPosition{tileX: p.gridHoverX, tileY: p.gridHoverY, contentNumber: 0})
	for line := 0; line < 4; line++ {
		currentPosition := position
		for cell := 0; cell < 4; cell++ {
			p.updateCorners(currentPosition)
			currentPosition = getRightPosition(currentPosition)
		}
		position = getBelowPosition(position)
	}

}

// check if a tile exists
func (p playArea) existsTile(position contentPosition) bool {
	return position.tileX >= 0 && position.tileX < globalGridWidth &&
		position.tileY >= 0 && position.tileY < globalGridHeight &&
		p.gridHasTile[position.tileY][position.tileX]
}

// get the content of a part of a tile
// the tile must exist
func (p playArea) getTileContent(position contentPosition) tileContent {
	return p.grid[position.tileY][position.tileX].content[position.contentNumber]
}

// check if D C B is an L shape of similar content
// A B
// D C
// C must exist
func (p playArea) isCorner(positionA, positionB, positionC, positionD contentPosition) bool {
	contentC := p.getTileContent(positionC)
	return p.existsTile(positionD) && p.getTileContent(positionD) == contentC &&
		p.existsTile(positionB) && p.getTileContent(positionB) == contentC &&
		(!p.existsTile(positionA) || p.getTileContent(positionA) != contentC)
}

// get positions relative to a given position
func getLeftPosition(position contentPosition) (res contentPosition) {

	res = position

	switch position.contentNumber {
	case 0:
		res.tileX--
		res.contentNumber = 1
	case 1:
		res.contentNumber = 0
	case 2:
		res.contentNumber = 3
	case 3:
		res.tileX--
		res.contentNumber = 2
	}

	return
}

func getRightPosition(position contentPosition) (res contentPosition) {

	res = position

	switch position.contentNumber {
	case 0:
		res.contentNumber = 1
	case 1:
		res.tileX++
		res.contentNumber = 0
	case 2:
		res.tileX++
		res.contentNumber = 3
	case 3:
		res.contentNumber = 2
	}

	return
}

func getAbovePosition(position contentPosition) (res contentPosition) {

	res = position

	switch position.contentNumber {
	case 0:
		res.tileY--
		res.contentNumber = 3
	case 1:
		res.tileY--
		res.contentNumber = 2
	case 2:
		res.contentNumber = 1
	case 3:
		res.contentNumber = 0
	}

	return
}

func getBelowPosition(position contentPosition) (res contentPosition) {

	res = position

	switch position.contentNumber {
	case 0:
		res.contentNumber = 3
	case 1:
		res.contentNumber = 2
	case 2:
		res.tileY++
		res.contentNumber = 1
	case 3:
		res.tileY++
		res.contentNumber = 0
	}

	return
}

func getAboveLeftPosition(position contentPosition) (res contentPosition) {

	res = position

	switch position.contentNumber {
	case 0:
		res.tileX--
		res.tileY--
		res.contentNumber = 2
	case 1:
		res.tileY--
		res.contentNumber = 3
	case 2:
		res.contentNumber = 0
	case 3:
		res.tileX--
		res.contentNumber = 1
	}

	return
}

func getAboveRightPosition(position contentPosition) (res contentPosition) {

	res = position

	switch position.contentNumber {
	case 0:
		res.tileY--
		res.contentNumber = 2
	case 1:
		res.tileY--
		res.tileX++
		res.contentNumber = 3
	case 2:
		res.tileX++
		res.contentNumber = 0
	case 3:
		res.contentNumber = 1
	}

	return
}

func getBelowRightPosition(position contentPosition) (res contentPosition) {

	res = position

	switch position.contentNumber {
	case 0:
		res.contentNumber = 2
	case 1:
		res.tileX++
		res.contentNumber = 3
	case 2:
		res.tileX++
		res.tileY++
		res.contentNumber = 0
	case 3:
		res.tileY++
		res.contentNumber = 1
	}

	return
}

func getBelowLeftPosition(position contentPosition) (res contentPosition) {

	res = position

	switch position.contentNumber {
	case 0:
		res.tileX--
		res.contentNumber = 2
	case 1:
		res.contentNumber = 3
	case 2:
		res.tileY++
		res.contentNumber = 0
	case 3:
		res.tileX--
		res.tileY++
		res.contentNumber = 1
	}

	return
}

// update corners info at a given position for drawing
func (p *playArea) updateCorners(position contentPosition) {
	if p.existsTile(position) {
		p.grid[position.tileY][position.tileX].corners[position.contentNumber][0] = p.isCorner(
			getAboveLeftPosition(position), getAbovePosition(position), position, getLeftPosition(position))
		p.grid[position.tileY][position.tileX].corners[position.contentNumber][1] = p.isCorner(
			getAboveRightPosition(position), getRightPosition(position), position, getAbovePosition(position))
		p.grid[position.tileY][position.tileX].corners[position.contentNumber][2] = p.isCorner(
			getBelowRightPosition(position), getBelowPosition(position), position, getRightPosition(position))
		p.grid[position.tileY][position.tileX].corners[position.contentNumber][3] = p.isCorner(
			getBelowLeftPosition(position), getLeftPosition(position), position, getBelowPosition(position))
	}

	//log.Print("Updated corners of ", position, " the result is ", p.grid[position.tileY][position.tileX].corners[position.contentNumber])
}

// update neighboring values for drawing
func (p *playArea) updateNeighborsFromPosition(position contentPosition) {
	neighbors := position.getNeighbors()
	for _, neighbor := range neighbors {
		if p.grid[position.tileY][position.tileX].content[position.contentNumber] == p.grid[neighbor.tileY][neighbor.tileX].content[neighbor.contentNumber] {
			p.grid[neighbor.tileY][neighbor.tileX].neighboring[neighbor.contentNumber] += neighborIncrement(position, neighbor)
			if position.tileX != neighbor.tileX || position.tileY != neighbor.tileY {
				p.grid[position.tileY][position.tileX].neighboring[position.contentNumber] += neighborIncrement(neighbor, position)
			}
		}
	}
}

// get the value that a neighbor brings to the neighboring of a contentPosition for drawing
func neighborIncrement(newNeighbor, position contentPosition) int {

	aboveIncrement := 1
	rightIncrement := 2
	belowIncrement := 4
	leftIncrement := 8

	if newNeighbor.tileX == position.tileX && newNeighbor.tileY == position.tileY {
		switch {
		case (newNeighbor.contentNumber == 0 && position.contentNumber == 1) ||
			(newNeighbor.contentNumber == 3 && position.contentNumber == 2):
			return leftIncrement
		case (newNeighbor.contentNumber == 1 && position.contentNumber == 0) ||
			(newNeighbor.contentNumber == 2 && position.contentNumber == 3):
			return rightIncrement
		case (newNeighbor.contentNumber == 0 && position.contentNumber == 3) ||
			(newNeighbor.contentNumber == 1 && position.contentNumber == 2):
			return aboveIncrement
		case (newNeighbor.contentNumber == 3 && position.contentNumber == 0) ||
			(newNeighbor.contentNumber == 2 && position.contentNumber == 1):
			return belowIncrement
		}
	}

	switch {
	case newNeighbor.tileX < position.tileX:
		return leftIncrement
	case newNeighbor.tileX > position.tileX:
		return rightIncrement
	case newNeighbor.tileY < position.tileY:
		return aboveIncrement
	case newNeighbor.tileY > position.tileY:
		return belowIncrement
	}

	return 0
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
		if p.grid[position.tileY][position.tileX].content[position.contentNumber] != contentDemonstration {
			for pos := 0; pos < len(p.grid[position.tileY][position.tileX].people[position.contentNumber]); pos++ {
				p.grid[position.tileY][position.tileX].people[position.contentNumber][pos] = rand.IntN(globalNumPeopleGraphics)
			}
		}
		p.grid[position.tileY][position.tileX].content[position.contentNumber] = contentDemonstration
		p.grid[position.tileY][position.tileX].contentGroupSize[position.contentNumber] = demonstrationSize
	}

	p.demonstrationSize = demonstrationSize
}

// check if two contents are similar for the sake of are size
func areSimiliarContents(contentA, contentB tileContent) bool {
	return contentA == contentB ||
		(contentA == contentCop && contentB == contentManyCops) ||
		(contentA == contentManyCops && contentB == contentCop)
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
				if areSimiliarContents(p.grid[n.tileY][n.tileX].content[n.contentNumber], content) {
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

	if content == contentCop && areaSize > p.maxCopsSize {
		p.maxCopsSize = areaSize
		p.maxCopsPosition = position
	}
}

// count the size of a type of content
func (p *playArea) setAreaContent(position contentPosition, newContent tileContent) {

	content := p.grid[position.tileY][position.tileX].content[position.contentNumber]

	nexts := []contentPosition{position}
	seen := map[contentPosition]bool{position: true}

	for len(nexts) > 0 {
		current := nexts[0]
		nexts = nexts[1:]

		neighbors := current.getNeighbors()
		for _, n := range neighbors {
			if !seen[n] && p.gridHasTile[n.tileY][n.tileX] {
				if areSimiliarContents(p.grid[n.tileY][n.tileX].content[n.contentNumber], content) {
					nexts = append(nexts, n)
					seen[n] = true
				}
			}
		}
	}

	for position := range seen {
		p.grid[position.tileY][position.tileX].content[position.contentNumber] = newContent
		switch newContent {
		case contentCop:
			p.grid[position.tileY][position.tileX].people[position.contentNumber] = copSet
		case contentManyCops:
			p.grid[position.tileY][position.tileX].people[position.contentNumber] = manyCopsSet
		}
	}
}

// check if the game is finished
func (p playArea) checkEndOfPlay() bool {
	handEmpty := !p.holdTile
	for pos := 0; pos < globalHandSize; pos++ {
		handEmpty = handEmpty && !p.handHasTile[pos]
	}
	return handEmpty || p.placedTiles >= globalGridWidth*globalGridHeight
}

/*
Content layout :
0 1
3 2
Same for layout for the corners of a content
*/
type tile struct {
	content          [4]tileContent
	contentGroupSize [4]int
	neighboring      [4]int
	corners          [4][4]bool
	people           [4][8]int
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

	// background
	opt := &ebiten.DrawImageOptions{}
	opt.GeoM.Translate(float64(tileX), float64(tileY))
	screen.DrawImage(tileImage, opt)

	// quarters
	// top left
	x, y := tileX+2, tileY+2
	drawTileQuarter(x, y, t.content[0], t.corners[0], t.neighboring[0], t.contentGroupSize[0], t.people[0], true, true, true, screen)
	// top right
	x += globalTileSize/2 - 4
	drawTileQuarter(x, y, t.content[1], t.corners[1], t.neighboring[1], t.contentGroupSize[1], t.people[1], false, false, true, screen)
	// bottom left
	x -= globalTileSize/2 - 4
	y += globalTileSize/2 - 4
	drawTileQuarter(x, y, t.content[3], t.corners[3], t.neighboring[3], t.contentGroupSize[3], t.people[3], false, true, false, screen)
	// bottom right
	x += globalTileSize/2 - 4
	drawTileQuarter(x, y, t.content[2], t.corners[2], t.neighboring[2], t.contentGroupSize[2], t.people[2], true, false, false, screen)
}

// draw one quarter of a tile
func drawTileQuarter(x, y int, content tileContent, corners [4]bool, neighboring int, groupSize int, people [8]int, isOdd bool, isLeft bool, isTop bool, screen *ebiten.Image) {
	if content == contentCop || content == contentPeople || content == contentDemonstration || content == contentManyCops {
		theImage := copsbackImage
		if content == contentPeople {
			theImage = peoplebackImage
		} else if content == contentDemonstration {
			theImage = protestbackImage
		} else if content == contentManyCops {
			theImage = manycopsbackImage
		}

		// background
		opt := &ebiten.DrawImageOptions{}
		opt.GeoM.Translate(float64(x), float64(y))
		imageX := neighboring * globalTileSize / 2
		screen.DrawImage(theImage.SubImage(image.Rect(imageX, 0, imageX+globalTileSize/2, globalTileSize/2)).(*ebiten.Image), opt)
		drawCorners(corners, x, y, screen, theImage)

		// people
		if groupSize <= 1 {
			opt = &ebiten.DrawImageOptions{}
			middleX := x - globalPeopleGraphicsWidth/2 + globalTileSize/4
			middleY := y - globalPeopleGraphicsHeight + 4 + globalTileSize/4
			opt.GeoM.Translate(float64(middleX), float64(middleY))
			peopleX := people[0] * globalPeopleGraphicsWidth
			screen.DrawImage(peopleImage.SubImage(image.Rect(peopleX, 0, peopleX+globalPeopleGraphicsWidth, globalPeopleGraphicsHeight)).(*ebiten.Image), opt)
		} else {

			numPeople := 0

			hasLeft := neighboring >= 8
			hasTop := neighboring%2 == 1
			//hasBottom := (neighboring/4)%2 == 1
			//hasRight := (neighboring/2)%2 == 1

			sepWidth := 16
			sepHeight := 13
			backMove := -(sepWidth + sepWidth/2)

			startX := x + sepWidth/4 + 1
			startY := y - globalPeopleGraphicsWidth
			if isLeft {
				startX -= 3
			} else {
				startX += 1
			}
			if isTop {
				startY -= 2
			} else {
				startY += 2
			}

			peopleX := func() int {
				res := people[numPeople] * globalPeopleGraphicsWidth
				numPeople++
				return res
			}

			if hasTop && hasLeft && !corners[0] {
				startXCorner := startX - sepWidth
				if isOdd {
					startXCorner += sepWidth / 2
				}
				startYCorner := startY - sepHeight
				opt = &ebiten.DrawImageOptions{}
				opt.GeoM.Translate(float64(startXCorner), float64(startYCorner))
				pX := peopleX()
				screen.DrawImage(peopleImage.SubImage(image.Rect(pX, 0, pX+globalPeopleGraphicsWidth, globalPeopleGraphicsHeight)).(*ebiten.Image), opt)
			}

			if hasTop {
				startXTop := startX + sepWidth/2
				startYTop := startY - sepHeight
				if !isOdd {
					startXTop -= sepWidth / 2
				}

				opt = &ebiten.DrawImageOptions{}
				opt.GeoM.Translate(float64(startXTop), float64(startYTop))
				pX := peopleX()
				screen.DrawImage(peopleImage.SubImage(image.Rect(pX, 0, pX+globalPeopleGraphicsWidth, globalPeopleGraphicsHeight)).(*ebiten.Image), opt)

				if !isOdd {
					opt.GeoM.Translate(float64(sepWidth), 0)
					pX := peopleX()
					screen.DrawImage(peopleImage.SubImage(image.Rect(pX, 0, pX+globalPeopleGraphicsWidth, globalPeopleGraphicsHeight)).(*ebiten.Image), opt)
				}
			}

			if !isOdd {
				startX += sepWidth / 2
			}
			if hasLeft {
				startX -= sepWidth
			}

			opt = &ebiten.DrawImageOptions{}
			opt.GeoM.Translate(float64(startX), float64(startY))
			pX := peopleX()
			screen.DrawImage(peopleImage.SubImage(image.Rect(pX, 0, pX+globalPeopleGraphicsWidth, globalPeopleGraphicsHeight)).(*ebiten.Image), opt)

			if hasLeft {
				opt.GeoM.Translate(float64(sepWidth), 0)
				pX := peopleX()
				screen.DrawImage(peopleImage.SubImage(image.Rect(pX, 0, pX+globalPeopleGraphicsWidth, globalPeopleGraphicsHeight)).(*ebiten.Image), opt)
			}

			opt.GeoM.Translate(float64(sepWidth), 0)
			if !isOdd {
				opt.GeoM.Translate(float64(backMove), float64(sepHeight))
				if hasLeft {
					opt.GeoM.Translate(-float64(sepWidth), 0)
				}
			}
			pX = peopleX()
			screen.DrawImage(peopleImage.SubImage(image.Rect(pX, 0, pX+globalPeopleGraphicsWidth, globalPeopleGraphicsHeight)).(*ebiten.Image), opt)

			opt.GeoM.Translate(float64(sepWidth), 0)
			if isOdd {
				opt.GeoM.Translate(float64(backMove), float64(sepHeight))
				if hasLeft {
					opt.GeoM.Translate(-float64(sepWidth), 0)
				}
			}
			pX = peopleX()
			screen.DrawImage(peopleImage.SubImage(image.Rect(pX, 0, pX+globalPeopleGraphicsWidth, globalPeopleGraphicsHeight)).(*ebiten.Image), opt)

			if hasLeft {
				opt.GeoM.Translate(float64(sepWidth), 0)
				pX := peopleX()
				screen.DrawImage(peopleImage.SubImage(image.Rect(pX, 0, pX+globalPeopleGraphicsWidth, globalPeopleGraphicsHeight)).(*ebiten.Image), opt)
			}
		}
	}
}

// draw the corners of a tile content if needed
func drawCorners(corners [4]bool, x, y int, screen *ebiten.Image, theImage *ebiten.Image) {
	for numCorner, isCorner := range corners {
		if isCorner {
			xShift, yShift := 0, 0
			if numCorner == 1 || numCorner == 2 {
				xShift = globalTileSize / 2
			}
			if numCorner == 2 || numCorner == 3 {
				yShift = globalTileSize / 2
			}

			opt := &ebiten.DrawImageOptions{}
			opt.GeoM.Rotate((float64(numCorner) * math.Pi / 2))
			opt.GeoM.Translate(float64(x+xShift), float64(y+yShift))
			screen.DrawImage(theImage.SubImage(image.Rect(16*globalTileSize/2, 0, 17*globalTileSize/2, globalTileSize/2)).(*ebiten.Image), opt)
		}
	}
}

type tileContent = int

const (
	contentCity int = iota
	contentCop
	contentNature
	contentPeople
	contentDemonstration
	contentManyCops
)
