package mathalgos

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

type FixedLengthCoding struct {
	alphabet   []string
	charToCode map[string]string
	codeToChar map[string]string
	codeLength int
}

func NewFixedLengthCoding(input string) *FixedLengthCoding {
	alphabetSet := make(map[string]struct{})
	for _, char := range input {
		alphabetSet[string(char)] = struct{}{}
	}
	alphabet := make([]string, 0, len(alphabetSet))
	for char := range alphabetSet {
		alphabet = append(alphabet, char)
	}
	sort.Strings(alphabet)

	codeLength := int(math.Ceil(math.Log2(float64(len(alphabet)))))
	charToCode := make(map[string]string)
	codeToChar := make(map[string]string)
	for i, char := range alphabet {
		code := fmt.Sprintf("%0*b", codeLength, i)
		charToCode[char] = code
		codeToChar[code] = char
	}

	return &FixedLengthCoding{
		alphabet:   alphabet,
		charToCode: charToCode,
		codeToChar: codeToChar,
		codeLength: codeLength,
	}
}

func RecreateFromAlphabet(alphabet map[string]string) *FixedLengthCoding {
	codeToChar := make(map[string]string)
	for char, code := range alphabet {
		codeToChar[code] = char
	}

	var codeLength int
	for _, code := range alphabet {
		codeLength = len(code)
		break
	}

	return &FixedLengthCoding{
		charToCode: alphabet,
		codeToChar: codeToChar,
		codeLength: codeLength,
	}
}

func (f *FixedLengthCoding) Encode(input string) string {
	var encoded strings.Builder
	for _, char := range input {
		encoded.WriteString(f.charToCode[string(char)])
	}
	return encoded.String()
}

func (f *FixedLengthCoding) Decode(encoded string) string {
	var decoded strings.Builder
	for i := 0; i < len(encoded); i += f.codeLength {
		code := encoded[i : i+f.codeLength]
		decoded.WriteString(f.codeToChar[code])
	}
	return decoded.String()
}

func (f *FixedLengthCoding) GetAlphabetDict() map[string]string {
	return f.charToCode
}

func (f *FixedLengthCoding) AverageCodeLength() int {
	return f.codeLength
}

type ProbabilityCalculating struct {
	string       string
	letterCounts map[string]int
	totalLetters int
}

func NewProbabilityCalculating(input string) *ProbabilityCalculating {
	letterCounts := make(map[string]int)
	for _, char := range input {
		letterCounts[string(char)]++
	}
	return &ProbabilityCalculating{
		string:       input,
		letterCounts: letterCounts,
		totalLetters: len(input),
	}
}

func (p *ProbabilityCalculating) GetProbabilities() map[string]float64 {
	probabilities := make(map[string]float64)
	for letter, count := range p.letterCounts {
		probabilities[letter] = float64(count) / float64(p.totalLetters)
	}
	return probabilities
}

type ShennonFanoCoding struct {
	probabilityCalculating *ProbabilityCalculating
	charToCode             map[string]string
	codeToChar             map[string]string
}

func NewShennonFanoCoding(input string) *ShennonFanoCoding {
	probabilityCalculating := NewProbabilityCalculating(input)
	sortedSymbols := make([]struct {
		char string
		prob float64
	}, 0, len(probabilityCalculating.letterCounts))

	for char, prob := range probabilityCalculating.GetProbabilities() {
		sortedSymbols = append(sortedSymbols, struct {
			char string
			prob float64
		}{char, prob})
	}
	sort.Slice(sortedSymbols, func(i, j int) bool {
		return sortedSymbols[i].prob > sortedSymbols[j].prob
	})

	charToCode := make(map[string]string)
	s := &ShennonFanoCoding{
		probabilityCalculating: probabilityCalculating,
		charToCode:             charToCode,
		codeToChar:             make(map[string]string),
	}
	s.createCodeTree(sortedSymbols, "", charToCode)

	for char, code := range charToCode {
		s.codeToChar[code] = char
	}

	return s
}

func RecreateFromCodes(codes map[string]string) *ShennonFanoCoding {
	codeToChar := make(map[string]string)
	for char, code := range codes {
		codeToChar[code] = char
	}

	return &ShennonFanoCoding{
		charToCode: codes,
		codeToChar: codeToChar,
	}
}

func (s *ShennonFanoCoding) createCodeTree(symbols []struct {
	char string
	prob float64
}, prefix string, charToCode map[string]string) {
	if len(symbols) == 1 {
		charToCode[symbols[0].char] = prefix
		return
	}
	total := 0.0
	for _, symbol := range symbols {
		total += symbol.prob
	}
	runningSum := 0.0
	splitIndex := 0
	for i, symbol := range symbols {
		runningSum += symbol.prob
		if runningSum*2 >= total {
			splitIndex = i
			break
		}
	}
	left := symbols[:splitIndex+1]
	right := symbols[splitIndex+1:]
	s.createCodeTree(left, prefix+"0", charToCode)
	s.createCodeTree(right, prefix+"1", charToCode)
}

func (s *ShennonFanoCoding) Encode(input string) string {
	var encoded strings.Builder
	for _, char := range input {
		encoded.WriteString(s.charToCode[string(char)])
	}
	return encoded.String()
}

func (s *ShennonFanoCoding) Decode(encoded string) string {
	var decoded strings.Builder
	code := ""
	for _, bit := range encoded {
		code += string(bit)
		if char, exists := s.codeToChar[code]; exists {
			decoded.WriteString(char)
			code = ""
		}
	}
	return decoded.String()
}

func (s *ShennonFanoCoding) GetAlphabetDict() map[string]string {
	return s.charToCode
}

func (s *ShennonFanoCoding) AverageCodeLength() float64 {
	probabilities := s.probabilityCalculating.GetProbabilities()
	totalLength := 0.0
	for char, code := range s.charToCode {
		totalLength += float64(len(code)) * probabilities[char]
	}
	return totalLength
}
