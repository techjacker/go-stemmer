package main

import (
	"bufio"
	"bytes"
	"unicode"
)

func stemScanWords(data []byte, atEOF bool) (int, []byte, error) {
	advance, token, err := bufio.ScanWords(data, atEOF)
	return advance, bytes.TrimRightFunc(token, unicode.IsPunct), err
}

// CountStems returns a map containing counts of each stem
func CountStems(in []byte) (map[string]int, error) {
	counter := NewCounter()
	scanner := bufio.NewScanner(bytes.NewReader(in))
	scanner.Split(stemScanWords)

	for scanner.Scan() {
		word := scanner.Bytes()
		if stem := Stem(word); stem != nil {
			counter.Add(stem)
		}
	}
	return counter.GetOccurrences(), scanner.Err()
}
