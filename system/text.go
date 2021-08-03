package system

import (
	"ecs/component"
	"ecs/primitive"
	"fmt"
)

// TextRenderSystem provides rendered text for entities
type TextRenderSystem struct {
	Entities []*textRenderEntity
}

// textRenderable identifies entities that can render text
type textRenderable interface {
	Entity
	GetText() *component.TextComponent
	GetPosition() *component.PositionComponent
}

// textRenderEntity used to reference an entity that has text render components
type textRenderEntity struct {
	id int
	*component.PositionComponent
	*component.TextComponent
}

// AddEntity will add entities that have text render components
func (t *TextRenderSystem) AddEntity(e primitive.Entity) {
	if tr, ok := e.(textRenderable); ok {
		t.Entities = append(t.Entities, &textRenderEntity{
			id:                tr.GetId(),
			PositionComponent: tr.GetPosition(),
			TextComponent:     tr.GetText(),
		})
	}
}

// Update will process text rendering on every related entity
func (t *TextRenderSystem) Update() {
	for _, text := range t.Entities {
		fmt.Printf("\n=== RENDERING: \"%s\" with font: \"%s\" at size: %d position X: %d, Y: %d ====\n\n", text.Text, text.Font, text.Size, text.X, text.Y)
	}
}
