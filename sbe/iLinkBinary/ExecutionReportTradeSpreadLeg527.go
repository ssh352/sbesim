// Generated SBE (Simple Binary Encoding) message codec

package ilinkbinary

import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
)

type ExecutionReportTradeSpreadLeg527 struct {
	SeqNum                uint32
	UUID                  uint64
	ExecID                [40]byte
	SenderID              [20]byte
	ClOrdID               [20]byte
	Volatility            Decimal64NULL
	PartyDetailsListReqID uint64
	LastPx                PRICE9
	OrderID               uint64
	UnderlyingPx          PRICENULL9
	TransactTime          uint64
	SendingTimeEpoch      uint64
	SecExecID             uint64
	Location              [5]byte
	OptionDelta           Decimal32NULL
	TimeToExpiration      Decimal32NULL
	RiskFreeRate          Decimal32NULL
	SecurityID            int32
	LastQty               uint32
	CumQty                uint32
	SideTradeID           uint32
	TradeDate             uint16
	OrdStatus             OrdStatusTrdEnum
	ExecType              [1]byte
	OrdType               OrderTypeEnum
	Side                  SideReqEnum
	PossRetransFlag       BooleanFlagEnum
	NoFills               []ExecutionReportTradeSpreadLeg527NoFills
	NoOrderEvents         []ExecutionReportTradeSpreadLeg527NoOrderEvents
}
type ExecutionReportTradeSpreadLeg527NoFills struct {
	FillPx        PRICE9
	FillQty       uint32
	FillExecID    [2]byte
	FillYieldType uint8
}
type ExecutionReportTradeSpreadLeg527NoOrderEvents struct {
	OrderEventPx     PRICE9
	OrderEventText   [5]byte
	OrderEventExecID uint32
	OrderEventQty    uint32
	OrderEventType   OrderEventTypeEnum
	OrderEventReason uint8
}

