package component

// VelocityComponent this component is useful for giving an entity momentum and direction
type VelocityComponent struct {
	X, Y int
}

// GetVelocity is the identity method for VelocityComponent
func (v *VelocityComponent) GetVelocity() *VelocityComponent {
	return v
}
