package stringLib

import (
	"fmt"
	"strings"
)

// most functions from strings lib are functions that does not changes original string
// but returns changed string in result
// there some usefull funcs:

func LetsTalkBoutStringLibs() {
	strForPresentation := "AbobuS AbobuS AboltuS AbormotuS PukuS KakuS"
	fmt.Println("str before changes:", strForPresentation)
	
	// upper and lowercase
	fmt.Println("str after lowercase:", strings.ToLower(strForPresentation))
	fmt.Println("str after uppercase:", strings.ToUpper(strForPresentation))
	
	// Contains, ContainsAny.
	// contains answer on question: Does string have substring in it
	fmt.Println("does our str have substr bolt? Answer:", strings.Contains(strForPresentation, "bolt"))
	
	// ContainsAny instead can amswer does string have any symbol from another string
	fmt.Println("does our string have ant symbol from string 'АбобусАболтус'", strings.ContainsAny(strForPresentation, "АбобусАболтус"))
	
	// Also i found useful func Count, that answer gow many substrs into str
	fmt.Println("how many 'bo' in str?", strings.Count(strForPresentation, "bo"))
	
	// Some funcs for sep str:
	// Cut and Split series of funcs:
	// Cut, CutPrefix, CutSuffix, they are returns modified string that cutted off subst that we wanted
	// Also they are returns bool value reports were substr found or not
	cutResultPrefix, cutResultSufix, cutResultBool := strings.Cut(strForPresentation, "AboltuS")
	fmt.Println("lets cut out from str 'AbobuS' thats what we got:", cutResultPrefix, "|", cutResultSufix, "|",cutResultBool)

	// CutPrefix, trying to erase prefix if prefix is like substr
	cutResultPrefix, cutResultBool = strings.CutPrefix(strForPresentation, "PukuS")
	fmt.Println("lets cut out from strs start 'PukuS' thats what we got:", cutResultPrefix, "|",cutResultBool)

	// same for CutSufix
	cutResultSufix, cutResultBool = strings.CutSuffix(strForPresentation, " KakuS")
	fmt.Println("lets cut out from strs end ' KakuS' thats what we got:", cutResultSufix, "|",cutResultBool)

	// Split is series of funcs that sep strings many times and does not report about were substr found or not
	// Split[Seq After N], Split returns slice with substring separated with input str.
	// If we add After, it'll include separator like a, b, c sep(,) => ["a,", " b,", " c"]
	// If we added Seq, we'll get iterarot on separated str, kinda for single use.
	// Also there option like N. It adds new arg 'n'
	// n < 0, return all
	// n = 0, return nil
	// n > 0, return first n from result
	fmt.Println("so lets split str on ' A'", strings.Split(strForPresentation," A"), len(strings.Split(strForPresentation," A")))
	fmt.Println("split after", strings.SplitAfter(strForPresentation, " "), len(strings.SplitAfter(strForPresentation, " ")))
	fmt.Print("split seq: ")
	for iter := range strings.SplitSeq(strForPresentation, " A") {
		fmt.Printf("%s | ", iter)
	}

	// Important family of funcs are Index, its like contains, but in result you'll get first index of substr ins str or -1 if there no substr in str
	// if we'll add suffix Any, we'll search for any symbol of substr in str
	// Suffix Byte will search for instance of byte in byte slice interpretion of str
	// Suffix Func take as argument function that take as argument Rune, and return bool. Function will search for first rune that makes func true
	// Suffix Rune make us looking for rune, cannot be used with Last
	// Prefix Last make us looking for last instance instand of first
	fmt.Println("\nLats ask first and lust S in sentence: first",
		strings.IndexRune(strForPresentation, 'S'),
		"last",
		strings.LastIndexAny(strForPresentation, "S"))

	// Trim functions erase head and tail symbols matched with sample str
	// Left prefix erase only suffix that matched with sample
	// Right sufix erase only suffix that matched with sample
	// Func suffix erase all symbols that makes func (rune) bool return true can be combined with left|right
	// Space suffix erase all non-important space symbols
	fmt.Println("Trim is func that erase symbols from str: ndfsgajappqpwpqpaaaaaaaeeEEIuoi, lets try trim vowel letters:",
		strings.Trim("ndfsgajappqpwpqpaaaaaaaeeEEIuoi", "aouieAUIOE"))

	// Replace funcs are useful and have simple sintaxis: str, substrToReplace, replaceStr, n, if n < 0, it'll replace all
	fmt.Println("Lets replace bolt in our sentence on 'REDUCKTED'", strings.Replace(strForPresentation, "bolt", "REDUCKTED", 1))
	fmt.Println("Lets replace all bo in our sentence on 'REDUCKTED'", strings.ReplaceAll(strForPresentation, "bo", "REDUCKTED"))

	// also there kinda funny funcs:
	// Lines, map string on \n, keeping this symbol in it
	for line := range strings.Lines(strings.ReplaceAll(strForPresentation, " ", "\n")) {
		fmt.Printf("so this is line after replace all spaces into \\n. Line: %v", line)
	}

	// Repeat func returns str repeated n times
	fmt.Println(strings.Repeat(strForPresentation, 4))

	// Map use some func for every rune in string
	fmt.Println(strings.Map(func (input rune) rune {
		if input >= 'a' && input <= 'z' {
			return '0'
		}
		return input
	}, strForPresentation))
}
