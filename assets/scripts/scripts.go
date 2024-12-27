package scripts

import (
	"fmt"

	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/filter"
)

type notFoundError struct {
	scriptError
}

func newNotFoundError(ct string) notFoundError {
	return notFoundError{
		scriptError: newScriptError(fmt.Sprintf("component type '%s' not found", ct)),
	}
}

func MustFindComponent[T any](w donburi.World, c *donburi.ComponentType[T]) *T {
	e, ok := donburi.NewQuery(filter.Contains(c)).First(w)
	if !ok {
		panic(newNotFoundError(c.String()))
	}

	return c.Get(e)
}

func MustFindComponents[T any](w donburi.World, c *donburi.ComponentType[T]) []*T {
	var t []*T
	donburi.NewQuery(filter.Contains(c)).Each(w, func(e *donburi.Entry) {
		t = append(t, c.Get(e))
	})

	return t
}

func MustFindEntries(w donburi.World, c donburi.IComponentType) []*donburi.Entry {
	var t []*donburi.Entry
	donburi.NewQuery(filter.Contains(c)).Each(w, func(e *donburi.Entry) {
		t = append(t, e)
	})

	return t
}

func MustFindEntry(w donburi.World, c donburi.IComponentType) *donburi.Entry {
	t, ok := donburi.NewQuery(filter.Contains(c)).First(w)
	if !ok {
		panic(newNotFoundError(c.Name()))
	}

	return t
}
