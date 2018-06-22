package main

import (
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"time"

	"github.com/tbruyelle/hipchat-go/hipchat"
)

type User struct {
	Name       string
	LastActive time.Time
}

func main() {
	apikeyPtr := flag.String("apikey", "", "The API key to access HipChat (Required)")
	flag.Parse()

	if *apikeyPtr == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	c := hipchat.NewClient(*apikeyPtr)
	opt := &hipchat.UserListOptions{}
	user, _, err := c.User.List(opt)
	if err != nil {
		panic(err)
	}

	allUser := []User{}
	for _, user := range user {
		u, _, err := c.User.View(strconv.Itoa(user.ID))
		if err != nil {
			panic(err)
		}
		t, err := time.Parse("2006-01-02T15:04:05+0000", u.LastActive)

		if err != nil {
			t, _ = time.Parse("2006-01-02T15:04:05.000Z", "1990-11-12T11:45:26.371Z")
		}
		allUser = append(allUser, User{Name: u.Name, LastActive: t})
	}

	// Sort the list of users by their last activity
	sort.Slice(allUser, func(i, j int) bool { return allUser[i].LastActive.Before(allUser[j].LastActive) })

	fmt.Printf("Found %d user\n", len(allUser))
	for _, user := range allUser {
		fmt.Printf("%s was last active at %q\n", user.Name, user.LastActive)
	}

}
