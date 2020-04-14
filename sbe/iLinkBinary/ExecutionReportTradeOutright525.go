// Generated SBE (Simple Binary Encoding) message codec

package ilinkbinary

import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
)

type ExecutionReportTradeOutright525 struct {
	SeqNum                uint32
	UUID                  uint64
	ExecID                [40]byte
	SenderID              [20]byte
	ClOrdID               [20]byte
	PartyDetailsListReqID uint64
	LastPx                PRICE9
	OrderID               uint64
	Price                 PRICE9
	StopPx                PRICENULL9
	TransactTime          uint64
	SendingTimeEpoch      uint64
	OrderRequestID        uint64
	SecExecID             uint64
	CrossID               uint64
	HostCrossID           uint64
	Location              [5]byte
	SecurityID            int32
	OrderQty              uint32
	LastQty               uint32
	CumQty                uint32
	MDTradeEntryID        uint32
	SideTradeID           uint32
	TradeLinkID           uint32
	LeavesQty             uint32
	TradeDate             uint16
	ExpireDate            uint16
	OrdStatus             OrdStatusTrdEnum
	ExecType              [1]byte
	OrdType               OrderTypeEnum
	Side                  SideReqEnum
	TimeInForce           TimeInForceEnum
	ManualOrderIndicator  ManualOrdIndReqEnum
	PossRetransFlag       BooleanFlagEnum
	AggressorIndicator    BooleanFlagEnum
	CrossType             uint8
	ExecInst              ExecInst
	ExecutionMode         ExecModeEnum
	LiquidityFlag         BooleanNULLEnum
	ManagedOrder          BooleanNULLEnum
	ShortSaleType         ShortSaleTypeEnum
	Ownership             uint8
	NoFills               []ExecutionReportTradeOutright525NoFills
	NoOrderEvents         []ExecutionReportTradeOutright525NoOrderEvents
}
type ExecutionReportTradeOutright525NoFills struct {
	FillPx        PRICE9
	FillQty       uint32
	FillExecID    [2]byte
	FillYieldType uint8
}
type ExecutionReportTradeOutright525NoOrderEvents struct {
	OrderEventPx     PRICE9
	OrderEventText   [5]byte
	OrderEventExecID uint32
	OrderEventQty    uint32
	OrderEventType   OrderEventTypeEnum
	OrderEventReason uint8
}

func (e *ExecutionReportTradeOutright525) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := e.RangeCheck(e.SbeSchemaVersion(), e.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	if err := _m.WriteUint32(_w, e.SeqNum); err != nil {
		return err
	}
	if err := _m.WriteUint64(_w, e.UUID); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, e.ExecID[:]); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, e.SenderID[:]); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, e.ClOrdID[:]); err != nil {
		return err
	}
	if err := _m.WriteUint64(_w, e.PartyDetailsListReqID); err != nil {
		return err
	}
	if err := e.LastPx.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteUint64(_w, e.OrderID); err != nil {
		return err
	}
	if err := e.Price.Encode(_m, _w); err != nil {
		return err
	}
	if err := e.StopPx.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteUint64(_w, e.TransactTime); err != nil {
		return err
	}
	if err := _m.WriteUint64(_w, e.SendingTimeEpoch); err != nil {
		return err
	}
	if err := _m.WriteUint64(_w, e.OrderRequestID); err != nil {
		return err
	}
	if err := _m.WriteUint64(_w, e.SecExecID); err != nil {
		return err
	}
	if err := _m.WriteUint64(_w, e.CrossID); err != nil {
		return err
	}
	if err := _m.WriteUint64(_w, e.HostCrossID); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, e.Location[:]); err != nil {
		return err
	}
	if err := _m.WriteInt32(_w, e.SecurityID); err != nil {
		return err
	}
	if err := _m.WriteUint32(_w, e.OrderQty); err != nil {
		return err
	}
	if err := _m.WriteUint32(_w, e.LastQty); err != nil {
		return err
	}
	if err := _m.WriteUint32(_w, e.CumQty); err != nil {
		return err
	}
	if err := _m.WriteUint32(_w, e.MDTradeEntryID); err != nil {
		return err
	}
	if err := _m.WriteUint32(_w, e.SideTradeID); err != nil {
		return err
	}
	if err := _m.WriteUint32(_w, e.TradeLinkID); err != nil {
		return err
	}
	if err := _m.WriteUint32(_w, e.LeavesQty); err != nil {
		return err
	}
	if err := _m.WriteUint16(_w, e.TradeDate); err != nil {
		return err
	}
	if err := _m.WriteUint16(_w, e.ExpireDate); err != nil {
		return err
	}
	if err := e.OrdStatus.Encode(_m, _w); err != nil {
		return err
	}
	if err := e.OrdType.Encode(_m, _w); err != nil {
		return err
	}
	if err := e.Side.Encode(_m, _w); err != nil {
		return err
	}
	if err := e.TimeInForce.Encode(_m, _w); err != nil {
		return err
	}
	if err := e.ManualOrderIndicator.Encode(_m, _w); err != nil {
		return err
	}
	if err := e.PossRetransFlag.Encode(_m, _w); err != nil {
		return err
	}
	if err := e.AggressorIndicator.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteUint8(_w, e.CrossType); err != nil {
		return err
	}
	if err := e.ExecInst.Encode(_m, _w); err != nil {
		return err
	}
	if err := e.ExecutionMode.Encode(_m, _w); err != nil {
		return err
	}
	if err := e.LiquidityFlag.Encode(_m, _w); err != nil {
		return err
	}
	if err := e.ManagedOrder.Encode(_m, _w); err != nil {
		return err
	}
	if err := e.ShortSaleType.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteUint8(_w, e.Ownership); err != nil {
		return err
	}
	var NoFillsBlockLength uint16 = 15
	if err := _m.WriteUint16(_w, NoFillsBlockLength); err != nil {
		return err
	}
	var NoFillsNumInGroup uint8 = uint8(len(e.NoFills))
	if err := _m.WriteUint8(_w, NoFillsNumInGroup); err != nil {
		return err
	}
	for _, prop := range e.NoFills {
		if err := prop.Encode(_m, _w); err != nil {
			return err
		}
	}
	var NoOrderEventsBlockLength uint16 = 23
	if err := _m.WriteUint16(_w, NoOrderEventsBlockLength); err != nil {
		return err
	}
	var NoOrderEventsNumInGroup uint8 = uint8(len(e.NoOrderEvents))
	if err := _m.WriteUint8(_w, NoOrderEventsNumInGroup); err != nil {
		return err
	}
	for _, prop := range e.NoOrderEvents {
		if err := prop.Encode(_m, _w); err != nil {
			return err
		}
	}
	return nil
}

