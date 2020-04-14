// Generated SBE (Simple Binary Encoding) message codec

package ilinkbinary

import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
)

type OrderCancelRequest516 struct {
	OrderID               uint64
	PartyDetailsListReqID uint64
	ManualOrderIndicator  ManualOrdIndReqEnum
	SeqNum                uint32
	SenderID              [20]byte
	ClOrdID               [20]byte
	OrderRequestID        uint64
	SendingTimeEpoch      uint64
	Location              [5]byte
	SecurityID            int32
	Side                  SideReqEnum
	LiquidityFlag         BooleanNULLEnum
}

func (o *OrderCancelRequest516) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := o.RangeCheck(o.SbeSchemaVersion(), o.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	if err := _m.WriteUint64(_w, o.OrderID); err != nil {
		return err
	}
	if err := _m.WriteUint64(_w, o.PartyDetailsListReqID); err != nil {
		return err
	}
	if err := o.ManualOrderIndicator.Encode(_m, _w); err != nil {
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
	if err := _m.WriteUint64(_w, o.OrderRequestID); err != nil {
		return err
	}
	if err := _m.WriteUint64(_w, o.SendingTimeEpoch); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, o.Location[:]); err != nil {
		return err
	}
	if err := _m.WriteInt32(_w, o.SecurityID); err != nil {
		return err
	}
	if err := o.Side.Encode(_m, _w); err != nil {
		return err
	}
	if err := o.LiquidityFlag.Encode(_m, _w); err != nil {
		return err
	}
	return nil
}

