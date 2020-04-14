// Generated SBE (Simple Binary Encoding) message codec

package ilinkbinary

import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
)

type ExecutionReportTradeAddendumSpread549 struct {
	SeqNum                   uint32
	UUID                     uint64
	ExecID                   [40]byte
	SenderID                 [20]byte
	ClOrdID                  [20]byte
	PartyDetailsListReqID    uint64
	LastPx                   PRICE9
	OrderID                  uint64
	TransactTime             uint64
	SendingTimeEpoch         uint64
	SecExecID                uint64
	OrigSecondaryExecutionID uint64
	Location                 [5]byte
	SecurityID               int32
	MDTradeEntryID           uint32
	LastQty                  uint32
	SideTradeID              uint32
	OrigSideTradeID          uint32
	TradeDate                uint16
	OrdStatus                OrdStatusTrdCxlEnum
	ExecType                 ExecTypTrdCxlEnum
	OrdType                  OrderTypeEnum
	Side                     SideReqEnum
	ManualOrderIndicator     ManualOrdIndReqEnum
	PossRetransFlag          BooleanFlagEnum
	TotalNumSecurities       uint8
	ExecInst                 ExecInst
	ExecutionMode            ExecModeEnum
	LiquidityFlag            BooleanNULLEnum
	ManagedOrder             BooleanNULLEnum
	ShortSaleType            ShortSaleTypeEnum
	NoFills                  []ExecutionReportTradeAddendumSpread549NoFills
	NoLegs                   []ExecutionReportTradeAddendumSpread549NoLegs
	NoOrderEvents            []ExecutionReportTradeAddendumSpread549NoOrderEvents
}
type ExecutionReportTradeAddendumSpread549NoFills struct {
	FillPx        PRICE9
	FillQty       uint32
	FillExecID    [2]byte
	FillYieldType uint8
}
type ExecutionReportTradeAddendumSpread549NoLegs struct {
	LegExecID     uint64
	LegLastPx     PRICE9
	LegExecRefID  uint64
	LegTradeID    uint32
	LegTradeRefID uint32
	LegSecurityID int32
	LegLastQty    uint32
	LegSide       SideReqEnum
}
type ExecutionReportTradeAddendumSpread549NoOrderEvents struct {
	OrderEventPx             PRICE9
	OrderEventText           [5]byte
	OrderEventExecID         uint32
	OrderEventQty            uint32
	OrderEventType           TradeAddendumEnum
	OrderEventReason         uint8
	OriginalOrderEventExecID uint32
}

