// Generated SBE (Simple Binary Encoding) message codec

package ilinkbinary

import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
)

type ExecutionReportTradeAddendumOutright548 struct {
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
	LastQty                  uint32
	SideTradeID              uint32
	OrigSideTradeID          uint32
	TradeDate                uint16
	OrdStatus                OrdStatusTrdCxlEnum
	ExecType                 ExecTypTrdCxlEnum
	Side                     SideReqEnum
	ManualOrderIndicator     ManualOrdIndReqEnum
	PossRetransFlag          BooleanFlagEnum
	ExecInst                 ExecInst
	ExecutionMode            ExecModeEnum
	LiquidityFlag            BooleanNULLEnum
	ManagedOrder             BooleanNULLEnum
	ShortSaleType            ShortSaleTypeEnum
	NoFills                  []ExecutionReportTradeAddendumOutright548NoFills
	NoOrderEvents            []ExecutionReportTradeAddendumOutright548NoOrderEvents
}
type ExecutionReportTradeAddendumOutright548NoFills struct {
	FillPx        PRICE9
	FillQty       uint32
	FillExecID    [2]byte
	FillYieldType uint8
}
type ExecutionReportTradeAddendumOutright548NoOrderEvents struct {
	OrderEventPx             PRICE9
	OrderEventText           [5]byte
	OrderEventExecID         uint32
	OrderEventQty            uint32
	OrderEventType           TradeAddendumEnum
	OrderEventReason         uint8
	OriginalOrderEventExecID uint32
}

func (e *ExecutionReportTradeAddendumOutright548) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
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
	if err := e.Side.Encode(_m, _w); err != nil {
		return err
	}
	if err := e.ManualOrderIndicator.Encode(_m, _w); err != nil {
		return err
	}
	if err := e.PossRetransFlag.Encode(_m, _w); err != nil {
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

func (e *ExecutionReportTradeAddendumOutright548) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
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
			e.NoFills = make([]ExecutionReportTradeAddendumOutright548NoFills, NoFillsNumInGroup)
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
			e.NoOrderEvents = make([]ExecutionReportTradeAddendumOutright548NoOrderEvents, NoOrderEventsNumInGroup)
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

func (e *ExecutionReportTradeAddendumOutright548) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
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
	if err := e.Side.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if err := e.ManualOrderIndicator.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if err := e.PossRetransFlag.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
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
	for _, prop := range e.NoOrderEvents {
		if err := prop.RangeCheck(actingVersion, schemaVersion); err != nil {
			return err
		}
	}
	return nil
}

func ExecutionReportTradeAddendumOutright548Init(e *ExecutionReportTradeAddendumOutright548) {
	e.OrigSecondaryExecutionID = 18446744073709551615
	e.OrigSideTradeID = 4294967295
	e.TradeDate = 65535
	return
}

func (e *ExecutionReportTradeAddendumOutright548NoFills) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
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

func (e *ExecutionReportTradeAddendumOutright548NoFills) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint) error {
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

func (e *ExecutionReportTradeAddendumOutright548NoFills) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
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

func ExecutionReportTradeAddendumOutright548NoFillsInit(e *ExecutionReportTradeAddendumOutright548NoFills) {
	return
}

func (e *ExecutionReportTradeAddendumOutright548NoOrderEvents) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
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

func (e *ExecutionReportTradeAddendumOutright548NoOrderEvents) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint) error {
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

func (e *ExecutionReportTradeAddendumOutright548NoOrderEvents) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
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

func ExecutionReportTradeAddendumOutright548NoOrderEventsInit(e *ExecutionReportTradeAddendumOutright548NoOrderEvents) {
	for idx := 0; idx < 5; idx++ {
		e.OrderEventText[idx] = 0
	}
	e.OriginalOrderEventExecID = 4294967295
	return
}

func (*ExecutionReportTradeAddendumOutright548) SbeBlockLength() (blockLength uint16) {
	return 181
}

func (*ExecutionReportTradeAddendumOutright548) SbeTemplateId() (templateId uint16) {
	return 548
}

func (*ExecutionReportTradeAddendumOutright548) SbeSchemaId() (schemaId uint16) {
	return 8
}

func (*ExecutionReportTradeAddendumOutright548) SbeSchemaVersion() (schemaVersion uint16) {
	return 5
}

func (*ExecutionReportTradeAddendumOutright548) SbeSemanticType() (semanticType []byte) {
	return []byte("8")
}

func (*ExecutionReportTradeAddendumOutright548) SeqNumId() uint16 {
	return 9726
}

func (*ExecutionReportTradeAddendumOutright548) SeqNumSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeAddendumOutright548) SeqNumInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.SeqNumSinceVersion()
}

