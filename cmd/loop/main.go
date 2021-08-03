package main

import (
	"ecs/entity"
	"ecs/primitive"
	"ecs/system"
	"ecs/world"
	"fmt"
)

func main() {
	fmt.Println("ECS in Go")
	// init the world
	w := world.NewWorld()

	// add systems
	w.AddSystem(&system.MovementSystem{})
	w.AddSystem(&system.TextRenderSystem{})
	w.AddSystem(&system.HealthSystem{})

	// create entities
	b := entity.NewBall(
		primitive.Vec2{X: 0, Y: 0},
		primitive.Vec2{X: 5, Y: 30},
		100,
	)
	w.AddEntity(b)

	// a text entity
	t := entity.NewTitle(
		primitive.Vec2{X: 0, Y: 0},
		primitive.Vec2{X: 2, Y: 5},
		"Hello, World!", "Sans Sarif", 20)
	w.AddEntity(t)

	// run the systems
	w.Run()
}
