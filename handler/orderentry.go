package handler

import (
	"context"
	"net"
	fix "sbe/sbe/iLinkBinary"
)

func New() OrderEntry {
	return &orderEntry{}
}

// Lotto ...
type OrderEntry interface {
	OnOrderNew(ctx context.Context, no *fix.NewOrderSingle514, con *net.Conn) error
	OnOrderCancel(ctx context.Context, co *fix.OrderCancelRequest516, con *net.Conn) error
	OnOrderModify(ctx context.Context, mo *fix.OrderCancelReplaceRequest515, con *net.Conn) error
}

type orderEntry struct {
}

func (o *orderEntry) OnOrderNew(ctx context.Context, no *fix.NewOrderSingle514, con *net.Conn) error{
	return nil
}

func (o *orderEntry) OnOrderCancel(ctx context.Context, co *fix.OrderCancelRequest516, con *net.Conn) error {
	return nil
}

func (o *orderEntry) OnOrderModify(ctx context.Context, mo *fix.OrderCancelReplaceRequest515, con *net.Conn) error{
	return nil
}