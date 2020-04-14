// Generated SBE (Simple Binary Encoding) message codec

package ilinkbinary

import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
)

type ExecutionReportStatus532 struct {
	SeqNum                uint32
	UUID                  uint64
	Text                  [256]byte
	ExecID                [40]byte
	SenderID              [20]byte
	ClOrdID               [20]byte
	PartyDetailsListReqID uint64
	OrderID               uint64
	Price                 PRICENULL9
	StopPx                PRICENULL9
	TransactTime          uint64
	SendingTimeEpoch      uint64
	OrderRequestID        uint64
	OrdStatusReqID        uint64
	MassStatusReqID       uint64
	CrossID               uint64
	HostCrossID           uint64
	Location              [5]byte
	SecurityID            int32
	OrderQty              uint32
	CumQty                uint32
	LeavesQty             uint32
	MinQty                uint32
	DisplayQty            uint32
	ExpireDate            uint16
	OrdStatus             OrderStatusEnum
	ExecType              [1]byte
	OrdType               OrderTypeEnum
	Side                  SideReqEnum
	TimeInForce           TimeInForceEnum
	ManualOrderIndicator  ManualOrdIndReqEnum
	PossRetransFlag       BooleanFlagEnum
	LastRptRequested      BooleanNULLEnum
	CrossType             uint8
	ExecInst              ExecInst
	ExecutionMode         ExecModeEnum
	LiquidityFlag         BooleanNULLEnum
	ManagedOrder          BooleanNULLEnum
	ShortSaleType         ShortSaleTypeEnum
}

