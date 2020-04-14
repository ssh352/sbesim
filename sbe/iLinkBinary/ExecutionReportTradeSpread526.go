// Generated SBE (Simple Binary Encoding) message codec

package ilinkbinary

import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
)

type ExecutionReportTradeSpread526 struct {
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
	TotalNumSecurities    uint8
	ExecInst              ExecInst
	ExecutionMode         ExecModeEnum
	LiquidityFlag         BooleanNULLEnum
	ShortSaleType         ShortSaleTypeEnum
	NoFills               []ExecutionReportTradeSpread526NoFills
	NoLegs                []ExecutionReportTradeSpread526NoLegs
	NoOrderEvents         []ExecutionReportTradeSpread526NoOrderEvents
}
type ExecutionReportTradeSpread526NoFills struct {
	FillPx        PRICE9
	FillQty       uint32
	FillExecID    [2]byte
	FillYieldType uint8
}
type ExecutionReportTradeSpread526NoLegs struct {
	LegExecID     uint64
	LegLastPx     PRICE9
	LegSecurityID int32
	LegTradeID    uint32
	LegLastQty    uint32
	LegSide       SideReqEnum
}
type ExecutionReportTradeSpread526NoOrderEvents struct {
	OrderEventPx     PRICE9
	OrderEventText   [5]byte
	OrderEventExecID uint32
	OrderEventQty    uint32
	OrderEventType   OrderEventTypeEnum
	OrderEventReason uint8
}

