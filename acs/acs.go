package acs

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type API struct {
	Endpoint string
	APIToken string
}

func (a *API) Lookup(card string) (Identity, error) {
	u, err := url.Parse(a.Endpoint)
	if err != nil {
		return Identity{}, fmt.Errorf("bad endpoint: %s: %w", a.Endpoint, err)
	}
	u.Path = "/api/v1/permissions"

	body, err := json.Marshal(struct {
		Token string `json:"api_token"`
		Card  string `json:"card_id"`
	}{
		Token: a.APIToken,
		Card:  card,
	})
	if err != nil {
		return Identity{}, fmt.Errorf("unable to marshal json: %w", err)
	}

	resp, err := http.Post(u.String(), "application/json", bytes.NewReader(body))
	if err != nil {
		return Identity{}, fmt.Errorf("could not request identity: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return Identity{}, fmt.Errorf("unexpected status-code from acs: %d", resp.StatusCode)
	}

	id := &Identity{}
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(id)
	if err != nil {
		return Identity{}, fmt.Errorf("unable to parse acs response: %w", err)
	}

	return *id, nil
}