func (e *ExecutionReportTradeSpreadLeg527) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
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
	if err := e.Volatility.Encode(_m, _w); err != nil {
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
	if err := e.UnderlyingPx.Encode(_m, _w); err != nil {
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
	if err := _m.WriteBytes(_w, e.Location[:]); err != nil {
		return err
	}
	if err := e.OptionDelta.Encode(_m, _w); err != nil {
		return err
	}
	if err := e.TimeToExpiration.Encode(_m, _w); err != nil {
		return err
	}
	if err := e.RiskFreeRate.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteInt32(_w, e.SecurityID); err != nil {
		return err
	}
	if err := _m.WriteUint32(_w, e.LastQty); err != nil {
		return err
	}
	if err := _m.WriteUint32(_w, e.CumQty); err != nil {
		return err
	}
	if err := _m.WriteUint32(_w, e.SideTradeID); err != nil {
		return err
	}
	if err := _m.WriteUint16(_w, e.TradeDate); err != nil {
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
	if err := e.PossRetransFlag.Encode(_m, _w); err != nil {
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

func (e *ExecutionReportTradeSpreadLeg527) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
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
	if e.VolatilityInActingVersion(actingVersion) {
		if err := e.Volatility.Decode(_m, _r, actingVersion); err != nil {
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
	if e.UnderlyingPxInActingVersion(actingVersion) {
		if err := e.UnderlyingPx.Decode(_m, _r, actingVersion); err != nil {
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
	if !e.LocationInActingVersion(actingVersion) {
		for idx := 0; idx < 5; idx++ {
			e.Location[idx] = e.LocationNullValue()
		}
	} else {
		if err := _m.ReadBytes(_r, e.Location[:]); err != nil {
			return err
		}
	}
	if e.OptionDeltaInActingVersion(actingVersion) {
		if err := e.OptionDelta.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if e.TimeToExpirationInActingVersion(actingVersion) {
		if err := e.TimeToExpiration.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if e.RiskFreeRateInActingVersion(actingVersion) {
		if err := e.RiskFreeRate.Decode(_m, _r, actingVersion); err != nil {
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
	if !e.CumQtyInActingVersion(actingVersion) {
		e.CumQty = e.CumQtyNullValue()
	} else {
		if err := _m.ReadUint32(_r, &e.CumQty); err != nil {
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
	if e.PossRetransFlagInActingVersion(actingVersion) {
		if err := e.PossRetransFlag.Decode(_m, _r, actingVersion); err != nil {
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
			e.NoFills = make([]ExecutionReportTradeSpreadLeg527NoFills, NoFillsNumInGroup)
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
			e.NoOrderEvents = make([]ExecutionReportTradeSpreadLeg527NoOrderEvents, NoOrderEventsNumInGroup)
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

func (e *ExecutionReportTradeSpreadLeg527) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
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
	if e.CumQtyInActingVersion(actingVersion) {
		if e.CumQty < e.CumQtyMinValue() || e.CumQty > e.CumQtyMaxValue() {
			return fmt.Errorf("Range check failed on e.CumQty (%v < %v > %v)", e.CumQtyMinValue(), e.CumQty, e.CumQtyMaxValue())
		}
	}
	if e.SideTradeIDInActingVersion(actingVersion) {
		if e.SideTradeID < e.SideTradeIDMinValue() || e.SideTradeID > e.SideTradeIDMaxValue() {
			return fmt.Errorf("Range check failed on e.SideTradeID (%v < %v > %v)", e.SideTradeIDMinValue(), e.SideTradeID, e.SideTradeIDMaxValue())
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
	if err := e.OrdType.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if err := e.Side.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if err := e.PossRetransFlag.RangeCheck(actingVersion, schemaVersion); err != nil {
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

func ExecutionReportTradeSpreadLeg527Init(e *ExecutionReportTradeSpreadLeg527) {
	e.TradeDate = 65535
	e.ExecType[0] = 70
	return
}

func (e *ExecutionReportTradeSpreadLeg527NoFills) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
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

func (e *ExecutionReportTradeSpreadLeg527NoFills) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint) error {
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

func (e *ExecutionReportTradeSpreadLeg527NoFills) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
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

func ExecutionReportTradeSpreadLeg527NoFillsInit(e *ExecutionReportTradeSpreadLeg527NoFills) {
	return
}

func (e *ExecutionReportTradeSpreadLeg527NoOrderEvents) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
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

func (e *ExecutionReportTradeSpreadLeg527NoOrderEvents) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint) error {
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

func (e *ExecutionReportTradeSpreadLeg527NoOrderEvents) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
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

func ExecutionReportTradeSpreadLeg527NoOrderEventsInit(e *ExecutionReportTradeSpreadLeg527NoOrderEvents) {
	for idx := 0; idx < 5; idx++ {
		e.OrderEventText[idx] = 0
	}
	return
}

func (*ExecutionReportTradeSpreadLeg527) SbeBlockLength() (blockLength uint16) {
	return 199
}

func (*ExecutionReportTradeSpreadLeg527) SbeTemplateId() (templateId uint16) {
	return 527
}

func (*ExecutionReportTradeSpreadLeg527) SbeSchemaId() (schemaId uint16) {
	return 8
}

func (*ExecutionReportTradeSpreadLeg527) SbeSchemaVersion() (schemaVersion uint16) {
	return 5
}

func (*ExecutionReportTradeSpreadLeg527) SbeSemanticType() (semanticType []byte) {
	return []byte("8")
}

func (*ExecutionReportTradeSpreadLeg527) SeqNumId() uint16 {
	return 9726
}

func (*ExecutionReportTradeSpreadLeg527) SeqNumSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeSpreadLeg527) SeqNumInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.SeqNumSinceVersion()
}

func (*ExecutionReportTradeSpreadLeg527) SeqNumDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeSpreadLeg527) SeqNumMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeSpreadLeg527) SeqNumMinValue() uint32 {
	return 0
}

func (*ExecutionReportTradeSpreadLeg527) SeqNumMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*ExecutionReportTradeSpreadLeg527) SeqNumNullValue() uint32 {
	return math.MaxUint32
}

func (*ExecutionReportTradeSpreadLeg527) UUIDId() uint16 {
	return 39001
}

func (*ExecutionReportTradeSpreadLeg527) UUIDSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeSpreadLeg527) UUIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.UUIDSinceVersion()
}

func (*ExecutionReportTradeSpreadLeg527) UUIDDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeSpreadLeg527) UUIDMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeSpreadLeg527) UUIDMinValue() uint64 {
	return 0
}

func (*ExecutionReportTradeSpreadLeg527) UUIDMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*ExecutionReportTradeSpreadLeg527) UUIDNullValue() uint64 {
	return math.MaxUint64
}

func (*ExecutionReportTradeSpreadLeg527) ExecIDId() uint16 {
	return 17
}

func (*ExecutionReportTradeSpreadLeg527) ExecIDSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeSpreadLeg527) ExecIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.ExecIDSinceVersion()
}

func (*ExecutionReportTradeSpreadLeg527) ExecIDDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeSpreadLeg527) ExecIDMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeSpreadLeg527) ExecIDMinValue() byte {
	return byte(32)
}

func (*ExecutionReportTradeSpreadLeg527) ExecIDMaxValue() byte {
	return byte(126)
}

func (*ExecutionReportTradeSpreadLeg527) ExecIDNullValue() byte {
	return 0
}

func (e *ExecutionReportTradeSpreadLeg527) ExecIDCharacterEncoding() string {
	return "US-ASCII"
}

func (*ExecutionReportTradeSpreadLeg527) SenderIDId() uint16 {
	return 5392
}

func (*ExecutionReportTradeSpreadLeg527) SenderIDSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeSpreadLeg527) SenderIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.SenderIDSinceVersion()
}

func (*ExecutionReportTradeSpreadLeg527) SenderIDDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeSpreadLeg527) SenderIDMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeSpreadLeg527) SenderIDMinValue() byte {
	return byte(32)
}

func (*ExecutionReportTradeSpreadLeg527) SenderIDMaxValue() byte {
	return byte(126)
}

func (*ExecutionReportTradeSpreadLeg527) SenderIDNullValue() byte {
	return 0
}

func (e *ExecutionReportTradeSpreadLeg527) SenderIDCharacterEncoding() string {
	return "US-ASCII"
}

func (*ExecutionReportTradeSpreadLeg527) ClOrdIDId() uint16 {
	return 11
}

func (*ExecutionReportTradeSpreadLeg527) ClOrdIDSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeSpreadLeg527) ClOrdIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.ClOrdIDSinceVersion()
}