func (e *ExecutionReportTradeSpread526) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
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
	if err := _m.WriteUint8(_w, e.TotalNumSecurities); err != nil {
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
	if err := e.ShortSaleType.Encode(_m, _w); err != nil {
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
	var NoLegsBlockLength uint16 = 29
	if err := _m.WriteUint16(_w, NoLegsBlockLength); err != nil {
		return err
	}
	var NoLegsNumInGroup uint8 = uint8(len(e.NoLegs))
	if err := _m.WriteUint8(_w, NoLegsNumInGroup); err != nil {
		return err
	}
	for _, prop := range e.NoLegs {
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

func (e *ExecutionReportTradeSpread526) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
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
	if !e.TotalNumSecuritiesInActingVersion(actingVersion) {
		e.TotalNumSecurities = e.TotalNumSecuritiesNullValue()
	} else {
		if err := _m.ReadUint8(_r, &e.TotalNumSecurities); err != nil {
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
	if e.ShortSaleTypeInActingVersion(actingVersion) {
		if err := e.ShortSaleType.Decode(_m, _r, actingVersion); err != nil {
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
			e.NoFills = make([]ExecutionReportTradeSpread526NoFills, NoFillsNumInGroup)
		}
		for i, _ := range e.NoFills {
			if err := e.NoFills[i].Decode(_m, _r, actingVersion, uint(NoFillsBlockLength)); err != nil {
				return err
			}
		}
	}

	if e.NoLegsInActingVersion(actingVersion) {
		var NoLegsBlockLength uint16
		if err := _m.ReadUint16(_r, &NoLegsBlockLength); err != nil {
			return err
		}
		var NoLegsNumInGroup uint8
		if err := _m.ReadUint8(_r, &NoLegsNumInGroup); err != nil {
			return err
		}
		if cap(e.NoLegs) < int(NoLegsNumInGroup) {
			e.NoLegs = make([]ExecutionReportTradeSpread526NoLegs, NoLegsNumInGroup)
		}
		for i, _ := range e.NoLegs {
			if err := e.NoLegs[i].Decode(_m, _r, actingVersion, uint(NoLegsBlockLength)); err != nil {
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
			e.NoOrderEvents = make([]ExecutionReportTradeSpread526NoOrderEvents, NoOrderEventsNumInGroup)
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

func (e *ExecutionReportTradeSpread526) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
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
	if e.TotalNumSecuritiesInActingVersion(actingVersion) {
		if e.TotalNumSecurities < e.TotalNumSecuritiesMinValue() || e.TotalNumSecurities > e.TotalNumSecuritiesMaxValue() {
			return fmt.Errorf("Range check failed on e.TotalNumSecurities (%v < %v > %v)", e.TotalNumSecuritiesMinValue(), e.TotalNumSecurities, e.TotalNumSecuritiesMaxValue())
		}
	}
	if err := e.ExecutionMode.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if err := e.LiquidityFlag.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if err := e.ShortSaleType.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	for _, prop := range e.NoFills {
		if err := prop.RangeCheck(actingVersion, schemaVersion); err != nil {
			return err
		}
	}
	for _, prop := range e.NoLegs {
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

func ExecutionReportTradeSpread526Init(e *ExecutionReportTradeSpread526) {
	e.CrossID = 18446744073709551615
	e.HostCrossID = 18446744073709551615
	e.TradeDate = 65535
	e.ExpireDate = 65535
	e.ExecType[0] = 70
	e.CrossType = 255
	return
}

func (e *ExecutionReportTradeSpread526NoFills) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
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

func (e *ExecutionReportTradeSpread526NoFills) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint) error {
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

func (e *ExecutionReportTradeSpread526NoFills) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
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

func ExecutionReportTradeSpread526NoFillsInit(e *ExecutionReportTradeSpread526NoFills) {
	return
}

func (e *ExecutionReportTradeSpread526NoLegs) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint64(_w, e.LegExecID); err != nil {
		return err
	}
	if err := e.LegLastPx.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteInt32(_w, e.LegSecurityID); err != nil {
		return err
	}
	if err := _m.WriteUint32(_w, e.LegTradeID); err != nil {
		return err
	}
	if err := _m.WriteUint32(_w, e.LegLastQty); err != nil {
		return err
	}
	if err := e.LegSide.Encode(_m, _w); err != nil {
		return err
	}
	return nil
}

func (e *ExecutionReportTradeSpread526NoLegs) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint) error {
	if !e.LegExecIDInActingVersion(actingVersion) {
		e.LegExecID = e.LegExecIDNullValue()
	} else {
		if err := _m.ReadUint64(_r, &e.LegExecID); err != nil {
			return err
		}
	}
	if e.LegLastPxInActingVersion(actingVersion) {
		if err := e.LegLastPx.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if !e.LegSecurityIDInActingVersion(actingVersion) {
		e.LegSecurityID = e.LegSecurityIDNullValue()
	} else {
		if err := _m.ReadInt32(_r, &e.LegSecurityID); err != nil {
			return err
		}
	}
	if !e.LegTradeIDInActingVersion(actingVersion) {
		e.LegTradeID = e.LegTradeIDNullValue()
	} else {
		if err := _m.ReadUint32(_r, &e.LegTradeID); err != nil {
			return err
		}
	}
	if !e.LegLastQtyInActingVersion(actingVersion) {
		e.LegLastQty = e.LegLastQtyNullValue()
	} else {
		if err := _m.ReadUint32(_r, &e.LegLastQty); err != nil {
			return err
		}
	}
	if e.LegSideInActingVersion(actingVersion) {
		if err := e.LegSide.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if actingVersion > e.SbeSchemaVersion() && blockLength > e.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-e.SbeBlockLength()))
	}
	return nil
}

func (e *ExecutionReportTradeSpread526NoLegs) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if e.LegExecIDInActingVersion(actingVersion) {
		if e.LegExecID < e.LegExecIDMinValue() || e.LegExecID > e.LegExecIDMaxValue() {
			return fmt.Errorf("Range check failed on e.LegExecID (%v < %v > %v)", e.LegExecIDMinValue(), e.LegExecID, e.LegExecIDMaxValue())
		}
	}
	if e.LegSecurityIDInActingVersion(actingVersion) {
		if e.LegSecurityID < e.LegSecurityIDMinValue() || e.LegSecurityID > e.LegSecurityIDMaxValue() {
			return fmt.Errorf("Range check failed on e.LegSecurityID (%v < %v > %v)", e.LegSecurityIDMinValue(), e.LegSecurityID, e.LegSecurityIDMaxValue())
		}
	}
	if e.LegTradeIDInActingVersion(actingVersion) {
		if e.LegTradeID < e.LegTradeIDMinValue() || e.LegTradeID > e.LegTradeIDMaxValue() {
			return fmt.Errorf("Range check failed on e.LegTradeID (%v < %v > %v)", e.LegTradeIDMinValue(), e.LegTradeID, e.LegTradeIDMaxValue())
		}
	}
	if e.LegLastQtyInActingVersion(actingVersion) {
		if e.LegLastQty < e.LegLastQtyMinValue() || e.LegLastQty > e.LegLastQtyMaxValue() {
			return fmt.Errorf("Range check failed on e.LegLastQty (%v < %v > %v)", e.LegLastQtyMinValue(), e.LegLastQty, e.LegLastQtyMaxValue())
		}
	}
	if err := e.LegSide.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	return nil
}

func ExecutionReportTradeSpread526NoLegsInit(e *ExecutionReportTradeSpread526NoLegs) {
	return
}

func (e *ExecutionReportTradeSpread526NoOrderEvents) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
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

func (e *ExecutionReportTradeSpread526NoOrderEvents) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint) error {
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

func (e *ExecutionReportTradeSpread526NoOrderEvents) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
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

func ExecutionReportTradeSpread526NoOrderEventsInit(e *ExecutionReportTradeSpread526NoOrderEvents) {
	for idx := 0; idx < 5; idx++ {
		e.OrderEventText[idx] = 0
	}
	return
}

func (*ExecutionReportTradeSpread526) SbeBlockLength() (blockLength uint16) {
	return 230
}

func (*ExecutionReportTradeSpread526) SbeTemplateId() (templateId uint16) {
	return 526
}

func (*ExecutionReportTradeSpread526) SbeSchemaId() (schemaId uint16) {
	return 8
}

func (*ExecutionReportTradeSpread526) SbeSchemaVersion() (schemaVersion uint16) {
	return 5
}

func (*ExecutionReportTradeSpread526) SbeSemanticType() (semanticType []byte) {
	return []byte("8")
}

func (*ExecutionReportTradeSpread526) SeqNumId() uint16 {
	return 9726
}

func (*ExecutionReportTradeSpread526) SeqNumSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeSpread526) SeqNumInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.SeqNumSinceVersion()
}

func (*ExecutionReportTradeSpread526) SeqNumDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeSpread526) SeqNumMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeSpread526) SeqNumMinValue() uint32 {
	return 0
}

func (*ExecutionReportTradeSpread526) SeqNumMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*ExecutionReportTradeSpread526) SeqNumNullValue() uint32 {
	return math.MaxUint32
}

func (*ExecutionReportTradeSpread526) UUIDId() uint16 {
	return 39001
}

func (*ExecutionReportTradeSpread526) UUIDSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeSpread526) UUIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.UUIDSinceVersion()
}

