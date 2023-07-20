package main

import (
	shopify "github.com/bluettipower/go-shopify-graphql/v8"
)

func defaultClient() *shopify.Client {
	return shopify.NewDefaultClient()
}
