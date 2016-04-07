package main

import (
	_ "./fonts"
	"github.com/engoengine/engo"
)

func main() {
	options := engo.RunOptions{
		Title:         "Test",
		Fullscreen:    false,
		Width:         1280,
		Height:        780,
		VSync:         false,
		ScaleOnResize: true,
	}

	engo.Run(options, &Scene{})
}