func (e *ExecutionReportStatus532) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
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
	if err := _m.WriteBytes(_w, e.Text[:]); err != nil {
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
	if err := _m.WriteUint64(_w, e.OrdStatusReqID); err != nil {
		return err
	}
	if err := _m.WriteUint64(_w, e.MassStatusReqID); err != nil {
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
	if err := _m.WriteUint32(_w, e.CumQty); err != nil {
		return err
	}
	if err := _m.WriteUint32(_w, e.LeavesQty); err != nil {
		return err
	}
	if err := _m.WriteUint32(_w, e.MinQty); err != nil {
		return err
	}
	if err := _m.WriteUint32(_w, e.DisplayQty); err != nil {
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
	if err := e.LastRptRequested.Encode(_m, _w); err != nil {
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
	return nil
}

func (e *ExecutionReportStatus532) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
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
	if !e.TextInActingVersion(actingVersion) {
		for idx := 0; idx < 256; idx++ {
			e.Text[idx] = e.TextNullValue()
		}
	} else {
		if err := _m.ReadBytes(_r, e.Text[:]); err != nil {
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
	if !e.OrdStatusReqIDInActingVersion(actingVersion) {
		e.OrdStatusReqID = e.OrdStatusReqIDNullValue()
	} else {
		if err := _m.ReadUint64(_r, &e.OrdStatusReqID); err != nil {
			return err
		}
	}
	if !e.MassStatusReqIDInActingVersion(actingVersion) {
		e.MassStatusReqID = e.MassStatusReqIDNullValue()
	} else {
		if err := _m.ReadUint64(_r, &e.MassStatusReqID); err != nil {
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
	if !e.CumQtyInActingVersion(actingVersion) {
		e.CumQty = e.CumQtyNullValue()
	} else {
		if err := _m.ReadUint32(_r, &e.CumQty); err != nil {
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
	if !e.MinQtyInActingVersion(actingVersion) {
		e.MinQty = e.MinQtyNullValue()
	} else {
		if err := _m.ReadUint32(_r, &e.MinQty); err != nil {
			return err
		}
	}
	if !e.DisplayQtyInActingVersion(actingVersion) {
		e.DisplayQty = e.DisplayQtyNullValue()
	} else {
		if err := _m.ReadUint32(_r, &e.DisplayQty); err != nil {
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
	e.ExecType[0] = 73
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
	if e.LastRptRequestedInActingVersion(actingVersion) {
		if err := e.LastRptRequested.Decode(_m, _r, actingVersion); err != nil {
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
	if actingVersion > e.SbeSchemaVersion() && blockLength > e.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-e.SbeBlockLength()))
	}
	if doRangeCheck {
		if err := e.RangeCheck(actingVersion, e.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	return nil
}

func (e *ExecutionReportStatus532) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
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
	if e.TextInActingVersion(actingVersion) {
		for idx := 0; idx < 256; idx++ {
			if e.Text[idx] < e.TextMinValue() || e.Text[idx] > e.TextMaxValue() {
				return fmt.Errorf("Range check failed on e.Text[%d] (%v < %v > %v)", idx, e.TextMinValue(), e.Text[idx], e.TextMaxValue())
			}
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
	if e.OrdStatusReqIDInActingVersion(actingVersion) {
		if e.OrdStatusReqID != e.OrdStatusReqIDNullValue() && (e.OrdStatusReqID < e.OrdStatusReqIDMinValue() || e.OrdStatusReqID > e.OrdStatusReqIDMaxValue()) {
			return fmt.Errorf("Range check failed on e.OrdStatusReqID (%v < %v > %v)", e.OrdStatusReqIDMinValue(), e.OrdStatusReqID, e.OrdStatusReqIDMaxValue())
		}
	}
	if e.MassStatusReqIDInActingVersion(actingVersion) {
		if e.MassStatusReqID != e.MassStatusReqIDNullValue() && (e.MassStatusReqID < e.MassStatusReqIDMinValue() || e.MassStatusReqID > e.MassStatusReqIDMaxValue()) {
			return fmt.Errorf("Range check failed on e.MassStatusReqID (%v < %v > %v)", e.MassStatusReqIDMinValue(), e.MassStatusReqID, e.MassStatusReqIDMaxValue())
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
	if e.CumQtyInActingVersion(actingVersion) {
		if e.CumQty < e.CumQtyMinValue() || e.CumQty > e.CumQtyMaxValue() {
			return fmt.Errorf("Range check failed on e.CumQty (%v < %v > %v)", e.CumQtyMinValue(), e.CumQty, e.CumQtyMaxValue())
		}
	}
	if e.LeavesQtyInActingVersion(actingVersion) {
		if e.LeavesQty < e.LeavesQtyMinValue() || e.LeavesQty > e.LeavesQtyMaxValue() {
			return fmt.Errorf("Range check failed on e.LeavesQty (%v < %v > %v)", e.LeavesQtyMinValue(), e.LeavesQty, e.LeavesQtyMaxValue())
		}
	}
	if e.MinQtyInActingVersion(actingVersion) {
		if e.MinQty != e.MinQtyNullValue() && (e.MinQty < e.MinQtyMinValue() || e.MinQty > e.MinQtyMaxValue()) {
			return fmt.Errorf("Range check failed on e.MinQty (%v < %v > %v)", e.MinQtyMinValue(), e.MinQty, e.MinQtyMaxValue())
		}
	}
	if e.DisplayQtyInActingVersion(actingVersion) {
		if e.DisplayQty != e.DisplayQtyNullValue() && (e.DisplayQty < e.DisplayQtyMinValue() || e.DisplayQty > e.DisplayQtyMaxValue()) {
			return fmt.Errorf("Range check failed on e.DisplayQty (%v < %v > %v)", e.DisplayQtyMinValue(), e.DisplayQty, e.DisplayQtyMaxValue())
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
	if err := e.LastRptRequested.RangeCheck(actingVersion, schemaVersion); err != nil {
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
	return nil
}

func ExecutionReportStatus532Init(e *ExecutionReportStatus532) {
	for idx := 0; idx < 256; idx++ {
		e.Text[idx] = 0
	}
	e.OrdStatusReqID = 18446744073709551615
	e.MassStatusReqID = 18446744073709551615
	e.CrossID = 18446744073709551615
	e.HostCrossID = 18446744073709551615
	e.MinQty = 4294967295
	e.DisplayQty = 4294967295
	e.ExpireDate = 65535
	e.ExecType[0] = 73
	e.CrossType = 255
	return
}

func (*ExecutionReportStatus532) SbeBlockLength() (blockLength uint16) {
	return 480
}

func (*ExecutionReportStatus532) SbeTemplateId() (templateId uint16) {
	return 532
}

func (*ExecutionReportStatus532) SbeSchemaId() (schemaId uint16) {
	return 8
}

func (*ExecutionReportStatus532) SbeSchemaVersion() (schemaVersion uint16) {
	return 5
}

func (*ExecutionReportStatus532) SbeSemanticType() (semanticType []byte) {
	return []byte("8")
}

func (*ExecutionReportStatus532) SeqNumId() uint16 {
	return 9726
}

func (*ExecutionReportStatus532) SeqNumSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportStatus532) SeqNumInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.SeqNumSinceVersion()
}

func (*ExecutionReportStatus532) SeqNumDeprecated() uint16 {
	return 0
}

func (*ExecutionReportStatus532) SeqNumMetaAttribute(meta int) string {
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

func (*ExecutionReportStatus532) SeqNumMinValue() uint32 {
	return 0
}

func (*ExecutionReportStatus532) SeqNumMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*ExecutionReportStatus532) SeqNumNullValue() uint32 {
	return math.MaxUint32
}

func (*ExecutionReportStatus532) UUIDId() uint16 {
	return 39001
}

func (*ExecutionReportStatus532) UUIDSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportStatus532) UUIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.UUIDSinceVersion()
}

func (*ExecutionReportStatus532) UUIDDeprecated() uint16 {
	return 0
}

func (*ExecutionReportStatus532) UUIDMetaAttribute(meta int) string {
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

func (*ExecutionReportStatus532) UUIDMinValue() uint64 {
	return 0
}

func (*ExecutionReportStatus532) UUIDMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*ExecutionReportStatus532) UUIDNullValue() uint64 {
	return math.MaxUint64
}

func (*ExecutionReportStatus532) TextId() uint16 {
	return 58
}

func (*ExecutionReportStatus532) TextSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportStatus532) TextInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.TextSinceVersion()
}

func (*ExecutionReportStatus532) TextDeprecated() uint16 {
	return 0
}

func (*ExecutionReportStatus532) TextMetaAttribute(meta int) string {
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

func (*ExecutionReportStatus532) TextMinValue() byte {
	return byte(32)
}

func (*ExecutionReportStatus532) TextMaxValue() byte {
	return byte(126)
}

func (*ExecutionReportStatus532) TextNullValue() byte {
	return 0
}

func (e *ExecutionReportStatus532) TextCharacterEncoding() string {
	return "US-ASCII"
}

func (*ExecutionReportStatus532) ExecIDId() uint16 {
	return 17
}

func (*ExecutionReportStatus532) ExecIDSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportStatus532) ExecIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.ExecIDSinceVersion()
}

func (*ExecutionReportStatus532) ExecIDDeprecated() uint16 {
	return 0
}

func (*ExecutionReportStatus532) ExecIDMetaAttribute(meta int) string {
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

func (*ExecutionReportStatus532) ExecIDMinValue() byte {
	return byte(32)
}

func (*ExecutionReportStatus532) ExecIDMaxValue() byte {
	return byte(126)
}

func (*ExecutionReportStatus532) ExecIDNullValue() byte {
	return 0
}

func (e *ExecutionReportStatus532) ExecIDCharacterEncoding() string {
	return "US-ASCII"
}

func (*ExecutionReportStatus532) SenderIDId() uint16 {
	return 5392
}

func (*ExecutionReportStatus532) SenderIDSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportStatus532) SenderIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.SenderIDSinceVersion()
}

func (*ExecutionReportStatus532) SenderIDDeprecated() uint16 {
	return 0
}

func (*ExecutionReportStatus532) SenderIDMetaAttribute(meta int) string {
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

func (*ExecutionReportStatus532) SenderIDMinValue() byte {
	return byte(32)
}

func (*ExecutionReportStatus532) SenderIDMaxValue() byte {
	return byte(126)
}

func (*ExecutionReportStatus532) SenderIDNullValue() byte {
	return 0
}

func (e *ExecutionReportStatus532) SenderIDCharacterEncoding() string {
	return "US-ASCII"
}

func (*ExecutionReportStatus532) ClOrdIDId() uint16 {
	return 11
}

func (*ExecutionReportStatus532) ClOrdIDSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportStatus532) ClOrdIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.ClOrdIDSinceVersion()
}

func (*ExecutionReportStatus532) ClOrdIDDeprecated() uint16 {
	return 0
}

func (*ExecutionReportStatus532) ClOrdIDMetaAttribute(meta int) string {
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

func (*ExecutionReportStatus532) ClOrdIDMinValue() byte {
	return byte(32)
}

func (*ExecutionReportStatus532) ClOrdIDMaxValue() byte {
	return byte(126)
}

func (*ExecutionReportStatus532) ClOrdIDNullValue() byte {
	return 0
}

func (e *ExecutionReportStatus532) ClOrdIDCharacterEncoding() string {
	return "US-ASCII"
}

func (*ExecutionReportStatus532) PartyDetailsListReqIDId() uint16 {
	return 1505
}

func (*ExecutionReportStatus532) PartyDetailsListReqIDSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportStatus532) PartyDetailsListReqIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.PartyDetailsListReqIDSinceVersion()
}

func (*ExecutionReportStatus532) PartyDetailsListReqIDDeprecated() uint16 {
	return 0
}

func (*ExecutionReportStatus532) PartyDetailsListReqIDMetaAttribute(meta int) string {
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

func (*ExecutionReportStatus532) PartyDetailsListReqIDMinValue() uint64 {
	return 0
}

func (*ExecutionReportStatus532) PartyDetailsListReqIDMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*ExecutionReportStatus532) PartyDetailsListReqIDNullValue() uint64 {
	return math.MaxUint64
}

func (*ExecutionReportStatus532) OrderIDId() uint16 {
	return 37
}

func (*ExecutionReportStatus532) OrderIDSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportStatus532) OrderIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.OrderIDSinceVersion()
}

func (*ExecutionReportStatus532) OrderIDDeprecated() uint16 {
	return 0
}

func (*ExecutionReportStatus532) OrderIDMetaAttribute(meta int) string {
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

func (*ExecutionReportStatus532) OrderIDMinValue() uint64 {
	return 0
}

func (*ExecutionReportStatus532) OrderIDMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*ExecutionReportStatus532) OrderIDNullValue() uint64 {
	return math.MaxUint64
}

func (*ExecutionReportStatus532) PriceId() uint16 {
	return 44
}

func (*ExecutionReportStatus532) PriceSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportStatus532) PriceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.PriceSinceVersion()
}

func (*ExecutionReportStatus532) PriceDeprecated() uint16 {
	return 0
}

func (*ExecutionReportStatus532) PriceMetaAttribute(meta int) string {
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

func (*ExecutionReportStatus532) StopPxId() uint16 {
	return 99
}

func (*ExecutionReportStatus532) StopPxSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportStatus532) StopPxInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.StopPxSinceVersion()
}

func (*ExecutionReportStatus532) StopPxDeprecated() uint16 {
	return 0
}

func (*ExecutionReportStatus532) StopPxMetaAttribute(meta int) string {
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

func (*ExecutionReportStatus532) TransactTimeId() uint16 {
	return 60
}

func (*ExecutionReportStatus532) TransactTimeSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportStatus532) TransactTimeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.TransactTimeSinceVersion()
}

func (*ExecutionReportStatus532) TransactTimeDeprecated() uint16 {
	return 0
}

func (*ExecutionReportStatus532) TransactTimeMetaAttribute(meta int) string {
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

func (*ExecutionReportStatus532) TransactTimeMinValue() uint64 {
	return 0
}

func (*ExecutionReportStatus532) TransactTimeMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*ExecutionReportStatus532) TransactTimeNullValue() uint64 {
	return math.MaxUint64
}

func (*ExecutionReportStatus532) SendingTimeEpochId() uint16 {
	return 5297
}

func (*ExecutionReportStatus532) SendingTimeEpochSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportStatus532) SendingTimeEpochInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.SendingTimeEpochSinceVersion()
}