func (e *ExecutionReportTradeOutright525) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	if !e.SeqNumInActingVersion(actingVersion) {
		e.SeqNum = e.SeqNumNullValue()
	} else {
		if err := _m.ReadUint32(_r, &e.SeqNum); err != nil {
			return err
		}
	}
	if !e.UUIDInActingVersion(actingVersion) {
		e.UUID = e.UUIDNullValue()
	} else {
		if err := _m.ReadUint64(_r, &e.UUID); err != nil {
			return err
		}
	}
	if !e.ExecIDInActingVersion(actingVersion) {
		for idx := 0; idx < 40; idx++ {
			e.ExecID[idx] = e.ExecIDNullValue()
		}
	} else {
		if err := _m.ReadBytes(_r, e.ExecID[:]); err != nil {
			return err
		}
	}
	if !e.SenderIDInActingVersion(actingVersion) {
		for idx := 0; idx < 20; idx++ {
			e.SenderID[idx] = e.SenderIDNullValue()
		}
	} else {
		if err := _m.ReadBytes(_r, e.SenderID[:]); err != nil {
			return err
		}
	}
	if !e.ClOrdIDInActingVersion(actingVersion) {
		for idx := 0; idx < 20; idx++ {
			e.ClOrdID[idx] = e.ClOrdIDNullValue()
		}
	} else {
		if err := _m.ReadBytes(_r, e.ClOrdID[:]); err != nil {
			return err
		}
	}
	if !e.PartyDetailsListReqIDInActingVersion(actingVersion) {
		e.PartyDetailsListReqID = e.PartyDetailsListReqIDNullValue()
	} else {
		if err := _m.ReadUint64(_r, &e.PartyDetailsListReqID); err != nil {
			return err
		}
	}
	if e.LastPxInActingVersion(actingVersion) {
		if err := e.LastPx.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if !e.OrderIDInActingVersion(actingVersion) {
		e.OrderID = e.OrderIDNullValue()
	} else {
		if err := _m.ReadUint64(_r, &e.OrderID); err != nil {
			return err
		}
	}
	if e.PriceInActingVersion(actingVersion) {
		if err := e.Price.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if e.StopPxInActingVersion(actingVersion) {
		if err := e.StopPx.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if !e.TransactTimeInActingVersion(actingVersion) {
		e.TransactTime = e.TransactTimeNullValue()
	} else {
		if err := _m.ReadUint64(_r, &e.TransactTime); err != nil {
			return err
		}
	}
	if !e.SendingTimeEpochInActingVersion(actingVersion) {
		e.SendingTimeEpoch = e.SendingTimeEpochNullValue()
	} else {
		if err := _m.ReadUint64(_r, &e.SendingTimeEpoch); err != nil {
			return err
		}
	}
	if !e.OrderRequestIDInActingVersion(actingVersion) {
		e.OrderRequestID = e.OrderRequestIDNullValue()
	} else {
		if err := _m.ReadUint64(_r, &e.OrderRequestID); err != nil {
			return err
		}
	}
	if !e.SecExecIDInActingVersion(actingVersion) {
		e.SecExecID = e.SecExecIDNullValue()
	} else {
		if err := _m.ReadUint64(_r, &e.SecExecID); err != nil {
			return err
		}
	}
	if !e.CrossIDInActingVersion(actingVersion) {
		e.CrossID = e.CrossIDNullValue()
	} else {
		if err := _m.ReadUint64(_r, &e.CrossID); err != nil {
			return err
		}
	}
	if !e.HostCrossIDInActingVersion(actingVersion) {
		e.HostCrossID = e.HostCrossIDNullValue()
	} else {
		if err := _m.ReadUint64(_r, &e.HostCrossID); err != nil {
			return err
		}
	}
	if !e.LocationInActingVersion(actingVersion) {
		for idx := 0; idx < 5; idx++ {
			e.Location[idx] = e.LocationNullValue()
		}
	} else {
		if err := _m.ReadBytes(_r, e.Location[:]); err != nil {
			return err
		}
	}
	if !e.SecurityIDInActingVersion(actingVersion) {
		e.SecurityID = e.SecurityIDNullValue()
	} else {
		if err := _m.ReadInt32(_r, &e.SecurityID); err != nil {
			return err
		}
	}
	if !e.OrderQtyInActingVersion(actingVersion) {
		e.OrderQty = e.OrderQtyNullValue()
	} else {
		if err := _m.ReadUint32(_r, &e.OrderQty); err != nil {
			return err
		}
	}
	if !e.LastQtyInActingVersion(actingVersion) {
		e.LastQty = e.LastQtyNullValue()
	} else {
		if err := _m.ReadUint32(_r, &e.LastQty); err != nil {
			return err
		}
	}
	if !e.CumQtyInActingVersion(actingVersion) {
		e.CumQty = e.CumQtyNullValue()
	} else {
		if err := _m.ReadUint32(_r, &e.CumQty); err != nil {
			return err
		}
	}
	if !e.MDTradeEntryIDInActingVersion(actingVersion) {
		e.MDTradeEntryID = e.MDTradeEntryIDNullValue()
	} else {
		if err := _m.ReadUint32(_r, &e.MDTradeEntryID); err != nil {
			return err
		}
	}
	if !e.SideTradeIDInActingVersion(actingVersion) {
		e.SideTradeID = e.SideTradeIDNullValue()
	} else {
		if err := _m.ReadUint32(_r, &e.SideTradeID); err != nil {
			return err
		}
	}
	if !e.TradeLinkIDInActingVersion(actingVersion) {
		e.TradeLinkID = e.TradeLinkIDNullValue()
	} else {
		if err := _m.ReadUint32(_r, &e.TradeLinkID); err != nil {
			return err
		}
	}
	if !e.LeavesQtyInActingVersion(actingVersion) {
		e.LeavesQty = e.LeavesQtyNullValue()
	} else {
		if err := _m.ReadUint32(_r, &e.LeavesQty); err != nil {
			return err
		}
	}
	if !e.TradeDateInActingVersion(actingVersion) {
		e.TradeDate = e.TradeDateNullValue()
	} else {
		if err := _m.ReadUint16(_r, &e.TradeDate); err != nil {
			return err
		}
	}
	if !e.ExpireDateInActingVersion(actingVersion) {
		e.ExpireDate = e.ExpireDateNullValue()
	} else {
		if err := _m.ReadUint16(_r, &e.ExpireDate); err != nil {
			return err
		}
	}
	if e.OrdStatusInActingVersion(actingVersion) {
		if err := e.OrdStatus.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	e.ExecType[0] = 70
	if e.OrdTypeInActingVersion(actingVersion) {
		if err := e.OrdType.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if e.SideInActingVersion(actingVersion) {
		if err := e.Side.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if e.TimeInForceInActingVersion(actingVersion) {
		if err := e.TimeInForce.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if e.ManualOrderIndicatorInActingVersion(actingVersion) {
		if err := e.ManualOrderIndicator.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if e.PossRetransFlagInActingVersion(actingVersion) {
		if err := e.PossRetransFlag.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if e.AggressorIndicatorInActingVersion(actingVersion) {
		if err := e.AggressorIndicator.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if !e.CrossTypeInActingVersion(actingVersion) {
		e.CrossType = e.CrossTypeNullValue()
	} else {
		if err := _m.ReadUint8(_r, &e.CrossType); err != nil {
			return err
		}
	}
	if e.ExecInstInActingVersion(actingVersion) {
		if err := e.ExecInst.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if e.ExecutionModeInActingVersion(actingVersion) {
		if err := e.ExecutionMode.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if e.LiquidityFlagInActingVersion(actingVersion) {
		if err := e.LiquidityFlag.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if e.ManagedOrderInActingVersion(actingVersion) {
		if err := e.ManagedOrder.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if e.ShortSaleTypeInActingVersion(actingVersion) {
		if err := e.ShortSaleType.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if !e.OwnershipInActingVersion(actingVersion) {
		e.Ownership = e.OwnershipNullValue()
	} else {
		if err := _m.ReadUint8(_r, &e.Ownership); err != nil {
			return err
		}
	}
	if actingVersion > e.SbeSchemaVersion() && blockLength > e.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-e.SbeBlockLength()))
	}

	if e.NoFillsInActingVersion(actingVersion) {
		var NoFillsBlockLength uint16
		if err := _m.ReadUint16(_r, &NoFillsBlockLength); err != nil {
			return err
		}
		var NoFillsNumInGroup uint8
		if err := _m.ReadUint8(_r, &NoFillsNumInGroup); err != nil {
			return err
		}
		if cap(e.NoFills) < int(NoFillsNumInGroup) {
			e.NoFills = make([]ExecutionReportTradeOutright525NoFills, NoFillsNumInGroup)
		}
		for i, _ := range e.NoFills {
			if err := e.NoFills[i].Decode(_m, _r, actingVersion, uint(NoFillsBlockLength)); err != nil {
				return err
			}
		}
	}

	if e.NoOrderEventsInActingVersion(actingVersion) {
		var NoOrderEventsBlockLength uint16
		if err := _m.ReadUint16(_r, &NoOrderEventsBlockLength); err != nil {
			return err
		}
		var NoOrderEventsNumInGroup uint8
		if err := _m.ReadUint8(_r, &NoOrderEventsNumInGroup); err != nil {
			return err
		}
		if cap(e.NoOrderEvents) < int(NoOrderEventsNumInGroup) {
			e.NoOrderEvents = make([]ExecutionReportTradeOutright525NoOrderEvents, NoOrderEventsNumInGroup)
		}
		for i, _ := range e.NoOrderEvents {
			if err := e.NoOrderEvents[i].Decode(_m, _r, actingVersion, uint(NoOrderEventsBlockLength)); err != nil {
				return err
			}
		}
	}
	if doRangeCheck {
		if err := e.RangeCheck(actingVersion, e.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	return nil
}

func (e *ExecutionReportTradeOutright525) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if e.SeqNumInActingVersion(actingVersion) {
		if e.SeqNum < e.SeqNumMinValue() || e.SeqNum > e.SeqNumMaxValue() {
			return fmt.Errorf("Range check failed on e.SeqNum (%v < %v > %v)", e.SeqNumMinValue(), e.SeqNum, e.SeqNumMaxValue())
		}
	}
	if e.UUIDInActingVersion(actingVersion) {
		if e.UUID < e.UUIDMinValue() || e.UUID > e.UUIDMaxValue() {
			return fmt.Errorf("Range check failed on e.UUID (%v < %v > %v)", e.UUIDMinValue(), e.UUID, e.UUIDMaxValue())
		}
	}
	if e.ExecIDInActingVersion(actingVersion) {
		for idx := 0; idx < 40; idx++ {
			if e.ExecID[idx] < e.ExecIDMinValue() || e.ExecID[idx] > e.ExecIDMaxValue() {
				return fmt.Errorf("Range check failed on e.ExecID[%d] (%v < %v > %v)", idx, e.ExecIDMinValue(), e.ExecID[idx], e.ExecIDMaxValue())
			}
		}
	}
	if e.SenderIDInActingVersion(actingVersion) {
		for idx := 0; idx < 20; idx++ {
			if e.SenderID[idx] < e.SenderIDMinValue() || e.SenderID[idx] > e.SenderIDMaxValue() {
				return fmt.Errorf("Range check failed on e.SenderID[%d] (%v < %v > %v)", idx, e.SenderIDMinValue(), e.SenderID[idx], e.SenderIDMaxValue())
			}
		}
	}
	if e.ClOrdIDInActingVersion(actingVersion) {
		for idx := 0; idx < 20; idx++ {
			if e.ClOrdID[idx] < e.ClOrdIDMinValue() || e.ClOrdID[idx] > e.ClOrdIDMaxValue() {
				return fmt.Errorf("Range check failed on e.ClOrdID[%d] (%v < %v > %v)", idx, e.ClOrdIDMinValue(), e.ClOrdID[idx], e.ClOrdIDMaxValue())
			}
		}
	}
	if e.PartyDetailsListReqIDInActingVersion(actingVersion) {
		if e.PartyDetailsListReqID < e.PartyDetailsListReqIDMinValue() || e.PartyDetailsListReqID > e.PartyDetailsListReqIDMaxValue() {
			return fmt.Errorf("Range check failed on e.PartyDetailsListReqID (%v < %v > %v)", e.PartyDetailsListReqIDMinValue(), e.PartyDetailsListReqID, e.PartyDetailsListReqIDMaxValue())
		}
	}
	if e.OrderIDInActingVersion(actingVersion) {
		if e.OrderID < e.OrderIDMinValue() || e.OrderID > e.OrderIDMaxValue() {
			return fmt.Errorf("Range check failed on e.OrderID (%v < %v > %v)", e.OrderIDMinValue(), e.OrderID, e.OrderIDMaxValue())
		}
	}
	if e.TransactTimeInActingVersion(actingVersion) {
		if e.TransactTime < e.TransactTimeMinValue() || e.TransactTime > e.TransactTimeMaxValue() {
			return fmt.Errorf("Range check failed on e.TransactTime (%v < %v > %v)", e.TransactTimeMinValue(), e.TransactTime, e.TransactTimeMaxValue())
		}
	}
	if e.SendingTimeEpochInActingVersion(actingVersion) {
		if e.SendingTimeEpoch < e.SendingTimeEpochMinValue() || e.SendingTimeEpoch > e.SendingTimeEpochMaxValue() {
			return fmt.Errorf("Range check failed on e.SendingTimeEpoch (%v < %v > %v)", e.SendingTimeEpochMinValue(), e.SendingTimeEpoch, e.SendingTimeEpochMaxValue())
		}
	}
	if e.OrderRequestIDInActingVersion(actingVersion) {
		if e.OrderRequestID < e.OrderRequestIDMinValue() || e.OrderRequestID > e.OrderRequestIDMaxValue() {
			return fmt.Errorf("Range check failed on e.OrderRequestID (%v < %v > %v)", e.OrderRequestIDMinValue(), e.OrderRequestID, e.OrderRequestIDMaxValue())
		}
	}
	if e.SecExecIDInActingVersion(actingVersion) {
		if e.SecExecID < e.SecExecIDMinValue() || e.SecExecID > e.SecExecIDMaxValue() {
			return fmt.Errorf("Range check failed on e.SecExecID (%v < %v > %v)", e.SecExecIDMinValue(), e.SecExecID, e.SecExecIDMaxValue())
		}
	}
	if e.CrossIDInActingVersion(actingVersion) {
		if e.CrossID != e.CrossIDNullValue() && (e.CrossID < e.CrossIDMinValue() || e.CrossID > e.CrossIDMaxValue()) {
			return fmt.Errorf("Range check failed on e.CrossID (%v < %v > %v)", e.CrossIDMinValue(), e.CrossID, e.CrossIDMaxValue())
		}
	}
	if e.HostCrossIDInActingVersion(actingVersion) {
		if e.HostCrossID != e.HostCrossIDNullValue() && (e.HostCrossID < e.HostCrossIDMinValue() || e.HostCrossID > e.HostCrossIDMaxValue()) {
			return fmt.Errorf("Range check failed on e.HostCrossID (%v < %v > %v)", e.HostCrossIDMinValue(), e.HostCrossID, e.HostCrossIDMaxValue())
		}
	}
	if e.LocationInActingVersion(actingVersion) {
		for idx := 0; idx < 5; idx++ {
			if e.Location[idx] < e.LocationMinValue() || e.Location[idx] > e.LocationMaxValue() {
				return fmt.Errorf("Range check failed on e.Location[%d] (%v < %v > %v)", idx, e.LocationMinValue(), e.Location[idx], e.LocationMaxValue())
			}
		}
	}
	if e.SecurityIDInActingVersion(actingVersion) {
		if e.SecurityID < e.SecurityIDMinValue() || e.SecurityID > e.SecurityIDMaxValue() {
			return fmt.Errorf("Range check failed on e.SecurityID (%v < %v > %v)", e.SecurityIDMinValue(), e.SecurityID, e.SecurityIDMaxValue())
		}
	}
	if e.OrderQtyInActingVersion(actingVersion) {
		if e.OrderQty < e.OrderQtyMinValue() || e.OrderQty > e.OrderQtyMaxValue() {
			return fmt.Errorf("Range check failed on e.OrderQty (%v < %v > %v)", e.OrderQtyMinValue(), e.OrderQty, e.OrderQtyMaxValue())
		}
	}
	if e.LastQtyInActingVersion(actingVersion) {
		if e.LastQty < e.LastQtyMinValue() || e.LastQty > e.LastQtyMaxValue() {
			return fmt.Errorf("Range check failed on e.LastQty (%v < %v > %v)", e.LastQtyMinValue(), e.LastQty, e.LastQtyMaxValue())
		}
	}
	if e.CumQtyInActingVersion(actingVersion) {
		if e.CumQty < e.CumQtyMinValue() || e.CumQty > e.CumQtyMaxValue() {
			return fmt.Errorf("Range check failed on e.CumQty (%v < %v > %v)", e.CumQtyMinValue(), e.CumQty, e.CumQtyMaxValue())
		}
	}
	if e.MDTradeEntryIDInActingVersion(actingVersion) {
		if e.MDTradeEntryID < e.MDTradeEntryIDMinValue() || e.MDTradeEntryID > e.MDTradeEntryIDMaxValue() {
			return fmt.Errorf("Range check failed on e.MDTradeEntryID (%v < %v > %v)", e.MDTradeEntryIDMinValue(), e.MDTradeEntryID, e.MDTradeEntryIDMaxValue())
		}
	}
	if e.SideTradeIDInActingVersion(actingVersion) {
		if e.SideTradeID < e.SideTradeIDMinValue() || e.SideTradeID > e.SideTradeIDMaxValue() {
			return fmt.Errorf("Range check failed on e.SideTradeID (%v < %v > %v)", e.SideTradeIDMinValue(), e.SideTradeID, e.SideTradeIDMaxValue())
		}
	}
	if e.TradeLinkIDInActingVersion(actingVersion) {
		if e.TradeLinkID != e.TradeLinkIDNullValue() && (e.TradeLinkID < e.TradeLinkIDMinValue() || e.TradeLinkID > e.TradeLinkIDMaxValue()) {
			return fmt.Errorf("Range check failed on e.TradeLinkID (%v < %v > %v)", e.TradeLinkIDMinValue(), e.TradeLinkID, e.TradeLinkIDMaxValue())
		}
	}
	if e.LeavesQtyInActingVersion(actingVersion) {
		if e.LeavesQty < e.LeavesQtyMinValue() || e.LeavesQty > e.LeavesQtyMaxValue() {
			return fmt.Errorf("Range check failed on e.LeavesQty (%v < %v > %v)", e.LeavesQtyMinValue(), e.LeavesQty, e.LeavesQtyMaxValue())
		}
	}
	if e.TradeDateInActingVersion(actingVersion) {
		if e.TradeDate != e.TradeDateNullValue() && (e.TradeDate < e.TradeDateMinValue() || e.TradeDate > e.TradeDateMaxValue()) {
			return fmt.Errorf("Range check failed on e.TradeDate (%v < %v > %v)", e.TradeDateMinValue(), e.TradeDate, e.TradeDateMaxValue())
		}
	}
	if e.ExpireDateInActingVersion(actingVersion) {
		if e.ExpireDate != e.ExpireDateNullValue() && (e.ExpireDate < e.ExpireDateMinValue() || e.ExpireDate > e.ExpireDateMaxValue()) {
			return fmt.Errorf("Range check failed on e.ExpireDate (%v < %v > %v)", e.ExpireDateMinValue(), e.ExpireDate, e.ExpireDateMaxValue())
		}
	}
	if err := e.OrdStatus.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if err := e.OrdType.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if err := e.Side.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if err := e.TimeInForce.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if err := e.ManualOrderIndicator.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if err := e.PossRetransFlag.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if err := e.AggressorIndicator.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if e.CrossTypeInActingVersion(actingVersion) {
		if e.CrossType != e.CrossTypeNullValue() && (e.CrossType < e.CrossTypeMinValue() || e.CrossType > e.CrossTypeMaxValue()) {
			return fmt.Errorf("Range check failed on e.CrossType (%v < %v > %v)", e.CrossTypeMinValue(), e.CrossType, e.CrossTypeMaxValue())
		}
	}
	if err := e.ExecutionMode.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if err := e.LiquidityFlag.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if err := e.ManagedOrder.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if err := e.ShortSaleType.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if e.OwnershipInActingVersion(actingVersion) {
		if e.Ownership < e.OwnershipMinValue() || e.Ownership > e.OwnershipMaxValue() {
			return fmt.Errorf("Range check failed on e.Ownership (%v < %v > %v)", e.OwnershipMinValue(), e.Ownership, e.OwnershipMaxValue())
		}
	}
	for _, prop := range e.NoFills {
		if err := prop.RangeCheck(actingVersion, schemaVersion); err != nil {
			return err
		}
	}
	for _, prop := range e.NoOrderEvents {
		if err := prop.RangeCheck(actingVersion, schemaVersion); err != nil {
			return err
		}
	}
	return nil
}

func ExecutionReportTradeOutright525Init(e *ExecutionReportTradeOutright525) {
	e.CrossID = 18446744073709551615
	e.HostCrossID = 18446744073709551615
	e.TradeLinkID = 4294967295
	e.TradeDate = 65535
	e.ExpireDate = 65535
	e.ExecType[0] = 70
	e.CrossType = 255
	return
}

func (e *ExecutionReportTradeOutright525NoFills) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := e.FillPx.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteUint32(_w, e.FillQty); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, e.FillExecID[:]); err != nil {
		return err
	}
	if err := _m.WriteUint8(_w, e.FillYieldType); err != nil {
		return err
	}
	return nil
}

func (e *ExecutionReportTradeOutright525NoFills) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint) error {
	if e.FillPxInActingVersion(actingVersion) {
		if err := e.FillPx.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if !e.FillQtyInActingVersion(actingVersion) {
		e.FillQty = e.FillQtyNullValue()
	} else {
		if err := _m.ReadUint32(_r, &e.FillQty); err != nil {
			return err
		}
	}
	if !e.FillExecIDInActingVersion(actingVersion) {
		for idx := 0; idx < 2; idx++ {
			e.FillExecID[idx] = e.FillExecIDNullValue()
		}
	} else {
		if err := _m.ReadBytes(_r, e.FillExecID[:]); err != nil {
			return err
		}
	}
	if !e.FillYieldTypeInActingVersion(actingVersion) {
		e.FillYieldType = e.FillYieldTypeNullValue()
	} else {
		if err := _m.ReadUint8(_r, &e.FillYieldType); err != nil {
			return err
		}
	}
	if actingVersion > e.SbeSchemaVersion() && blockLength > e.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-e.SbeBlockLength()))
	}
	return nil
}

func (e *ExecutionReportTradeOutright525NoFills) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if e.FillQtyInActingVersion(actingVersion) {
		if e.FillQty < e.FillQtyMinValue() || e.FillQty > e.FillQtyMaxValue() {
			return fmt.Errorf("Range check failed on e.FillQty (%v < %v > %v)", e.FillQtyMinValue(), e.FillQty, e.FillQtyMaxValue())
		}
	}
	if e.FillExecIDInActingVersion(actingVersion) {
		for idx := 0; idx < 2; idx++ {
			if e.FillExecID[idx] < e.FillExecIDMinValue() || e.FillExecID[idx] > e.FillExecIDMaxValue() {
				return fmt.Errorf("Range check failed on e.FillExecID[%d] (%v < %v > %v)", idx, e.FillExecIDMinValue(), e.FillExecID[idx], e.FillExecIDMaxValue())
			}
		}
	}
	if e.FillYieldTypeInActingVersion(actingVersion) {
		if e.FillYieldType < e.FillYieldTypeMinValue() || e.FillYieldType > e.FillYieldTypeMaxValue() {
			return fmt.Errorf("Range check failed on e.FillYieldType (%v < %v > %v)", e.FillYieldTypeMinValue(), e.FillYieldType, e.FillYieldTypeMaxValue())
		}
	}
	return nil
}

func ExecutionReportTradeOutright525NoFillsInit(e *ExecutionReportTradeOutright525NoFills) {
	return
}

func (e *ExecutionReportTradeOutright525NoOrderEvents) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := e.OrderEventPx.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, e.OrderEventText[:]); err != nil {
		return err
	}
	if err := _m.WriteUint32(_w, e.OrderEventExecID); err != nil {
		return err
	}
	if err := _m.WriteUint32(_w, e.OrderEventQty); err != nil {
		return err
	}
	if err := e.OrderEventType.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteUint8(_w, e.OrderEventReason); err != nil {
		return err
	}
	return nil
}

func (e *ExecutionReportTradeOutright525NoOrderEvents) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint) error {
	if e.OrderEventPxInActingVersion(actingVersion) {
		if err := e.OrderEventPx.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if !e.OrderEventTextInActingVersion(actingVersion) {
		for idx := 0; idx < 5; idx++ {
			e.OrderEventText[idx] = e.OrderEventTextNullValue()
		}
	} else {
		if err := _m.ReadBytes(_r, e.OrderEventText[:]); err != nil {
			return err
		}
	}
	if !e.OrderEventExecIDInActingVersion(actingVersion) {
		e.OrderEventExecID = e.OrderEventExecIDNullValue()
	} else {
		if err := _m.ReadUint32(_r, &e.OrderEventExecID); err != nil {
			return err
		}
	}
	if !e.OrderEventQtyInActingVersion(actingVersion) {
		e.OrderEventQty = e.OrderEventQtyNullValue()
	} else {
		if err := _m.ReadUint32(_r, &e.OrderEventQty); err != nil {
			return err
		}
	}
	if e.OrderEventTypeInActingVersion(actingVersion) {
		if err := e.OrderEventType.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if !e.OrderEventReasonInActingVersion(actingVersion) {
		e.OrderEventReason = e.OrderEventReasonNullValue()
	} else {
		if err := _m.ReadUint8(_r, &e.OrderEventReason); err != nil {
			return err
		}
	}
	if actingVersion > e.SbeSchemaVersion() && blockLength > e.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-e.SbeBlockLength()))
	}
	return nil
}

func (e *ExecutionReportTradeOutright525NoOrderEvents) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if e.OrderEventTextInActingVersion(actingVersion) {
		for idx := 0; idx < 5; idx++ {
			if e.OrderEventText[idx] < e.OrderEventTextMinValue() || e.OrderEventText[idx] > e.OrderEventTextMaxValue() {
				return fmt.Errorf("Range check failed on e.OrderEventText[%d] (%v < %v > %v)", idx, e.OrderEventTextMinValue(), e.OrderEventText[idx], e.OrderEventTextMaxValue())
			}
		}
	}
	if e.OrderEventExecIDInActingVersion(actingVersion) {
		if e.OrderEventExecID < e.OrderEventExecIDMinValue() || e.OrderEventExecID > e.OrderEventExecIDMaxValue() {
			return fmt.Errorf("Range check failed on e.OrderEventExecID (%v < %v > %v)", e.OrderEventExecIDMinValue(), e.OrderEventExecID, e.OrderEventExecIDMaxValue())
		}
	}
	if e.OrderEventQtyInActingVersion(actingVersion) {
		if e.OrderEventQty < e.OrderEventQtyMinValue() || e.OrderEventQty > e.OrderEventQtyMaxValue() {
			return fmt.Errorf("Range check failed on e.OrderEventQty (%v < %v > %v)", e.OrderEventQtyMinValue(), e.OrderEventQty, e.OrderEventQtyMaxValue())
		}
	}
	if err := e.OrderEventType.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if e.OrderEventReasonInActingVersion(actingVersion) {
		if e.OrderEventReason < e.OrderEventReasonMinValue() || e.OrderEventReason > e.OrderEventReasonMaxValue() {
			return fmt.Errorf("Range check failed on e.OrderEventReason (%v < %v > %v)", e.OrderEventReasonMinValue(), e.OrderEventReason, e.OrderEventReasonMaxValue())
		}
	}
	return nil
}

func ExecutionReportTradeOutright525NoOrderEventsInit(e *ExecutionReportTradeOutright525NoOrderEvents) {
	for idx := 0; idx < 5; idx++ {
		e.OrderEventText[idx] = 0
	}
	return
}

func (*ExecutionReportTradeOutright525) SbeBlockLength() (blockLength uint16) {
	return 235
}

func (*ExecutionReportTradeOutright525) SbeTemplateId() (templateId uint16) {
	return 525
}

func (*ExecutionReportTradeOutright525) SbeSchemaId() (schemaId uint16) {
	return 8
}

func (*ExecutionReportTradeOutright525) SbeSchemaVersion() (schemaVersion uint16) {
	return 5
}

func (*ExecutionReportTradeOutright525) SbeSemanticType() (semanticType []byte) {
	return []byte("8")
}

func (*ExecutionReportTradeOutright525) SeqNumId() uint16 {
	return 9726
}

func (*ExecutionReportTradeOutright525) SeqNumSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeOutright525) SeqNumInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.SeqNumSinceVersion()
}

