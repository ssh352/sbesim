package handler

import (
	"context"
	"github.com/google/uuid"
	"sbe/entity"
	fix "sbe/sbe/iLinkBinary"
	"time"
)

func New() OrderEntry {
	return &orderEntry{
		uniqID:  0,
	}
}

//go:generate mockgen -destination=../mock/handler/mock_orderentry.go -package=mockhandler  sbe/handler OrderEntry
type OrderEntry interface {
	OnOrderNew(ctx context.Context, no *fix.NewOrderSingle514, s entity.SBESession) error
	OnOrderCancel(ctx context.Context, co *fix.OrderCancelRequest516, s entity.SBESession) error
	OnOrderModify(ctx context.Context, mo *fix.OrderCancelReplaceRequest515, s entity.SBESession) error
}

type orderEntry struct {
	uniqID uint64
}

func (o *orderEntry) getUniqID() uint64 {
	ret := o.uniqID
	o.uniqID = o.uniqID + 1
	return ret
}

func (o *orderEntry) OnOrderNew(ctx context.Context, no *fix.NewOrderSingle514, s entity.SBESession) error{
	ack := fix.ExecutionReportNew522{}
	ack.UUID = s.GetUUID()
	copy(ack.ExecID[:], uuid.New().String())
	ack.ClOrdID = no.ClOrdID
	ack.PartyDetailsListReqID = 0
	ack.OrderID = o.getUniqID()
	ack.Price.Mantissa = 100
	ack.TransactTime = uint64(time.Now().UTC().UnixNano())
	ack.OrderRequestID = o.getUniqID()
	copy(ack.Location[:], "US/IL")
	ack.SecurityID = no.SecurityID
	ack.OrderQty = no.OrderQty
	ack.ExpireDate = 59999
	// status, execType are constant
	ack.OrdType = fix.OrderTypeEnum(no.OrdType)
	ack.Side = no.Side
	ack.TimeInForce = no.TimeInForce
	ack.ManualOrderIndicator = no.ManualOrderIndicator
	ack.PossRetransFlag = fix.BooleanFlag.False
	ack.SeqNum = s.GetSeqNo()
	seq := ack.SeqNum
	if err := s.Send(&ack); err != nil {
		return err
	}
	s.SetSeqNo(seq + 1)
	fill := fix.ExecutionReportTradeOutright525{}
	fill.SeqNum = s.GetSeqNo()
	fill.UUID = s.GetUUID()
	copy(fill.ExecID[:], uuid.New().String())
	copy(fill.SenderID[:], "CME")
	fill.ClOrdID = no.ClOrdID
	fill.PartyDetailsListReqID = 0
	fill.OrderID = ack.OrderID
	fill.LastPx = ack.Price
	fill.TransactTime = uint64(time.Now().UTC().Nanosecond())
	fill.SendingTimeEpoch = fill.TransactTime
	fill.SecExecID = 1
	copy(fill.Location[:], "US/IL")
	fill.SecurityID = no.SecurityID
	fill.LastQty = no.OrderQty
	fill.OrderQty = no.OrderQty
	fill.CumQty = no.OrderQty
	fill.MDTradeEntryID = uint32(o.getUniqID())
	fill.SideTradeID = uint32(o.getUniqID())
	fill.LeavesQty = 0
	today := time.Now().UTC().Unix() / (3600 * 24 )
	fill.TradeDate = uint16(today)
	fill.ExpireDate = fill.TradeDate + 10
	fill.OrdStatus = fix.OrdStatusTrd.Filled
	fill.OrdType = ack.OrdType
	fill.Side = ack.Side
	fill.TimeInForce = no.TimeInForce
	fill.ManualOrderIndicator = ack.ManualOrderIndicator
	fill.PossRetransFlag = ack.PossRetransFlag
	fill.AggressorIndicator = fix.BooleanFlag.False
	var reason [2]byte
	copy(reason[:], "na")
	fill.NoFills = append(fill.NoFills, fix.ExecutionReportTradeOutright525NoFills{
		FillPx: fix.PRICE9{
			Mantissa:  100,
			Exponent:  1,
		},
		FillQty: no.OrderQty,
		FillYieldType:  4,
		FillExecID:  reason,
	})
	if err := s.Send(&fill); err != nil {
		return err
	}
	s.SetSeqNo(fill.SeqNum + 1)
	return nil
}

func (o *orderEntry) OnOrderCancel(ctx context.Context, co *fix.OrderCancelRequest516, s entity.SBESession) error {
	return nil
}

func (o *orderEntry) OnOrderModify(ctx context.Context, mo *fix.OrderCancelReplaceRequest515, s entity.SBESession) error{
	return nil
}