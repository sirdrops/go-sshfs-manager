package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

var filepath string = "config.json"

func checkFileCreateIfNotExist(path string) (*os.File, error) {
	f, err := os.Open(path)

	if err != nil {
		f, err := os.Create(path)
		if err != nil {
			return f, err
		}
	}
	return f, nil
}

func main() {

	f, err := checkFileCreateIfNotExist(filepath)
	if err != nil {
		fmt.Println("Could not open or create file")
	} else {
		fmt.Println("File is ready to go")
		defer f.Close()
	}
	byteValue, _ := io.ReadAll(f)
	var result map[string]interface{}
	json.Unmarshal([]byte(byteValue), &result)

	// fmt.Println(result["servers"])
	// servers := result["servers"].(map[string]interface{})
	servers := result["servers"].([]interface{})
	fmt.Println(servers)

	// b1 := make([]byte, 1000)
	// contents, err := f.Read(b1)
	// if err != nil {
	// 	fmt.Println("Cos zlego sie stanelo")
	// } else {
	// 	fmt.Printf("%s", string(b1[:contents]))
	// }
}
