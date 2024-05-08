package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
)

var filepath string = "config.json"

// var logFile, _ = os.Create("my_log.log")
var logFile, _ = os.OpenFile("my_log.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
var mylog = log.New(logFile, "", log.Lshortfile|log.LstdFlags)

type Servers struct {
	Servers []Server `json:"servers"`
}
type Server struct {
	User    string `json:"user"`
	Address string `json:"address"`
}

var servers Servers

// Try to open file, if file doesn't exist try to create new one
func openFileCreateIfNotExist(path string) (*os.File, error) {
	// f, err := os.Open(path)
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		return f, err
		// fmt.Println("Could not open file, trying to create new one")
		// f, err := os.Create(path)
		// if err != nil {
		// 	return f, err
		// }
	}
	return f, nil
}

func main() {

	f, err := openFileCreateIfNotExist(filepath)
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
	mylog.Println(input)

	// fmt.Println(servers.Servers[0].Address)
	// fmt.Println(servers.Servers[input].Address)
	choice := fmt.Sprintf("%s@%s\n", servers.Servers[input].User, servers.Servers[input].Address)
	fmt.Println(choice)
	sshfsString := fmt.Sprintf("sudo sshfs -o allow_other,default_permission %s//home/%s/ /path/to_remote_folder/", strings.TrimRight(choice, "\n"), servers.Servers[input].User)

	// cmdStruct := exec.Command("echo", "Tutaj bedzie komenda do sshfs most likely")
	// cmdStruct := exec.Command("echo", choice)
	cmdStruct := exec.Command("pwd")
	out, err := cmdStruct.Output()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print(string(out))
	fmt.Println(sshfsString)
	// this will be probably a string combo to use sshfs hope it will work with remote keygen
	// "sudo", "sshfs", "-o", "allow_other,default_permission", choice + "//home/" + servers.Servers[input].User + "/" + "remote_server" <<< tutaj pewnie raczej jakas sciezka dodatkowa
}
