package component

// TextComponent used to add text to an entity for display
type TextComponent struct {
	Font string
	Text string
	Size int
}

// GetText is the identity component for TextComponent
func (t *TextComponent) GetText() *TextComponent {
	return t
}
