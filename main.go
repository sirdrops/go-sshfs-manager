package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

var filepath string = "config.json"

type Servers struct {
	Servers []Server `json:"servers"`
}
type Server struct {
	User    string `json:"user"`
	Address string `json:"address"`
}

var servers Servers

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
	// var result map[string]interface{}
	// json.Unmarshal([]byte(byteValue), &result)
	json.Unmarshal(byteValue, &servers)
	fmt.Println(len(servers.Servers))

	for key, value := range servers.Servers {
		fmt.Println(key, value)
	}
	// Trzeba bedzie zrobic strukta chyba jednak
	// servers := result["servers"].([]interface{})
	// fmt.Println(servers)
	// for key, value := range servers {
	// 	fmt.Println("key:", key, "value", value)
	// fmt.Println(servers[key])
	// }

}