func (*ExecutionReportTradeSpreadLeg527) ClOrdIDDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeSpreadLeg527) ClOrdIDMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeSpreadLeg527) ClOrdIDMinValue() byte {
	return byte(32)
}

func (*ExecutionReportTradeSpreadLeg527) ClOrdIDMaxValue() byte {
	return byte(126)
}

func (*ExecutionReportTradeSpreadLeg527) ClOrdIDNullValue() byte {
	return 0
}

func (e *ExecutionReportTradeSpreadLeg527) ClOrdIDCharacterEncoding() string {
	return "US-ASCII"
}

func (*ExecutionReportTradeSpreadLeg527) VolatilityId() uint16 {
	return 1188
}

func (*ExecutionReportTradeSpreadLeg527) VolatilitySinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeSpreadLeg527) VolatilityInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.VolatilitySinceVersion()
}

func (*ExecutionReportTradeSpreadLeg527) VolatilityDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeSpreadLeg527) VolatilityMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return "float"
	case 4:
		return "required"
	}
	return ""
}

func (*ExecutionReportTradeSpreadLeg527) PartyDetailsListReqIDId() uint16 {
	return 1505
}

func (*ExecutionReportTradeSpreadLeg527) PartyDetailsListReqIDSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeSpreadLeg527) PartyDetailsListReqIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.PartyDetailsListReqIDSinceVersion()
}

func (*ExecutionReportTradeSpreadLeg527) PartyDetailsListReqIDDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeSpreadLeg527) PartyDetailsListReqIDMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeSpreadLeg527) PartyDetailsListReqIDMinValue() uint64 {
	return 0
}

