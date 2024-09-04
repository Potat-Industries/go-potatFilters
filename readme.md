[PotatBotat](https://potat.app/)'s chat filters, but in Go.

![Ermm](https://cdn.7tv.app/emote/637b3fba1d4a267ec1cd3364/4x.avif)

[pkg.go.dev](https://pkg.go.dev/github.com/Potat-Industries/go-potatFilters)

## install

```sh
go get github.com/Potat-Industries/go-potatFilters
```

## Filters

This package includes the following filters:

- `FilterRacismGeneral` 
- `FilterRacismGeneralSlurs`
- `FilterRacismRacoonWord`
- `FilterRacismNWord`
- `FilterRacismCWord`
- `FilterRacismTWord`
- `FilterRacismFWord`
- `FilterRacismNonEnglishSlurs`
- `FilterViolence`
- `FilterSelfHarm`
- `FilterSexualHarassment`
- `FilterSexism`
- `FilterAbleism`
- `FilterAdvertising`
- `FilterAgeTos`
- `filterEnd`

and the following companions:

- `FilterRacism`: FilterRacismGeneral | FilterRacismGeneralSlurs | FilterRacismRacoonWord | FilterRacismNWord | FilterRacismCWord | FilterRacismTWord | FilterRacismFWord | FilterRacismNonEnglishSlurs
- `FilterStrict`:  FilterRacism | FilterViolence | FilterSelfHarm | FilterSexualHarassment | FilterSexism | FilterAgeTos
- `FilterStrictEnglish`: FilterStrict &^ FilterRacismNonEnglishSlurs
- `FilterAll`: FilterStrict | FilterAdvertising | FilterAbleism

Most of the filters are not perfect, and there may be some false positives. But it should be good enough for most Twitch chats.

For most cases, you can use the `FilterStrict` or `FilterStrictEnglish`, You can also mix filters using `Filter.add(AnotherFilter)`;

## examples

### Test: A simple CLI program that checks if a message is "bad"

```go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Potat-Industries/go-potatFilters"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Enter text: ")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)

		if potatFilters.Test(text, potatFilters.FilterStrict) {
			fmt.Println("Bad text detected BAND")

			continue
		}

		fmt.Println("Good")
	}
}
```

### Match: Identifying why the message is considered 'bad.'

```go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Potat-Industries/go-potatFilters"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Enter text: ")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)

		filter := potatFilters.Match(text, potatFilters.FilterAll) // FilterStrict is usually what you want

		if filter == 0 {
			fmt.Printf("Good Text: %s\n", text)

			continue
		}

		fmt.Printf(
			"Bad Text: %s, filter: %s, Is Racism: %t, Regexps Count: %d\n",
			text,
			filter,
			potatFilters.FilterRacism.Has(filter),
			len(filter.Regexp()),
		)
	}
}
```
example output:
```
Enter text: kys
Bad Text: kys, filter: violence, Is Racism: false, Regexps Count: 2
Enter text: retarded
Bad Text: retarded, filter: ableism, Is Racism: false, Regexps Count: 1
Enter text: im 10 years old
Bad Text: im 10 years old, filter: age-tos, Is Racism: false, Regexps Count: 1
Enter text: Good Text
Good Text: Good Text
```

## Example ReplaceConfusable
Removes accents and replaces confusable characters with their English equivalents.

```go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Potat-Industries/go-potatFilters"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Enter text: ")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)

		fmt.Printf("Clean text: %s\n", potatFilters.ReplaceConfusable(text))
	}
}
```
example output:
```
Enter text: 𝑨𝔄ᗄ𝖠𝗔ꓯ𝞐🄐🄰Ꭿ𐊠𝕬𝜜𝐴ꓮᎪ
Clean text: AAAAAAAAAAAAAAAA
Enter text: 𝐃𝑫𝕯𝖣𝔇𝘿ꭰⅅ𝒟ꓓ⑯⒃⒗ᴁ'𝑨𝔄ᗄ𝖠𝗔ꓯ𝞐🄐🄰Ꭿ𐊠𝕬𝜜𝐴ꓮᎪ31&$#
Clean text: DDDDDDDDDD161616AE'AAAAAAAAAAAAAAAA31&$#
```

