// Generated SBE (Simple Binary Encoding) message codec

package ilinkbinary

import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
)

type ExecutionReportNew522 struct {
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

func (e *ExecutionReportNew522) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
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

func (e *ExecutionReportNew522) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
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
	e.OrdStatus[0] = 48
	e.ExecType[0] = 48
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

func (e *ExecutionReportNew522) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
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

func ExecutionReportNew522Init(e *ExecutionReportNew522) {
	e.CrossID = 18446744073709551615
	e.HostCrossID = 18446744073709551615
	e.MinQty = 4294967295
	e.DisplayQty = 4294967295
	e.ExpireDate = 65535
	e.DelayDuration = 65535
	e.OrdStatus[0] = 48
	e.ExecType[0] = 48
	e.CrossType = 255
	e.DelayToTime = 18446744073709551615
	return
}

func (*ExecutionReportNew522) SbeBlockLength() (blockLength uint16) {
	return 209
}

func (*ExecutionReportNew522) SbeTemplateId() (templateId uint16) {
	return 522
}

func (*ExecutionReportNew522) SbeSchemaId() (schemaId uint16) {
	return 8
}

func (*ExecutionReportNew522) SbeSchemaVersion() (schemaVersion uint16) {
	return 5
}

func (*ExecutionReportNew522) SbeSemanticType() (semanticType []byte) {
	return []byte("8")
}

func (*ExecutionReportNew522) SeqNumId() uint16 {
	return 9726
}

func (*ExecutionReportNew522) SeqNumSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportNew522) SeqNumInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.SeqNumSinceVersion()
}

func (*ExecutionReportNew522) SeqNumDeprecated() uint16 {
	return 0
}

func (*ExecutionReportNew522) SeqNumMetaAttribute(meta int) string {
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

func (*ExecutionReportNew522) SeqNumMinValue() uint32 {
	return 0
}

func (*ExecutionReportNew522) SeqNumMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*ExecutionReportNew522) SeqNumNullValue() uint32 {
	return math.MaxUint32
}

func (*ExecutionReportNew522) UUIDId() uint16 {
	return 39001
}

func (*ExecutionReportNew522) UUIDSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportNew522) UUIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.UUIDSinceVersion()
}

func (*ExecutionReportNew522) UUIDDeprecated() uint16 {
	return 0
}

func (*ExecutionReportNew522) UUIDMetaAttribute(meta int) string {
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

func (*ExecutionReportNew522) UUIDMinValue() uint64 {
	return 0
}

func (*ExecutionReportNew522) UUIDMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*ExecutionReportNew522) UUIDNullValue() uint64 {
	return math.MaxUint64
}

func (*ExecutionReportNew522) ExecIDId() uint16 {
	return 17
}

func (*ExecutionReportNew522) ExecIDSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportNew522) ExecIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.ExecIDSinceVersion()
}

func (*ExecutionReportNew522) ExecIDDeprecated() uint16 {
	return 0
}

func (*ExecutionReportNew522) ExecIDMetaAttribute(meta int) string {
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

func (*ExecutionReportNew522) ExecIDMinValue() byte {
	return byte(32)
}

func (*ExecutionReportNew522) ExecIDMaxValue() byte {
	return byte(126)
}

func (*ExecutionReportNew522) ExecIDNullValue() byte {
	return 0
}

func (e *ExecutionReportNew522) ExecIDCharacterEncoding() string {
	return "US-ASCII"
}

func (*ExecutionReportNew522) SenderIDId() uint16 {
	return 5392
}

func (*ExecutionReportNew522) SenderIDSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportNew522) SenderIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.SenderIDSinceVersion()
}

func (*ExecutionReportNew522) SenderIDDeprecated() uint16 {
	return 0
}

func (*ExecutionReportNew522) SenderIDMetaAttribute(meta int) string {
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

func (*ExecutionReportNew522) SenderIDMinValue() byte {
	return byte(32)
}

func (*ExecutionReportNew522) SenderIDMaxValue() byte {
	return byte(126)
}

func (*ExecutionReportNew522) SenderIDNullValue() byte {
	return 0
}

func (e *ExecutionReportNew522) SenderIDCharacterEncoding() string {
	return "US-ASCII"
}

func (*ExecutionReportNew522) ClOrdIDId() uint16 {
	return 11
}

func (*ExecutionReportNew522) ClOrdIDSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportNew522) ClOrdIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.ClOrdIDSinceVersion()
}

func (*ExecutionReportNew522) ClOrdIDDeprecated() uint16 {
	return 0
}

