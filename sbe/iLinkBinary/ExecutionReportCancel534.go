// Generated SBE (Simple Binary Encoding) message codec

package ilinkbinary

import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
)

type ExecutionReportCancel534 struct {
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
	ExecRestatementReason ExecReasonEnum
	CrossType             uint8
	ExecInst              ExecInst
	ExecutionMode         ExecModeEnum
	LiquidityFlag         BooleanNULLEnum
	ManagedOrder          BooleanNULLEnum
	ShortSaleType         ShortSaleTypeEnum
	DelayToTime           uint64
}

func (e *ExecutionReportCancel534) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
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
	if err := e.ExecRestatementReason.Encode(_m, _w); err != nil {
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

func (e *ExecutionReportCancel534) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
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
	e.OrdStatus[0] = 52
	e.ExecType[0] = 52
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
	if e.ExecRestatementReasonInActingVersion(actingVersion) {
		if err := e.ExecRestatementReason.Decode(_m, _r, actingVersion); err != nil {
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

func (e *ExecutionReportCancel534) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
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
	if err := e.ExecRestatementReason.RangeCheck(actingVersion, schemaVersion); err != nil {
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

func ExecutionReportCancel534Init(e *ExecutionReportCancel534) {
	e.CrossID = 18446744073709551615
	e.HostCrossID = 18446744073709551615
	e.MinQty = 4294967295
	e.DisplayQty = 4294967295
	e.ExpireDate = 65535
	e.DelayDuration = 65535
	e.OrdStatus[0] = 52
	e.ExecType[0] = 52
	e.CrossType = 255
	e.DelayToTime = 18446744073709551615
	return
}

func (*ExecutionReportCancel534) SbeBlockLength() (blockLength uint16) {
	return 214
}

func (*ExecutionReportCancel534) SbeTemplateId() (templateId uint16) {
	return 534
}

func (*ExecutionReportCancel534) SbeSchemaId() (schemaId uint16) {
	return 8
}

func (*ExecutionReportCancel534) SbeSchemaVersion() (schemaVersion uint16) {
	return 5
}

func (*ExecutionReportCancel534) SbeSemanticType() (semanticType []byte) {
	return []byte("8")
}

func (*ExecutionReportCancel534) SeqNumId() uint16 {
	return 9726
}

func (*ExecutionReportCancel534) SeqNumSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportCancel534) SeqNumInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.SeqNumSinceVersion()
}

func (*ExecutionReportCancel534) SeqNumDeprecated() uint16 {
	return 0
}

func (*ExecutionReportCancel534) SeqNumMetaAttribute(meta int) string {
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

func (*ExecutionReportCancel534) SeqNumMinValue() uint32 {
	return 0
}

func (*ExecutionReportCancel534) SeqNumMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*ExecutionReportCancel534) SeqNumNullValue() uint32 {
	return math.MaxUint32
}

func (*ExecutionReportCancel534) UUIDId() uint16 {
	return 39001
}

func (*ExecutionReportCancel534) UUIDSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportCancel534) UUIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.UUIDSinceVersion()
}

func (*ExecutionReportCancel534) UUIDDeprecated() uint16 {
	return 0
}

func (*ExecutionReportCancel534) UUIDMetaAttribute(meta int) string {
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

func (*ExecutionReportCancel534) UUIDMinValue() uint64 {
	return 0
}

func (*ExecutionReportCancel534) UUIDMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*ExecutionReportCancel534) UUIDNullValue() uint64 {
	return math.MaxUint64
}

func (*ExecutionReportCancel534) ExecIDId() uint16 {
	return 17
}

func (*ExecutionReportCancel534) ExecIDSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportCancel534) ExecIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.ExecIDSinceVersion()
}

func (*ExecutionReportCancel534) ExecIDDeprecated() uint16 {
	return 0
}

func (*ExecutionReportCancel534) ExecIDMetaAttribute(meta int) string {
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

func (*ExecutionReportCancel534) ExecIDMinValue() byte {
	return byte(32)
}

func (*ExecutionReportCancel534) ExecIDMaxValue() byte {
	return byte(126)
}

func (*ExecutionReportCancel534) ExecIDNullValue() byte {
	return 0
}

func (e *ExecutionReportCancel534) ExecIDCharacterEncoding() string {
	return "US-ASCII"
}

func (*ExecutionReportCancel534) SenderIDId() uint16 {
	return 5392
}

func (*ExecutionReportCancel534) SenderIDSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportCancel534) SenderIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.SenderIDSinceVersion()
}

