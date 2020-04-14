// Generated SBE (Simple Binary Encoding) message codec

package ilinkbinary

import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
)

type ExecutionReportModify531 struct {
	SeqNum                uint32
	UUID                  uint64
	ExecID                [40]byte
	SenderID              [20]byte
	ClOrdID               [20]byte
	PartyDetailsListReqID uint64
	OrderID               uint64
	Price                 PRICE9
	StopPx                PRICENULL9
	TransactTime          uint64
	SendingTimeEpoch      uint64
	OrderRequestID        uint64
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
	DelayDuration         uint16
	OrdStatus             [1]byte
	ExecType              [1]byte
	OrdType               OrderTypeEnum
	Side                  SideReqEnum
	TimeInForce           TimeInForceEnum
	ManualOrderIndicator  ManualOrdIndReqEnum
	PossRetransFlag       BooleanFlagEnum
	SplitMsg              SplitMsgEnum
	CrossType             uint8
	ExecInst              ExecInst
	ExecutionMode         ExecModeEnum
	LiquidityFlag         BooleanNULLEnum
	ManagedOrder          BooleanNULLEnum
	ShortSaleType         ShortSaleTypeEnum
	DelayToTime           uint64
}

func (e *ExecutionReportModify531) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
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
	if err := _m.WriteUint16(_w, e.DelayDuration); err != nil {
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
	if err := e.SplitMsg.Encode(_m, _w); err != nil {
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
	if err := _m.WriteUint64(_w, e.DelayToTime); err != nil {
		return err
	}
	return nil
}

func (e *ExecutionReportModify531) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
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
	if !e.DelayDurationInActingVersion(actingVersion) {
		e.DelayDuration = e.DelayDurationNullValue()
	} else {
		if err := _m.ReadUint16(_r, &e.DelayDuration); err != nil {
			return err
		}
	}
	e.OrdStatus[0] = 53
	e.ExecType[0] = 53
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
	if e.SplitMsgInActingVersion(actingVersion) {
		if err := e.SplitMsg.Decode(_m, _r, actingVersion); err != nil {
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
	if !e.DelayToTimeInActingVersion(actingVersion) {
		e.DelayToTime = e.DelayToTimeNullValue()
	} else {
		if err := _m.ReadUint64(_r, &e.DelayToTime); err != nil {
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

func (e *ExecutionReportModify531) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
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
	if e.DelayDurationInActingVersion(actingVersion) {
		if e.DelayDuration != e.DelayDurationNullValue() && (e.DelayDuration < e.DelayDurationMinValue() || e.DelayDuration > e.DelayDurationMaxValue()) {
			return fmt.Errorf("Range check failed on e.DelayDuration (%v < %v > %v)", e.DelayDurationMinValue(), e.DelayDuration, e.DelayDurationMaxValue())
		}
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
	if err := e.SplitMsg.RangeCheck(actingVersion, schemaVersion); err != nil {
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
	if e.DelayToTimeInActingVersion(actingVersion) {
		if e.DelayToTime != e.DelayToTimeNullValue() && (e.DelayToTime < e.DelayToTimeMinValue() || e.DelayToTime > e.DelayToTimeMaxValue()) {
			return fmt.Errorf("Range check failed on e.DelayToTime (%v < %v > %v)", e.DelayToTimeMinValue(), e.DelayToTime, e.DelayToTimeMaxValue())
		}
	}
	return nil
}

func ExecutionReportModify531Init(e *ExecutionReportModify531) {
	e.CrossID = 18446744073709551615
	e.HostCrossID = 18446744073709551615
	e.MinQty = 4294967295
	e.DisplayQty = 4294967295
	e.ExpireDate = 65535
	e.DelayDuration = 65535
	e.OrdStatus[0] = 53
	e.ExecType[0] = 53
	e.CrossType = 255
	e.DelayToTime = 18446744073709551615
	return
}

func (*ExecutionReportModify531) SbeBlockLength() (blockLength uint16) {
	return 217
}

func (*ExecutionReportModify531) SbeTemplateId() (templateId uint16) {
	return 531
}

func (*ExecutionReportModify531) SbeSchemaId() (schemaId uint16) {
	return 8
}

func (*ExecutionReportModify531) SbeSchemaVersion() (schemaVersion uint16) {
	return 5
}

func (*ExecutionReportModify531) SbeSemanticType() (semanticType []byte) {
	return []byte("8")
}

func (*ExecutionReportModify531) SeqNumId() uint16 {
	return 9726
}

func (*ExecutionReportModify531) SeqNumSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportModify531) SeqNumInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.SeqNumSinceVersion()
}

func (*ExecutionReportModify531) SeqNumDeprecated() uint16 {
	return 0
}

func (*ExecutionReportModify531) SeqNumMetaAttribute(meta int) string {
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

func (*ExecutionReportModify531) SeqNumMinValue() uint32 {
	return 0
}

func (*ExecutionReportModify531) SeqNumMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*ExecutionReportModify531) SeqNumNullValue() uint32 {
	return math.MaxUint32
}

func (*ExecutionReportModify531) UUIDId() uint16 {
	return 39001
}

func (*ExecutionReportModify531) UUIDSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportModify531) UUIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.UUIDSinceVersion()
}

func (*ExecutionReportModify531) UUIDDeprecated() uint16 {
	return 0
}

func (*ExecutionReportModify531) UUIDMetaAttribute(meta int) string {
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

func (*ExecutionReportModify531) UUIDMinValue() uint64 {
	return 0
}

func (*ExecutionReportModify531) UUIDMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*ExecutionReportModify531) UUIDNullValue() uint64 {
	return math.MaxUint64
}

func (*ExecutionReportModify531) ExecIDId() uint16 {
	return 17
}

func (*ExecutionReportModify531) ExecIDSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportModify531) ExecIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.ExecIDSinceVersion()
}

