// Generated SBE (Simple Binary Encoding) message codec

package ilinkbinary

import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
)

type OrderCancelReject535 struct {
	SeqNum                uint32
	UUID                  uint64
	Text                  [256]byte
	ExecID                [40]byte
	SenderID              [20]byte
	ClOrdID               [20]byte
	PartyDetailsListReqID uint64
	OrdStatus             [1]byte
	CxlRejResponseTo      [1]byte
	OrderID               uint64
	TransactTime          uint64
	SendingTimeEpoch      uint64
	OrderRequestID        uint64
	Location              [5]byte
	CxlRejReason          uint16
	DelayDuration         uint16
	ManualOrderIndicator  ManualOrdIndReqEnum
	PossRetransFlag       BooleanFlagEnum
	SplitMsg              SplitMsgEnum
	LiquidityFlag         BooleanNULLEnum
	DelayToTime           uint64
}

func (o *OrderCancelReject535) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
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

func (o *OrderCancelReject535) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
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
	o.OrdStatus[0] = 85
	o.CxlRejResponseTo[0] = 49
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

func (o *OrderCancelReject535) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
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

func OrderCancelReject535Init(o *OrderCancelReject535) {
	for idx := 0; idx < 256; idx++ {
		o.Text[idx] = 0
	}
	o.OrdStatus[0] = 85
	o.CxlRejResponseTo[0] = 49
	o.DelayDuration = 65535
	o.DelayToTime = 18446744073709551615
	return
}

func (*OrderCancelReject535) SbeBlockLength() (blockLength uint16) {
	return 409
}

func (*OrderCancelReject535) SbeTemplateId() (templateId uint16) {
	return 535
}

func (*OrderCancelReject535) SbeSchemaId() (schemaId uint16) {
	return 8
}

func (*OrderCancelReject535) SbeSchemaVersion() (schemaVersion uint16) {
	return 5
}

func (*OrderCancelReject535) SbeSemanticType() (semanticType []byte) {
	return []byte("9")
}

func (*OrderCancelReject535) SeqNumId() uint16 {
	return 9726
}

func (*OrderCancelReject535) SeqNumSinceVersion() uint16 {
	return 0
}

func (o *OrderCancelReject535) SeqNumInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.SeqNumSinceVersion()
}

func (*OrderCancelReject535) SeqNumDeprecated() uint16 {
	return 0
}

func (*OrderCancelReject535) SeqNumMetaAttribute(meta int) string {
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

func (*OrderCancelReject535) SeqNumMinValue() uint32 {
	return 0
}

func (*OrderCancelReject535) SeqNumMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*OrderCancelReject535) SeqNumNullValue() uint32 {
	return math.MaxUint32
}

func (*OrderCancelReject535) UUIDId() uint16 {
	return 39001
}

func (*OrderCancelReject535) UUIDSinceVersion() uint16 {
	return 0
}

func (o *OrderCancelReject535) UUIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.UUIDSinceVersion()
}

func (*OrderCancelReject535) UUIDDeprecated() uint16 {
	return 0
}

func (*OrderCancelReject535) UUIDMetaAttribute(meta int) string {
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

func (*OrderCancelReject535) UUIDMinValue() uint64 {
	return 0
}

func (*OrderCancelReject535) UUIDMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*OrderCancelReject535) UUIDNullValue() uint64 {
	return math.MaxUint64
}

func (*OrderCancelReject535) TextId() uint16 {
	return 58
}

func (*OrderCancelReject535) TextSinceVersion() uint16 {
	return 0
}

func (o *OrderCancelReject535) TextInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.TextSinceVersion()
}

func (*OrderCancelReject535) TextDeprecated() uint16 {
	return 0
}

func (*OrderCancelReject535) TextMetaAttribute(meta int) string {
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

func (*OrderCancelReject535) TextMinValue() byte {
	return byte(32)
}

func (*OrderCancelReject535) TextMaxValue() byte {
	return byte(126)
}

func (*OrderCancelReject535) TextNullValue() byte {
	return 0
}

func (o *OrderCancelReject535) TextCharacterEncoding() string {
	return "US-ASCII"
}

func (*OrderCancelReject535) ExecIDId() uint16 {
	return 17
}

func (*OrderCancelReject535) ExecIDSinceVersion() uint16 {
	return 0
}

func (o *OrderCancelReject535) ExecIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.ExecIDSinceVersion()
}

