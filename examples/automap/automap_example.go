package main

import (
	"fmt"
	"github.com/zegl/goriak"
)

type User struct {
	ID   int
	Name string
}

func main() {
	con, err := goriak.Connect(goriak.ConnectOpts{Address: "127.0.0.1"})

	if err != nil {
		panic(err)
	}

	user := User{
		ID:   400,
		Name: "FooBar",
	}

	// Save our User object to Riak
	_, err = goriak.Bucket("bucket", "bucketType").
		Set(user).
		Key("user-400").
		Run(con)

	if err != nil {
		panic(err)
	}

	// Retrieve the same object from Riak
	var getUser User

	resp, err := goriak.Bucket("bucket", "bucketType").
		Get("user-400", &getUser).
		Run(con)

	if err != nil {
		panic(err)
	}

	if resp.NotFound {
		panic("Not found")
	}

	fmt.Printf("%+v", getUser)
}
