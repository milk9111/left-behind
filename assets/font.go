package assets

import (
	"bytes"

	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"golang.org/x/image/font/gofont/goregular"
)

var (
	FontGoregular *text.GoTextFaceSource
)

func init() {
	FontGoregular = mustLoadFont(goregular.TTF)
}

func mustLoadFont(b []byte) *text.GoTextFaceSource {
	s, err := text.NewGoTextFaceSource(bytes.NewReader(b))
	if err != nil {
		panic(err)
	}

	return s
}
