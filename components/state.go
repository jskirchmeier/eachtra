package components

// State represents the current state of the system
// It also allows for modifications to that state
// it's represented as an interface witch will be implemented by the
// controlling structure (adventure), to eliminate circular dependencies
// Turn - current turn number - starting with 1
// Location - the current location of the user
type State interface {
	Turn() int

	Location() Location
	NewLocation(l Location)

	// TODO - container
	// TODO - inventory
}
