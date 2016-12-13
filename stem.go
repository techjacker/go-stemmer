package main

import (
	"bytes"
	"regexp"
	"sort"
)

var vowels = []byte("aeiou")

// Stem returns the stem of a word
func Stem(in []byte) []byte {
	s := bytes.ToLower(in)
	seq := genCVSequence(s)
	mes := measureSequence(seq)
	vCount := bytes.Count(seq, []byte("v"))

	if isStopWord(in) {
		return nil
	}

	// porter step 1a
	s = rmPlural(s)

	// porter step 1b
	// s = rmPastParticles(s)

	// porter step 4
	if mes > 1 && vCount > 2 {
		s = rmSuffix(s, ReplacementSuffixes4)
	}

	// not in porter
	if mes > 0 {
		s = rmSuffix(s, ReplacementSuffixes5)
	}
	// not in porter
	s = rmAdverb(s)

	return s
}

func isStopWord(in []byte) bool {
	ins := string(in)
	sort.Strings(stopWords)
	i := sort.Search(len(stopWords),
		func(i int) bool { return stopWords[i] >= ins })
	if i < len(stopWords) && stopWords[i] == ins {
		return true
	}
	return false
}

func rmPlural(in []byte) []byte {
	suffix := []byte("s")
	re1, _ := regexp.Compile("^(.+?)(ss|i)es$")
	re2, _ := regexp.Compile("^(.+?)([^s])s$")
	if bytes.HasSuffix(in, suffix) {
		if re1.Match(in) {
			return re1.ReplaceAll(in, []byte("$1$2"))
		}
		if re2.Match(in) {
			return bytes.TrimSuffix(in, suffix)
		}
	}
	return in
}

func rmSuffix(in []byte, replacements map[string]string) []byte {
	for k, v := range replacements {
		ks := []byte(k)
		vs := []byte(v)
		if bytes.HasSuffix(in, ks) {
			b := bytes.TrimSuffix(in, ks)
			b = append(b, vs...)
			return b
		}
	}
	return in
}

func rmAdverb(in []byte) []byte {
	suffixes := [][]byte{[]byte("ly"), []byte("li")}
	for _, suf := range suffixes {
		if bytes.HasSuffix(in, suf) {
			return bytes.TrimSuffix(in, suf)
		}
	}
	return in
}