func (*ExecutionReportTradeSpreadLeg527) PartyDetailsListReqIDMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*ExecutionReportTradeSpreadLeg527) PartyDetailsListReqIDNullValue() uint64 {
	return math.MaxUint64
}

func (*ExecutionReportTradeSpreadLeg527) LastPxId() uint16 {
	return 31
}

func (*ExecutionReportTradeSpreadLeg527) LastPxSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeSpreadLeg527) LastPxInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.LastPxSinceVersion()
}

func (*ExecutionReportTradeSpreadLeg527) LastPxDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeSpreadLeg527) LastPxMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeSpreadLeg527) OrderIDId() uint16 {
	return 37
}

func (*ExecutionReportTradeSpreadLeg527) OrderIDSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeSpreadLeg527) OrderIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.OrderIDSinceVersion()
}

func (*ExecutionReportTradeSpreadLeg527) OrderIDDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeSpreadLeg527) OrderIDMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeSpreadLeg527) OrderIDMinValue() uint64 {
	return 0
}

func (*ExecutionReportTradeSpreadLeg527) OrderIDMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*ExecutionReportTradeSpreadLeg527) OrderIDNullValue() uint64 {
	return math.MaxUint64
}

func (*ExecutionReportTradeSpreadLeg527) UnderlyingPxId() uint16 {
	return 810
}

func (*ExecutionReportTradeSpreadLeg527) UnderlyingPxSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeSpreadLeg527) UnderlyingPxInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.UnderlyingPxSinceVersion()
}

func (*ExecutionReportTradeSpreadLeg527) UnderlyingPxDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeSpreadLeg527) UnderlyingPxMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeSpreadLeg527) TransactTimeId() uint16 {
	return 60
}

func (*ExecutionReportTradeSpreadLeg527) TransactTimeSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeSpreadLeg527) TransactTimeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.TransactTimeSinceVersion()
}

func (*ExecutionReportTradeSpreadLeg527) TransactTimeDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeSpreadLeg527) TransactTimeMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeSpreadLeg527) TransactTimeMinValue() uint64 {
	return 0
}

func (*ExecutionReportTradeSpreadLeg527) TransactTimeMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*ExecutionReportTradeSpreadLeg527) TransactTimeNullValue() uint64 {
	return math.MaxUint64
}

func (*ExecutionReportTradeSpreadLeg527) SendingTimeEpochId() uint16 {
	return 5297
}

func (*ExecutionReportTradeSpreadLeg527) SendingTimeEpochSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeSpreadLeg527) SendingTimeEpochInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.SendingTimeEpochSinceVersion()
}

func (*ExecutionReportTradeSpreadLeg527) SendingTimeEpochDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeSpreadLeg527) SendingTimeEpochMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeSpreadLeg527) SendingTimeEpochMinValue() uint64 {
	return 0
}

func (*ExecutionReportTradeSpreadLeg527) SendingTimeEpochMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*ExecutionReportTradeSpreadLeg527) SendingTimeEpochNullValue() uint64 {
	return math.MaxUint64
}

func (*ExecutionReportTradeSpreadLeg527) SecExecIDId() uint16 {
	return 527
}

func (*ExecutionReportTradeSpreadLeg527) SecExecIDSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeSpreadLeg527) SecExecIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.SecExecIDSinceVersion()
}

func (*ExecutionReportTradeSpreadLeg527) SecExecIDDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeSpreadLeg527) SecExecIDMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeSpreadLeg527) SecExecIDMinValue() uint64 {
	return 0
}

func (*ExecutionReportTradeSpreadLeg527) SecExecIDMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*ExecutionReportTradeSpreadLeg527) SecExecIDNullValue() uint64 {
	return math.MaxUint64
}

func (*ExecutionReportTradeSpreadLeg527) LocationId() uint16 {
	return 9537
}

func (*ExecutionReportTradeSpreadLeg527) LocationSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeSpreadLeg527) LocationInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.LocationSinceVersion()
}

func (*ExecutionReportTradeSpreadLeg527) LocationDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeSpreadLeg527) LocationMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeSpreadLeg527) LocationMinValue() byte {
	return byte(32)
}

