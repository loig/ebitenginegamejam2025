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

const (
	// tiles characteristics
	globalTileSize = 80
	// grid characteristics
	globalGridWidth           = 5 //7
	globalGridHeight          = 5 //9
	globalGridX               = 20
	globalGridY               = 20
	globalDemonstrationStartX = globalGridWidth / 2
	globalDemonstrationStartY = globalGridHeight / 2
	// hand characteristics
	globalHandSize = 3
	globalHandX    = globalGridX + globalGridWidth*globalTileSize + 20
	globalHandY    = globalGridY
	globalHandSep  = 20
	// deck characteristics
	globalDeckSize = globalGridWidth*globalGridHeight - 1
	// tile set characteristics
	globalNumCops   = 20
	globalNumPeople = 80
	globalNumNature = 20
	globalNumTiles  = 70
	// score characteristics
	globalScoreX = globalGridX
	globalScoreY = globalGridY + globalGridHeight*globalTileSize + 40 + globalTimeHeight
	// time characteristics
	globalAllowedTime = 180
	globalTimeX       = globalGridX
	globalTimeY       = globalGridY + globalGridHeight*globalTileSize + 20
	globalTimeHeight  = 40
	globalTimeWidth   = globalGridWidth * globalTileSize
	// screen size
	globalWidth  = 800
	globalHeight = 600
)
