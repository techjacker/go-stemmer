package main

var stopWords = []string{"by", "through", "are", "and", "then", "the", "that", "is", "our", "of", "this", "in", "my"}

// var ReplacementSuffixes2 = map[string]string{
// 	"through": "",
// 	"are": "",
// 	"and": "",
// 	"then": "",
// 	"then": "",
// }

// (m>0)
var ReplacementSuffixes2 = map[string]string{
	"ational": "ate",
	"tional":  "tion",
	"enci":    "ence",
	"anci":    "ance",
	"izer":    "ize",
	"abli":    "able",
	"alli":    "al",
	"entli":   "ent",
	"eli":     "e",
	"ousli":   "ous",
	"ization": "ize",
	"ation":   "ate",
	"ator":    "ate",
	"alism":   "al",
	"iveness": "ive",
	"fulness": "ful",
	"ousness": "ous",
	"aliti":   "al",
	"iviti":   "ive",
	"biliti":  "ble",
}

// (m>0)
var ReplacementSuffixes3 = map[string]string{
	"icate": "ic",
	"ative": "",
	"alize": "al",
	"iciti": "ic",
	"ical":  "ic",
	"ful":   "",
	"ness":  "",
}

// (m>1)
// (m>1 and (*S or *T)) ION ->     adoption       ->  adopt
// v count > 2 eg, friendlier
//                 ccvvcccvvc
var ReplacementSuffixes4 = map[string]string{
	"al":    "",
	"ance":  "",
	"ence":  "",
	"er":    "",
	"ic":    "",
	"able":  "",
	"ible":  "",
	"ant":   "",
	"ement": "",
	"ment":  "",
	"ent":   "",
	"ou":    "",
	"ism":   "",
	"ate":   "",
	"iti":   "",
	"ous":   "",
	"ive":   "",
	"ize":   "",
}

// (m>0) and (*v* - the stem contains a vowel.), eg ccvc
var ReplacementSuffixes5 = map[string]string{
	"ery": "er",
	"ify": "",
}
