// Generated SBE (Simple Binary Encoding) message codec

package ilinkbinary

import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
)

type ExecutionReportTradeAddendumSpreadLeg550 struct {
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
	ManualOrderIndicator     ManualOrdIndReqEnum
	PossRetransFlag          BooleanFlagEnum
	Side                     SideReqEnum
	NoFills                  []ExecutionReportTradeAddendumSpreadLeg550NoFills
	NoOrderEvents            []ExecutionReportTradeAddendumSpreadLeg550NoOrderEvents
}
type ExecutionReportTradeAddendumSpreadLeg550NoFills struct {
	FillPx        PRICE9
	FillQty       uint32
	FillExecID    [2]byte
	FillYieldType uint8
}
type ExecutionReportTradeAddendumSpreadLeg550NoOrderEvents struct {
	OrderEventPx             PRICE9
	OrderEventText           [5]byte
	OrderEventExecID         uint32
	OrderEventQty            uint32
	OrderEventType           TradeAddendumEnum
	OrderEventReason         uint8
	OriginalOrderEventExecID uint32
}

func (e *ExecutionReportTradeAddendumSpreadLeg550) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
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
	if err := e.ManualOrderIndicator.Encode(_m, _w); err != nil {
		return err
	}
	if err := e.PossRetransFlag.Encode(_m, _w); err != nil {
		return err
	}
	if err := e.Side.Encode(_m, _w); err != nil {
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

func (e *ExecutionReportTradeAddendumSpreadLeg550) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
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
	if e.SideInActingVersion(actingVersion) {
		if err := e.Side.Decode(_m, _r, actingVersion); err != nil {
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
			e.NoFills = make([]ExecutionReportTradeAddendumSpreadLeg550NoFills, NoFillsNumInGroup)
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
			e.NoOrderEvents = make([]ExecutionReportTradeAddendumSpreadLeg550NoOrderEvents, NoOrderEventsNumInGroup)
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

func (e *ExecutionReportTradeAddendumSpreadLeg550) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
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
	if err := e.ManualOrderIndicator.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if err := e.PossRetransFlag.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if err := e.Side.RangeCheck(actingVersion, schemaVersion); err != nil {
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

func ExecutionReportTradeAddendumSpreadLeg550Init(e *ExecutionReportTradeAddendumSpreadLeg550) {
	e.OrigSecondaryExecutionID = 18446744073709551615
	e.OrigSideTradeID = 4294967295
	e.TradeDate = 65535
	return
}

func (e *ExecutionReportTradeAddendumSpreadLeg550NoFills) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
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

func (e *ExecutionReportTradeAddendumSpreadLeg550NoFills) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint) error {
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

func (e *ExecutionReportTradeAddendumSpreadLeg550NoFills) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
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

func ExecutionReportTradeAddendumSpreadLeg550NoFillsInit(e *ExecutionReportTradeAddendumSpreadLeg550NoFills) {
	return
}

func (e *ExecutionReportTradeAddendumSpreadLeg550NoOrderEvents) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
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

func (e *ExecutionReportTradeAddendumSpreadLeg550NoOrderEvents) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint) error {
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

func (e *ExecutionReportTradeAddendumSpreadLeg550NoOrderEvents) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
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

func ExecutionReportTradeAddendumSpreadLeg550NoOrderEventsInit(e *ExecutionReportTradeAddendumSpreadLeg550NoOrderEvents) {
	for idx := 0; idx < 5; idx++ {
		e.OrderEventText[idx] = 0
	}
	e.OriginalOrderEventExecID = 4294967295
	return
}

func (*ExecutionReportTradeAddendumSpreadLeg550) SbeBlockLength() (blockLength uint16) {
	return 176
}

func (*ExecutionReportTradeAddendumSpreadLeg550) SbeTemplateId() (templateId uint16) {
	return 550
}

func (*ExecutionReportTradeAddendumSpreadLeg550) SbeSchemaId() (schemaId uint16) {
	return 8
}

func (*ExecutionReportTradeAddendumSpreadLeg550) SbeSchemaVersion() (schemaVersion uint16) {
	return 5
}

func (*ExecutionReportTradeAddendumSpreadLeg550) SbeSemanticType() (semanticType []byte) {
	return []byte("8")
}

func (*ExecutionReportTradeAddendumSpreadLeg550) SeqNumId() uint16 {
	return 9726
}

func (*ExecutionReportTradeAddendumSpreadLeg550) SeqNumSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeAddendumSpreadLeg550) SeqNumInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.SeqNumSinceVersion()
}

func (*ExecutionReportTradeAddendumSpreadLeg550) SeqNumDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeAddendumSpreadLeg550) SeqNumMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeAddendumSpreadLeg550) SeqNumMinValue() uint32 {
	return 0
}

