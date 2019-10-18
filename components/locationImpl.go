package components

import (
	"eachtra/utils"
)

type BaseLocation struct {
	Description string
	Executor    func(st State, cmd utils.StringQueue) (newState State, description string)
}

func (b *BaseLocation) Introduction() string {
	return b.Description
}

func (b *BaseLocation) Execute(st State, cmd utils.StringQueue) (newState State, description string) {
	if b.Executor == nil {
		return st, "Still here \n" + b.Description
	}
	return b.Executor(st, cmd)
}
