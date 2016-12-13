package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func leftPad2Len(s string, padStr string, overallLen int) string {
	var padCountInt int
	padCountInt = 1 + ((overallLen - len(padStr)) / len(padStr))
	var retStr = strings.Repeat(padStr, padCountInt) + s
	return retStr[(len(retStr) - overallLen):]
}

func rightPad2Len(s string, padStr string, overallLen int) string {
	var padCountInt int
	padCountInt = 1 + ((overallLen - len(padStr)) / len(padStr))
	var retStr = s + strings.Repeat(padStr, padCountInt)
	return retStr[:overallLen]
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter text: ")
	text, _ := reader.ReadBytes('\n')

	stats, err := CountStems(text)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		return
	}

	hks := "STEM"
	hvs := leftPad2Len("COUNT", " ", 26-len(hks))
	fmt.Printf("\n%v%v\n", hks, hvs)
	fmt.Println(strings.Repeat("-", 26))
	for k, v := range stats {
		ks := rightPad2Len(fmt.Sprintf("%v", k), " ", 25)
		fmt.Printf("%v%v\n", ks, v)
	}
	fmt.Println("")
}
