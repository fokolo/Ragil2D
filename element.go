package main

import (
	"fmt"
	"reflect"

	"github.com/veandco/go-sdl2/sdl"
)

type element struct {
	scale int32
	x int32
	y int32
	active bool
	components []component
}

type component interface {
	onUpdate() error
	onDraw(renderer *sdl.Renderer) error
}

func (elem *element) draw(renderer *sdl.Renderer) error {
	for _, comp := range elem.components {
		err := comp.onDraw(renderer)
		if err != nil {
			return err
		}
	}

	return nil
}

func (elem *element) update() error {
	for _, comp := range elem.components {
		err := comp.onUpdate()
		if err != nil {
			return err
		}
	}

	return nil
}

func (elem *element) addComponent(new component) {
	for _, comp := range elem.components {
		if (reflect.TypeOf(comp) == reflect.TypeOf(new)) {
			panic(fmt.Sprintf(
				"attempt to add new component with existing type %v",
				reflect.TypeOf(new)))
		}
	}
	elem.components = append(elem.components, new)
}


func (elem *element) getComponent(withType component) component {
	for _, comp := range elem.components {
		if (reflect.TypeOf(comp) == reflect.TypeOf(withType)) {
			return comp
		}
	}

	panic(fmt.Sprintf("no component with type %v", reflect.TypeOf(withType)))
}