func (*ExecutionReportCancel534) SenderIDDeprecated() uint16 {
	return 0
}

func (*ExecutionReportCancel534) SenderIDMetaAttribute(meta int) string {
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

func (*ExecutionReportCancel534) SenderIDMinValue() byte {
	return byte(32)
}

func (*ExecutionReportCancel534) SenderIDMaxValue() byte {
	return byte(126)
}

func (*ExecutionReportCancel534) SenderIDNullValue() byte {
	return 0
}

func (e *ExecutionReportCancel534) SenderIDCharacterEncoding() string {
	return "US-ASCII"
}

func (*ExecutionReportCancel534) ClOrdIDId() uint16 {
	return 11
}

func (*ExecutionReportCancel534) ClOrdIDSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportCancel534) ClOrdIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.ClOrdIDSinceVersion()
}

func (*ExecutionReportCancel534) ClOrdIDDeprecated() uint16 {
	return 0
}

func (*ExecutionReportCancel534) ClOrdIDMetaAttribute(meta int) string {
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

func (*ExecutionReportCancel534) ClOrdIDMinValue() byte {
	return byte(32)
}

func (*ExecutionReportCancel534) ClOrdIDMaxValue() byte {
	return byte(126)
}

func (*ExecutionReportCancel534) ClOrdIDNullValue() byte {
	return 0
}

func (e *ExecutionReportCancel534) ClOrdIDCharacterEncoding() string {
	return "US-ASCII"
}

func (*ExecutionReportCancel534) PartyDetailsListReqIDId() uint16 {
	return 1505
}

func (*ExecutionReportCancel534) PartyDetailsListReqIDSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportCancel534) PartyDetailsListReqIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.PartyDetailsListReqIDSinceVersion()
}

func (*ExecutionReportCancel534) PartyDetailsListReqIDDeprecated() uint16 {
	return 0
}

func (*ExecutionReportCancel534) PartyDetailsListReqIDMetaAttribute(meta int) string {
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

func (*ExecutionReportCancel534) PartyDetailsListReqIDMinValue() uint64 {
	return 0
}

func (*ExecutionReportCancel534) PartyDetailsListReqIDMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*ExecutionReportCancel534) PartyDetailsListReqIDNullValue() uint64 {
	return math.MaxUint64
}

func (*ExecutionReportCancel534) OrderIDId() uint16 {
	return 37
}

func (*ExecutionReportCancel534) OrderIDSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportCancel534) OrderIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.OrderIDSinceVersion()
}

func (*ExecutionReportCancel534) OrderIDDeprecated() uint16 {
	return 0
}

func (*ExecutionReportCancel534) OrderIDMetaAttribute(meta int) string {
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

func (*ExecutionReportCancel534) OrderIDMinValue() uint64 {
	return 0
}

func (*ExecutionReportCancel534) OrderIDMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*ExecutionReportCancel534) OrderIDNullValue() uint64 {
	return math.MaxUint64
}

func (*ExecutionReportCancel534) PriceId() uint16 {
	return 44
}

func (*ExecutionReportCancel534) PriceSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportCancel534) PriceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.PriceSinceVersion()
}

func (*ExecutionReportCancel534) PriceDeprecated() uint16 {
	return 0
}

func (*ExecutionReportCancel534) PriceMetaAttribute(meta int) string {
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

func (*ExecutionReportCancel534) StopPxId() uint16 {
	return 99
}

func (*ExecutionReportCancel534) StopPxSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportCancel534) StopPxInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.StopPxSinceVersion()
}

func (*ExecutionReportCancel534) StopPxDeprecated() uint16 {
	return 0
}

func (*ExecutionReportCancel534) StopPxMetaAttribute(meta int) string {
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

func (*ExecutionReportCancel534) TransactTimeId() uint16 {
	return 60
}

func (*ExecutionReportCancel534) TransactTimeSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportCancel534) TransactTimeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.TransactTimeSinceVersion()
}

func (*ExecutionReportCancel534) TransactTimeDeprecated() uint16 {
	return 0
}

