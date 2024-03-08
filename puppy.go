package puppy

import (
	"fmt"

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

func FromV1() {
	fmt.Println("I am from version 1")
}