func (*ExecutionReportTradeSpread526) UUIDDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeSpread526) UUIDMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeSpread526) UUIDMinValue() uint64 {
	return 0
}

func (*ExecutionReportTradeSpread526) UUIDMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*ExecutionReportTradeSpread526) UUIDNullValue() uint64 {
	return math.MaxUint64
}

func (*ExecutionReportTradeSpread526) ExecIDId() uint16 {
	return 17
}

func (*ExecutionReportTradeSpread526) ExecIDSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeSpread526) ExecIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.ExecIDSinceVersion()
}

func (*ExecutionReportTradeSpread526) ExecIDDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeSpread526) ExecIDMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeSpread526) ExecIDMinValue() byte {
	return byte(32)
}

func (*ExecutionReportTradeSpread526) ExecIDMaxValue() byte {
	return byte(126)
}

func (*ExecutionReportTradeSpread526) ExecIDNullValue() byte {
	return 0
}

func (e *ExecutionReportTradeSpread526) ExecIDCharacterEncoding() string {
	return "US-ASCII"
}

func (*ExecutionReportTradeSpread526) SenderIDId() uint16 {
	return 5392
}

func (*ExecutionReportTradeSpread526) SenderIDSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeSpread526) SenderIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.SenderIDSinceVersion()
}

func (*ExecutionReportTradeSpread526) SenderIDDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeSpread526) SenderIDMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeSpread526) SenderIDMinValue() byte {
	return byte(32)
}

func (*ExecutionReportTradeSpread526) SenderIDMaxValue() byte {
	return byte(126)
}

func (*ExecutionReportTradeSpread526) SenderIDNullValue() byte {
	return 0
}

func (e *ExecutionReportTradeSpread526) SenderIDCharacterEncoding() string {
	return "US-ASCII"
}

func (*ExecutionReportTradeSpread526) ClOrdIDId() uint16 {
	return 11
}

func (*ExecutionReportTradeSpread526) ClOrdIDSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeSpread526) ClOrdIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.ClOrdIDSinceVersion()
}

func (*ExecutionReportTradeSpread526) ClOrdIDDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeSpread526) ClOrdIDMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeSpread526) ClOrdIDMinValue() byte {
	return byte(32)
}

func (*ExecutionReportTradeSpread526) ClOrdIDMaxValue() byte {
	return byte(126)
}

func (*ExecutionReportTradeSpread526) ClOrdIDNullValue() byte {
	return 0
}

func (e *ExecutionReportTradeSpread526) ClOrdIDCharacterEncoding() string {
	return "US-ASCII"
}

func (*ExecutionReportTradeSpread526) PartyDetailsListReqIDId() uint16 {
	return 1505
}

func (*ExecutionReportTradeSpread526) PartyDetailsListReqIDSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeSpread526) PartyDetailsListReqIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.PartyDetailsListReqIDSinceVersion()
}

func (*ExecutionReportTradeSpread526) PartyDetailsListReqIDDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeSpread526) PartyDetailsListReqIDMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeSpread526) PartyDetailsListReqIDMinValue() uint64 {
	return 0
}

func (*ExecutionReportTradeSpread526) PartyDetailsListReqIDMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*ExecutionReportTradeSpread526) PartyDetailsListReqIDNullValue() uint64 {
	return math.MaxUint64
}

func (*ExecutionReportTradeSpread526) LastPxId() uint16 {
	return 31
}

func (*ExecutionReportTradeSpread526) LastPxSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeSpread526) LastPxInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.LastPxSinceVersion()
}

func (*ExecutionReportTradeSpread526) LastPxDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeSpread526) LastPxMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeSpread526) OrderIDId() uint16 {
	return 37
}

func (*ExecutionReportTradeSpread526) OrderIDSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeSpread526) OrderIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.OrderIDSinceVersion()
}

func (*ExecutionReportTradeSpread526) OrderIDDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeSpread526) OrderIDMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeSpread526) OrderIDMinValue() uint64 {
	return 0
}

func (*ExecutionReportTradeSpread526) OrderIDMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*ExecutionReportTradeSpread526) OrderIDNullValue() uint64 {
	return math.MaxUint64
}

func (*ExecutionReportTradeSpread526) PriceId() uint16 {
	return 44
}

func (*ExecutionReportTradeSpread526) PriceSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeSpread526) PriceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.PriceSinceVersion()
}

func (*ExecutionReportTradeSpread526) PriceDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeSpread526) PriceMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeSpread526) StopPxId() uint16 {
	return 99
}

func (*ExecutionReportTradeSpread526) StopPxSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeSpread526) StopPxInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.StopPxSinceVersion()
}

func (*ExecutionReportTradeSpread526) StopPxDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeSpread526) StopPxMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeSpread526) TransactTimeId() uint16 {
	return 60
}

func (*ExecutionReportTradeSpread526) TransactTimeSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeSpread526) TransactTimeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.TransactTimeSinceVersion()
}

func (*ExecutionReportTradeSpread526) TransactTimeDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeSpread526) TransactTimeMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeSpread526) TransactTimeMinValue() uint64 {
	return 0
}

func (*ExecutionReportTradeSpread526) TransactTimeMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*ExecutionReportTradeSpread526) TransactTimeNullValue() uint64 {
	return math.MaxUint64
}

func (*ExecutionReportTradeSpread526) SendingTimeEpochId() uint16 {
	return 5297
}