func (*ExecutionReportModify531) ExecIDDeprecated() uint16 {
	return 0
}

func (*ExecutionReportModify531) ExecIDMetaAttribute(meta int) string {
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

func (*ExecutionReportModify531) ExecIDMinValue() byte {
	return byte(32)
}

func (*ExecutionReportModify531) ExecIDMaxValue() byte {
	return byte(126)
}

func (*ExecutionReportModify531) ExecIDNullValue() byte {
	return 0
}

func (e *ExecutionReportModify531) ExecIDCharacterEncoding() string {
	return "US-ASCII"
}

func (*ExecutionReportModify531) SenderIDId() uint16 {
	return 5392
}

func (*ExecutionReportModify531) SenderIDSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportModify531) SenderIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.SenderIDSinceVersion()
}

func (*ExecutionReportModify531) SenderIDDeprecated() uint16 {
	return 0
}

func (*ExecutionReportModify531) SenderIDMetaAttribute(meta int) string {
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

func (*ExecutionReportModify531) SenderIDMinValue() byte {
	return byte(32)
}

func (*ExecutionReportModify531) SenderIDMaxValue() byte {
	return byte(126)
}

func (*ExecutionReportModify531) SenderIDNullValue() byte {
	return 0
}

func (e *ExecutionReportModify531) SenderIDCharacterEncoding() string {
	return "US-ASCII"
}

func (*ExecutionReportModify531) ClOrdIDId() uint16 {
	return 11
}

func (*ExecutionReportModify531) ClOrdIDSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportModify531) ClOrdIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.ClOrdIDSinceVersion()
}

func (*ExecutionReportModify531) ClOrdIDDeprecated() uint16 {
	return 0
}