func (*ExecutionReportStatus532) SendingTimeEpochDeprecated() uint16 {
	return 0
}

func (*ExecutionReportStatus532) SendingTimeEpochMetaAttribute(meta int) string {
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

func (*ExecutionReportStatus532) SendingTimeEpochMinValue() uint64 {
	return 0
}

func (*ExecutionReportStatus532) SendingTimeEpochMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*ExecutionReportStatus532) SendingTimeEpochNullValue() uint64 {
	return math.MaxUint64
}

func (*ExecutionReportStatus532) OrderRequestIDId() uint16 {
	return 2422
}

func (*ExecutionReportStatus532) OrderRequestIDSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportStatus532) OrderRequestIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.OrderRequestIDSinceVersion()
}

func (*ExecutionReportStatus532) OrderRequestIDDeprecated() uint16 {
	return 0
}

func (*ExecutionReportStatus532) OrderRequestIDMetaAttribute(meta int) string {
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

func (*ExecutionReportStatus532) OrderRequestIDMinValue() uint64 {
	return 0
}

func (*ExecutionReportStatus532) OrderRequestIDMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*ExecutionReportStatus532) OrderRequestIDNullValue() uint64 {
	return math.MaxUint64
}

func (*ExecutionReportStatus532) OrdStatusReqIDId() uint16 {
	return 790
}

