package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	//"github.com/dbysani/dynamic-rest-api/mappingfiles"
)

// set of fields that are key value pairs
type inputJson struct {
	Fields []map[string]string
}

// set of field mapping definitions
type mappingFile struct {
	Fields []map[string]string
}

func loadJsonInput(filename string) map[string]any {

	jsonStr, err := ioutil.ReadFile(filename)

	if err != nil {
		log.Fatal("Unable to read input json", err)
	}

	var result map[string]any
	json.Unmarshal([]byte(jsonStr), &result)

	fields := result["fields"].(map[string]any)
	return fields
}

func mapAPIFields(mappingFileName string) []map[string]string {
	content, err := ioutil.ReadFile(mappingFileName)
	if err != nil {
		log.Fatal("Error opening file", err)
	}

	var amappingFile mappingFile
	err = json.Unmarshal(content, &amappingFile)

	if err != nil {
		log.Fatal("Error during unmarshalling mapping file")
	}

	return amappingFile.Fields
}

func transformInput(jsonInputFileName string, mappingFileName string) {

	inputFields := loadJsonInput(jsonInputFileName)
	mapp := mapAPIFields(mappingFileName)

	transformedFields := make(map[string]any)

	for k, v := range inputFields {

		for _, mappingItem := range mapp {
			if mappingItem["sourceName"] == k {
				transformedFields[mappingItem["destinationName"]] = v
			}
		}
	}

	fmt.Println("input fields", inputFields)

	fmt.Println("transformed fields", transformedFields)
}

func main() {

	fmt.Println("---Transforming input1 with mapping file V1: \n")
	transformInput("./input1.json", "./mappingfiles/mapping-V1.json")

	fmt.Println("\n---Transforming input2 with mapping file V1.1: \n")
	transformInput("./input2.json", "./mappingfiles/mapping-V1.1.json")

	fmt.Println("\n---Backward compatibility with new mapping file: \n")
	transformInput("./input1.json", "./mappingfiles/mapping-V1.1.json")
}
