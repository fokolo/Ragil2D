package main

import (
	"fmt"
	"reflect"

	"github.com/veandco/go-sdl2/gfx"
	"github.com/veandco/go-sdl2/sdl"
)

type vector struct {
	x float64
	y float64
}

type element struct {
	scale int32
	position vector
	radius float64
	velocity vector
	active bool
	components []component
	collisionCircles []circle
}

type component interface {
	onUpdate() error
	onDraw(renderer *sdl.Renderer) error
}

func (elem *element) addVelocity(velocity vector)  {
	elem.velocity.x += gameDelta * velocity.x
	elem.velocity.y += gameDelta * velocity.y
}

func (elem *element) setVelocity(velocity vector)  {
	elem.velocity.x = velocity.x
	elem.velocity.y = velocity.y
}

func drawCircle(renderer *sdl.Renderer, draw circle) error {
	gfx.CircleRGBA(renderer, int32(draw.center.x), int32(draw.center.y), int32(draw.radius), 255, 0, 0, 255)

	return nil
}


func (elem *element) draw(renderer *sdl.Renderer) error {
	for _, comp := range elem.components {
		err := comp.onDraw(renderer)
		if err != nil {
			return err
		}
	}

	drawCircle(renderer, elem.getCollisionCircle())
	for _, colCircle := range elem.collisionCircles {
		drawCircle(renderer, colCircle)	
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

func (elem *element) getCollisionCircle() circle {
	return circle{center: elem.position, radius: elem.radius}
}


func (elem *element) getComponent(withType component) component {
	for _, comp := range elem.components {
		if (reflect.TypeOf(comp) == reflect.TypeOf(withType)) {
			return comp
		}
	}

	panic(fmt.Sprintf("no component with type %v", reflect.TypeOf(withType)))
}