func (*ExecutionReportTradeAddendumOutright548) SeqNumDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeAddendumOutright548) SeqNumMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeAddendumOutright548) SeqNumMinValue() uint32 {
	return 0
}

func (*ExecutionReportTradeAddendumOutright548) SeqNumMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*ExecutionReportTradeAddendumOutright548) SeqNumNullValue() uint32 {
	return math.MaxUint32
}

func (*ExecutionReportTradeAddendumOutright548) UUIDId() uint16 {
	return 39001
}

func (*ExecutionReportTradeAddendumOutright548) UUIDSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeAddendumOutright548) UUIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.UUIDSinceVersion()
}

func (*ExecutionReportTradeAddendumOutright548) UUIDDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeAddendumOutright548) UUIDMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeAddendumOutright548) UUIDMinValue() uint64 {
	return 0
}

func (*ExecutionReportTradeAddendumOutright548) UUIDMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*ExecutionReportTradeAddendumOutright548) UUIDNullValue() uint64 {
	return math.MaxUint64
}

func (*ExecutionReportTradeAddendumOutright548) ExecIDId() uint16 {
	return 17
}

func (*ExecutionReportTradeAddendumOutright548) ExecIDSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeAddendumOutright548) ExecIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.ExecIDSinceVersion()
}

func (*ExecutionReportTradeAddendumOutright548) ExecIDDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeAddendumOutright548) ExecIDMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeAddendumOutright548) ExecIDMinValue() byte {
	return byte(32)
}

func (*ExecutionReportTradeAddendumOutright548) ExecIDMaxValue() byte {
	return byte(126)
}

func (*ExecutionReportTradeAddendumOutright548) ExecIDNullValue() byte {
	return 0
}

func (e *ExecutionReportTradeAddendumOutright548) ExecIDCharacterEncoding() string {
	return "US-ASCII"
}

func (*ExecutionReportTradeAddendumOutright548) SenderIDId() uint16 {
	return 5392
}

func (*ExecutionReportTradeAddendumOutright548) SenderIDSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeAddendumOutright548) SenderIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.SenderIDSinceVersion()
}

func (*ExecutionReportTradeAddendumOutright548) SenderIDDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeAddendumOutright548) SenderIDMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeAddendumOutright548) SenderIDMinValue() byte {
	return byte(32)
}

func (*ExecutionReportTradeAddendumOutright548) SenderIDMaxValue() byte {
	return byte(126)
}

func (*ExecutionReportTradeAddendumOutright548) SenderIDNullValue() byte {
	return 0
}

func (e *ExecutionReportTradeAddendumOutright548) SenderIDCharacterEncoding() string {
	return "US-ASCII"
}

func (*ExecutionReportTradeAddendumOutright548) ClOrdIDId() uint16 {
	return 11
}

func (*ExecutionReportTradeAddendumOutright548) ClOrdIDSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeAddendumOutright548) ClOrdIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.ClOrdIDSinceVersion()
}

func (*ExecutionReportTradeAddendumOutright548) ClOrdIDDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeAddendumOutright548) ClOrdIDMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeAddendumOutright548) ClOrdIDMinValue() byte {
	return byte(32)
}

