package component

// PositionComponent is a component useful for setting position in 2D space
type PositionComponent struct {
	X, Y int
}

// GetPosition is an identity method used to return itself. This helps
// systems to identify with which entities they should be concerned
func (p *PositionComponent) GetPosition() *PositionComponent {
	return p
}
