package entity

import (
	"ecs/component"
	"ecs/primitive"
)

// ball is an entity composed of components
type ball struct {
	ID int
	*component.PositionComponent
	*component.VelocityComponent
	*component.HealthComponent
}

// NewBall will construct a new Ball entity
func NewBall(pos, vel primitive.Vec2, h int) *ball {
	return &ball{
		ID: genID(),
		PositionComponent: &component.PositionComponent{
			X: pos.X,
			Y: pos.Y,
		},
		VelocityComponent: &component.VelocityComponent{
			X: vel.X,
			Y: vel.Y,
		},
		HealthComponent: &component.HealthComponent{
			HP: h,
		},
	}
}

// GetId returns the entity id which implements Entity
func (b *ball) GetId() int {
	return b.ID
}
