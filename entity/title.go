package entity

import (
	"ecs/component"
	"ecs/primitive"
)

// title is an entity composed of components
type title struct {
	ID int
	*component.TextComponent
	*component.PositionComponent
	*component.VelocityComponent
}

// NewTitle will construct a new Title entity
func NewTitle(pos, vel primitive.Vec2, text, font string, size int) *title {
	return &title{
		ID:                genID(),
		PositionComponent: &component.PositionComponent{X: pos.X, Y: pos.Y},
		VelocityComponent: &component.VelocityComponent{X: vel.X, Y: vel.Y},
		TextComponent: &component.TextComponent{
			Font: font,
			Text: text,
			Size: size,
		},
	}
}

// GetId returns the entity id and implements Entity
func (t *title) GetId() int {
	return t.ID
}
