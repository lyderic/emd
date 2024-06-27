package cmd

type Configuration struct {
	Style    string `yaml:"style"`
	WordWrap int    `yaml:"wrap"`
}

var c = Configuration{
	Style:    "dark",
	WordWrap: 80,
}
