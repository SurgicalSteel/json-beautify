package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

// aimed to beautify your wrecked JSON
// to make it beauty like your crush and your ex

func main() {
	fileFlag := flag.String("f", "", "Define JSON file path that you want to beautify")
	flag.Parse()
	rawJsonFile, err := os.Open(*fileFlag)
	if err != nil {
		fmt.Println("😢 Error on opening file :", err.Error())
		return
	}
	defer rawJsonFile.Close()
	rawJsonBody, err := ioutil.ReadAll(rawJsonFile)
	if err != nil {
		fmt.Println("😢 Error on reading file :", err.Error())
		return
	}
	jsonData, err := validateJSON(string(rawJsonBody))
	if err != nil {
		fmt.Println("😢 Your JSON is invalid", err.Error())
		return
	}

	rawJsonData, err := json.MarshalIndent(jsonData, "", "\t")
	if err != nil {
		fmt.Println("😢 Error occured while beautifying your JSON", err.Error())
		return
	}

	beautifiedPath := fmt.Sprintf("beautified-%s", *fileFlag)
	beautifiedFile, err := os.Create(beautifiedPath)
	if err != nil {
		fmt.Println("😢 Error occured while creating new file", err.Error())
		return
	}
	beautifiedFile.Write(rawJsonData)

	fmt.Println("😉 Done beautifying your JSON")
	defer beautifiedFile.Close()
}
func validateJSON(body string) (map[string]interface{}, error) {
	var temporaryJSONMap map[string]interface{}
	err := json.Unmarshal([]byte(body), &temporaryJSONMap)
	return temporaryJSONMap, err
}