func (*ExecutionReportTradeAddendumSpreadLeg550) SeqNumMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*ExecutionReportTradeAddendumSpreadLeg550) SeqNumNullValue() uint32 {
	return math.MaxUint32
}

func (*ExecutionReportTradeAddendumSpreadLeg550) UUIDId() uint16 {
	return 39001
}

func (*ExecutionReportTradeAddendumSpreadLeg550) UUIDSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeAddendumSpreadLeg550) UUIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.UUIDSinceVersion()
}

func (*ExecutionReportTradeAddendumSpreadLeg550) UUIDDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeAddendumSpreadLeg550) UUIDMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeAddendumSpreadLeg550) UUIDMinValue() uint64 {
	return 0
}

func (*ExecutionReportTradeAddendumSpreadLeg550) UUIDMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*ExecutionReportTradeAddendumSpreadLeg550) UUIDNullValue() uint64 {
	return math.MaxUint64
}

func (*ExecutionReportTradeAddendumSpreadLeg550) ExecIDId() uint16 {
	return 17
}

func (*ExecutionReportTradeAddendumSpreadLeg550) ExecIDSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeAddendumSpreadLeg550) ExecIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.ExecIDSinceVersion()
}

func (*ExecutionReportTradeAddendumSpreadLeg550) ExecIDDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeAddendumSpreadLeg550) ExecIDMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeAddendumSpreadLeg550) ExecIDMinValue() byte {
	return byte(32)
}

func (*ExecutionReportTradeAddendumSpreadLeg550) ExecIDMaxValue() byte {
	return byte(126)
}

func (*ExecutionReportTradeAddendumSpreadLeg550) ExecIDNullValue() byte {
	return 0
}

func (e *ExecutionReportTradeAddendumSpreadLeg550) ExecIDCharacterEncoding() string {
	return "US-ASCII"
}

func (*ExecutionReportTradeAddendumSpreadLeg550) SenderIDId() uint16 {
	return 5392
}

func (*ExecutionReportTradeAddendumSpreadLeg550) SenderIDSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeAddendumSpreadLeg550) SenderIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.SenderIDSinceVersion()
}

func (*ExecutionReportTradeAddendumSpreadLeg550) SenderIDDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeAddendumSpreadLeg550) SenderIDMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeAddendumSpreadLeg550) SenderIDMinValue() byte {
	return byte(32)
}

func (*ExecutionReportTradeAddendumSpreadLeg550) SenderIDMaxValue() byte {
	return byte(126)
}

func (*ExecutionReportTradeAddendumSpreadLeg550) SenderIDNullValue() byte {
	return 0
}

func (e *ExecutionReportTradeAddendumSpreadLeg550) SenderIDCharacterEncoding() string {
	return "US-ASCII"
}

func (*ExecutionReportTradeAddendumSpreadLeg550) ClOrdIDId() uint16 {
	return 11
}

func (*ExecutionReportTradeAddendumSpreadLeg550) ClOrdIDSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeAddendumSpreadLeg550) ClOrdIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.ClOrdIDSinceVersion()
}

