package utils

// func InitializePatternConfig(customPatterns []SupportedPatterns) (*PatternConfig, error) {
// 	supportedPatternsMap := make(map[string]SupportedPatterns)
// 	for _, pattern := range SUPPORTED_PATTERNS {
// 		supportedPatternsMap[pattern.Identifier] = pattern
// 	}

// 	for _, custom := range customPatterns {
// 		isValid, err := IsValidSupportPattern(custom)
// 		if !isValid {
// 			return nil, err
// 		}

// 		supportedPattern, exists := supportedPatternsMap[custom.Identifier]
// 		if !exists {
// 			return nil, fmt.Errorf("custom pattern %s is not supported", custom.Identifier)
// 		}

// 		if custom.MaxConsecutiveInstances > supportedPattern.MaxConsecutiveInstances {
// 			return nil, fmt.Errorf("custom pattern %s has maxConsecutiveInstances greater than supported", custom.Identifier)
// 		}

// 		if custom.MaxInstanceInSequence > supportedPattern.MaxInstanceInSequence {
// 			return nil, fmt.Errorf("custom pattern %s has maxInstanceInSequence greater than supported", custom.Identifier)
// 		}
// 	}

// 	return &PatternConfig{customPatterns}, nil
// }
// func IsValidSupportPattern(custom SupportedPatterns) (bool, error) {
// 	if custom.Identifier == "" {
// 		return false, errors.New("identifier is required")
// 	}
// 	if custom.MaxConsecutiveInstances <= 0 {
// 		return false, errors.New("maxConsecutiveInstances should be greater than 0")
// 	}
// 	if custom.MaxInstanceInSequence <= 0 {
// 		return false, errors.New("maxInstanceInSequence should be greater than 0")
// 	}
// 	return true, nil
// }