func (*ExecutionReportTradeSpread526) SendingTimeEpochSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeSpread526) SendingTimeEpochInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.SendingTimeEpochSinceVersion()
}

func (*ExecutionReportTradeSpread526) SendingTimeEpochDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeSpread526) SendingTimeEpochMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeSpread526) SendingTimeEpochMinValue() uint64 {
	return 0
}

func (*ExecutionReportTradeSpread526) SendingTimeEpochMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*ExecutionReportTradeSpread526) SendingTimeEpochNullValue() uint64 {
	return math.MaxUint64
}

func (*ExecutionReportTradeSpread526) OrderRequestIDId() uint16 {
	return 2422
}

func (*ExecutionReportTradeSpread526) OrderRequestIDSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeSpread526) OrderRequestIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.OrderRequestIDSinceVersion()
}

func (*ExecutionReportTradeSpread526) OrderRequestIDDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeSpread526) OrderRequestIDMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeSpread526) OrderRequestIDMinValue() uint64 {
	return 0
}

func (*ExecutionReportTradeSpread526) OrderRequestIDMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*ExecutionReportTradeSpread526) OrderRequestIDNullValue() uint64 {
	return math.MaxUint64
}

func (*ExecutionReportTradeSpread526) SecExecIDId() uint16 {
	return 527
}

func (*ExecutionReportTradeSpread526) SecExecIDSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeSpread526) SecExecIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.SecExecIDSinceVersion()
}

func (*ExecutionReportTradeSpread526) SecExecIDDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeSpread526) SecExecIDMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeSpread526) SecExecIDMinValue() uint64 {
	return 0
}

func (*ExecutionReportTradeSpread526) SecExecIDMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*ExecutionReportTradeSpread526) SecExecIDNullValue() uint64 {
	return math.MaxUint64
}

func (*ExecutionReportTradeSpread526) CrossIDId() uint16 {
	return 548
}

func (*ExecutionReportTradeSpread526) CrossIDSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeSpread526) CrossIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.CrossIDSinceVersion()
}

func (*ExecutionReportTradeSpread526) CrossIDDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeSpread526) CrossIDMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeSpread526) CrossIDMinValue() uint64 {
	return 0
}

func (*ExecutionReportTradeSpread526) CrossIDMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*ExecutionReportTradeSpread526) CrossIDNullValue() uint64 {
	return 18446744073709551615
}

func (*ExecutionReportTradeSpread526) HostCrossIDId() uint16 {
	return 961
}

func (*ExecutionReportTradeSpread526) HostCrossIDSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeSpread526) HostCrossIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.HostCrossIDSinceVersion()
}

func (*ExecutionReportTradeSpread526) HostCrossIDDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeSpread526) HostCrossIDMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeSpread526) HostCrossIDMinValue() uint64 {
	return 0
}

func (*ExecutionReportTradeSpread526) HostCrossIDMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*ExecutionReportTradeSpread526) HostCrossIDNullValue() uint64 {
	return 18446744073709551615
}

func (*ExecutionReportTradeSpread526) LocationId() uint16 {
	return 9537
}

func (*ExecutionReportTradeSpread526) LocationSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeSpread526) LocationInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.LocationSinceVersion()
}

func (*ExecutionReportTradeSpread526) LocationDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeSpread526) LocationMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeSpread526) LocationMinValue() byte {
	return byte(32)
}

func (*ExecutionReportTradeSpread526) LocationMaxValue() byte {
	return byte(126)
}

func (*ExecutionReportTradeSpread526) LocationNullValue() byte {
	return 0
}

func (e *ExecutionReportTradeSpread526) LocationCharacterEncoding() string {
	return "US-ASCII"
}

func (*ExecutionReportTradeSpread526) SecurityIDId() uint16 {
	return 48
}

func (*ExecutionReportTradeSpread526) SecurityIDSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeSpread526) SecurityIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.SecurityIDSinceVersion()
}

func (*ExecutionReportTradeSpread526) SecurityIDDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeSpread526) SecurityIDMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeSpread526) SecurityIDMinValue() int32 {
	return math.MinInt32 + 1
}

func (*ExecutionReportTradeSpread526) SecurityIDMaxValue() int32 {
	return math.MaxInt32
}

func (*ExecutionReportTradeSpread526) SecurityIDNullValue() int32 {
	return math.MinInt32
}

func (*ExecutionReportTradeSpread526) OrderQtyId() uint16 {
	return 38
}

func (*ExecutionReportTradeSpread526) OrderQtySinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeSpread526) OrderQtyInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.OrderQtySinceVersion()
}

func (*ExecutionReportTradeSpread526) OrderQtyDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeSpread526) OrderQtyMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeSpread526) OrderQtyMinValue() uint32 {
	return 0
}

func (*ExecutionReportTradeSpread526) OrderQtyMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*ExecutionReportTradeSpread526) OrderQtyNullValue() uint32 {
	return math.MaxUint32
}

func (*ExecutionReportTradeSpread526) LastQtyId() uint16 {
	return 32
}

func (*ExecutionReportTradeSpread526) LastQtySinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeSpread526) LastQtyInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.LastQtySinceVersion()
}

func (*ExecutionReportTradeSpread526) LastQtyDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeSpread526) LastQtyMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeSpread526) LastQtyMinValue() uint32 {
	return 0
}

func (*ExecutionReportTradeSpread526) LastQtyMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*ExecutionReportTradeSpread526) LastQtyNullValue() uint32 {
	return math.MaxUint32
}