func (e *ExecutionReportTradeAddendumSpread549) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
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
	if err := _m.WriteUint64(_w, e.TransactTime); err != nil {
		return err
	}
	if err := _m.WriteUint64(_w, e.SendingTimeEpoch); err != nil {
		return err
	}
	if err := _m.WriteUint64(_w, e.SecExecID); err != nil {
		return err
	}
	if err := _m.WriteUint64(_w, e.OrigSecondaryExecutionID); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, e.Location[:]); err != nil {
		return err
	}
	if err := _m.WriteInt32(_w, e.SecurityID); err != nil {
		return err
	}
	if err := _m.WriteUint32(_w, e.MDTradeEntryID); err != nil {
		return err
	}
	if err := _m.WriteUint32(_w, e.LastQty); err != nil {
		return err
	}
	if err := _m.WriteUint32(_w, e.SideTradeID); err != nil {
		return err
	}
	if err := _m.WriteUint32(_w, e.OrigSideTradeID); err != nil {
		return err
	}
	if err := _m.WriteUint16(_w, e.TradeDate); err != nil {
		return err
	}
	if err := e.OrdStatus.Encode(_m, _w); err != nil {
		return err
	}
	if err := e.ExecType.Encode(_m, _w); err != nil {
		return err
	}
	if err := e.OrdType.Encode(_m, _w); err != nil {
		return err
	}
	if err := e.Side.Encode(_m, _w); err != nil {
		return err
	}
	if err := e.ManualOrderIndicator.Encode(_m, _w); err != nil {
		return err
	}
	if err := e.PossRetransFlag.Encode(_m, _w); err != nil {
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
	if err := e.ManagedOrder.Encode(_m, _w); err != nil {
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
	var NoLegsBlockLength uint16 = 41
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
	var NoOrderEventsBlockLength uint16 = 27
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

func (e *ExecutionReportTradeAddendumSpread549) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
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
	if !e.SecExecIDInActingVersion(actingVersion) {
		e.SecExecID = e.SecExecIDNullValue()
	} else {
		if err := _m.ReadUint64(_r, &e.SecExecID); err != nil {
			return err
		}
	}
	if !e.OrigSecondaryExecutionIDInActingVersion(actingVersion) {
		e.OrigSecondaryExecutionID = e.OrigSecondaryExecutionIDNullValue()
	} else {
		if err := _m.ReadUint64(_r, &e.OrigSecondaryExecutionID); err != nil {
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
	if !e.MDTradeEntryIDInActingVersion(actingVersion) {
		e.MDTradeEntryID = e.MDTradeEntryIDNullValue()
	} else {
		if err := _m.ReadUint32(_r, &e.MDTradeEntryID); err != nil {
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
	if !e.SideTradeIDInActingVersion(actingVersion) {
		e.SideTradeID = e.SideTradeIDNullValue()
	} else {
		if err := _m.ReadUint32(_r, &e.SideTradeID); err != nil {
			return err
		}
	}
	if !e.OrigSideTradeIDInActingVersion(actingVersion) {
		e.OrigSideTradeID = e.OrigSideTradeIDNullValue()
	} else {
		if err := _m.ReadUint32(_r, &e.OrigSideTradeID); err != nil {
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
	if e.OrdStatusInActingVersion(actingVersion) {
		if err := e.OrdStatus.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if e.ExecTypeInActingVersion(actingVersion) {
		if err := e.ExecType.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
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
			e.NoFills = make([]ExecutionReportTradeAddendumSpread549NoFills, NoFillsNumInGroup)
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
			e.NoLegs = make([]ExecutionReportTradeAddendumSpread549NoLegs, NoLegsNumInGroup)
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
			e.NoOrderEvents = make([]ExecutionReportTradeAddendumSpread549NoOrderEvents, NoOrderEventsNumInGroup)
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

func (e *ExecutionReportTradeAddendumSpread549) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
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
	if e.SecExecIDInActingVersion(actingVersion) {
		if e.SecExecID < e.SecExecIDMinValue() || e.SecExecID > e.SecExecIDMaxValue() {
			return fmt.Errorf("Range check failed on e.SecExecID (%v < %v > %v)", e.SecExecIDMinValue(), e.SecExecID, e.SecExecIDMaxValue())
		}
	}
	if e.OrigSecondaryExecutionIDInActingVersion(actingVersion) {
		if e.OrigSecondaryExecutionID != e.OrigSecondaryExecutionIDNullValue() && (e.OrigSecondaryExecutionID < e.OrigSecondaryExecutionIDMinValue() || e.OrigSecondaryExecutionID > e.OrigSecondaryExecutionIDMaxValue()) {
			return fmt.Errorf("Range check failed on e.OrigSecondaryExecutionID (%v < %v > %v)", e.OrigSecondaryExecutionIDMinValue(), e.OrigSecondaryExecutionID, e.OrigSecondaryExecutionIDMaxValue())
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
	if e.MDTradeEntryIDInActingVersion(actingVersion) {
		if e.MDTradeEntryID < e.MDTradeEntryIDMinValue() || e.MDTradeEntryID > e.MDTradeEntryIDMaxValue() {
			return fmt.Errorf("Range check failed on e.MDTradeEntryID (%v < %v > %v)", e.MDTradeEntryIDMinValue(), e.MDTradeEntryID, e.MDTradeEntryIDMaxValue())
		}
	}
	if e.LastQtyInActingVersion(actingVersion) {
		if e.LastQty < e.LastQtyMinValue() || e.LastQty > e.LastQtyMaxValue() {
			return fmt.Errorf("Range check failed on e.LastQty (%v < %v > %v)", e.LastQtyMinValue(), e.LastQty, e.LastQtyMaxValue())
		}
	}
	if e.SideTradeIDInActingVersion(actingVersion) {
		if e.SideTradeID < e.SideTradeIDMinValue() || e.SideTradeID > e.SideTradeIDMaxValue() {
			return fmt.Errorf("Range check failed on e.SideTradeID (%v < %v > %v)", e.SideTradeIDMinValue(), e.SideTradeID, e.SideTradeIDMaxValue())
		}
	}
	if e.OrigSideTradeIDInActingVersion(actingVersion) {
		if e.OrigSideTradeID != e.OrigSideTradeIDNullValue() && (e.OrigSideTradeID < e.OrigSideTradeIDMinValue() || e.OrigSideTradeID > e.OrigSideTradeIDMaxValue()) {
			return fmt.Errorf("Range check failed on e.OrigSideTradeID (%v < %v > %v)", e.OrigSideTradeIDMinValue(), e.OrigSideTradeID, e.OrigSideTradeIDMaxValue())
		}
	}
	if e.TradeDateInActingVersion(actingVersion) {
		if e.TradeDate != e.TradeDateNullValue() && (e.TradeDate < e.TradeDateMinValue() || e.TradeDate > e.TradeDateMaxValue()) {
			return fmt.Errorf("Range check failed on e.TradeDate (%v < %v > %v)", e.TradeDateMinValue(), e.TradeDate, e.TradeDateMaxValue())
		}
	}
	if err := e.OrdStatus.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if err := e.ExecType.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if err := e.OrdType.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if err := e.Side.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if err := e.ManualOrderIndicator.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if err := e.PossRetransFlag.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
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
	if err := e.ManagedOrder.RangeCheck(actingVersion, schemaVersion); err != nil {
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

func ExecutionReportTradeAddendumSpread549Init(e *ExecutionReportTradeAddendumSpread549) {
	e.OrigSecondaryExecutionID = 18446744073709551615
	e.OrigSideTradeID = 4294967295
	e.TradeDate = 65535
	return
}

func (e *ExecutionReportTradeAddendumSpread549NoFills) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
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

func (e *ExecutionReportTradeAddendumSpread549NoFills) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint) error {
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

func (e *ExecutionReportTradeAddendumSpread549NoFills) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
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

func ExecutionReportTradeAddendumSpread549NoFillsInit(e *ExecutionReportTradeAddendumSpread549NoFills) {
	return
}

func (e *ExecutionReportTradeAddendumSpread549NoLegs) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint64(_w, e.LegExecID); err != nil {
		return err
	}
	if err := e.LegLastPx.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteUint64(_w, e.LegExecRefID); err != nil {
		return err
	}
	if err := _m.WriteUint32(_w, e.LegTradeID); err != nil {
		return err
	}
	if err := _m.WriteUint32(_w, e.LegTradeRefID); err != nil {
		return err
	}
	if err := _m.WriteInt32(_w, e.LegSecurityID); err != nil {
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

func (e *ExecutionReportTradeAddendumSpread549NoLegs) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint) error {
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
	if !e.LegExecRefIDInActingVersion(actingVersion) {
		e.LegExecRefID = e.LegExecRefIDNullValue()
	} else {
		if err := _m.ReadUint64(_r, &e.LegExecRefID); err != nil {
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
	if !e.LegTradeRefIDInActingVersion(actingVersion) {
		e.LegTradeRefID = e.LegTradeRefIDNullValue()
	} else {
		if err := _m.ReadUint32(_r, &e.LegTradeRefID); err != nil {
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

func (e *ExecutionReportTradeAddendumSpread549NoLegs) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if e.LegExecIDInActingVersion(actingVersion) {
		if e.LegExecID < e.LegExecIDMinValue() || e.LegExecID > e.LegExecIDMaxValue() {
			return fmt.Errorf("Range check failed on e.LegExecID (%v < %v > %v)", e.LegExecIDMinValue(), e.LegExecID, e.LegExecIDMaxValue())
		}
	}
	if e.LegExecRefIDInActingVersion(actingVersion) {
		if e.LegExecRefID != e.LegExecRefIDNullValue() && (e.LegExecRefID < e.LegExecRefIDMinValue() || e.LegExecRefID > e.LegExecRefIDMaxValue()) {
			return fmt.Errorf("Range check failed on e.LegExecRefID (%v < %v > %v)", e.LegExecRefIDMinValue(), e.LegExecRefID, e.LegExecRefIDMaxValue())
		}
	}
	if e.LegTradeIDInActingVersion(actingVersion) {
		if e.LegTradeID < e.LegTradeIDMinValue() || e.LegTradeID > e.LegTradeIDMaxValue() {
			return fmt.Errorf("Range check failed on e.LegTradeID (%v < %v > %v)", e.LegTradeIDMinValue(), e.LegTradeID, e.LegTradeIDMaxValue())
		}
	}
	if e.LegTradeRefIDInActingVersion(actingVersion) {
		if e.LegTradeRefID != e.LegTradeRefIDNullValue() && (e.LegTradeRefID < e.LegTradeRefIDMinValue() || e.LegTradeRefID > e.LegTradeRefIDMaxValue()) {
			return fmt.Errorf("Range check failed on e.LegTradeRefID (%v < %v > %v)", e.LegTradeRefIDMinValue(), e.LegTradeRefID, e.LegTradeRefIDMaxValue())
		}
	}
	if e.LegSecurityIDInActingVersion(actingVersion) {
		if e.LegSecurityID < e.LegSecurityIDMinValue() || e.LegSecurityID > e.LegSecurityIDMaxValue() {
			return fmt.Errorf("Range check failed on e.LegSecurityID (%v < %v > %v)", e.LegSecurityIDMinValue(), e.LegSecurityID, e.LegSecurityIDMaxValue())
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

func ExecutionReportTradeAddendumSpread549NoLegsInit(e *ExecutionReportTradeAddendumSpread549NoLegs) {
	e.LegExecRefID = 18446744073709551615
	e.LegTradeRefID = 4294967295
	return
}

func (e *ExecutionReportTradeAddendumSpread549NoOrderEvents) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
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
	if err := _m.WriteUint32(_w, e.OriginalOrderEventExecID); err != nil {
		return err
	}
	return nil
}

func (e *ExecutionReportTradeAddendumSpread549NoOrderEvents) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint) error {
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
	if !e.OriginalOrderEventExecIDInActingVersion(actingVersion) {
		e.OriginalOrderEventExecID = e.OriginalOrderEventExecIDNullValue()
	} else {
		if err := _m.ReadUint32(_r, &e.OriginalOrderEventExecID); err != nil {
			return err
		}
	}
	if actingVersion > e.SbeSchemaVersion() && blockLength > e.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-e.SbeBlockLength()))
	}
	return nil
}

func (e *ExecutionReportTradeAddendumSpread549NoOrderEvents) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
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
	if e.OriginalOrderEventExecIDInActingVersion(actingVersion) {
		if e.OriginalOrderEventExecID != e.OriginalOrderEventExecIDNullValue() && (e.OriginalOrderEventExecID < e.OriginalOrderEventExecIDMinValue() || e.OriginalOrderEventExecID > e.OriginalOrderEventExecIDMaxValue()) {
			return fmt.Errorf("Range check failed on e.OriginalOrderEventExecID (%v < %v > %v)", e.OriginalOrderEventExecIDMinValue(), e.OriginalOrderEventExecID, e.OriginalOrderEventExecIDMaxValue())
		}
	}
	return nil
}

func ExecutionReportTradeAddendumSpread549NoOrderEventsInit(e *ExecutionReportTradeAddendumSpread549NoOrderEvents) {
	for idx := 0; idx < 5; idx++ {
		e.OrderEventText[idx] = 0
	}
	e.OriginalOrderEventExecID = 4294967295
	return
}

func (*ExecutionReportTradeAddendumSpread549) SbeBlockLength() (blockLength uint16) {
	return 187
}

func (*ExecutionReportTradeAddendumSpread549) SbeTemplateId() (templateId uint16) {
	return 549
}

func (*ExecutionReportTradeAddendumSpread549) SbeSchemaId() (schemaId uint16) {
	return 8
}

func (*ExecutionReportTradeAddendumSpread549) SbeSchemaVersion() (schemaVersion uint16) {
	return 5
}

func (*ExecutionReportTradeAddendumSpread549) SbeSemanticType() (semanticType []byte) {
	return []byte("8")
}

func (*ExecutionReportTradeAddendumSpread549) SeqNumId() uint16 {
	return 9726
}

func (*ExecutionReportTradeAddendumSpread549) SeqNumSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeAddendumSpread549) SeqNumInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.SeqNumSinceVersion()
}

func (*ExecutionReportTradeAddendumSpread549) SeqNumDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeAddendumSpread549) SeqNumMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeAddendumSpread549) SeqNumMinValue() uint32 {
	return 0
}

func (*ExecutionReportTradeAddendumSpread549) SeqNumMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*ExecutionReportTradeAddendumSpread549) SeqNumNullValue() uint32 {
	return math.MaxUint32
}

func (*ExecutionReportTradeAddendumSpread549) UUIDId() uint16 {
	return 39001
}

func (*ExecutionReportTradeAddendumSpread549) UUIDSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeAddendumSpread549) UUIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.UUIDSinceVersion()
}

func (*ExecutionReportTradeAddendumSpread549) UUIDDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeAddendumSpread549) UUIDMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeAddendumSpread549) UUIDMinValue() uint64 {
	return 0
}

func (*ExecutionReportTradeAddendumSpread549) UUIDMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*ExecutionReportTradeAddendumSpread549) UUIDNullValue() uint64 {
	return math.MaxUint64
}

func (*ExecutionReportTradeAddendumSpread549) ExecIDId() uint16 {
	return 17
}

func (*ExecutionReportTradeAddendumSpread549) ExecIDSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeAddendumSpread549) ExecIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.ExecIDSinceVersion()
}

func (*ExecutionReportTradeAddendumSpread549) ExecIDDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeAddendumSpread549) ExecIDMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeAddendumSpread549) ExecIDMinValue() byte {
	return byte(32)
}

