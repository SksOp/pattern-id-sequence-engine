package utils

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
)

func isCustomePatternValid(CustomeSupport CustomeSupport) error {
	for key, value := range CustomeSupport.Customs {
		if strings.Contains(key, "N") || strings.Contains(key, "Y") || strings.Contains(key, "M") || strings.Contains(key, "D") {
			return fmt.Errorf("trying to override reserve keyword %s", key)
		} else {
			if len(key) != len(value) {
				return fmt.Errorf("length of custom key and length of its crossponding value must")
			}
		}

	}
	return nil
}

// PUBLIC
type CustomeSupport struct {
	Customs map[string]string
}

// PUBLIC
func GetNextId(val ...interface{}) (string, error) {
	if len(val) < 1 {
		return "", errors.New("pattern is required")
	}
	pattern, ok := val[0].(string)
	if !ok {
		return "", fmt.Errorf("type Error: first argument must we a sting")
	}
	if len(val) == 1 {
		return getNextSequenceId(pattern, "", CustomeSupport{})
	}
	lastGen, ok := val[1].(string)
	if !ok {
		return "", fmt.Errorf("type Error: second argument must be string")
	}
	if len(val) == 2 {
		return getNextSequenceId(pattern, lastGen, CustomeSupport{})
	}
	customeSupport, ok := val[2].(CustomeSupport)
	if !ok {
		return "", fmt.Errorf("type Error: second Argument Must be a CustomeSupport")
	}

	if len(val) == 3 {
		return getNextSequenceId(pattern, lastGen, customeSupport)
	}
	if len(val) > 3 {
		return "", errors.New("too many arguments")
	}
	return "", fmt.Errorf("invalid")
}

func getNextSequenceId(pattern string, lastGeneratedSequence string, customPatten CustomeSupport) (string, error) {
	err := isCustomePatternValid(customPatten)
	if err != nil {
		return "", err
	}

	patterns := ParsePattern(pattern)
	result := ""
	if pattern == "" {
		return "", errors.New("pattern is required")
	}
	year := time.Now().Year()
	month := time.Now().Month()
	day := time.Now().Day()
	_isNPresent := 0
	for _, _pattern := range patterns {
		found := false
		/* this is the repeting character*/
		if strings.HasPrefix(_pattern, "N") {
			if lastGeneratedSequence == "" || len(lastGeneratedSequence) == 0 {
				_isNPresent++
				if len(_pattern) > 18 {
					return "", fmt.Errorf("error Too big to increment the value: Max allowed is 18 but provided string contains %d N sequence", len(_pattern))
				}
				k := ""
				for i := 0; i < len(_pattern); i++ {
					k += "0"
				}
				found = true
				result += k
				continue
			} else {

				found = true
				_isNPresent++
				i := strings.Index(pattern, "N")
				if i == -1 {
					return "", fmt.Errorf("pattern Error : String dosent contains N ")
				}

				length := 0
				for j := i; j < len(pattern) && pattern[j] == 'N'; j++ {
					length++
				}
				if i+length > len(lastGeneratedSequence) {
					return "", fmt.Errorf("error: last generated Sequence is shoter then pattern")
				}

				// fmt.Println(i, length, i+length)
				substr := lastGeneratedSequence[i : i+length]
				// fmt.Println(substr, i, i+length-1)
				_N, err := strconv.Atoi(substr)
				if err != nil {
					return "", err
				}
				_N += 1
				extract := fmt.Sprintf("%0*d", length, _N)
				result += extract

				continue
			}
		} else {

			for key, value := range customPatten.Customs {
				//dont add break here
				if len(key) == len(_pattern) && key == _pattern {
					found = true
					result += value
				}
				continue
			}
		}

		if strings.HasPrefix(_pattern, "Y") && len(_pattern) <= 4 {
			var ln int = len(_pattern)
			for i := 0; i < len(_pattern); i += ln {
				// _year := strconv.Itoa(year % int(math.Pow10(ln)))
				result += fmt.Sprintf("%0*d", ln, year%int(math.Pow10(ln)))
			}
			found = true
			continue
		}

		if len(_pattern) >= 2 && len(_pattern)%2 == 0 && _pattern[0:2] == "MM" {
			for i := 0; i < len(_pattern); i += 2 {
				result += fmt.Sprintf("%02d", month)
			}
			found = true
			continue
		}
		if len(_pattern) >= 2 && len(_pattern)%2 == 0 && _pattern[0:2] == "DD" {
			for i := 0; i < len(_pattern); i += 2 {
				result += fmt.Sprintf("%02d", day)
			}
			found = true
			continue
		}
		if len(_pattern) >= 2 && _pattern[0:2] == "YY" && len(_pattern) == 2 {
			_year := strconv.Itoa(year % 100)
			result += _year
			found = true
			continue
		}

		if !found {
			return "", fmt.Errorf("custom pattern %s is not supported: value not specified", _pattern)
		}
	}
	if _isNPresent == 1 {
		if len(result) > len(lastGeneratedSequence) && lastGeneratedSequence != "" && len(lastGeneratedSequence) != 0 {
			return "", fmt.Errorf("overflow")
		}
		return result, nil
	}
	if _isNPresent == 0 {
		return "", fmt.Errorf("pattern Error : N not found in the syntex")
	}
	return "", fmt.Errorf("Pattern Error: Too many occurence of N ")
}
