// Generated SBE (Simple Binary Encoding) message codec

package ilinkbinary

import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
)

type OrderCancelReplaceReject536 struct {
	SeqNum                uint32
	UUID                  uint64
	Text                  [256]byte
	ExecID                [40]byte
	SenderID              [20]byte
	ClOrdID               [20]byte
	PartyDetailsListReqID uint64
	OrderID               uint64
	TransactTime          uint64
	SendingTimeEpoch      uint64
	OrderRequestID        uint64
	Location              [5]byte
	CxlRejReason          uint16
	DelayDuration         uint16
	OrdStatus             [1]byte
	CxlRejResponseTo      [1]byte
	ManualOrderIndicator  ManualOrdIndReqEnum
	PossRetransFlag       BooleanFlagEnum
	SplitMsg              SplitMsgEnum
	LiquidityFlag         BooleanNULLEnum
	DelayToTime           uint64
}

func (o *OrderCancelReplaceReject536) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := o.RangeCheck(o.SbeSchemaVersion(), o.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	if err := _m.WriteUint32(_w, o.SeqNum); err != nil {
		return err
	}
	if err := _m.WriteUint64(_w, o.UUID); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, o.Text[:]); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, o.ExecID[:]); err != nil {
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
	if err := _m.WriteUint64(_w, o.TransactTime); err != nil {
		return err
	}
	if err := _m.WriteUint64(_w, o.SendingTimeEpoch); err != nil {
		return err
	}
	if err := _m.WriteUint64(_w, o.OrderRequestID); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, o.Location[:]); err != nil {
		return err
	}
	if err := _m.WriteUint16(_w, o.CxlRejReason); err != nil {
		return err
	}
	if err := _m.WriteUint16(_w, o.DelayDuration); err != nil {
		return err
	}
	if err := o.ManualOrderIndicator.Encode(_m, _w); err != nil {
		return err
	}
	if err := o.PossRetransFlag.Encode(_m, _w); err != nil {
		return err
	}
	if err := o.SplitMsg.Encode(_m, _w); err != nil {
		return err
	}
	if err := o.LiquidityFlag.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteUint64(_w, o.DelayToTime); err != nil {
		return err
	}
	return nil
}