func (*ExecutionReportTradeSpread526) CumQtyId() uint16 {
	return 14
}

func (*ExecutionReportTradeSpread526) CumQtySinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeSpread526) CumQtyInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.CumQtySinceVersion()
}

func (*ExecutionReportTradeSpread526) CumQtyDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeSpread526) CumQtyMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeSpread526) CumQtyMinValue() uint32 {
	return 0
}

func (*ExecutionReportTradeSpread526) CumQtyMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*ExecutionReportTradeSpread526) CumQtyNullValue() uint32 {
	return math.MaxUint32
}

func (*ExecutionReportTradeSpread526) MDTradeEntryIDId() uint16 {
	return 37711
}

func (*ExecutionReportTradeSpread526) MDTradeEntryIDSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeSpread526) MDTradeEntryIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.MDTradeEntryIDSinceVersion()
}

func (*ExecutionReportTradeSpread526) MDTradeEntryIDDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeSpread526) MDTradeEntryIDMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeSpread526) MDTradeEntryIDMinValue() uint32 {
	return 0
}

func (*ExecutionReportTradeSpread526) MDTradeEntryIDMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*ExecutionReportTradeSpread526) MDTradeEntryIDNullValue() uint32 {
	return math.MaxUint32
}

func (*ExecutionReportTradeSpread526) SideTradeIDId() uint16 {
	return 1506
}

func (*ExecutionReportTradeSpread526) SideTradeIDSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeSpread526) SideTradeIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.SideTradeIDSinceVersion()
}

func (*ExecutionReportTradeSpread526) SideTradeIDDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeSpread526) SideTradeIDMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeSpread526) SideTradeIDMinValue() uint32 {
	return 0
}

func (*ExecutionReportTradeSpread526) SideTradeIDMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*ExecutionReportTradeSpread526) SideTradeIDNullValue() uint32 {
	return math.MaxUint32
}

func (*ExecutionReportTradeSpread526) LeavesQtyId() uint16 {
	return 151
}

func (*ExecutionReportTradeSpread526) LeavesQtySinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeSpread526) LeavesQtyInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.LeavesQtySinceVersion()
}

func (*ExecutionReportTradeSpread526) LeavesQtyDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeSpread526) LeavesQtyMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeSpread526) LeavesQtyMinValue() uint32 {
	return 0
}

func (*ExecutionReportTradeSpread526) LeavesQtyMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*ExecutionReportTradeSpread526) LeavesQtyNullValue() uint32 {
	return math.MaxUint32
}

func (*ExecutionReportTradeSpread526) TradeDateId() uint16 {
	return 75
}

func (*ExecutionReportTradeSpread526) TradeDateSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeSpread526) TradeDateInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.TradeDateSinceVersion()
}

func (*ExecutionReportTradeSpread526) TradeDateDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeSpread526) TradeDateMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeSpread526) TradeDateMinValue() uint16 {
	return 0
}

func (*ExecutionReportTradeSpread526) TradeDateMaxValue() uint16 {
	return math.MaxUint16 - 1
}

func (*ExecutionReportTradeSpread526) TradeDateNullValue() uint16 {
	return 65535
}

func (*ExecutionReportTradeSpread526) ExpireDateId() uint16 {
	return 432
}

func (*ExecutionReportTradeSpread526) ExpireDateSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeSpread526) ExpireDateInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.ExpireDateSinceVersion()
}

func (*ExecutionReportTradeSpread526) ExpireDateDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeSpread526) ExpireDateMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeSpread526) ExpireDateMinValue() uint16 {
	return 0
}

func (*ExecutionReportTradeSpread526) ExpireDateMaxValue() uint16 {
	return math.MaxUint16 - 1
}

func (*ExecutionReportTradeSpread526) ExpireDateNullValue() uint16 {
	return 65535
}

func (*ExecutionReportTradeSpread526) OrdStatusId() uint16 {
	return 39
}

func (*ExecutionReportTradeSpread526) OrdStatusSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeSpread526) OrdStatusInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.OrdStatusSinceVersion()
}

func (*ExecutionReportTradeSpread526) OrdStatusDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeSpread526) OrdStatusMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeSpread526) ExecTypeId() uint16 {
	return 150
}

func (*ExecutionReportTradeSpread526) ExecTypeSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeSpread526) ExecTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.ExecTypeSinceVersion()
}

func (*ExecutionReportTradeSpread526) ExecTypeDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeSpread526) ExecTypeMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeSpread526) ExecTypeMinValue() byte {
	return byte(32)
}

func (*ExecutionReportTradeSpread526) ExecTypeMaxValue() byte {
	return byte(126)
}

func (*ExecutionReportTradeSpread526) ExecTypeNullValue() byte {
	return 0
}

func (*ExecutionReportTradeSpread526) OrdTypeId() uint16 {
	return 40
}

func (*ExecutionReportTradeSpread526) OrdTypeSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeSpread526) OrdTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.OrdTypeSinceVersion()
}

func (*ExecutionReportTradeSpread526) OrdTypeDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeSpread526) OrdTypeMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeSpread526) SideId() uint16 {
	return 54
}

func (*ExecutionReportTradeSpread526) SideSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeSpread526) SideInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.SideSinceVersion()
}

func (*ExecutionReportTradeSpread526) SideDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeSpread526) SideMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeSpread526) TimeInForceId() uint16 {
	return 59
}

