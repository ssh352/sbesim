package handler

import (
	"context"
	"sbe/sbe/fix"
)

// Lotto ...
type OrderEntry interface {
	OnOrderNew(ctx context.Context, no *fix.NewOrder) error
	OnOrderCancel(ctx context.Context, co *fix.OrderCancelRequest) error
	OnOrderModify(ctx context.Context, mo *fix.OrderCancelReplaceRequest) error
}