func (*ExecutionReportStatus532) OrdStatusReqIDSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportStatus532) OrdStatusReqIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.OrdStatusReqIDSinceVersion()
}

func (*ExecutionReportStatus532) OrdStatusReqIDDeprecated() uint16 {
	return 0
}

func (*ExecutionReportStatus532) OrdStatusReqIDMetaAttribute(meta int) string {
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

func (*ExecutionReportStatus532) OrdStatusReqIDMinValue() uint64 {
	return 0
}

func (*ExecutionReportStatus532) OrdStatusReqIDMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*ExecutionReportStatus532) OrdStatusReqIDNullValue() uint64 {
	return 18446744073709551615
}

func (*ExecutionReportStatus532) MassStatusReqIDId() uint16 {
	return 584
}

func (*ExecutionReportStatus532) MassStatusReqIDSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportStatus532) MassStatusReqIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.MassStatusReqIDSinceVersion()
}

func (*ExecutionReportStatus532) MassStatusReqIDDeprecated() uint16 {
	return 0
}

func (*ExecutionReportStatus532) MassStatusReqIDMetaAttribute(meta int) string {
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

func (*ExecutionReportStatus532) MassStatusReqIDMinValue() uint64 {
	return 0
}

func (*ExecutionReportStatus532) MassStatusReqIDMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*ExecutionReportStatus532) MassStatusReqIDNullValue() uint64 {
	return 18446744073709551615
}

