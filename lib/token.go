// Largely copied from https://gist.github.com/guumaster/2c7f48ac3567ae6c456f4020c857c375

package lib

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"

	oauth2ns "github.com/nmrshll/oauth2-noserver"
	"golang.org/x/oauth2"
)

var (
	TokenNotFoundErr = errors.New("token not found")
	TokenWebErr      = errors.New("error getting token from web")
	TokenOpenErr     = errors.New("error opening token file")
	TokenSaveErr     = errors.New("error saving token file")
)

type TokenManager struct {
	conf                *oauth2.Config
	token               *oauth2.Token
	originalAccessToken string
	filepath            string
}

func NewTokenManager(conf *oauth2.Config, filepath string) (*TokenManager, error) {
	t := &TokenManager{
		conf:     conf,
		filepath: filepath,
	}

	isNewToken := false
	tok, err := t.getFromFile()
	if errors.Is(err, TokenOpenErr) || errors.Is(err, TokenNotFoundErr) {
		tok, err = t.getFromWeb()
		if err != nil {
			return nil, fmt.Errorf("error getting token from web: %w", err)
		}
		isNewToken = true
	}
	if err != nil {
		return nil, fmt.Errorf("error getting token: %w", err)
	}

	t.token = tok
	t.originalAccessToken = tok.AccessToken

	// This will refresh the token when needed
	ts := t.TokenSource(context.Background())
	newTok, err := ts.Token()
	if err != nil {
		return nil, err
	}

	tokenRefreshed := tok.AccessToken != newTok.AccessToken

	if isNewToken || tokenRefreshed {
		t.token = newTok
		err = t.save()
		if err != nil {
			return nil, err
		}
	}

	return t, nil
}

func (t *TokenManager) TokenSource(ctx context.Context) oauth2.TokenSource {
	return t.conf.TokenSource(ctx, t.token)
}

// Save stores the token in json file.
func (t *TokenManager) save() error {
	fmt.Printf("Saving token to file: %s\n", t.filepath)
	f, err := os.OpenFile(t.filepath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return fmt.Errorf("%q: %w", err, TokenSaveErr)
	}

	defer f.Close()
	enc := json.NewEncoder(f)
	enc.SetIndent("", "  ")
	err = enc.Encode(t.token)
	if err != nil {
		return fmt.Errorf("%q: %w", err, TokenSaveErr)
	}

	return nil
}

// getFromFile retrieves a token from a local file.
func (t *TokenManager) getFromFile() (*oauth2.Token, error) {
	f, err := os.Open(t.filepath)
	if err != nil {
		return nil, fmt.Errorf("%q: %w", err, TokenOpenErr)
	}
	defer f.Close()
	tok := new(oauth2.Token)
	err = json.NewDecoder(f).Decode(tok)
	if err != nil {
		return nil, fmt.Errorf("%q: %w", err, TokenOpenErr)
	}

	t.originalAccessToken = tok.AccessToken

	return tok, err
}

// getFromWeb Starts a local server and the oauth flow
func (t *TokenManager) getFromWeb() (*oauth2.Token, error) {
	client, err := oauth2ns.AuthenticateUser(t.conf)
	if err != nil {
		return nil, fmt.Errorf("%q: %w", err, TokenWebErr)
	}
	return client.Token, nil
}
