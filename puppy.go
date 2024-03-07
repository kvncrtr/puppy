package puppy

import (
	"github.com/kvncrtr/dog"
)

func Bark() string {
	return "Woof"
}

func Barks() string {
	return "WOOF! WOOF! WOOF!"
}

func BigBark() string {
	return dog.WhenGrownUp(Bark())
}

func BigBarks() string {
	return dog.WhenGrownUp(Barks())
}