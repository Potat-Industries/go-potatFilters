package potatFilters

import (
	"regexp"
	"strconv"
	"testing"
	"unicode/utf8"
)

func TestConfusableSpaces(t *testing.T) {
	testCases := []struct {
		original string
		expected string
	}{
		{"\t\n\r\f ", "     "},
		{"\u2420\u2422\u2423", "   "},
		{"\u00a0\u1680\u2000\u2001\u2002\u2003\u2004\u2005\u2006\u2007\u2008\u2009\u200a\u202f\u205f\u3000", "                "},
	}

	for _, tc := range testCases {
		actual := ReplaceConfusable(tc.original)
		if actual != tc.expected {
			t.Errorf("replaceConfusable(%s) = %s, want %s", strconv.Quote(tc.original), strconv.Quote(actual), strconv.Quote(tc.expected))
		}
	}
}

func TestConfusableRemovesZeroWidth(t *testing.T) {
	testCases := []struct {
		original string
		expected string
	}{
		{"\u180E", ""},
		{"\u200b", ""},
		{"\u200c", ""},
		{"\u200d", ""},
		{"\u2060", ""},
		{"\uFEFF", ""},
	}

	for _, tc := range testCases {
		actual := ReplaceConfusable(tc.original)
		if actual != tc.expected {
			t.Errorf("replaceConfusable(%s) = %s, want %s", strconv.Quote(tc.original), strconv.Quote(actual), strconv.Quote(tc.expected))
		}
	}
}

func TestConfusableRemovesAccents(t *testing.T) {
	testCases := []struct {
		original string
		expected string
	}{
		{"\u0300\u0301\u0302\u0303\u0304\u0305\u0306\u0307\u0308\u0309\u030a\u030b", ""},
		{"\u030c\u030d\u030e\u030f\u0310\u0311\u0312\u0313\u0314\u0315\u0316", ""},
		{"\u0317\u0318\u0319\u031a\u031b\u031c\u031d\u031e\u031f\u0320\u0321", ""},
		{"\u0322\u0323\u0324\u0325\u0326\u0327\u0328\u0329\u032a\u032b\u032c", ""},
		{"\u032d\u032e\u032f\u0330\u0331\u0332\u0333\u0334\u0335\u0336\u0337", ""},
		{"\u0338\u0339\u033a\u033b\u033c\u033d\u033e\u033f\u0340\u0341\u0342", ""},
		{"\u0343\u0344\u0345\u0346\u0347\u0348\u0349\u034a\u034b\u034c\u034d", ""},
		{"\u034e\u034f\u0350\u0351\u0352\u0353\u0354\u0355\u0356\u0357\u0358", ""},
		{"\u0359\u035a\u035b\u035c\u035d\u035e\u035f\u0360\u0361\u0362\u0363", ""},
		{"\u0364\u0365\u0366\u0367\u0368\u0369\u036a\u036b\u036c\u036d\u036e", ""},
		{"\u036f", ""},
	}

	for _, tc := range testCases {
		actual := ReplaceConfusable(tc.original)
		if actual != tc.expected {
			t.Errorf("replaceConfusable(%s) = %s, want %s", strconv.Quote(tc.original), strconv.Quote(actual), strconv.Quote(tc.expected))
		}
	}
}

func TestConfusableMultiCharactersConfusable(t *testing.T) {
	testCases := []struct {
		original string
		expected string
	}{
		{"æ", "ae"},
		{"Œ", "OE"},
		{"œ", "oe"},
		{"ᒆ", "pi"},
		{"ǋ", "Nj"},
		{"⑪⑾⒒⓫", "11111111"},
		{"⑫⑿⒓⓬", "12121212"},
		{"⑬⒀⒔⓭", "13131313"},
		{"⑭⒁⒕⓮", "14141414"},
		{"⑮⒂⒖⓯", "15151515"},
		{"⑯⒃⒗⓰", "16161616"},
		{"⑰⒄⒘⓱", "17171717"},
		{"⑱⒅⒙⓲", "18181818"},
		{"⑲⒆⒚⓳", "19191919"},
		{"⑳⒇⒛⓴", "20202020"},
	}

	for _, tc := range testCases {
		actual := ReplaceConfusable(tc.original)
		if actual != tc.expected {
			t.Errorf("replaceConfusable(%s) = %s, want %s", strconv.Quote(tc.original), strconv.Quote(actual), strconv.Quote(tc.expected))
		}
	}
}

