package assets

import (
	_ "embed"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	//go:embed shaders/outline.kage
	outline_kage []byte
)

var (
	ShaderOutline *ebiten.Shader
)

func init() {
	ShaderOutline = mustShader(outline_kage)
}

func mustShader(b []byte) *ebiten.Shader {
	shader, err := ebiten.NewShader(b)
	if err != nil {
		panic(err)
	}

	return shader
}
