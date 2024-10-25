package potatFilters

import (
	"testing"
)

func TestTestGoodText(t *testing.T) {
	goodTextKappa := []string{
		"test",
		"ryanpotat",
		"fart",
		"potat",
		"potatfart",
		"ermm",
		"twitch",
		"good text",
		"this text is a nothingburger",
		"you truck potat into the bed with the power of a v8",
		"i enjoyed my stay xqcL",
		"01jb0ncqqk7nvx448xhped0676",
		"01ffryc79r0007p57xyw0bjdjq",
	}

	for _, text := range goodTextKappa {
		if Test(text, FilterAll) {
			t.Errorf("Test(%s) = true, want false", text)
		}
	}
}

func TestBadText(t *testing.T) {
	badText := []string{
		"women are objects",
		"retarded",
		"kys",
		"𝙆𝚈s",
		"kill your self",
		"blackface",
		"drink bleach",
		"i'm 12 years old",
		"ոiɡɡаs",
		"tura",
		"szuwak",
		"pidor",
		"nickeurs",
		"nickers",
		"behead your self",
		"bhead",
		"pedo",
		"amerykaniec",
		"kitajec",
	}

	for _, text := range badText {
		if !Test(text, FilterAll) {
			t.Errorf("Test(%s) = false, want true", text)
		}
	}
}
