package day8

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

var day = Day{}

func TestPart1(t *testing.T) {
	result := day.Part1(strings.Split(testInput, "\n"))
	if result != 26 {
		t.Error("digits in the output values should be 26 but were", result)
	}
}

func TestPart2(t *testing.T) {
	result := day.Part2(strings.Split(testInput, "\n"))
	if result != 61229 {
		t.Error("sum of all output values should be 61229 but was", result)
	}
}

func TestIntersection(t *testing.T) {
	assert.Equal(t, "de", Intersection("abcde", "defgh"))
	assert.Equal(t, "", Intersection("ab", "cd"))
	assert.Equal(t, "a", Intersection("a", "a"))
	assert.Equal(t, "", Intersection("", "a"))
	assert.Equal(t, "", Intersection("a", ""))
}

func TestExclude(t *testing.T) {
	assert.Equal(t, "", Exclude("", "a"))
	assert.Equal(t, "", Exclude("a", "a"))
	assert.Equal(t, "a", Exclude("a", ""))
	assert.Equal(t, "c", Exclude("abc", "ab"))
}

func TestUnion(t *testing.T) {
	assert.Equal(t, "a", UnionDedupe("a", ""))
	assert.Equal(t, "a", UnionDedupe("", "aa"))
	assert.Equal(t, "abcd", UnionDedupe("ab", "acd"))
}

func TestDisabmiguateCandidates(t *testing.T) {
	result := []SegMapping{{"a": "c"}, {"a": "d"}}
	assert.Equal(t, result, DisambiguateCandidates(SegMapping{"a": "cd"}, "a"))
}

func TestFindValidCandidates(t *testing.T) {
	var segmap SegMapping = SegMapping{"a": "cf",
		"b": "cf",
		"c": "eg",
		"d": "a",
		"e": "bd",
		"f": "bd",
		"g": "eg"}
	result := FindValidCandidates([]SegMapping{segmap}, []string{"cdfeb", "fcadb", "cdfeb", "cdbafe"})
	fmt.Println(result)
}

const testInputSingle = `acedgfb cdfbe gcdfa fbcad dab cefabd cdfgeb eafb cagedb ab | cdfeb fcadb cdfeb cdbafe

`

/*
cdfeb
candidates =
 0 = a -> cf
 1 = b -> cf
 2 = c -> eg
 3 = d -> a
 4 = e -> bd
 5 = f -> bd
 6 = g -> eg
*/

const testInput = `be cfbegad cbdgef fgaecd cgeb fdcge agebfd fecdb fabcd edb | fdgacbe cefdb cefbgd gcbe
edbfga begcd cbg gc gcadebf fbgde acbgfd abcde gfcbed gfec | fcgedb cgb dgebacf gc
fgaebd cg bdaec gdafb agbcfd gdcbef bgcad gfac gcb cdgabef | cg cg fdcagb cbg
fbegcd cbd adcefb dageb afcb bc aefdc ecdab fgdeca fcdbega | efabcd cedba gadfec cb
aecbfdg fbg gf bafeg dbefa fcge gcbea fcaegb dgceab fcbdga | gecf egdcabf bgf bfgea
fgeab ca afcebg bdacfeg cfaedg gcfdb baec bfadeg bafgc acf | gebdcfa ecba ca fadegcb
dbcfg fgd bdegcaf fgec aegbdf ecdfab fbedc dacgb gdcebf gf | cefg dcbef fcge gbcadfe
bdfegc cbegaf gecbf dfcage bdacg ed bedf ced adcbefg gebcd | ed bcgafe cdgba cbgef
egadfb cdbfeg cegd fecab cgb gbdefca cg fgcdab egfdb bfceg | gbdfcae bgc cg cgb
gcafb gcf dcaebfg ecagb gf abcdeg gaef cafbge fdbac fegbdc | fgae cfgab fg bagce

`