func (*ExecutionReportTradeAddendumSpread549) ExecIDMaxValue() byte {
	return byte(126)
}

func (*ExecutionReportTradeAddendumSpread549) ExecIDNullValue() byte {
	return 0
}

func (e *ExecutionReportTradeAddendumSpread549) ExecIDCharacterEncoding() string {
	return "US-ASCII"
}

func (*ExecutionReportTradeAddendumSpread549) SenderIDId() uint16 {
	return 5392
}

func (*ExecutionReportTradeAddendumSpread549) SenderIDSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeAddendumSpread549) SenderIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.SenderIDSinceVersion()
}

func (*ExecutionReportTradeAddendumSpread549) SenderIDDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeAddendumSpread549) SenderIDMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeAddendumSpread549) SenderIDMinValue() byte {
	return byte(32)
}

func (*ExecutionReportTradeAddendumSpread549) SenderIDMaxValue() byte {
	return byte(126)
}

func (*ExecutionReportTradeAddendumSpread549) SenderIDNullValue() byte {
	return 0
}

func (e *ExecutionReportTradeAddendumSpread549) SenderIDCharacterEncoding() string {
	return "US-ASCII"
}

func (*ExecutionReportTradeAddendumSpread549) ClOrdIDId() uint16 {
	return 11
}

