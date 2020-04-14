// Generated SBE (Simple Binary Encoding) message codec

package ilinkbinary

import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
)

type ExecutionReportElimination524 struct {
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
	CumQty                uint32
	OrderQty              uint32
	MinQty                uint32
	DisplayQty            uint32
	OrdStatus             [1]byte
	ExecType              [1]byte
	ExpireDate            uint16
	OrdType               OrderTypeEnum
	Side                  SideReqEnum
	TimeInForce           TimeInForceEnum
	ManualOrderIndicator  ManualOrdIndReqEnum
	PossRetransFlag       BooleanFlagEnum
	CrossType             uint8
	ExecInst              ExecInst
	ExecutionMode         ExecModeEnum
	LiquidityFlag         BooleanNULLEnum
	ManagedOrder          BooleanNULLEnum
	ShortSaleType         ShortSaleTypeEnum
}

func (e *ExecutionReportElimination524) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
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
	if err := _m.WriteUint32(_w, e.CumQty); err != nil {
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

func (e *ExecutionReportElimination524) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
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
	if !e.CumQtyInActingVersion(actingVersion) {
		e.CumQty = e.CumQtyNullValue()
	} else {
		if err := _m.ReadUint32(_r, &e.CumQty); err != nil {
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
	e.OrdStatus[0] = 67
	e.ExecType[0] = 67
	if !e.ExpireDateInActingVersion(actingVersion) {
		e.ExpireDate = e.ExpireDateNullValue()
	} else {
		if err := _m.ReadUint16(_r, &e.ExpireDate); err != nil {
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

func (e *ExecutionReportElimination524) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
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
	if e.CumQtyInActingVersion(actingVersion) {
		if e.CumQty < e.CumQtyMinValue() || e.CumQty > e.CumQtyMaxValue() {
			return fmt.Errorf("Range check failed on e.CumQty (%v < %v > %v)", e.CumQtyMinValue(), e.CumQty, e.CumQtyMaxValue())
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

func ExecutionReportElimination524Init(e *ExecutionReportElimination524) {
	e.CrossID = 18446744073709551615
	e.HostCrossID = 18446744073709551615
	e.MinQty = 4294967295
	e.DisplayQty = 4294967295
	e.OrdStatus[0] = 67
	e.ExecType[0] = 67
	e.ExpireDate = 65535
	e.CrossType = 255
	return
}

func (*ExecutionReportElimination524) SbeBlockLength() (blockLength uint16) {
	return 202
}

func (*ExecutionReportElimination524) SbeTemplateId() (templateId uint16) {
	return 524
}

func (*ExecutionReportElimination524) SbeSchemaId() (schemaId uint16) {
	return 8
}

func (*ExecutionReportElimination524) SbeSchemaVersion() (schemaVersion uint16) {
	return 5
}

func (*ExecutionReportElimination524) SbeSemanticType() (semanticType []byte) {
	return []byte("8")
}

func (*ExecutionReportElimination524) SeqNumId() uint16 {
	return 9726
}

func (*ExecutionReportElimination524) SeqNumSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportElimination524) SeqNumInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.SeqNumSinceVersion()
}

func (*ExecutionReportElimination524) SeqNumDeprecated() uint16 {
	return 0
}

func (*ExecutionReportElimination524) SeqNumMetaAttribute(meta int) string {
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

func (*ExecutionReportElimination524) SeqNumMinValue() uint32 {
	return 0
}

func (*ExecutionReportElimination524) SeqNumMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*ExecutionReportElimination524) SeqNumNullValue() uint32 {
	return math.MaxUint32
}

func (*ExecutionReportElimination524) UUIDId() uint16 {
	return 39001
}

func (*ExecutionReportElimination524) UUIDSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportElimination524) UUIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.UUIDSinceVersion()
}

func (*ExecutionReportElimination524) UUIDDeprecated() uint16 {
	return 0
}

func (*ExecutionReportElimination524) UUIDMetaAttribute(meta int) string {
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

func (*ExecutionReportElimination524) UUIDMinValue() uint64 {
	return 0
}

func (*ExecutionReportElimination524) UUIDMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*ExecutionReportElimination524) UUIDNullValue() uint64 {
	return math.MaxUint64
}

func (*ExecutionReportElimination524) ExecIDId() uint16 {
	return 17
}

func (*ExecutionReportElimination524) ExecIDSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportElimination524) ExecIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.ExecIDSinceVersion()
}

func (*ExecutionReportElimination524) ExecIDDeprecated() uint16 {
	return 0
}

func (*ExecutionReportElimination524) ExecIDMetaAttribute(meta int) string {
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

func (*ExecutionReportElimination524) ExecIDMinValue() byte {
	return byte(32)
}

func (*ExecutionReportElimination524) ExecIDMaxValue() byte {
	return byte(126)
}

func (*ExecutionReportElimination524) ExecIDNullValue() byte {
	return 0
}

func (e *ExecutionReportElimination524) ExecIDCharacterEncoding() string {
	return "US-ASCII"
}

func (*ExecutionReportElimination524) SenderIDId() uint16 {
	return 5392
}

func (*ExecutionReportElimination524) SenderIDSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportElimination524) SenderIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.SenderIDSinceVersion()
}

