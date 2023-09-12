package main

import (
	"fmt"
	"os"

	yaml "gopkg.in/yaml.v3"
)

func main() {
	file, err := os.ReadFile("values.yaml")
	if err != nil {
		panic(err)
	}
	body := make(map[string]interface{})
	err = yaml.Unmarshal(file, body)
	if err != nil {
		panic(err)
	}

	processData(body)
}

func processData(data map[string]interface{}) {
	for key, value := range data {
		dataValue, ok := data[key]
		if !ok {
			continue
		}
		if key == "image" {
			if _, ok := data[key].(map[string]interface{}); ok {
				nestedData, ok := dataValue.(map[string]interface{})
				if !ok {
					continue
				}
				reg, ok := nestedData["registry"]
				if !ok {
					continue
				}
				repo, ok := nestedData["repository"]
				if !ok {
					continue
				}
				tag, ok := nestedData["tag"]
				if !ok {
					continue
				}

				fmt.Printf("%s/%s:%s", reg, repo, tag)
				return
			}
		} else {
			// Determine the type of the value in the schema
			switch value.(type) {
			case map[string]interface{}:
				// Nested object, recursively process
				if nestedData, ok := dataValue.(map[string]interface{}); ok {
					processData(nestedData)
				} else {
				}
			case []interface{}:
				// Array of values, process each element
				if arrayData, ok := dataValue.([]interface{}); ok {
					for _, element := range arrayData {
						if nestedData, ok := element.(map[string]interface{}); ok {
							processData(nestedData)
						} else {
						}
					}
				}
			default:
				// Leaf node, do something with the data
			}
		}
	}
}
