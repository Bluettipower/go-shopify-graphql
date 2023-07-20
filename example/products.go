package main

import (
	"context"
	"fmt"

	"github.com/bluettipower/go-shopify-graphql/v8"
)

func products(client *shopify.Client) {
	// Get products
	products, err := client.Product.List(context.Background(), "")
	if err != nil {
		panic(err)
	}

	// Print out the result
	for _, p := range products {
		fmt.Println(p.Title)
	}
}