func (*ExecutionReportElimination524) SenderIDDeprecated() uint16 {
	return 0
}

func (*ExecutionReportElimination524) SenderIDMetaAttribute(meta int) string {
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

func (*ExecutionReportElimination524) SenderIDMinValue() byte {
	return byte(32)
}

func (*ExecutionReportElimination524) SenderIDMaxValue() byte {
	return byte(126)
}

func (*ExecutionReportElimination524) SenderIDNullValue() byte {
	return 0
}

func (e *ExecutionReportElimination524) SenderIDCharacterEncoding() string {
	return "US-ASCII"
}

func (*ExecutionReportElimination524) ClOrdIDId() uint16 {
	return 11
}

func (*ExecutionReportElimination524) ClOrdIDSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportElimination524) ClOrdIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.ClOrdIDSinceVersion()
}

func (*ExecutionReportElimination524) ClOrdIDDeprecated() uint16 {
	return 0
}

func (*ExecutionReportElimination524) ClOrdIDMetaAttribute(meta int) string {
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

func (*ExecutionReportElimination524) ClOrdIDMinValue() byte {
	return byte(32)
}

func (*ExecutionReportElimination524) ClOrdIDMaxValue() byte {
	return byte(126)
}

func (*ExecutionReportElimination524) ClOrdIDNullValue() byte {
	return 0
}

func (e *ExecutionReportElimination524) ClOrdIDCharacterEncoding() string {
	return "US-ASCII"
}

func (*ExecutionReportElimination524) PartyDetailsListReqIDId() uint16 {
	return 1505
}

func (*ExecutionReportElimination524) PartyDetailsListReqIDSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportElimination524) PartyDetailsListReqIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.PartyDetailsListReqIDSinceVersion()
}

func (*ExecutionReportElimination524) PartyDetailsListReqIDDeprecated() uint16 {
	return 0
}

func (*ExecutionReportElimination524) PartyDetailsListReqIDMetaAttribute(meta int) string {
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

func (*ExecutionReportElimination524) PartyDetailsListReqIDMinValue() uint64 {
	return 0
}

func (*ExecutionReportElimination524) PartyDetailsListReqIDMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*ExecutionReportElimination524) PartyDetailsListReqIDNullValue() uint64 {
	return math.MaxUint64
}

func (*ExecutionReportElimination524) OrderIDId() uint16 {
	return 37
}

func (*ExecutionReportElimination524) OrderIDSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportElimination524) OrderIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.OrderIDSinceVersion()
}

func (*ExecutionReportElimination524) OrderIDDeprecated() uint16 {
	return 0
}

func (*ExecutionReportElimination524) OrderIDMetaAttribute(meta int) string {
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

func (*ExecutionReportElimination524) OrderIDMinValue() uint64 {
	return 0
}

func (*ExecutionReportElimination524) OrderIDMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*ExecutionReportElimination524) OrderIDNullValue() uint64 {
	return math.MaxUint64
}