func (*ExecutionReportTradeOutright525) SeqNumDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeOutright525) SeqNumMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return "int"
	case 4:
		return "required"
	}
	return ""
}

func (*ExecutionReportTradeOutright525) SeqNumMinValue() uint32 {
	return 0
}

func (*ExecutionReportTradeOutright525) SeqNumMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*ExecutionReportTradeOutright525) SeqNumNullValue() uint32 {
	return math.MaxUint32
}

func (*ExecutionReportTradeOutright525) UUIDId() uint16 {
	return 39001
}

func (*ExecutionReportTradeOutright525) UUIDSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeOutright525) UUIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.UUIDSinceVersion()
}

func (*ExecutionReportTradeOutright525) UUIDDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeOutright525) UUIDMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return "int"
	case 4:
		return "required"
	}
	return ""
}

func (*ExecutionReportTradeOutright525) UUIDMinValue() uint64 {
	return 0
}

func (*ExecutionReportTradeOutright525) UUIDMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*ExecutionReportTradeOutright525) UUIDNullValue() uint64 {
	return math.MaxUint64
}

func (*ExecutionReportTradeOutright525) ExecIDId() uint16 {
	return 17
}

func (*ExecutionReportTradeOutright525) ExecIDSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeOutright525) ExecIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.ExecIDSinceVersion()
}