func (*ExecutionReportTradeAddendumSpread549) ClOrdIDSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeAddendumSpread549) ClOrdIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.ClOrdIDSinceVersion()
}

func (*ExecutionReportTradeAddendumSpread549) ClOrdIDDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeAddendumSpread549) ClOrdIDMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeAddendumSpread549) ClOrdIDMinValue() byte {
	return byte(32)
}

func (*ExecutionReportTradeAddendumSpread549) ClOrdIDMaxValue() byte {
	return byte(126)
}

func (*ExecutionReportTradeAddendumSpread549) ClOrdIDNullValue() byte {
	return 0
}

func (e *ExecutionReportTradeAddendumSpread549) ClOrdIDCharacterEncoding() string {
	return "US-ASCII"
}

func (*ExecutionReportTradeAddendumSpread549) PartyDetailsListReqIDId() uint16 {
	return 1505
}

func (*ExecutionReportTradeAddendumSpread549) PartyDetailsListReqIDSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeAddendumSpread549) PartyDetailsListReqIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.PartyDetailsListReqIDSinceVersion()
}

func (*ExecutionReportTradeAddendumSpread549) PartyDetailsListReqIDDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeAddendumSpread549) PartyDetailsListReqIDMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeAddendumSpread549) PartyDetailsListReqIDMinValue() uint64 {
	return 0
}

func (*ExecutionReportTradeAddendumSpread549) PartyDetailsListReqIDMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*ExecutionReportTradeAddendumSpread549) PartyDetailsListReqIDNullValue() uint64 {
	return math.MaxUint64
}

func (*ExecutionReportTradeAddendumSpread549) LastPxId() uint16 {
	return 31
}

func (*ExecutionReportTradeAddendumSpread549) LastPxSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeAddendumSpread549) LastPxInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.LastPxSinceVersion()
}

func (*ExecutionReportTradeAddendumSpread549) LastPxDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeAddendumSpread549) LastPxMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeAddendumSpread549) OrderIDId() uint16 {
	return 37
}

func (*ExecutionReportTradeAddendumSpread549) OrderIDSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeAddendumSpread549) OrderIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.OrderIDSinceVersion()
}

func (*ExecutionReportTradeAddendumSpread549) OrderIDDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeAddendumSpread549) OrderIDMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeAddendumSpread549) OrderIDMinValue() uint64 {
	return 0
}

func (*ExecutionReportTradeAddendumSpread549) OrderIDMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*ExecutionReportTradeAddendumSpread549) OrderIDNullValue() uint64 {
	return math.MaxUint64
}

func (*ExecutionReportTradeAddendumSpread549) TransactTimeId() uint16 {
	return 60
}

func (*ExecutionReportTradeAddendumSpread549) TransactTimeSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeAddendumSpread549) TransactTimeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.TransactTimeSinceVersion()
}

func (*ExecutionReportTradeAddendumSpread549) TransactTimeDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeAddendumSpread549) TransactTimeMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeAddendumSpread549) TransactTimeMinValue() uint64 {
	return 0
}

func (*ExecutionReportTradeAddendumSpread549) TransactTimeMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*ExecutionReportTradeAddendumSpread549) TransactTimeNullValue() uint64 {
	return math.MaxUint64
}

func (*ExecutionReportTradeAddendumSpread549) SendingTimeEpochId() uint16 {
	return 5297
}

func (*ExecutionReportTradeAddendumSpread549) SendingTimeEpochSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeAddendumSpread549) SendingTimeEpochInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.SendingTimeEpochSinceVersion()
}

func (*ExecutionReportTradeAddendumSpread549) SendingTimeEpochDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeAddendumSpread549) SendingTimeEpochMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeAddendumSpread549) SendingTimeEpochMinValue() uint64 {
	return 0
}

func (*ExecutionReportTradeAddendumSpread549) SendingTimeEpochMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*ExecutionReportTradeAddendumSpread549) SendingTimeEpochNullValue() uint64 {
	return math.MaxUint64
}

func (*ExecutionReportTradeAddendumSpread549) SecExecIDId() uint16 {
	return 527
}

func (*ExecutionReportTradeAddendumSpread549) SecExecIDSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeAddendumSpread549) SecExecIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.SecExecIDSinceVersion()
}

func (*ExecutionReportTradeAddendumSpread549) SecExecIDDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeAddendumSpread549) SecExecIDMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeAddendumSpread549) SecExecIDMinValue() uint64 {
	return 0
}

func (*ExecutionReportTradeAddendumSpread549) SecExecIDMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*ExecutionReportTradeAddendumSpread549) SecExecIDNullValue() uint64 {
	return math.MaxUint64
}

func (*ExecutionReportTradeAddendumSpread549) OrigSecondaryExecutionIDId() uint16 {
	return 9703
}

func (*ExecutionReportTradeAddendumSpread549) OrigSecondaryExecutionIDSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeAddendumSpread549) OrigSecondaryExecutionIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.OrigSecondaryExecutionIDSinceVersion()
}

func (*ExecutionReportTradeAddendumSpread549) OrigSecondaryExecutionIDDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeAddendumSpread549) OrigSecondaryExecutionIDMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeAddendumSpread549) OrigSecondaryExecutionIDMinValue() uint64 {
	return 0
}

func (*ExecutionReportTradeAddendumSpread549) OrigSecondaryExecutionIDMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*ExecutionReportTradeAddendumSpread549) OrigSecondaryExecutionIDNullValue() uint64 {
	return 18446744073709551615
}

func (*ExecutionReportTradeAddendumSpread549) LocationId() uint16 {
	return 9537
}

func (*ExecutionReportTradeAddendumSpread549) LocationSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeAddendumSpread549) LocationInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.LocationSinceVersion()
}

func (*ExecutionReportTradeAddendumSpread549) LocationDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeAddendumSpread549) LocationMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeAddendumSpread549) LocationMinValue() byte {
	return byte(32)
}

func (*ExecutionReportTradeAddendumSpread549) LocationMaxValue() byte {
	return byte(126)
}

func (*ExecutionReportTradeAddendumSpread549) LocationNullValue() byte {
	return 0
}

