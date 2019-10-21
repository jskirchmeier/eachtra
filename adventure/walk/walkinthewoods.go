package walk

import (
	"eachtra/components"
)

var (
	Name        = "walk"
	Description = "a simple walk in the woods"
	Entrance    = components.BaseLocation{
		Description: "You are standing at the end of the road, a thick ancient forest spreads out before you with a narrow path leading to the north",
	}
)

func init() {
	Entrance.Add(components.ExitCommand)
}