func (*ExecutionReportStatus532) CrossIDId() uint16 {
	return 548
}

func (*ExecutionReportStatus532) CrossIDSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportStatus532) CrossIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.CrossIDSinceVersion()
}

func (*ExecutionReportStatus532) CrossIDDeprecated() uint16 {
	return 0
}

func (*ExecutionReportStatus532) CrossIDMetaAttribute(meta int) string {
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

func (*ExecutionReportStatus532) CrossIDMinValue() uint64 {
	return 0
}

func (*ExecutionReportStatus532) CrossIDMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*ExecutionReportStatus532) CrossIDNullValue() uint64 {
	return 18446744073709551615
}

func (*ExecutionReportStatus532) HostCrossIDId() uint16 {
	return 961
}

func (*ExecutionReportStatus532) HostCrossIDSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportStatus532) HostCrossIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.HostCrossIDSinceVersion()
}

func (*ExecutionReportStatus532) HostCrossIDDeprecated() uint16 {
	return 0
}

func (*ExecutionReportStatus532) HostCrossIDMetaAttribute(meta int) string {
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

func (*ExecutionReportStatus532) HostCrossIDMinValue() uint64 {
	return 0
}

func (*ExecutionReportStatus532) HostCrossIDMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*ExecutionReportStatus532) HostCrossIDNullValue() uint64 {
	return 18446744073709551615
}

func (*ExecutionReportStatus532) LocationId() uint16 {
	return 9537
}