func (e *ExecutionReportTradeAddendumSpread549) LocationCharacterEncoding() string {
	return "US-ASCII"
}

func (*ExecutionReportTradeAddendumSpread549) SecurityIDId() uint16 {
	return 48
}

func (*ExecutionReportTradeAddendumSpread549) SecurityIDSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeAddendumSpread549) SecurityIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.SecurityIDSinceVersion()
}

func (*ExecutionReportTradeAddendumSpread549) SecurityIDDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeAddendumSpread549) SecurityIDMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeAddendumSpread549) SecurityIDMinValue() int32 {
	return math.MinInt32 + 1
}

func (*ExecutionReportTradeAddendumSpread549) SecurityIDMaxValue() int32 {
	return math.MaxInt32
}

func (*ExecutionReportTradeAddendumSpread549) SecurityIDNullValue() int32 {
	return math.MinInt32
}

func (*ExecutionReportTradeAddendumSpread549) MDTradeEntryIDId() uint16 {
	return 37711
}

func (*ExecutionReportTradeAddendumSpread549) MDTradeEntryIDSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeAddendumSpread549) MDTradeEntryIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.MDTradeEntryIDSinceVersion()
}

func (*ExecutionReportTradeAddendumSpread549) MDTradeEntryIDDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeAddendumSpread549) MDTradeEntryIDMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeAddendumSpread549) MDTradeEntryIDMinValue() uint32 {
	return 0
}

func (*ExecutionReportTradeAddendumSpread549) MDTradeEntryIDMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*ExecutionReportTradeAddendumSpread549) MDTradeEntryIDNullValue() uint32 {
	return math.MaxUint32
}

func (*ExecutionReportTradeAddendumSpread549) LastQtyId() uint16 {
	return 32
}

func (*ExecutionReportTradeAddendumSpread549) LastQtySinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeAddendumSpread549) LastQtyInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.LastQtySinceVersion()
}

func (*ExecutionReportTradeAddendumSpread549) LastQtyDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeAddendumSpread549) LastQtyMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeAddendumSpread549) LastQtyMinValue() uint32 {
	return 0
}

func (*ExecutionReportTradeAddendumSpread549) LastQtyMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*ExecutionReportTradeAddendumSpread549) LastQtyNullValue() uint32 {
	return math.MaxUint32
}

func (*ExecutionReportTradeAddendumSpread549) SideTradeIDId() uint16 {
	return 1506
}

func (*ExecutionReportTradeAddendumSpread549) SideTradeIDSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeAddendumSpread549) SideTradeIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.SideTradeIDSinceVersion()
}

func (*ExecutionReportTradeAddendumSpread549) SideTradeIDDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeAddendumSpread549) SideTradeIDMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeAddendumSpread549) SideTradeIDMinValue() uint32 {
	return 0
}

func (*ExecutionReportTradeAddendumSpread549) SideTradeIDMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*ExecutionReportTradeAddendumSpread549) SideTradeIDNullValue() uint32 {
	return math.MaxUint32
}

func (*ExecutionReportTradeAddendumSpread549) OrigSideTradeIDId() uint16 {
	return 1507
}

func (*ExecutionReportTradeAddendumSpread549) OrigSideTradeIDSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeAddendumSpread549) OrigSideTradeIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.OrigSideTradeIDSinceVersion()
}

func (*ExecutionReportTradeAddendumSpread549) OrigSideTradeIDDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeAddendumSpread549) OrigSideTradeIDMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeAddendumSpread549) OrigSideTradeIDMinValue() uint32 {
	return 0
}

func (*ExecutionReportTradeAddendumSpread549) OrigSideTradeIDMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*ExecutionReportTradeAddendumSpread549) OrigSideTradeIDNullValue() uint32 {
	return 4294967295
}

func (*ExecutionReportTradeAddendumSpread549) TradeDateId() uint16 {
	return 75
}

func (*ExecutionReportTradeAddendumSpread549) TradeDateSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeAddendumSpread549) TradeDateInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.TradeDateSinceVersion()
}

func (*ExecutionReportTradeAddendumSpread549) TradeDateDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeAddendumSpread549) TradeDateMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeAddendumSpread549) TradeDateMinValue() uint16 {
	return 0
}

func (*ExecutionReportTradeAddendumSpread549) TradeDateMaxValue() uint16 {
	return math.MaxUint16 - 1
}

func (*ExecutionReportTradeAddendumSpread549) TradeDateNullValue() uint16 {
	return 65535
}

func (*ExecutionReportTradeAddendumSpread549) OrdStatusId() uint16 {
	return 39
}

func (*ExecutionReportTradeAddendumSpread549) OrdStatusSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeAddendumSpread549) OrdStatusInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.OrdStatusSinceVersion()
}

func (*ExecutionReportTradeAddendumSpread549) OrdStatusDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeAddendumSpread549) OrdStatusMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeAddendumSpread549) ExecTypeId() uint16 {
	return 150
}

func (*ExecutionReportTradeAddendumSpread549) ExecTypeSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeAddendumSpread549) ExecTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.ExecTypeSinceVersion()
}

func (*ExecutionReportTradeAddendumSpread549) ExecTypeDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeAddendumSpread549) ExecTypeMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeAddendumSpread549) OrdTypeId() uint16 {
	return 40
}

func (*ExecutionReportTradeAddendumSpread549) OrdTypeSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeAddendumSpread549) OrdTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.OrdTypeSinceVersion()
}

func (*ExecutionReportTradeAddendumSpread549) OrdTypeDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeAddendumSpread549) OrdTypeMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeAddendumSpread549) SideId() uint16 {
	return 54
}

func (*ExecutionReportTradeAddendumSpread549) SideSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeAddendumSpread549) SideInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.SideSinceVersion()
}

func (*ExecutionReportTradeAddendumSpread549) SideDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeAddendumSpread549) SideMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeAddendumSpread549) ManualOrderIndicatorId() uint16 {
	return 1028
}

func (*ExecutionReportTradeAddendumSpread549) ManualOrderIndicatorSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeAddendumSpread549) ManualOrderIndicatorInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.ManualOrderIndicatorSinceVersion()
}

func (*ExecutionReportTradeAddendumSpread549) ManualOrderIndicatorDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeAddendumSpread549) ManualOrderIndicatorMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeAddendumSpread549) PossRetransFlagId() uint16 {
	return 9765
}

