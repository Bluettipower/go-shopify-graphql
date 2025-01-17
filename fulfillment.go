package shopify

import (
	"context"
	"fmt"

	"github.com/r0busta/go-shopify-graphql-model/v3/graph/model"
)

//go:generate mockgen -destination=./mock/fulfillment_service.go -package=mock . FulfillmentService
type FulfillmentService interface {
	// Create(ctx context.Context, input model.FulfillmentV2Input) (*model.Fulfillment, error)
	Create(ctx context.Context, query string, input model.FulfillmentV2Input) (*model.Fulfillment, error)
	Update(ctx context.Context, query string, input FulfillmentTrackingInfoUpdateV2Input) error
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

type mutationFulfillmentTrackingInfoUpdateV2 struct {
	FulfillmentTrackingInfoUpdateV2Result struct {
		UserErrors []model.UserError `json:"userErrors,omitempty"`
	} `graphql:"fulfillmentTrackingInfoUpdateV2($fulfillmentId: fulfillmentId, $trackingInfoInput: trackingInfoInput, $notifyCustomer: notifyCustomer)" json:"fulfillmentTrackingInfoUpdateV2"`
}

type FulfillmentTrackingInfoUpdateV2Input struct {
	FulfillmentID     string             `json:"fulfillmentId,omitempty"`
	NotifyCustomer    *bool              `json:"notifyCustomer,omitempty,omitempty"`
	TrackingInfoInput *TrackingInfoInput `json:"trackingInfoInput,omitempty"`
}

type TrackingInfoInput struct {
	Company *string  `json:"company,omitempty"`
	Number  *string  `json:"number,omitempty"`
	Numbers []string `json:"numbers,omitempty"`
	Url     *string  `json:"url,omitempty"`
	Urls    []string `json:"urls,omitempty"`
}

func (s *FulfillmentServiceOp) Create(ctx context.Context, query string, fulfillment model.FulfillmentV2Input) (*model.Fulfillment, error) {
	m := mutationFulfillmentCreateV2{}
	vars := map[string]interface{}{
		"fulfillment": fulfillment,
	}

	err := s.client.gql.MutateString(ctx, query, vars, &m)
	if err != nil {
		return nil, fmt.Errorf("mutation: %w", err)
	}
	if len(m.FulfillmentCreateV2Result.UserErrors) > 0 {
		return nil, fmt.Errorf("UserErrors: %+v", m.FulfillmentCreateV2Result.UserErrors)
	}
	return &m.FulfillmentCreateV2Result.Fulfillment, nil
}

func (s *FulfillmentServiceOp) Update(ctx context.Context, query string, fulfillment FulfillmentTrackingInfoUpdateV2Input) error {
	m := mutationFulfillmentTrackingInfoUpdateV2{}

	vars := map[string]interface{}{
		"fulfillmentId":     fulfillment.FulfillmentID,
		"notifyCustomer":    fulfillment.NotifyCustomer,
		"trackingInfoInput": fulfillment.TrackingInfoInput,
	}
	err := s.client.gql.MutateString(ctx, query, vars, &m)
	if err != nil {
		return fmt.Errorf("mutation: %w", err)
	}
	if len(m.FulfillmentTrackingInfoUpdateV2Result.UserErrors) > 0 {
		return fmt.Errorf("UserErrors: %+v", m.FulfillmentTrackingInfoUpdateV2Result.UserErrors)
	}
	return nil
}