func (*ExecutionReportTradeAddendumOutright548) ClOrdIDMaxValue() byte {
	return byte(126)
}

func (*ExecutionReportTradeAddendumOutright548) ClOrdIDNullValue() byte {
	return 0
}

func (e *ExecutionReportTradeAddendumOutright548) ClOrdIDCharacterEncoding() string {
	return "US-ASCII"
}

func (*ExecutionReportTradeAddendumOutright548) PartyDetailsListReqIDId() uint16 {
	return 1505
}

func (*ExecutionReportTradeAddendumOutright548) PartyDetailsListReqIDSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeAddendumOutright548) PartyDetailsListReqIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.PartyDetailsListReqIDSinceVersion()
}

func (*ExecutionReportTradeAddendumOutright548) PartyDetailsListReqIDDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeAddendumOutright548) PartyDetailsListReqIDMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeAddendumOutright548) PartyDetailsListReqIDMinValue() uint64 {
	return 0
}

func (*ExecutionReportTradeAddendumOutright548) PartyDetailsListReqIDMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*ExecutionReportTradeAddendumOutright548) PartyDetailsListReqIDNullValue() uint64 {
	return math.MaxUint64
}

func (*ExecutionReportTradeAddendumOutright548) LastPxId() uint16 {
	return 31
}

func (*ExecutionReportTradeAddendumOutright548) LastPxSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeAddendumOutright548) LastPxInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.LastPxSinceVersion()
}

func (*ExecutionReportTradeAddendumOutright548) LastPxDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeAddendumOutright548) LastPxMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeAddendumOutright548) OrderIDId() uint16 {
	return 37
}

func (*ExecutionReportTradeAddendumOutright548) OrderIDSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeAddendumOutright548) OrderIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.OrderIDSinceVersion()
}

func (*ExecutionReportTradeAddendumOutright548) OrderIDDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeAddendumOutright548) OrderIDMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeAddendumOutright548) OrderIDMinValue() uint64 {
	return 0
}

func (*ExecutionReportTradeAddendumOutright548) OrderIDMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*ExecutionReportTradeAddendumOutright548) OrderIDNullValue() uint64 {
	return math.MaxUint64
}

func (*ExecutionReportTradeAddendumOutright548) TransactTimeId() uint16 {
	return 60
}

func (*ExecutionReportTradeAddendumOutright548) TransactTimeSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeAddendumOutright548) TransactTimeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.TransactTimeSinceVersion()
}

func (*ExecutionReportTradeAddendumOutright548) TransactTimeDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeAddendumOutright548) TransactTimeMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeAddendumOutright548) TransactTimeMinValue() uint64 {
	return 0
}

func (*ExecutionReportTradeAddendumOutright548) TransactTimeMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*ExecutionReportTradeAddendumOutright548) TransactTimeNullValue() uint64 {
	return math.MaxUint64
}

func (*ExecutionReportTradeAddendumOutright548) SendingTimeEpochId() uint16 {
	return 5297
}

func (*ExecutionReportTradeAddendumOutright548) SendingTimeEpochSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeAddendumOutright548) SendingTimeEpochInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.SendingTimeEpochSinceVersion()
}

func (*ExecutionReportTradeAddendumOutright548) SendingTimeEpochDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeAddendumOutright548) SendingTimeEpochMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeAddendumOutright548) SendingTimeEpochMinValue() uint64 {
	return 0
}

func (*ExecutionReportTradeAddendumOutright548) SendingTimeEpochMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*ExecutionReportTradeAddendumOutright548) SendingTimeEpochNullValue() uint64 {
	return math.MaxUint64
}

func (*ExecutionReportTradeAddendumOutright548) SecExecIDId() uint16 {
	return 527
}

func (*ExecutionReportTradeAddendumOutright548) SecExecIDSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeAddendumOutright548) SecExecIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.SecExecIDSinceVersion()
}