func (o *OrderCancelRequest516) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	if !o.OrderIDInActingVersion(actingVersion) {
		o.OrderID = o.OrderIDNullValue()
	} else {
		if err := _m.ReadUint64(_r, &o.OrderID); err != nil {
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
	if o.ManualOrderIndicatorInActingVersion(actingVersion) {
		if err := o.ManualOrderIndicator.Decode(_m, _r, actingVersion); err != nil {
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
	if o.LiquidityFlagInActingVersion(actingVersion) {
		if err := o.LiquidityFlag.Decode(_m, _r, actingVersion); err != nil {
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

func (o *OrderCancelRequest516) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if o.OrderIDInActingVersion(actingVersion) {
		if o.OrderID < o.OrderIDMinValue() || o.OrderID > o.OrderIDMaxValue() {
			return fmt.Errorf("Range check failed on o.OrderID (%v < %v > %v)", o.OrderIDMinValue(), o.OrderID, o.OrderIDMaxValue())
		}
	}
	if o.PartyDetailsListReqIDInActingVersion(actingVersion) {
		if o.PartyDetailsListReqID < o.PartyDetailsListReqIDMinValue() || o.PartyDetailsListReqID > o.PartyDetailsListReqIDMaxValue() {
			return fmt.Errorf("Range check failed on o.PartyDetailsListReqID (%v < %v > %v)", o.PartyDetailsListReqIDMinValue(), o.PartyDetailsListReqID, o.PartyDetailsListReqIDMaxValue())
		}
	}
	if err := o.ManualOrderIndicator.RangeCheck(actingVersion, schemaVersion); err != nil {
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
	if o.SecurityIDInActingVersion(actingVersion) {
		if o.SecurityID < o.SecurityIDMinValue() || o.SecurityID > o.SecurityIDMaxValue() {
			return fmt.Errorf("Range check failed on o.SecurityID (%v < %v > %v)", o.SecurityIDMinValue(), o.SecurityID, o.SecurityIDMaxValue())
		}
	}
	if err := o.Side.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if err := o.LiquidityFlag.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	return nil
}

func OrderCancelRequest516Init(o *OrderCancelRequest516) {
	return
}

func (*OrderCancelRequest516) SbeBlockLength() (blockLength uint16) {
	return 88
}

func (*OrderCancelRequest516) SbeTemplateId() (templateId uint16) {
	return 516
}

func (*OrderCancelRequest516) SbeSchemaId() (schemaId uint16) {
	return 8
}

func (*OrderCancelRequest516) SbeSchemaVersion() (schemaVersion uint16) {
	return 5
}

func (*OrderCancelRequest516) SbeSemanticType() (semanticType []byte) {
	return []byte("F")
}

func (*OrderCancelRequest516) OrderIDId() uint16 {
	return 37
}

func (*OrderCancelRequest516) OrderIDSinceVersion() uint16 {
	return 0
}

func (o *OrderCancelRequest516) OrderIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.OrderIDSinceVersion()
}

func (*OrderCancelRequest516) OrderIDDeprecated() uint16 {
	return 0
}

func (*OrderCancelRequest516) OrderIDMetaAttribute(meta int) string {
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

func (*OrderCancelRequest516) OrderIDMinValue() uint64 {
	return 0
}

func (*OrderCancelRequest516) OrderIDMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*OrderCancelRequest516) OrderIDNullValue() uint64 {
	return math.MaxUint64
}

func (*OrderCancelRequest516) PartyDetailsListReqIDId() uint16 {
	return 1505
}

func (*OrderCancelRequest516) PartyDetailsListReqIDSinceVersion() uint16 {
	return 0
}

func (o *OrderCancelRequest516) PartyDetailsListReqIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.PartyDetailsListReqIDSinceVersion()
}

func (*OrderCancelRequest516) PartyDetailsListReqIDDeprecated() uint16 {
	return 0
}

func (*OrderCancelRequest516) PartyDetailsListReqIDMetaAttribute(meta int) string {
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

func (*OrderCancelRequest516) PartyDetailsListReqIDMinValue() uint64 {
	return 0
}

func (*OrderCancelRequest516) PartyDetailsListReqIDMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*OrderCancelRequest516) PartyDetailsListReqIDNullValue() uint64 {
	return math.MaxUint64
}

func (*OrderCancelRequest516) ManualOrderIndicatorId() uint16 {
	return 1028
}

func (*OrderCancelRequest516) ManualOrderIndicatorSinceVersion() uint16 {
	return 0
}

func (o *OrderCancelRequest516) ManualOrderIndicatorInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.ManualOrderIndicatorSinceVersion()
}

func (*OrderCancelRequest516) ManualOrderIndicatorDeprecated() uint16 {
	return 0
}

func (*OrderCancelRequest516) ManualOrderIndicatorMetaAttribute(meta int) string {
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

func (*OrderCancelRequest516) SeqNumId() uint16 {
	return 9726
}

func (*OrderCancelRequest516) SeqNumSinceVersion() uint16 {
	return 0
}

func (o *OrderCancelRequest516) SeqNumInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.SeqNumSinceVersion()
}

func (*OrderCancelRequest516) SeqNumDeprecated() uint16 {
	return 0
}

func (*OrderCancelRequest516) SeqNumMetaAttribute(meta int) string {
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

func (*OrderCancelRequest516) SeqNumMinValue() uint32 {
	return 0
}

func (*OrderCancelRequest516) SeqNumMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*OrderCancelRequest516) SeqNumNullValue() uint32 {
	return math.MaxUint32
}

func (*OrderCancelRequest516) SenderIDId() uint16 {
	return 5392
}

func (*OrderCancelRequest516) SenderIDSinceVersion() uint16 {
	return 0
}

func (o *OrderCancelRequest516) SenderIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.SenderIDSinceVersion()
}

func (*OrderCancelRequest516) SenderIDDeprecated() uint16 {
	return 0
}

func (*OrderCancelRequest516) SenderIDMetaAttribute(meta int) string {
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

func (*OrderCancelRequest516) SenderIDMinValue() byte {
	return byte(32)
}

func (*OrderCancelRequest516) SenderIDMaxValue() byte {
	return byte(126)
}

func (*OrderCancelRequest516) SenderIDNullValue() byte {
	return 0
}

func (o *OrderCancelRequest516) SenderIDCharacterEncoding() string {
	return "US-ASCII"
}

func (*OrderCancelRequest516) ClOrdIDId() uint16 {
	return 11
}

func (*OrderCancelRequest516) ClOrdIDSinceVersion() uint16 {
	return 0
}

func (o *OrderCancelRequest516) ClOrdIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.ClOrdIDSinceVersion()
}

func (*OrderCancelRequest516) ClOrdIDDeprecated() uint16 {
	return 0
}

func (*OrderCancelRequest516) ClOrdIDMetaAttribute(meta int) string {
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

func (*OrderCancelRequest516) ClOrdIDMinValue() byte {
	return byte(32)
}

func (*OrderCancelRequest516) ClOrdIDMaxValue() byte {
	return byte(126)
}

func (*OrderCancelRequest516) ClOrdIDNullValue() byte {
	return 0
}

func (o *OrderCancelRequest516) ClOrdIDCharacterEncoding() string {
	return "US-ASCII"
}

func (*OrderCancelRequest516) OrderRequestIDId() uint16 {
	return 2422
}

func (*OrderCancelRequest516) OrderRequestIDSinceVersion() uint16 {
	return 0
}

func (o *OrderCancelRequest516) OrderRequestIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.OrderRequestIDSinceVersion()
}

func (*OrderCancelRequest516) OrderRequestIDDeprecated() uint16 {
	return 0
}

func (*OrderCancelRequest516) OrderRequestIDMetaAttribute(meta int) string {
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

func (*OrderCancelRequest516) OrderRequestIDMinValue() uint64 {
	return 0
}

func (*OrderCancelRequest516) OrderRequestIDMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*OrderCancelRequest516) OrderRequestIDNullValue() uint64 {
	return math.MaxUint64
}

func (*OrderCancelRequest516) SendingTimeEpochId() uint16 {
	return 5297
}

func (*OrderCancelRequest516) SendingTimeEpochSinceVersion() uint16 {
	return 0
}

func (o *OrderCancelRequest516) SendingTimeEpochInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.SendingTimeEpochSinceVersion()
}

func (*OrderCancelRequest516) SendingTimeEpochDeprecated() uint16 {
	return 0
}

func (*OrderCancelRequest516) SendingTimeEpochMetaAttribute(meta int) string {
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

func (*OrderCancelRequest516) SendingTimeEpochMinValue() uint64 {
	return 0
}

func (*OrderCancelRequest516) SendingTimeEpochMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*OrderCancelRequest516) SendingTimeEpochNullValue() uint64 {
	return math.MaxUint64
}

func (*OrderCancelRequest516) LocationId() uint16 {
	return 9537
}

func (*OrderCancelRequest516) LocationSinceVersion() uint16 {
	return 0
}

func (o *OrderCancelRequest516) LocationInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.LocationSinceVersion()
}

