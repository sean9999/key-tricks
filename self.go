package main

import "github.com/google/go-github/v55/github"

func (g GappyMachine) Self() (*github.User, error) {
	record, _, err := g.client.Users.Get(g.ctx, "")
	return record, err
}
