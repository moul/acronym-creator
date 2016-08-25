package actor

// ACTOR - Acronym CreaTOR :)

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

type Creator struct {
	columns [][]string
}

type Acronym struct {
	Acronym     string
	Combination string
	Highlighted string
	Score       float64
}

func New(columns [][]string) Creator {
	return Creator{
		columns: columns,
	}
}

func newAcronym(acronym, combination string) Acronym {
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

func getMatchingScore(a, b string) (int, string) {
	score := 0
	c := string(a)
	minJ := 0
	for i := 0; i < len(c); i++ {
		found := false
		for j := minJ; j < len(b); j++ {
			if c[i] == b[j] {
				c = c[:i] + strings.ToUpper(fmt.Sprintf("%c", c[i])) + c[i+1:]
				found = true
				minJ = j
				break
			}
		}
		if !found {
			score--
		}
	}
	return score, c
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
		acronyms[acronym] = append(acronyms[acronym], newAcronym(acronym, strings.Join(combination, " ")))

		dict, err := c.getDictionary()
		if err != nil {
			return nil, err
		}

		joinedCombination := strings.Join(combination, " ")
		re := regexp.MustCompile("^" + query + "$")
		for _, line := range dict {
			if !re.MatchString(line) {
				continue
			}
			score, highlighted := getMatchingScore(line, joinedCombination)
			if score > -3 {
				acr := newAcronym(acronym, strings.Join(combination, " "))
				acr.Highlighted = highlighted
				acr.Score = float64(score)
				acronyms[acronym] = append(acronyms[acronym], acr)
			}

			// fmt.Printf("%s\r\t\t\t%s\t\t%s\r\t\t\t\t\t\t\t\t\t%d\n", line, joinedCombination, highlighted, score)
			// FIXME: iterate over letters, check the index in the combination, put in uppercase, and count the amount of lowercase, if the percent is enough, add the word !
		}
	}

	return acronyms, nil
}
