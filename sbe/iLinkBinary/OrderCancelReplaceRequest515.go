// Generated SBE (Simple Binary Encoding) message codec

package ilinkbinary

import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
)

type OrderCancelReplaceRequest515 struct {
	Price                 PRICENULL9
	OrderQty              uint32
	SecurityID            int32
	Side                  SideReqEnum
	SeqNum                uint32
	SenderID              [20]byte
	ClOrdID               [20]byte
	PartyDetailsListReqID uint64
	OrderID               uint64
	StopPx                PRICENULL9
	OrderRequestID        uint64
	SendingTimeEpoch      uint64
	Location              [5]byte
	MinQty                uint32
	DisplayQty            uint32
	ExpireDate            uint16
	OrdType               OrderTypeReqEnum
	TimeInForce           TimeInForceEnum
	ManualOrderIndicator  ManualOrdIndReqEnum
	OFMOverride           OFMOverrideReqEnum
	ExecInst              ExecInst
	ExecutionMode         ExecModeEnum
	LiquidityFlag         BooleanNULLEnum
	ManagedOrder          BooleanNULLEnum
	ShortSaleType         ShortSaleTypeEnum
}

func (o *OrderCancelReplaceRequest515) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := o.RangeCheck(o.SbeSchemaVersion(), o.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	if err := o.Price.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteUint32(_w, o.OrderQty); err != nil {
		return err
	}
	if err := _m.WriteInt32(_w, o.SecurityID); err != nil {
		return err
	}
	if err := o.Side.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteUint32(_w, o.SeqNum); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, o.SenderID[:]); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, o.ClOrdID[:]); err != nil {
		return err
	}
	if err := _m.WriteUint64(_w, o.PartyDetailsListReqID); err != nil {
		return err
	}
	if err := _m.WriteUint64(_w, o.OrderID); err != nil {
		return err
	}
	if err := o.StopPx.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteUint64(_w, o.OrderRequestID); err != nil {
		return err
	}
	if err := _m.WriteUint64(_w, o.SendingTimeEpoch); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, o.Location[:]); err != nil {
		return err
	}
	if err := _m.WriteUint32(_w, o.MinQty); err != nil {
		return err
	}
	if err := _m.WriteUint32(_w, o.DisplayQty); err != nil {
		return err
	}
	if err := _m.WriteUint16(_w, o.ExpireDate); err != nil {
		return err
	}
	if err := o.OrdType.Encode(_m, _w); err != nil {
		return err
	}
	if err := o.TimeInForce.Encode(_m, _w); err != nil {
		return err
	}
	if err := o.ManualOrderIndicator.Encode(_m, _w); err != nil {
		return err
	}
	if err := o.OFMOverride.Encode(_m, _w); err != nil {
		return err
	}
	if err := o.ExecInst.Encode(_m, _w); err != nil {
		return err
	}
	if err := o.ExecutionMode.Encode(_m, _w); err != nil {
		return err
	}
	if err := o.LiquidityFlag.Encode(_m, _w); err != nil {
		return err
	}
	if err := o.ManagedOrder.Encode(_m, _w); err != nil {
		return err
	}
	if err := o.ShortSaleType.Encode(_m, _w); err != nil {
		return err
	}
	return nil
}