func (*ExecutionReportTradeAddendumSpreadLeg550) ClOrdIDDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeAddendumSpreadLeg550) ClOrdIDMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeAddendumSpreadLeg550) ClOrdIDMinValue() byte {
	return byte(32)
}

func (*ExecutionReportTradeAddendumSpreadLeg550) ClOrdIDMaxValue() byte {
	return byte(126)
}

func (*ExecutionReportTradeAddendumSpreadLeg550) ClOrdIDNullValue() byte {
	return 0
}

func (e *ExecutionReportTradeAddendumSpreadLeg550) ClOrdIDCharacterEncoding() string {
	return "US-ASCII"
}

func (*ExecutionReportTradeAddendumSpreadLeg550) PartyDetailsListReqIDId() uint16 {
	return 1505
}

func (*ExecutionReportTradeAddendumSpreadLeg550) PartyDetailsListReqIDSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeAddendumSpreadLeg550) PartyDetailsListReqIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.PartyDetailsListReqIDSinceVersion()
}

func (*ExecutionReportTradeAddendumSpreadLeg550) PartyDetailsListReqIDDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeAddendumSpreadLeg550) PartyDetailsListReqIDMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeAddendumSpreadLeg550) PartyDetailsListReqIDMinValue() uint64 {
	return 0
}

func (*ExecutionReportTradeAddendumSpreadLeg550) PartyDetailsListReqIDMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*ExecutionReportTradeAddendumSpreadLeg550) PartyDetailsListReqIDNullValue() uint64 {
	return math.MaxUint64
}

func (*ExecutionReportTradeAddendumSpreadLeg550) LastPxId() uint16 {
	return 31
}

func (*ExecutionReportTradeAddendumSpreadLeg550) LastPxSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeAddendumSpreadLeg550) LastPxInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.LastPxSinceVersion()
}

func (*ExecutionReportTradeAddendumSpreadLeg550) LastPxDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeAddendumSpreadLeg550) LastPxMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeAddendumSpreadLeg550) OrderIDId() uint16 {
	return 37
}

func (*ExecutionReportTradeAddendumSpreadLeg550) OrderIDSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeAddendumSpreadLeg550) OrderIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.OrderIDSinceVersion()
}

func (*ExecutionReportTradeAddendumSpreadLeg550) OrderIDDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeAddendumSpreadLeg550) OrderIDMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeAddendumSpreadLeg550) OrderIDMinValue() uint64 {
	return 0
}

func (*ExecutionReportTradeAddendumSpreadLeg550) OrderIDMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*ExecutionReportTradeAddendumSpreadLeg550) OrderIDNullValue() uint64 {
	return math.MaxUint64
}

func (*ExecutionReportTradeAddendumSpreadLeg550) TransactTimeId() uint16 {
	return 60
}

func (*ExecutionReportTradeAddendumSpreadLeg550) TransactTimeSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeAddendumSpreadLeg550) TransactTimeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.TransactTimeSinceVersion()
}

func (*ExecutionReportTradeAddendumSpreadLeg550) TransactTimeDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeAddendumSpreadLeg550) TransactTimeMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeAddendumSpreadLeg550) TransactTimeMinValue() uint64 {
	return 0
}

func (*ExecutionReportTradeAddendumSpreadLeg550) TransactTimeMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*ExecutionReportTradeAddendumSpreadLeg550) TransactTimeNullValue() uint64 {
	return math.MaxUint64
}

func (*ExecutionReportTradeAddendumSpreadLeg550) SendingTimeEpochId() uint16 {
	return 5297
}

func (*ExecutionReportTradeAddendumSpreadLeg550) SendingTimeEpochSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeAddendumSpreadLeg550) SendingTimeEpochInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.SendingTimeEpochSinceVersion()
}

func (*ExecutionReportTradeAddendumSpreadLeg550) SendingTimeEpochDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeAddendumSpreadLeg550) SendingTimeEpochMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeAddendumSpreadLeg550) SendingTimeEpochMinValue() uint64 {
	return 0
}

