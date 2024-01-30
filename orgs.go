package main

import "github.com/google/go-github/v55/github"

func (gh GappyMachine) Organizations(user string, opts *github.ListOptions) ([]*github.Organization, error) {
	orgs, _, err := gh.client.Organizations.List(gh.ctx, user, opts)
	return orgs, err
}