func (o *OrderCancelReplaceRequest515) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	if o.PriceInActingVersion(actingVersion) {
		if err := o.Price.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if !o.OrderQtyInActingVersion(actingVersion) {
		o.OrderQty = o.OrderQtyNullValue()
	} else {
		if err := _m.ReadUint32(_r, &o.OrderQty); err != nil {
			return err
		}
	}
	if !o.SecurityIDInActingVersion(actingVersion) {
		o.SecurityID = o.SecurityIDNullValue()
	} else {
		if err := _m.ReadInt32(_r, &o.SecurityID); err != nil {
			return err
		}
	}
	if o.SideInActingVersion(actingVersion) {
		if err := o.Side.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if !o.SeqNumInActingVersion(actingVersion) {
		o.SeqNum = o.SeqNumNullValue()
	} else {
		if err := _m.ReadUint32(_r, &o.SeqNum); err != nil {
			return err
		}
	}
	if !o.SenderIDInActingVersion(actingVersion) {
		for idx := 0; idx < 20; idx++ {
			o.SenderID[idx] = o.SenderIDNullValue()
		}
	} else {
		if err := _m.ReadBytes(_r, o.SenderID[:]); err != nil {
			return err
		}
	}
	if !o.ClOrdIDInActingVersion(actingVersion) {
		for idx := 0; idx < 20; idx++ {
			o.ClOrdID[idx] = o.ClOrdIDNullValue()
		}
	} else {
		if err := _m.ReadBytes(_r, o.ClOrdID[:]); err != nil {
			return err
		}
	}
	if !o.PartyDetailsListReqIDInActingVersion(actingVersion) {
		o.PartyDetailsListReqID = o.PartyDetailsListReqIDNullValue()
	} else {
		if err := _m.ReadUint64(_r, &o.PartyDetailsListReqID); err != nil {
			return err
		}
	}
	if !o.OrderIDInActingVersion(actingVersion) {
		o.OrderID = o.OrderIDNullValue()
	} else {
		if err := _m.ReadUint64(_r, &o.OrderID); err != nil {
			return err
		}
	}
	if o.StopPxInActingVersion(actingVersion) {
		if err := o.StopPx.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if !o.OrderRequestIDInActingVersion(actingVersion) {
		o.OrderRequestID = o.OrderRequestIDNullValue()
	} else {
		if err := _m.ReadUint64(_r, &o.OrderRequestID); err != nil {
			return err
		}
	}
	if !o.SendingTimeEpochInActingVersion(actingVersion) {
		o.SendingTimeEpoch = o.SendingTimeEpochNullValue()
	} else {
		if err := _m.ReadUint64(_r, &o.SendingTimeEpoch); err != nil {
			return err
		}
	}
	if !o.LocationInActingVersion(actingVersion) {
		for idx := 0; idx < 5; idx++ {
			o.Location[idx] = o.LocationNullValue()
		}
	} else {
		if err := _m.ReadBytes(_r, o.Location[:]); err != nil {
			return err
		}
	}
	if !o.MinQtyInActingVersion(actingVersion) {
		o.MinQty = o.MinQtyNullValue()
	} else {
		if err := _m.ReadUint32(_r, &o.MinQty); err != nil {
			return err
		}
	}
	if !o.DisplayQtyInActingVersion(actingVersion) {
		o.DisplayQty = o.DisplayQtyNullValue()
	} else {
		if err := _m.ReadUint32(_r, &o.DisplayQty); err != nil {
			return err
		}
	}
	if !o.ExpireDateInActingVersion(actingVersion) {
		o.ExpireDate = o.ExpireDateNullValue()
	} else {
		if err := _m.ReadUint16(_r, &o.ExpireDate); err != nil {
			return err
		}
	}
	if o.OrdTypeInActingVersion(actingVersion) {
		if err := o.OrdType.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if o.TimeInForceInActingVersion(actingVersion) {
		if err := o.TimeInForce.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if o.ManualOrderIndicatorInActingVersion(actingVersion) {
		if err := o.ManualOrderIndicator.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if o.OFMOverrideInActingVersion(actingVersion) {
		if err := o.OFMOverride.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if o.ExecInstInActingVersion(actingVersion) {
		if err := o.ExecInst.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if o.ExecutionModeInActingVersion(actingVersion) {
		if err := o.ExecutionMode.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if o.LiquidityFlagInActingVersion(actingVersion) {
		if err := o.LiquidityFlag.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if o.ManagedOrderInActingVersion(actingVersion) {
		if err := o.ManagedOrder.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if o.ShortSaleTypeInActingVersion(actingVersion) {
		if err := o.ShortSaleType.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if actingVersion > o.SbeSchemaVersion() && blockLength > o.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-o.SbeBlockLength()))
	}
	if doRangeCheck {
		if err := o.RangeCheck(actingVersion, o.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	return nil
}

func (o *OrderCancelReplaceRequest515) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if o.OrderQtyInActingVersion(actingVersion) {
		if o.OrderQty < o.OrderQtyMinValue() || o.OrderQty > o.OrderQtyMaxValue() {
			return fmt.Errorf("Range check failed on o.OrderQty (%v < %v > %v)", o.OrderQtyMinValue(), o.OrderQty, o.OrderQtyMaxValue())
		}
	}
	if o.SecurityIDInActingVersion(actingVersion) {
		if o.SecurityID < o.SecurityIDMinValue() || o.SecurityID > o.SecurityIDMaxValue() {
			return fmt.Errorf("Range check failed on o.SecurityID (%v < %v > %v)", o.SecurityIDMinValue(), o.SecurityID, o.SecurityIDMaxValue())
		}
	}
	if err := o.Side.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if o.SeqNumInActingVersion(actingVersion) {
		if o.SeqNum < o.SeqNumMinValue() || o.SeqNum > o.SeqNumMaxValue() {
			return fmt.Errorf("Range check failed on o.SeqNum (%v < %v > %v)", o.SeqNumMinValue(), o.SeqNum, o.SeqNumMaxValue())
		}
	}
	if o.SenderIDInActingVersion(actingVersion) {
		for idx := 0; idx < 20; idx++ {
			if o.SenderID[idx] < o.SenderIDMinValue() || o.SenderID[idx] > o.SenderIDMaxValue() {
				return fmt.Errorf("Range check failed on o.SenderID[%d] (%v < %v > %v)", idx, o.SenderIDMinValue(), o.SenderID[idx], o.SenderIDMaxValue())
			}
		}
	}
	if o.ClOrdIDInActingVersion(actingVersion) {
		for idx := 0; idx < 20; idx++ {
			if o.ClOrdID[idx] < o.ClOrdIDMinValue() || o.ClOrdID[idx] > o.ClOrdIDMaxValue() {
				return fmt.Errorf("Range check failed on o.ClOrdID[%d] (%v < %v > %v)", idx, o.ClOrdIDMinValue(), o.ClOrdID[idx], o.ClOrdIDMaxValue())
			}
		}
	}
	if o.PartyDetailsListReqIDInActingVersion(actingVersion) {
		if o.PartyDetailsListReqID < o.PartyDetailsListReqIDMinValue() || o.PartyDetailsListReqID > o.PartyDetailsListReqIDMaxValue() {
			return fmt.Errorf("Range check failed on o.PartyDetailsListReqID (%v < %v > %v)", o.PartyDetailsListReqIDMinValue(), o.PartyDetailsListReqID, o.PartyDetailsListReqIDMaxValue())
		}
	}
	if o.OrderIDInActingVersion(actingVersion) {
		if o.OrderID < o.OrderIDMinValue() || o.OrderID > o.OrderIDMaxValue() {
			return fmt.Errorf("Range check failed on o.OrderID (%v < %v > %v)", o.OrderIDMinValue(), o.OrderID, o.OrderIDMaxValue())
		}
	}
	if o.OrderRequestIDInActingVersion(actingVersion) {
		if o.OrderRequestID < o.OrderRequestIDMinValue() || o.OrderRequestID > o.OrderRequestIDMaxValue() {
			return fmt.Errorf("Range check failed on o.OrderRequestID (%v < %v > %v)", o.OrderRequestIDMinValue(), o.OrderRequestID, o.OrderRequestIDMaxValue())
		}
	}
	if o.SendingTimeEpochInActingVersion(actingVersion) {
		if o.SendingTimeEpoch < o.SendingTimeEpochMinValue() || o.SendingTimeEpoch > o.SendingTimeEpochMaxValue() {
			return fmt.Errorf("Range check failed on o.SendingTimeEpoch (%v < %v > %v)", o.SendingTimeEpochMinValue(), o.SendingTimeEpoch, o.SendingTimeEpochMaxValue())
		}
	}
	if o.LocationInActingVersion(actingVersion) {
		for idx := 0; idx < 5; idx++ {
			if o.Location[idx] < o.LocationMinValue() || o.Location[idx] > o.LocationMaxValue() {
				return fmt.Errorf("Range check failed on o.Location[%d] (%v < %v > %v)", idx, o.LocationMinValue(), o.Location[idx], o.LocationMaxValue())
			}
		}
	}
	if o.MinQtyInActingVersion(actingVersion) {
		if o.MinQty != o.MinQtyNullValue() && (o.MinQty < o.MinQtyMinValue() || o.MinQty > o.MinQtyMaxValue()) {
			return fmt.Errorf("Range check failed on o.MinQty (%v < %v > %v)", o.MinQtyMinValue(), o.MinQty, o.MinQtyMaxValue())
		}
	}
	if o.DisplayQtyInActingVersion(actingVersion) {
		if o.DisplayQty != o.DisplayQtyNullValue() && (o.DisplayQty < o.DisplayQtyMinValue() || o.DisplayQty > o.DisplayQtyMaxValue()) {
			return fmt.Errorf("Range check failed on o.DisplayQty (%v < %v > %v)", o.DisplayQtyMinValue(), o.DisplayQty, o.DisplayQtyMaxValue())
		}
	}
	if o.ExpireDateInActingVersion(actingVersion) {
		if o.ExpireDate != o.ExpireDateNullValue() && (o.ExpireDate < o.ExpireDateMinValue() || o.ExpireDate > o.ExpireDateMaxValue()) {
			return fmt.Errorf("Range check failed on o.ExpireDate (%v < %v > %v)", o.ExpireDateMinValue(), o.ExpireDate, o.ExpireDateMaxValue())
		}
	}
	if err := o.OrdType.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if err := o.TimeInForce.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if err := o.ManualOrderIndicator.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if err := o.OFMOverride.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if err := o.ExecutionMode.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if err := o.LiquidityFlag.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if err := o.ManagedOrder.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if err := o.ShortSaleType.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	return nil
}

func OrderCancelReplaceRequest515Init(o *OrderCancelReplaceRequest515) {
	o.MinQty = 4294967295
	o.DisplayQty = 4294967295
	o.ExpireDate = 65535
	return
}

func (*OrderCancelReplaceRequest515) SbeBlockLength() (blockLength uint16) {
	return 125
}

func (*OrderCancelReplaceRequest515) SbeTemplateId() (templateId uint16) {
	return 515
}

func (*OrderCancelReplaceRequest515) SbeSchemaId() (schemaId uint16) {
	return 8
}

func (*OrderCancelReplaceRequest515) SbeSchemaVersion() (schemaVersion uint16) {
	return 5
}

func (*OrderCancelReplaceRequest515) SbeSemanticType() (semanticType []byte) {
	return []byte("G")
}

func (*OrderCancelReplaceRequest515) PriceId() uint16 {
	return 44
}

func (*OrderCancelReplaceRequest515) PriceSinceVersion() uint16 {
	return 0
}

func (o *OrderCancelReplaceRequest515) PriceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.PriceSinceVersion()
}

func (*OrderCancelReplaceRequest515) PriceDeprecated() uint16 {
	return 0
}

func (*OrderCancelReplaceRequest515) PriceMetaAttribute(meta int) string {
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

func (*OrderCancelReplaceRequest515) OrderQtyId() uint16 {
	return 38
}

func (*OrderCancelReplaceRequest515) OrderQtySinceVersion() uint16 {
	return 0
}

func (o *OrderCancelReplaceRequest515) OrderQtyInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.OrderQtySinceVersion()
}

func (*OrderCancelReplaceRequest515) OrderQtyDeprecated() uint16 {
	return 0
}

func (*OrderCancelReplaceRequest515) OrderQtyMetaAttribute(meta int) string {
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

func (*OrderCancelReplaceRequest515) OrderQtyMinValue() uint32 {
	return 0
}

func (*OrderCancelReplaceRequest515) OrderQtyMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*OrderCancelReplaceRequest515) OrderQtyNullValue() uint32 {
	return math.MaxUint32
}

func (*OrderCancelReplaceRequest515) SecurityIDId() uint16 {
	return 48
}

func (*OrderCancelReplaceRequest515) SecurityIDSinceVersion() uint16 {
	return 0
}

func (o *OrderCancelReplaceRequest515) SecurityIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.SecurityIDSinceVersion()
}

func (*OrderCancelReplaceRequest515) SecurityIDDeprecated() uint16 {
	return 0
}

func (*OrderCancelReplaceRequest515) SecurityIDMetaAttribute(meta int) string {
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

func (*OrderCancelReplaceRequest515) SecurityIDMinValue() int32 {
	return math.MinInt32 + 1
}

func (*OrderCancelReplaceRequest515) SecurityIDMaxValue() int32 {
	return math.MaxInt32
}

func (*OrderCancelReplaceRequest515) SecurityIDNullValue() int32 {
	return math.MinInt32
}

func (*OrderCancelReplaceRequest515) SideId() uint16 {
	return 54
}

func (*OrderCancelReplaceRequest515) SideSinceVersion() uint16 {
	return 0
}

func (o *OrderCancelReplaceRequest515) SideInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.SideSinceVersion()
}

func (*OrderCancelReplaceRequest515) SideDeprecated() uint16 {
	return 0
}

func (*OrderCancelReplaceRequest515) SideMetaAttribute(meta int) string {
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

func (*OrderCancelReplaceRequest515) SeqNumId() uint16 {
	return 9726
}

func (*OrderCancelReplaceRequest515) SeqNumSinceVersion() uint16 {
	return 0
}

func (o *OrderCancelReplaceRequest515) SeqNumInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.SeqNumSinceVersion()
}

func (*OrderCancelReplaceRequest515) SeqNumDeprecated() uint16 {
	return 0
}

func (*OrderCancelReplaceRequest515) SeqNumMetaAttribute(meta int) string {
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

func (*OrderCancelReplaceRequest515) SeqNumMinValue() uint32 {
	return 0
}

func (*OrderCancelReplaceRequest515) SeqNumMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*OrderCancelReplaceRequest515) SeqNumNullValue() uint32 {
	return math.MaxUint32
}

func (*OrderCancelReplaceRequest515) SenderIDId() uint16 {
	return 5392
}

func (*OrderCancelReplaceRequest515) SenderIDSinceVersion() uint16 {
	return 0
}

func (o *OrderCancelReplaceRequest515) SenderIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.SenderIDSinceVersion()
}

func (*OrderCancelReplaceRequest515) SenderIDDeprecated() uint16 {
	return 0
}

func (*OrderCancelReplaceRequest515) SenderIDMetaAttribute(meta int) string {
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

func (*OrderCancelReplaceRequest515) SenderIDMinValue() byte {
	return byte(32)
}

func (*OrderCancelReplaceRequest515) SenderIDMaxValue() byte {
	return byte(126)
}

func (*OrderCancelReplaceRequest515) SenderIDNullValue() byte {
	return 0
}

func (o *OrderCancelReplaceRequest515) SenderIDCharacterEncoding() string {
	return "US-ASCII"
}

func (*OrderCancelReplaceRequest515) ClOrdIDId() uint16 {
	return 11
}

func (*OrderCancelReplaceRequest515) ClOrdIDSinceVersion() uint16 {
	return 0
}

func (o *OrderCancelReplaceRequest515) ClOrdIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.ClOrdIDSinceVersion()
}

func (*OrderCancelReplaceRequest515) ClOrdIDDeprecated() uint16 {
	return 0
}

func (*OrderCancelReplaceRequest515) ClOrdIDMetaAttribute(meta int) string {
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

func (*OrderCancelReplaceRequest515) ClOrdIDMinValue() byte {
	return byte(32)
}

func (*OrderCancelReplaceRequest515) ClOrdIDMaxValue() byte {
	return byte(126)
}

func (*OrderCancelReplaceRequest515) ClOrdIDNullValue() byte {
	return 0
}

func (o *OrderCancelReplaceRequest515) ClOrdIDCharacterEncoding() string {
	return "US-ASCII"
}

func (*OrderCancelReplaceRequest515) PartyDetailsListReqIDId() uint16 {
	return 1505
}

func (*OrderCancelReplaceRequest515) PartyDetailsListReqIDSinceVersion() uint16 {
	return 0
}

func (o *OrderCancelReplaceRequest515) PartyDetailsListReqIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.PartyDetailsListReqIDSinceVersion()
}

func (*OrderCancelReplaceRequest515) PartyDetailsListReqIDDeprecated() uint16 {
	return 0
}

func (*OrderCancelReplaceRequest515) PartyDetailsListReqIDMetaAttribute(meta int) string {
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

func (*OrderCancelReplaceRequest515) PartyDetailsListReqIDMinValue() uint64 {
	return 0
}

func (*OrderCancelReplaceRequest515) PartyDetailsListReqIDMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*OrderCancelReplaceRequest515) PartyDetailsListReqIDNullValue() uint64 {
	return math.MaxUint64
}

func (*OrderCancelReplaceRequest515) OrderIDId() uint16 {
	return 37
}

func (*OrderCancelReplaceRequest515) OrderIDSinceVersion() uint16 {
	return 0
}

func (o *OrderCancelReplaceRequest515) OrderIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.OrderIDSinceVersion()
}

func (*OrderCancelReplaceRequest515) OrderIDDeprecated() uint16 {
	return 0
}

func (*OrderCancelReplaceRequest515) OrderIDMetaAttribute(meta int) string {
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

func (*OrderCancelReplaceRequest515) OrderIDMinValue() uint64 {
	return 0
}

func (*OrderCancelReplaceRequest515) OrderIDMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*OrderCancelReplaceRequest515) OrderIDNullValue() uint64 {
	return math.MaxUint64
}

func (*OrderCancelReplaceRequest515) StopPxId() uint16 {
	return 99
}

func (*OrderCancelReplaceRequest515) StopPxSinceVersion() uint16 {
	return 0
}

func (o *OrderCancelReplaceRequest515) StopPxInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.StopPxSinceVersion()
}

func (*OrderCancelReplaceRequest515) StopPxDeprecated() uint16 {
	return 0
}

func (*OrderCancelReplaceRequest515) StopPxMetaAttribute(meta int) string {
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

func (*OrderCancelReplaceRequest515) OrderRequestIDId() uint16 {
	return 2422
}

func (*OrderCancelReplaceRequest515) OrderRequestIDSinceVersion() uint16 {
	return 0
}

func (o *OrderCancelReplaceRequest515) OrderRequestIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.OrderRequestIDSinceVersion()
}

func (*OrderCancelReplaceRequest515) OrderRequestIDDeprecated() uint16 {
	return 0
}

func (*OrderCancelReplaceRequest515) OrderRequestIDMetaAttribute(meta int) string {
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

func (*OrderCancelReplaceRequest515) OrderRequestIDMinValue() uint64 {
	return 0
}

func (*OrderCancelReplaceRequest515) OrderRequestIDMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*OrderCancelReplaceRequest515) OrderRequestIDNullValue() uint64 {
	return math.MaxUint64
}

func (*OrderCancelReplaceRequest515) SendingTimeEpochId() uint16 {
	return 5297
}

func (*OrderCancelReplaceRequest515) SendingTimeEpochSinceVersion() uint16 {
	return 0
}

func (o *OrderCancelReplaceRequest515) SendingTimeEpochInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.SendingTimeEpochSinceVersion()
}

func (*OrderCancelReplaceRequest515) SendingTimeEpochDeprecated() uint16 {
	return 0
}

func (*OrderCancelReplaceRequest515) SendingTimeEpochMetaAttribute(meta int) string {
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

func (*OrderCancelReplaceRequest515) SendingTimeEpochMinValue() uint64 {
	return 0
}

func (*OrderCancelReplaceRequest515) SendingTimeEpochMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*OrderCancelReplaceRequest515) SendingTimeEpochNullValue() uint64 {
	return math.MaxUint64
}

func (*OrderCancelReplaceRequest515) LocationId() uint16 {
	return 9537
}

func (*OrderCancelReplaceRequest515) LocationSinceVersion() uint16 {
	return 0
}

func (o *OrderCancelReplaceRequest515) LocationInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.LocationSinceVersion()
}

func (*OrderCancelReplaceRequest515) LocationDeprecated() uint16 {
	return 0
}

func (*OrderCancelReplaceRequest515) LocationMetaAttribute(meta int) string {
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

func (*OrderCancelReplaceRequest515) LocationMinValue() byte {
	return byte(32)
}

func (*OrderCancelReplaceRequest515) LocationMaxValue() byte {
	return byte(126)
}

func (*OrderCancelReplaceRequest515) LocationNullValue() byte {
	return 0
}

func (o *OrderCancelReplaceRequest515) LocationCharacterEncoding() string {
	return "US-ASCII"
}

func (*OrderCancelReplaceRequest515) MinQtyId() uint16 {
	return 110
}

func (*OrderCancelReplaceRequest515) MinQtySinceVersion() uint16 {
	return 0
}

func (o *OrderCancelReplaceRequest515) MinQtyInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.MinQtySinceVersion()
}

func (*OrderCancelReplaceRequest515) MinQtyDeprecated() uint16 {
	return 0
}

func (*OrderCancelReplaceRequest515) MinQtyMetaAttribute(meta int) string {
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

func (*OrderCancelReplaceRequest515) MinQtyMinValue() uint32 {
	return 0
}

func (*OrderCancelReplaceRequest515) MinQtyMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*OrderCancelReplaceRequest515) MinQtyNullValue() uint32 {
	return 4294967295
}

func (*OrderCancelReplaceRequest515) DisplayQtyId() uint16 {
	return 1138
}

func (*OrderCancelReplaceRequest515) DisplayQtySinceVersion() uint16 {
	return 0
}

func (o *OrderCancelReplaceRequest515) DisplayQtyInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.DisplayQtySinceVersion()
}

func (*OrderCancelReplaceRequest515) DisplayQtyDeprecated() uint16 {
	return 0
}

func (*OrderCancelReplaceRequest515) DisplayQtyMetaAttribute(meta int) string {
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

func (*OrderCancelReplaceRequest515) DisplayQtyMinValue() uint32 {
	return 0
}

func (*OrderCancelReplaceRequest515) DisplayQtyMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*OrderCancelReplaceRequest515) DisplayQtyNullValue() uint32 {
	return 4294967295
}

func (*OrderCancelReplaceRequest515) ExpireDateId() uint16 {
	return 432
}

func (*OrderCancelReplaceRequest515) ExpireDateSinceVersion() uint16 {
	return 0
}

func (o *OrderCancelReplaceRequest515) ExpireDateInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.ExpireDateSinceVersion()
}

