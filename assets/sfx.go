package assets

import (
	"bytes"
	_ "embed"
	"io"

	"github.com/hajimehoshi/ebiten/v2/audio/wav"
)

var (
	//go:embed sfx/bad_move.wav
	badMove_wav []byte
	//go:embed sfx/goal_reached.wav
	goalReached_wav []byte
)

var (
	SFXBadMove     AudioClip
	SFXGoalReached AudioClip
)

func init() {
	SFXBadMove = mustAudioClip(badMove_wav)
	SFXGoalReached = mustAudioClip(goalReached_wav)
}

type AudioClip []byte

func mustAudioClip(b []byte) []byte {
	m, err := wav.DecodeF32(bytes.NewBuffer(b))
	if err != nil {
		panic(err)
	}

	mb, err := io.ReadAll(m)
	if err != nil {
		panic(err)
	}

	return mb
}
