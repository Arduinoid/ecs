package component

// HealthComponent is useful for giving an entity a way to track and affect health
type HealthComponent struct {
	HP int
}

// GetHealth is an identity method for HealthComponent
func (h *HealthComponent) GetHealth() *HealthComponent {
	return h
}