func (*ExecutionReportTradeOutright525) ExecIDDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeOutright525) ExecIDMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return "String"
	case 4:
		return "required"
	}
	return ""
}

func (*ExecutionReportTradeOutright525) ExecIDMinValue() byte {
	return byte(32)
}

func (*ExecutionReportTradeOutright525) ExecIDMaxValue() byte {
	return byte(126)
}

func (*ExecutionReportTradeOutright525) ExecIDNullValue() byte {
	return 0
}

func (e *ExecutionReportTradeOutright525) ExecIDCharacterEncoding() string {
	return "US-ASCII"
}

func (*ExecutionReportTradeOutright525) SenderIDId() uint16 {
	return 5392
}

func (*ExecutionReportTradeOutright525) SenderIDSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeOutright525) SenderIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.SenderIDSinceVersion()
}

func (*ExecutionReportTradeOutright525) SenderIDDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeOutright525) SenderIDMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return "String"
	case 4:
		return "required"
	}
	return ""
}

func (*ExecutionReportTradeOutright525) SenderIDMinValue() byte {
	return byte(32)
}

func (*ExecutionReportTradeOutright525) SenderIDMaxValue() byte {
	return byte(126)
}

func (*ExecutionReportTradeOutright525) SenderIDNullValue() byte {
	return 0
}