func (*ExecutionReportTradeAddendumOutright548) SecExecIDDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeAddendumOutright548) SecExecIDMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeAddendumOutright548) SecExecIDMinValue() uint64 {
	return 0
}

func (*ExecutionReportTradeAddendumOutright548) SecExecIDMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*ExecutionReportTradeAddendumOutright548) SecExecIDNullValue() uint64 {
	return math.MaxUint64
}

func (*ExecutionReportTradeAddendumOutright548) OrigSecondaryExecutionIDId() uint16 {
	return 9703
}

func (*ExecutionReportTradeAddendumOutright548) OrigSecondaryExecutionIDSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeAddendumOutright548) OrigSecondaryExecutionIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.OrigSecondaryExecutionIDSinceVersion()
}

func (*ExecutionReportTradeAddendumOutright548) OrigSecondaryExecutionIDDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeAddendumOutright548) OrigSecondaryExecutionIDMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeAddendumOutright548) OrigSecondaryExecutionIDMinValue() uint64 {
	return 0
}

func (*ExecutionReportTradeAddendumOutright548) OrigSecondaryExecutionIDMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*ExecutionReportTradeAddendumOutright548) OrigSecondaryExecutionIDNullValue() uint64 {
	return 18446744073709551615
}

func (*ExecutionReportTradeAddendumOutright548) LocationId() uint16 {
	return 9537
}

func (*ExecutionReportTradeAddendumOutright548) LocationSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeAddendumOutright548) LocationInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.LocationSinceVersion()
}

func (*ExecutionReportTradeAddendumOutright548) LocationDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeAddendumOutright548) LocationMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeAddendumOutright548) LocationMinValue() byte {
	return byte(32)
}

func (*ExecutionReportTradeAddendumOutright548) LocationMaxValue() byte {
	return byte(126)
}

func (*ExecutionReportTradeAddendumOutright548) LocationNullValue() byte {
	return 0
}

func (e *ExecutionReportTradeAddendumOutright548) LocationCharacterEncoding() string {
	return "US-ASCII"
}

func (*ExecutionReportTradeAddendumOutright548) SecurityIDId() uint16 {
	return 48
}

func (*ExecutionReportTradeAddendumOutright548) SecurityIDSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeAddendumOutright548) SecurityIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.SecurityIDSinceVersion()
}

func (*ExecutionReportTradeAddendumOutright548) SecurityIDDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeAddendumOutright548) SecurityIDMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeAddendumOutright548) SecurityIDMinValue() int32 {
	return math.MinInt32 + 1
}

func (*ExecutionReportTradeAddendumOutright548) SecurityIDMaxValue() int32 {
	return math.MaxInt32
}

func (*ExecutionReportTradeAddendumOutright548) SecurityIDNullValue() int32 {
	return math.MinInt32
}

func (*ExecutionReportTradeAddendumOutright548) LastQtyId() uint16 {
	return 32
}

func (*ExecutionReportTradeAddendumOutright548) LastQtySinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeAddendumOutright548) LastQtyInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.LastQtySinceVersion()
}

func (*ExecutionReportTradeAddendumOutright548) LastQtyDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeAddendumOutright548) LastQtyMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeAddendumOutright548) LastQtyMinValue() uint32 {
	return 0
}

func (*ExecutionReportTradeAddendumOutright548) LastQtyMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*ExecutionReportTradeAddendumOutright548) LastQtyNullValue() uint32 {
	return math.MaxUint32
}

func (*ExecutionReportTradeAddendumOutright548) SideTradeIDId() uint16 {
	return 1506
}

func (*ExecutionReportTradeAddendumOutright548) SideTradeIDSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeAddendumOutright548) SideTradeIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.SideTradeIDSinceVersion()
}

func (*ExecutionReportTradeAddendumOutright548) SideTradeIDDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeAddendumOutright548) SideTradeIDMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeAddendumOutright548) SideTradeIDMinValue() uint32 {
	return 0
}