func (*ExecutionReportNew522) ClOrdIDMetaAttribute(meta int) string {
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

func (*ExecutionReportNew522) ClOrdIDMinValue() byte {
	return byte(32)
}

func (*ExecutionReportNew522) ClOrdIDMaxValue() byte {
	return byte(126)
}

func (*ExecutionReportNew522) ClOrdIDNullValue() byte {
	return 0
}

func (e *ExecutionReportNew522) ClOrdIDCharacterEncoding() string {
	return "US-ASCII"
}

func (*ExecutionReportNew522) PartyDetailsListReqIDId() uint16 {
	return 1505
}

func (*ExecutionReportNew522) PartyDetailsListReqIDSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportNew522) PartyDetailsListReqIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.PartyDetailsListReqIDSinceVersion()
}

func (*ExecutionReportNew522) PartyDetailsListReqIDDeprecated() uint16 {
	return 0
}

func (*ExecutionReportNew522) PartyDetailsListReqIDMetaAttribute(meta int) string {
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

func (*ExecutionReportNew522) PartyDetailsListReqIDMinValue() uint64 {
	return 0
}

func (*ExecutionReportNew522) PartyDetailsListReqIDMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*ExecutionReportNew522) PartyDetailsListReqIDNullValue() uint64 {
	return math.MaxUint64
}

func (*ExecutionReportNew522) OrderIDId() uint16 {
	return 37
}

func (*ExecutionReportNew522) OrderIDSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportNew522) OrderIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.OrderIDSinceVersion()
}

func (*ExecutionReportNew522) OrderIDDeprecated() uint16 {
	return 0
}

func (*ExecutionReportNew522) OrderIDMetaAttribute(meta int) string {
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

func (*ExecutionReportNew522) OrderIDMinValue() uint64 {
	return 0
}

func (*ExecutionReportNew522) OrderIDMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*ExecutionReportNew522) OrderIDNullValue() uint64 {
	return math.MaxUint64
}

func (*ExecutionReportNew522) PriceId() uint16 {
	return 44
}

func (*ExecutionReportNew522) PriceSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportNew522) PriceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.PriceSinceVersion()
}

func (*ExecutionReportNew522) PriceDeprecated() uint16 {
	return 0
}

func (*ExecutionReportNew522) PriceMetaAttribute(meta int) string {
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

func (*ExecutionReportNew522) StopPxId() uint16 {
	return 99
}

func (*ExecutionReportNew522) StopPxSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportNew522) StopPxInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.StopPxSinceVersion()
}

func (*ExecutionReportNew522) StopPxDeprecated() uint16 {
	return 0
}

func (*ExecutionReportNew522) StopPxMetaAttribute(meta int) string {
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

func (*ExecutionReportNew522) TransactTimeId() uint16 {
	return 60
}

func (*ExecutionReportNew522) TransactTimeSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportNew522) TransactTimeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.TransactTimeSinceVersion()
}

func (*ExecutionReportNew522) TransactTimeDeprecated() uint16 {
	return 0
}

func (*ExecutionReportNew522) TransactTimeMetaAttribute(meta int) string {
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

func (*ExecutionReportNew522) TransactTimeMinValue() uint64 {
	return 0
}

func (*ExecutionReportNew522) TransactTimeMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*ExecutionReportNew522) TransactTimeNullValue() uint64 {
	return math.MaxUint64
}

func (*ExecutionReportNew522) SendingTimeEpochId() uint16 {
	return 5297
}

func (*ExecutionReportNew522) SendingTimeEpochSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportNew522) SendingTimeEpochInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.SendingTimeEpochSinceVersion()
}

func (*ExecutionReportNew522) SendingTimeEpochDeprecated() uint16 {
	return 0
}

func (*ExecutionReportNew522) SendingTimeEpochMetaAttribute(meta int) string {
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

func (*ExecutionReportNew522) SendingTimeEpochMinValue() uint64 {
	return 0
}

func (*ExecutionReportNew522) SendingTimeEpochMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*ExecutionReportNew522) SendingTimeEpochNullValue() uint64 {
	return math.MaxUint64
}

func (*ExecutionReportNew522) OrderRequestIDId() uint16 {
	return 2422
}

func (*ExecutionReportNew522) OrderRequestIDSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportNew522) OrderRequestIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.OrderRequestIDSinceVersion()
}

func (*ExecutionReportNew522) OrderRequestIDDeprecated() uint16 {
	return 0
}

func (*ExecutionReportNew522) OrderRequestIDMetaAttribute(meta int) string {
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

func (*ExecutionReportNew522) OrderRequestIDMinValue() uint64 {
	return 0
}

func (*ExecutionReportNew522) OrderRequestIDMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*ExecutionReportNew522) OrderRequestIDNullValue() uint64 {
	return math.MaxUint64
}