func (*ExecutionReportStatus532) LocationSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportStatus532) LocationInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.LocationSinceVersion()
}

func (*ExecutionReportStatus532) LocationDeprecated() uint16 {
	return 0
}

func (*ExecutionReportStatus532) LocationMetaAttribute(meta int) string {
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

func (*ExecutionReportStatus532) LocationMinValue() byte {
	return byte(32)
}

func (*ExecutionReportStatus532) LocationMaxValue() byte {
	return byte(126)
}

func (*ExecutionReportStatus532) LocationNullValue() byte {
	return 0
}

func (e *ExecutionReportStatus532) LocationCharacterEncoding() string {
	return "US-ASCII"
}

func (*ExecutionReportStatus532) SecurityIDId() uint16 {
	return 48
}

func (*ExecutionReportStatus532) SecurityIDSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportStatus532) SecurityIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.SecurityIDSinceVersion()
}

func (*ExecutionReportStatus532) SecurityIDDeprecated() uint16 {
	return 0
}

func (*ExecutionReportStatus532) SecurityIDMetaAttribute(meta int) string {
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

func (*ExecutionReportStatus532) SecurityIDMinValue() int32 {
	return math.MinInt32 + 1
}

func (*ExecutionReportStatus532) SecurityIDMaxValue() int32 {
	return math.MaxInt32
}

func (*ExecutionReportStatus532) SecurityIDNullValue() int32 {
	return math.MinInt32
}

func (*ExecutionReportStatus532) OrderQtyId() uint16 {
	return 38
}

func (*ExecutionReportStatus532) OrderQtySinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportStatus532) OrderQtyInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.OrderQtySinceVersion()
}

func (*ExecutionReportStatus532) OrderQtyDeprecated() uint16 {
	return 0
}

func (*ExecutionReportStatus532) OrderQtyMetaAttribute(meta int) string {
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

func (*ExecutionReportStatus532) OrderQtyMinValue() uint32 {
	return 0
}

func (*ExecutionReportStatus532) OrderQtyMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*ExecutionReportStatus532) OrderQtyNullValue() uint32 {
	return math.MaxUint32
}

func (*ExecutionReportStatus532) CumQtyId() uint16 {
	return 14
}

func (*ExecutionReportStatus532) CumQtySinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportStatus532) CumQtyInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.CumQtySinceVersion()
}

func (*ExecutionReportStatus532) CumQtyDeprecated() uint16 {
	return 0
}

func (*ExecutionReportStatus532) CumQtyMetaAttribute(meta int) string {
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

func (*ExecutionReportStatus532) CumQtyMinValue() uint32 {
	return 0
}

func (*ExecutionReportStatus532) CumQtyMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*ExecutionReportStatus532) CumQtyNullValue() uint32 {
	return math.MaxUint32
}

func (*ExecutionReportStatus532) LeavesQtyId() uint16 {
	return 151
}

func (*ExecutionReportStatus532) LeavesQtySinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportStatus532) LeavesQtyInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.LeavesQtySinceVersion()
}

func (*ExecutionReportStatus532) LeavesQtyDeprecated() uint16 {
	return 0
}

func (*ExecutionReportStatus532) LeavesQtyMetaAttribute(meta int) string {
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

func (*ExecutionReportStatus532) LeavesQtyMinValue() uint32 {
	return 0
}

func (*ExecutionReportStatus532) LeavesQtyMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*ExecutionReportStatus532) LeavesQtyNullValue() uint32 {
	return math.MaxUint32
}

func (*ExecutionReportStatus532) MinQtyId() uint16 {
	return 110
}

func (*ExecutionReportStatus532) MinQtySinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportStatus532) MinQtyInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.MinQtySinceVersion()
}

func (*ExecutionReportStatus532) MinQtyDeprecated() uint16 {
	return 0
}

func (*ExecutionReportStatus532) MinQtyMetaAttribute(meta int) string {
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

func (*ExecutionReportStatus532) MinQtyMinValue() uint32 {
	return 0
}

func (*ExecutionReportStatus532) MinQtyMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*ExecutionReportStatus532) MinQtyNullValue() uint32 {
	return 4294967295
}