func (e *ExecutionReportTradeOutright525) SenderIDCharacterEncoding() string {
	return "US-ASCII"
}

func (*ExecutionReportTradeOutright525) ClOrdIDId() uint16 {
	return 11
}

func (*ExecutionReportTradeOutright525) ClOrdIDSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeOutright525) ClOrdIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.ClOrdIDSinceVersion()
}

func (*ExecutionReportTradeOutright525) ClOrdIDDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeOutright525) ClOrdIDMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return "String"
	case 4:
		return "required"
	}
	return ""
}

func (*ExecutionReportTradeOutright525) ClOrdIDMinValue() byte {
	return byte(32)
}

func (*ExecutionReportTradeOutright525) ClOrdIDMaxValue() byte {
	return byte(126)
}

func (*ExecutionReportTradeOutright525) ClOrdIDNullValue() byte {
	return 0
}

func (e *ExecutionReportTradeOutright525) ClOrdIDCharacterEncoding() string {
	return "US-ASCII"
}

func (*ExecutionReportTradeOutright525) PartyDetailsListReqIDId() uint16 {
	return 1505
}

func (*ExecutionReportTradeOutright525) PartyDetailsListReqIDSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeOutright525) PartyDetailsListReqIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.PartyDetailsListReqIDSinceVersion()
}

func (*ExecutionReportTradeOutright525) PartyDetailsListReqIDDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeOutright525) PartyDetailsListReqIDMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return "int"
	case 4:
		return "required"
	}
	return ""
}

func (*ExecutionReportTradeOutright525) PartyDetailsListReqIDMinValue() uint64 {
	return 0
}

func (*ExecutionReportTradeOutright525) PartyDetailsListReqIDMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*ExecutionReportTradeOutright525) PartyDetailsListReqIDNullValue() uint64 {
	return math.MaxUint64
}

func (*ExecutionReportTradeOutright525) LastPxId() uint16 {
	return 31
}

func (*ExecutionReportTradeOutright525) LastPxSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeOutright525) LastPxInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.LastPxSinceVersion()
}

func (*ExecutionReportTradeOutright525) LastPxDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeOutright525) LastPxMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return "Price"
	case 4:
		return "required"
	}
	return ""
}

func (*ExecutionReportTradeOutright525) OrderIDId() uint16 {
	return 37
}

func (*ExecutionReportTradeOutright525) OrderIDSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeOutright525) OrderIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.OrderIDSinceVersion()
}

func (*ExecutionReportTradeOutright525) OrderIDDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeOutright525) OrderIDMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return "int"
	case 4:
		return "required"
	}
	return ""
}

func (*ExecutionReportTradeOutright525) OrderIDMinValue() uint64 {
	return 0
}

func (*ExecutionReportTradeOutright525) OrderIDMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*ExecutionReportTradeOutright525) OrderIDNullValue() uint64 {
	return math.MaxUint64
}

func (*ExecutionReportTradeOutright525) PriceId() uint16 {
	return 44
}

func (*ExecutionReportTradeOutright525) PriceSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeOutright525) PriceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.PriceSinceVersion()
}

func (*ExecutionReportTradeOutright525) PriceDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeOutright525) PriceMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return "Price"
	case 4:
		return "required"
	}
	return ""
}

func (*ExecutionReportTradeOutright525) StopPxId() uint16 {
	return 99
}

func (*ExecutionReportTradeOutright525) StopPxSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeOutright525) StopPxInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.StopPxSinceVersion()
}

func (*ExecutionReportTradeOutright525) StopPxDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeOutright525) StopPxMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return "Price"
	case 4:
		return "required"
	}
	return ""
}

func (*ExecutionReportTradeOutright525) TransactTimeId() uint16 {
	return 60
}

func (*ExecutionReportTradeOutright525) TransactTimeSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeOutright525) TransactTimeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.TransactTimeSinceVersion()
}

func (*ExecutionReportTradeOutright525) TransactTimeDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeOutright525) TransactTimeMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return "int"
	case 4:
		return "required"
	}
	return ""
}

func (*ExecutionReportTradeOutright525) TransactTimeMinValue() uint64 {
	return 0
}

func (*ExecutionReportTradeOutright525) TransactTimeMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*ExecutionReportTradeOutright525) TransactTimeNullValue() uint64 {
	return math.MaxUint64
}

func (*ExecutionReportTradeOutright525) SendingTimeEpochId() uint16 {
	return 5297
}

func (*ExecutionReportTradeOutright525) SendingTimeEpochSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeOutright525) SendingTimeEpochInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.SendingTimeEpochSinceVersion()
}

func (*ExecutionReportTradeOutright525) SendingTimeEpochDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeOutright525) SendingTimeEpochMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return "int"
	case 4:
		return "required"
	}
	return ""
}

func (*ExecutionReportTradeOutright525) SendingTimeEpochMinValue() uint64 {
	return 0
}

func (*ExecutionReportTradeOutright525) SendingTimeEpochMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*ExecutionReportTradeOutright525) SendingTimeEpochNullValue() uint64 {
	return math.MaxUint64
}

func (*ExecutionReportTradeOutright525) OrderRequestIDId() uint16 {
	return 2422
}

func (*ExecutionReportTradeOutright525) OrderRequestIDSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeOutright525) OrderRequestIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.OrderRequestIDSinceVersion()
}

func (*ExecutionReportTradeOutright525) OrderRequestIDDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeOutright525) OrderRequestIDMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return "int"
	case 4:
		return "required"
	}
	return ""
}