func (*OrderCancelReplaceRequest515) ExpireDateDeprecated() uint16 {
	return 0
}

func (*OrderCancelReplaceRequest515) ExpireDateMetaAttribute(meta int) string {
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

func (*OrderCancelReplaceRequest515) ExpireDateMinValue() uint16 {
	return 0
}

func (*OrderCancelReplaceRequest515) ExpireDateMaxValue() uint16 {
	return math.MaxUint16 - 1
}

func (*OrderCancelReplaceRequest515) ExpireDateNullValue() uint16 {
	return 65535
}

func (*OrderCancelReplaceRequest515) OrdTypeId() uint16 {
	return 40
}

func (*OrderCancelReplaceRequest515) OrdTypeSinceVersion() uint16 {
	return 0
}

func (o *OrderCancelReplaceRequest515) OrdTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.OrdTypeSinceVersion()
}

func (*OrderCancelReplaceRequest515) OrdTypeDeprecated() uint16 {
	return 0
}

func (*OrderCancelReplaceRequest515) OrdTypeMetaAttribute(meta int) string {
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

func (*OrderCancelReplaceRequest515) TimeInForceId() uint16 {
	return 59
}

func (*OrderCancelReplaceRequest515) TimeInForceSinceVersion() uint16 {
	return 0
}

func (o *OrderCancelReplaceRequest515) TimeInForceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.TimeInForceSinceVersion()
}