func (*ExecutionReportTradeAddendumSpreadLeg550) SendingTimeEpochMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*ExecutionReportTradeAddendumSpreadLeg550) SendingTimeEpochNullValue() uint64 {
	return math.MaxUint64
}

func (*ExecutionReportTradeAddendumSpreadLeg550) SecExecIDId() uint16 {
	return 527
}

func (*ExecutionReportTradeAddendumSpreadLeg550) SecExecIDSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeAddendumSpreadLeg550) SecExecIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.SecExecIDSinceVersion()
}

func (*ExecutionReportTradeAddendumSpreadLeg550) SecExecIDDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeAddendumSpreadLeg550) SecExecIDMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeAddendumSpreadLeg550) SecExecIDMinValue() uint64 {
	return 0
}

func (*ExecutionReportTradeAddendumSpreadLeg550) SecExecIDMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*ExecutionReportTradeAddendumSpreadLeg550) SecExecIDNullValue() uint64 {
	return math.MaxUint64
}

func (*ExecutionReportTradeAddendumSpreadLeg550) OrigSecondaryExecutionIDId() uint16 {
	return 9703
}

func (*ExecutionReportTradeAddendumSpreadLeg550) OrigSecondaryExecutionIDSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeAddendumSpreadLeg550) OrigSecondaryExecutionIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.OrigSecondaryExecutionIDSinceVersion()
}

func (*ExecutionReportTradeAddendumSpreadLeg550) OrigSecondaryExecutionIDDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeAddendumSpreadLeg550) OrigSecondaryExecutionIDMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeAddendumSpreadLeg550) OrigSecondaryExecutionIDMinValue() uint64 {
	return 0
}

func (*ExecutionReportTradeAddendumSpreadLeg550) OrigSecondaryExecutionIDMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*ExecutionReportTradeAddendumSpreadLeg550) OrigSecondaryExecutionIDNullValue() uint64 {
	return 18446744073709551615
}

func (*ExecutionReportTradeAddendumSpreadLeg550) LocationId() uint16 {
	return 9537
}

func (*ExecutionReportTradeAddendumSpreadLeg550) LocationSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeAddendumSpreadLeg550) LocationInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.LocationSinceVersion()
}

func (*ExecutionReportTradeAddendumSpreadLeg550) LocationDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeAddendumSpreadLeg550) LocationMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeAddendumSpreadLeg550) LocationMinValue() byte {
	return byte(32)
}

func (*ExecutionReportTradeAddendumSpreadLeg550) LocationMaxValue() byte {
	return byte(126)
}

func (*ExecutionReportTradeAddendumSpreadLeg550) LocationNullValue() byte {
	return 0
}

func (e *ExecutionReportTradeAddendumSpreadLeg550) LocationCharacterEncoding() string {
	return "US-ASCII"
}

func (*ExecutionReportTradeAddendumSpreadLeg550) SecurityIDId() uint16 {
	return 48
}

func (*ExecutionReportTradeAddendumSpreadLeg550) SecurityIDSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeAddendumSpreadLeg550) SecurityIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.SecurityIDSinceVersion()
}

func (*ExecutionReportTradeAddendumSpreadLeg550) SecurityIDDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeAddendumSpreadLeg550) SecurityIDMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeAddendumSpreadLeg550) SecurityIDMinValue() int32 {
	return math.MinInt32 + 1
}

func (*ExecutionReportTradeAddendumSpreadLeg550) SecurityIDMaxValue() int32 {
	return math.MaxInt32
}

func (*ExecutionReportTradeAddendumSpreadLeg550) SecurityIDNullValue() int32 {
	return math.MinInt32
}

func (*ExecutionReportTradeAddendumSpreadLeg550) LastQtyId() uint16 {
	return 32
}

func (*ExecutionReportTradeAddendumSpreadLeg550) LastQtySinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeAddendumSpreadLeg550) LastQtyInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.LastQtySinceVersion()
}

func (*ExecutionReportTradeAddendumSpreadLeg550) LastQtyDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeAddendumSpreadLeg550) LastQtyMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeAddendumSpreadLeg550) LastQtyMinValue() uint32 {
	return 0
}

