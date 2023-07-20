# go-shopify-graphql

A simple client using the Shopify GraphQL Admin API.

## Getting started

Hello World example

### 0. Setup

```bash
export STORE_API_KEY=<private_app_api_key>
export STORE_PASSWORD=<private_app_access_token>
export STORE_NAME=<store_name>
```

### 1. Program

```go
package main

import (
	"context"
	"fmt"

	shopify "github.com/bluettipower/go-shopify-graphql/v8"
)

func main() {
	// Create client
	client := shopify.NewDefaultClient()

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
```

### 3. Run

```bash
go run .
```