func (*ExecutionReportCancel534) TransactTimeMetaAttribute(meta int) string {
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

func (*ExecutionReportCancel534) TransactTimeMinValue() uint64 {
	return 0
}

func (*ExecutionReportCancel534) TransactTimeMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*ExecutionReportCancel534) TransactTimeNullValue() uint64 {
	return math.MaxUint64
}

func (*ExecutionReportCancel534) SendingTimeEpochId() uint16 {
	return 5297
}

func (*ExecutionReportCancel534) SendingTimeEpochSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportCancel534) SendingTimeEpochInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.SendingTimeEpochSinceVersion()
}

func (*ExecutionReportCancel534) SendingTimeEpochDeprecated() uint16 {
	return 0
}

func (*ExecutionReportCancel534) SendingTimeEpochMetaAttribute(meta int) string {
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

func (*ExecutionReportCancel534) SendingTimeEpochMinValue() uint64 {
	return 0
}

func (*ExecutionReportCancel534) SendingTimeEpochMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*ExecutionReportCancel534) SendingTimeEpochNullValue() uint64 {
	return math.MaxUint64
}

func (*ExecutionReportCancel534) OrderRequestIDId() uint16 {
	return 2422
}

func (*ExecutionReportCancel534) OrderRequestIDSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportCancel534) OrderRequestIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.OrderRequestIDSinceVersion()
}

func (*ExecutionReportCancel534) OrderRequestIDDeprecated() uint16 {
	return 0
}

func (*ExecutionReportCancel534) OrderRequestIDMetaAttribute(meta int) string {
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

func (*ExecutionReportCancel534) OrderRequestIDMinValue() uint64 {
	return 0
}

func (*ExecutionReportCancel534) OrderRequestIDMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*ExecutionReportCancel534) OrderRequestIDNullValue() uint64 {
	return math.MaxUint64
}

func (*ExecutionReportCancel534) CrossIDId() uint16 {
	return 548
}

func (*ExecutionReportCancel534) CrossIDSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportCancel534) CrossIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.CrossIDSinceVersion()
}

func (*ExecutionReportCancel534) CrossIDDeprecated() uint16 {
	return 0
}

func (*ExecutionReportCancel534) CrossIDMetaAttribute(meta int) string {
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

func (*ExecutionReportCancel534) CrossIDMinValue() uint64 {
	return 0
}

func (*ExecutionReportCancel534) CrossIDMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*ExecutionReportCancel534) CrossIDNullValue() uint64 {
	return 18446744073709551615
}

func (*ExecutionReportCancel534) HostCrossIDId() uint16 {
	return 961
}

func (*ExecutionReportCancel534) HostCrossIDSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportCancel534) HostCrossIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.HostCrossIDSinceVersion()
}

func (*ExecutionReportCancel534) HostCrossIDDeprecated() uint16 {
	return 0
}

func (*ExecutionReportCancel534) HostCrossIDMetaAttribute(meta int) string {
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

func (*ExecutionReportCancel534) HostCrossIDMinValue() uint64 {
	return 0
}

func (*ExecutionReportCancel534) HostCrossIDMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*ExecutionReportCancel534) HostCrossIDNullValue() uint64 {
	return 18446744073709551615
}

func (*ExecutionReportCancel534) LocationId() uint16 {
	return 9537
}

func (*ExecutionReportCancel534) LocationSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportCancel534) LocationInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.LocationSinceVersion()
}

func (*ExecutionReportCancel534) LocationDeprecated() uint16 {
	return 0
}

func (*ExecutionReportCancel534) LocationMetaAttribute(meta int) string {
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

func (*ExecutionReportCancel534) LocationMinValue() byte {
	return byte(32)
}

func (*ExecutionReportCancel534) LocationMaxValue() byte {
	return byte(126)
}

func (*ExecutionReportCancel534) LocationNullValue() byte {
	return 0
}

func (e *ExecutionReportCancel534) LocationCharacterEncoding() string {
	return "US-ASCII"
}

func (*ExecutionReportCancel534) SecurityIDId() uint16 {
	return 48
}

func (*ExecutionReportCancel534) SecurityIDSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportCancel534) SecurityIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.SecurityIDSinceVersion()
}

func (*ExecutionReportCancel534) SecurityIDDeprecated() uint16 {
	return 0
}