func (*ExecutionReportModify531) ClOrdIDMetaAttribute(meta int) string {
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

func (*ExecutionReportModify531) ClOrdIDMinValue() byte {
	return byte(32)
}

func (*ExecutionReportModify531) ClOrdIDMaxValue() byte {
	return byte(126)
}

func (*ExecutionReportModify531) ClOrdIDNullValue() byte {
	return 0
}

func (e *ExecutionReportModify531) ClOrdIDCharacterEncoding() string {
	return "US-ASCII"
}

func (*ExecutionReportModify531) PartyDetailsListReqIDId() uint16 {
	return 1505
}

func (*ExecutionReportModify531) PartyDetailsListReqIDSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportModify531) PartyDetailsListReqIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.PartyDetailsListReqIDSinceVersion()
}

func (*ExecutionReportModify531) PartyDetailsListReqIDDeprecated() uint16 {
	return 0
}

func (*ExecutionReportModify531) PartyDetailsListReqIDMetaAttribute(meta int) string {
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

func (*ExecutionReportModify531) PartyDetailsListReqIDMinValue() uint64 {
	return 0
}

func (*ExecutionReportModify531) PartyDetailsListReqIDMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*ExecutionReportModify531) PartyDetailsListReqIDNullValue() uint64 {
	return math.MaxUint64
}

func (*ExecutionReportModify531) OrderIDId() uint16 {
	return 37
}

func (*ExecutionReportModify531) OrderIDSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportModify531) OrderIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.OrderIDSinceVersion()
}

func (*ExecutionReportModify531) OrderIDDeprecated() uint16 {
	return 0
}

func (*ExecutionReportModify531) OrderIDMetaAttribute(meta int) string {
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

func (*ExecutionReportModify531) OrderIDMinValue() uint64 {
	return 0
}

func (*ExecutionReportModify531) OrderIDMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*ExecutionReportModify531) OrderIDNullValue() uint64 {
	return math.MaxUint64
}

func (*ExecutionReportModify531) PriceId() uint16 {
	return 44
}

func (*ExecutionReportModify531) PriceSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportModify531) PriceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.PriceSinceVersion()
}

func (*ExecutionReportModify531) PriceDeprecated() uint16 {
	return 0
}

func (*ExecutionReportModify531) PriceMetaAttribute(meta int) string {
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

func (*ExecutionReportModify531) StopPxId() uint16 {
	return 99
}

func (*ExecutionReportModify531) StopPxSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportModify531) StopPxInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.StopPxSinceVersion()
}

func (*ExecutionReportModify531) StopPxDeprecated() uint16 {
	return 0
}

func (*ExecutionReportModify531) StopPxMetaAttribute(meta int) string {
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

func (*ExecutionReportModify531) TransactTimeId() uint16 {
	return 60
}

func (*ExecutionReportModify531) TransactTimeSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportModify531) TransactTimeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.TransactTimeSinceVersion()
}

func (*ExecutionReportModify531) TransactTimeDeprecated() uint16 {
	return 0
}

func (*ExecutionReportModify531) TransactTimeMetaAttribute(meta int) string {
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

func (*ExecutionReportModify531) TransactTimeMinValue() uint64 {
	return 0
}

func (*ExecutionReportModify531) TransactTimeMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*ExecutionReportModify531) TransactTimeNullValue() uint64 {
	return math.MaxUint64
}

func (*ExecutionReportModify531) SendingTimeEpochId() uint16 {
	return 5297
}

func (*ExecutionReportModify531) SendingTimeEpochSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportModify531) SendingTimeEpochInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.SendingTimeEpochSinceVersion()
}

func (*ExecutionReportModify531) SendingTimeEpochDeprecated() uint16 {
	return 0
}

func (*ExecutionReportModify531) SendingTimeEpochMetaAttribute(meta int) string {
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

func (*ExecutionReportModify531) SendingTimeEpochMinValue() uint64 {
	return 0
}

func (*ExecutionReportModify531) SendingTimeEpochMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*ExecutionReportModify531) SendingTimeEpochNullValue() uint64 {
	return math.MaxUint64
}

func (*ExecutionReportModify531) OrderRequestIDId() uint16 {
	return 2422
}

