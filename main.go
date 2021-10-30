package main

import (
	"fmt"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	screenWidth  = 800
	screenHeight = 600

	targetTicksPerSecond = 60
)

var gameDelta float64;

func main() {
	fmt.Println("Starting...")
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		fmt.Println("initializing SDL:", err)
		return
	}

	window, err := sdl.CreateWindow(
		"Ragil2D",
		sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		screenWidth, screenHeight,
		sdl.WINDOW_OPENGL)
	if err != nil {
		fmt.Println("initializing window:", err)
		return
	}
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Println("initializing renderer:", err)
		return
	}
	defer renderer.Destroy()

	var elements []*element

	plr := newPlayer(renderer)
	elements = append(elements, plr)

	for i := 0; true; i++ {
		frameStartTime := time.Now()
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				return
			}
		}

		renderer.SetDrawColor(255, 255, 255, 255)
		renderer.Clear()

		for _, element := range(elements) {
			if element.active {
				err = element.update()
				if err != nil {
					fmt.Println("Error: updating element:", err)
					return
				}
				err = element.draw(renderer)
				if err != nil {
					fmt.Println("Error: drawing element:", err)
					return
				}
			}
		}
		renderer.Present()

		gameDelta = time.Since(frameStartTime).Seconds() * targetTicksPerSecond
	}
}
