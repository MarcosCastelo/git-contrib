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
	var email string
	flag.StringVar(&folder, "add", "", "add a new folder to scan for Git repositories")
	flag.StringVar(&email, "email", "your@email.com", "the email to scan")
	flag.Parse()

	if folder != "" {
		utils.Scan(folder)
		endingTime := time.Now().UTC()
		fmt.Println(endingTime.Sub(startingTime))
		return
	}

	utils.Stats(email)
	endingTime := time.Now().UTC()
	fmt.Println(endingTime.Sub(startingTime))
}
