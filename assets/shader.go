package assets

import (
	_ "embed"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	//go:embed shaders/outline.kage
	outline_kage []byte
	//go:embed shaders/glossy_gradient.kage
	glossyGradient_kage []byte
	//go:embed shaders/rainbow_flows.kage
	rainbowFlows_kage []byte
	//go:embed shaders/windy_grass.kage
	windyGrass_kage []byte
)

var (
	ShaderOutline        *ebiten.Shader
	ShaderGlossyGradient *ebiten.Shader
	ShaderRainbowFlows   *ebiten.Shader
	ShaderWindyGrass     *ebiten.Shader
)

func init() {
	ShaderOutline = mustShader(outline_kage)
	ShaderGlossyGradient = mustShader(glossyGradient_kage)
	ShaderRainbowFlows = mustShader(rainbowFlows_kage)
	ShaderWindyGrass = mustShader(windyGrass_kage)
}

func mustShader(b []byte) *ebiten.Shader {
	shader, err := ebiten.NewShader(b)
	if err != nil {
		panic(err)
	}

	return shader
}