func (*ExecutionReportModify531) OrderRequestIDSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportModify531) OrderRequestIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.OrderRequestIDSinceVersion()
}

func (*ExecutionReportModify531) OrderRequestIDDeprecated() uint16 {
	return 0
}

func (*ExecutionReportModify531) OrderRequestIDMetaAttribute(meta int) string {
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

func (*ExecutionReportModify531) OrderRequestIDMinValue() uint64 {
	return 0
}

func (*ExecutionReportModify531) OrderRequestIDMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*ExecutionReportModify531) OrderRequestIDNullValue() uint64 {
	return math.MaxUint64
}

func (*ExecutionReportModify531) CrossIDId() uint16 {
	return 548
}

func (*ExecutionReportModify531) CrossIDSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportModify531) CrossIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.CrossIDSinceVersion()
}

func (*ExecutionReportModify531) CrossIDDeprecated() uint16 {
	return 0
}

func (*ExecutionReportModify531) CrossIDMetaAttribute(meta int) string {
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

func (*ExecutionReportModify531) CrossIDMinValue() uint64 {
	return 0
}

func (*ExecutionReportModify531) CrossIDMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*ExecutionReportModify531) CrossIDNullValue() uint64 {
	return 18446744073709551615
}

func (*ExecutionReportModify531) HostCrossIDId() uint16 {
	return 961
}

func (*ExecutionReportModify531) HostCrossIDSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportModify531) HostCrossIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.HostCrossIDSinceVersion()
}

func (*ExecutionReportModify531) HostCrossIDDeprecated() uint16 {
	return 0
}

func (*ExecutionReportModify531) HostCrossIDMetaAttribute(meta int) string {
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

func (*ExecutionReportModify531) HostCrossIDMinValue() uint64 {
	return 0
}

func (*ExecutionReportModify531) HostCrossIDMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*ExecutionReportModify531) HostCrossIDNullValue() uint64 {
	return 18446744073709551615
}

func (*ExecutionReportModify531) LocationId() uint16 {
	return 9537
}

func (*ExecutionReportModify531) LocationSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportModify531) LocationInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.LocationSinceVersion()
}

func (*ExecutionReportModify531) LocationDeprecated() uint16 {
	return 0
}

func (*ExecutionReportModify531) LocationMetaAttribute(meta int) string {
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

func (*ExecutionReportModify531) LocationMinValue() byte {
	return byte(32)
}

func (*ExecutionReportModify531) LocationMaxValue() byte {
	return byte(126)
}

func (*ExecutionReportModify531) LocationNullValue() byte {
	return 0
}

func (e *ExecutionReportModify531) LocationCharacterEncoding() string {
	return "US-ASCII"
}

func (*ExecutionReportModify531) SecurityIDId() uint16 {
	return 48
}

func (*ExecutionReportModify531) SecurityIDSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportModify531) SecurityIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.SecurityIDSinceVersion()
}

func (*ExecutionReportModify531) SecurityIDDeprecated() uint16 {
	return 0
}

func (*ExecutionReportModify531) SecurityIDMetaAttribute(meta int) string {
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

func (*ExecutionReportModify531) SecurityIDMinValue() int32 {
	return math.MinInt32 + 1
}

func (*ExecutionReportModify531) SecurityIDMaxValue() int32 {
	return math.MaxInt32
}

func (*ExecutionReportModify531) SecurityIDNullValue() int32 {
	return math.MinInt32
}

func (*ExecutionReportModify531) OrderQtyId() uint16 {
	return 38
}

func (*ExecutionReportModify531) OrderQtySinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportModify531) OrderQtyInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.OrderQtySinceVersion()
}

func (*ExecutionReportModify531) OrderQtyDeprecated() uint16 {
	return 0
}

func (*ExecutionReportModify531) OrderQtyMetaAttribute(meta int) string {
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

func (*ExecutionReportModify531) OrderQtyMinValue() uint32 {
	return 0
}

func (*ExecutionReportModify531) OrderQtyMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*ExecutionReportModify531) OrderQtyNullValue() uint32 {
	return math.MaxUint32
}