func (*ExecutionReportStatus532) DisplayQtyId() uint16 {
	return 1138
}

func (*ExecutionReportStatus532) DisplayQtySinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportStatus532) DisplayQtyInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.DisplayQtySinceVersion()
}

func (*ExecutionReportStatus532) DisplayQtyDeprecated() uint16 {
	return 0
}

func (*ExecutionReportStatus532) DisplayQtyMetaAttribute(meta int) string {
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

func (*ExecutionReportStatus532) DisplayQtyMinValue() uint32 {
	return 0
}

func (*ExecutionReportStatus532) DisplayQtyMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*ExecutionReportStatus532) DisplayQtyNullValue() uint32 {
	return 4294967295
}

func (*ExecutionReportStatus532) ExpireDateId() uint16 {
	return 432
}

func (*ExecutionReportStatus532) ExpireDateSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportStatus532) ExpireDateInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.ExpireDateSinceVersion()
}

func (*ExecutionReportStatus532) ExpireDateDeprecated() uint16 {
	return 0
}

func (*ExecutionReportStatus532) ExpireDateMetaAttribute(meta int) string {
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

func (*ExecutionReportStatus532) ExpireDateMinValue() uint16 {
	return 0
}

func (*ExecutionReportStatus532) ExpireDateMaxValue() uint16 {
	return math.MaxUint16 - 1
}

func (*ExecutionReportStatus532) ExpireDateNullValue() uint16 {
	return 65535
}

func (*ExecutionReportStatus532) OrdStatusId() uint16 {
	return 39
}

func (*ExecutionReportStatus532) OrdStatusSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportStatus532) OrdStatusInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.OrdStatusSinceVersion()
}

func (*ExecutionReportStatus532) OrdStatusDeprecated() uint16 {
	return 0
}

func (*ExecutionReportStatus532) OrdStatusMetaAttribute(meta int) string {
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

func (*ExecutionReportStatus532) ExecTypeId() uint16 {
	return 150
}

func (*ExecutionReportStatus532) ExecTypeSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportStatus532) ExecTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.ExecTypeSinceVersion()
}

func (*ExecutionReportStatus532) ExecTypeDeprecated() uint16 {
	return 0
}

func (*ExecutionReportStatus532) ExecTypeMetaAttribute(meta int) string {
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

func (*ExecutionReportStatus532) ExecTypeMinValue() byte {
	return byte(32)
}

func (*ExecutionReportStatus532) ExecTypeMaxValue() byte {
	return byte(126)
}

func (*ExecutionReportStatus532) ExecTypeNullValue() byte {
	return 0
}

func (*ExecutionReportStatus532) OrdTypeId() uint16 {
	return 40
}

func (*ExecutionReportStatus532) OrdTypeSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportStatus532) OrdTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.OrdTypeSinceVersion()
}

func (*ExecutionReportStatus532) OrdTypeDeprecated() uint16 {
	return 0
}

func (*ExecutionReportStatus532) OrdTypeMetaAttribute(meta int) string {
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

func (*ExecutionReportStatus532) SideId() uint16 {
	return 54
}

func (*ExecutionReportStatus532) SideSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportStatus532) SideInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.SideSinceVersion()
}

func (*ExecutionReportStatus532) SideDeprecated() uint16 {
	return 0
}

func (*ExecutionReportStatus532) SideMetaAttribute(meta int) string {
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

func (*ExecutionReportStatus532) TimeInForceId() uint16 {
	return 59
}

func (*ExecutionReportStatus532) TimeInForceSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportStatus532) TimeInForceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.TimeInForceSinceVersion()
}

func (*ExecutionReportStatus532) TimeInForceDeprecated() uint16 {
	return 0
}

func (*ExecutionReportStatus532) TimeInForceMetaAttribute(meta int) string {
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

func (*ExecutionReportStatus532) ManualOrderIndicatorId() uint16 {
	return 1028
}

func (*ExecutionReportStatus532) ManualOrderIndicatorSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportStatus532) ManualOrderIndicatorInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.ManualOrderIndicatorSinceVersion()
}

