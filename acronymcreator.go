package actor

// ACTOR - Acronym CreaTOR :)

type Creator struct {
	columns [][]string
}

type Acronym struct {
}

func New(columns [][]string) Creator {
	return Creator{
		columns: columns,
	}
}

func (c *Creator) CreateAcronyms() []Acronym {
	acronyms := []Acronym{}
	return acronyms
}