func (*ExecutionReportTradeAddendumSpreadLeg550) LastQtyMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*ExecutionReportTradeAddendumSpreadLeg550) LastQtyNullValue() uint32 {
	return math.MaxUint32
}

func (*ExecutionReportTradeAddendumSpreadLeg550) SideTradeIDId() uint16 {
	return 1506
}

func (*ExecutionReportTradeAddendumSpreadLeg550) SideTradeIDSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeAddendumSpreadLeg550) SideTradeIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.SideTradeIDSinceVersion()
}

func (*ExecutionReportTradeAddendumSpreadLeg550) SideTradeIDDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeAddendumSpreadLeg550) SideTradeIDMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeAddendumSpreadLeg550) SideTradeIDMinValue() uint32 {
	return 0
}

func (*ExecutionReportTradeAddendumSpreadLeg550) SideTradeIDMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*ExecutionReportTradeAddendumSpreadLeg550) SideTradeIDNullValue() uint32 {
	return math.MaxUint32
}

func (*ExecutionReportTradeAddendumSpreadLeg550) OrigSideTradeIDId() uint16 {
	return 1507
}

func (*ExecutionReportTradeAddendumSpreadLeg550) OrigSideTradeIDSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeAddendumSpreadLeg550) OrigSideTradeIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.OrigSideTradeIDSinceVersion()
}

func (*ExecutionReportTradeAddendumSpreadLeg550) OrigSideTradeIDDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeAddendumSpreadLeg550) OrigSideTradeIDMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeAddendumSpreadLeg550) OrigSideTradeIDMinValue() uint32 {
	return 0
}

func (*ExecutionReportTradeAddendumSpreadLeg550) OrigSideTradeIDMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*ExecutionReportTradeAddendumSpreadLeg550) OrigSideTradeIDNullValue() uint32 {
	return 4294967295
}

func (*ExecutionReportTradeAddendumSpreadLeg550) TradeDateId() uint16 {
	return 75
}

func (*ExecutionReportTradeAddendumSpreadLeg550) TradeDateSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeAddendumSpreadLeg550) TradeDateInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.TradeDateSinceVersion()
}

func (*ExecutionReportTradeAddendumSpreadLeg550) TradeDateDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeAddendumSpreadLeg550) TradeDateMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeAddendumSpreadLeg550) TradeDateMinValue() uint16 {
	return 0
}

func (*ExecutionReportTradeAddendumSpreadLeg550) TradeDateMaxValue() uint16 {
	return math.MaxUint16 - 1
}

func (*ExecutionReportTradeAddendumSpreadLeg550) TradeDateNullValue() uint16 {
	return 65535
}

func (*ExecutionReportTradeAddendumSpreadLeg550) OrdStatusId() uint16 {
	return 39
}

func (*ExecutionReportTradeAddendumSpreadLeg550) OrdStatusSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeAddendumSpreadLeg550) OrdStatusInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.OrdStatusSinceVersion()
}

func (*ExecutionReportTradeAddendumSpreadLeg550) OrdStatusDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeAddendumSpreadLeg550) OrdStatusMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeAddendumSpreadLeg550) ExecTypeId() uint16 {
	return 150
}

func (*ExecutionReportTradeAddendumSpreadLeg550) ExecTypeSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeAddendumSpreadLeg550) ExecTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.ExecTypeSinceVersion()
}

func (*ExecutionReportTradeAddendumSpreadLeg550) ExecTypeDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeAddendumSpreadLeg550) ExecTypeMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeAddendumSpreadLeg550) ManualOrderIndicatorId() uint16 {
	return 1028
}

func (*ExecutionReportTradeAddendumSpreadLeg550) ManualOrderIndicatorSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeAddendumSpreadLeg550) ManualOrderIndicatorInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.ManualOrderIndicatorSinceVersion()
}

