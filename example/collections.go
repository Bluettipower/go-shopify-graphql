package main

import (
	"context"
	"fmt"

	"github.com/bluettipower/go-shopify-graphql/v8"
)

func collections(client *shopify.Client) {
	// Get all collections
	collections, err := client.Collection.ListAll(context.Background())
	if err != nil {
		panic(err)
	}

	// Print out the result
	for _, c := range collections {
		fmt.Println(c.Handle)
	}
}