func (*OrderCancelReject535) ExecIDDeprecated() uint16 {
	return 0
}

func (*OrderCancelReject535) ExecIDMetaAttribute(meta int) string {
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

func (*OrderCancelReject535) ExecIDMinValue() byte {
	return byte(32)
}

func (*OrderCancelReject535) ExecIDMaxValue() byte {
	return byte(126)
}

func (*OrderCancelReject535) ExecIDNullValue() byte {
	return 0
}

func (o *OrderCancelReject535) ExecIDCharacterEncoding() string {
	return "US-ASCII"
}

func (*OrderCancelReject535) SenderIDId() uint16 {
	return 5392
}

func (*OrderCancelReject535) SenderIDSinceVersion() uint16 {
	return 0
}

func (o *OrderCancelReject535) SenderIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.SenderIDSinceVersion()
}

func (*OrderCancelReject535) SenderIDDeprecated() uint16 {
	return 0
}

func (*OrderCancelReject535) SenderIDMetaAttribute(meta int) string {
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

func (*OrderCancelReject535) SenderIDMinValue() byte {
	return byte(32)
}

func (*OrderCancelReject535) SenderIDMaxValue() byte {
	return byte(126)
}

func (*OrderCancelReject535) SenderIDNullValue() byte {
	return 0
}

func (o *OrderCancelReject535) SenderIDCharacterEncoding() string {
	return "US-ASCII"
}

func (*OrderCancelReject535) ClOrdIDId() uint16 {
	return 11
}

func (*OrderCancelReject535) ClOrdIDSinceVersion() uint16 {
	return 0
}

func (o *OrderCancelReject535) ClOrdIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.ClOrdIDSinceVersion()
}

func (*OrderCancelReject535) ClOrdIDDeprecated() uint16 {
	return 0
}

func (*OrderCancelReject535) ClOrdIDMetaAttribute(meta int) string {
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

func (*OrderCancelReject535) ClOrdIDMinValue() byte {
	return byte(32)
}

func (*OrderCancelReject535) ClOrdIDMaxValue() byte {
	return byte(126)
}

func (*OrderCancelReject535) ClOrdIDNullValue() byte {
	return 0
}

func (o *OrderCancelReject535) ClOrdIDCharacterEncoding() string {
	return "US-ASCII"
}

func (*OrderCancelReject535) PartyDetailsListReqIDId() uint16 {
	return 1505
}

func (*OrderCancelReject535) PartyDetailsListReqIDSinceVersion() uint16 {
	return 0
}

func (o *OrderCancelReject535) PartyDetailsListReqIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.PartyDetailsListReqIDSinceVersion()
}

func (*OrderCancelReject535) PartyDetailsListReqIDDeprecated() uint16 {
	return 0
}

func (*OrderCancelReject535) PartyDetailsListReqIDMetaAttribute(meta int) string {
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

func (*OrderCancelReject535) PartyDetailsListReqIDMinValue() uint64 {
	return 0
}

func (*OrderCancelReject535) PartyDetailsListReqIDMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*OrderCancelReject535) PartyDetailsListReqIDNullValue() uint64 {
	return math.MaxUint64
}

func (*OrderCancelReject535) OrdStatusId() uint16 {
	return 39
}

func (*OrderCancelReject535) OrdStatusSinceVersion() uint16 {
	return 0
}

func (o *OrderCancelReject535) OrdStatusInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.OrdStatusSinceVersion()
}

func (*OrderCancelReject535) OrdStatusDeprecated() uint16 {
	return 0
}

func (*OrderCancelReject535) OrdStatusMetaAttribute(meta int) string {
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

func (*OrderCancelReject535) OrdStatusMinValue() byte {
	return byte(32)
}

func (*OrderCancelReject535) OrdStatusMaxValue() byte {
	return byte(126)
}

func (*OrderCancelReject535) OrdStatusNullValue() byte {
	return 0
}

func (*OrderCancelReject535) CxlRejResponseToId() uint16 {
	return 434
}

func (*OrderCancelReject535) CxlRejResponseToSinceVersion() uint16 {
	return 0
}

func (o *OrderCancelReject535) CxlRejResponseToInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.CxlRejResponseToSinceVersion()
}