func (*ExecutionReportTradeSpreadLeg527) LocationMaxValue() byte {
	return byte(126)
}

func (*ExecutionReportTradeSpreadLeg527) LocationNullValue() byte {
	return 0
}

func (e *ExecutionReportTradeSpreadLeg527) LocationCharacterEncoding() string {
	return "US-ASCII"
}

func (*ExecutionReportTradeSpreadLeg527) OptionDeltaId() uint16 {
	return 811
}

func (*ExecutionReportTradeSpreadLeg527) OptionDeltaSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeSpreadLeg527) OptionDeltaInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.OptionDeltaSinceVersion()
}

func (*ExecutionReportTradeSpreadLeg527) OptionDeltaDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeSpreadLeg527) OptionDeltaMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return "float"
	case 4:
		return "required"
	}
	return ""
}

func (*ExecutionReportTradeSpreadLeg527) TimeToExpirationId() uint16 {
	return 1189
}

func (*ExecutionReportTradeSpreadLeg527) TimeToExpirationSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeSpreadLeg527) TimeToExpirationInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.TimeToExpirationSinceVersion()
}

func (*ExecutionReportTradeSpreadLeg527) TimeToExpirationDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeSpreadLeg527) TimeToExpirationMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return "float"
	case 4:
		return "required"
	}
	return ""
}

func (*ExecutionReportTradeSpreadLeg527) RiskFreeRateId() uint16 {
	return 1190
}

func (*ExecutionReportTradeSpreadLeg527) RiskFreeRateSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeSpreadLeg527) RiskFreeRateInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.RiskFreeRateSinceVersion()
}

func (*ExecutionReportTradeSpreadLeg527) RiskFreeRateDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeSpreadLeg527) RiskFreeRateMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return "float"
	case 4:
		return "required"
	}
	return ""
}

func (*ExecutionReportTradeSpreadLeg527) SecurityIDId() uint16 {
	return 48
}

func (*ExecutionReportTradeSpreadLeg527) SecurityIDSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeSpreadLeg527) SecurityIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.SecurityIDSinceVersion()
}

func (*ExecutionReportTradeSpreadLeg527) SecurityIDDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeSpreadLeg527) SecurityIDMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeSpreadLeg527) SecurityIDMinValue() int32 {
	return math.MinInt32 + 1
}

func (*ExecutionReportTradeSpreadLeg527) SecurityIDMaxValue() int32 {
	return math.MaxInt32
}

func (*ExecutionReportTradeSpreadLeg527) SecurityIDNullValue() int32 {
	return math.MinInt32
}

func (*ExecutionReportTradeSpreadLeg527) LastQtyId() uint16 {
	return 32
}

func (*ExecutionReportTradeSpreadLeg527) LastQtySinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeSpreadLeg527) LastQtyInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.LastQtySinceVersion()
}

func (*ExecutionReportTradeSpreadLeg527) LastQtyDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeSpreadLeg527) LastQtyMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeSpreadLeg527) LastQtyMinValue() uint32 {
	return 0
}

func (*ExecutionReportTradeSpreadLeg527) LastQtyMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*ExecutionReportTradeSpreadLeg527) LastQtyNullValue() uint32 {
	return math.MaxUint32
}

func (*ExecutionReportTradeSpreadLeg527) CumQtyId() uint16 {
	return 14
}

func (*ExecutionReportTradeSpreadLeg527) CumQtySinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeSpreadLeg527) CumQtyInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.CumQtySinceVersion()
}

func (*ExecutionReportTradeSpreadLeg527) CumQtyDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeSpreadLeg527) CumQtyMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeSpreadLeg527) CumQtyMinValue() uint32 {
	return 0
}

func (*ExecutionReportTradeSpreadLeg527) CumQtyMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*ExecutionReportTradeSpreadLeg527) CumQtyNullValue() uint32 {
	return math.MaxUint32
}

func (*ExecutionReportTradeSpreadLeg527) SideTradeIDId() uint16 {
	return 1506
}

func (*ExecutionReportTradeSpreadLeg527) SideTradeIDSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeSpreadLeg527) SideTradeIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.SideTradeIDSinceVersion()
}

