package patternidsequenceengine

import "fmt"

func Example() {
	// create a custom pattern

	customPatters := []SupportedPatterns{
		{"C", 10, 1}, // classes ->
		{"P", 6, 2},  // pincode
	}
	support := InitialiseSupports(customPatters)
	seqService := InitialisePatternService(&support)

	// based on school the pattern witll change

	pattern := "YYYYMMDDCCPPPP"
	lastID := "2019010101000001"
	/*
		where YYYY -> 2019
		MM -> 01
		DD -> 01
		CCCC -> 01
		PP -> 000001
	*/

	nextID, err := seqService.CreateNextSequenceID(pattern, lastID)

}
