package main

import "github.com/google/go-github/v55/github"

func (g GappyMachine) Gists() ([]*github.Gist, error) {
	records, _, err := g.client.Gists.ListAll(g.ctx, nil)
	return records, err
}
