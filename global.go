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
	globalTileSize = 40
	// grid characteristics
	globalGridWidth  = 7
	globalGridHeight = 9
	globalGridX      = 20
	globalGridY      = 20
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
	// screen size
	globalWidth  = 800
	globalHeight = 600
)
