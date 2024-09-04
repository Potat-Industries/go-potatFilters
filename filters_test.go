package potatFilters

import (
	"testing"
)

func TestFilterHasRegexp(t *testing.T) {
	filters := []Filter{
		FilterRacismGeneral,
		FilterRacismGeneralSlurs,
		FilterRacismRacoonWord,
		FilterRacismNWord,
		FilterRacismCWord,
		FilterRacismTWord,
		FilterRacismFWord,
		FilterRacismNonEnglishSlurs,
		FilterViolence,
		FilterSelfHarm,
		FilterSexualHarassment,
		FilterSexism,
		FilterAbleism,
		FilterAdvertising,
		FilterAgeTos,
	}
	all := FilterAll.Regexp()
	count := 0

	for _, filter := range filters {
		t.Run(filter.String(), func(t *testing.T) {
			regex := filter.Regexp()

			if len(regex) == 0 {
				t.Errorf("filter %s has no regexp", filter.String())
			}

			for _, r := range regex {
				if r == nil {
					t.Errorf("filter %s has nil regexp", filter.String())
				}
			}

			count += len(regex)
		})
	}

	if count != len(all) {
		t.Errorf("expected len(FilterAll.Regexp()) to be %d regexps, got %d", len(all), count)
	}
}