func (*ExecutionReportTradeAddendumOutright548) SideTradeIDMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*ExecutionReportTradeAddendumOutright548) SideTradeIDNullValue() uint32 {
	return math.MaxUint32
}

func (*ExecutionReportTradeAddendumOutright548) OrigSideTradeIDId() uint16 {
	return 1507
}

func (*ExecutionReportTradeAddendumOutright548) OrigSideTradeIDSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeAddendumOutright548) OrigSideTradeIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.OrigSideTradeIDSinceVersion()
}

func (*ExecutionReportTradeAddendumOutright548) OrigSideTradeIDDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeAddendumOutright548) OrigSideTradeIDMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeAddendumOutright548) OrigSideTradeIDMinValue() uint32 {
	return 0
}

func (*ExecutionReportTradeAddendumOutright548) OrigSideTradeIDMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*ExecutionReportTradeAddendumOutright548) OrigSideTradeIDNullValue() uint32 {
	return 4294967295
}

func (*ExecutionReportTradeAddendumOutright548) TradeDateId() uint16 {
	return 75
}

func (*ExecutionReportTradeAddendumOutright548) TradeDateSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeAddendumOutright548) TradeDateInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.TradeDateSinceVersion()
}

func (*ExecutionReportTradeAddendumOutright548) TradeDateDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeAddendumOutright548) TradeDateMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeAddendumOutright548) TradeDateMinValue() uint16 {
	return 0
}

func (*ExecutionReportTradeAddendumOutright548) TradeDateMaxValue() uint16 {
	return math.MaxUint16 - 1
}

func (*ExecutionReportTradeAddendumOutright548) TradeDateNullValue() uint16 {
	return 65535
}

func (*ExecutionReportTradeAddendumOutright548) OrdStatusId() uint16 {
	return 39
}

func (*ExecutionReportTradeAddendumOutright548) OrdStatusSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeAddendumOutright548) OrdStatusInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.OrdStatusSinceVersion()
}

func (*ExecutionReportTradeAddendumOutright548) OrdStatusDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeAddendumOutright548) OrdStatusMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeAddendumOutright548) ExecTypeId() uint16 {
	return 150
}

func (*ExecutionReportTradeAddendumOutright548) ExecTypeSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeAddendumOutright548) ExecTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.ExecTypeSinceVersion()
}

func (*ExecutionReportTradeAddendumOutright548) ExecTypeDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeAddendumOutright548) ExecTypeMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeAddendumOutright548) SideId() uint16 {
	return 54
}

func (*ExecutionReportTradeAddendumOutright548) SideSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeAddendumOutright548) SideInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.SideSinceVersion()
}

func (*ExecutionReportTradeAddendumOutright548) SideDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeAddendumOutright548) SideMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeAddendumOutright548) ManualOrderIndicatorId() uint16 {
	return 1028
}

func (*ExecutionReportTradeAddendumOutright548) ManualOrderIndicatorSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeAddendumOutright548) ManualOrderIndicatorInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.ManualOrderIndicatorSinceVersion()
}

func (*ExecutionReportTradeAddendumOutright548) ManualOrderIndicatorDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeAddendumOutright548) ManualOrderIndicatorMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeAddendumOutright548) PossRetransFlagId() uint16 {
	return 9765
}

func (*ExecutionReportTradeAddendumOutright548) PossRetransFlagSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeAddendumOutright548) PossRetransFlagInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.PossRetransFlagSinceVersion()
}

func (*ExecutionReportTradeAddendumOutright548) PossRetransFlagDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeAddendumOutright548) PossRetransFlagMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeAddendumOutright548) ExecInstId() uint16 {
	return 18
}

func (*ExecutionReportTradeAddendumOutright548) ExecInstSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeAddendumOutright548) ExecInstInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.ExecInstSinceVersion()
}

