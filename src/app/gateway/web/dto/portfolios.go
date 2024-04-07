package dto

type CreatePortfolio struct {
	Name string `json:"name"`
	From string `json:"from"`
	To   string `json:"to"`
}

type UpdatePortfolio struct {
	Name string `json:"name"`
	From string `json:"from"`
	To   string `json:"to"`
}

type PortfolioPartial struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	From string `json:"from"`
	To   string `json:"to"`
}

type PortfolioFull struct {
	Id      string   `json:"id"`
	Name    string   `json:"name"`
	From    string   `json:"from"`
	To      string   `json:"to"`
	Sectors []Sector `json:"sectors"`
	Groups  []Group  `json:"groups"`
	Asset   []Asset  `json:"assets"`
}

type Sector struct {
}

type Group struct {
}

type Asset struct {
}
