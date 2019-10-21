package adventure

import (
	"eachtra/adventure/walk"
	"eachtra/components"
	"eachtra/utils"
	"errors"
	"sort"
	"strings"
)

type StartFunc func() (startingLocation components.Location)

type factory struct {
	name        string
	description string
	starting    components.Location
}

var adventures []factory

// any new adventures have to added here, a registration method called from within the created adventure causes a circular dependency
// by convention each adventure should export the following
// Name - string - unique name to lookup the adventure
// Description - string
// Entrance - Location - where the adventure starts
// TODO:
// Inventory - Container - what the user starts with
func init() {
	adventures = append(adventures, factory{
		name:        walk.Name,
		description: walk.Description,
		starting:    walk.Entrance,
	})
}

func List() (l []string) {
	for _, a := range adventures {
		l = append(l, a.name+" - "+a.description)
	}
	sort.Strings(l)
	return
}

func New(player, name string) (*Adventure, error) {
	for _, f := range adventures {
		if strings.EqualFold(f.name, name) {
			start := f.starting
			return &Adventure{
				Player:          player,
				Name:            f.name,
				Description:     f.description,
				currentLocation: start,
			}, nil
		}
	}
	return nil, errors.New("no adventures of that Name can be found")

}

// Adventure ties everything together and processes the turns
type Adventure struct {
	Player          string
	Name            string
	Description     string
	currentLocation components.Location
	history         []string
}

func (a *Adventure) Turn() int {
	return len(a.history)
}

func (a *Adventure) Location() components.Location {
	return a.currentLocation
}

func (a *Adventure) NewLocation(location components.Location) {
	a.currentLocation = location
}

func (a *Adventure) Execute(cmd string) (resp string, err error) {
	c := utils.NewStringQueueFromLine(strings.ToLower(cmd))
	desc, err := a.currentLocation.Execute(a, c)
	if err == nil {
		a.history = append(a.history, cmd)
	}
	return desc, err
}
