package potatFilters_test

import (
	"fmt"

	"github.com/Potat-Industries/go-potatFilters"
)

func ExampleReplaceConfusable() {
	fmt.Println(potatFilters.ReplaceConfusable("æŒǋ"))

	// Output: aeOENj
}

func ExampleReplaceConfusable_flashBacks() {
	fmt.Println(potatFilters.ReplaceConfusable("𝗻ɡɡ"))

	// Output: ngg
}
