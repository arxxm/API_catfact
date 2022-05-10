package catfacts

type Breed struct {
	Breed   string
	Country string
	Origin  string
	Coat    string
	Pattern string
}

type CatFact struct {
	Fact   string `json:"fact"`
	Length int    `json:"length"`
}

type pagination struct {
	Total       int     `json:"total"`
	PerPage     string  `json:"per_page"`
	CurrentPage int     `json:"current_page"`
	LastPage    int     `json:"last_page"`
	From        int     `json:"from"`
	To          int     `json:"to"`
	NextPageURL *string `json:"next_page_url"`
	PrevPageURL *string `json:"prev_page_url"`
}

type breeds struct {
	pagination
	Data []Breed `json:"data"`
}

type facts struct {
	pagination
	Data []CatFact `json:"data"`
}