func (*ExecutionReportStatus532) ManualOrderIndicatorDeprecated() uint16 {
	return 0
}

func (*ExecutionReportStatus532) ManualOrderIndicatorMetaAttribute(meta int) string {
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

func (*ExecutionReportStatus532) PossRetransFlagId() uint16 {
	return 9765
}

func (*ExecutionReportStatus532) PossRetransFlagSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportStatus532) PossRetransFlagInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.PossRetransFlagSinceVersion()
}

func (*ExecutionReportStatus532) PossRetransFlagDeprecated() uint16 {
	return 0
}

func (*ExecutionReportStatus532) PossRetransFlagMetaAttribute(meta int) string {
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

func (*ExecutionReportStatus532) LastRptRequestedId() uint16 {
	return 912
}

func (*ExecutionReportStatus532) LastRptRequestedSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportStatus532) LastRptRequestedInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.LastRptRequestedSinceVersion()
}

func (*ExecutionReportStatus532) LastRptRequestedDeprecated() uint16 {
	return 0
}

func (*ExecutionReportStatus532) LastRptRequestedMetaAttribute(meta int) string {
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

func (*ExecutionReportStatus532) CrossTypeId() uint16 {
	return 549
}

func (*ExecutionReportStatus532) CrossTypeSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportStatus532) CrossTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.CrossTypeSinceVersion()
}

func (*ExecutionReportStatus532) CrossTypeDeprecated() uint16 {
	return 0
}

func (*ExecutionReportStatus532) CrossTypeMetaAttribute(meta int) string {
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

func (*ExecutionReportStatus532) CrossTypeMinValue() uint8 {
	return 0
}

func (*ExecutionReportStatus532) CrossTypeMaxValue() uint8 {
	return math.MaxUint8 - 1
}

func (*ExecutionReportStatus532) CrossTypeNullValue() uint8 {
	return 255
}

func (*ExecutionReportStatus532) ExecInstId() uint16 {
	return 18
}

func (*ExecutionReportStatus532) ExecInstSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportStatus532) ExecInstInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.ExecInstSinceVersion()
}

func (*ExecutionReportStatus532) ExecInstDeprecated() uint16 {
	return 0
}

func (*ExecutionReportStatus532) ExecInstMetaAttribute(meta int) string {
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

func (*ExecutionReportStatus532) ExecutionModeId() uint16 {
	return 5906
}

func (*ExecutionReportStatus532) ExecutionModeSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportStatus532) ExecutionModeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.ExecutionModeSinceVersion()
}

func (*ExecutionReportStatus532) ExecutionModeDeprecated() uint16 {
	return 0
}

func (*ExecutionReportStatus532) ExecutionModeMetaAttribute(meta int) string {
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

func (*ExecutionReportStatus532) LiquidityFlagId() uint16 {
	return 9373
}

func (*ExecutionReportStatus532) LiquidityFlagSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportStatus532) LiquidityFlagInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.LiquidityFlagSinceVersion()
}

func (*ExecutionReportStatus532) LiquidityFlagDeprecated() uint16 {
	return 0
}

func (*ExecutionReportStatus532) LiquidityFlagMetaAttribute(meta int) string {
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

func (*ExecutionReportStatus532) ManagedOrderId() uint16 {
	return 6881
}

func (*ExecutionReportStatus532) ManagedOrderSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportStatus532) ManagedOrderInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.ManagedOrderSinceVersion()
}

func (*ExecutionReportStatus532) ManagedOrderDeprecated() uint16 {
	return 0
}

func (*ExecutionReportStatus532) ManagedOrderMetaAttribute(meta int) string {
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

func (*ExecutionReportStatus532) ShortSaleTypeId() uint16 {
	return 5409
}

func (*ExecutionReportStatus532) ShortSaleTypeSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportStatus532) ShortSaleTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.ShortSaleTypeSinceVersion()
}

func (*ExecutionReportStatus532) ShortSaleTypeDeprecated() uint16 {
	return 0
}

func (*ExecutionReportStatus532) ShortSaleTypeMetaAttribute(meta int) string {
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