func (*ExecutionReportModify531) CumQtyId() uint16 {
	return 14
}

func (*ExecutionReportModify531) CumQtySinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportModify531) CumQtyInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.CumQtySinceVersion()
}

func (*ExecutionReportModify531) CumQtyDeprecated() uint16 {
	return 0
}

func (*ExecutionReportModify531) CumQtyMetaAttribute(meta int) string {
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

func (*ExecutionReportModify531) CumQtyMinValue() uint32 {
	return 0
}

func (*ExecutionReportModify531) CumQtyMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*ExecutionReportModify531) CumQtyNullValue() uint32 {
	return math.MaxUint32
}

func (*ExecutionReportModify531) LeavesQtyId() uint16 {
	return 151
}

func (*ExecutionReportModify531) LeavesQtySinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportModify531) LeavesQtyInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.LeavesQtySinceVersion()
}

func (*ExecutionReportModify531) LeavesQtyDeprecated() uint16 {
	return 0
}

func (*ExecutionReportModify531) LeavesQtyMetaAttribute(meta int) string {
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

func (*ExecutionReportModify531) LeavesQtyMinValue() uint32 {
	return 0
}

func (*ExecutionReportModify531) LeavesQtyMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*ExecutionReportModify531) LeavesQtyNullValue() uint32 {
	return math.MaxUint32
}

func (*ExecutionReportModify531) MinQtyId() uint16 {
	return 110
}

func (*ExecutionReportModify531) MinQtySinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportModify531) MinQtyInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.MinQtySinceVersion()
}

func (*ExecutionReportModify531) MinQtyDeprecated() uint16 {
	return 0
}

func (*ExecutionReportModify531) MinQtyMetaAttribute(meta int) string {
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

func (*ExecutionReportModify531) MinQtyMinValue() uint32 {
	return 0
}

func (*ExecutionReportModify531) MinQtyMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*ExecutionReportModify531) MinQtyNullValue() uint32 {
	return 4294967295
}

func (*ExecutionReportModify531) DisplayQtyId() uint16 {
	return 1138
}

func (*ExecutionReportModify531) DisplayQtySinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportModify531) DisplayQtyInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.DisplayQtySinceVersion()
}

func (*ExecutionReportModify531) DisplayQtyDeprecated() uint16 {
	return 0
}

func (*ExecutionReportModify531) DisplayQtyMetaAttribute(meta int) string {
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

func (*ExecutionReportModify531) DisplayQtyMinValue() uint32 {
	return 0
}

func (*ExecutionReportModify531) DisplayQtyMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*ExecutionReportModify531) DisplayQtyNullValue() uint32 {
	return 4294967295
}

func (*ExecutionReportModify531) ExpireDateId() uint16 {
	return 432
}

func (*ExecutionReportModify531) ExpireDateSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportModify531) ExpireDateInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.ExpireDateSinceVersion()
}

func (*ExecutionReportModify531) ExpireDateDeprecated() uint16 {
	return 0
}

func (*ExecutionReportModify531) ExpireDateMetaAttribute(meta int) string {
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

func (*ExecutionReportModify531) ExpireDateMinValue() uint16 {
	return 0
}

func (*ExecutionReportModify531) ExpireDateMaxValue() uint16 {
	return math.MaxUint16 - 1
}

func (*ExecutionReportModify531) ExpireDateNullValue() uint16 {
	return 65535
}

func (*ExecutionReportModify531) DelayDurationId() uint16 {
	return 5904
}

func (*ExecutionReportModify531) DelayDurationSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportModify531) DelayDurationInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.DelayDurationSinceVersion()
}

func (*ExecutionReportModify531) DelayDurationDeprecated() uint16 {
	return 0
}