func (*ExecutionReportTradeOutright525) OrderRequestIDMinValue() uint64 {
	return 0
}

func (*ExecutionReportTradeOutright525) OrderRequestIDMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*ExecutionReportTradeOutright525) OrderRequestIDNullValue() uint64 {
	return math.MaxUint64
}

func (*ExecutionReportTradeOutright525) SecExecIDId() uint16 {
	return 527
}

func (*ExecutionReportTradeOutright525) SecExecIDSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeOutright525) SecExecIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.SecExecIDSinceVersion()
}

func (*ExecutionReportTradeOutright525) SecExecIDDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeOutright525) SecExecIDMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return "int"
	case 4:
		return "required"
	}
	return ""
}

func (*ExecutionReportTradeOutright525) SecExecIDMinValue() uint64 {
	return 0
}

func (*ExecutionReportTradeOutright525) SecExecIDMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*ExecutionReportTradeOutright525) SecExecIDNullValue() uint64 {
	return math.MaxUint64
}

func (*ExecutionReportTradeOutright525) CrossIDId() uint16 {
	return 548
}

func (*ExecutionReportTradeOutright525) CrossIDSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeOutright525) CrossIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.CrossIDSinceVersion()
}

func (*ExecutionReportTradeOutright525) CrossIDDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeOutright525) CrossIDMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return "int"
	case 4:
		return "optional"
	}
	return ""
}

func (*ExecutionReportTradeOutright525) CrossIDMinValue() uint64 {
	return 0
}

func (*ExecutionReportTradeOutright525) CrossIDMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*ExecutionReportTradeOutright525) CrossIDNullValue() uint64 {
	return 18446744073709551615
}

func (*ExecutionReportTradeOutright525) HostCrossIDId() uint16 {
	return 961
}

func (*ExecutionReportTradeOutright525) HostCrossIDSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeOutright525) HostCrossIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.HostCrossIDSinceVersion()
}

func (*ExecutionReportTradeOutright525) HostCrossIDDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeOutright525) HostCrossIDMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return "int"
	case 4:
		return "optional"
	}
	return ""
}

func (*ExecutionReportTradeOutright525) HostCrossIDMinValue() uint64 {
	return 0
}

func (*ExecutionReportTradeOutright525) HostCrossIDMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*ExecutionReportTradeOutright525) HostCrossIDNullValue() uint64 {
	return 18446744073709551615
}

func (*ExecutionReportTradeOutright525) LocationId() uint16 {
	return 9537
}

func (*ExecutionReportTradeOutright525) LocationSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeOutright525) LocationInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.LocationSinceVersion()
}

func (*ExecutionReportTradeOutright525) LocationDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeOutright525) LocationMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return "String"
	case 4:
		return "required"
	}
	return ""
}

func (*ExecutionReportTradeOutright525) LocationMinValue() byte {
	return byte(32)
}

func (*ExecutionReportTradeOutright525) LocationMaxValue() byte {
	return byte(126)
}

func (*ExecutionReportTradeOutright525) LocationNullValue() byte {
	return 0
}

func (e *ExecutionReportTradeOutright525) LocationCharacterEncoding() string {
	return "US-ASCII"
}

func (*ExecutionReportTradeOutright525) SecurityIDId() uint16 {
	return 48
}

func (*ExecutionReportTradeOutright525) SecurityIDSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeOutright525) SecurityIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.SecurityIDSinceVersion()
}

func (*ExecutionReportTradeOutright525) SecurityIDDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeOutright525) SecurityIDMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return "int"
	case 4:
		return "required"
	}
	return ""
}

func (*ExecutionReportTradeOutright525) SecurityIDMinValue() int32 {
	return math.MinInt32 + 1
}

func (*ExecutionReportTradeOutright525) SecurityIDMaxValue() int32 {
	return math.MaxInt32
}

func (*ExecutionReportTradeOutright525) SecurityIDNullValue() int32 {
	return math.MinInt32
}

func (*ExecutionReportTradeOutright525) OrderQtyId() uint16 {
	return 38
}

func (*ExecutionReportTradeOutright525) OrderQtySinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeOutright525) OrderQtyInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.OrderQtySinceVersion()
}

func (*ExecutionReportTradeOutright525) OrderQtyDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeOutright525) OrderQtyMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return "int"
	case 4:
		return "required"
	}
	return ""
}

func (*ExecutionReportTradeOutright525) OrderQtyMinValue() uint32 {
	return 0
}

func (*ExecutionReportTradeOutright525) OrderQtyMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*ExecutionReportTradeOutright525) OrderQtyNullValue() uint32 {
	return math.MaxUint32
}

func (*ExecutionReportTradeOutright525) LastQtyId() uint16 {
	return 32
}

func (*ExecutionReportTradeOutright525) LastQtySinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeOutright525) LastQtyInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.LastQtySinceVersion()
}

func (*ExecutionReportTradeOutright525) LastQtyDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeOutright525) LastQtyMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return "int"
	case 4:
		return "required"
	}
	return ""
}

func (*ExecutionReportTradeOutright525) LastQtyMinValue() uint32 {
	return 0
}

func (*ExecutionReportTradeOutright525) LastQtyMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*ExecutionReportTradeOutright525) LastQtyNullValue() uint32 {
	return math.MaxUint32
}

func (*ExecutionReportTradeOutright525) CumQtyId() uint16 {
	return 14
}

func (*ExecutionReportTradeOutright525) CumQtySinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeOutright525) CumQtyInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.CumQtySinceVersion()
}

func (*ExecutionReportTradeOutright525) CumQtyDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeOutright525) CumQtyMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return "int"
	case 4:
		return "required"
	}
	return ""
}

func (*ExecutionReportTradeOutright525) CumQtyMinValue() uint32 {
	return 0
}

func (*ExecutionReportTradeOutright525) CumQtyMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*ExecutionReportTradeOutright525) CumQtyNullValue() uint32 {
	return math.MaxUint32
}

func (*ExecutionReportTradeOutright525) MDTradeEntryIDId() uint16 {
	return 37711
}

func (*ExecutionReportTradeOutright525) MDTradeEntryIDSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeOutright525) MDTradeEntryIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.MDTradeEntryIDSinceVersion()
}

func (*ExecutionReportTradeOutright525) MDTradeEntryIDDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeOutright525) MDTradeEntryIDMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return "int"
	case 4:
		return "required"
	}
	return ""
}

func (*ExecutionReportTradeOutright525) MDTradeEntryIDMinValue() uint32 {
	return 0
}

func (*ExecutionReportTradeOutright525) MDTradeEntryIDMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*ExecutionReportTradeOutright525) MDTradeEntryIDNullValue() uint32 {
	return math.MaxUint32
}

func (*ExecutionReportTradeOutright525) SideTradeIDId() uint16 {
	return 1506
}

func (*ExecutionReportTradeOutright525) SideTradeIDSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeOutright525) SideTradeIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.SideTradeIDSinceVersion()
}

func (*ExecutionReportTradeOutright525) SideTradeIDDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeOutright525) SideTradeIDMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return "int"
	case 4:
		return "required"
	}
	return ""
}

func (*ExecutionReportTradeOutright525) SideTradeIDMinValue() uint32 {
	return 0
}

func (*ExecutionReportTradeOutright525) SideTradeIDMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*ExecutionReportTradeOutright525) SideTradeIDNullValue() uint32 {
	return math.MaxUint32
}

func (*ExecutionReportTradeOutright525) TradeLinkIDId() uint16 {
	return 820
}

func (*ExecutionReportTradeOutright525) TradeLinkIDSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeOutright525) TradeLinkIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.TradeLinkIDSinceVersion()
}

func (*ExecutionReportTradeOutright525) TradeLinkIDDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeOutright525) TradeLinkIDMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return "int"
	case 4:
		return "optional"
	}
	return ""
}

func (*ExecutionReportTradeOutright525) TradeLinkIDMinValue() uint32 {
	return 0
}

func (*ExecutionReportTradeOutright525) TradeLinkIDMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*ExecutionReportTradeOutright525) TradeLinkIDNullValue() uint32 {
	return 4294967295
}

func (*ExecutionReportTradeOutright525) LeavesQtyId() uint16 {
	return 151
}

func (*ExecutionReportTradeOutright525) LeavesQtySinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeOutright525) LeavesQtyInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.LeavesQtySinceVersion()
}

func (*ExecutionReportTradeOutright525) LeavesQtyDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeOutright525) LeavesQtyMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return "int"
	case 4:
		return "required"
	}
	return ""
}

func (*ExecutionReportTradeOutright525) LeavesQtyMinValue() uint32 {
	return 0
}

func (*ExecutionReportTradeOutright525) LeavesQtyMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*ExecutionReportTradeOutright525) LeavesQtyNullValue() uint32 {
	return math.MaxUint32
}

func (*ExecutionReportTradeOutright525) TradeDateId() uint16 {
	return 75
}

