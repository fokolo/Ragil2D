package main

import (
	"math"

	"github.com/veandco/go-sdl2/sdl"
)

type circle struct {
	center vector
	radius float64
}

type collision struct {
	container *element
}

func newCollision(container *element) (*collision, error) {
	cl := &collision{container: container}

	return cl, nil
}

func (*collision) onDraw(renderer *sdl.Renderer) error {
	return nil
}

func collides(c1, c2 circle) bool {
	dist := math.Sqrt(math.Pow(c2.center.x-c1.center.x, 2) +
		math.Pow(c2.center.y-c1.center.y, 2))

	return dist <= c1.radius+c2.radius
}


func (cl *collision) onUpdate() error {
	for _, collisionCircle := range cl.container.collisionCircles {
		if (collides(cl.container.getCollisionCircle(), collisionCircle)) {
			invertedVelocity := vector{x: -5 * cl.container.velocity.x, y: -5 * cl.container.velocity.y}
			cl.container.addVelocity(invertedVelocity)
		}		
	}
	return nil
}
