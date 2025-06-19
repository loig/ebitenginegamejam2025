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

import "math/rand"

func genTileSet() (tileSet [globalNumTiles]tile) {
	// add cops
	for addedCops := 0; addedCops < globalNumCops; addedCops++ {
		tilePosition := rand.Intn(globalNumTiles)
		tileSide := rand.Intn(4)

		for tileSet[tilePosition].getContentAtSide(tileSide) != contentCity {
			tileSide++
			if tileSide >= 4 {
				tileSide = 0
				tilePosition = (tilePosition + 1) % globalNumTiles
			}
		}

		tileSet[tilePosition].setContentAtSide(tileSide, contentCop)
	}

	// add people
	for addedPeople := 0; addedPeople < globalNumPeople; addedPeople++ {
		tilePosition := rand.Intn(globalNumTiles)
		tileSide := rand.Intn(4)

		for tileSet[tilePosition].getContentAtSide(tileSide) != contentCity {
			tileSide++
			if tileSide >= 4 {
				tileSide = 0
				tilePosition = (tilePosition + 1) % globalNumTiles
			}
		}

		tileSet[tilePosition].setContentAtSide(tileSide, contentPeople)
	}

	// add nature
	for addedNature := 0; addedNature < globalNumNature; addedNature++ {
		tilePosition := rand.Intn(globalNumTiles)
		tileSide := rand.Intn(4)

		for tileSet[tilePosition].getContentAtSide(tileSide) != contentCity {
			tileSide++
			if tileSide >= 4 {
				tileSide = 0
				tilePosition = (tilePosition + 1) % globalNumTiles
			}
		}

		tileSet[tilePosition].setContentAtSide(tileSide, contentNature)
	}

	return
}

func getDeck() (deck [globalDeckSize]tile) {
	allTiles := genTileSet()
	for pos := 0; pos < len(deck); pos++ {
		deck[pos] = allTiles[pos]
	}
	return
}