func (*ExecutionReportTradeSpreadLeg527) SideTradeIDDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeSpreadLeg527) SideTradeIDMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeSpreadLeg527) SideTradeIDMinValue() uint32 {
	return 0
}

func (*ExecutionReportTradeSpreadLeg527) SideTradeIDMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*ExecutionReportTradeSpreadLeg527) SideTradeIDNullValue() uint32 {
	return math.MaxUint32
}

func (*ExecutionReportTradeSpreadLeg527) TradeDateId() uint16 {
	return 75
}

func (*ExecutionReportTradeSpreadLeg527) TradeDateSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeSpreadLeg527) TradeDateInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.TradeDateSinceVersion()
}

func (*ExecutionReportTradeSpreadLeg527) TradeDateDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeSpreadLeg527) TradeDateMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeSpreadLeg527) TradeDateMinValue() uint16 {
	return 0
}

func (*ExecutionReportTradeSpreadLeg527) TradeDateMaxValue() uint16 {
	return math.MaxUint16 - 1
}

func (*ExecutionReportTradeSpreadLeg527) TradeDateNullValue() uint16 {
	return 65535
}

func (*ExecutionReportTradeSpreadLeg527) OrdStatusId() uint16 {
	return 39
}

func (*ExecutionReportTradeSpreadLeg527) OrdStatusSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeSpreadLeg527) OrdStatusInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.OrdStatusSinceVersion()
}

func (*ExecutionReportTradeSpreadLeg527) OrdStatusDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeSpreadLeg527) OrdStatusMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeSpreadLeg527) ExecTypeId() uint16 {
	return 150
}

func (*ExecutionReportTradeSpreadLeg527) ExecTypeSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeSpreadLeg527) ExecTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.ExecTypeSinceVersion()
}

func (*ExecutionReportTradeSpreadLeg527) ExecTypeDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeSpreadLeg527) ExecTypeMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeSpreadLeg527) ExecTypeMinValue() byte {
	return byte(32)
}

func (*ExecutionReportTradeSpreadLeg527) ExecTypeMaxValue() byte {
	return byte(126)
}

func (*ExecutionReportTradeSpreadLeg527) ExecTypeNullValue() byte {
	return 0
}

func (*ExecutionReportTradeSpreadLeg527) OrdTypeId() uint16 {
	return 40
}

func (*ExecutionReportTradeSpreadLeg527) OrdTypeSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeSpreadLeg527) OrdTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.OrdTypeSinceVersion()
}

func (*ExecutionReportTradeSpreadLeg527) OrdTypeDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeSpreadLeg527) OrdTypeMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeSpreadLeg527) SideId() uint16 {
	return 54
}

func (*ExecutionReportTradeSpreadLeg527) SideSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeSpreadLeg527) SideInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.SideSinceVersion()
}

func (*ExecutionReportTradeSpreadLeg527) SideDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeSpreadLeg527) SideMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeSpreadLeg527) PossRetransFlagId() uint16 {
	return 9765
}

func (*ExecutionReportTradeSpreadLeg527) PossRetransFlagSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeSpreadLeg527) PossRetransFlagInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.PossRetransFlagSinceVersion()
}

func (*ExecutionReportTradeSpreadLeg527) PossRetransFlagDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeSpreadLeg527) PossRetransFlagMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeSpreadLeg527NoFills) FillPxId() uint16 {
	return 1364
}

func (*ExecutionReportTradeSpreadLeg527NoFills) FillPxSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeSpreadLeg527NoFills) FillPxInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.FillPxSinceVersion()
}

func (*ExecutionReportTradeSpreadLeg527NoFills) FillPxDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeSpreadLeg527NoFills) FillPxMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeSpreadLeg527NoFills) FillQtyId() uint16 {
	return 1365
}

func (*ExecutionReportTradeSpreadLeg527NoFills) FillQtySinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeSpreadLeg527NoFills) FillQtyInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.FillQtySinceVersion()
}

func (*ExecutionReportTradeSpreadLeg527NoFills) FillQtyDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeSpreadLeg527NoFills) FillQtyMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeSpreadLeg527NoFills) FillQtyMinValue() uint32 {
	return 0
}