func (*ExecutionReportElimination524) PriceId() uint16 {
	return 44
}

func (*ExecutionReportElimination524) PriceSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportElimination524) PriceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.PriceSinceVersion()
}

func (*ExecutionReportElimination524) PriceDeprecated() uint16 {
	return 0
}

func (*ExecutionReportElimination524) PriceMetaAttribute(meta int) string {
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

func (*ExecutionReportElimination524) StopPxId() uint16 {
	return 99
}

func (*ExecutionReportElimination524) StopPxSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportElimination524) StopPxInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.StopPxSinceVersion()
}

func (*ExecutionReportElimination524) StopPxDeprecated() uint16 {
	return 0
}

func (*ExecutionReportElimination524) StopPxMetaAttribute(meta int) string {
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

func (*ExecutionReportElimination524) TransactTimeId() uint16 {
	return 60
}

func (*ExecutionReportElimination524) TransactTimeSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportElimination524) TransactTimeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.TransactTimeSinceVersion()
}

func (*ExecutionReportElimination524) TransactTimeDeprecated() uint16 {
	return 0
}

func (*ExecutionReportElimination524) TransactTimeMetaAttribute(meta int) string {
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

func (*ExecutionReportElimination524) TransactTimeMinValue() uint64 {
	return 0
}

func (*ExecutionReportElimination524) TransactTimeMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*ExecutionReportElimination524) TransactTimeNullValue() uint64 {
	return math.MaxUint64
}

func (*ExecutionReportElimination524) SendingTimeEpochId() uint16 {
	return 5297
}

func (*ExecutionReportElimination524) SendingTimeEpochSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportElimination524) SendingTimeEpochInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.SendingTimeEpochSinceVersion()
}

func (*ExecutionReportElimination524) SendingTimeEpochDeprecated() uint16 {
	return 0
}

func (*ExecutionReportElimination524) SendingTimeEpochMetaAttribute(meta int) string {
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

func (*ExecutionReportElimination524) SendingTimeEpochMinValue() uint64 {
	return 0
}

func (*ExecutionReportElimination524) SendingTimeEpochMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*ExecutionReportElimination524) SendingTimeEpochNullValue() uint64 {
	return math.MaxUint64
}

func (*ExecutionReportElimination524) OrderRequestIDId() uint16 {
	return 2422
}

func (*ExecutionReportElimination524) OrderRequestIDSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportElimination524) OrderRequestIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.OrderRequestIDSinceVersion()
}

func (*ExecutionReportElimination524) OrderRequestIDDeprecated() uint16 {
	return 0
}

func (*ExecutionReportElimination524) OrderRequestIDMetaAttribute(meta int) string {
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

func (*ExecutionReportElimination524) OrderRequestIDMinValue() uint64 {
	return 0
}

func (*ExecutionReportElimination524) OrderRequestIDMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*ExecutionReportElimination524) OrderRequestIDNullValue() uint64 {
	return math.MaxUint64
}

func (*ExecutionReportElimination524) CrossIDId() uint16 {
	return 548
}

func (*ExecutionReportElimination524) CrossIDSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportElimination524) CrossIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.CrossIDSinceVersion()
}

func (*ExecutionReportElimination524) CrossIDDeprecated() uint16 {
	return 0
}

func (*ExecutionReportElimination524) CrossIDMetaAttribute(meta int) string {
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

func (*ExecutionReportElimination524) CrossIDMinValue() uint64 {
	return 0
}

func (*ExecutionReportElimination524) CrossIDMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*ExecutionReportElimination524) CrossIDNullValue() uint64 {
	return 18446744073709551615
}

func (*ExecutionReportElimination524) HostCrossIDId() uint16 {
	return 961
}

func (*ExecutionReportElimination524) HostCrossIDSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportElimination524) HostCrossIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.HostCrossIDSinceVersion()
}

