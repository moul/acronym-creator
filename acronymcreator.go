package actor

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

// ACTOR - Acronym CreaTOR :)

type Creator struct {
	columns [][]string
}

type Acronym struct {
	Acronym     string
	Combination []string
	//Score       float64
}

func New(columns [][]string) Creator {
	return Creator{
		columns: columns,
	}
}

func newAcronym(acronym string, combination []string) Acronym {
	return Acronym{
		Acronym:     acronym,
		Combination: combination,
	}
}

func (c *Creator) getCombinations() [][]string {
	// FIXME: create real list of combinations
	combination := []string{}
	for _, column := range c.columns {
		combination = append(combination, column[0])
	}
	return [][]string{combination}
}

func (c *Creator) getDictionary() ([]string, error) {
	dictionary, err := os.Open("/usr/share/dict/words")
	if err != nil {
		return nil, err
	}
	defer dictionary.Close()
	var lines []string
	scanner := bufio.NewScanner(dictionary)
	for scanner.Scan() {
		lines = append(lines, stripCtlFromBytes(scanner.Text()))
	}
	return lines, nil
}

func (c *Creator) CreateAcronyms() (map[string][]Acronym, error) {
	acronyms := make(map[string][]Acronym, 0)

	for _, combination := range c.getCombinations() {
		acronym := ""
		query := ""
		for _, word := range combination {
			acronym += string(word[0])
			query += string(word[0]) + ".*"
		}
		if _, found := acronyms[acronym]; !found {
			acronyms[acronym] = make([]Acronym, 0)
		}
		acronyms[acronym] = append(acronyms[acronym], newAcronym(acronym, combination))

		dict, err := c.getDictionary()
		if err != nil {
			return nil, err
		}

		re := regexp.MustCompile("^" + query + "$")
		for _, line := range dict {
			if re.MatchString(line) {
				fmt.Println(line, combination)
				// FIXME: iterate over letters, check the index in the combination, put in uppercase, and count the amount of lowercase, if the percent is enough, add the word !
			}
		}
	}

	return acronyms, nil
}