func (*ExecutionReportTradeSpreadLeg527NoFills) FillQtyMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*ExecutionReportTradeSpreadLeg527NoFills) FillQtyNullValue() uint32 {
	return math.MaxUint32
}

func (*ExecutionReportTradeSpreadLeg527NoFills) FillExecIDId() uint16 {
	return 1363
}

func (*ExecutionReportTradeSpreadLeg527NoFills) FillExecIDSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeSpreadLeg527NoFills) FillExecIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.FillExecIDSinceVersion()
}

func (*ExecutionReportTradeSpreadLeg527NoFills) FillExecIDDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeSpreadLeg527NoFills) FillExecIDMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeSpreadLeg527NoFills) FillExecIDMinValue() byte {
	return byte(32)
}

func (*ExecutionReportTradeSpreadLeg527NoFills) FillExecIDMaxValue() byte {
	return byte(126)
}

func (*ExecutionReportTradeSpreadLeg527NoFills) FillExecIDNullValue() byte {
	return 0
}

func (e *ExecutionReportTradeSpreadLeg527NoFills) FillExecIDCharacterEncoding() string {
	return "US-ASCII"
}

func (*ExecutionReportTradeSpreadLeg527NoFills) FillYieldTypeId() uint16 {
	return 1622
}

func (*ExecutionReportTradeSpreadLeg527NoFills) FillYieldTypeSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeSpreadLeg527NoFills) FillYieldTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.FillYieldTypeSinceVersion()
}

func (*ExecutionReportTradeSpreadLeg527NoFills) FillYieldTypeDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeSpreadLeg527NoFills) FillYieldTypeMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeSpreadLeg527NoFills) FillYieldTypeMinValue() uint8 {
	return 0
}

func (*ExecutionReportTradeSpreadLeg527NoFills) FillYieldTypeMaxValue() uint8 {
	return math.MaxUint8 - 1
}

func (*ExecutionReportTradeSpreadLeg527NoFills) FillYieldTypeNullValue() uint8 {
	return math.MaxUint8
}

func (*ExecutionReportTradeSpreadLeg527NoOrderEvents) OrderEventPxId() uint16 {
	return 1799
}

func (*ExecutionReportTradeSpreadLeg527NoOrderEvents) OrderEventPxSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeSpreadLeg527NoOrderEvents) OrderEventPxInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.OrderEventPxSinceVersion()
}

func (*ExecutionReportTradeSpreadLeg527NoOrderEvents) OrderEventPxDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeSpreadLeg527NoOrderEvents) OrderEventPxMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeSpreadLeg527NoOrderEvents) OrderEventTextId() uint16 {
	return 1802
}

func (*ExecutionReportTradeSpreadLeg527NoOrderEvents) OrderEventTextSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeSpreadLeg527NoOrderEvents) OrderEventTextInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.OrderEventTextSinceVersion()
}

func (*ExecutionReportTradeSpreadLeg527NoOrderEvents) OrderEventTextDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeSpreadLeg527NoOrderEvents) OrderEventTextMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeSpreadLeg527NoOrderEvents) OrderEventTextMinValue() byte {
	return byte(32)
}

func (*ExecutionReportTradeSpreadLeg527NoOrderEvents) OrderEventTextMaxValue() byte {
	return byte(126)
}

func (*ExecutionReportTradeSpreadLeg527NoOrderEvents) OrderEventTextNullValue() byte {
	return 0
}

func (e *ExecutionReportTradeSpreadLeg527NoOrderEvents) OrderEventTextCharacterEncoding() string {
	return "US-ASCII"
}

func (*ExecutionReportTradeSpreadLeg527NoOrderEvents) OrderEventExecIDId() uint16 {
	return 1797
}

func (*ExecutionReportTradeSpreadLeg527NoOrderEvents) OrderEventExecIDSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeSpreadLeg527NoOrderEvents) OrderEventExecIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.OrderEventExecIDSinceVersion()
}

