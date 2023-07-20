package main

import (
	"context"
	"fmt"

	"github.com/bluettipower/go-shopify-graphql/v8"
	"github.com/r0busta/go-shopify-graphql-model/v3/graph/model"
)

func bulk(client *shopify.Client) {
	q := `
	{
		products{
			edges {
				node {
					id
					variants {
						edges {
							node {
								id
								media{
									edges {
										node {
											... on MediaImage {
												id
												image {
													url
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}`

	products := []*model.Product{}
	err := client.BulkOperation.BulkQuery(context.Background(), q, &products)
	if err != nil {
		panic(err)
	}

	// Print out the result
	for _, p := range products {
		for _, v := range p.Variants.Edges {
			for _, m := range v.Node.Media.Edges {
				fmt.Println(m.Node.(*model.MediaImage).Image.URL)
			}
		}
	}
}