func (*ExecutionReportModify531) DelayDurationMetaAttribute(meta int) string {
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

func (*ExecutionReportModify531) DelayDurationMinValue() uint16 {
	return 0
}

func (*ExecutionReportModify531) DelayDurationMaxValue() uint16 {
	return math.MaxUint16 - 1
}

func (*ExecutionReportModify531) DelayDurationNullValue() uint16 {
	return 65535
}

func (*ExecutionReportModify531) OrdStatusId() uint16 {
	return 39
}

func (*ExecutionReportModify531) OrdStatusSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportModify531) OrdStatusInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.OrdStatusSinceVersion()
}

func (*ExecutionReportModify531) OrdStatusDeprecated() uint16 {
	return 0
}

func (*ExecutionReportModify531) OrdStatusMetaAttribute(meta int) string {
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

func (*ExecutionReportModify531) OrdStatusMinValue() byte {
	return byte(32)
}

func (*ExecutionReportModify531) OrdStatusMaxValue() byte {
	return byte(126)
}

func (*ExecutionReportModify531) OrdStatusNullValue() byte {
	return 0
}

func (*ExecutionReportModify531) ExecTypeId() uint16 {
	return 150
}

func (*ExecutionReportModify531) ExecTypeSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportModify531) ExecTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.ExecTypeSinceVersion()
}

func (*ExecutionReportModify531) ExecTypeDeprecated() uint16 {
	return 0
}

func (*ExecutionReportModify531) ExecTypeMetaAttribute(meta int) string {
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

func (*ExecutionReportModify531) ExecTypeMinValue() byte {
	return byte(32)
}

func (*ExecutionReportModify531) ExecTypeMaxValue() byte {
	return byte(126)
}

func (*ExecutionReportModify531) ExecTypeNullValue() byte {
	return 0
}

func (*ExecutionReportModify531) OrdTypeId() uint16 {
	return 40
}

func (*ExecutionReportModify531) OrdTypeSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportModify531) OrdTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.OrdTypeSinceVersion()
}

func (*ExecutionReportModify531) OrdTypeDeprecated() uint16 {
	return 0
}

func (*ExecutionReportModify531) OrdTypeMetaAttribute(meta int) string {
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

func (*ExecutionReportModify531) SideId() uint16 {
	return 54
}

func (*ExecutionReportModify531) SideSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportModify531) SideInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.SideSinceVersion()
}

func (*ExecutionReportModify531) SideDeprecated() uint16 {
	return 0
}

func (*ExecutionReportModify531) SideMetaAttribute(meta int) string {
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

func (*ExecutionReportModify531) TimeInForceId() uint16 {
	return 59
}

func (*ExecutionReportModify531) TimeInForceSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportModify531) TimeInForceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.TimeInForceSinceVersion()
}

func (*ExecutionReportModify531) TimeInForceDeprecated() uint16 {
	return 0
}

func (*ExecutionReportModify531) TimeInForceMetaAttribute(meta int) string {
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

func (*ExecutionReportModify531) ManualOrderIndicatorId() uint16 {
	return 1028
}

func (*ExecutionReportModify531) ManualOrderIndicatorSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportModify531) ManualOrderIndicatorInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.ManualOrderIndicatorSinceVersion()
}

func (*ExecutionReportModify531) ManualOrderIndicatorDeprecated() uint16 {
	return 0
}

func (*ExecutionReportModify531) ManualOrderIndicatorMetaAttribute(meta int) string {
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

func (*ExecutionReportModify531) PossRetransFlagId() uint16 {
	return 9765
}

func (*ExecutionReportModify531) PossRetransFlagSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportModify531) PossRetransFlagInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.PossRetransFlagSinceVersion()
}

func (*ExecutionReportModify531) PossRetransFlagDeprecated() uint16 {
	return 0
}

func (*ExecutionReportModify531) PossRetransFlagMetaAttribute(meta int) string {
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

func (*ExecutionReportModify531) SplitMsgId() uint16 {
	return 9553
}

func (*ExecutionReportModify531) SplitMsgSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportModify531) SplitMsgInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.SplitMsgSinceVersion()
}