func (*ExecutionReportTradeOutright525) TradeDateSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeOutright525) TradeDateInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.TradeDateSinceVersion()
}

func (*ExecutionReportTradeOutright525) TradeDateDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeOutright525) TradeDateMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return "LocalMktDate"
	case 4:
		return "optional"
	}
	return ""
}

func (*ExecutionReportTradeOutright525) TradeDateMinValue() uint16 {
	return 0
}

func (*ExecutionReportTradeOutright525) TradeDateMaxValue() uint16 {
	return math.MaxUint16 - 1
}

func (*ExecutionReportTradeOutright525) TradeDateNullValue() uint16 {
	return 65535
}

func (*ExecutionReportTradeOutright525) ExpireDateId() uint16 {
	return 432
}

func (*ExecutionReportTradeOutright525) ExpireDateSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeOutright525) ExpireDateInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.ExpireDateSinceVersion()
}

func (*ExecutionReportTradeOutright525) ExpireDateDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeOutright525) ExpireDateMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return "LocalMktDate"
	case 4:
		return "optional"
	}
	return ""
}

func (*ExecutionReportTradeOutright525) ExpireDateMinValue() uint16 {
	return 0
}

func (*ExecutionReportTradeOutright525) ExpireDateMaxValue() uint16 {
	return math.MaxUint16 - 1
}

func (*ExecutionReportTradeOutright525) ExpireDateNullValue() uint16 {
	return 65535
}

func (*ExecutionReportTradeOutright525) OrdStatusId() uint16 {
	return 39
}

func (*ExecutionReportTradeOutright525) OrdStatusSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeOutright525) OrdStatusInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.OrdStatusSinceVersion()
}

func (*ExecutionReportTradeOutright525) OrdStatusDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeOutright525) OrdStatusMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return "int"
	case 4:
		return "required"
	}
	return ""
}

func (*ExecutionReportTradeOutright525) ExecTypeId() uint16 {
	return 150
}

func (*ExecutionReportTradeOutright525) ExecTypeSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeOutright525) ExecTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.ExecTypeSinceVersion()
}

func (*ExecutionReportTradeOutright525) ExecTypeDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeOutright525) ExecTypeMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return "char"
	case 4:
		return "constant"
	}
	return ""
}

func (*ExecutionReportTradeOutright525) ExecTypeMinValue() byte {
	return byte(32)
}

func (*ExecutionReportTradeOutright525) ExecTypeMaxValue() byte {
	return byte(126)
}

func (*ExecutionReportTradeOutright525) ExecTypeNullValue() byte {
	return 0
}

func (*ExecutionReportTradeOutright525) OrdTypeId() uint16 {
	return 40
}

func (*ExecutionReportTradeOutright525) OrdTypeSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeOutright525) OrdTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.OrdTypeSinceVersion()
}

func (*ExecutionReportTradeOutright525) OrdTypeDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeOutright525) OrdTypeMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return "char"
	case 4:
		return "required"
	}
	return ""
}

func (*ExecutionReportTradeOutright525) SideId() uint16 {
	return 54
}

func (*ExecutionReportTradeOutright525) SideSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeOutright525) SideInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.SideSinceVersion()
}

func (*ExecutionReportTradeOutright525) SideDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeOutright525) SideMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return "int"
	case 4:
		return "required"
	}
	return ""
}

func (*ExecutionReportTradeOutright525) TimeInForceId() uint16 {
	return 59
}

func (*ExecutionReportTradeOutright525) TimeInForceSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeOutright525) TimeInForceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.TimeInForceSinceVersion()
}

func (*ExecutionReportTradeOutright525) TimeInForceDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeOutright525) TimeInForceMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return "int"
	case 4:
		return "required"
	}
	return ""
}

func (*ExecutionReportTradeOutright525) ManualOrderIndicatorId() uint16 {
	return 1028
}

func (*ExecutionReportTradeOutright525) ManualOrderIndicatorSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeOutright525) ManualOrderIndicatorInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.ManualOrderIndicatorSinceVersion()
}

func (*ExecutionReportTradeOutright525) ManualOrderIndicatorDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeOutright525) ManualOrderIndicatorMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return "int"
	case 4:
		return "required"
	}
	return ""
}

func (*ExecutionReportTradeOutright525) PossRetransFlagId() uint16 {
	return 9765
}

func (*ExecutionReportTradeOutright525) PossRetransFlagSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeOutright525) PossRetransFlagInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.PossRetransFlagSinceVersion()
}

func (*ExecutionReportTradeOutright525) PossRetransFlagDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeOutright525) PossRetransFlagMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return "int"
	case 4:
		return "required"
	}
	return ""
}

func (*ExecutionReportTradeOutright525) AggressorIndicatorId() uint16 {
	return 1057
}

func (*ExecutionReportTradeOutright525) AggressorIndicatorSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeOutright525) AggressorIndicatorInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.AggressorIndicatorSinceVersion()
}

func (*ExecutionReportTradeOutright525) AggressorIndicatorDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeOutright525) AggressorIndicatorMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return "int"
	case 4:
		return "required"
	}
	return ""
}

func (*ExecutionReportTradeOutright525) CrossTypeId() uint16 {
	return 549
}

func (*ExecutionReportTradeOutright525) CrossTypeSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeOutright525) CrossTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.CrossTypeSinceVersion()
}

func (*ExecutionReportTradeOutright525) CrossTypeDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeOutright525) CrossTypeMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return "int"
	case 4:
		return "optional"
	}
	return ""
}

func (*ExecutionReportTradeOutright525) CrossTypeMinValue() uint8 {
	return 0
}

func (*ExecutionReportTradeOutright525) CrossTypeMaxValue() uint8 {
	return math.MaxUint8 - 1
}

func (*ExecutionReportTradeOutright525) CrossTypeNullValue() uint8 {
	return 255
}

func (*ExecutionReportTradeOutright525) ExecInstId() uint16 {
	return 18
}

func (*ExecutionReportTradeOutright525) ExecInstSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeOutright525) ExecInstInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.ExecInstSinceVersion()
}

func (*ExecutionReportTradeOutright525) ExecInstDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeOutright525) ExecInstMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return "MultipleCharValue"
	case 4:
		return "required"
	}
	return ""
}

func (*ExecutionReportTradeOutright525) ExecutionModeId() uint16 {
	return 5906
}

func (*ExecutionReportTradeOutright525) ExecutionModeSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeOutright525) ExecutionModeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.ExecutionModeSinceVersion()
}

func (*ExecutionReportTradeOutright525) ExecutionModeDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeOutright525) ExecutionModeMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return "char"
	case 4:
		return "required"
	}
	return ""
}

func (*ExecutionReportTradeOutright525) LiquidityFlagId() uint16 {
	return 9373
}

func (*ExecutionReportTradeOutright525) LiquidityFlagSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeOutright525) LiquidityFlagInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.LiquidityFlagSinceVersion()
}

func (*ExecutionReportTradeOutright525) LiquidityFlagDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeOutright525) LiquidityFlagMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return "int"
	case 4:
		return "required"
	}
	return ""
}

func (*ExecutionReportTradeOutright525) ManagedOrderId() uint16 {
	return 6881
}

func (*ExecutionReportTradeOutright525) ManagedOrderSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeOutright525) ManagedOrderInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.ManagedOrderSinceVersion()
}

func (*ExecutionReportTradeOutright525) ManagedOrderDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeOutright525) ManagedOrderMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return "int"
	case 4:
		return "required"
	}
	return ""
}

func (*ExecutionReportTradeOutright525) ShortSaleTypeId() uint16 {
	return 5409
}

func (*ExecutionReportTradeOutright525) ShortSaleTypeSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeOutright525) ShortSaleTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.ShortSaleTypeSinceVersion()
}

func (*ExecutionReportTradeOutright525) ShortSaleTypeDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeOutright525) ShortSaleTypeMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return "int"
	case 4:
		return "required"
	}
	return ""
}

func (*ExecutionReportTradeOutright525) OwnershipId() uint16 {
	return 7191
}

func (*ExecutionReportTradeOutright525) OwnershipSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeOutright525) OwnershipInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.OwnershipSinceVersion()
}

func (*ExecutionReportTradeOutright525) OwnershipDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeOutright525) OwnershipMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return "int"
	case 4:
		return "required"
	}
	return ""
}

func (*ExecutionReportTradeOutright525) OwnershipMinValue() uint8 {
	return 0
}

func (*ExecutionReportTradeOutright525) OwnershipMaxValue() uint8 {
	return math.MaxUint8 - 1
}

func (*ExecutionReportTradeOutright525) OwnershipNullValue() uint8 {
	return math.MaxUint8
}

func (*ExecutionReportTradeOutright525NoFills) FillPxId() uint16 {
	return 1364
}

