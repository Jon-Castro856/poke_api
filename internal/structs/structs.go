package structs

type CliCommand struct {
	Name        string
	Description string
	Callback    func(Name string, Cfg *Config) (Config, error)
	Cfg         Config
}

type Config struct {
	Back    string
	Forward string
}

type MapData struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}