func (*ExecutionReportTradeAddendumSpread549) PossRetransFlagSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeAddendumSpread549) PossRetransFlagInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.PossRetransFlagSinceVersion()
}

func (*ExecutionReportTradeAddendumSpread549) PossRetransFlagDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeAddendumSpread549) PossRetransFlagMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeAddendumSpread549) TotalNumSecuritiesId() uint16 {
	return 393
}

func (*ExecutionReportTradeAddendumSpread549) TotalNumSecuritiesSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeAddendumSpread549) TotalNumSecuritiesInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.TotalNumSecuritiesSinceVersion()
}

func (*ExecutionReportTradeAddendumSpread549) TotalNumSecuritiesDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeAddendumSpread549) TotalNumSecuritiesMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeAddendumSpread549) TotalNumSecuritiesMinValue() uint8 {
	return 0
}

func (*ExecutionReportTradeAddendumSpread549) TotalNumSecuritiesMaxValue() uint8 {
	return math.MaxUint8 - 1
}

func (*ExecutionReportTradeAddendumSpread549) TotalNumSecuritiesNullValue() uint8 {
	return math.MaxUint8
}

func (*ExecutionReportTradeAddendumSpread549) ExecInstId() uint16 {
	return 18
}

func (*ExecutionReportTradeAddendumSpread549) ExecInstSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeAddendumSpread549) ExecInstInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.ExecInstSinceVersion()
}

func (*ExecutionReportTradeAddendumSpread549) ExecInstDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeAddendumSpread549) ExecInstMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeAddendumSpread549) ExecutionModeId() uint16 {
	return 5906
}

func (*ExecutionReportTradeAddendumSpread549) ExecutionModeSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeAddendumSpread549) ExecutionModeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.ExecutionModeSinceVersion()
}

func (*ExecutionReportTradeAddendumSpread549) ExecutionModeDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeAddendumSpread549) ExecutionModeMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeAddendumSpread549) LiquidityFlagId() uint16 {
	return 9373
}

func (*ExecutionReportTradeAddendumSpread549) LiquidityFlagSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeAddendumSpread549) LiquidityFlagInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.LiquidityFlagSinceVersion()
}

func (*ExecutionReportTradeAddendumSpread549) LiquidityFlagDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeAddendumSpread549) LiquidityFlagMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeAddendumSpread549) ManagedOrderId() uint16 {
	return 6881
}

func (*ExecutionReportTradeAddendumSpread549) ManagedOrderSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeAddendumSpread549) ManagedOrderInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.ManagedOrderSinceVersion()
}

func (*ExecutionReportTradeAddendumSpread549) ManagedOrderDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeAddendumSpread549) ManagedOrderMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeAddendumSpread549) ShortSaleTypeId() uint16 {
	return 5409
}

func (*ExecutionReportTradeAddendumSpread549) ShortSaleTypeSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeAddendumSpread549) ShortSaleTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.ShortSaleTypeSinceVersion()
}

func (*ExecutionReportTradeAddendumSpread549) ShortSaleTypeDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeAddendumSpread549) ShortSaleTypeMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeAddendumSpread549NoFills) FillPxId() uint16 {
	return 1364
}

func (*ExecutionReportTradeAddendumSpread549NoFills) FillPxSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeAddendumSpread549NoFills) FillPxInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.FillPxSinceVersion()
}

func (*ExecutionReportTradeAddendumSpread549NoFills) FillPxDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeAddendumSpread549NoFills) FillPxMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeAddendumSpread549NoFills) FillQtyId() uint16 {
	return 1365
}

func (*ExecutionReportTradeAddendumSpread549NoFills) FillQtySinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeAddendumSpread549NoFills) FillQtyInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.FillQtySinceVersion()
}

func (*ExecutionReportTradeAddendumSpread549NoFills) FillQtyDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeAddendumSpread549NoFills) FillQtyMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeAddendumSpread549NoFills) FillQtyMinValue() uint32 {
	return 0
}

func (*ExecutionReportTradeAddendumSpread549NoFills) FillQtyMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*ExecutionReportTradeAddendumSpread549NoFills) FillQtyNullValue() uint32 {
	return math.MaxUint32
}

func (*ExecutionReportTradeAddendumSpread549NoFills) FillExecIDId() uint16 {
	return 1363
}

func (*ExecutionReportTradeAddendumSpread549NoFills) FillExecIDSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeAddendumSpread549NoFills) FillExecIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.FillExecIDSinceVersion()
}

func (*ExecutionReportTradeAddendumSpread549NoFills) FillExecIDDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeAddendumSpread549NoFills) FillExecIDMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeAddendumSpread549NoFills) FillExecIDMinValue() byte {
	return byte(32)
}

func (*ExecutionReportTradeAddendumSpread549NoFills) FillExecIDMaxValue() byte {
	return byte(126)
}

func (*ExecutionReportTradeAddendumSpread549NoFills) FillExecIDNullValue() byte {
	return 0
}

func (e *ExecutionReportTradeAddendumSpread549NoFills) FillExecIDCharacterEncoding() string {
	return "US-ASCII"
}

func (*ExecutionReportTradeAddendumSpread549NoFills) FillYieldTypeId() uint16 {
	return 1622
}

func (*ExecutionReportTradeAddendumSpread549NoFills) FillYieldTypeSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeAddendumSpread549NoFills) FillYieldTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.FillYieldTypeSinceVersion()
}

func (*ExecutionReportTradeAddendumSpread549NoFills) FillYieldTypeDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeAddendumSpread549NoFills) FillYieldTypeMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeAddendumSpread549NoFills) FillYieldTypeMinValue() uint8 {
	return 0
}

func (*ExecutionReportTradeAddendumSpread549NoFills) FillYieldTypeMaxValue() uint8 {
	return math.MaxUint8 - 1
}

func (*ExecutionReportTradeAddendumSpread549NoFills) FillYieldTypeNullValue() uint8 {
	return math.MaxUint8
}

func (*ExecutionReportTradeAddendumSpread549NoLegs) LegExecIDId() uint16 {
	return 1893
}

func (*ExecutionReportTradeAddendumSpread549NoLegs) LegExecIDSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeAddendumSpread549NoLegs) LegExecIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.LegExecIDSinceVersion()
}

func (*ExecutionReportTradeAddendumSpread549NoLegs) LegExecIDDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeAddendumSpread549NoLegs) LegExecIDMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeAddendumSpread549NoLegs) LegExecIDMinValue() uint64 {
	return 0
}

