package account

type Account struct {
	Puuid    string `json:"puuid"`
	GameName string `json:"gameName"`
	TagLine  string `json:"tagLine"`
}

func (Account) IsGetResponse() {}

type Region struct {
	Puuid  string `json:"puuid"`
	Game   string `json:"game"`
	Region string `json:"region"`
}

func (Region) IsGetResponse() {}