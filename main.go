package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/exec"
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
		fmt.Println("Could not open file, trying to create new one")
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
		fmt.Println("Could not create file")
	} else {
		fmt.Println("File is ready to go")
		defer f.Close()
	}
	byteValue, _ := io.ReadAll(f)
	json.Unmarshal(byteValue, &servers)

	for key, value := range servers.Servers {
		// fmt.Println(key, value)
		fmt.Printf("%d %s\n", key, value.Address)
	}
	var input int
	fmt.Print("Please enter number of server: ")
	fmt.Scan(&input)

	// fmt.Println(servers.Servers[0].Address)
	// fmt.Println(servers.Servers[input].Address)
	choice := fmt.Sprintf("%s@%s\n", servers.Servers[input].User, servers.Servers[input].Address)
	fmt.Println(choice)

	cmdStruct := exec.Command("echo", "hello")
	out, err := cmdStruct.Output()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print(string(out))
	// this will be probably a string combo to use sshfs hope it will work with remote keygen
	// "sudo", "sshfs", "-o", "allow_other,default_permission", choice + "//home/" + servers.Servers[input].User + "/"
}
