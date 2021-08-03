package entity

// EID this variable holds the last generated id
var EID int = 0

// genID generates an id used for entities
func genID() int {
	newId := EID
	EID++
	return newId
}