func (*ExecutionReportTradeAddendumSpreadLeg550) ManualOrderIndicatorDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeAddendumSpreadLeg550) ManualOrderIndicatorMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeAddendumSpreadLeg550) PossRetransFlagId() uint16 {
	return 9765
}

func (*ExecutionReportTradeAddendumSpreadLeg550) PossRetransFlagSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeAddendumSpreadLeg550) PossRetransFlagInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.PossRetransFlagSinceVersion()
}

func (*ExecutionReportTradeAddendumSpreadLeg550) PossRetransFlagDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeAddendumSpreadLeg550) PossRetransFlagMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeAddendumSpreadLeg550) SideId() uint16 {
	return 54
}

func (*ExecutionReportTradeAddendumSpreadLeg550) SideSinceVersion() uint16 {
	return 5
}

func (e *ExecutionReportTradeAddendumSpreadLeg550) SideInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.SideSinceVersion()
}

func (*ExecutionReportTradeAddendumSpreadLeg550) SideDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeAddendumSpreadLeg550) SideMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeAddendumSpreadLeg550NoFills) FillPxId() uint16 {
	return 1364
}

func (*ExecutionReportTradeAddendumSpreadLeg550NoFills) FillPxSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeAddendumSpreadLeg550NoFills) FillPxInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.FillPxSinceVersion()
}

func (*ExecutionReportTradeAddendumSpreadLeg550NoFills) FillPxDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeAddendumSpreadLeg550NoFills) FillPxMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeAddendumSpreadLeg550NoFills) FillQtyId() uint16 {
	return 1365
}

func (*ExecutionReportTradeAddendumSpreadLeg550NoFills) FillQtySinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeAddendumSpreadLeg550NoFills) FillQtyInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.FillQtySinceVersion()
}

func (*ExecutionReportTradeAddendumSpreadLeg550NoFills) FillQtyDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeAddendumSpreadLeg550NoFills) FillQtyMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeAddendumSpreadLeg550NoFills) FillQtyMinValue() uint32 {
	return 0
}

func (*ExecutionReportTradeAddendumSpreadLeg550NoFills) FillQtyMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*ExecutionReportTradeAddendumSpreadLeg550NoFills) FillQtyNullValue() uint32 {
	return math.MaxUint32
}

func (*ExecutionReportTradeAddendumSpreadLeg550NoFills) FillExecIDId() uint16 {
	return 1363
}

func (*ExecutionReportTradeAddendumSpreadLeg550NoFills) FillExecIDSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeAddendumSpreadLeg550NoFills) FillExecIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.FillExecIDSinceVersion()
}

func (*ExecutionReportTradeAddendumSpreadLeg550NoFills) FillExecIDDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeAddendumSpreadLeg550NoFills) FillExecIDMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeAddendumSpreadLeg550NoFills) FillExecIDMinValue() byte {
	return byte(32)
}

func (*ExecutionReportTradeAddendumSpreadLeg550NoFills) FillExecIDMaxValue() byte {
	return byte(126)
}

func (*ExecutionReportTradeAddendumSpreadLeg550NoFills) FillExecIDNullValue() byte {
	return 0
}

func (e *ExecutionReportTradeAddendumSpreadLeg550NoFills) FillExecIDCharacterEncoding() string {
	return "US-ASCII"
}

func (*ExecutionReportTradeAddendumSpreadLeg550NoFills) FillYieldTypeId() uint16 {
	return 1622
}

func (*ExecutionReportTradeAddendumSpreadLeg550NoFills) FillYieldTypeSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeAddendumSpreadLeg550NoFills) FillYieldTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.FillYieldTypeSinceVersion()
}

func (*ExecutionReportTradeAddendumSpreadLeg550NoFills) FillYieldTypeDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeAddendumSpreadLeg550NoFills) FillYieldTypeMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeAddendumSpreadLeg550NoFills) FillYieldTypeMinValue() uint8 {
	return 0
}

func (*ExecutionReportTradeAddendumSpreadLeg550NoFills) FillYieldTypeMaxValue() uint8 {
	return math.MaxUint8 - 1
}

