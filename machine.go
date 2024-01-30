package main

import (
	"context"

	"github.com/google/go-github/v55/github"
)

type Gappy interface {
	AuthenticationKeys() ([]*github.Key, error)
	GPGKeys() ([]*github.GPGKey, error)
	SigningKeys() ([]*github.SSHSigningKey, error)
	Gists() ([]*github.Gist, error)
	Self() (*github.User, error)
}

type GappyMachine struct {
	client *github.Client
	ctx    context.Context
	me     *github.User
}

var ZeroMachine GappyMachine = GappyMachine{}

func NewGitHubMachine(authToken string) (GappyMachine, error) {
	m := GappyMachine{
		client: github.NewClient(nil).WithAuthToken(authToken),
		ctx:    context.Background(),
	}
	me, err := m.Self()
	if err != nil {
		return ZeroMachine, err
	}
	m.me = me
	return m, nil
}
