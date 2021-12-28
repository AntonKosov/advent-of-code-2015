package main

import (
	"encoding/json"
	"fmt"

	"github.com/AntonKosov/advent-of-code-2015/aoc"
)

func main() {
	data := read()
	r := process(data)
	fmt.Printf("Answer: %v\n", r)
}

func read() (input string) {
	return aoc.ReadAllInput()[0]
}

func process(input string) int {
	parsed := parse(input)
	return int(objectSum(parsed))
}

func objectSum(node map[string]interface{}) float64 {
	var result float64
	for _, v := range node {
		switch v := v.(type) {
		case []interface{}:
			result += arraySum(v)
		case map[string]interface{}:
			result += objectSum(v)
		case string:
			if v == "red" {
				return 0
			}
		case float64:
			result += v
		default:
			panic("Unexpected type")
		}
	}
	return result
}

func arraySum(arr []interface{}) float64 {
	var result float64
	for _, v := range arr {
		switch v := v.(type) {
		case float64:
			result += v
		case string:
			// ignore all string values in arrays
		case map[string]interface{}:
			result += objectSum(v)
		case []interface{}:
			result += arraySum(v)
		default:
			panic("Unexpected type")
		}
	}

	return result
}

func parse(input string) map[string]interface{} {
	input = fmt.Sprintf(`{"input":%v}`, input)
	parsed := map[string]interface{}{}
	json.Unmarshal([]byte(input), &parsed)

	return parsed
}
