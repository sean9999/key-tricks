package main

import "github.com/google/go-github/v55/github"

func (g GappyMachine) SigningKeys() ([]*github.SSHSigningKey, error) {
	keys, _, err := g.client.Users.ListSSHSigningKeys(g.ctx, "", nil)
	return keys, err
}

func (g GappyMachine) GPGKeys() ([]*github.GPGKey, error) {
	keys, _, err := g.client.Users.ListGPGKeys(g.ctx, "", nil)
	return keys, err
}

func (g GappyMachine) AuthenticationKeys() ([]*github.Key, error) {
	keys, _, err := g.client.Users.ListKeys(g.ctx, "", nil)
	return keys, err
}

func (g GappyMachine) DeleteAuthenticationKey(key github.Key) (*github.Response, error) {
	return g.client.Users.DeleteKey(g.ctx, key.GetID())
}
