package reader

import (
	"embed"
)

//go:embed resources/default.json
var defaultDictionaryFile embed.FS

//go:embed resources/no-markup.json
var noMarkupDictionaryFile embed.FS

func readDefaultDictionary() ([]byte, error) {
	return defaultDictionaryFile.ReadFile("resources/default.json")
}

func readNoMarkupDictionary() ([]byte, error) {
	return noMarkupDictionaryFile.ReadFile("resources/no-markup.json")
}

func ReadJsonDictionary(resourceName string) ([]byte, error) {
	if resourceName == "no-markup" {
		return readNoMarkupDictionary()
	}

	return readDefaultDictionary()
}