func (*ExecutionReportTradeSpread526) TimeInForceSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeSpread526) TimeInForceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.TimeInForceSinceVersion()
}

func (*ExecutionReportTradeSpread526) TimeInForceDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeSpread526) TimeInForceMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeSpread526) ManualOrderIndicatorId() uint16 {
	return 1028
}

func (*ExecutionReportTradeSpread526) ManualOrderIndicatorSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeSpread526) ManualOrderIndicatorInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.ManualOrderIndicatorSinceVersion()
}

func (*ExecutionReportTradeSpread526) ManualOrderIndicatorDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeSpread526) ManualOrderIndicatorMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeSpread526) PossRetransFlagId() uint16 {
	return 9765
}

func (*ExecutionReportTradeSpread526) PossRetransFlagSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeSpread526) PossRetransFlagInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.PossRetransFlagSinceVersion()
}

func (*ExecutionReportTradeSpread526) PossRetransFlagDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeSpread526) PossRetransFlagMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeSpread526) AggressorIndicatorId() uint16 {
	return 1057
}

func (*ExecutionReportTradeSpread526) AggressorIndicatorSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeSpread526) AggressorIndicatorInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.AggressorIndicatorSinceVersion()
}

func (*ExecutionReportTradeSpread526) AggressorIndicatorDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeSpread526) AggressorIndicatorMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeSpread526) CrossTypeId() uint16 {
	return 549
}

func (*ExecutionReportTradeSpread526) CrossTypeSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeSpread526) CrossTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.CrossTypeSinceVersion()
}

func (*ExecutionReportTradeSpread526) CrossTypeDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeSpread526) CrossTypeMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeSpread526) CrossTypeMinValue() uint8 {
	return 0
}

func (*ExecutionReportTradeSpread526) CrossTypeMaxValue() uint8 {
	return math.MaxUint8 - 1
}

func (*ExecutionReportTradeSpread526) CrossTypeNullValue() uint8 {
	return 255
}

func (*ExecutionReportTradeSpread526) TotalNumSecuritiesId() uint16 {
	return 393
}

func (*ExecutionReportTradeSpread526) TotalNumSecuritiesSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeSpread526) TotalNumSecuritiesInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.TotalNumSecuritiesSinceVersion()
}

func (*ExecutionReportTradeSpread526) TotalNumSecuritiesDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeSpread526) TotalNumSecuritiesMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeSpread526) TotalNumSecuritiesMinValue() uint8 {
	return 0
}

func (*ExecutionReportTradeSpread526) TotalNumSecuritiesMaxValue() uint8 {
	return math.MaxUint8 - 1
}

func (*ExecutionReportTradeSpread526) TotalNumSecuritiesNullValue() uint8 {
	return math.MaxUint8
}

func (*ExecutionReportTradeSpread526) ExecInstId() uint16 {
	return 18
}

func (*ExecutionReportTradeSpread526) ExecInstSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeSpread526) ExecInstInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.ExecInstSinceVersion()
}

func (*ExecutionReportTradeSpread526) ExecInstDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeSpread526) ExecInstMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeSpread526) ExecutionModeId() uint16 {
	return 5906
}

func (*ExecutionReportTradeSpread526) ExecutionModeSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeSpread526) ExecutionModeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.ExecutionModeSinceVersion()
}

func (*ExecutionReportTradeSpread526) ExecutionModeDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeSpread526) ExecutionModeMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeSpread526) LiquidityFlagId() uint16 {
	return 9373
}

func (*ExecutionReportTradeSpread526) LiquidityFlagSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeSpread526) LiquidityFlagInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.LiquidityFlagSinceVersion()
}

func (*ExecutionReportTradeSpread526) LiquidityFlagDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeSpread526) LiquidityFlagMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeSpread526) ShortSaleTypeId() uint16 {
	return 5409
}

func (*ExecutionReportTradeSpread526) ShortSaleTypeSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeSpread526) ShortSaleTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.ShortSaleTypeSinceVersion()
}

func (*ExecutionReportTradeSpread526) ShortSaleTypeDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeSpread526) ShortSaleTypeMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeSpread526NoFills) FillPxId() uint16 {
	return 1364
}

func (*ExecutionReportTradeSpread526NoFills) FillPxSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeSpread526NoFills) FillPxInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.FillPxSinceVersion()
}

func (*ExecutionReportTradeSpread526NoFills) FillPxDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeSpread526NoFills) FillPxMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeSpread526NoFills) FillQtyId() uint16 {
	return 1365
}

func (*ExecutionReportTradeSpread526NoFills) FillQtySinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeSpread526NoFills) FillQtyInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.FillQtySinceVersion()
}

func (*ExecutionReportTradeSpread526NoFills) FillQtyDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeSpread526NoFills) FillQtyMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeSpread526NoFills) FillQtyMinValue() uint32 {
	return 0
}

func (*ExecutionReportTradeSpread526NoFills) FillQtyMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*ExecutionReportTradeSpread526NoFills) FillQtyNullValue() uint32 {
	return math.MaxUint32
}

func (*ExecutionReportTradeSpread526NoFills) FillExecIDId() uint16 {
	return 1363
}

func (*ExecutionReportTradeSpread526NoFills) FillExecIDSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeSpread526NoFills) FillExecIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.FillExecIDSinceVersion()
}

func (*ExecutionReportTradeSpread526NoFills) FillExecIDDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeSpread526NoFills) FillExecIDMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeSpread526NoFills) FillExecIDMinValue() byte {
	return byte(32)
}