func TestConfusableRemovesConfusable(t *testing.T) {
	re := regexp.MustCompile(`^[a-zA-Z0-9]+$`)

	testCases := []string{
		"⓿",
		"11⓵➊⑴¹𝟏𝟙１𝟷𝟣⒈𝟭1➀₁①❶⥠",
		"⓶⒉⑵➋ƻ²ᒿ𝟚２𝟮𝟤ᒾ𝟸Ƨ𝟐②ᴤ₂➁❷ᘝƨ",
		"³ȝჳⳌꞫ𝟑ℨ𝟛𝟯𝟥Ꝫ➌ЗȜ⓷ӠƷ３𝟹⑶⒊ʒʓǯǮƺ𝕴ᶾзᦡ➂③₃ᶚᴣᴟ❸ҘҙӬӡӭӟӞ",
		"𝟰𝟺𝟦𝟒➍ҶᏎ𝟜ҷ⓸ҸҹӴӵᶣ４чㄩ⁴➃₄④❹Ӌ⑷⒋",
		"𝟱⓹➎Ƽ𝟓𝟻𝟝𝟧５➄₅⑤⁵❺ƽ⑸⒌",
		"Ⳓ🄇𝟼Ꮾ𝟲𝟞𝟨𝟔➏⓺Ϭϭ⁶б６ᧈ⑥➅₆❻⑹⒍",
		"𝟕𝟟𝟩𝟳𝟽🄈⓻𐓒➐７⁷⑦₇❼➆⑺⒎",
		"𐌚🄉➑⓼８𝟠𝟪৪⁸₈𝟴➇⑧❽𝟾𝟖⑻⒏",
		"൭Ꝯ𝝑𝞋𝟅🄊𝟡𝟵Ⳋ⓽➒੧৭୨９𝟫𝟿𝟗⁹₉Գ➈⑨❾⑼⒐",
		"⓾❿➉➓🔟⑩⑽⒑",
		"⑪⑾⒒⓫",
		"⑫⑿⒓⓬",
		"⑬⒀⒔⓭",
		"⑭⒁⒕⓮",
		"⑮⒂⒖⓯",
		"⑯⒃⒗⓰",
		"⑰⒄⒘⓱",
		"⑱⒅⒙⓲",
		"⑲⒆⒚⓳",
		"⑳⒇⒛⓴",
		"æ",
		"Œ",
		"œ",
		"ᒆ",
		"ǋ",
		"ᴁ",
		"𝑨𝔄ᗄ𝖠𝗔ꓯ𝞐🄐🄰Ꭿ𐊠𝕬𝜜𝐴ꓮᎪ𝚨ꭺ𝝖🅐Å∀🇦₳🅰𝒜𝘈𝐀𝔸дǺᗅⒶＡΑᾋᗩĂÃÅǍȀȂĀȺĄʌΛλƛᴀᴬДАልÄₐᕱªǞӒΆẠẢẦẨẬẮẰẲẴẶᾸᾹᾺΆᾼᾈᾉᾊᾌᾍᾎᾏἈἉἊἋἌἍἎἏḀȦǠӐÀÁÂẤẪ𝛢𝓐𝙰𝘼ᗩ",
		"∂⍺ⓐձǟᵃᶏ⒜аɒａαȃȁคǎმäɑāɐąᾄẚạảǡầẵḁȧӑӓãåάὰάăẩằẳặᾀᾁᾂᾃᾅᾆᾰᾱᾲᾳᾴᶐᾶᾷἀἁἂἃἄἅἆἇᾇậắàáâấẫǻⱥ𝐚𝑎𝒂𝒶𝓪𝔞𝕒𝖆𝖺𝗮𝘢𝙖𝚊𝛂𝛼𝜶𝝰𝞪⍶",
		"🄑𝔙𝖁ꞵ𝛃𝛽𝜷𝝱𝞫Ᏸ𐌁𝑩𝕭🄱𐊡𝖡𝘽ꓐ𝗕𝘉𝜝𐊂𝚩𝐁𝛣𝝗𝐵𝙱𝔹Ᏼᏼ𝞑Ꞵ𝔅🅑฿𝓑ᗿᗾᗽ🅱ⒷＢвϐᗷƁ乃ßცჩ๖βɮБՅ๒ᙖʙᴮᵇጌḄℬΒВẞḂḆɃദᗹᗸᵝᙞᙟᙝᛒᙗᙘᴃ🇧",
		"ꮟᏏ𝐛𝘣𝒷𝔟𝓫𝖇𝖻𝑏𝙗𝕓𝒃𝗯𝚋♭ᑳᒈｂᖚᕹᕺⓑḃḅҍъḇƃɓƅᖯƄЬᑲþƂ⒝ЪᶀᑿᒀᒂᒁᑾьƀҌѢѣᔎ",
		"ꞆႠ℃🄒ᏟⲤ🄲ꓚ𐊢𐌂🅲𐐕🅒☾ČÇⒸＣↃƇᑕㄈ¢८↻ĈϾՇȻᙅᶜ⒞ĆҀĊ©टƆℂℭϹС匚ḈҪʗᑖᑡᑢᑣᑤᑥⅭ𝐂𝐶𝑪𝒞𝓒𝕮𝖢𝗖𝘊𝘾ᔍ",
		"🝌ｃⅽ𝐜𝑐𝒄𝒸𝓬𝔠𝕔𝖈𝖼𝗰𝘤𝙘𝚌ᴄϲⲥсꮯ𐐽ⲥ𐐽ꮯĉｃⓒćčċçҁƈḉȼↄсርᴄϲҫ꒝ςɽϛ𝙲ᑦ᧚𝐜𝑐𝒄𝒸𝓬𝔠𝕔𝖈𝖼𝗰𝘤𝙘𝚌₵🇨ᥴᒼⅽ",
		"🄓Ꭰ🄳𝔡𝖉𝔻𝗗𝘋𝙳𝐷𝓓𝐃𝑫𝕯𝖣𝔇𝘿ꭰⅅ𝒟ꓓ🅳🅓ⒹＤƉᗪƊÐԺᴅᴰↁḊĐÞⅮᗞᑯĎḌḐḒḎᗫᗬᗟᗠᶛᴆ🇩",
		"Ꮷ𝔡𝖉ᑯꓒ𝓭ᵭ₫ԃⓓｄḋďḍḑḓḏđƌɖɗᵈ⒟ԁⅾᶁԀᑺᑻᑼᑽᒄᑰᑱᶑ𝕕𝖽𝑑𝘥𝒅𝙙𝐝𝗱𝚍ⅆ𝒹ʠժ",
		"£ᙓ⋿∃ⴺꓱ𝐄𝐸𝔈𝕰𝖤𝘌𝙴𝛦𝜠ꭼ🄔🄴𝙀𝔼𐊆𝚬ꓰ𝝚𝞔𝓔𝑬𝗘🅴🅔ⒺΈＥƎἝᕮƐモЄᴇᴱᵉÉ乇ЁɆꂅ€ÈℰΕЕⴹᎬĒĔĖĘĚÊËԐỀẾỄỂẼḔḖẺȄȆẸỆȨḜḘḚἘἙἚἛἜῈΈӖὲέЀϵ🇪",
		"əәⅇꬲꞓ⋴𝛆𝛜𝜀𝜖𝜺𝝐𝝴𝞊𝞮𝟄ⲉꮛ𐐩ꞒⲈ⍷𝑒𝓮𝕖𝖊𝘦𝗲𝚎𝙚𝒆𝔢𝖾𝐞Ҿҿⓔｅ⒠èᧉéᶒêɘἔềếễ૯ǝєεēҽɛểẽḕḗĕėëẻěȅȇẹệȩɇₑęḝḙḛ℮еԑѐӗᥱёἐἑἒἓἕℯ",
		"ᖵꘘꓞꟻᖷ𝐅𝐹𝑭𝔽𝕱𝖥𝗙𝙁𝙵𝟊℉🄕🄵𐊇𝔉𝘍𐊥ꓝꞘ🅵🅕𝓕ⒻＦғҒᖴƑԲϝቻḞℱϜ₣🇫Ⅎ",
		"𝐟ᵮ𝑓𝒇𝒻𝓯𝔣𝕗𝖿𝗳𝙛𝚏ꬵꞙẝ𝖋ⓕｆƒḟʃբᶠ⒡ſꊰʄ∱ᶂ𝘧",
		"𝗚𝘎🄖ꓖᏳ🄶Ꮐᏻ𝔾𝓖𝑮𝕲ꮐ𝒢𝙂𝖦𝙶𝔊𝐺𝐆🅶🅖ⒼＧɢƓʛĢᘜᴳǴĠԌĜḠĞǦǤԍ₲🇬⅁",
		"ᶃᶢⓖｇǵĝḡğġǧģց૭ǥɠﻭﻮᵍ⒢ℊɡᧁ𝐠𝑔𝒈𝓰𝔤𝕘𝖌𝗀𝗴𝘨𝙜𝚐",
		"Ἤ🄗𝆦🄷𝜢ꓧ𝘏𝐻𝝜𝖧𐋏𝗛ꮋℍᎻℌⲎ𝑯𝞖🅷🅗ዞǶԋⒽＨĤᚺḢḦȞḤḨḪĦⱧҢңҤῊΉῌἨἩἪἫἭἮἯᾘᾙᾚᾛᾜᾝᾞᾟӉӈҥΉн卄♓𝓗ℋН𝐇𝙃𝙷ʜ𝛨Η𝚮ᕼӇᴴᵸ🇭",
		"ꞕ৸𝕳ꚕᏲℏӊԊꜧᏂҺ⒣ђⓗｈĥḣḧȟḥḩḫẖħⱨհһከኩኪካɦℎ𝐡𝒉𝒽𝓱𝔥𝕙𝖍𝗁𝗵𝘩𝙝𝚑իʰᑋᗁɧんɥ",
		"ⲒἿ🄘🄸ЇꀤᏆ🅸🅘إﺇٳأﺃٲٵⒾＩ៸ÌÍÎĨĪĬİÏḮỈǏȈȊỊĮḬƗェエῘῙῚΊἸἹἺἻἼἽἾⅠΪΊɪᶦᑊᥣ𝛪𝐈𝙄𝙸𝓵𝙡𝐼ᴵ𝚰𝑰🇮",
		"⍳ℹⅈ𝑖𝒊𝒾ı𝚤ɩιιͺ𝛊𝜄𝜾𝞲ꙇӏꭵᎥⓘｉìíîĩīĭïḯỉǐȉȋịḭῐῑῒΐῖῗἰἱἲⅰⅼ∣ⵏ￨׀ا١۱ߊᛁἳἴἵɨіὶίᶖ𝔦𝚒𝝸𝗂𝐢𝕚𝖎𝗶𝘪𝙞ίⁱᵢ𝓲⒤",
		"𝐉𝐽𝑱𝒥𝓙𝔍𝕁𝕵𝖩𝗝𝘑𝙅𝙹ꞲͿꓙ🄙🄹🅹🅙ⒿＪЈʝᒍנﾌĴʆวلյʖᴊᴶﻝጋɈⱼՂๅႱįᎫȷ丿ℐℑᒘᒙᒚᒛᒴᒵᒎᒏ🇯",
		"𝚥ꭻⅉⓙｊϳʲ⒥ɉĵǰјڶᶨ𝒿𝘫𝗷𝑗𝙟𝔧𝒋𝗃𝓳𝕛𝚓𝖏𝐣",
		"𝐊ꝄꝀ𝐾𝑲𝓚𝕶𝖪𝙺𝚱𝝟🄚𝗞🄺𝜥𝘒ꓗ𝙆𝕂Ⲕ𝔎𝛫Ꮶ𝞙𝒦🅺🅚₭ⓀＫĸḰќƘкҠκқҟӄʞҚКҡᴋᴷᵏ⒦ᛕЌጕḲΚKҜҝҞĶḴǨⱩϗӃ🇰",
		"ⓚꝁｋḱǩḳķḵƙⱪᶄ𝐤𝘬𝗄𝕜𝜅𝜘𝜿𝝒𝝹𝞌𝞳𝙠𝚔𝑘𝒌ϰ𝛋𝛞𝟆𝗸𝓴𝓀",
		"𝐋𝐿𝔏𝕃𝕷𝖫𝗟𝘓𝙇ﴼ🄛🄻𐐛Ⳑ𝑳𝙻𐑃𝓛ⳑꮮᏞꓡ🅻🅛ﺈ└ⓁւＬĿᒪ乚ՆʟꓶιԼᴸˡĹረḶₗΓլĻᄂⅬℒⱢᥧᥨᒻᒶᒷᶫﺎᒺᒹᒸᒫ⎳ㄥŁⱠﺄȽ🇱",
		"ⓛｌŀĺľḷḹļӀℓḽḻłﾚɭƚɫⱡ|Ɩ⒧ʅǀוןΙІ｜ᶩӏ𝓘𝕀𝖨𝗜𝘐𝐥𝑙𝒍𝓁𝔩𝕝𝖑𝗅𝗹𝘭𝚕𝜤𝝞ı𝚤ɩι𝛊𝜄𝜾𝞲",
		"ꮇ🄜🄼𐌑𐊰ꓟⲘᎷ🅼🅜ⓂＭмṂ൱ᗰ州ᘻო๓♏ʍᙏᴍᴹᵐ⒨ḾМṀ௱ⅯℳΜϺᛖӍӎ𝐌𝑀𝑴𝓜𝔐𝕄𝕸𝖬𝗠𝘔𝙈𝙼𝚳𝛭𝜧𝝡𝞛🇲",
		"₥ᵯ𝖒𝐦𝗆𝔪𝕞𝓂ⓜｍനᙢ൩ḿṁⅿϻṃጠɱ៳ᶆ𝙢𝓶𝚖𝑚𝗺᧕᧗",
		"𝇙𝇚𝇜🄝𝆧𝙉🄽ℕꓠ𝛮𝝢𝙽𝚴𝑵𝑁Ⲛ𝐍𝒩𝞜𝗡𝘕𝜨𝓝𝖭🅽₦🅝ЙЍⓃҋ៷ＮᴎɴƝᑎ几иՈռИהЛπᴺᶰŃ刀ክṄⁿÑПΝᴨոϖǸŇṆŅṊṈทŊӢӣӤӥћѝйᥢҊᴻ🇳",
		"ոռח𝒏𝓷𝙣𝑛𝖓𝔫𝗇𝚗𝗻ᥒⓝήｎǹᴒńñᾗηṅňṇɲņṋṉղຖՌƞŋ⒩ภกɳпŉлԉȠἠἡῃդᾐᾑᾒᾓᾔᾕᾖῄῆῇῂἢἣἤἥἦἧὴήበቡቢባቤብቦȵ𝛈𝜂𝜼𝝶𝞰𝕟𝘯𝐧𝓃ᶇᵰᥥ∩",
		"𝜽⭘🔿ꭴ⭕⏺🄁🄀Ꭴ𝚯𝚹𝛩𝛳𝜣𝜭𝝝𝝧𝞗𝞡ⴱᎾᏫ⍬𝞱𝝷𝛉𝟎𝜃θ𝟘𝑂𝑶𝓞𝔒𝕆𝕺𝗢𝘖𝙊𝛰㈇ꄲ🄞🔾🄾𐊒𝟬ꓳⲞ𐐄𐊫𐓂𝞞🅞⍥◯ⵁ⊖０⊝𝝤Ѳϴ𝚶𝜪ѺӦӨӪΌʘ𝐎ǑÒŎÓÔÕȌȎㇿ❍ⓄＯὋロ❤૦⊕ØФԾΘƠᴼᵒ⒪ŐÖₒ¤◊Φ〇ΟОՕଠഠ௦סỒỐỖỔṌȬṎŌṐṒȮȰȪỎỜỚỠỞỢỌỘǪǬǾƟⵔ߀៰⍜⎔⎕⦰⦱⦲⦳⦴⦵⦶⦷⦸⦹⦺⦻⦼⦽⦾⦿⧀⧁⧂⧃ὈὉὊὌὍ",
		"ంಂംං૦௦۵ℴ𝑜𝒐𝖔ꬽ𝝄𝛔𝜎𝝈𝞂ჿ𝚘০୦ዐ𝛐𝗈𝞼ဝⲟ𝙤၀𐐬𝔬𐓪𝓸🇴⍤○ϙ🅾𝒪𝖮𝟢𝟶𝙾𝘰𝗼𝕠𝜊𝐨𝝾𝞸ᐤⓞѳ᧐ᥲðｏఠᦞՓòөӧóºōôǒȏŏồốȍỗổõσṍȭṏὄṑṓȯȫ๏ᴏőöѻоዐǭȱ০୦٥౦೦൦๐໐οօᴑ०੦ỏơờớỡởợọộǫøǿɵծὀὁόὸόὂὃὅ",
		"🄟🄿ꓑ𝚸𝙿𝞠𝙋ꮲⲢ𝒫𝝦𝑃𝑷𝗣𝐏𐊕𝜬𝘗𝓟𝖯𝛲Ꮲ🅟Ҏ🅿ⓅＰƤᑭ尸Ṗրφքᴘᴾᵖ⒫ṔｱקРየᴩⱣℙΡῬᑸᑶᑷᑹᑬᑮ🇵₱",
		"ⲣҏ℗ⓟｐṕṗƥᵽῥρрƿǷῤ⍴𝓹𝓅𝐩𝑝𝒑𝔭𝕡𝖕𝗉𝗽𝘱𝙥𝚙𝛒𝝆𝞺𝜌𝞀",
		"🅀🄠Ꝗ🆀🅠ⓆＱℚⵕԚ𝐐𝑄𝑸𝒬𝓠𝚀𝘘𝙌𝖰𝕼𝔔𝗤🇶",
		"𝓆ꝗ𝗾ⓠｑգ⒬۹զᑫɋɊԛ𝗊𝑞𝘲𝕢𝚚𝒒𝖖𝐪𝔮𝓺𝙦",
		"℞🄡℟ꭱᏒ𐒴ꮢᎡꓣ🆁🅡ⓇＲᴙȒʀᖇя尺ŔЯરƦᴿዪṚɌʁℛℜℝṘŘȐṜŖṞⱤ𝐑𝑅𝑹𝓡𝕽𝖱𝗥𝘙𝙍𝚁ᚱ🇷ᴚ",
		"𝚛ꭇᣴℾ𝚪𝛤𝜞𝝘𝞒ⲄГᎱᒥꭈⲅꮁⓡｒŕṙřȑȓṛṝŗгՐɾᥬṟɍʳ⒭ɼѓᴦᶉ𝐫𝑟𝒓𝓇𝓻𝔯𝕣𝖗𝗋𝗿𝘳𝙧ᵲґᵣ",
		"🅂🄪🄢ꇙ𝓢𝗦Ꮪ𝒮Ꮥ𝚂𝐒ꓢ𝖲𝔖𝙎𐊖𝕾𐐠𝘚𝕊𝑆𝑺🆂🅢ⓈＳṨŞֆՏȘˢ⒮ЅṠŠŚṤŜṦṢടᔕᔖᔢᔡᔣᔤ",
		"ᣵⓢꜱ𐑈ꮪｓśṥŝṡšṧʂṣṩѕşșȿᶊక𝐬𝑠𝒔𝓈𝓼𝔰𝕤𝖘𝗌𝘀𝘴𝙨𝚜ގ🇸",
		"🅃🄣七ፒ𝜯🆃𐌕𝚻𝛵𝕋𝕿𝑻𐊱𐊗𝖳𝙏🝨𝝩𝞣𝚃𝘛𝑇ꓔ⟙𝐓Ⲧ𝗧⊤𝔗Ꭲꭲ𝒯🅣⏇⏉ⓉＴтҬҭƬイŦԵτᴛᵀｲፕϮŤ⊥ƮΤТ下ṪṬȚŢṰṮ丅丁ᐪ𝛕𝜏𝝉𝞃𝞽𝓣ㄒ🇹ጥ",
		"ⓣｔṫẗťṭțȶ੮էʇ†ţṱṯƭŧᵗ⒯ʈեƫ𝐭𝑡𝒕𝓉𝓽𝔱𝕥𝖙𝗍𝘁𝘵𝙩𝚝ナ",
		"🅄Џ🄤ሀꓴ𐓎꒤🆄🅤ŨŬŮᑗᑘǓǕǗǙⓊＵȖᑌ凵ƱմԱꓵЦŪՄƲᙀᵁᵘ⒰ŰપÜՍÙÚÛṸṺǛỦȔƯỪỨỮỬỰỤṲŲṶṴɄᥩᑧ∪ᘮ⋃𝐔𝑈𝑼𝒰𝓤𝔘𝕌𝖀𝖴𝗨𝘜𝙐𝚄🇺",
		"𝘂𝘶𝙪𝚞ꞟꭎꭒ𝛖𝜐𝝊𝞄𝞾𐓶ὺύⓤｕùũūừṷṹŭǖữᥙǚǜὗυΰนսʊǘǔúůᴜűųยûṻцሁüᵾᵤµʋủȕȗưứửựụṳṵʉῠῡῢΰῦῧὐὑϋύὒὓὔὕὖᥔ𝐮𝑢𝒖𝓊𝓾𝔲𝕦𝖚𝗎ᶙ",
		"𝑉𝒱𝕍𝗩🄥🅅ꓦ𝑽𝖵𝘝Ꮩ𝚅𝙑𝐕🆅🅥ⓋＶᐯѴᵛ⒱۷ṾⅴⅤṼ٧ⴸѶᐺᐻ🇻𝓥",
		"∨⌄⋁ⅴ𝐯𝑣𝒗𝓋𝔳𝕧𝖛𝗏ꮩሀⓥｖ𝜐𝝊ṽṿ౮งѵעᴠνטᵥѷ៴ᘁ𝙫𝚟𝛎𝜈𝝂𝝼𝞶𝘷𝘃𝓿",
		"𝐖𝑊𝓦𝔚𝕎𝖂𝖶𝗪𝙒𝚆🄦🅆ᏔᎳ𝑾ꓪ𝒲𝘞🆆Ⓦ🅦ｗＷẂᾧᗯᥕ山ѠຟచաЩШώщฬшᙎᵂʷ⒲ฝሠẄԜẀŴẆẈധᘺѿᙡƜ₩🇼",
		"𝐰ꝡ𝑤𝒘𝓌𝔀𝔴𝕨𝖜𝗐𝘄𝘸𝙬𝚠աẁꮃẃⓦ⍵ŵẇẅẘẉⱳὼὠὡὢὣωὤὥὦὧῲῳῴῶῷⱲѡԝᴡώᾠᾡᾢᾣᾤᾥᾦɯ𝝕𝟉𝞏",
		"ꭓꭕ𝛘𝜒𝝌𝞆𝟀ⲭ🞨𝑿𝛸🄧🞩🞪🅇🞫🞬𐌗Ⲭꓫ𝖃𝞦𝘟𐊐𝚾𝝬𝜲Ꭓ𐌢𝖷𝑋𝕏𝔛𐊴𝗫🆇🅧❌Ⓧ𝓧ＸẊ᙭χㄨ𝒳ӾჯӼҳЖΧҲᵡˣ⒳אሸẌꊼⅩХ╳᙮ᕁᕽⅹᚷⵝ𝙓𝚇乂𝐗🇽",
		"᙮ⅹ𝑥𝒙𝓍𝔵𝕩𝖝𝗑𝘅ᕁᕽⓧｘхẋ×ₓ⤫⤬⨯ẍᶍ𝙭ӽ𝘹𝐱𝚡⨰ﾒ𝔁",
		"𝒴🄨𝓨𝔜𝖄𝖸𝘠𝙔𝚼𝛶𝝪𝞤УᎩᎽⲨ𝚈𝑌𝗬𝐘ꓬ𝒀𝜰𐊲🆈🅨ⓎＹὛƳㄚʏ⅄ϔ￥¥ՎϓγץӲЧЎሃŸɎϤΥϒҮỲÝŶỸȲẎỶỴῨῩῪΎὙὝὟΫΎӮӰҰұ𝕐🇾",
		"𝐲𝑦𝒚𝓎𝔂𝔶𝕪𝖞𝗒𝘆𝘺𝙮𝚢ʏỿꭚγℽ𝛄𝛾𝜸𝝲𝞬🅈ᎽᎩⓨｙỳýŷỹȳẏÿỷуყẙỵƴɏᵞɣʸᶌү⒴ӳӱӯўУʎ",
		"🄩🅉ꓜ𝗭𝐙☡Ꮓ𝘡🆉🅩ⓏＺẔƵ乙ẐȤᶻ⒵ŹℤΖŻŽẒⱫ🇿",
		"𝑍𝒁𝒵𝓩𝖹𝙕𝚉𝚭𝛧𝜡𝝛𝞕ᵶꮓ𝐳𝑧𝒛𝓏𝔃𝔷𝕫𝖟𝗓𝘇𝘻𝙯𝚣ⓩｚźẑżžẓẕƶȥɀᴢጊʐⱬᶎʑᙆ",
	}

	for _, test := range testCases {
		actual := ReplaceConfusable(test)

		if utf8.RuneCountInString(actual) < utf8.RuneCountInString(test) {
			t.Errorf("replaceConfusable(%s) expected length > %d, got %d", test, utf8.RuneCountInString(test), utf8.RuneCountInString(actual))
		}

		if !re.MatchString(actual) {
			t.Errorf("replaceConfusable(%s) expected to match /^[a-zA-Z0-9]+$/, got %s", test, actual)
		}
	}
}