func (*ExecutionReportElimination524) HostCrossIDDeprecated() uint16 {
	return 0
}

func (*ExecutionReportElimination524) HostCrossIDMetaAttribute(meta int) string {
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

func (*ExecutionReportElimination524) HostCrossIDMinValue() uint64 {
	return 0
}

func (*ExecutionReportElimination524) HostCrossIDMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*ExecutionReportElimination524) HostCrossIDNullValue() uint64 {
	return 18446744073709551615
}

func (*ExecutionReportElimination524) LocationId() uint16 {
	return 9537
}

func (*ExecutionReportElimination524) LocationSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportElimination524) LocationInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.LocationSinceVersion()
}

func (*ExecutionReportElimination524) LocationDeprecated() uint16 {
	return 0
}

func (*ExecutionReportElimination524) LocationMetaAttribute(meta int) string {
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

func (*ExecutionReportElimination524) LocationMinValue() byte {
	return byte(32)
}

func (*ExecutionReportElimination524) LocationMaxValue() byte {
	return byte(126)
}

func (*ExecutionReportElimination524) LocationNullValue() byte {
	return 0
}

func (e *ExecutionReportElimination524) LocationCharacterEncoding() string {
	return "US-ASCII"
}

func (*ExecutionReportElimination524) SecurityIDId() uint16 {
	return 48
}

func (*ExecutionReportElimination524) SecurityIDSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportElimination524) SecurityIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.SecurityIDSinceVersion()
}

func (*ExecutionReportElimination524) SecurityIDDeprecated() uint16 {
	return 0
}

func (*ExecutionReportElimination524) SecurityIDMetaAttribute(meta int) string {
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

func (*ExecutionReportElimination524) SecurityIDMinValue() int32 {
	return math.MinInt32 + 1
}

func (*ExecutionReportElimination524) SecurityIDMaxValue() int32 {
	return math.MaxInt32
}

func (*ExecutionReportElimination524) SecurityIDNullValue() int32 {
	return math.MinInt32
}

func (*ExecutionReportElimination524) CumQtyId() uint16 {
	return 14
}

func (*ExecutionReportElimination524) CumQtySinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportElimination524) CumQtyInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.CumQtySinceVersion()
}

func (*ExecutionReportElimination524) CumQtyDeprecated() uint16 {
	return 0
}

func (*ExecutionReportElimination524) CumQtyMetaAttribute(meta int) string {
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

func (*ExecutionReportElimination524) CumQtyMinValue() uint32 {
	return 0
}

func (*ExecutionReportElimination524) CumQtyMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*ExecutionReportElimination524) CumQtyNullValue() uint32 {
	return math.MaxUint32
}

func (*ExecutionReportElimination524) OrderQtyId() uint16 {
	return 38
}

func (*ExecutionReportElimination524) OrderQtySinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportElimination524) OrderQtyInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.OrderQtySinceVersion()
}

func (*ExecutionReportElimination524) OrderQtyDeprecated() uint16 {
	return 0
}

func (*ExecutionReportElimination524) OrderQtyMetaAttribute(meta int) string {
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

func (*ExecutionReportElimination524) OrderQtyMinValue() uint32 {
	return 0
}

func (*ExecutionReportElimination524) OrderQtyMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*ExecutionReportElimination524) OrderQtyNullValue() uint32 {
	return math.MaxUint32
}

func (*ExecutionReportElimination524) MinQtyId() uint16 {
	return 110
}

func (*ExecutionReportElimination524) MinQtySinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportElimination524) MinQtyInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.MinQtySinceVersion()
}

func (*ExecutionReportElimination524) MinQtyDeprecated() uint16 {
	return 0
}

func (*ExecutionReportElimination524) MinQtyMetaAttribute(meta int) string {
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

func (*ExecutionReportElimination524) MinQtyMinValue() uint32 {
	return 0
}

func (*ExecutionReportElimination524) MinQtyMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*ExecutionReportElimination524) MinQtyNullValue() uint32 {
	return 4294967295
}