func (*ExecutionReportTradeSpread526NoFills) FillExecIDMaxValue() byte {
	return byte(126)
}

func (*ExecutionReportTradeSpread526NoFills) FillExecIDNullValue() byte {
	return 0
}

func (e *ExecutionReportTradeSpread526NoFills) FillExecIDCharacterEncoding() string {
	return "US-ASCII"
}

func (*ExecutionReportTradeSpread526NoFills) FillYieldTypeId() uint16 {
	return 1622
}

func (*ExecutionReportTradeSpread526NoFills) FillYieldTypeSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeSpread526NoFills) FillYieldTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.FillYieldTypeSinceVersion()
}

func (*ExecutionReportTradeSpread526NoFills) FillYieldTypeDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeSpread526NoFills) FillYieldTypeMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeSpread526NoFills) FillYieldTypeMinValue() uint8 {
	return 0
}

func (*ExecutionReportTradeSpread526NoFills) FillYieldTypeMaxValue() uint8 {
	return math.MaxUint8 - 1
}

func (*ExecutionReportTradeSpread526NoFills) FillYieldTypeNullValue() uint8 {
	return math.MaxUint8
}

func (*ExecutionReportTradeSpread526NoLegs) LegExecIDId() uint16 {
	return 1893
}

func (*ExecutionReportTradeSpread526NoLegs) LegExecIDSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeSpread526NoLegs) LegExecIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.LegExecIDSinceVersion()
}

func (*ExecutionReportTradeSpread526NoLegs) LegExecIDDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeSpread526NoLegs) LegExecIDMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeSpread526NoLegs) LegExecIDMinValue() uint64 {
	return 0
}

func (*ExecutionReportTradeSpread526NoLegs) LegExecIDMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*ExecutionReportTradeSpread526NoLegs) LegExecIDNullValue() uint64 {
	return math.MaxUint64
}

func (*ExecutionReportTradeSpread526NoLegs) LegLastPxId() uint16 {
	return 637
}

func (*ExecutionReportTradeSpread526NoLegs) LegLastPxSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeSpread526NoLegs) LegLastPxInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.LegLastPxSinceVersion()
}

func (*ExecutionReportTradeSpread526NoLegs) LegLastPxDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeSpread526NoLegs) LegLastPxMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeSpread526NoLegs) LegSecurityIDId() uint16 {
	return 602
}

func (*ExecutionReportTradeSpread526NoLegs) LegSecurityIDSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeSpread526NoLegs) LegSecurityIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.LegSecurityIDSinceVersion()
}

func (*ExecutionReportTradeSpread526NoLegs) LegSecurityIDDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeSpread526NoLegs) LegSecurityIDMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeSpread526NoLegs) LegSecurityIDMinValue() int32 {
	return math.MinInt32 + 1
}

func (*ExecutionReportTradeSpread526NoLegs) LegSecurityIDMaxValue() int32 {
	return math.MaxInt32
}

func (*ExecutionReportTradeSpread526NoLegs) LegSecurityIDNullValue() int32 {
	return math.MinInt32
}

func (*ExecutionReportTradeSpread526NoLegs) LegTradeIDId() uint16 {
	return 1894
}

func (*ExecutionReportTradeSpread526NoLegs) LegTradeIDSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeSpread526NoLegs) LegTradeIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.LegTradeIDSinceVersion()
}

func (*ExecutionReportTradeSpread526NoLegs) LegTradeIDDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeSpread526NoLegs) LegTradeIDMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeSpread526NoLegs) LegTradeIDMinValue() uint32 {
	return 0
}

func (*ExecutionReportTradeSpread526NoLegs) LegTradeIDMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*ExecutionReportTradeSpread526NoLegs) LegTradeIDNullValue() uint32 {
	return math.MaxUint32
}

func (*ExecutionReportTradeSpread526NoLegs) LegLastQtyId() uint16 {
	return 1418
}

func (*ExecutionReportTradeSpread526NoLegs) LegLastQtySinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeSpread526NoLegs) LegLastQtyInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.LegLastQtySinceVersion()
}

func (*ExecutionReportTradeSpread526NoLegs) LegLastQtyDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeSpread526NoLegs) LegLastQtyMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeSpread526NoLegs) LegLastQtyMinValue() uint32 {
	return 0
}

func (*ExecutionReportTradeSpread526NoLegs) LegLastQtyMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*ExecutionReportTradeSpread526NoLegs) LegLastQtyNullValue() uint32 {
	return math.MaxUint32
}

func (*ExecutionReportTradeSpread526NoLegs) LegSideId() uint16 {
	return 624
}

func (*ExecutionReportTradeSpread526NoLegs) LegSideSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeSpread526NoLegs) LegSideInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.LegSideSinceVersion()
}

func (*ExecutionReportTradeSpread526NoLegs) LegSideDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeSpread526NoLegs) LegSideMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeSpread526NoOrderEvents) OrderEventPxId() uint16 {
	return 1799
}

func (*ExecutionReportTradeSpread526NoOrderEvents) OrderEventPxSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeSpread526NoOrderEvents) OrderEventPxInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.OrderEventPxSinceVersion()
}

func (*ExecutionReportTradeSpread526NoOrderEvents) OrderEventPxDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeSpread526NoOrderEvents) OrderEventPxMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeSpread526NoOrderEvents) OrderEventTextId() uint16 {
	return 1802
}

