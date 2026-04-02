package main

import (
	_ "embed"
	"encoding/json"
	"os"
	"strings"
)

//go:embed locales.json
var localesData []byte

type messages struct {
	Help         string `json:"help"`
	InvalidArgs  string `json:"invalidArgs"`
	InvalidBeats string `json:"invalidBeats"`
	InvalidSteps string `json:"invalidSteps"`
	InvalidShift string `json:"invalidShift"`
}

var locales map[string]messages

func init() {
	if err := json.Unmarshal(localesData, &locales); err != nil {
		panic("failed to parse locales.json: " + err.Error())
	}
}

func getMessages() messages {
	lang := os.Getenv("LANG")
	if lang == "" {
		lang = os.Getenv("LC_ALL")
	}
	if lang == "" {
		lang = os.Getenv("LC_MESSAGES")
	}

	// extract locale code from e.g. "en_US.UTF-8" → "en_US"
	code := strings.Split(lang, ".")[0]

	if m, ok := locales[code]; ok {
		return m
	}
	return locales["en_GB"]
}
