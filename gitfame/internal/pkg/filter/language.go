package filter

import (
	"encoding/json"
	"strings"

	"gitlab.com/slon/shad-go/gitfame/configs"
)

type Language struct {
	Name       string   `json:"name"`
	Type       string   `json:"type"`
	Extensions []string `json:"extensions"`
}

var nameToLanguage = make(map[string]Language)

func init() {
	var languages []Language
	err := json.Unmarshal(configs.LanguageExtentions, &languages)
	if err != nil {
		panic(err)
	}

	for _, lang := range languages {
		nameToLanguage[strings.ToLower(lang.Name)] = lang
	}
}