func (*ExecutionReportTradeAddendumSpreadLeg550NoFills) FillYieldTypeNullValue() uint8 {
	return math.MaxUint8
}

func (*ExecutionReportTradeAddendumSpreadLeg550NoOrderEvents) OrderEventPxId() uint16 {
	return 1799
}

func (*ExecutionReportTradeAddendumSpreadLeg550NoOrderEvents) OrderEventPxSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeAddendumSpreadLeg550NoOrderEvents) OrderEventPxInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.OrderEventPxSinceVersion()
}

func (*ExecutionReportTradeAddendumSpreadLeg550NoOrderEvents) OrderEventPxDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeAddendumSpreadLeg550NoOrderEvents) OrderEventPxMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeAddendumSpreadLeg550NoOrderEvents) OrderEventTextId() uint16 {
	return 1802
}

func (*ExecutionReportTradeAddendumSpreadLeg550NoOrderEvents) OrderEventTextSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeAddendumSpreadLeg550NoOrderEvents) OrderEventTextInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.OrderEventTextSinceVersion()
}

func (*ExecutionReportTradeAddendumSpreadLeg550NoOrderEvents) OrderEventTextDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeAddendumSpreadLeg550NoOrderEvents) OrderEventTextMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeAddendumSpreadLeg550NoOrderEvents) OrderEventTextMinValue() byte {
	return byte(32)
}

func (*ExecutionReportTradeAddendumSpreadLeg550NoOrderEvents) OrderEventTextMaxValue() byte {
	return byte(126)
}

func (*ExecutionReportTradeAddendumSpreadLeg550NoOrderEvents) OrderEventTextNullValue() byte {
	return 0
}

func (e *ExecutionReportTradeAddendumSpreadLeg550NoOrderEvents) OrderEventTextCharacterEncoding() string {
	return "US-ASCII"
}

func (*ExecutionReportTradeAddendumSpreadLeg550NoOrderEvents) OrderEventExecIDId() uint16 {
	return 1797
}

func (*ExecutionReportTradeAddendumSpreadLeg550NoOrderEvents) OrderEventExecIDSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeAddendumSpreadLeg550NoOrderEvents) OrderEventExecIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.OrderEventExecIDSinceVersion()
}

func (*ExecutionReportTradeAddendumSpreadLeg550NoOrderEvents) OrderEventExecIDDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeAddendumSpreadLeg550NoOrderEvents) OrderEventExecIDMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeAddendumSpreadLeg550NoOrderEvents) OrderEventExecIDMinValue() uint32 {
	return 0
}

func (*ExecutionReportTradeAddendumSpreadLeg550NoOrderEvents) OrderEventExecIDMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*ExecutionReportTradeAddendumSpreadLeg550NoOrderEvents) OrderEventExecIDNullValue() uint32 {
	return math.MaxUint32
}

func (*ExecutionReportTradeAddendumSpreadLeg550NoOrderEvents) OrderEventQtyId() uint16 {
	return 1800
}

func (*ExecutionReportTradeAddendumSpreadLeg550NoOrderEvents) OrderEventQtySinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeAddendumSpreadLeg550NoOrderEvents) OrderEventQtyInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.OrderEventQtySinceVersion()
}

func (*ExecutionReportTradeAddendumSpreadLeg550NoOrderEvents) OrderEventQtyDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeAddendumSpreadLeg550NoOrderEvents) OrderEventQtyMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeAddendumSpreadLeg550NoOrderEvents) OrderEventQtyMinValue() uint32 {
	return 0
}

func (*ExecutionReportTradeAddendumSpreadLeg550NoOrderEvents) OrderEventQtyMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*ExecutionReportTradeAddendumSpreadLeg550NoOrderEvents) OrderEventQtyNullValue() uint32 {
	return math.MaxUint32
}

func (*ExecutionReportTradeAddendumSpreadLeg550NoOrderEvents) OrderEventTypeId() uint16 {
	return 1796
}

