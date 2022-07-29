package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/alicekaerast/pronounskit/lib"
	"golang.org/x/oauth2"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	clientID := os.Getenv("ZOOM_CLIENT")
	clientSecret := os.Getenv("ZOOM_SECRET")
	conf := &oauth2.Config{
		ClientID:     clientID,     // also known as client key sometimes
		ClientSecret: clientSecret, // also known as secret key
		Scopes:       []string{"account"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://zoom.us/oauth/authorize",
			TokenURL: "https://zoom.us/oauth/token",
		},
	}

	client, err := lib.AuthenticateUser(conf)
	if err != nil {
		log.Fatal(err)
	}

	// use client.Get / client.Post for further requests, the token will automatically be there
	resp, err := client.Get("https://zoom.us/v2/users/me")
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(resp.Body)
	user := lib.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Current pronouns are:", user.Pronouns)

	if len(os.Args) > 1 {
		newPronouns := os.Args[1]
		log.Println("Setting new pronouns to:", newPronouns)

		user.Pronouns = newPronouns

		jsonStr, err := json.Marshal(user)
		if err != nil {
			log.Fatal(err)
		}

		req, err := http.NewRequest("PATCH", "https://zoom.us/v2/users/me", bytes.NewBuffer(jsonStr))
		req.Header.Set("Content-Type", "application/json")

		if err != nil {
			log.Fatal(err)
		}
		resp, err = client.Do(req)
		if err != nil {
			log.Fatal(err)
		}
		body, _ = io.ReadAll(resp.Body)
		fmt.Println(string(body))
	}
}
