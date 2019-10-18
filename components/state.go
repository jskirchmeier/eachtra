package components

// State represents the current state of the system
// It also allows for modifications to that state
// Turn - current turn number - starting with 1
// Location - the current location of the user
type State struct {
	Turn int
	Location Location
	// TODO - container
	// TODO - inventory
}
