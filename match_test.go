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
		"ඞ",
		"im 20",
		"i'm 61",
		"3f9dfa4",
		"faq",
		"fdafagd",
	}

	for _, text := range goodTextKappa {
		if Test(text, FilterAll) {
			t.Errorf("Test(%s) = true, want false", text)
		}
	}
}

func TestBadText(t *testing.T) {
	badText := []string{
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
		"im actually 7",
		"im 8",
		"i'm actually 4",
		"im 12",
		"i'm 12",
		"i'm seven",
		"im thiRtEEn",
		"im actually under 14",
		"i'm actually below 12",
		"im actually underage",
		"pidor",
		"tranny",
		"women are objects",
		"women are nothing more than objects",
		"niger",
		"niga",
		"nigga",
		"nigger",
		"faggot",
		"rape",
		"RA͒P͒E͒D͒",
		"fa4",
		"faqq",
		"fag", // i mean we have to test it
		"faqq",
		"helloniggas :)",
		"negger",
		"nogger",
		"n1gger",
		"n111111gg3r",
	}

	for _, text := range badText {
		if !Test(text, FilterAll) {
			t.Errorf("Test(%s) = false, want true", text)
		}
	}
}
