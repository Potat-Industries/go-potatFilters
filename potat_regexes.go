package potatFilters

import "regexp"

// PotatBotat's chat filters but in Go.
var (
	potatGeneralRacism1  = regexp.MustCompile(`(?i)(\b((=[nhk])(n[i1!¡jl]b+[e3]r|nygg[e3]r|higger|kneeger)[s5z]?)\b)|((chinam[ae]n|ching[\W_]*chong))|((towel|rag|diaper)[\W_]*head[s5z]?)|((sheep|goat|donkey)\W?(fuck|shag)\w*)|((sand|dune)[\W_]*(n[i1!¡jl]g(!ht)|c[o0]{2}n|monk[iey]+)\w*)`)
	potatGeneralRacism2  = regexp.MustCompile(`(?i)((the h[o0]l[o0]caust|gen[o0]cide|there was)[\s\S]*?(the holocaust|genocide)[\s\S]*?((didn[ ''‘’´` + "`" + `]?t|never) happened|(is|was) a lie)|there was[\s\S]*?(no|n[ ''‘’´` + "`" + `]?t an?y?)[\s\S]*?\w*[\s\S]*?(genocide|holocaust))|in[\W_]*bred[s5z]?|filthy jew|bl[a4]cks?|africans? bastard|musl[i1]ms are (violent )?t[e3]rrorists?|r[e3]t[a4]rded m[0o]nkey|bl[a4]cks?|africans? (are|can be|were) (subhuman|primitive)|blackface`)
	potatGeneralSlurs    = regexp.MustCompile(`(?i)(\b(f|ph)[áàâäãåa@][g4]+[e3o0]*t*\b)|(\btr[a@4]nn[y¡i1!jl]es?|tr[a@4]nn|trans[vf]est[iy]te|transfag|tranny|trapsexual|she\W?males?)([s5z]?)|(\bbull)?d[yi]ke[s5z]?|(fudge\W?packer|muff\W?diver|(carpet|rug)\W?muncher|pillow\W?biter|shirt\W?lifter|shit\W?stabber|turd\W?burglar)|\bboiola\b|\bplayo\b|\bwomen are nothing more than objects\b|\bwomen are objects\b|\bholocaust\b|(\b[fḞḟ][a4@][g]\b|[fḞḟ]+[a4@]+[g]+[o0]+[t]+)`)
	potatRacoonWord      = regexp.MustCompile(`(?i)\bc[o0]{2}ns\b`)
	potatNWord           = regexp.MustCompile(`(?i)\b[Nnñ]+[i1|]+[ckgbB6934qğĜƃ၅5]+[e3]u?r+|nigg`)
	potatCWord           = regexp.MustCompile(`(?i)\bc+h+i+n+k+s*`)
	potatTWord           = regexp.MustCompile(`(?i)\bt+r+[a4]+n+(i+[e3]+|y+|[e3]+r+)s*`)
	potatFWord           = regexp.MustCompile(`(?i)\bf+[a4@]+[gbB6934qğĜƃ၅5]+([oei]+t+(r+y+|r+i+[e3]+)?)?s*`)
	potatNonEnglishSlurs = regexp.MustCompile(`(?i)amerykaniec|\bangol\b|arabus|asfalt|bambus|brudas|brudaska|Brytol|chachoł|chinol|ciapaty|czarnuch|fryc|gudłaj|helmut|japoniec|kacap|kacapka|kitajec|koszerny|kozojebca|kudłacz|makaroniarz|małpa|Moskal|negatyw|parch|pejsaty|rezun|Rusek|Ruska|skośnooki|syfiara|syfiarz|szkop|szmatogłowy|szuwaks|szwab|szwabka|\bturas\b|wietnamiec|żabojad|żółtek|żydek|Żydzisko|zabojad|zoltek|zydek|zydzisko|matoglowy|chachol|szuwak|\btura\b|pidor`)

	potatViolence1 = regexp.MustCompile(`(?i)\b(h[i1!¡jl]tl[e3]r|kill your self|kms|kys|simp|incel)\b|i[''‘’´` + "`" + `]? (ll|will|( a)?m(m?a| go(ing to|nna))?|wan(t to|na))( \S+)? (k([i1!jl.\-_]{3}|\\?[^a-z\d\s]ll)|shoot|murder|hang|lynch|poison) ((y+[o0ua]+|u+))(r( \S+)? family)?|cut (y([o0u]+r|o)|ur)\W?sel(f|ves)|should(a|\W?ve| have)* ((k[i1!jl.\-_](ll|lled)|hanged|hung|shot|shoot|exterminated|suicided|roped( \Wsel\w+)? (into|off|from)|drowned|necked) (y([o0u]+r|o)|ur|the[my]|dem)\W?sel(f|ves)|aborted (y([o0ua]+r?|o)|ur?))|go (die|jump (off|out|from))|should ?n[o''‘’´` + "`" + `]?t (be|stay) alive|\br[a4@]p[il1]st\b|\br[a4]p[e3]\b|\bbhead\b|\bbehead\b`)
	potatViolence2 = regexp.MustCompile(`(?i)((k[i1]l+|[e3]nd|sh[0oO]+t)\s?(y[0oO]ur?)\s?(s[e3]l+f)?)`)

	potatSelfHarm = regexp.MustCompile(`(?i)(drink (poison|bleach)|slit (y([o0u]+r|o)|ur)|r[a4@]p[e3]\W?(toy|meat|doll|slut|bait|slave|material|[s5$z](l|[^\w\s])([uv]|[^\w\s])[t7]|wh([o0]|[^\w\s])((r|[^\w\s])[e3]|[o0][a@4e3])|hole|face|body|pig)[s5z]?|com+it suicide|end your( own)? life|take your( own)? life|p[e3]d[o0]ph[i1]l[e3]|p[e0]d[o0]|eat (a|my) (dick|cock|penis)|sieg heil|heil hitler)`)

	potatSexualHarassment1 = regexp.MustCompile(`(?i)pull the [^\s.]+( dam\w*| fuck\S+)? trigger(?: (?:on(?: (?:yo)?ur\W?self))?)?|blows? (\w+\W+(?:\.)){1,4}(?:brains? out)?|play in (some )?traffic|(get|be) raped(?: (?:with| on\b))?(?: (?:can be raped|meant to be raped))?|r[a4@]p(?:[e3][sd]?\b|[i1!¡jl]ng) (her|him|the[my]|them)|throats? (cut|ripped|slit)|pedo|pedobear|(lick|eat|suck|gobble) (my|your|his|her|their|our) (cock|dick|balls|cum|kok|coc)|^get (cancer|a(?: \w+)? tumor|AIDS|HIV|covid\w*|coronavirus|sick)|\b(suck|lick|gobble|consume|eat)\b.*?\b(my|these)\b.*?\b(cock|penis|dick|balls|nuts)|((show|flash|expose) (you|your|those|them) (tits|boobs|breasts|ass|cock|pussy|vagina|crotch))`)
	potatSexualHarassment2 = regexp.MustCompile(`(?i)(\b(stick|shove|insert|force|put)\b.*?\b(in (my|their|h[eris]{2}|your))\b.*?\b(ass|butt|vagina|asshole|cunt)\b)|(\b([1li][0o][il][il]s?|[1li][0o]l[il]cons?)\b)`)

	potatSexism = regexp.MustCompile(`(?i)\bwom[e3a]n\s*(belong|should|go)\s*(in|2|go|be|to)?\s*(the\s*)?(k[i1]tch[e3]n|c[0o]{2}k|cl[e3]an)(ing)?\b`)

	potatAbleism = regexp.MustCompile(`(?i)(r+[\W_]*[e3a4@i1!¡jlw]*[\W_]*[t7]+[\W_]*[a4@e3]*[\W_]*r+[\W_]*[dt]+[\W_]*([e3i1!¡jl]+[\W_]*[dt]+[\W_]*)?([s5z]|retarded\w*|ation)?|((th|d|it[\W_]*i?s)((is|at[\W_]*|[ts]|it)[ ''‘’´` + "`" + `]?i?s)?|ese|ose|em) autis(t(ic|ism)?|m)|retard|ass[\W_]?burger)`)

	potatAdvertising = regexp.MustCompile(`(?i)(?:[f]+[o0]+[l1]+[o0]+[w]+[s]?|raid|host|w(a|4)tch|view|ch(e|3)ck|j(o|0)in|(?:go|come)\s?to)\s(?:o(?:ut|n))?\s?(?:m[ye]|us|him|her|them)\s(?:stream|channel|live|out\b)|(?:i'?m|we're|us|s?he'?s?|they'?re)\s?(?:live|streaming)|(f|f[o0][wl]|flw|[f]+[o0]+[1l]+[o0]+[w]+[s]?)\s*(4|f(o|0)r)\s*(f|f[o0][wl]|flw|[f]+[o0]+[l1]+[o0]+[w]+[s]?)|[f]+[o0]+[l1]+[o0]+[w]+[s]? (me|them|us|him|her|b(4|a)ck)|f[0o]ll[0o]w[i1]ng ([e3]v[3e]ry(one|1)|anyone)?\s?(b[a4]ck)?`)

	potatAgeTos = regexp.MustCompile(`(?i)\b(?:(?:i|my age)\s*['’]?\s*(?:am|'m|m| is|will be)\s*(?:(under|below)\s*)?(?:less\s*than\s*)?\s+(1[0-4]|[5-9]|(?:one|two|three|four|five|six|seven|eight|nine|ten|eleven|twelve|thirteen)))`)
)
