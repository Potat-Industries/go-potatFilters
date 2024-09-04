package potatFilters

import (
	"regexp"
	"strings"
)

type Filter int

const (
	FilterRacismGeneral Filter = 1 << iota
	FilterRacismGeneralSlurs
	FilterRacismRacoonWord
	FilterRacismNWord
	FilterRacismCWord
	FilterRacismTWord
	FilterRacismFWord
	FilterRacismNonEnglishSlurs
	FilterViolence
	FilterSelfHarm
	FilterSexualHarassment
	FilterSexism
	FilterAbleism
	FilterAdvertising
	FilterAgeTos
	filterEnd

	FilterRacism = FilterRacismGeneral | FilterRacismGeneralSlurs | FilterRacismRacoonWord | FilterRacismNWord | FilterRacismCWord | FilterRacismTWord | FilterRacismFWord | FilterRacismNonEnglishSlurs

	FilterStrict        = FilterRacism | FilterViolence | FilterSelfHarm | FilterSexualHarassment | FilterSexism | FilterAgeTos
	FilterStrictEnglish = FilterStrict &^ FilterRacismNonEnglishSlurs

	FilterAll = FilterStrict | FilterAdvertising | FilterAbleism
)

var filtersNames = map[Filter]string{
	FilterRacismGeneral:         "racism-general",
	FilterRacismGeneralSlurs:    "racism-slurs",
	FilterRacismRacoonWord:      "racism-racoon-word",
	FilterRacismNWord:           "racism-n-word",
	FilterRacismCWord:           "racism-c-word",
	FilterRacismTWord:           "racism-t-word",
	FilterRacismFWord:           "racism-f-word",
	FilterRacismNonEnglishSlurs: "racism-non-english-slurs",

	FilterViolence:         "violence",
	FilterSelfHarm:         "self-harm",
	FilterSexism:           "sexism",
	FilterSexualHarassment: "sexual-harassment",
	FilterAgeTos:           "age-tos",

	FilterAdvertising: "advertising",
	FilterAbleism:     "ableism",
}

// todo: merge regexps
var filtersRegexp = map[Filter][]*regexp.Regexp{
	FilterRacismGeneral:         {potatGeneralRacism1, potatGeneralRacism2},
	FilterRacismGeneralSlurs:    {potatGeneralSlurs},
	FilterRacismRacoonWord:      {potatRacoonWord},
	FilterRacismNWord:           {potatNWord},
	FilterRacismCWord:           {potatCWord},
	FilterRacismTWord:           {potatTWord},
	FilterRacismFWord:           {potatFWord},
	FilterRacismNonEnglishSlurs: {potatNonEnglishSlurs},

	FilterViolence: {potatViolence1, potatViolence2},

	FilterSelfHarm: {potatSelfHarm},
	FilterSexism:   {potatSexism},

	FilterSexualHarassment: {potatSexualHarassment1, potatSexualHarassment2},

	FilterAgeTos: {potatAgeTos},

	FilterAdvertising: {potatAdvertising},

	FilterAbleism: {potatAbleism},
}

func (f Filter) Has(filter Filter) bool {
	return f&filter == filter
}

func (f Filter) Add(filter Filter) Filter {
	return f | filter
}

func (f Filter) Remove(filter Filter) Filter {
	return f &^ filter
}
func (f Filter) Regexp() []*regexp.Regexp {
	reg, has := filtersRegexp[f]

	if has {
		return reg
	}

	reg = []*regexp.Regexp{}
	filters := f.Decompose()

	for _, filter := range filters {
		reg = append(reg, filtersRegexp[filter]...)
	}

	return reg
}

// breaks filter info into its components.
func (f Filter) Decompose() []Filter {
	var filters []Filter

	for i := Filter(1); i < filterEnd; i <<= 1 {
		if f.Has(i) {
			filters = append(filters, i)
		}
	}

	return filters
}

func (f Filter) String() string {
	name := filtersNames[f]

	if name != "" {
		return name
	}

	stringBuilder := strings.Builder{}
	filters := f.Decompose()

	for i, filter := range filters {
		if i > 0 {
			stringBuilder.WriteString(", ")
		}

		stringBuilder.WriteString(filter.String())
	}

	return stringBuilder.String()
}
