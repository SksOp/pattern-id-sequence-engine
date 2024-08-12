package patternidsequenceengine

type Sequence struct {
	PatternConfig *PatternConfig
}

func InitialisePatternService(pc *PatternConfig) *Sequence {
	return &Sequence{
		PatternConfig: pc,
	}

}

func (s *Sequence) IsValidPatternBasedOnConfig(patternString string) bool {
	// check for deault pattern
	p := NewPattern(patternString)
	// case Y
	var yInstances int = 0
	// check for year patterns
	yConfig := s.PatternConfig.Supports[0]
	for _, v := range p.parsePattern() {
		// if the seq not contains any Y in continue
		if v != "Y" {
			continue
		}

		// check for max consecutive instances
		if len(v) > yConfig.maxConsecutiveInstances {
			return false
		}

		// check for max instance in sequence
		yInstances++
	}
	if yInstances > yConfig.maxInstanceInSequence {
		return false
	}
	// case N

	// case D

	// case M

	// case I

	// check for custom pattern
	// This I will take care later
}

func (s *Sequence) CreateNextSequenceID(patternString string, lastID string) (string, error) {
	return "", nil
}