func (*ExecutionReportNew522) CrossIDId() uint16 {
	return 548
}

func (*ExecutionReportNew522) CrossIDSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportNew522) CrossIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.CrossIDSinceVersion()
}

func (*ExecutionReportNew522) CrossIDDeprecated() uint16 {
	return 0
}

func (*ExecutionReportNew522) CrossIDMetaAttribute(meta int) string {
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

func (*ExecutionReportNew522) CrossIDMinValue() uint64 {
	return 0
}

func (*ExecutionReportNew522) CrossIDMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*ExecutionReportNew522) CrossIDNullValue() uint64 {
	return 18446744073709551615
}

func (*ExecutionReportNew522) HostCrossIDId() uint16 {
	return 961
}

func (*ExecutionReportNew522) HostCrossIDSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportNew522) HostCrossIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.HostCrossIDSinceVersion()
}

func (*ExecutionReportNew522) HostCrossIDDeprecated() uint16 {
	return 0
}

func (*ExecutionReportNew522) HostCrossIDMetaAttribute(meta int) string {
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

func (*ExecutionReportNew522) HostCrossIDMinValue() uint64 {
	return 0
}

func (*ExecutionReportNew522) HostCrossIDMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*ExecutionReportNew522) HostCrossIDNullValue() uint64 {
	return 18446744073709551615
}

func (*ExecutionReportNew522) LocationId() uint16 {
	return 9537
}

func (*ExecutionReportNew522) LocationSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportNew522) LocationInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.LocationSinceVersion()
}

func (*ExecutionReportNew522) LocationDeprecated() uint16 {
	return 0
}

func (*ExecutionReportNew522) LocationMetaAttribute(meta int) string {
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

func (*ExecutionReportNew522) LocationMinValue() byte {
	return byte(32)
}

func (*ExecutionReportNew522) LocationMaxValue() byte {
	return byte(126)
}

func (*ExecutionReportNew522) LocationNullValue() byte {
	return 0
}

func (e *ExecutionReportNew522) LocationCharacterEncoding() string {
	return "US-ASCII"
}

func (*ExecutionReportNew522) SecurityIDId() uint16 {
	return 48
}

func (*ExecutionReportNew522) SecurityIDSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportNew522) SecurityIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.SecurityIDSinceVersion()
}

func (*ExecutionReportNew522) SecurityIDDeprecated() uint16 {
	return 0
}

func (*ExecutionReportNew522) SecurityIDMetaAttribute(meta int) string {
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

func (*ExecutionReportNew522) SecurityIDMinValue() int32 {
	return math.MinInt32 + 1
}

func (*ExecutionReportNew522) SecurityIDMaxValue() int32 {
	return math.MaxInt32
}

func (*ExecutionReportNew522) SecurityIDNullValue() int32 {
	return math.MinInt32
}

func (*ExecutionReportNew522) OrderQtyId() uint16 {
	return 38
}

func (*ExecutionReportNew522) OrderQtySinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportNew522) OrderQtyInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.OrderQtySinceVersion()
}

func (*ExecutionReportNew522) OrderQtyDeprecated() uint16 {
	return 0
}

func (*ExecutionReportNew522) OrderQtyMetaAttribute(meta int) string {
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

func (*ExecutionReportNew522) OrderQtyMinValue() uint32 {
	return 0
}

func (*ExecutionReportNew522) OrderQtyMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*ExecutionReportNew522) OrderQtyNullValue() uint32 {
	return math.MaxUint32
}

func (*ExecutionReportNew522) MinQtyId() uint16 {
	return 110
}

func (*ExecutionReportNew522) MinQtySinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportNew522) MinQtyInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.MinQtySinceVersion()
}

func (*ExecutionReportNew522) MinQtyDeprecated() uint16 {
	return 0
}

func (*ExecutionReportNew522) MinQtyMetaAttribute(meta int) string {
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

func (*ExecutionReportNew522) MinQtyMinValue() uint32 {
	return 0
}

func (*ExecutionReportNew522) MinQtyMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*ExecutionReportNew522) MinQtyNullValue() uint32 {
	return 4294967295
}

func (*ExecutionReportNew522) DisplayQtyId() uint16 {
	return 1138
}

func (*ExecutionReportNew522) DisplayQtySinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportNew522) DisplayQtyInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.DisplayQtySinceVersion()
}

func (*ExecutionReportNew522) DisplayQtyDeprecated() uint16 {
	return 0
}

