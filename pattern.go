package patternidsequenceengine

type Pattern struct {
	pattern string //"YYYYMMDDIIIICC"
}

func NewPattern(pattern string) *Pattern {
	return &Pattern{
		pattern: pattern,
	}
}

func (p *Pattern) parsePattern() []string {
	/* break YYYYDDMMIIIICC into
	   - YYYY
	   - DD
	   - MM
	   - IIII
	   - CC
	*/

	return []string{"YYYY", "DD", "MM", "IIII", "CC"}

}

func (p *Pattern) fillInPatternBasedOnConfigAndParsedValue() {

}
