package main

import (
	"encoding/json"
	"log"
	"os"
)

type iFinalRes interface {
	writeJSON()
}

type FinalResult struct {
	FinalRes string
}

func finalResult() iFinalRes {
	charCount := charCounterResult()
	locNames := episodeLocationsResult()
	res := []interface{}{charCount, locNames}
	dataBytes, err := json.Marshal(res)
	if err != nil {
		log.Fatalf(err.Error())
	}

	return &FinalResult{FinalRes: string(dataBytes)}
}

func (f *FinalResult) writeJSON() {
	// return os.WriteFile("result", data, 0644)
	file, _ := os.Create("result.json")
	defer file.Close()

	_, err := file.WriteString(f.FinalRes)
	if err != nil {
		log.Fatalf(err.Error())
	}
}