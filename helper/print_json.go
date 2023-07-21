package helper

import (
	"encoding/json"
	"fmt"
)

func PrintJson(v interface{}) {
	empJSON, _ := json.MarshalIndent(v, "", "  ")
	fmt.Print(string(empJSON))
}