func (*ExecutionReportCancel534) SecurityIDMetaAttribute(meta int) string {
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

func (*ExecutionReportCancel534) SecurityIDMinValue() int32 {
	return math.MinInt32 + 1
}

func (*ExecutionReportCancel534) SecurityIDMaxValue() int32 {
	return math.MaxInt32
}

func (*ExecutionReportCancel534) SecurityIDNullValue() int32 {
	return math.MinInt32
}

func (*ExecutionReportCancel534) OrderQtyId() uint16 {
	return 38
}

func (*ExecutionReportCancel534) OrderQtySinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportCancel534) OrderQtyInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.OrderQtySinceVersion()
}

func (*ExecutionReportCancel534) OrderQtyDeprecated() uint16 {
	return 0
}

func (*ExecutionReportCancel534) OrderQtyMetaAttribute(meta int) string {
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

func (*ExecutionReportCancel534) OrderQtyMinValue() uint32 {
	return 0
}

func (*ExecutionReportCancel534) OrderQtyMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*ExecutionReportCancel534) OrderQtyNullValue() uint32 {
	return math.MaxUint32
}

func (*ExecutionReportCancel534) CumQtyId() uint16 {
	return 14
}

func (*ExecutionReportCancel534) CumQtySinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportCancel534) CumQtyInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.CumQtySinceVersion()
}

func (*ExecutionReportCancel534) CumQtyDeprecated() uint16 {
	return 0
}

func (*ExecutionReportCancel534) CumQtyMetaAttribute(meta int) string {
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

func (*ExecutionReportCancel534) CumQtyMinValue() uint32 {
	return 0
}

func (*ExecutionReportCancel534) CumQtyMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*ExecutionReportCancel534) CumQtyNullValue() uint32 {
	return math.MaxUint32
}

func (*ExecutionReportCancel534) MinQtyId() uint16 {
	return 110
}

func (*ExecutionReportCancel534) MinQtySinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportCancel534) MinQtyInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.MinQtySinceVersion()
}

func (*ExecutionReportCancel534) MinQtyDeprecated() uint16 {
	return 0
}

func (*ExecutionReportCancel534) MinQtyMetaAttribute(meta int) string {
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

func (*ExecutionReportCancel534) MinQtyMinValue() uint32 {
	return 0
}

func (*ExecutionReportCancel534) MinQtyMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*ExecutionReportCancel534) MinQtyNullValue() uint32 {
	return 4294967295
}

func (*ExecutionReportCancel534) DisplayQtyId() uint16 {
	return 1138
}

func (*ExecutionReportCancel534) DisplayQtySinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportCancel534) DisplayQtyInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.DisplayQtySinceVersion()
}

func (*ExecutionReportCancel534) DisplayQtyDeprecated() uint16 {
	return 0
}

func (*ExecutionReportCancel534) DisplayQtyMetaAttribute(meta int) string {
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

func (*ExecutionReportCancel534) DisplayQtyMinValue() uint32 {
	return 0
}

func (*ExecutionReportCancel534) DisplayQtyMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*ExecutionReportCancel534) DisplayQtyNullValue() uint32 {
	return 4294967295
}

func (*ExecutionReportCancel534) ExpireDateId() uint16 {
	return 432
}

func (*ExecutionReportCancel534) ExpireDateSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportCancel534) ExpireDateInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.ExpireDateSinceVersion()
}

func (*ExecutionReportCancel534) ExpireDateDeprecated() uint16 {
	return 0
}

func (*ExecutionReportCancel534) ExpireDateMetaAttribute(meta int) string {
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

func (*ExecutionReportCancel534) ExpireDateMinValue() uint16 {
	return 0
}

func (*ExecutionReportCancel534) ExpireDateMaxValue() uint16 {
	return math.MaxUint16 - 1
}

func (*ExecutionReportCancel534) ExpireDateNullValue() uint16 {
	return 65535
}

func (*ExecutionReportCancel534) DelayDurationId() uint16 {
	return 5904
}

func (*ExecutionReportCancel534) DelayDurationSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportCancel534) DelayDurationInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.DelayDurationSinceVersion()
}

func (*ExecutionReportCancel534) DelayDurationDeprecated() uint16 {
	return 0
}

