package dictionary

import (
	"encoding/json"

	"github.com/stscoundrel/old-norse-dictionary-go/internal/reader"
)

type DictionaryEntry struct {
	Headword    string   `json:"a"`
	Definitions []string `json:"b"`
}

const (
	DICTIONARY_NAME           = "default"
	NO_MARKUP_DICTIONARY_NAME = "no-markup"
)

func parseDictionary(bytes []byte) ([]DictionaryEntry, error) {
	var entries []DictionaryEntry

	err := json.Unmarshal(bytes, &entries)

	return entries, err
}

func readDictionary(dictionaryName string) ([]DictionaryEntry, error) {
	rawDictionary, readError := reader.ReadJsonDictionary(dictionaryName)

	if readError != nil {
		return []DictionaryEntry{}, readError
	}

	dictionary, parseError := parseDictionary(rawDictionary)

	if parseError != nil {
		return []DictionaryEntry{}, parseError
	}

	return dictionary, nil
}

func GetDictionary() ([]DictionaryEntry, error) {
	return readDictionary(DICTIONARY_NAME)
}

func GetNoMarkupDictionary() ([]DictionaryEntry, error) {
	return readDictionary(NO_MARKUP_DICTIONARY_NAME)

}
