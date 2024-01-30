package main

import (
	"fmt"
	"time"

	"github.com/google/go-github/v55/github"
)

const PERIOD = time.Duration(time.Hour * 24)
const GRACE = PERIOD * 1
const TOO_OLD = PERIOD * 500
const SEND_WARNING = PERIOD * 200

func (g GappyMachine) WarningEmail(user *github.User, key github.Key) (string, string, string) {
	recipient := user.GetEmail()
	subject := "Pending SSH key deletion"
	body := fmt.Sprintf("Dear %s,\n\nYour SSH key %q (%s) will be deleted soon.\n", *user.Name, key.GetTitle(), key.GetKey())
	return recipient, subject, body
}

func BusinessLogic(g GappyMachine, key github.Key) {

	keyAge := time.Since(key.GetCreatedAt().Time)

	switch {
	case keyAge > TOO_OLD+GRACE:
		g.DeleteAuthenticationKey(key)
	case keyAge > SEND_WARNING && keyAge < TOO_OLD:
		g.WarningEmail(g.me, key)
	}

}
