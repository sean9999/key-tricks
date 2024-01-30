package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"time"
)

var (
	//go:embed auth_token.private.txt
	authToken string
)

func PrettyPrint(thing any) string {
	oo, _ := json.MarshalIndent(thing, "", "\t")
	return string(oo)
}

func barfOn(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	gh, err := NewGitHubMachine(authToken)
	barfOn(err)
	keys, err := gh.AuthenticationKeys()
	barfOn(err)

	for _, key := range keys {
		ageInDays := int(time.Since(key.GetCreatedAt().Time).Hours() / 24)
		fmt.Printf("\nthe key %q was created %d days ago\n", key.GetTitle(), ageInDays)
		if ageInDays >= 0 {
			//resp, err := gh.DeleteAuthenticationKey(*key)
			//fmt.Println(resp, err)

			fmt.Println("I would like to delete the key ", key.GetTitle(), key.GetID(), *key.Key)
		}

	}

	fmt.Println(err)

}