func (*ExecutionReportTradeOutright525NoFills) FillPxSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeOutright525NoFills) FillPxInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.FillPxSinceVersion()
}

func (*ExecutionReportTradeOutright525NoFills) FillPxDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeOutright525NoFills) FillPxMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return "Price"
	case 4:
		return "required"
	}
	return ""
}

func (*ExecutionReportTradeOutright525NoFills) FillQtyId() uint16 {
	return 1365
}

func (*ExecutionReportTradeOutright525NoFills) FillQtySinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeOutright525NoFills) FillQtyInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.FillQtySinceVersion()
}

func (*ExecutionReportTradeOutright525NoFills) FillQtyDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeOutright525NoFills) FillQtyMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return "int"
	case 4:
		return "required"
	}
	return ""
}

func (*ExecutionReportTradeOutright525NoFills) FillQtyMinValue() uint32 {
	return 0
}

func (*ExecutionReportTradeOutright525NoFills) FillQtyMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*ExecutionReportTradeOutright525NoFills) FillQtyNullValue() uint32 {
	return math.MaxUint32
}

func (*ExecutionReportTradeOutright525NoFills) FillExecIDId() uint16 {
	return 1363
}

func (*ExecutionReportTradeOutright525NoFills) FillExecIDSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeOutright525NoFills) FillExecIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.FillExecIDSinceVersion()
}

func (*ExecutionReportTradeOutright525NoFills) FillExecIDDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeOutright525NoFills) FillExecIDMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return "String"
	case 4:
		return "required"
	}
	return ""
}

func (*ExecutionReportTradeOutright525NoFills) FillExecIDMinValue() byte {
	return byte(32)
}

func (*ExecutionReportTradeOutright525NoFills) FillExecIDMaxValue() byte {
	return byte(126)
}

func (*ExecutionReportTradeOutright525NoFills) FillExecIDNullValue() byte {
	return 0
}

func (e *ExecutionReportTradeOutright525NoFills) FillExecIDCharacterEncoding() string {
	return "US-ASCII"
}

func (*ExecutionReportTradeOutright525NoFills) FillYieldTypeId() uint16 {
	return 1622
}

func (*ExecutionReportTradeOutright525NoFills) FillYieldTypeSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeOutright525NoFills) FillYieldTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.FillYieldTypeSinceVersion()
}

func (*ExecutionReportTradeOutright525NoFills) FillYieldTypeDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeOutright525NoFills) FillYieldTypeMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return "int"
	case 4:
		return "required"
	}
	return ""
}

func (*ExecutionReportTradeOutright525NoFills) FillYieldTypeMinValue() uint8 {
	return 0
}

func (*ExecutionReportTradeOutright525NoFills) FillYieldTypeMaxValue() uint8 {
	return math.MaxUint8 - 1
}

func (*ExecutionReportTradeOutright525NoFills) FillYieldTypeNullValue() uint8 {
	return math.MaxUint8
}

func (*ExecutionReportTradeOutright525NoOrderEvents) OrderEventPxId() uint16 {
	return 1799
}

func (*ExecutionReportTradeOutright525NoOrderEvents) OrderEventPxSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeOutright525NoOrderEvents) OrderEventPxInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.OrderEventPxSinceVersion()
}

func (*ExecutionReportTradeOutright525NoOrderEvents) OrderEventPxDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeOutright525NoOrderEvents) OrderEventPxMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return "Price"
	case 4:
		return "required"
	}
	return ""
}

func (*ExecutionReportTradeOutright525NoOrderEvents) OrderEventTextId() uint16 {
	return 1802
}

func (*ExecutionReportTradeOutright525NoOrderEvents) OrderEventTextSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeOutright525NoOrderEvents) OrderEventTextInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.OrderEventTextSinceVersion()
}

func (*ExecutionReportTradeOutright525NoOrderEvents) OrderEventTextDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeOutright525NoOrderEvents) OrderEventTextMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return "String"
	case 4:
		return "optional"
	}
	return ""
}

func (*ExecutionReportTradeOutright525NoOrderEvents) OrderEventTextMinValue() byte {
	return byte(32)
}

func (*ExecutionReportTradeOutright525NoOrderEvents) OrderEventTextMaxValue() byte {
	return byte(126)
}

func (*ExecutionReportTradeOutright525NoOrderEvents) OrderEventTextNullValue() byte {
	return 0
}

func (e *ExecutionReportTradeOutright525NoOrderEvents) OrderEventTextCharacterEncoding() string {
	return "US-ASCII"
}

func (*ExecutionReportTradeOutright525NoOrderEvents) OrderEventExecIDId() uint16 {
	return 1797
}

func (*ExecutionReportTradeOutright525NoOrderEvents) OrderEventExecIDSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeOutright525NoOrderEvents) OrderEventExecIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.OrderEventExecIDSinceVersion()
}

func (*ExecutionReportTradeOutright525NoOrderEvents) OrderEventExecIDDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeOutright525NoOrderEvents) OrderEventExecIDMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return "int"
	case 4:
		return "required"
	}
	return ""
}

func (*ExecutionReportTradeOutright525NoOrderEvents) OrderEventExecIDMinValue() uint32 {
	return 0
}

func (*ExecutionReportTradeOutright525NoOrderEvents) OrderEventExecIDMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*ExecutionReportTradeOutright525NoOrderEvents) OrderEventExecIDNullValue() uint32 {
	return math.MaxUint32
}

func (*ExecutionReportTradeOutright525NoOrderEvents) OrderEventQtyId() uint16 {
	return 1800
}

func (*ExecutionReportTradeOutright525NoOrderEvents) OrderEventQtySinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeOutright525NoOrderEvents) OrderEventQtyInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.OrderEventQtySinceVersion()
}

func (*ExecutionReportTradeOutright525NoOrderEvents) OrderEventQtyDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeOutright525NoOrderEvents) OrderEventQtyMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return "int"
	case 4:
		return "required"
	}
	return ""
}

func (*ExecutionReportTradeOutright525NoOrderEvents) OrderEventQtyMinValue() uint32 {
	return 0
}

func (*ExecutionReportTradeOutright525NoOrderEvents) OrderEventQtyMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*ExecutionReportTradeOutright525NoOrderEvents) OrderEventQtyNullValue() uint32 {
	return math.MaxUint32
}

func (*ExecutionReportTradeOutright525NoOrderEvents) OrderEventTypeId() uint16 {
	return 1796
}

func (*ExecutionReportTradeOutright525NoOrderEvents) OrderEventTypeSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeOutright525NoOrderEvents) OrderEventTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.OrderEventTypeSinceVersion()
}

func (*ExecutionReportTradeOutright525NoOrderEvents) OrderEventTypeDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeOutright525NoOrderEvents) OrderEventTypeMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return "int"
	case 4:
		return "required"
	}
	return ""
}

func (*ExecutionReportTradeOutright525NoOrderEvents) OrderEventReasonId() uint16 {
	return 1798
}

func (*ExecutionReportTradeOutright525NoOrderEvents) OrderEventReasonSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeOutright525NoOrderEvents) OrderEventReasonInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.OrderEventReasonSinceVersion()
}

func (*ExecutionReportTradeOutright525NoOrderEvents) OrderEventReasonDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeOutright525NoOrderEvents) OrderEventReasonMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return "int"
	case 4:
		return "required"
	}
	return ""
}

func (*ExecutionReportTradeOutright525NoOrderEvents) OrderEventReasonMinValue() uint8 {
	return 0
}

func (*ExecutionReportTradeOutright525NoOrderEvents) OrderEventReasonMaxValue() uint8 {
	return math.MaxUint8 - 1
}

func (*ExecutionReportTradeOutright525NoOrderEvents) OrderEventReasonNullValue() uint8 {
	return math.MaxUint8
}

func (*ExecutionReportTradeOutright525) NoFillsId() uint16 {
	return 1362
}

func (*ExecutionReportTradeOutright525) NoFillsSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeOutright525) NoFillsInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.NoFillsSinceVersion()
}

func (*ExecutionReportTradeOutright525) NoFillsDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeOutright525NoFills) SbeBlockLength() (blockLength uint) {
	return 15
}

func (*ExecutionReportTradeOutright525NoFills) SbeSchemaVersion() (schemaVersion uint16) {
	return 5
}

func (*ExecutionReportTradeOutright525) NoOrderEventsId() uint16 {
	return 1795
}

func (*ExecutionReportTradeOutright525) NoOrderEventsSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeOutright525) NoOrderEventsInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.NoOrderEventsSinceVersion()
}

func (*ExecutionReportTradeOutright525) NoOrderEventsDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeOutright525NoOrderEvents) SbeBlockLength() (blockLength uint) {
	return 23
}

func (*ExecutionReportTradeOutright525NoOrderEvents) SbeSchemaVersion() (schemaVersion uint16) {
	return 5
}
