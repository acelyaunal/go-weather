package main

import "strings"

func lc(s string) string { return strings.ToLower(strings.TrimSpace(s)) }

var iso2ToName = map[string]string{
	"de": "germany", "tr": "turkey", "us": "united states", "gb": "united kingdom", "uk": "united kingdom",
	"fr": "france", "es": "spain", "it": "italy", "nl": "netherlands", "be": "belgium", "at": "austria",
	"ch": "switzerland", "ru": "russia", "ua": "ukraine", "gr": "greece", "pl": "poland",
}

var trSynonyms = map[string]string{
	"almanya": "germany",
	"türkiye": "turkey", "turkiye": "turkey",
	"abd": "united states", "amerika": "united states",
	"ingiltere": "united kingdom", "bk": "united kingdom",
	"fransa": "france", "i̇talya": "italy", "italya": "italy",
	"i̇spanya": "spain", "ispanya": "spain",
}

func normalizeCountryToken(tok string) string {
	t := lc(tok)
	if v, ok := trSynonyms[t]; ok {
		return v
	}
	if len(t) == 2 {
		if v, ok := iso2ToName[t]; ok {
			return v
		}
	}
	return t
}

func countryMatch(resultCountry, want string) bool {
	c := lc(resultCountry)
	w := normalizeCountryToken(want)
	if c == w {
		return true
	}
	return strings.Contains(c, w)
}