func (*ExecutionReportElimination524) DisplayQtyId() uint16 {
	return 1138
}

func (*ExecutionReportElimination524) DisplayQtySinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportElimination524) DisplayQtyInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.DisplayQtySinceVersion()
}

func (*ExecutionReportElimination524) DisplayQtyDeprecated() uint16 {
	return 0
}

func (*ExecutionReportElimination524) DisplayQtyMetaAttribute(meta int) string {
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

func (*ExecutionReportElimination524) DisplayQtyMinValue() uint32 {
	return 0
}

func (*ExecutionReportElimination524) DisplayQtyMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*ExecutionReportElimination524) DisplayQtyNullValue() uint32 {
	return 4294967295
}

func (*ExecutionReportElimination524) OrdStatusId() uint16 {
	return 39
}

func (*ExecutionReportElimination524) OrdStatusSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportElimination524) OrdStatusInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.OrdStatusSinceVersion()
}

func (*ExecutionReportElimination524) OrdStatusDeprecated() uint16 {
	return 0
}

func (*ExecutionReportElimination524) OrdStatusMetaAttribute(meta int) string {
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

func (*ExecutionReportElimination524) OrdStatusMinValue() byte {
	return byte(32)
}

func (*ExecutionReportElimination524) OrdStatusMaxValue() byte {
	return byte(126)
}

func (*ExecutionReportElimination524) OrdStatusNullValue() byte {
	return 0
}

func (*ExecutionReportElimination524) ExecTypeId() uint16 {
	return 150
}

func (*ExecutionReportElimination524) ExecTypeSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportElimination524) ExecTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.ExecTypeSinceVersion()
}

func (*ExecutionReportElimination524) ExecTypeDeprecated() uint16 {
	return 0
}

func (*ExecutionReportElimination524) ExecTypeMetaAttribute(meta int) string {
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

func (*ExecutionReportElimination524) ExecTypeMinValue() byte {
	return byte(32)
}

func (*ExecutionReportElimination524) ExecTypeMaxValue() byte {
	return byte(126)
}

func (*ExecutionReportElimination524) ExecTypeNullValue() byte {
	return 0
}

func (*ExecutionReportElimination524) ExpireDateId() uint16 {
	return 432
}

func (*ExecutionReportElimination524) ExpireDateSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportElimination524) ExpireDateInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.ExpireDateSinceVersion()
}

func (*ExecutionReportElimination524) ExpireDateDeprecated() uint16 {
	return 0
}

func (*ExecutionReportElimination524) ExpireDateMetaAttribute(meta int) string {
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

func (*ExecutionReportElimination524) ExpireDateMinValue() uint16 {
	return 0
}

func (*ExecutionReportElimination524) ExpireDateMaxValue() uint16 {
	return math.MaxUint16 - 1
}

func (*ExecutionReportElimination524) ExpireDateNullValue() uint16 {
	return 65535
}

func (*ExecutionReportElimination524) OrdTypeId() uint16 {
	return 40
}

func (*ExecutionReportElimination524) OrdTypeSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportElimination524) OrdTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.OrdTypeSinceVersion()
}

func (*ExecutionReportElimination524) OrdTypeDeprecated() uint16 {
	return 0
}

func (*ExecutionReportElimination524) OrdTypeMetaAttribute(meta int) string {
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

func (*ExecutionReportElimination524) SideId() uint16 {
	return 54
}

func (*ExecutionReportElimination524) SideSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportElimination524) SideInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.SideSinceVersion()
}

func (*ExecutionReportElimination524) SideDeprecated() uint16 {
	return 0
}

func (*ExecutionReportElimination524) SideMetaAttribute(meta int) string {
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

func (*ExecutionReportElimination524) TimeInForceId() uint16 {
	return 59
}

func (*ExecutionReportElimination524) TimeInForceSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportElimination524) TimeInForceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.TimeInForceSinceVersion()
}

func (*ExecutionReportElimination524) TimeInForceDeprecated() uint16 {
	return 0
}

