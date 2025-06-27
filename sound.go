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
	"io"
	"log"
	"math"
	"time"

	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/mp3"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
)

//go:embed title.mp3
var titleMusicBytes []byte
var titleMusic *audio.InfiniteLoop
var titleMusicPlayer *audio.Player

//go:embed intro.mp3
var introMusicBytes []byte
var introMusic *audio.InfiniteLoop
var introMusicPlayer *audio.Player

//go:embed theme.mp3
var themeMusicBytes []byte
var themeMusic *audio.InfiniteLoop
var themeMusicPlayer *audio.Player

//go:embed select.wav
var soundSelectBytes []byte
var soundSelect []byte

//go:embed mvt.wav
var soundMvtBytes []byte
var soundMvt []byte

//go:embed drop.wav
var soundDropBytes []byte
var soundDrop []byte

//go:embed achievement.wav
var soundAchievementBytes []byte
var soundAchievement []byte

const (
	soundSelectID int = iota
	soundMvtID
	soundDropID
	soundAchievementID
	numSounds
)

type soundManager struct {
	audioContext *audio.Context
	NextSounds   [numSounds]bool
	musicPlayer  *audio.Player
}

// play requested sounds
func (s *soundManager) PlaySounds() {
	for sound, play := range s.NextSounds {
		if play {
			s.playSound(sound)
			s.NextSounds[sound] = false
		}
	}
}

// play a sound
func (s soundManager) playSound(sound int) {
	var soundBytes []byte
	switch sound {
	case soundSelectID:
		soundBytes = soundSelect
	case soundMvtID:
		soundBytes = soundMvt
	case soundDropID:
		soundBytes = soundDrop
	case soundAchievementID:
		soundBytes = soundAchievement
	}

	if len(soundBytes) > 0 {
		soundPlayer := s.audioContext.NewPlayerFromBytes(soundBytes)
		soundPlayer.SetVolume(0.15)
		soundPlayer.Play()
	}
}

// loop the music
func (s *soundManager) UpdateMusic(volume float64) {
	if s.musicPlayer != nil {
		if !s.musicPlayer.IsPlaying() {
			s.musicPlayer.Rewind()
			s.musicPlayer.Play()
		}
		s.musicPlayer.SetVolume(volume)
	}
}

// stop the music
func (s *soundManager) StopMusic() {
	if s.musicPlayer != nil {
		s.musicPlayer.Pause()
	}
}

// set the music
type musicTrack = int

const (
	titleMusicTrack int = iota
	introMusicTrack
	themeMusicTrack
)

func (s *soundManager) ChangeMusic(track musicTrack) {
	s.StopMusic()
	switch track {
	case titleMusicTrack:
		s.musicPlayer = titleMusicPlayer
	case introMusicTrack:
		s.musicPlayer = introMusicPlayer
	case themeMusicTrack:
		s.musicPlayer = themeMusicPlayer
	}
}

// load all audio assets
func InitAudio() (manager soundManager) {

	var error error
	manager.audioContext = audio.NewContext(44100)

	// music
	musicmp3, error := mp3.DecodeWithSampleRate(manager.audioContext.SampleRate(), bytes.NewReader(titleMusicBytes))
	if error != nil {
		log.Panic("Audio problem:", error)
	}
	tduration, _ := time.ParseDuration("18s")
	duration := tduration.Seconds()
	introBytes := int64(math.Round(duration * 4 * float64(44100)))
	tduration, _ = time.ParseDuration("18s")
	duration = tduration.Seconds()
	mainBytes := int64(math.Round(duration * 4 * float64(44100)))
	titleMusic = audio.NewInfiniteLoopWithIntro(musicmp3, introBytes, mainBytes)
	titleMusicPlayer, error = manager.audioContext.NewPlayer(titleMusic)
	if error != nil {
		log.Panic("Audio problem:", error)
	}

	musicmp3, error = mp3.DecodeWithSampleRate(manager.audioContext.SampleRate(), bytes.NewReader(introMusicBytes))
	if error != nil {
		log.Panic("Audio problem:", error)
	}
	tduration, _ = time.ParseDuration("4s")
	duration = tduration.Seconds()
	introBytes = int64(math.Round(duration * 4 * float64(44100)))
	tduration, _ = time.ParseDuration("32s")
	duration = tduration.Seconds()
	mainBytes = int64(math.Round(duration * 4 * float64(44100)))
	introMusic = audio.NewInfiniteLoopWithIntro(musicmp3, introBytes, mainBytes)
	introMusicPlayer, error = manager.audioContext.NewPlayer(introMusic)
	if error != nil {
		log.Panic("Audio problem:", error)
	}

	musicmp3, error = mp3.DecodeWithSampleRate(manager.audioContext.SampleRate(), bytes.NewReader(themeMusicBytes))
	if error != nil {
		log.Panic("Audio problem:", error)
	}
	tduration, _ = time.ParseDuration("20s")
	duration = tduration.Seconds()
	introBytes = int64(math.Round(duration * 4 * float64(44100)))
	tduration, _ = time.ParseDuration("32s")
	duration = tduration.Seconds()
	mainBytes = int64(math.Round(duration * 4 * float64(44100)))
	themeMusic = audio.NewInfiniteLoopWithIntro(musicmp3, introBytes, mainBytes)
	themeMusicPlayer, error = manager.audioContext.NewPlayer(themeMusic)
	if error != nil {
		log.Panic("Audio problem:", error)
	}

	// sounds
	sound, error := wav.DecodeWithSampleRate(manager.audioContext.SampleRate(), bytes.NewReader(soundSelectBytes))
	if error != nil {
		log.Panic("Audio problem:", error)
	}
	soundSelect, error = io.ReadAll(sound)
	if error != nil {
		log.Panic("Audio problem:", error)
	}

	sound, error = wav.DecodeWithSampleRate(manager.audioContext.SampleRate(), bytes.NewReader(soundMvtBytes))
	if error != nil {
		log.Panic("Audio problem:", error)
	}
	soundMvt, error = io.ReadAll(sound)
	if error != nil {
		log.Panic("Audio problem:", error)
	}

	sound, error = wav.DecodeWithSampleRate(manager.audioContext.SampleRate(), bytes.NewReader(soundDropBytes))
	if error != nil {
		log.Panic("Audio problem:", error)
	}
	soundDrop, error = io.ReadAll(sound)
	if error != nil {
		log.Panic("Audio problem:", error)
	}

	sound, error = wav.DecodeWithSampleRate(manager.audioContext.SampleRate(), bytes.NewReader(soundAchievementBytes))
	if error != nil {
		log.Panic("Audio problem:", error)
	}
	soundAchievement, error = io.ReadAll(sound)
	if error != nil {
		log.Panic("Audio problem:", error)
	}
	return
}
