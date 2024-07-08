package main

import (
	"encoding/json"
	"fmt"
	"time"
)

func main() {
	data := map[string]interface{}{
		"intValue":    1234,
		"boolValue":   true,
		"stringValue": "hello!",
		"dateValue":   time.Date(2024, 6, 29, 10, 35, 0, 0, time.UTC),
		"objectValue": map[string]interface{}{
			"arrayValue": []int{1, 2, 3, 4},
		},
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("could not marshal json: %s\n", err)
		return
	}

	//fmt.Printf("json data: %s\n", jsonData) will do the same work as the two line of codes becaise the %s verb converts the byte response into strings.
	stringData := string(jsonData)
	fmt.Println(stringData)
}
