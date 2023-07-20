package shopify

import (
	"context"
	"fmt"

	"github.com/r0busta/go-shopify-graphql-model/v3/graph/model"
)

//go:generate mockgen -destination=./mock/fulfillment_service.go -package=mock . FulfillmentService
type FulfillmentService interface {
	Create(ctx context.Context, input model.FulfillmentV2Input) (*model.Fulfillment, error)
}

type FulfillmentServiceOp struct {
	client *Client
}

var _ FulfillmentService = &FulfillmentServiceOp{}

type mutationFulfillmentCreateV2 struct {
	FulfillmentCreateV2Result struct {
		Fulfillment model.Fulfillment `json:"fulfillment,omitempty"`
		UserErrors  []model.UserError `json:"userErrors,omitempty"`
	} `graphql:"fulfillmentCreateV2(fulfillment: $fulfillment)" json:"fulfillmentCreateV2"`
}

func (s *FulfillmentServiceOp) Create(ctx context.Context, fulfillment model.FulfillmentV2Input) (*model.Fulfillment, error) {
	m := mutationFulfillmentCreateV2{}

	vars := map[string]interface{}{
		"fulfillment": fulfillment,
	}
	err := s.client.gql.Mutate(ctx, &m, vars)
	if err != nil {
		return nil, fmt.Errorf("mutation: %w", err)
	}

	if len(m.FulfillmentCreateV2Result.UserErrors) > 0 {
		return nil, fmt.Errorf("UserErrors: %+v", m.FulfillmentCreateV2Result.UserErrors)
	}

	return &m.FulfillmentCreateV2Result.Fulfillment, nil
}
