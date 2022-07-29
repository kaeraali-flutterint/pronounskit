package main

import (
	"github.com/alicekaerast/pronounskit/lib"
	"golang.org/x/oauth2"
	"log"
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

	m, err := lib.NewTokenManager(conf, "pronounkit.json")
	if err != nil {
		log.Fatalf("error getting token: %v", err)
	}

	zoomUser := lib.GetZoomPronouns(m)

	if len(os.Args) > 1 {
		newPronouns := os.Args[1]
		log.Println("Setting new pronouns to:", newPronouns)

		zoomUser.Pronouns = newPronouns
		lib.SetZoomPronouns(m, zoomUser)

	}
}