func (*ExecutionReportElimination524) TimeInForceMetaAttribute(meta int) string {
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

func (*ExecutionReportElimination524) ManualOrderIndicatorId() uint16 {
	return 1028
}

func (*ExecutionReportElimination524) ManualOrderIndicatorSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportElimination524) ManualOrderIndicatorInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.ManualOrderIndicatorSinceVersion()
}

func (*ExecutionReportElimination524) ManualOrderIndicatorDeprecated() uint16 {
	return 0
}

func (*ExecutionReportElimination524) ManualOrderIndicatorMetaAttribute(meta int) string {
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

func (*ExecutionReportElimination524) PossRetransFlagId() uint16 {
	return 9765
}

func (*ExecutionReportElimination524) PossRetransFlagSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportElimination524) PossRetransFlagInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.PossRetransFlagSinceVersion()
}

func (*ExecutionReportElimination524) PossRetransFlagDeprecated() uint16 {
	return 0
}

func (*ExecutionReportElimination524) PossRetransFlagMetaAttribute(meta int) string {
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

func (*ExecutionReportElimination524) CrossTypeId() uint16 {
	return 549
}

func (*ExecutionReportElimination524) CrossTypeSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportElimination524) CrossTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.CrossTypeSinceVersion()
}

func (*ExecutionReportElimination524) CrossTypeDeprecated() uint16 {
	return 0
}

func (*ExecutionReportElimination524) CrossTypeMetaAttribute(meta int) string {
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

func (*ExecutionReportElimination524) CrossTypeMinValue() uint8 {
	return 0
}

func (*ExecutionReportElimination524) CrossTypeMaxValue() uint8 {
	return math.MaxUint8 - 1
}

func (*ExecutionReportElimination524) CrossTypeNullValue() uint8 {
	return 255
}

func (*ExecutionReportElimination524) ExecInstId() uint16 {
	return 18
}

func (*ExecutionReportElimination524) ExecInstSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportElimination524) ExecInstInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.ExecInstSinceVersion()
}

func (*ExecutionReportElimination524) ExecInstDeprecated() uint16 {
	return 0
}

func (*ExecutionReportElimination524) ExecInstMetaAttribute(meta int) string {
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

func (*ExecutionReportElimination524) ExecutionModeId() uint16 {
	return 5906
}

func (*ExecutionReportElimination524) ExecutionModeSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportElimination524) ExecutionModeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.ExecutionModeSinceVersion()
}

func (*ExecutionReportElimination524) ExecutionModeDeprecated() uint16 {
	return 0
}

func (*ExecutionReportElimination524) ExecutionModeMetaAttribute(meta int) string {
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

func (*ExecutionReportElimination524) LiquidityFlagId() uint16 {
	return 9373
}

func (*ExecutionReportElimination524) LiquidityFlagSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportElimination524) LiquidityFlagInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.LiquidityFlagSinceVersion()
}

func (*ExecutionReportElimination524) LiquidityFlagDeprecated() uint16 {
	return 0
}

func (*ExecutionReportElimination524) LiquidityFlagMetaAttribute(meta int) string {
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

func (*ExecutionReportElimination524) ManagedOrderId() uint16 {
	return 6881
}

func (*ExecutionReportElimination524) ManagedOrderSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportElimination524) ManagedOrderInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.ManagedOrderSinceVersion()
}

func (*ExecutionReportElimination524) ManagedOrderDeprecated() uint16 {
	return 0
}

func (*ExecutionReportElimination524) ManagedOrderMetaAttribute(meta int) string {
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

func (*ExecutionReportElimination524) ShortSaleTypeId() uint16 {
	return 5409
}

func (*ExecutionReportElimination524) ShortSaleTypeSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportElimination524) ShortSaleTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.ShortSaleTypeSinceVersion()
}

func (*ExecutionReportElimination524) ShortSaleTypeDeprecated() uint16 {
	return 0
}

func (*ExecutionReportElimination524) ShortSaleTypeMetaAttribute(meta int) string {
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