func (*ExecutionReportCancel534) DelayDurationMetaAttribute(meta int) string {
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

func (*ExecutionReportCancel534) DelayDurationMinValue() uint16 {
	return 0
}

func (*ExecutionReportCancel534) DelayDurationMaxValue() uint16 {
	return math.MaxUint16 - 1
}

func (*ExecutionReportCancel534) DelayDurationNullValue() uint16 {
	return 65535
}

func (*ExecutionReportCancel534) OrdStatusId() uint16 {
	return 39
}

func (*ExecutionReportCancel534) OrdStatusSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportCancel534) OrdStatusInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.OrdStatusSinceVersion()
}

func (*ExecutionReportCancel534) OrdStatusDeprecated() uint16 {
	return 0
}

func (*ExecutionReportCancel534) OrdStatusMetaAttribute(meta int) string {
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

func (*ExecutionReportCancel534) OrdStatusMinValue() byte {
	return byte(32)
}

func (*ExecutionReportCancel534) OrdStatusMaxValue() byte {
	return byte(126)
}

func (*ExecutionReportCancel534) OrdStatusNullValue() byte {
	return 0
}

func (*ExecutionReportCancel534) ExecTypeId() uint16 {
	return 150
}

func (*ExecutionReportCancel534) ExecTypeSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportCancel534) ExecTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.ExecTypeSinceVersion()
}

func (*ExecutionReportCancel534) ExecTypeDeprecated() uint16 {
	return 0
}

func (*ExecutionReportCancel534) ExecTypeMetaAttribute(meta int) string {
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

func (*ExecutionReportCancel534) ExecTypeMinValue() byte {
	return byte(32)
}

func (*ExecutionReportCancel534) ExecTypeMaxValue() byte {
	return byte(126)
}

func (*ExecutionReportCancel534) ExecTypeNullValue() byte {
	return 0
}

func (*ExecutionReportCancel534) OrdTypeId() uint16 {
	return 40
}

func (*ExecutionReportCancel534) OrdTypeSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportCancel534) OrdTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.OrdTypeSinceVersion()
}

func (*ExecutionReportCancel534) OrdTypeDeprecated() uint16 {
	return 0
}

func (*ExecutionReportCancel534) OrdTypeMetaAttribute(meta int) string {
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

func (*ExecutionReportCancel534) SideId() uint16 {
	return 54
}

func (*ExecutionReportCancel534) SideSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportCancel534) SideInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.SideSinceVersion()
}

func (*ExecutionReportCancel534) SideDeprecated() uint16 {
	return 0
}

func (*ExecutionReportCancel534) SideMetaAttribute(meta int) string {
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

func (*ExecutionReportCancel534) TimeInForceId() uint16 {
	return 59
}

func (*ExecutionReportCancel534) TimeInForceSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportCancel534) TimeInForceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.TimeInForceSinceVersion()
}

func (*ExecutionReportCancel534) TimeInForceDeprecated() uint16 {
	return 0
}

func (*ExecutionReportCancel534) TimeInForceMetaAttribute(meta int) string {
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

func (*ExecutionReportCancel534) ManualOrderIndicatorId() uint16 {
	return 1028
}

func (*ExecutionReportCancel534) ManualOrderIndicatorSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportCancel534) ManualOrderIndicatorInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.ManualOrderIndicatorSinceVersion()
}

func (*ExecutionReportCancel534) ManualOrderIndicatorDeprecated() uint16 {
	return 0
}

func (*ExecutionReportCancel534) ManualOrderIndicatorMetaAttribute(meta int) string {
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

func (*ExecutionReportCancel534) PossRetransFlagId() uint16 {
	return 9765
}

func (*ExecutionReportCancel534) PossRetransFlagSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportCancel534) PossRetransFlagInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.PossRetransFlagSinceVersion()
}

func (*ExecutionReportCancel534) PossRetransFlagDeprecated() uint16 {
	return 0
}

func (*ExecutionReportCancel534) PossRetransFlagMetaAttribute(meta int) string {
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

func (*ExecutionReportCancel534) SplitMsgId() uint16 {
	return 9553
}

func (*ExecutionReportCancel534) SplitMsgSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportCancel534) SplitMsgInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.SplitMsgSinceVersion()
}

func (*ExecutionReportCancel534) SplitMsgDeprecated() uint16 {
	return 0
}

