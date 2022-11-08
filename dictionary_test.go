package dictionary

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDictionariesHaveExpectedAmountOfEntries(t *testing.T) {
	expectedEntries := 35207
	defaultResult, _ := GetDictionary()
	noMarkupResult, _ := GetNoMarkupDictionary()

	assert.Equal(t, expectedEntries, len(defaultResult))
	assert.Equal(t, expectedEntries, len(noMarkupResult))
	assert.Equal(t, len(defaultResult), len(noMarkupResult))
}

func TestDefaultDictionaryHasExpectedEntries(t *testing.T) {
	result, _ := GetDictionary()

	expected1 := DictionaryEntry{
		Headword: "af-kymi",
		Definitions: []string{
			"a, m. <i>nook</i>, Ísl. ii. 471 (paper MS.); kymi, id., is now freq.",
		},
	}

	expected2 := DictionaryEntry{
		Headword: "át-frekr",
		Definitions: []string{
			"adj. <i>greedy, voracious,</i> Hkv. 2. 41.",
		},
	}

	expected3 := DictionaryEntry{
		Headword: "tor-velda",
		Definitions: []string{
			"<strong>1.</strong> d, <i>to make difficulties;</i> t. e-t fyrir e-m, Ld. 238.",
			"<strong>2.</strong> u, f. <i>a difficulty,</i> Rb. 336. <strong>torvelda-laust,</strong> adj. <i>without difficulties,</i> Bs. i. 307.",
		},
	}

	assert.Equal(t, expected1, result[100])
	assert.Equal(t, expected2, result[1989])
	assert.Equal(t, expected3, result[30000])
}

func TestNoMarkupDictionaryHasExpectedEntries(t *testing.T) {
	result, _ := GetNoMarkupDictionary()

	expected1 := DictionaryEntry{
		Headword: "af-kymi",
		Definitions: []string{
			"a, m. nook, Ísl. ii. 471 (paper MS.); kymi, id., is now freq.",
		},
	}

	expected2 := DictionaryEntry{
		Headword: "át-frekr",
		Definitions: []string{
			"adj. greedy, voracious, Hkv. 2. 41.",
		},
	}

	expected3 := DictionaryEntry{
		Headword: "tor-velda",
		Definitions: []string{
			"1. d, to make difficulties; t. e-t fyrir e-m, Ld. 238.",
			"2. u, f. a difficulty, Rb. 336. torvelda-laust, adj. without difficulties, Bs. i. 307.",
		},
	}

	assert.Equal(t, expected1, result[100])
	assert.Equal(t, expected2, result[1989])
	assert.Equal(t, expected3, result[30000])
}