func (*OrderCancelReject535) CxlRejResponseToDeprecated() uint16 {
	return 0
}

func (*OrderCancelReject535) CxlRejResponseToMetaAttribute(meta int) string {
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

func (*OrderCancelReject535) CxlRejResponseToMinValue() byte {
	return byte(32)
}

func (*OrderCancelReject535) CxlRejResponseToMaxValue() byte {
	return byte(126)
}

func (*OrderCancelReject535) CxlRejResponseToNullValue() byte {
	return 0
}

func (*OrderCancelReject535) OrderIDId() uint16 {
	return 37
}

func (*OrderCancelReject535) OrderIDSinceVersion() uint16 {
	return 0
}

func (o *OrderCancelReject535) OrderIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.OrderIDSinceVersion()
}

func (*OrderCancelReject535) OrderIDDeprecated() uint16 {
	return 0
}

func (*OrderCancelReject535) OrderIDMetaAttribute(meta int) string {
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

func (*OrderCancelReject535) OrderIDMinValue() uint64 {
	return 0
}

func (*OrderCancelReject535) OrderIDMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*OrderCancelReject535) OrderIDNullValue() uint64 {
	return math.MaxUint64
}

func (*OrderCancelReject535) TransactTimeId() uint16 {
	return 60
}

func (*OrderCancelReject535) TransactTimeSinceVersion() uint16 {
	return 0
}

func (o *OrderCancelReject535) TransactTimeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.TransactTimeSinceVersion()
}

func (*OrderCancelReject535) TransactTimeDeprecated() uint16 {
	return 0
}

func (*OrderCancelReject535) TransactTimeMetaAttribute(meta int) string {
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

func (*OrderCancelReject535) TransactTimeMinValue() uint64 {
	return 0
}

func (*OrderCancelReject535) TransactTimeMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*OrderCancelReject535) TransactTimeNullValue() uint64 {
	return math.MaxUint64
}

func (*OrderCancelReject535) SendingTimeEpochId() uint16 {
	return 5297
}

func (*OrderCancelReject535) SendingTimeEpochSinceVersion() uint16 {
	return 0
}

func (o *OrderCancelReject535) SendingTimeEpochInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.SendingTimeEpochSinceVersion()
}

func (*OrderCancelReject535) SendingTimeEpochDeprecated() uint16 {
	return 0
}

func (*OrderCancelReject535) SendingTimeEpochMetaAttribute(meta int) string {
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

func (*OrderCancelReject535) SendingTimeEpochMinValue() uint64 {
	return 0
}

func (*OrderCancelReject535) SendingTimeEpochMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*OrderCancelReject535) SendingTimeEpochNullValue() uint64 {
	return math.MaxUint64
}

func (*OrderCancelReject535) OrderRequestIDId() uint16 {
	return 2422
}

func (*OrderCancelReject535) OrderRequestIDSinceVersion() uint16 {
	return 0
}

func (o *OrderCancelReject535) OrderRequestIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.OrderRequestIDSinceVersion()
}

func (*OrderCancelReject535) OrderRequestIDDeprecated() uint16 {
	return 0
}

func (*OrderCancelReject535) OrderRequestIDMetaAttribute(meta int) string {
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

func (*OrderCancelReject535) OrderRequestIDMinValue() uint64 {
	return 0
}

func (*OrderCancelReject535) OrderRequestIDMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*OrderCancelReject535) OrderRequestIDNullValue() uint64 {
	return math.MaxUint64
}

func (*OrderCancelReject535) LocationId() uint16 {
	return 9537
}

func (*OrderCancelReject535) LocationSinceVersion() uint16 {
	return 0
}

func (o *OrderCancelReject535) LocationInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.LocationSinceVersion()
}

func (*OrderCancelReject535) LocationDeprecated() uint16 {
	return 0
}

func (*OrderCancelReject535) LocationMetaAttribute(meta int) string {
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

func (*OrderCancelReject535) LocationMinValue() byte {
	return byte(32)
}

func (*OrderCancelReject535) LocationMaxValue() byte {
	return byte(126)
}

func (*OrderCancelReject535) LocationNullValue() byte {
	return 0
}

func (o *OrderCancelReject535) LocationCharacterEncoding() string {
	return "US-ASCII"
}

func (*OrderCancelReject535) CxlRejReasonId() uint16 {
	return 102
}

func (*OrderCancelReject535) CxlRejReasonSinceVersion() uint16 {
	return 0
}

func (o *OrderCancelReject535) CxlRejReasonInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.CxlRejReasonSinceVersion()
}

