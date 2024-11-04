package api

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

const apiUrl = "https://api.github.com/users/%s/events"

func FetchUserActivity(user string) {
	res, err := http.Get(fmt.Sprintf(apiUrl, user))

	if err != nil {
		log.Fatal(err)
	}

	if res.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(res.Body)

		if err != nil{
			log.Fatal(err)
		}

		json := string(bodyBytes)

		fmt.Printf("%v", json)

	}else{
		log.Fatal("Could not fetch data")
	}
}