func (*ExecutionReportModify531) SplitMsgDeprecated() uint16 {
	return 0
}

func (*ExecutionReportModify531) SplitMsgMetaAttribute(meta int) string {
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

func (*ExecutionReportModify531) CrossTypeId() uint16 {
	return 549
}

func (*ExecutionReportModify531) CrossTypeSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportModify531) CrossTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.CrossTypeSinceVersion()
}

func (*ExecutionReportModify531) CrossTypeDeprecated() uint16 {
	return 0
}

func (*ExecutionReportModify531) CrossTypeMetaAttribute(meta int) string {
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

func (*ExecutionReportModify531) CrossTypeMinValue() uint8 {
	return 0
}

func (*ExecutionReportModify531) CrossTypeMaxValue() uint8 {
	return math.MaxUint8 - 1
}

func (*ExecutionReportModify531) CrossTypeNullValue() uint8 {
	return 255
}

func (*ExecutionReportModify531) ExecInstId() uint16 {
	return 18
}

func (*ExecutionReportModify531) ExecInstSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportModify531) ExecInstInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.ExecInstSinceVersion()
}

func (*ExecutionReportModify531) ExecInstDeprecated() uint16 {
	return 0
}

func (*ExecutionReportModify531) ExecInstMetaAttribute(meta int) string {
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

func (*ExecutionReportModify531) ExecutionModeId() uint16 {
	return 5906
}

func (*ExecutionReportModify531) ExecutionModeSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportModify531) ExecutionModeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.ExecutionModeSinceVersion()
}

func (*ExecutionReportModify531) ExecutionModeDeprecated() uint16 {
	return 0
}

func (*ExecutionReportModify531) ExecutionModeMetaAttribute(meta int) string {
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

func (*ExecutionReportModify531) LiquidityFlagId() uint16 {
	return 9373
}

func (*ExecutionReportModify531) LiquidityFlagSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportModify531) LiquidityFlagInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.LiquidityFlagSinceVersion()
}

func (*ExecutionReportModify531) LiquidityFlagDeprecated() uint16 {
	return 0
}

func (*ExecutionReportModify531) LiquidityFlagMetaAttribute(meta int) string {
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

func (*ExecutionReportModify531) ManagedOrderId() uint16 {
	return 6881
}

func (*ExecutionReportModify531) ManagedOrderSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportModify531) ManagedOrderInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.ManagedOrderSinceVersion()
}

func (*ExecutionReportModify531) ManagedOrderDeprecated() uint16 {
	return 0
}

func (*ExecutionReportModify531) ManagedOrderMetaAttribute(meta int) string {
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

func (*ExecutionReportModify531) ShortSaleTypeId() uint16 {
	return 5409
}

func (*ExecutionReportModify531) ShortSaleTypeSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportModify531) ShortSaleTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.ShortSaleTypeSinceVersion()
}

func (*ExecutionReportModify531) ShortSaleTypeDeprecated() uint16 {
	return 0
}

func (*ExecutionReportModify531) ShortSaleTypeMetaAttribute(meta int) string {
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

func (*ExecutionReportModify531) DelayToTimeId() uint16 {
	return 7552
}

func (*ExecutionReportModify531) DelayToTimeSinceVersion() uint16 {
	return 4
}

func (e *ExecutionReportModify531) DelayToTimeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.DelayToTimeSinceVersion()
}

func (*ExecutionReportModify531) DelayToTimeDeprecated() uint16 {
	return 0
}

func (*ExecutionReportModify531) DelayToTimeMetaAttribute(meta int) string {
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

func (*ExecutionReportModify531) DelayToTimeMinValue() uint64 {
	return 0
}

func (*ExecutionReportModify531) DelayToTimeMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*ExecutionReportModify531) DelayToTimeNullValue() uint64 {
	return 18446744073709551615
}