func (*ExecutionReportTradeSpreadLeg527NoOrderEvents) OrderEventExecIDDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeSpreadLeg527NoOrderEvents) OrderEventExecIDMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeSpreadLeg527NoOrderEvents) OrderEventExecIDMinValue() uint32 {
	return 0
}

func (*ExecutionReportTradeSpreadLeg527NoOrderEvents) OrderEventExecIDMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*ExecutionReportTradeSpreadLeg527NoOrderEvents) OrderEventExecIDNullValue() uint32 {
	return math.MaxUint32
}

func (*ExecutionReportTradeSpreadLeg527NoOrderEvents) OrderEventQtyId() uint16 {
	return 1800
}

func (*ExecutionReportTradeSpreadLeg527NoOrderEvents) OrderEventQtySinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeSpreadLeg527NoOrderEvents) OrderEventQtyInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.OrderEventQtySinceVersion()
}

func (*ExecutionReportTradeSpreadLeg527NoOrderEvents) OrderEventQtyDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeSpreadLeg527NoOrderEvents) OrderEventQtyMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeSpreadLeg527NoOrderEvents) OrderEventQtyMinValue() uint32 {
	return 0
}

func (*ExecutionReportTradeSpreadLeg527NoOrderEvents) OrderEventQtyMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*ExecutionReportTradeSpreadLeg527NoOrderEvents) OrderEventQtyNullValue() uint32 {
	return math.MaxUint32
}

func (*ExecutionReportTradeSpreadLeg527NoOrderEvents) OrderEventTypeId() uint16 {
	return 1796
}

func (*ExecutionReportTradeSpreadLeg527NoOrderEvents) OrderEventTypeSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeSpreadLeg527NoOrderEvents) OrderEventTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.OrderEventTypeSinceVersion()
}

func (*ExecutionReportTradeSpreadLeg527NoOrderEvents) OrderEventTypeDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeSpreadLeg527NoOrderEvents) OrderEventTypeMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeSpreadLeg527NoOrderEvents) OrderEventReasonId() uint16 {
	return 1798
}

func (*ExecutionReportTradeSpreadLeg527NoOrderEvents) OrderEventReasonSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeSpreadLeg527NoOrderEvents) OrderEventReasonInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.OrderEventReasonSinceVersion()
}

func (*ExecutionReportTradeSpreadLeg527NoOrderEvents) OrderEventReasonDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeSpreadLeg527NoOrderEvents) OrderEventReasonMetaAttribute(meta int) string {
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

func (*ExecutionReportTradeSpreadLeg527NoOrderEvents) OrderEventReasonMinValue() uint8 {
	return 0
}

func (*ExecutionReportTradeSpreadLeg527NoOrderEvents) OrderEventReasonMaxValue() uint8 {
	return math.MaxUint8 - 1
}

func (*ExecutionReportTradeSpreadLeg527NoOrderEvents) OrderEventReasonNullValue() uint8 {
	return math.MaxUint8
}

func (*ExecutionReportTradeSpreadLeg527) NoFillsId() uint16 {
	return 1362
}

func (*ExecutionReportTradeSpreadLeg527) NoFillsSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeSpreadLeg527) NoFillsInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.NoFillsSinceVersion()
}

func (*ExecutionReportTradeSpreadLeg527) NoFillsDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeSpreadLeg527NoFills) SbeBlockLength() (blockLength uint) {
	return 15
}

func (*ExecutionReportTradeSpreadLeg527NoFills) SbeSchemaVersion() (schemaVersion uint16) {
	return 5
}

func (*ExecutionReportTradeSpreadLeg527) NoOrderEventsId() uint16 {
	return 1795
}

func (*ExecutionReportTradeSpreadLeg527) NoOrderEventsSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTradeSpreadLeg527) NoOrderEventsInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.NoOrderEventsSinceVersion()
}

func (*ExecutionReportTradeSpreadLeg527) NoOrderEventsDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTradeSpreadLeg527NoOrderEvents) SbeBlockLength() (blockLength uint) {
	return 23
}

func (*ExecutionReportTradeSpreadLeg527NoOrderEvents) SbeSchemaVersion() (schemaVersion uint16) {
	return 5
}