func (*ExecutionReportCancel534) SplitMsgMetaAttribute(meta int) string {
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

func (*ExecutionReportCancel534) ExecRestatementReasonId() uint16 {
	return 378
}

func (*ExecutionReportCancel534) ExecRestatementReasonSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportCancel534) ExecRestatementReasonInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.ExecRestatementReasonSinceVersion()
}

func (*ExecutionReportCancel534) ExecRestatementReasonDeprecated() uint16 {
	return 0
}

func (*ExecutionReportCancel534) ExecRestatementReasonMetaAttribute(meta int) string {
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

func (*ExecutionReportCancel534) CrossTypeId() uint16 {
	return 549
}

func (*ExecutionReportCancel534) CrossTypeSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportCancel534) CrossTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.CrossTypeSinceVersion()
}

func (*ExecutionReportCancel534) CrossTypeDeprecated() uint16 {
	return 0
}

func (*ExecutionReportCancel534) CrossTypeMetaAttribute(meta int) string {
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

func (*ExecutionReportCancel534) CrossTypeMinValue() uint8 {
	return 0
}

func (*ExecutionReportCancel534) CrossTypeMaxValue() uint8 {
	return math.MaxUint8 - 1
}

func (*ExecutionReportCancel534) CrossTypeNullValue() uint8 {
	return 255
}

func (*ExecutionReportCancel534) ExecInstId() uint16 {
	return 18
}

func (*ExecutionReportCancel534) ExecInstSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportCancel534) ExecInstInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.ExecInstSinceVersion()
}

func (*ExecutionReportCancel534) ExecInstDeprecated() uint16 {
	return 0
}

func (*ExecutionReportCancel534) ExecInstMetaAttribute(meta int) string {
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

func (*ExecutionReportCancel534) ExecutionModeId() uint16 {
	return 5906
}

func (*ExecutionReportCancel534) ExecutionModeSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportCancel534) ExecutionModeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.ExecutionModeSinceVersion()
}

func (*ExecutionReportCancel534) ExecutionModeDeprecated() uint16 {
	return 0
}

func (*ExecutionReportCancel534) ExecutionModeMetaAttribute(meta int) string {
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

func (*ExecutionReportCancel534) LiquidityFlagId() uint16 {
	return 9373
}

func (*ExecutionReportCancel534) LiquidityFlagSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportCancel534) LiquidityFlagInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.LiquidityFlagSinceVersion()
}

func (*ExecutionReportCancel534) LiquidityFlagDeprecated() uint16 {
	return 0
}

func (*ExecutionReportCancel534) LiquidityFlagMetaAttribute(meta int) string {
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

func (*ExecutionReportCancel534) ManagedOrderId() uint16 {
	return 6881
}

func (*ExecutionReportCancel534) ManagedOrderSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportCancel534) ManagedOrderInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.ManagedOrderSinceVersion()
}

func (*ExecutionReportCancel534) ManagedOrderDeprecated() uint16 {
	return 0
}

func (*ExecutionReportCancel534) ManagedOrderMetaAttribute(meta int) string {
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

func (*ExecutionReportCancel534) ShortSaleTypeId() uint16 {
	return 5409
}

func (*ExecutionReportCancel534) ShortSaleTypeSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportCancel534) ShortSaleTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.ShortSaleTypeSinceVersion()
}

func (*ExecutionReportCancel534) ShortSaleTypeDeprecated() uint16 {
	return 0
}

func (*ExecutionReportCancel534) ShortSaleTypeMetaAttribute(meta int) string {
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

func (*ExecutionReportCancel534) DelayToTimeId() uint16 {
	return 7552
}

func (*ExecutionReportCancel534) DelayToTimeSinceVersion() uint16 {
	return 4
}

func (e *ExecutionReportCancel534) DelayToTimeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.DelayToTimeSinceVersion()
}

func (*ExecutionReportCancel534) DelayToTimeDeprecated() uint16 {
	return 0
}

func (*ExecutionReportCancel534) DelayToTimeMetaAttribute(meta int) string {
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

func (*ExecutionReportCancel534) DelayToTimeMinValue() uint64 {
	return 0
}

func (*ExecutionReportCancel534) DelayToTimeMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*ExecutionReportCancel534) DelayToTimeNullValue() uint64 {
	return 18446744073709551615
}
