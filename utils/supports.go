package utils

// type PatternConfig struct {
// 	Supports []SupportedPatterns
// }

// type SupportedPatterns struct {
// 	Identifier              string // e.g. "Y"
// 	MaxConsecutiveInstances int    // e.g. 4    // YYYYYYYYY
// 	MaxInstanceInSequence   int    // e.g. 1 // of we want year to be defines only once in the sequence
// }

// var SUPPORTED_PATTERNS = []SupportedPatterns{
// 	{"Y", 4, 10}, // year
// 	{"N", 10, 1}, // number
// 	{"D", 2, 1},  // day
// 	{"M", 2, 1},  // month
// 	{"I", 10, 1}, // identifier
// }

// func InitialiseSupports(customSupportedPatters []SupportedPatterns) PatternConfig {
// 	return PatternConfig{
// 		Supports: append(customSupportedPatters, SUPPORTED_PATTERNS...),
// 	}
// }
