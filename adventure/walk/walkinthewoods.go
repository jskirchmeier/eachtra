package walk

import (
	"eachtra/components"
	"eachtra/utils"
)

var (
	Name = "walk"
	Description = "a simple walk in the woods"
	Entrance = &components.BaseLocation{"at the end of the road, a thick ancient forest spreads out before you with a narrow path leading to the north", entranceFunc}
)

func entranceFunc(st components.State, cmd utils.StringQueue) (newState components.State, description string) {

	return st, "still here, no options are available"
}
