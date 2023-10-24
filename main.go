package main

import (
	"flag"
	"fmt"
	"git-contrib/utils"
	"time"
)

func main() {
	startingTime := time.Now().UTC()

	var folder string
	var name string
	flag.StringVar(&folder, "add", "", "add a new folder to scan for Git repositories")
	flag.StringVar(&name, "name", "username", "the username to scan")
	flag.Parse()

	if folder != "" {
		utils.Scan(folder)
		endingTime := time.Now().UTC()
		fmt.Println(endingTime.Sub(startingTime))
		return
	}

	utils.Stats(name)
	endingTime := time.Now().UTC()
	fmt.Println(endingTime.Sub(startingTime))
}