func (*ExecutionReportTradeAddendumSpread549NoLegs) LegExecIDMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*ExecutionReportTradeAddendumSpread549NoLegs) LegExecIDNullValue() uint64 {
	return math.MaxUint64
}

func (*ExecutionReportTradeAddendumSpread549NoLegs) LegLastPxId() uint16 {
	return 637
}

func (*ExecutionReportTradeAddendumSpread549NoLegs) LegLastPxSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeAddendumSpread549NoLegs) LegLastPxInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.LegLastPxSinceVersion()
}

func (*ExecutionReportTradeAddendumSpread549NoLegs) LegLastPxDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeAddendumSpread549NoLegs) LegLastPxMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeAddendumSpread549NoLegs) LegExecRefIDId() uint16 {
	return 1901
}

func (*ExecutionReportTradeAddendumSpread549NoLegs) LegExecRefIDSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeAddendumSpread549NoLegs) LegExecRefIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.LegExecRefIDSinceVersion()
}

func (*ExecutionReportTradeAddendumSpread549NoLegs) LegExecRefIDDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeAddendumSpread549NoLegs) LegExecRefIDMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeAddendumSpread549NoLegs) LegExecRefIDMinValue() uint64 {
	return 0
}

func (*ExecutionReportTradeAddendumSpread549NoLegs) LegExecRefIDMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*ExecutionReportTradeAddendumSpread549NoLegs) LegExecRefIDNullValue() uint64 {
	return 18446744073709551615
}

func (*ExecutionReportTradeAddendumSpread549NoLegs) LegTradeIDId() uint16 {
	return 1894
}

func (*ExecutionReportTradeAddendumSpread549NoLegs) LegTradeIDSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeAddendumSpread549NoLegs) LegTradeIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.LegTradeIDSinceVersion()
}

func (*ExecutionReportTradeAddendumSpread549NoLegs) LegTradeIDDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeAddendumSpread549NoLegs) LegTradeIDMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeAddendumSpread549NoLegs) LegTradeIDMinValue() uint32 {
	return 0
}

func (*ExecutionReportTradeAddendumSpread549NoLegs) LegTradeIDMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*ExecutionReportTradeAddendumSpread549NoLegs) LegTradeIDNullValue() uint32 {
	return math.MaxUint32
}

func (*ExecutionReportTradeAddendumSpread549NoLegs) LegTradeRefIDId() uint16 {
	return 39023
}

func (*ExecutionReportTradeAddendumSpread549NoLegs) LegTradeRefIDSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeAddendumSpread549NoLegs) LegTradeRefIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.LegTradeRefIDSinceVersion()
}

func (*ExecutionReportTradeAddendumSpread549NoLegs) LegTradeRefIDDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeAddendumSpread549NoLegs) LegTradeRefIDMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeAddendumSpread549NoLegs) LegTradeRefIDMinValue() uint32 {
	return 0
}

func (*ExecutionReportTradeAddendumSpread549NoLegs) LegTradeRefIDMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*ExecutionReportTradeAddendumSpread549NoLegs) LegTradeRefIDNullValue() uint32 {
	return 4294967295
}

func (*ExecutionReportTradeAddendumSpread549NoLegs) LegSecurityIDId() uint16 {
	return 602
}

func (*ExecutionReportTradeAddendumSpread549NoLegs) LegSecurityIDSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeAddendumSpread549NoLegs) LegSecurityIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.LegSecurityIDSinceVersion()
}

func (*ExecutionReportTradeAddendumSpread549NoLegs) LegSecurityIDDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeAddendumSpread549NoLegs) LegSecurityIDMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeAddendumSpread549NoLegs) LegSecurityIDMinValue() int32 {
	return math.MinInt32 + 1
}

func (*ExecutionReportTradeAddendumSpread549NoLegs) LegSecurityIDMaxValue() int32 {
	return math.MaxInt32
}

func (*ExecutionReportTradeAddendumSpread549NoLegs) LegSecurityIDNullValue() int32 {
	return math.MinInt32
}

func (*ExecutionReportTradeAddendumSpread549NoLegs) LegLastQtyId() uint16 {
	return 1418
}

func (*ExecutionReportTradeAddendumSpread549NoLegs) LegLastQtySinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeAddendumSpread549NoLegs) LegLastQtyInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.LegLastQtySinceVersion()
}

func (*ExecutionReportTradeAddendumSpread549NoLegs) LegLastQtyDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeAddendumSpread549NoLegs) LegLastQtyMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeAddendumSpread549NoLegs) LegLastQtyMinValue() uint32 {
	return 0
}

func (*ExecutionReportTradeAddendumSpread549NoLegs) LegLastQtyMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*ExecutionReportTradeAddendumSpread549NoLegs) LegLastQtyNullValue() uint32 {
	return math.MaxUint32
}

func (*ExecutionReportTradeAddendumSpread549NoLegs) LegSideId() uint16 {
	return 624
}

func (*ExecutionReportTradeAddendumSpread549NoLegs) LegSideSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeAddendumSpread549NoLegs) LegSideInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.LegSideSinceVersion()
}

func (*ExecutionReportTradeAddendumSpread549NoLegs) LegSideDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeAddendumSpread549NoLegs) LegSideMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeAddendumSpread549NoOrderEvents) OrderEventPxId() uint16 {
	return 1799
}

func (*ExecutionReportTradeAddendumSpread549NoOrderEvents) OrderEventPxSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeAddendumSpread549NoOrderEvents) OrderEventPxInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.OrderEventPxSinceVersion()
}

func (*ExecutionReportTradeAddendumSpread549NoOrderEvents) OrderEventPxDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeAddendumSpread549NoOrderEvents) OrderEventPxMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeAddendumSpread549NoOrderEvents) OrderEventTextId() uint16 {
	return 1802
}

func (*ExecutionReportTradeAddendumSpread549NoOrderEvents) OrderEventTextSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeAddendumSpread549NoOrderEvents) OrderEventTextInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.OrderEventTextSinceVersion()
}

func (*ExecutionReportTradeAddendumSpread549NoOrderEvents) OrderEventTextDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeAddendumSpread549NoOrderEvents) OrderEventTextMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeAddendumSpread549NoOrderEvents) OrderEventTextMinValue() byte {
	return byte(32)
}

func (*ExecutionReportTradeAddendumSpread549NoOrderEvents) OrderEventTextMaxValue() byte {
	return byte(126)
}