func (*ExecutionReportNew522) DisplayQtyMetaAttribute(meta int) string {
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

func (*ExecutionReportNew522) DisplayQtyMinValue() uint32 {
	return 0
}

func (*ExecutionReportNew522) DisplayQtyMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*ExecutionReportNew522) DisplayQtyNullValue() uint32 {
	return 4294967295
}

func (*ExecutionReportNew522) ExpireDateId() uint16 {
	return 432
}

func (*ExecutionReportNew522) ExpireDateSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportNew522) ExpireDateInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.ExpireDateSinceVersion()
}

func (*ExecutionReportNew522) ExpireDateDeprecated() uint16 {
	return 0
}

func (*ExecutionReportNew522) ExpireDateMetaAttribute(meta int) string {
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

func (*ExecutionReportNew522) ExpireDateMinValue() uint16 {
	return 0
}

func (*ExecutionReportNew522) ExpireDateMaxValue() uint16 {
	return math.MaxUint16 - 1
}

func (*ExecutionReportNew522) ExpireDateNullValue() uint16 {
	return 65535
}

func (*ExecutionReportNew522) DelayDurationId() uint16 {
	return 5904
}

func (*ExecutionReportNew522) DelayDurationSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportNew522) DelayDurationInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.DelayDurationSinceVersion()
}

func (*ExecutionReportNew522) DelayDurationDeprecated() uint16 {
	return 0
}

func (*ExecutionReportNew522) DelayDurationMetaAttribute(meta int) string {
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

func (*ExecutionReportNew522) DelayDurationMinValue() uint16 {
	return 0
}

func (*ExecutionReportNew522) DelayDurationMaxValue() uint16 {
	return math.MaxUint16 - 1
}

func (*ExecutionReportNew522) DelayDurationNullValue() uint16 {
	return 65535
}

func (*ExecutionReportNew522) OrdStatusId() uint16 {
	return 39
}

func (*ExecutionReportNew522) OrdStatusSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportNew522) OrdStatusInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.OrdStatusSinceVersion()
}

func (*ExecutionReportNew522) OrdStatusDeprecated() uint16 {
	return 0
}

func (*ExecutionReportNew522) OrdStatusMetaAttribute(meta int) string {
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

func (*ExecutionReportNew522) OrdStatusMinValue() byte {
	return byte(32)
}

func (*ExecutionReportNew522) OrdStatusMaxValue() byte {
	return byte(126)
}

func (*ExecutionReportNew522) OrdStatusNullValue() byte {
	return 0
}

func (*ExecutionReportNew522) ExecTypeId() uint16 {
	return 150
}

func (*ExecutionReportNew522) ExecTypeSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportNew522) ExecTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.ExecTypeSinceVersion()
}

func (*ExecutionReportNew522) ExecTypeDeprecated() uint16 {
	return 0
}

func (*ExecutionReportNew522) ExecTypeMetaAttribute(meta int) string {
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

func (*ExecutionReportNew522) ExecTypeMinValue() byte {
	return byte(32)
}

func (*ExecutionReportNew522) ExecTypeMaxValue() byte {
	return byte(126)
}

func (*ExecutionReportNew522) ExecTypeNullValue() byte {
	return 0
}

func (*ExecutionReportNew522) OrdTypeId() uint16 {
	return 40
}

func (*ExecutionReportNew522) OrdTypeSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportNew522) OrdTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.OrdTypeSinceVersion()
}

func (*ExecutionReportNew522) OrdTypeDeprecated() uint16 {
	return 0
}

func (*ExecutionReportNew522) OrdTypeMetaAttribute(meta int) string {
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

func (*ExecutionReportNew522) SideId() uint16 {
	return 54
}

func (*ExecutionReportNew522) SideSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportNew522) SideInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.SideSinceVersion()
}

func (*ExecutionReportNew522) SideDeprecated() uint16 {
	return 0
}

func (*ExecutionReportNew522) SideMetaAttribute(meta int) string {
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

func (*ExecutionReportNew522) TimeInForceId() uint16 {
	return 59
}

func (*ExecutionReportNew522) TimeInForceSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportNew522) TimeInForceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.TimeInForceSinceVersion()
}

func (*ExecutionReportNew522) TimeInForceDeprecated() uint16 {
	return 0
}

func (*ExecutionReportNew522) TimeInForceMetaAttribute(meta int) string {
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

func (*ExecutionReportNew522) ManualOrderIndicatorId() uint16 {
	return 1028
}

func (*ExecutionReportNew522) ManualOrderIndicatorSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportNew522) ManualOrderIndicatorInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.ManualOrderIndicatorSinceVersion()
}

func (*ExecutionReportNew522) ManualOrderIndicatorDeprecated() uint16 {
	return 0
}

