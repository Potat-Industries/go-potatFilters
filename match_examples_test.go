package potatFilters_test

import (
	"fmt"

	"github.com/Potat-Industries/go-potatFilters"
)

func ExampleTest_filterAll() {
	fmt.Println(potatFilters.Test("retarded", potatFilters.FilterAll))

	// Output: true
}

func ExampleTest_filterStrict() {
	fmt.Println(potatFilters.Test("retarded", potatFilters.FilterStrict))

	// Output: false
}

func ExampleTest_filterRemove() {
	fmt.Println(potatFilters.Test("retarded", potatFilters.FilterAll.Remove(potatFilters.FilterRacism)))

	// Output: true
}

func ExampleMatch_goodText() {
	fmt.Println(potatFilters.Match("good text", potatFilters.FilterAll) == 0)

	// Output: true
}

func ExampleMatch_badText() {
	fmt.Println(potatFilters.Match("women belong to the kitchen", potatFilters.FilterAll).Has(potatFilters.FilterSexism))

	// Output: true
}

func ExampleMatch_filterString() {
	fmt.Println(potatFilters.Match("nigg", potatFilters.FilterAll))

	// Output: racism-n-word
}