func (*ExecutionReportTradeAddendumOutright548) ExecInstDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeAddendumOutright548) ExecInstMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeAddendumOutright548) ExecutionModeId() uint16 {
	return 5906
}

func (*ExecutionReportTradeAddendumOutright548) ExecutionModeSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeAddendumOutright548) ExecutionModeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.ExecutionModeSinceVersion()
}

func (*ExecutionReportTradeAddendumOutright548) ExecutionModeDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeAddendumOutright548) ExecutionModeMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeAddendumOutright548) LiquidityFlagId() uint16 {
	return 9373
}

func (*ExecutionReportTradeAddendumOutright548) LiquidityFlagSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeAddendumOutright548) LiquidityFlagInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.LiquidityFlagSinceVersion()
}

func (*ExecutionReportTradeAddendumOutright548) LiquidityFlagDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeAddendumOutright548) LiquidityFlagMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeAddendumOutright548) ManagedOrderId() uint16 {
	return 6881
}

func (*ExecutionReportTradeAddendumOutright548) ManagedOrderSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeAddendumOutright548) ManagedOrderInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.ManagedOrderSinceVersion()
}

func (*ExecutionReportTradeAddendumOutright548) ManagedOrderDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeAddendumOutright548) ManagedOrderMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeAddendumOutright548) ShortSaleTypeId() uint16 {
	return 5409
}

func (*ExecutionReportTradeAddendumOutright548) ShortSaleTypeSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeAddendumOutright548) ShortSaleTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.ShortSaleTypeSinceVersion()
}

func (*ExecutionReportTradeAddendumOutright548) ShortSaleTypeDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeAddendumOutright548) ShortSaleTypeMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeAddendumOutright548NoFills) FillPxId() uint16 {
	return 1364
}

func (*ExecutionReportTradeAddendumOutright548NoFills) FillPxSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeAddendumOutright548NoFills) FillPxInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.FillPxSinceVersion()
}

func (*ExecutionReportTradeAddendumOutright548NoFills) FillPxDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeAddendumOutright548NoFills) FillPxMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeAddendumOutright548NoFills) FillQtyId() uint16 {
	return 1365
}

func (*ExecutionReportTradeAddendumOutright548NoFills) FillQtySinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeAddendumOutright548NoFills) FillQtyInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.FillQtySinceVersion()
}

func (*ExecutionReportTradeAddendumOutright548NoFills) FillQtyDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeAddendumOutright548NoFills) FillQtyMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeAddendumOutright548NoFills) FillQtyMinValue() uint32 {
	return 0
}

func (*ExecutionReportTradeAddendumOutright548NoFills) FillQtyMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*ExecutionReportTradeAddendumOutright548NoFills) FillQtyNullValue() uint32 {
	return math.MaxUint32
}

func (*ExecutionReportTradeAddendumOutright548NoFills) FillExecIDId() uint16 {
	return 1363
}

func (*ExecutionReportTradeAddendumOutright548NoFills) FillExecIDSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeAddendumOutright548NoFills) FillExecIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.FillExecIDSinceVersion()
}

func (*ExecutionReportTradeAddendumOutright548NoFills) FillExecIDDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeAddendumOutright548NoFills) FillExecIDMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeAddendumOutright548NoFills) FillExecIDMinValue() byte {
	return byte(32)
}

func (*ExecutionReportTradeAddendumOutright548NoFills) FillExecIDMaxValue() byte {
	return byte(126)
}

func (*ExecutionReportTradeAddendumOutright548NoFills) FillExecIDNullValue() byte {
	return 0
}

func (e *ExecutionReportTradeAddendumOutright548NoFills) FillExecIDCharacterEncoding() string {
	return "US-ASCII"
}

func (*ExecutionReportTradeAddendumOutright548NoFills) FillYieldTypeId() uint16 {
	return 1622
}

