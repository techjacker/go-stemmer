package main

import "bytes"

// unicode.ToLower(in[1])
func isConsonant(in []rune, i int) bool {
	if bytes.IndexRune(vowels, in[i]) != -1 {
		return false
	}
	if in[i] == 'y' {
		if i == 0 {
			return true
		}
		return !isConsonant(in, i-1)
	}
	return true
}

func genCVSequence(in []byte) []byte {
	var buf bytes.Buffer
	inR := bytes.Runes(in)
	for i := range in {
		if isConsonant(inR, i) {
			buf.Write([]byte("c"))
		} else {
			buf.Write([]byte("v"))
		}
	}
	return buf.Bytes()
}

func measureSequence(in []byte) int {
	return bytes.Count(in, []byte("vc"))
}

// *v* - the stem contains a vowel
func containsVowel(seq []byte) bool {
	return bytes.Contains(seq, []byte("v"))
}
