package assets

import (
	"bytes"
	_ "embed"

	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"golang.org/x/image/font/gofont/goregular"
)

var (
	//go:embed fonts/Roboto-VariableFont_wdth,wght.ttf
	roboto_ttf []byte
)

var (
	FontGoregular *text.GoTextFaceSource
	FontRoboto    *text.GoTextFaceSource
)

func init() {
	FontGoregular = mustLoadFont(goregular.TTF)
	FontRoboto = mustLoadFont(roboto_ttf)
}

func mustLoadFont(b []byte) *text.GoTextFaceSource {
	s, err := text.NewGoTextFaceSource(bytes.NewReader(b))
	if err != nil {
		panic(err)
	}

	return s
}