func (*OrderCancelReject535) CxlRejReasonDeprecated() uint16 {
	return 0
}

func (*OrderCancelReject535) CxlRejReasonMetaAttribute(meta int) string {
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

func (*OrderCancelReject535) CxlRejReasonMinValue() uint16 {
	return 0
}

func (*OrderCancelReject535) CxlRejReasonMaxValue() uint16 {
	return math.MaxUint16 - 1
}

func (*OrderCancelReject535) CxlRejReasonNullValue() uint16 {
	return math.MaxUint16
}

func (*OrderCancelReject535) DelayDurationId() uint16 {
	return 5904
}

func (*OrderCancelReject535) DelayDurationSinceVersion() uint16 {
	return 0
}

func (o *OrderCancelReject535) DelayDurationInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.DelayDurationSinceVersion()
}

func (*OrderCancelReject535) DelayDurationDeprecated() uint16 {
	return 0
}

func (*OrderCancelReject535) DelayDurationMetaAttribute(meta int) string {
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

func (*OrderCancelReject535) DelayDurationMinValue() uint16 {
	return 0
}

func (*OrderCancelReject535) DelayDurationMaxValue() uint16 {
	return math.MaxUint16 - 1
}

func (*OrderCancelReject535) DelayDurationNullValue() uint16 {
	return 65535
}

func (*OrderCancelReject535) ManualOrderIndicatorId() uint16 {
	return 1028
}

func (*OrderCancelReject535) ManualOrderIndicatorSinceVersion() uint16 {
	return 0
}

func (o *OrderCancelReject535) ManualOrderIndicatorInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.ManualOrderIndicatorSinceVersion()
}

func (*OrderCancelReject535) ManualOrderIndicatorDeprecated() uint16 {
	return 0
}

func (*OrderCancelReject535) ManualOrderIndicatorMetaAttribute(meta int) string {
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

func (*OrderCancelReject535) PossRetransFlagId() uint16 {
	return 9765
}

func (*OrderCancelReject535) PossRetransFlagSinceVersion() uint16 {
	return 0
}

func (o *OrderCancelReject535) PossRetransFlagInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.PossRetransFlagSinceVersion()
}

func (*OrderCancelReject535) PossRetransFlagDeprecated() uint16 {
	return 0
}

func (*OrderCancelReject535) PossRetransFlagMetaAttribute(meta int) string {
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

func (*OrderCancelReject535) SplitMsgId() uint16 {
	return 9553
}

func (*OrderCancelReject535) SplitMsgSinceVersion() uint16 {
	return 0
}

func (o *OrderCancelReject535) SplitMsgInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.SplitMsgSinceVersion()
}

func (*OrderCancelReject535) SplitMsgDeprecated() uint16 {
	return 0
}

func (*OrderCancelReject535) SplitMsgMetaAttribute(meta int) string {
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

func (*OrderCancelReject535) LiquidityFlagId() uint16 {
	return 9373
}

func (*OrderCancelReject535) LiquidityFlagSinceVersion() uint16 {
	return 0
}

func (o *OrderCancelReject535) LiquidityFlagInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.LiquidityFlagSinceVersion()
}

func (*OrderCancelReject535) LiquidityFlagDeprecated() uint16 {
	return 0
}

func (*OrderCancelReject535) LiquidityFlagMetaAttribute(meta int) string {
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

func (*OrderCancelReject535) DelayToTimeId() uint16 {
	return 7552
}

func (*OrderCancelReject535) DelayToTimeSinceVersion() uint16 {
	return 4
}

func (o *OrderCancelReject535) DelayToTimeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.DelayToTimeSinceVersion()
}

func (*OrderCancelReject535) DelayToTimeDeprecated() uint16 {
	return 0
}

func (*OrderCancelReject535) DelayToTimeMetaAttribute(meta int) string {
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

func (*OrderCancelReject535) DelayToTimeMinValue() uint64 {
	return 0
}

func (*OrderCancelReject535) DelayToTimeMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*OrderCancelReject535) DelayToTimeNullValue() uint64 {
	return 18446744073709551615
}