func (*OrderCancelReplaceRequest515) TimeInForceDeprecated() uint16 {
	return 0
}

func (*OrderCancelReplaceRequest515) TimeInForceMetaAttribute(meta int) string {
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

func (*OrderCancelReplaceRequest515) ManualOrderIndicatorId() uint16 {
	return 1028
}

func (*OrderCancelReplaceRequest515) ManualOrderIndicatorSinceVersion() uint16 {
	return 0
}

func (o *OrderCancelReplaceRequest515) ManualOrderIndicatorInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.ManualOrderIndicatorSinceVersion()
}

func (*OrderCancelReplaceRequest515) ManualOrderIndicatorDeprecated() uint16 {
	return 0
}

func (*OrderCancelReplaceRequest515) ManualOrderIndicatorMetaAttribute(meta int) string {
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

func (*OrderCancelReplaceRequest515) OFMOverrideId() uint16 {
	return 9768
}

func (*OrderCancelReplaceRequest515) OFMOverrideSinceVersion() uint16 {
	return 0
}

func (o *OrderCancelReplaceRequest515) OFMOverrideInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.OFMOverrideSinceVersion()
}

func (*OrderCancelReplaceRequest515) OFMOverrideDeprecated() uint16 {
	return 0
}

func (*OrderCancelReplaceRequest515) OFMOverrideMetaAttribute(meta int) string {
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

func (*OrderCancelReplaceRequest515) ExecInstId() uint16 {
	return 18
}

func (*OrderCancelReplaceRequest515) ExecInstSinceVersion() uint16 {
	return 0
}

func (o *OrderCancelReplaceRequest515) ExecInstInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.ExecInstSinceVersion()
}

