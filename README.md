# Old Norse Dictionary

Old Norse to English Dictionary for Golang. The dictionary consists of 35 000+ Old Norse words with English translations.

Based on the classic dictionary by Richard Cleasby and Gudbrand Vigfusson. If you find this one too abbreviated, academic or hard to read, you might want to check out [A Concise Dictionary of Old Icelandic](https://github.com/stscoundrel/old-icelandic-dictionary-go)

### Install

`go get github.com/stscoundrel/old-norse-dictionary-go`

### Usage

The dictionary comes in two variants:
- Default dictionary has HTML markup `<i>` and `<b>` to match look of the original book.
- No-markup version has the same content without any additional formatting tags.


```go
package main

import (
    "fmt"

    dictionary "github.com/stscoundrel/old-norse-dictionary-go"
)

func main() {
  defaultDictionary, defaltErr := dictionary.GetDictionary()
  noMarkupDictionary, noMarkupErr := dictionary.GetNoMarkupDictionary()
  
  // Error handling as you please.
  if defaltErr != nil {
    fmt.Println(err)
  }
  
  // Contains 35 000+ DictionaryEntries.
  for _, entry := range defaultDictionary {
    fmt.Println(entry.Headword)
  }
  
  // Headwords wont differ in dictionaries.
  fmt.Println(defaultDictionary[1989].Headword)  // át-frekr
  fmt.Println(noMarkupDictionary[1989].Headword) // át-frekr
  
  // But definitions markup will differ
  fmt.Println(defaultDictionary[1989].Definitions[0])  // adj. <i>greedy, voracious,</i> Hkv. 2. 41.
  fmt.Println(noMarkupDictionary[1989].Definitions[0]) // adj. greedy, voracious, Hkv. 2. 41.
}
```

The entries are structs of:

```go
type DictionaryEntry struct {
  Headword          string
  Definitions       []string
}

```

### About Cleasby & Vigfusson Dictionary

"Icelandic-English" dictionary was started by Richard Cleasby and finished by Gudbrand Vigfusson. It was published in 1874, which leads to there being many public domain versions of the book available.
