package patternidsequenceengine

type PatternConfig struct {
	Supports []SupportedPatterns
}

type SupportedPatterns struct {
	indetifier              string // e.g. "Y"
	maxConsecutiveInstances int    // e.g. 4    // YYYYYYYYY
	maxInstanceInSequence   int    // e.g. 1 // of we want year to be defines only once in the sequence
}

var SUPPORTED_PATTERNS = []SupportedPatterns{
	{"Y", 4, 10}, //year
	{"N", 10, 1}, // number
	{"D", 2, 1},  // day
	{"M", 2, 1},  // month
	{"I", 10, 1}, // identifier
}

func InitialiseSupports(customSupportedPatters []SupportedPatterns) PatternConfig {
	return PatternConfig{
		Supports: append(customSupportedPatters, SUPPORTED_PATTERNS...),
	}

}