func (*ExecutionReportTradeAddendumOutright548NoFills) FillYieldTypeSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeAddendumOutright548NoFills) FillYieldTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.FillYieldTypeSinceVersion()
}

func (*ExecutionReportTradeAddendumOutright548NoFills) FillYieldTypeDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeAddendumOutright548NoFills) FillYieldTypeMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeAddendumOutright548NoFills) FillYieldTypeMinValue() uint8 {
	return 0
}

func (*ExecutionReportTradeAddendumOutright548NoFills) FillYieldTypeMaxValue() uint8 {
	return math.MaxUint8 - 1
}

func (*ExecutionReportTradeAddendumOutright548NoFills) FillYieldTypeNullValue() uint8 {
	return math.MaxUint8
}

func (*ExecutionReportTradeAddendumOutright548NoOrderEvents) OrderEventPxId() uint16 {
	return 1799
}

func (*ExecutionReportTradeAddendumOutright548NoOrderEvents) OrderEventPxSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeAddendumOutright548NoOrderEvents) OrderEventPxInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.OrderEventPxSinceVersion()
}

func (*ExecutionReportTradeAddendumOutright548NoOrderEvents) OrderEventPxDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeAddendumOutright548NoOrderEvents) OrderEventPxMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeAddendumOutright548NoOrderEvents) OrderEventTextId() uint16 {
	return 1802
}

func (*ExecutionReportTradeAddendumOutright548NoOrderEvents) OrderEventTextSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeAddendumOutright548NoOrderEvents) OrderEventTextInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.OrderEventTextSinceVersion()
}

func (*ExecutionReportTradeAddendumOutright548NoOrderEvents) OrderEventTextDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeAddendumOutright548NoOrderEvents) OrderEventTextMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeAddendumOutright548NoOrderEvents) OrderEventTextMinValue() byte {
	return byte(32)
}

func (*ExecutionReportTradeAddendumOutright548NoOrderEvents) OrderEventTextMaxValue() byte {
	return byte(126)
}

func (*ExecutionReportTradeAddendumOutright548NoOrderEvents) OrderEventTextNullValue() byte {
	return 0
}

func (e *ExecutionReportTradeAddendumOutright548NoOrderEvents) OrderEventTextCharacterEncoding() string {
	return "US-ASCII"
}

func (*ExecutionReportTradeAddendumOutright548NoOrderEvents) OrderEventExecIDId() uint16 {
	return 1797
}

func (*ExecutionReportTradeAddendumOutright548NoOrderEvents) OrderEventExecIDSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeAddendumOutright548NoOrderEvents) OrderEventExecIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.OrderEventExecIDSinceVersion()
}

func (*ExecutionReportTradeAddendumOutright548NoOrderEvents) OrderEventExecIDDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeAddendumOutright548NoOrderEvents) OrderEventExecIDMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeAddendumOutright548NoOrderEvents) OrderEventExecIDMinValue() uint32 {
	return 0
}

func (*ExecutionReportTradeAddendumOutright548NoOrderEvents) OrderEventExecIDMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*ExecutionReportTradeAddendumOutright548NoOrderEvents) OrderEventExecIDNullValue() uint32 {
	return math.MaxUint32
}

func (*ExecutionReportTradeAddendumOutright548NoOrderEvents) OrderEventQtyId() uint16 {
	return 1800
}

func (*ExecutionReportTradeAddendumOutright548NoOrderEvents) OrderEventQtySinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeAddendumOutright548NoOrderEvents) OrderEventQtyInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.OrderEventQtySinceVersion()
}

func (*ExecutionReportTradeAddendumOutright548NoOrderEvents) OrderEventQtyDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeAddendumOutright548NoOrderEvents) OrderEventQtyMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeAddendumOutright548NoOrderEvents) OrderEventQtyMinValue() uint32 {
	return 0
}

