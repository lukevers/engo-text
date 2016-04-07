package main

import (
	"./fonts"
	"./text"
	"github.com/engoengine/ecs"
	"github.com/engoengine/engo"
	"image/color"
	"log"
)

type Scene struct{}

func (s *Scene) Hide() {}
func (s *Scene) Show() {}
func (s *Scene) Type() string {
	return "Scene"
}
func (s *Scene) Preload() {}
func (s *Scene) Setup(w *ecs.World) {
	engo.SetBackground(color.Black)

	w.AddSystem(&engo.MouseSystem{})
	w.AddSystem(&engo.RenderSystem{})

	button1 := text.New(text.Text{
		Text:  "button1",
		Size:  25,
		Font:  fonts.FONT_PRIMARY,
		Scale: engo.Point{1, 1},
		Color: text.Color{
			BG: color.Transparent,
			FG: color.White,
		},
		Position: engo.Point{25, 25},
	})

	button1.OnEnter(func(entity *ecs.Entity, dt float32) {
		engo.SetCursor(engo.Hand)
		log.Println("x")
	})

	button1.OnLeave(func(entity *ecs.Entity, dt float32) {
		engo.SetCursor(nil)
	})

	w.AddEntity(button1.Entity(w))

	button2 := text.New(text.Text{
		Text:  "button2",
		Size:  25,
		Font:  fonts.FONT_PRIMARY,
		Scale: engo.Point{1, 1},
		Color: text.Color{
			BG: color.Transparent,
			FG: color.White,
		},
		Position: engo.Point{25, 100},
	})

	button2.OnEnter(func(entity *ecs.Entity, dt float32) {
		engo.SetCursor(engo.Hand)
		log.Println("y")
	})

	button2.OnLeave(func(entity *ecs.Entity, dt float32) {
		engo.SetCursor(nil)
	})

	w.AddEntity(button2.Entity(w))

}
