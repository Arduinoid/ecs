package world

import (
	"ecs/primitive"
	"time"
)

// Entity interface for entities to implement
type Entity interface {
	GetId() int
}

// System is used to identify various systems that can be added to a world
type System interface {
	Update()
	AddEntity(primitive.Entity)
}

// World is used as the main structure that glues together systems and running those on entities
type World struct {
	Systems []System
}

// NewWorld construct a new world for adding and running systems and entities
func NewWorld() *World {
	return &World{}
}

// Run will start an endless loop and step through all the systems
func (w *World) Run() {
	for {
		w.StepSystems()
		time.Sleep(1 * time.Second)
	}
}

// AddSystem will register a system to the world which can then be ran on entities
func (w *World) AddSystem(s System) {
	w.Systems = append(w.Systems, s)
}

// AddEntity will add a new entity to the world and add it to it's related systems
func (w *World) AddEntity(e primitive.Entity) {
	for _, s := range w.Systems {
		s.AddEntity(e)
	}
}

// StepSystems will run through each registered system and invoke that system's update method
func (w *World) StepSystems() {
	for _, s := range w.Systems {
		s.Update()
	}
}