func (*ExecutionReportTradeAddendumSpreadLeg550NoOrderEvents) OrderEventTypeSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeAddendumSpreadLeg550NoOrderEvents) OrderEventTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.OrderEventTypeSinceVersion()
}

func (*ExecutionReportTradeAddendumSpreadLeg550NoOrderEvents) OrderEventTypeDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeAddendumSpreadLeg550NoOrderEvents) OrderEventTypeMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeAddendumSpreadLeg550NoOrderEvents) OrderEventReasonId() uint16 {
	return 1798
}

func (*ExecutionReportTradeAddendumSpreadLeg550NoOrderEvents) OrderEventReasonSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeAddendumSpreadLeg550NoOrderEvents) OrderEventReasonInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.OrderEventReasonSinceVersion()
}

func (*ExecutionReportTradeAddendumSpreadLeg550NoOrderEvents) OrderEventReasonDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeAddendumSpreadLeg550NoOrderEvents) OrderEventReasonMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeAddendumSpreadLeg550NoOrderEvents) OrderEventReasonMinValue() uint8 {
	return 0
}

func (*ExecutionReportTradeAddendumSpreadLeg550NoOrderEvents) OrderEventReasonMaxValue() uint8 {
	return math.MaxUint8 - 1
}

func (*ExecutionReportTradeAddendumSpreadLeg550NoOrderEvents) OrderEventReasonNullValue() uint8 {
	return math.MaxUint8
}

func (*ExecutionReportTradeAddendumSpreadLeg550NoOrderEvents) OriginalOrderEventExecIDId() uint16 {
	return 6555
}

func (*ExecutionReportTradeAddendumSpreadLeg550NoOrderEvents) OriginalOrderEventExecIDSinceVersion() uint16 {
	return 3
}

func (e *ExecutionReportTradeAddendumSpreadLeg550NoOrderEvents) OriginalOrderEventExecIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.OriginalOrderEventExecIDSinceVersion()
}

func (*ExecutionReportTradeAddendumSpreadLeg550NoOrderEvents) OriginalOrderEventExecIDDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeAddendumSpreadLeg550NoOrderEvents) OriginalOrderEventExecIDMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeAddendumSpreadLeg550NoOrderEvents) OriginalOrderEventExecIDMinValue() uint32 {
	return 0
}

func (*ExecutionReportTradeAddendumSpreadLeg550NoOrderEvents) OriginalOrderEventExecIDMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*ExecutionReportTradeAddendumSpreadLeg550NoOrderEvents) OriginalOrderEventExecIDNullValue() uint32 {
	return 4294967295
}

func (*ExecutionReportTradeAddendumSpreadLeg550) NoFillsId() uint16 {
	return 1362
}

func (*ExecutionReportTradeAddendumSpreadLeg550) NoFillsSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeAddendumSpreadLeg550) NoFillsInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.NoFillsSinceVersion()
}

func (*ExecutionReportTradeAddendumSpreadLeg550) NoFillsDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeAddendumSpreadLeg550NoFills) SbeBlockLength() (blockLength uint) {
	return 15
}

func (*ExecutionReportTradeAddendumSpreadLeg550NoFills) SbeSchemaVersion() (schemaVersion uint16) {
	return 5
}

func (*ExecutionReportTradeAddendumSpreadLeg550) NoOrderEventsId() uint16 {
	return 1795
}

func (*ExecutionReportTradeAddendumSpreadLeg550) NoOrderEventsSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeAddendumSpreadLeg550) NoOrderEventsInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.NoOrderEventsSinceVersion()
}

func (*ExecutionReportTradeAddendumSpreadLeg550) NoOrderEventsDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeAddendumSpreadLeg550NoOrderEvents) SbeBlockLength() (blockLength uint) {
	return 27
}

func (*ExecutionReportTradeAddendumSpreadLeg550NoOrderEvents) SbeSchemaVersion() (schemaVersion uint16) {
	return 5
}
