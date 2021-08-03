package primitive

// Entity interface for entities to implement
type Entity interface {
	GetId() int
}

// Vec2 is useful for representing 2D data such as a point or velocity
type Vec2 struct {
	X, Y int
}
