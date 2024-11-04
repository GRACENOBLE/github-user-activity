package main

import (
	"fmt"

	"github.com/noble/github-user-activity/api"
)

func main() {
	data := api.FetchUserActivity("GRACENOBLE")
	fmt.Println(data)
}