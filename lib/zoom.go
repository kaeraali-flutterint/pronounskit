package lib

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"golang.org/x/oauth2"
	"io"
	"log"
	"net/http"
)

func GetZoomPronouns(m *TokenManager) ZoomUser {
	ctx := context.Background()
	ts := m.TokenSource(ctx)

	client := oauth2.NewClient(ctx, ts)

	// use client.Get / client.Post for further requests, the token will automatically be there
	resp, err := client.Get("https://zoom.us/v2/users/me")
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(resp.Body)
	user := ZoomUser{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Current pronouns are:", user.Pronouns)
	return user
}

func SetZoomPronouns(m *TokenManager, u ZoomUser) {
	ctx := context.Background()
	ts := m.TokenSource(ctx)

	client := oauth2.NewClient(ctx, ts)

	jsonStr, err := json.Marshal(u)
	if err != nil {
		log.Fatal(err)
	}

	req, err := http.NewRequest("PATCH", "https://zoom.us/v2/users/me", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		log.Fatal(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	body, _ := io.ReadAll(resp.Body)
	fmt.Println(string(body))
}
