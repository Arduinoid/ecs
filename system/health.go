package system

import (
	"ecs/component"
	"ecs/primitive"
)

// HealthSystem used to manage health based behavior on an entity
type HealthSystem struct {
	Entities []*healthEntity
}

// healthable used to check if an entity as the health component
type healthable interface {
	GetHealth() *component.HealthComponent
}

// healthEntity used to retain a reference to an entity that has a health
type healthEntity struct {
	id int
	*component.HealthComponent
}

func (h *HealthSystem) Update() {

}

// AddEntity will add an entity that has health to be processed by the system
func (h *HealthSystem) AddEntity(e primitive.Entity) {
	if he, ok := e.(healthable); ok {
		h.Entities = append(h.Entities, &healthEntity{
			id:              e.GetId(),
			HealthComponent: he.GetHealth(),
		},
		)
	}
}
