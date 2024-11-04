package main

import (
	"fmt"
	"github.com/noble/github-user-activity/api"
)

func main() {
	fmt.Print("github-user-activity: ")
	var username string
	fmt.Scanf("%s", &username)
	api.FetchUserActivity(username)
}