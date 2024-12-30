package archetype

import (
	"github.com/milk9111/left-behind/assets/scripts"
	"github.com/milk9111/left-behind/component"
	"github.com/yohamta/donburi"
)

func NewStickyTranslation(w donburi.World) *donburi.Entry {
	e := w.Entry(w.Create(
		component.InputHandler,
		component.Start,
	))

	stickyTranslation := scripts.NewStickyTranslation()
	component.InputHandler.SetValue(e, component.InputHandlerData{
		Handler: stickyTranslation,
	})

	component.Start.SetValue(e, component.StartData{
		Handler: stickyTranslation,
	})

	return e
}
