package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
)

/* Taking JSON input, find all numbers and sum them.
 * Can assume that there are no sttings that include numbers
 * The very naive solution is just to a regex to find all numbers
 */

var inputFile = flag.String("inputFile", "inputs/day12.input", "Relative file path to use as input.")

func main() {
	flag.Parse()
	bytes, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		return
	}
	contents := string(bytes)

	regex := regexp.MustCompile("(-?\\d+)")
	matches := regex.FindAllString(contents, -1)
	sum := 0
	for i := range matches {
		val, _ := strconv.Atoi(matches[i])
		sum += val
	}

	fmt.Println(sum)

	// Part B requires excluding objects that have a property with value "red"
	var result interface{}
	json.Unmarshal(bytes, &result)
	switch data := result.(type) {
	case map[string]interface{}:
		fmt.Println(sumObject(data))
	case []interface{}:
		fmt.Println(sumList(data))
	default:
		fmt.Println(":shrug:")
	}
}

func sumList(list []interface{}) int {
	total := 0

	for i := range list {
		switch val := list[i].(type) {
		case int:
			total += val
		case int64:
			total += int(val)
		case float64:
			total += int(val)
		case []interface{}:
			total += sumList(val)
		case map[string]interface{}:
			total += sumObject(val)
		}
	}
	return total
}

func sumObject(obj map[string]interface{}) int {
	total := 0
	for _, v := range obj {
		switch val := v.(type) {
		case int:
			total += val
		case int64:
			total += int(val)
		case float64:
			total += int(val)
		case string:
			if val == "red" {
				return 0
			}
		case []interface{}:
			total += sumList(val)
		case map[string]interface{}:
			total += sumObject(val)
		}
	}
	return total

}
