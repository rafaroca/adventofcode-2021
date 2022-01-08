package day8

import (
	"sort"
	"strconv"
	"strings"
)

type Day struct{}

type SegMapping map[string]string

func (t Day) Part1(input []string) int {
	onefourseveneights := 0
	for _, rawLine := range input {
		_, outputs := parseSignalsAndOutput(rawLine)
		for _, output := range outputs {
			switch len(output) {
			case 2, 3, 4, 7:
				onefourseveneights++
			}
		}
	}
	return onefourseveneights
}

func parseSignalsAndOutput(rawLine string) (signals, outputs []string) {
	if len(rawLine) == 0 {
		return
	}
	linePart := strings.Split(rawLine, " | ")
	signals = strings.Split(linePart[0], " ")
	outputs = strings.Split(linePart[1], " ")
	return
}

var segs = map[int]string{0: "abcefg", 1: "cf", 2: "acdeg", 3: "acdfg", 4: "bcdf", 5: "abdfg", 6: "abdefg", 7: "acf", 8: "abcdefg", 9: "abcdfg"}

func (t Day) Part2(input []string) int {
	outputSum := 0

	for _, rawLine := range input {
		// map of letter to possible candidate letter a -> abcdefg, b -> abcdefg, ...
		candidates := InitializeCandidates()
		signals, outputs := parseSignalsAndOutput(rawLine)
		// remove possible candidates on each signal. "eb" then e -> cf, b -> cf. All others may not contain cf anymore
		for _, signal := range signals {
			switch length := len(signal); length {
			case 2, 3, 4, 7:
				{
					segsForLength := SegForLength(length)
					ReduceToCandidates(&candidates, signal, segsForLength)
					RemoveCandidates(&candidates, Exclude(segs[8], signal), segsForLength)
				}
			}
		}

		// create possible mappings from candidate list since they are still ambiguous
		var candidatesPermutations []SegMapping
		candidatesPermutations = append(candidatesPermutations, candidates)
		validCandidateList := FindValidCandidates(candidatesPermutations, signals)

		// decode output, sum up
		outputFigures := ""
		for _, output := range outputs {
			outputFigures += DecodeOutput(&validCandidateList, output)
		}
		output, _ := strconv.Atoi(outputFigures)
		outputSum += output
	}

	return outputSum
}

func FindValidCandidates(candidatesPerm []SegMapping, signals []string) SegMapping {
	currentCandidate := candidatesPerm[0]
	// When there is ambiguity in our candidate, disambiguate and append to the slice, then recur
	if ambiguousKey := FirstAmbiguousMappingKey(currentCandidate); ambiguousKey != "" {
		disambiguatedCandidates := DisambiguateCandidates(currentCandidate, ambiguousKey)
		return FindValidCandidates(append(candidatesPerm[1:], disambiguatedCandidates...), signals)
	}

	// current candidate is disambiguated, check if all signals can be resolved to valid numbers
	if CandidateMapsAllSignalsToValidNumbers(currentCandidate, signals) {
		return currentCandidate
	} else {
		return FindValidCandidates(candidatesPerm[1:], signals)
	}
}

func CandidateMapsAllSignalsToValidNumbers(candidate SegMapping, signals []string) bool {
	for _, signal := range signals {
		if DecodeOutput(&candidate, signal) == "" {
			return false
		}
	}
	return true
}

func FirstAmbiguousMappingKey(candidate SegMapping) string {
	for key, mappings := range candidate {
		if len(mappings) > 1 {
			return key
		}
	}
	return ""
}

func DisambiguateCandidates(candidate SegMapping, ambiguousKey string) (disambiguatedCandidates []SegMapping) {
	for _, possibleMapping := range candidate[ambiguousKey] {
		disambiguatedCandidate := CopyMap(&candidate)
		disambiguatedCandidate[ambiguousKey] = string(possibleMapping)
		for k, v := range disambiguatedCandidate {
			if v == string(possibleMapping) {
				continue
			}
			disambiguatedCandidate[k] = Exclude(v, string(possibleMapping))
		}
		disambiguatedCandidates = append(disambiguatedCandidates, disambiguatedCandidate)
	}
	return disambiguatedCandidates
}

func CopyMap(inputMap *SegMapping) SegMapping {
	copyMap := make(SegMapping)
	for key, value := range *inputMap {
		copyMap[key] = value
	}
	return copyMap
}
func DecodeOutput(candidates *SegMapping, output string) (outputFigure string) {
	outputSegs := ""
	for _, char := range output {
		targetSegs := (*candidates)[string(char)]
		outputSegs = UnionDedupe(outputSegs, targetSegs)
	}
	for figure, segs := range segs {
		if outputSegs == segs {
			outputFigure = strconv.Itoa(figure)
		}
	}
	return
}

func RemoveCandidates(candidates *SegMapping, invertedSignal string, targetSegs string) {
	for _, seg := range invertedSignal {
		currentCandidatesForSeg := (*candidates)[string(seg)]
		(*candidates)[string(seg)] = Exclude(currentCandidatesForSeg, targetSegs)
	}
}

func ReduceToCandidates(candidates *SegMapping, signal string, targetSegs string) {
	for _, seg := range signal {
		currentCandidatesForSeg := (*candidates)[string(seg)]
		(*candidates)[string(seg)] = Intersection(currentCandidatesForSeg, targetSegs)
	}
}

func SegForLength(length int) string {
	switch length {
	case 2:
		return segs[1]
	case 3:
		return segs[7]
	case 4:
		return segs[4]
	case 7:
		return segs[8]
	default:
		return segs[8] // no clear indication on actual number, can be any connection
	}
}

func Exclude(a string, b string) (exclusion string) {
	for _, char := range a {
		if !strings.Contains(b, string(char)) {
			exclusion += string(char)
		}
	}
	return
}

func Intersection(a, b string) (intersection string) {
	for _, char := range a {
		if strings.Contains(b, string(char)) {
			intersection += string(char)
		}
	}
	return
}

func UnionDedupe(a, b string) string {
	unions := strings.Split(a+b, "")
	sort.Strings(unions)
	unions = DedupSorted(unions)
	return strings.Join(unions, "")
}

func DedupSorted(input []string) (dedup []string) {
	if len(input) <= 1 {
		return input
	}
	dedup = []string{input[0]}
	for i := 1; i < len(input); i++ {
		if input[i] != input[i-1] {
			dedup = append(dedup, input[i])
		}
	}
	return
}

func InitializeCandidates() SegMapping {
	candidates := make(SegMapping)
	for _, segment := range strings.Split(segs[8], "") {
		candidates[segment] = segs[8]
	}
	return candidates
}

func (t Day) InputFilename() string {
	return "day8/input"
}