func (*ExecutionReportTradeAddendumOutright548NoOrderEvents) OrderEventQtyMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*ExecutionReportTradeAddendumOutright548NoOrderEvents) OrderEventQtyNullValue() uint32 {
	return math.MaxUint32
}

func (*ExecutionReportTradeAddendumOutright548NoOrderEvents) OrderEventTypeId() uint16 {
	return 1796
}

func (*ExecutionReportTradeAddendumOutright548NoOrderEvents) OrderEventTypeSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeAddendumOutright548NoOrderEvents) OrderEventTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.OrderEventTypeSinceVersion()
}

func (*ExecutionReportTradeAddendumOutright548NoOrderEvents) OrderEventTypeDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeAddendumOutright548NoOrderEvents) OrderEventTypeMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeAddendumOutright548NoOrderEvents) OrderEventReasonId() uint16 {
	return 1798
}

func (*ExecutionReportTradeAddendumOutright548NoOrderEvents) OrderEventReasonSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeAddendumOutright548NoOrderEvents) OrderEventReasonInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.OrderEventReasonSinceVersion()
}

func (*ExecutionReportTradeAddendumOutright548NoOrderEvents) OrderEventReasonDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeAddendumOutright548NoOrderEvents) OrderEventReasonMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeAddendumOutright548NoOrderEvents) OrderEventReasonMinValue() uint8 {
	return 0
}

func (*ExecutionReportTradeAddendumOutright548NoOrderEvents) OrderEventReasonMaxValue() uint8 {
	return math.MaxUint8 - 1
}

func (*ExecutionReportTradeAddendumOutright548NoOrderEvents) OrderEventReasonNullValue() uint8 {
	return math.MaxUint8
}

func (*ExecutionReportTradeAddendumOutright548NoOrderEvents) OriginalOrderEventExecIDId() uint16 {
	return 6555
}

func (*ExecutionReportTradeAddendumOutright548NoOrderEvents) OriginalOrderEventExecIDSinceVersion() uint16 {
	return 3
}

func (e *ExecutionReportTradeAddendumOutright548NoOrderEvents) OriginalOrderEventExecIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.OriginalOrderEventExecIDSinceVersion()
}

func (*ExecutionReportTradeAddendumOutright548NoOrderEvents) OriginalOrderEventExecIDDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeAddendumOutright548NoOrderEvents) OriginalOrderEventExecIDMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeAddendumOutright548NoOrderEvents) OriginalOrderEventExecIDMinValue() uint32 {
	return 0
}

func (*ExecutionReportTradeAddendumOutright548NoOrderEvents) OriginalOrderEventExecIDMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*ExecutionReportTradeAddendumOutright548NoOrderEvents) OriginalOrderEventExecIDNullValue() uint32 {
	return 4294967295
}

func (*ExecutionReportTradeAddendumOutright548) NoFillsId() uint16 {
	return 1362
}

func (*ExecutionReportTradeAddendumOutright548) NoFillsSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeAddendumOutright548) NoFillsInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.NoFillsSinceVersion()
}

func (*ExecutionReportTradeAddendumOutright548) NoFillsDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeAddendumOutright548NoFills) SbeBlockLength() (blockLength uint) {
	return 15
}

func (*ExecutionReportTradeAddendumOutright548NoFills) SbeSchemaVersion() (schemaVersion uint16) {
	return 5
}

func (*ExecutionReportTradeAddendumOutright548) NoOrderEventsId() uint16 {
	return 1795
}

func (*ExecutionReportTradeAddendumOutright548) NoOrderEventsSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeAddendumOutright548) NoOrderEventsInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.NoOrderEventsSinceVersion()
}

func (*ExecutionReportTradeAddendumOutright548) NoOrderEventsDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeAddendumOutright548NoOrderEvents) SbeBlockLength() (blockLength uint) {
	return 27
}

func (*ExecutionReportTradeAddendumOutright548NoOrderEvents) SbeSchemaVersion() (schemaVersion uint16) {
	return 5
}
