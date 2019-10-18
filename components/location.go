package components

import (
	"eachtra/utils"
)

// Location defines a place the user can be in, they always have to be in a Location
// A location is also a container that can hold objects
// Introduction - a short description upon entering the new location
// Executes - executes the cmd upon the state, cmd is the
type Location interface {
	// Introduction describes the location, setting the scene
	Introduction() string

	// Execute acts upon the entered command, the results may change state or simply return information
	// returns the new state (which may be the same as the old) and a description of the results of the
	// command
	Execute(st State, cmd utils.StringQueue) (newState State, description string)
}
