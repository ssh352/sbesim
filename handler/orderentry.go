package handler

import (
	"context"
	"sbe/entity"
	fix "sbe/sbe/iLinkBinary"
)

func New() OrderEntry {
	return &orderEntry{}
}

// Lotto ...
type OrderEntry interface {
	OnOrderNew(ctx context.Context, no *fix.NewOrderSingle514, s entity.SBESession) error
	OnOrderCancel(ctx context.Context, co *fix.OrderCancelRequest516, s entity.SBESession) error
	OnOrderModify(ctx context.Context, mo *fix.OrderCancelReplaceRequest515, s entity.SBESession) error
}

type orderEntry struct {
}

func (o *orderEntry) OnOrderNew(ctx context.Context, no *fix.NewOrderSingle514, s entity.SBESession) error{
	return nil
}

func (o *orderEntry) OnOrderCancel(ctx context.Context, co *fix.OrderCancelRequest516, s entity.SBESession) error {
	return nil
}

func (o *orderEntry) OnOrderModify(ctx context.Context, mo *fix.OrderCancelReplaceRequest515, s entity.SBESession) error{
	return nil
}