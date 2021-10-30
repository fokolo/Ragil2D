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

func collides(c1, c2 circle) (bool, float64) {
	dist := math.Sqrt(math.Pow(c2.center.x-c1.center.x, 2) +
		math.Pow(c2.center.y-c1.center.y, 2))
	
	touchingDistance := c1.radius+c2.radius
	return dist <= touchingDistance, touchingDistance - dist
}

func (cl *collision) moveOutOfTheWay() {
	
}


func (cl *collision) onUpdate() error {
	for _, collisionCircle := range cl.container.collisionCircles {
		isColide, correctionNeeded := collides(cl.container.getCollisionCircle(), collisionCircle)
		if (isColide) {
			invertedVelocity := vector{x: -1 * cl.container.velocity.x, y: -1 * cl.container.velocity.y}
			cl.container.position.x += correctionNeeded / 2 // TODO: this is not really good need to fix the math
			cl.container.position.y += correctionNeeded / 2
			cl.container.setVelocity(invertedVelocity)
		}		
	}
	return nil
}