func (*OrderCancelReplaceRequest515) ExecInstDeprecated() uint16 {
	return 0
}

func (*OrderCancelReplaceRequest515) ExecInstMetaAttribute(meta int) string {
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

func (*OrderCancelReplaceRequest515) ExecutionModeId() uint16 {
	return 5906
}

func (*OrderCancelReplaceRequest515) ExecutionModeSinceVersion() uint16 {
	return 0
}

func (o *OrderCancelReplaceRequest515) ExecutionModeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.ExecutionModeSinceVersion()
}

func (*OrderCancelReplaceRequest515) ExecutionModeDeprecated() uint16 {
	return 0
}

func (*OrderCancelReplaceRequest515) ExecutionModeMetaAttribute(meta int) string {
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

func (*OrderCancelReplaceRequest515) LiquidityFlagId() uint16 {
	return 9373
}

func (*OrderCancelReplaceRequest515) LiquidityFlagSinceVersion() uint16 {
	return 0
}

func (o *OrderCancelReplaceRequest515) LiquidityFlagInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.LiquidityFlagSinceVersion()
}

func (*OrderCancelReplaceRequest515) LiquidityFlagDeprecated() uint16 {
	return 0
}

func (*OrderCancelReplaceRequest515) LiquidityFlagMetaAttribute(meta int) string {
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

func (*OrderCancelReplaceRequest515) ManagedOrderId() uint16 {
	return 6881
}

func (*OrderCancelReplaceRequest515) ManagedOrderSinceVersion() uint16 {
	return 0
}

func (o *OrderCancelReplaceRequest515) ManagedOrderInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.ManagedOrderSinceVersion()
}

func (*OrderCancelReplaceRequest515) ManagedOrderDeprecated() uint16 {
	return 0
}

func (*OrderCancelReplaceRequest515) ManagedOrderMetaAttribute(meta int) string {
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

func (*OrderCancelReplaceRequest515) ShortSaleTypeId() uint16 {
	return 5409
}

func (*OrderCancelReplaceRequest515) ShortSaleTypeSinceVersion() uint16 {
	return 0
}

func (o *OrderCancelReplaceRequest515) ShortSaleTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.ShortSaleTypeSinceVersion()
}

func (*OrderCancelReplaceRequest515) ShortSaleTypeDeprecated() uint16 {
	return 0
}

func (*OrderCancelReplaceRequest515) ShortSaleTypeMetaAttribute(meta int) string {
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
