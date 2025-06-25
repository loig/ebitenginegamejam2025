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
	globalGridWidth           = 7 //7
	globalGridHeight          = 7 //9
	globalGridX               = 40
	globalGridY               = 100
	globalDemonstrationStartX = globalGridWidth / 2
	globalDemonstrationStartY = globalGridHeight / 2
	// hand characteristics
	globalHandSize = 3
	globalHandX    = globalGridX + globalGridWidth*globalTileSize + 80
	globalHandY    = 180
	globalHandSep  = 80
	// deck characteristics
	globalDeckSize = globalGridWidth*globalGridHeight - 1
	// tile set characteristics
	globalNumCops   = 50
	globalNumPeople = 80
	globalNumNature = 0
	globalNumTiles  = 70
	// score characteristics
	globalPlayScoreX = globalGridX
	globalPlayScoreY = globalGridY + globalGridHeight*globalTileSize + 40 + globalTimeHeight
	// time characteristics
	globalAllowedTime = 180
	globalTimeX       = globalGridX
	globalTimeY       = globalGridY + globalGridHeight*globalTileSize + 20
	globalTimeHeight  = 40
	globalTimeWidth   = globalGridWidth * globalTileSize
	// achievements characteristics
	globalLiveDisplayTimeAchievement  = 120
	globalLiveDisplayAchievementsX    = globalHandX
	globalLiveDisplayAchievementsY    = globalHeight - 20
	globalLiveDisplayAchievementsSize = 20
	globalLiveDisplayAchievementsSep  = 10
	// title screen characteristics
	globalTitleMaxScoreX = 20
	globalTitleMaxScoreY = globalHeight - 20
	// end screen characteristics
	globalEndDisplayAchievementTime = 120
	globalEndDisplayAchievementX    = globalWidth/2 - 100
	globalEndDisplayAchievementY    = globalHeight / 2
	globalEndScoreX                 = globalWidth/2 - 100
	globalEndScoreY                 = 40
	globalEndMaxScoreX              = globalWidth/2 - 100
	globalEndMaxScoreY              = 60
	// intro characteristics
	globalIntroY    = 100.0
	globalIntroSep  = 50.0
	globalIntroTime = 300
	// language selection characteristics
	globalFlagWidth  = 99
	globalFlagHeight = 80
	// people characteristics
	globalNumPeopleGraphics    = 15
	globalPeopleGraphicsWidth  = 16
	globalPeopleGraphicsHeight = 32
	// screen size
	globalWidth  = 800
	globalHeight = 800
)
