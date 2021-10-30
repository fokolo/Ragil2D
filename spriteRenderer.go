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

type spriteRenderer struct {
	texture  *sdl.Texture
	container *element
}

func newSpriteRenderer(container *element, renderer *sdl.Renderer, imgPath string) (*spriteRenderer, error) {
	sr := &spriteRenderer{container: container}

	img, err := img.Load(imgPath)
	if err != nil {
		return sr, fmt.Errorf("loading player sprite: %v", err)
	}
	defer img.Free()

	sr.texture, err = renderer.CreateTextureFromSurface(img)
	if err != nil {
		return sr, fmt.Errorf("creating player texture: %v", err)
	}

	return sr, nil
}

func (sr *spriteRenderer) onDraw(renderer *sdl.Renderer) error {
	width := spriteWidth/sr.container.scale
	height := spriteHeight/sr.container.scale

	renderer.Copy(sr.texture,
		&sdl.Rect{X: 0, Y: 0, W: spriteWidth, H: spriteHeight},
		&sdl.Rect{X: int32(sr.container.position.x) - (width / 2), Y: int32(sr.container.position.y) - (height / 2), W: width, H: height})
		
	return nil
}

func (sr *spriteRenderer) onUpdate() error {
	return nil
}