func (o *OrderCancelReplaceReject536) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	if !o.SeqNumInActingVersion(actingVersion) {
		o.SeqNum = o.SeqNumNullValue()
	} else {
		if err := _m.ReadUint32(_r, &o.SeqNum); err != nil {
			return err
		}
	}
	if !o.UUIDInActingVersion(actingVersion) {
		o.UUID = o.UUIDNullValue()
	} else {
		if err := _m.ReadUint64(_r, &o.UUID); err != nil {
			return err
		}
	}
	if !o.TextInActingVersion(actingVersion) {
		for idx := 0; idx < 256; idx++ {
			o.Text[idx] = o.TextNullValue()
		}
	} else {
		if err := _m.ReadBytes(_r, o.Text[:]); err != nil {
			return err
		}
	}
	if !o.ExecIDInActingVersion(actingVersion) {
		for idx := 0; idx < 40; idx++ {
			o.ExecID[idx] = o.ExecIDNullValue()
		}
	} else {
		if err := _m.ReadBytes(_r, o.ExecID[:]); err != nil {
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
	if !o.TransactTimeInActingVersion(actingVersion) {
		o.TransactTime = o.TransactTimeNullValue()
	} else {
		if err := _m.ReadUint64(_r, &o.TransactTime); err != nil {
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
	if !o.OrderRequestIDInActingVersion(actingVersion) {
		o.OrderRequestID = o.OrderRequestIDNullValue()
	} else {
		if err := _m.ReadUint64(_r, &o.OrderRequestID); err != nil {
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
	if !o.CxlRejReasonInActingVersion(actingVersion) {
		o.CxlRejReason = o.CxlRejReasonNullValue()
	} else {
		if err := _m.ReadUint16(_r, &o.CxlRejReason); err != nil {
			return err
		}
	}
	if !o.DelayDurationInActingVersion(actingVersion) {
		o.DelayDuration = o.DelayDurationNullValue()
	} else {
		if err := _m.ReadUint16(_r, &o.DelayDuration); err != nil {
			return err
		}
	}
	o.OrdStatus[0] = 85
	o.CxlRejResponseTo[0] = 50
	if o.ManualOrderIndicatorInActingVersion(actingVersion) {
		if err := o.ManualOrderIndicator.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if o.PossRetransFlagInActingVersion(actingVersion) {
		if err := o.PossRetransFlag.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if o.SplitMsgInActingVersion(actingVersion) {
		if err := o.SplitMsg.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if o.LiquidityFlagInActingVersion(actingVersion) {
		if err := o.LiquidityFlag.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if !o.DelayToTimeInActingVersion(actingVersion) {
		o.DelayToTime = o.DelayToTimeNullValue()
	} else {
		if err := _m.ReadUint64(_r, &o.DelayToTime); err != nil {
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

func (o *OrderCancelReplaceReject536) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if o.SeqNumInActingVersion(actingVersion) {
		if o.SeqNum < o.SeqNumMinValue() || o.SeqNum > o.SeqNumMaxValue() {
			return fmt.Errorf("Range check failed on o.SeqNum (%v < %v > %v)", o.SeqNumMinValue(), o.SeqNum, o.SeqNumMaxValue())
		}
	}
	if o.UUIDInActingVersion(actingVersion) {
		if o.UUID < o.UUIDMinValue() || o.UUID > o.UUIDMaxValue() {
			return fmt.Errorf("Range check failed on o.UUID (%v < %v > %v)", o.UUIDMinValue(), o.UUID, o.UUIDMaxValue())
		}
	}
	if o.TextInActingVersion(actingVersion) {
		for idx := 0; idx < 256; idx++ {
			if o.Text[idx] < o.TextMinValue() || o.Text[idx] > o.TextMaxValue() {
				return fmt.Errorf("Range check failed on o.Text[%d] (%v < %v > %v)", idx, o.TextMinValue(), o.Text[idx], o.TextMaxValue())
			}
		}
	}
	if o.ExecIDInActingVersion(actingVersion) {
		for idx := 0; idx < 40; idx++ {
			if o.ExecID[idx] < o.ExecIDMinValue() || o.ExecID[idx] > o.ExecIDMaxValue() {
				return fmt.Errorf("Range check failed on o.ExecID[%d] (%v < %v > %v)", idx, o.ExecIDMinValue(), o.ExecID[idx], o.ExecIDMaxValue())
			}
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
	if o.TransactTimeInActingVersion(actingVersion) {
		if o.TransactTime < o.TransactTimeMinValue() || o.TransactTime > o.TransactTimeMaxValue() {
			return fmt.Errorf("Range check failed on o.TransactTime (%v < %v > %v)", o.TransactTimeMinValue(), o.TransactTime, o.TransactTimeMaxValue())
		}
	}
	if o.SendingTimeEpochInActingVersion(actingVersion) {
		if o.SendingTimeEpoch < o.SendingTimeEpochMinValue() || o.SendingTimeEpoch > o.SendingTimeEpochMaxValue() {
			return fmt.Errorf("Range check failed on o.SendingTimeEpoch (%v < %v > %v)", o.SendingTimeEpochMinValue(), o.SendingTimeEpoch, o.SendingTimeEpochMaxValue())
		}
	}
	if o.OrderRequestIDInActingVersion(actingVersion) {
		if o.OrderRequestID < o.OrderRequestIDMinValue() || o.OrderRequestID > o.OrderRequestIDMaxValue() {
			return fmt.Errorf("Range check failed on o.OrderRequestID (%v < %v > %v)", o.OrderRequestIDMinValue(), o.OrderRequestID, o.OrderRequestIDMaxValue())
		}
	}
	if o.LocationInActingVersion(actingVersion) {
		for idx := 0; idx < 5; idx++ {
			if o.Location[idx] < o.LocationMinValue() || o.Location[idx] > o.LocationMaxValue() {
				return fmt.Errorf("Range check failed on o.Location[%d] (%v < %v > %v)", idx, o.LocationMinValue(), o.Location[idx], o.LocationMaxValue())
			}
		}
	}
	if o.CxlRejReasonInActingVersion(actingVersion) {
		if o.CxlRejReason < o.CxlRejReasonMinValue() || o.CxlRejReason > o.CxlRejReasonMaxValue() {
			return fmt.Errorf("Range check failed on o.CxlRejReason (%v < %v > %v)", o.CxlRejReasonMinValue(), o.CxlRejReason, o.CxlRejReasonMaxValue())
		}
	}
	if o.DelayDurationInActingVersion(actingVersion) {
		if o.DelayDuration != o.DelayDurationNullValue() && (o.DelayDuration < o.DelayDurationMinValue() || o.DelayDuration > o.DelayDurationMaxValue()) {
			return fmt.Errorf("Range check failed on o.DelayDuration (%v < %v > %v)", o.DelayDurationMinValue(), o.DelayDuration, o.DelayDurationMaxValue())
		}
	}
	if err := o.ManualOrderIndicator.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if err := o.PossRetransFlag.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if err := o.SplitMsg.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if err := o.LiquidityFlag.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if o.DelayToTimeInActingVersion(actingVersion) {
		if o.DelayToTime != o.DelayToTimeNullValue() && (o.DelayToTime < o.DelayToTimeMinValue() || o.DelayToTime > o.DelayToTimeMaxValue()) {
			return fmt.Errorf("Range check failed on o.DelayToTime (%v < %v > %v)", o.DelayToTimeMinValue(), o.DelayToTime, o.DelayToTimeMaxValue())
		}
	}
	return nil
}

func OrderCancelReplaceReject536Init(o *OrderCancelReplaceReject536) {
	for idx := 0; idx < 256; idx++ {
		o.Text[idx] = 0
	}
	o.DelayDuration = 65535
	o.OrdStatus[0] = 85
	o.CxlRejResponseTo[0] = 50
	o.DelayToTime = 18446744073709551615
	return
}

func (*OrderCancelReplaceReject536) SbeBlockLength() (blockLength uint16) {
	return 409
}

func (*OrderCancelReplaceReject536) SbeTemplateId() (templateId uint16) {
	return 536
}

func (*OrderCancelReplaceReject536) SbeSchemaId() (schemaId uint16) {
	return 8
}

func (*OrderCancelReplaceReject536) SbeSchemaVersion() (schemaVersion uint16) {
	return 5
}

func (*OrderCancelReplaceReject536) SbeSemanticType() (semanticType []byte) {
	return []byte("9")
}

func (*OrderCancelReplaceReject536) SeqNumId() uint16 {
	return 9726
}

func (*OrderCancelReplaceReject536) SeqNumSinceVersion() uint16 {
	return 0
}

func (o *OrderCancelReplaceReject536) SeqNumInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.SeqNumSinceVersion()
}

func (*OrderCancelReplaceReject536) SeqNumDeprecated() uint16 {
	return 0
}

func (*OrderCancelReplaceReject536) SeqNumMetaAttribute(meta int) string {
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

func (*OrderCancelReplaceReject536) SeqNumMinValue() uint32 {
	return 0
}

func (*OrderCancelReplaceReject536) SeqNumMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*OrderCancelReplaceReject536) SeqNumNullValue() uint32 {
	return math.MaxUint32
}

func (*OrderCancelReplaceReject536) UUIDId() uint16 {
	return 39001
}

func (*OrderCancelReplaceReject536) UUIDSinceVersion() uint16 {
	return 0
}

func (o *OrderCancelReplaceReject536) UUIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.UUIDSinceVersion()
}

func (*OrderCancelReplaceReject536) UUIDDeprecated() uint16 {
	return 0
}

func (*OrderCancelReplaceReject536) UUIDMetaAttribute(meta int) string {
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

func (*OrderCancelReplaceReject536) UUIDMinValue() uint64 {
	return 0
}

func (*OrderCancelReplaceReject536) UUIDMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*OrderCancelReplaceReject536) UUIDNullValue() uint64 {
	return math.MaxUint64
}

func (*OrderCancelReplaceReject536) TextId() uint16 {
	return 58
}

func (*OrderCancelReplaceReject536) TextSinceVersion() uint16 {
	return 0
}

func (o *OrderCancelReplaceReject536) TextInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.TextSinceVersion()
}

func (*OrderCancelReplaceReject536) TextDeprecated() uint16 {
	return 0
}

func (*OrderCancelReplaceReject536) TextMetaAttribute(meta int) string {
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

func (*OrderCancelReplaceReject536) TextMinValue() byte {
	return byte(32)
}

func (*OrderCancelReplaceReject536) TextMaxValue() byte {
	return byte(126)
}

func (*OrderCancelReplaceReject536) TextNullValue() byte {
	return 0
}

func (o *OrderCancelReplaceReject536) TextCharacterEncoding() string {
	return "US-ASCII"
}

func (*OrderCancelReplaceReject536) ExecIDId() uint16 {
	return 17
}

func (*OrderCancelReplaceReject536) ExecIDSinceVersion() uint16 {
	return 0
}

func (o *OrderCancelReplaceReject536) ExecIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.ExecIDSinceVersion()
}

func (*OrderCancelReplaceReject536) ExecIDDeprecated() uint16 {
	return 0
}

func (*OrderCancelReplaceReject536) ExecIDMetaAttribute(meta int) string {
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

func (*OrderCancelReplaceReject536) ExecIDMinValue() byte {
	return byte(32)
}

func (*OrderCancelReplaceReject536) ExecIDMaxValue() byte {
	return byte(126)
}

func (*OrderCancelReplaceReject536) ExecIDNullValue() byte {
	return 0
}

func (o *OrderCancelReplaceReject536) ExecIDCharacterEncoding() string {
	return "US-ASCII"
}

func (*OrderCancelReplaceReject536) SenderIDId() uint16 {
	return 5392
}

func (*OrderCancelReplaceReject536) SenderIDSinceVersion() uint16 {
	return 0
}

func (o *OrderCancelReplaceReject536) SenderIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.SenderIDSinceVersion()
}

func (*OrderCancelReplaceReject536) SenderIDDeprecated() uint16 {
	return 0
}

func (*OrderCancelReplaceReject536) SenderIDMetaAttribute(meta int) string {
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

func (*OrderCancelReplaceReject536) SenderIDMinValue() byte {
	return byte(32)
}

func (*OrderCancelReplaceReject536) SenderIDMaxValue() byte {
	return byte(126)
}

func (*OrderCancelReplaceReject536) SenderIDNullValue() byte {
	return 0
}

func (o *OrderCancelReplaceReject536) SenderIDCharacterEncoding() string {
	return "US-ASCII"
}

func (*OrderCancelReplaceReject536) ClOrdIDId() uint16 {
	return 11
}

func (*OrderCancelReplaceReject536) ClOrdIDSinceVersion() uint16 {
	return 0
}

func (o *OrderCancelReplaceReject536) ClOrdIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.ClOrdIDSinceVersion()
}

func (*OrderCancelReplaceReject536) ClOrdIDDeprecated() uint16 {
	return 0
}

func (*OrderCancelReplaceReject536) ClOrdIDMetaAttribute(meta int) string {
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

func (*OrderCancelReplaceReject536) ClOrdIDMinValue() byte {
	return byte(32)
}

func (*OrderCancelReplaceReject536) ClOrdIDMaxValue() byte {
	return byte(126)
}

func (*OrderCancelReplaceReject536) ClOrdIDNullValue() byte {
	return 0
}

func (o *OrderCancelReplaceReject536) ClOrdIDCharacterEncoding() string {
	return "US-ASCII"
}

func (*OrderCancelReplaceReject536) PartyDetailsListReqIDId() uint16 {
	return 1505
}

func (*OrderCancelReplaceReject536) PartyDetailsListReqIDSinceVersion() uint16 {
	return 0
}

func (o *OrderCancelReplaceReject536) PartyDetailsListReqIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.PartyDetailsListReqIDSinceVersion()
}

func (*OrderCancelReplaceReject536) PartyDetailsListReqIDDeprecated() uint16 {
	return 0
}

func (*OrderCancelReplaceReject536) PartyDetailsListReqIDMetaAttribute(meta int) string {
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

func (*OrderCancelReplaceReject536) PartyDetailsListReqIDMinValue() uint64 {
	return 0
}

func (*OrderCancelReplaceReject536) PartyDetailsListReqIDMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*OrderCancelReplaceReject536) PartyDetailsListReqIDNullValue() uint64 {
	return math.MaxUint64
}

func (*OrderCancelReplaceReject536) OrderIDId() uint16 {
	return 37
}

func (*OrderCancelReplaceReject536) OrderIDSinceVersion() uint16 {
	return 0
}

func (o *OrderCancelReplaceReject536) OrderIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.OrderIDSinceVersion()
}

func (*OrderCancelReplaceReject536) OrderIDDeprecated() uint16 {
	return 0
}

func (*OrderCancelReplaceReject536) OrderIDMetaAttribute(meta int) string {
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

func (*OrderCancelReplaceReject536) OrderIDMinValue() uint64 {
	return 0
}

func (*OrderCancelReplaceReject536) OrderIDMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*OrderCancelReplaceReject536) OrderIDNullValue() uint64 {
	return math.MaxUint64
}

func (*OrderCancelReplaceReject536) TransactTimeId() uint16 {
	return 60
}

func (*OrderCancelReplaceReject536) TransactTimeSinceVersion() uint16 {
	return 0
}

func (o *OrderCancelReplaceReject536) TransactTimeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.TransactTimeSinceVersion()
}

func (*OrderCancelReplaceReject536) TransactTimeDeprecated() uint16 {
	return 0
}

func (*OrderCancelReplaceReject536) TransactTimeMetaAttribute(meta int) string {
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

func (*OrderCancelReplaceReject536) TransactTimeMinValue() uint64 {
	return 0
}

func (*OrderCancelReplaceReject536) TransactTimeMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*OrderCancelReplaceReject536) TransactTimeNullValue() uint64 {
	return math.MaxUint64
}

func (*OrderCancelReplaceReject536) SendingTimeEpochId() uint16 {
	return 5297
}

func (*OrderCancelReplaceReject536) SendingTimeEpochSinceVersion() uint16 {
	return 0
}

func (o *OrderCancelReplaceReject536) SendingTimeEpochInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.SendingTimeEpochSinceVersion()
}

func (*OrderCancelReplaceReject536) SendingTimeEpochDeprecated() uint16 {
	return 0
}

func (*OrderCancelReplaceReject536) SendingTimeEpochMetaAttribute(meta int) string {
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

func (*OrderCancelReplaceReject536) SendingTimeEpochMinValue() uint64 {
	return 0
}

func (*OrderCancelReplaceReject536) SendingTimeEpochMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*OrderCancelReplaceReject536) SendingTimeEpochNullValue() uint64 {
	return math.MaxUint64
}

func (*OrderCancelReplaceReject536) OrderRequestIDId() uint16 {
	return 2422
}

func (*OrderCancelReplaceReject536) OrderRequestIDSinceVersion() uint16 {
	return 0
}

func (o *OrderCancelReplaceReject536) OrderRequestIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.OrderRequestIDSinceVersion()
}

func (*OrderCancelReplaceReject536) OrderRequestIDDeprecated() uint16 {
	return 0
}

func (*OrderCancelReplaceReject536) OrderRequestIDMetaAttribute(meta int) string {
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

func (*OrderCancelReplaceReject536) OrderRequestIDMinValue() uint64 {
	return 0
}

func (*OrderCancelReplaceReject536) OrderRequestIDMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*OrderCancelReplaceReject536) OrderRequestIDNullValue() uint64 {
	return math.MaxUint64
}

func (*OrderCancelReplaceReject536) LocationId() uint16 {
	return 9537
}

func (*OrderCancelReplaceReject536) LocationSinceVersion() uint16 {
	return 0
}

func (o *OrderCancelReplaceReject536) LocationInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.LocationSinceVersion()
}

func (*OrderCancelReplaceReject536) LocationDeprecated() uint16 {
	return 0
}

func (*OrderCancelReplaceReject536) LocationMetaAttribute(meta int) string {
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

func (*OrderCancelReplaceReject536) LocationMinValue() byte {
	return byte(32)
}

func (*OrderCancelReplaceReject536) LocationMaxValue() byte {
	return byte(126)
}

func (*OrderCancelReplaceReject536) LocationNullValue() byte {
	return 0
}

func (o *OrderCancelReplaceReject536) LocationCharacterEncoding() string {
	return "US-ASCII"
}

func (*OrderCancelReplaceReject536) CxlRejReasonId() uint16 {
	return 102
}

func (*OrderCancelReplaceReject536) CxlRejReasonSinceVersion() uint16 {
	return 0
}

func (o *OrderCancelReplaceReject536) CxlRejReasonInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.CxlRejReasonSinceVersion()
}

func (*OrderCancelReplaceReject536) CxlRejReasonDeprecated() uint16 {
	return 0
}

func (*OrderCancelReplaceReject536) CxlRejReasonMetaAttribute(meta int) string {
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

func (*OrderCancelReplaceReject536) CxlRejReasonMinValue() uint16 {
	return 0
}

func (*OrderCancelReplaceReject536) CxlRejReasonMaxValue() uint16 {
	return math.MaxUint16 - 1
}

func (*OrderCancelReplaceReject536) CxlRejReasonNullValue() uint16 {
	return math.MaxUint16
}

func (*OrderCancelReplaceReject536) DelayDurationId() uint16 {
	return 5904
}

func (*OrderCancelReplaceReject536) DelayDurationSinceVersion() uint16 {
	return 0
}

func (o *OrderCancelReplaceReject536) DelayDurationInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.DelayDurationSinceVersion()
}

func (*OrderCancelReplaceReject536) DelayDurationDeprecated() uint16 {
	return 0
}

func (*OrderCancelReplaceReject536) DelayDurationMetaAttribute(meta int) string {
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

func (*OrderCancelReplaceReject536) DelayDurationMinValue() uint16 {
	return 0
}

func (*OrderCancelReplaceReject536) DelayDurationMaxValue() uint16 {
	return math.MaxUint16 - 1
}

func (*OrderCancelReplaceReject536) DelayDurationNullValue() uint16 {
	return 65535
}

func (*OrderCancelReplaceReject536) OrdStatusId() uint16 {
	return 39
}

func (*OrderCancelReplaceReject536) OrdStatusSinceVersion() uint16 {
	return 0
}

func (o *OrderCancelReplaceReject536) OrdStatusInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.OrdStatusSinceVersion()
}

func (*OrderCancelReplaceReject536) OrdStatusDeprecated() uint16 {
	return 0
}

func (*OrderCancelReplaceReject536) OrdStatusMetaAttribute(meta int) string {
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

func (*OrderCancelReplaceReject536) OrdStatusMinValue() byte {
	return byte(32)
}

func (*OrderCancelReplaceReject536) OrdStatusMaxValue() byte {
	return byte(126)
}

func (*OrderCancelReplaceReject536) OrdStatusNullValue() byte {
	return 0
}

func (*OrderCancelReplaceReject536) CxlRejResponseToId() uint16 {
	return 434
}

func (*OrderCancelReplaceReject536) CxlRejResponseToSinceVersion() uint16 {
	return 0
}

func (o *OrderCancelReplaceReject536) CxlRejResponseToInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.CxlRejResponseToSinceVersion()
}

func (*OrderCancelReplaceReject536) CxlRejResponseToDeprecated() uint16 {
	return 0
}

func (*OrderCancelReplaceReject536) CxlRejResponseToMetaAttribute(meta int) string {
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

func (*OrderCancelReplaceReject536) CxlRejResponseToMinValue() byte {
	return byte(32)
}

func (*OrderCancelReplaceReject536) CxlRejResponseToMaxValue() byte {
	return byte(126)
}

func (*OrderCancelReplaceReject536) CxlRejResponseToNullValue() byte {
	return 0
}

func (*OrderCancelReplaceReject536) ManualOrderIndicatorId() uint16 {
	return 1028
}

func (*OrderCancelReplaceReject536) ManualOrderIndicatorSinceVersion() uint16 {
	return 0
}

func (o *OrderCancelReplaceReject536) ManualOrderIndicatorInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.ManualOrderIndicatorSinceVersion()
}

func (*OrderCancelReplaceReject536) ManualOrderIndicatorDeprecated() uint16 {
	return 0
}

func (*OrderCancelReplaceReject536) ManualOrderIndicatorMetaAttribute(meta int) string {
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

func (*OrderCancelReplaceReject536) PossRetransFlagId() uint16 {
	return 9765
}

func (*OrderCancelReplaceReject536) PossRetransFlagSinceVersion() uint16 {
	return 0
}

func (o *OrderCancelReplaceReject536) PossRetransFlagInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.PossRetransFlagSinceVersion()
}

func (*OrderCancelReplaceReject536) PossRetransFlagDeprecated() uint16 {
	return 0
}

func (*OrderCancelReplaceReject536) PossRetransFlagMetaAttribute(meta int) string {
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

func (*OrderCancelReplaceReject536) SplitMsgId() uint16 {
	return 9553
}

func (*OrderCancelReplaceReject536) SplitMsgSinceVersion() uint16 {
	return 0
}

func (o *OrderCancelReplaceReject536) SplitMsgInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.SplitMsgSinceVersion()
}

func (*OrderCancelReplaceReject536) SplitMsgDeprecated() uint16 {
	return 0
}

func (*OrderCancelReplaceReject536) SplitMsgMetaAttribute(meta int) string {
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

func (*OrderCancelReplaceReject536) LiquidityFlagId() uint16 {
	return 9373
}

func (*OrderCancelReplaceReject536) LiquidityFlagSinceVersion() uint16 {
	return 0
}

func (o *OrderCancelReplaceReject536) LiquidityFlagInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.LiquidityFlagSinceVersion()
}

func (*OrderCancelReplaceReject536) LiquidityFlagDeprecated() uint16 {
	return 0
}

func (*OrderCancelReplaceReject536) LiquidityFlagMetaAttribute(meta int) string {
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

func (*OrderCancelReplaceReject536) DelayToTimeId() uint16 {
	return 7552
}

func (*OrderCancelReplaceReject536) DelayToTimeSinceVersion() uint16 {
	return 4
}

func (o *OrderCancelReplaceReject536) DelayToTimeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.DelayToTimeSinceVersion()
}

func (*OrderCancelReplaceReject536) DelayToTimeDeprecated() uint16 {
	return 0
}

func (*OrderCancelReplaceReject536) DelayToTimeMetaAttribute(meta int) string {
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

func (*OrderCancelReplaceReject536) DelayToTimeMinValue() uint64 {
	return 0
}

func (*OrderCancelReplaceReject536) DelayToTimeMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*OrderCancelReplaceReject536) DelayToTimeNullValue() uint64 {
	return 18446744073709551615
}
