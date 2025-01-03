package scripts

import (
	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/milk9111/left-behind/event"
	"github.com/yohamta/donburi"
)

type PauseMenu struct {
	widget *widget.Widget

	paused bool
}

func NewPauseMenu(widget *widget.Widget) *PauseMenu {
	return &PauseMenu{
		widget: widget,
	}
}

func (p *PauseMenu) Start(w donburi.World) {
	p.widget.Visibility = widget.Visibility_Hide_Blocking
}

func (p *PauseMenu) Update(w donburi.World) {
	if !inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		return
	}

	p.paused = !p.paused

	if p.paused {
		p.widget.Visibility = widget.Visibility_Show
		event.PausedGame.Publish(w, event.PausedGameData{})
	} else {
		p.widget.Visibility = widget.Visibility_Hide_Blocking
		event.UnpausedGame.Publish(w, event.UnpausedGameData{})
	}
}

func (p *PauseMenu) OnClick(w donburi.World) {
	event.ExitedGame.Publish(w, event.ExitedGameData{})
}
