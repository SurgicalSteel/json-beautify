package main

// MIT License
//
// Copyright (c) 2018 Yuwono Bangun Nagoro
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sync"
)

var wg sync.WaitGroup

// aimed to beautify your wrecked JSON
// to make it beauty like your crush and your ex

func main() {
	args := os.Args

	if len(args) <= 1 {
		fmt.Println("😠 JSON files must be defined!")
		return
	}
	fileNames := args[1:]

	for _, fn := range fileNames {
		wg.Add(1)
		go func(fileName string) {
			defer wg.Done()
			rawJsonFile, err := os.Open(fileName)
			if err != nil {
				fmt.Printf("😢 Error on opening file %s\n%s\n", fileName, err.Error())
				return
			}
			defer rawJsonFile.Close()
			rawJsonBody, err := ioutil.ReadAll(rawJsonFile)
			if err != nil {
				fmt.Printf("😢 Error on reading file %s\n%s\n", fileName, err.Error())
				return
			}
			jsonData, err := validateJSON(string(rawJsonBody))
			if err != nil {
				fmt.Printf("😢 Your JSON on file %s is invalid\n%s\n", fileName, err.Error())
				return
			}

			rawJsonData, err := json.MarshalIndent(jsonData, "", "\t")
			if err != nil {
				fmt.Printf("😢 Error occured while beautifying your JSON on file %s\n%s\n", fileName, err.Error())
				return
			}

			beautifiedPath := fmt.Sprintf("%s", fileName)
			beautifiedFile, err := os.Create(beautifiedPath)
			if err != nil {
				fmt.Printf("😢 Error occured while creating new file %s\n%s\n", fileName, err.Error())
				return
			}
			beautifiedFile.Write(rawJsonData)

			fmt.Printf("😉 Done beautifying your JSON on file %s\n", fileName)
			defer beautifiedFile.Close()
		}(fn)
	}

	wg.Wait()
}
func validateJSON(body string) (map[string]interface{}, error) {
	var temporaryJSONMap map[string]interface{}
	err := json.Unmarshal([]byte(body), &temporaryJSONMap)
	return temporaryJSONMap, err
}
