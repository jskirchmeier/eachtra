package components

import (
	"eachtra/utils"
	"errors"
)

var ExitAdventure = errors.New("thank you for playing, please come again")

// Location defines a place the user can be in, they always have to be in a Location
// A location is also a container that can hold objects
// Introduction - a short description upon entering the new location
// Executes - executes the cmd upon the state, cmd is the
type Location interface {
	// Introduction describes the location, setting the scene
	Introduction() string

	// Execute acts upon the entered command, the results may change state or simply return information
	// returns a description of the results of the command, and the location where the user is
	// an error if the command could not be executed, the error EXIT_ADVENTURE terminates the game
	Execute(st State, cmd utils.StringQueue) (description string, err error)
}

// BaseLocation implements the basic requirements for a 'normal' location.
// Can be used to implement most locations.  More elaborate locations can extend this or
// ignore it all together
type BaseLocation struct {
	Commander

	Description string
}

func (b BaseLocation) Introduction() string {
	return b.Description
}