func (*ExecutionReportNew522) ManualOrderIndicatorMetaAttribute(meta int) string {
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

func (*ExecutionReportNew522) PossRetransFlagId() uint16 {
	return 9765
}

func (*ExecutionReportNew522) PossRetransFlagSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportNew522) PossRetransFlagInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.PossRetransFlagSinceVersion()
}

func (*ExecutionReportNew522) PossRetransFlagDeprecated() uint16 {
	return 0
}

func (*ExecutionReportNew522) PossRetransFlagMetaAttribute(meta int) string {
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

func (*ExecutionReportNew522) SplitMsgId() uint16 {
	return 9553
}

func (*ExecutionReportNew522) SplitMsgSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportNew522) SplitMsgInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.SplitMsgSinceVersion()
}

func (*ExecutionReportNew522) SplitMsgDeprecated() uint16 {
	return 0
}

func (*ExecutionReportNew522) SplitMsgMetaAttribute(meta int) string {
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

func (*ExecutionReportNew522) CrossTypeId() uint16 {
	return 549
}

func (*ExecutionReportNew522) CrossTypeSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportNew522) CrossTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.CrossTypeSinceVersion()
}

func (*ExecutionReportNew522) CrossTypeDeprecated() uint16 {
	return 0
}

func (*ExecutionReportNew522) CrossTypeMetaAttribute(meta int) string {
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

func (*ExecutionReportNew522) CrossTypeMinValue() uint8 {
	return 0
}

func (*ExecutionReportNew522) CrossTypeMaxValue() uint8 {
	return math.MaxUint8 - 1
}

func (*ExecutionReportNew522) CrossTypeNullValue() uint8 {
	return 255
}

func (*ExecutionReportNew522) ExecInstId() uint16 {
	return 18
}

func (*ExecutionReportNew522) ExecInstSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportNew522) ExecInstInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.ExecInstSinceVersion()
}

func (*ExecutionReportNew522) ExecInstDeprecated() uint16 {
	return 0
}

func (*ExecutionReportNew522) ExecInstMetaAttribute(meta int) string {
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

func (*ExecutionReportNew522) ExecutionModeId() uint16 {
	return 5906
}

func (*ExecutionReportNew522) ExecutionModeSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportNew522) ExecutionModeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.ExecutionModeSinceVersion()
}

func (*ExecutionReportNew522) ExecutionModeDeprecated() uint16 {
	return 0
}

func (*ExecutionReportNew522) ExecutionModeMetaAttribute(meta int) string {
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

func (*ExecutionReportNew522) LiquidityFlagId() uint16 {
	return 9373
}

func (*ExecutionReportNew522) LiquidityFlagSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportNew522) LiquidityFlagInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.LiquidityFlagSinceVersion()
}

func (*ExecutionReportNew522) LiquidityFlagDeprecated() uint16 {
	return 0
}

func (*ExecutionReportNew522) LiquidityFlagMetaAttribute(meta int) string {
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

func (*ExecutionReportNew522) ManagedOrderId() uint16 {
	return 6881
}

func (*ExecutionReportNew522) ManagedOrderSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportNew522) ManagedOrderInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.ManagedOrderSinceVersion()
}

func (*ExecutionReportNew522) ManagedOrderDeprecated() uint16 {
	return 0
}

func (*ExecutionReportNew522) ManagedOrderMetaAttribute(meta int) string {
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

func (*ExecutionReportNew522) ShortSaleTypeId() uint16 {
	return 5409
}

func (*ExecutionReportNew522) ShortSaleTypeSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportNew522) ShortSaleTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.ShortSaleTypeSinceVersion()
}

func (*ExecutionReportNew522) ShortSaleTypeDeprecated() uint16 {
	return 0
}

func (*ExecutionReportNew522) ShortSaleTypeMetaAttribute(meta int) string {
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

func (*ExecutionReportNew522) DelayToTimeId() uint16 {
	return 7552
}

func (*ExecutionReportNew522) DelayToTimeSinceVersion() uint16 {
	return 4
}

func (e *ExecutionReportNew522) DelayToTimeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.DelayToTimeSinceVersion()
}

func (*ExecutionReportNew522) DelayToTimeDeprecated() uint16 {
	return 0
}

func (*ExecutionReportNew522) DelayToTimeMetaAttribute(meta int) string {
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

func (*ExecutionReportNew522) DelayToTimeMinValue() uint64 {
	return 0
}

func (*ExecutionReportNew522) DelayToTimeMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*ExecutionReportNew522) DelayToTimeNullValue() uint64 {
	return 18446744073709551615
}
