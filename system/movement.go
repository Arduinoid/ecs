package system

import (
	"ecs/component"
	"ecs/primitive"
	"fmt"
)

// moveable is used to identify entities that can be affected by the MovementSystem
type moveable interface {
	primitive.Entity
	GetPosition() *component.PositionComponent
	GetVelocity() *component.VelocityComponent
}

// MovementSystem will provide entities with movement behavior
type MovementSystem struct {
	Entities []*movementEntity
}

// movementEntity is used to represent a Moveable Entity
type movementEntity struct {
	id int
	*component.PositionComponent
	*component.VelocityComponent
}

// AddEntity will store a reference to an entity that has movement components
func (m *MovementSystem) AddEntity(e primitive.Entity) {
	if t, ok := e.(moveable); ok {
		m.Entities = append(m.Entities, &movementEntity{
			id:                t.GetId(),
			PositionComponent: t.GetPosition(),
			VelocityComponent: t.GetVelocity(),
		})
	}
}

// Update will should be called every step and will update all related entities
func (m *MovementSystem) Update() {
	for _, e := range m.Entities {
		fmt.Printf("Entity: %d, Start POS X: %d, Y: %d\n", e.id, e.PositionComponent.X, e.PositionComponent.Y)
		e.PositionComponent.X += e.VelocityComponent.X
		e.PositionComponent.Y += e.VelocityComponent.Y
		fmt.Printf("Entity: %d, End POS X: %d, Y: %d\n", e.id, e.PositionComponent.X, e.PositionComponent.Y)
	}
}
