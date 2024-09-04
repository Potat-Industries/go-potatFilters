// PotatBot's chat filters, but in Go
package potatFilters

// Test checks if the string matches any of the filters
// Returns false if the string doesn't match any of the filters
func Test(input string, filters Filter) bool {
	regexps := filters.Regexp()

	if len(regexps) == 0 {
		return false
	}

	for _, r := range regexps {
		if r.MatchString(input) {
			return true
		}
	}

	return false
}

// Match applies the specified filters to the input string.
//
// Parameters:
// - input: The string to be filtered.
// - filters: The filters to be applied.
//
// Returns:
// - The filter that was matched, or 0 if no filter was matched.
func Match(input string, filters Filter) Filter {
	allFilters := filters.Decompose()

	for _, filter := range allFilters {
		if Test(input, filter) {
			return filter
		}
	}

	return 0
}
