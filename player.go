package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	spriteWidth = 319
	spriteHeight = 486
)

type player struct {
	tex *sdl.Texture
	scale int32
	x int32
	y int32
}


func newPlayer(renderer *sdl.Renderer) (p player, err error) {
	img, err := img.Load("sprites/TempleRun/Idle__000.png")
	if err != nil {
		return player{}, fmt.Errorf("loading player sprite: %v", err)
	}
	defer img.Free()
	p.tex, err = renderer.CreateTextureFromSurface(img)
	if err != nil {
		return player{}, fmt.Errorf("creating player texture: %v", err)
	}

	p.scale = 5

	return p, nil
}

func (p *player) setLocation(x int32, y int32) {
	p.x = x
	p.y = y
}

func (p *player) draw(renderer *sdl.Renderer) {
	renderer.Copy(p.tex,
		&sdl.Rect{X: 0, Y: 0, W: spriteWidth, H: spriteHeight},
		&sdl.Rect{X: p.x, Y: p.y, W: spriteWidth/p.scale, H: spriteHeight/p.scale})
}
