package text

import (
	"../fonts"
	"github.com/engoengine/ecs"
	"github.com/engoengine/engo"
)

type Text struct {
	Text     string
	Size     float64
	Font     fonts.Font
	Scale    engo.Point
	Color    Color
	Position engo.Point

	control *TextControlSystem
	entity  *ecs.Entity
	render  *engo.RenderComponent
	space   *engo.SpaceComponent
}

func New(t Text) *Text {
	t.control = &TextControlSystem{}
	return &t
}

func (t *Text) Render() {
	font := &engo.Font{
		Size: t.Size,
		BG:   t.Color.BG,
		FG:   t.Color.FG,
		TTF:  fonts.Get(t.Font),
	}

	t.render = engo.NewRenderComponent(font.Render(t.Text), t.Scale, "text")
	t.entity.AddComponent(t.render)

	x, y, _ := font.TextDimensions(t.Text)
	t.entity.AddComponent(&engo.SpaceComponent{
		Position: t.Position,
		Width:    float32(x),
		Height:   float32(y),
	})
}

func (t *Text) Entity(w *ecs.World) *ecs.Entity {
	w.AddSystem(t.control)

	t.entity = ecs.NewEntity([]string{
		"RenderSystem",
		"MouseSystem",
		"TextControlSystem",
	})

	t.Render()
	t.entity.AddComponent(&engo.MouseComponent{})

	return t.entity
}

func (t *Text) OnClicked(fn func(entity *ecs.Entity, dt float32)) {
	t.control.Clicked = fn
}

func (t *Text) OnReleased(fn func(entity *ecs.Entity, dt float32)) {
	t.control.Released = fn
}

func (t *Text) OnHovered(fn func(entity *ecs.Entity, dt float32)) {
	t.control.Hovered = fn
}

func (t *Text) OnDragged(fn func(entity *ecs.Entity, dt float32)) {
	t.control.Dragged = fn
}

func (t *Text) OnRightClicked(fn func(entity *ecs.Entity, dt float32)) {
	t.control.RightClicked = fn
}

func (t *Text) OnRightReleased(fn func(entity *ecs.Entity, dt float32)) {
	t.control.RightReleased = fn
}

func (t *Text) OnEnter(fn func(entity *ecs.Entity, dt float32)) {
	t.control.Enter = fn
}

func (t *Text) OnLeave(fn func(entity *ecs.Entity, dt float32)) {
	t.control.Leave = fn
}
