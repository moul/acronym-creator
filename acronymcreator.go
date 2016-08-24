package actor

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

func (c *Creator) CreateAcronyms() map[string][]Acronym {
	acronyms := make(map[string][]Acronym, 0)

	for _, combination := range c.getCombinations() {
		acronym := ""
		for _, word := range combination {
			acronym += string(word[0])
		}
		if _, found := acronyms[acronym]; !found {
			acronyms[acronym] = make([]Acronym, 0)
		}
		acronyms[acronym] = append(acronyms[acronym], newAcronym(acronym, combination))
	}

	return acronyms
}