func (*ExecutionReportTradeAddendumSpread549NoOrderEvents) OrderEventTextNullValue() byte {
	return 0
}

func (e *ExecutionReportTradeAddendumSpread549NoOrderEvents) OrderEventTextCharacterEncoding() string {
	return "US-ASCII"
}

func (*ExecutionReportTradeAddendumSpread549NoOrderEvents) OrderEventExecIDId() uint16 {
	return 1797
}

func (*ExecutionReportTradeAddendumSpread549NoOrderEvents) OrderEventExecIDSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeAddendumSpread549NoOrderEvents) OrderEventExecIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.OrderEventExecIDSinceVersion()
}

func (*ExecutionReportTradeAddendumSpread549NoOrderEvents) OrderEventExecIDDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeAddendumSpread549NoOrderEvents) OrderEventExecIDMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeAddendumSpread549NoOrderEvents) OrderEventExecIDMinValue() uint32 {
	return 0
}

func (*ExecutionReportTradeAddendumSpread549NoOrderEvents) OrderEventExecIDMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*ExecutionReportTradeAddendumSpread549NoOrderEvents) OrderEventExecIDNullValue() uint32 {
	return math.MaxUint32
}

func (*ExecutionReportTradeAddendumSpread549NoOrderEvents) OrderEventQtyId() uint16 {
	return 1800
}

func (*ExecutionReportTradeAddendumSpread549NoOrderEvents) OrderEventQtySinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeAddendumSpread549NoOrderEvents) OrderEventQtyInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.OrderEventQtySinceVersion()
}

func (*ExecutionReportTradeAddendumSpread549NoOrderEvents) OrderEventQtyDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeAddendumSpread549NoOrderEvents) OrderEventQtyMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeAddendumSpread549NoOrderEvents) OrderEventQtyMinValue() uint32 {
	return 0
}

func (*ExecutionReportTradeAddendumSpread549NoOrderEvents) OrderEventQtyMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*ExecutionReportTradeAddendumSpread549NoOrderEvents) OrderEventQtyNullValue() uint32 {
	return math.MaxUint32
}

func (*ExecutionReportTradeAddendumSpread549NoOrderEvents) OrderEventTypeId() uint16 {
	return 1796
}

func (*ExecutionReportTradeAddendumSpread549NoOrderEvents) OrderEventTypeSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeAddendumSpread549NoOrderEvents) OrderEventTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.OrderEventTypeSinceVersion()
}

func (*ExecutionReportTradeAddendumSpread549NoOrderEvents) OrderEventTypeDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeAddendumSpread549NoOrderEvents) OrderEventTypeMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeAddendumSpread549NoOrderEvents) OrderEventReasonId() uint16 {
	return 1798
}

func (*ExecutionReportTradeAddendumSpread549NoOrderEvents) OrderEventReasonSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeAddendumSpread549NoOrderEvents) OrderEventReasonInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.OrderEventReasonSinceVersion()
}

func (*ExecutionReportTradeAddendumSpread549NoOrderEvents) OrderEventReasonDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeAddendumSpread549NoOrderEvents) OrderEventReasonMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeAddendumSpread549NoOrderEvents) OrderEventReasonMinValue() uint8 {
	return 0
}

func (*ExecutionReportTradeAddendumSpread549NoOrderEvents) OrderEventReasonMaxValue() uint8 {
	return math.MaxUint8 - 1
}

func (*ExecutionReportTradeAddendumSpread549NoOrderEvents) OrderEventReasonNullValue() uint8 {
	return math.MaxUint8
}

func (*ExecutionReportTradeAddendumSpread549NoOrderEvents) OriginalOrderEventExecIDId() uint16 {
	return 6555
}

func (*ExecutionReportTradeAddendumSpread549NoOrderEvents) OriginalOrderEventExecIDSinceVersion() uint16 {
	return 3
}

func (e *ExecutionReportTradeAddendumSpread549NoOrderEvents) OriginalOrderEventExecIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.OriginalOrderEventExecIDSinceVersion()
}

func (*ExecutionReportTradeAddendumSpread549NoOrderEvents) OriginalOrderEventExecIDDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeAddendumSpread549NoOrderEvents) OriginalOrderEventExecIDMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeAddendumSpread549NoOrderEvents) OriginalOrderEventExecIDMinValue() uint32 {
	return 0
}

func (*ExecutionReportTradeAddendumSpread549NoOrderEvents) OriginalOrderEventExecIDMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*ExecutionReportTradeAddendumSpread549NoOrderEvents) OriginalOrderEventExecIDNullValue() uint32 {
	return 4294967295
}

func (*ExecutionReportTradeAddendumSpread549) NoFillsId() uint16 {
	return 1362
}

func (*ExecutionReportTradeAddendumSpread549) NoFillsSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeAddendumSpread549) NoFillsInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.NoFillsSinceVersion()
}

func (*ExecutionReportTradeAddendumSpread549) NoFillsDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeAddendumSpread549NoFills) SbeBlockLength() (blockLength uint) {
	return 15
}

func (*ExecutionReportTradeAddendumSpread549NoFills) SbeSchemaVersion() (schemaVersion uint16) {
	return 5
}

func (*ExecutionReportTradeAddendumSpread549) NoLegsId() uint16 {
	return 555
}

func (*ExecutionReportTradeAddendumSpread549) NoLegsSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeAddendumSpread549) NoLegsInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.NoLegsSinceVersion()
}

func (*ExecutionReportTradeAddendumSpread549) NoLegsDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeAddendumSpread549NoLegs) SbeBlockLength() (blockLength uint) {
	return 41
}

func (*ExecutionReportTradeAddendumSpread549NoLegs) SbeSchemaVersion() (schemaVersion uint16) {
	return 5
}

func (*ExecutionReportTradeAddendumSpread549) NoOrderEventsId() uint16 {
	return 1795
}

func (*ExecutionReportTradeAddendumSpread549) NoOrderEventsSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeAddendumSpread549) NoOrderEventsInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.NoOrderEventsSinceVersion()
}

func (*ExecutionReportTradeAddendumSpread549) NoOrderEventsDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeAddendumSpread549NoOrderEvents) SbeBlockLength() (blockLength uint) {
	return 27
}

func (*ExecutionReportTradeAddendumSpread549NoOrderEvents) SbeSchemaVersion() (schemaVersion uint16) {
	return 5
}
