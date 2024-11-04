package api

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/noble/github-user-activity/datatypes"
)

const apiUrl = "https://api.github.com/users/%s/events"

func FetchUserActivity(user string) []datatypes.UserActivity {
	res, err := http.Get(fmt.Sprintf(apiUrl, user))

	if err != nil {
		log.Fatal(err)
	}

	if res.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(res.Body)

		if err != nil{
			log.Fatal(err)
		}

		var userActivity []datatypes.UserActivity

		err = json.Unmarshal(bodyBytes, &userActivity)
		if err != nil{
			log.Fatal(err)
		}

		return userActivity

	}else{
		log.Fatal("Could not fetch data")
	}

	return []datatypes.UserActivity{}
}