func (*OrderCancelRequest516) LocationDeprecated() uint16 {
	return 0
}

func (*OrderCancelRequest516) LocationMetaAttribute(meta int) string {
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

func (*OrderCancelRequest516) LocationMinValue() byte {
	return byte(32)
}

func (*OrderCancelRequest516) LocationMaxValue() byte {
	return byte(126)
}

func (*OrderCancelRequest516) LocationNullValue() byte {
	return 0
}

func (o *OrderCancelRequest516) LocationCharacterEncoding() string {
	return "US-ASCII"
}

func (*OrderCancelRequest516) SecurityIDId() uint16 {
	return 48
}

func (*OrderCancelRequest516) SecurityIDSinceVersion() uint16 {
	return 0
}

func (o *OrderCancelRequest516) SecurityIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.SecurityIDSinceVersion()
}

func (*OrderCancelRequest516) SecurityIDDeprecated() uint16 {
	return 0
}

func (*OrderCancelRequest516) SecurityIDMetaAttribute(meta int) string {
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

func (*OrderCancelRequest516) SecurityIDMinValue() int32 {
	return math.MinInt32 + 1
}

func (*OrderCancelRequest516) SecurityIDMaxValue() int32 {
	return math.MaxInt32
}

func (*OrderCancelRequest516) SecurityIDNullValue() int32 {
	return math.MinInt32
}

func (*OrderCancelRequest516) SideId() uint16 {
	return 54
}

func (*OrderCancelRequest516) SideSinceVersion() uint16 {
	return 0
}

func (o *OrderCancelRequest516) SideInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.SideSinceVersion()
}

func (*OrderCancelRequest516) SideDeprecated() uint16 {
	return 0
}

func (*OrderCancelRequest516) SideMetaAttribute(meta int) string {
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

func (*OrderCancelRequest516) LiquidityFlagId() uint16 {
	return 9373
}

func (*OrderCancelRequest516) LiquidityFlagSinceVersion() uint16 {
	return 0
}

func (o *OrderCancelRequest516) LiquidityFlagInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.LiquidityFlagSinceVersion()
}

func (*OrderCancelRequest516) LiquidityFlagDeprecated() uint16 {
	return 0
}

func (*OrderCancelRequest516) LiquidityFlagMetaAttribute(meta int) string {
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