func (*ExecutionReportTradeSpread526NoOrderEvents) OrderEventTextSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeSpread526NoOrderEvents) OrderEventTextInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.OrderEventTextSinceVersion()
}

func (*ExecutionReportTradeSpread526NoOrderEvents) OrderEventTextDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeSpread526NoOrderEvents) OrderEventTextMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeSpread526NoOrderEvents) OrderEventTextMinValue() byte {
	return byte(32)
}

func (*ExecutionReportTradeSpread526NoOrderEvents) OrderEventTextMaxValue() byte {
	return byte(126)
}

func (*ExecutionReportTradeSpread526NoOrderEvents) OrderEventTextNullValue() byte {
	return 0
}

func (e *ExecutionReportTradeSpread526NoOrderEvents) OrderEventTextCharacterEncoding() string {
	return "US-ASCII"
}

func (*ExecutionReportTradeSpread526NoOrderEvents) OrderEventExecIDId() uint16 {
	return 1797
}

func (*ExecutionReportTradeSpread526NoOrderEvents) OrderEventExecIDSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeSpread526NoOrderEvents) OrderEventExecIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.OrderEventExecIDSinceVersion()
}

func (*ExecutionReportTradeSpread526NoOrderEvents) OrderEventExecIDDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeSpread526NoOrderEvents) OrderEventExecIDMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeSpread526NoOrderEvents) OrderEventExecIDMinValue() uint32 {
	return 0
}

func (*ExecutionReportTradeSpread526NoOrderEvents) OrderEventExecIDMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*ExecutionReportTradeSpread526NoOrderEvents) OrderEventExecIDNullValue() uint32 {
	return math.MaxUint32
}

func (*ExecutionReportTradeSpread526NoOrderEvents) OrderEventQtyId() uint16 {
	return 1800
}

func (*ExecutionReportTradeSpread526NoOrderEvents) OrderEventQtySinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeSpread526NoOrderEvents) OrderEventQtyInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.OrderEventQtySinceVersion()
}

func (*ExecutionReportTradeSpread526NoOrderEvents) OrderEventQtyDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeSpread526NoOrderEvents) OrderEventQtyMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeSpread526NoOrderEvents) OrderEventQtyMinValue() uint32 {
	return 0
}

func (*ExecutionReportTradeSpread526NoOrderEvents) OrderEventQtyMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*ExecutionReportTradeSpread526NoOrderEvents) OrderEventQtyNullValue() uint32 {
	return math.MaxUint32
}

func (*ExecutionReportTradeSpread526NoOrderEvents) OrderEventTypeId() uint16 {
	return 1796
}

func (*ExecutionReportTradeSpread526NoOrderEvents) OrderEventTypeSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeSpread526NoOrderEvents) OrderEventTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.OrderEventTypeSinceVersion()
}

func (*ExecutionReportTradeSpread526NoOrderEvents) OrderEventTypeDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeSpread526NoOrderEvents) OrderEventTypeMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeSpread526NoOrderEvents) OrderEventReasonId() uint16 {
	return 1798
}

func (*ExecutionReportTradeSpread526NoOrderEvents) OrderEventReasonSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeSpread526NoOrderEvents) OrderEventReasonInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.OrderEventReasonSinceVersion()
}

func (*ExecutionReportTradeSpread526NoOrderEvents) OrderEventReasonDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeSpread526NoOrderEvents) OrderEventReasonMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeSpread526NoOrderEvents) OrderEventReasonMinValue() uint8 {
	return 0
}

func (*ExecutionReportTradeSpread526NoOrderEvents) OrderEventReasonMaxValue() uint8 {
	return math.MaxUint8 - 1
}

func (*ExecutionReportTradeSpread526NoOrderEvents) OrderEventReasonNullValue() uint8 {
	return math.MaxUint8
}

func (*ExecutionReportTradeSpread526) NoFillsId() uint16 {
	return 1362
}

func (*ExecutionReportTradeSpread526) NoFillsSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeSpread526) NoFillsInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.NoFillsSinceVersion()
}

func (*ExecutionReportTradeSpread526) NoFillsDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeSpread526NoFills) SbeBlockLength() (blockLength uint) {
	return 15
}

func (*ExecutionReportTradeSpread526NoFills) SbeSchemaVersion() (schemaVersion uint16) {
	return 5
}

func (*ExecutionReportTradeSpread526) NoLegsId() uint16 {
	return 555
}

func (*ExecutionReportTradeSpread526) NoLegsSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeSpread526) NoLegsInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.NoLegsSinceVersion()
}

func (*ExecutionReportTradeSpread526) NoLegsDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeSpread526NoLegs) SbeBlockLength() (blockLength uint) {
	return 29
}

func (*ExecutionReportTradeSpread526NoLegs) SbeSchemaVersion() (schemaVersion uint16) {
	return 5
}

func (*ExecutionReportTradeSpread526) NoOrderEventsId() uint16 {
	return 1795
}

func (*ExecutionReportTradeSpread526) NoOrderEventsSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeSpread526) NoOrderEventsInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.NoOrderEventsSinceVersion()
}

func (*ExecutionReportTradeSpread526) NoOrderEventsDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeSpread526NoOrderEvents) SbeBlockLength() (blockLength uint) {
	return 23
}

func (*ExecutionReportTradeSpread526NoOrderEvents) SbeSchemaVersion() (schemaVersion uint16) {
	return 5
}
