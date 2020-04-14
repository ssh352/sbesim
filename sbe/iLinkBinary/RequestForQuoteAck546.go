// Generated SBE (Simple Binary Encoding) message codec

package ilinkbinary

import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
)

type RequestForQuoteAck546 struct {
	SeqNum                uint32
	UUID                  uint64
	Text                  [256]byte
	SenderID              [20]byte
	ExchangeQuoteReqID    [17]byte
	PartyDetailsListReqID uint64
	RequestTime           uint64
	SendingTimeEpoch      uint64
	QuoteReqID            uint64
	Location              [5]byte
	QuoteRejectReason     uint16
	DelayDuration         uint16
	QuoteStatus           QuoteAckStatusEnum
	ManualOrderIndicator  ManualOrdIndReqEnum
	SplitMsg              SplitMsgEnum
	PossRetransFlag       BooleanFlagEnum
	DelayToTime           uint64
}

func (r *RequestForQuoteAck546) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := r.RangeCheck(r.SbeSchemaVersion(), r.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	if err := _m.WriteUint32(_w, r.SeqNum); err != nil {
		return err
	}
	if err := _m.WriteUint64(_w, r.UUID); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, r.Text[:]); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, r.SenderID[:]); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, r.ExchangeQuoteReqID[:]); err != nil {
		return err
	}
	if err := _m.WriteUint64(_w, r.PartyDetailsListReqID); err != nil {
		return err
	}
	if err := _m.WriteUint64(_w, r.RequestTime); err != nil {
		return err
	}
	if err := _m.WriteUint64(_w, r.SendingTimeEpoch); err != nil {
		return err
	}
	if err := _m.WriteUint64(_w, r.QuoteReqID); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, r.Location[:]); err != nil {
		return err
	}
	if err := _m.WriteUint16(_w, r.QuoteRejectReason); err != nil {
		return err
	}
	if err := _m.WriteUint16(_w, r.DelayDuration); err != nil {
		return err
	}
	if err := r.QuoteStatus.Encode(_m, _w); err != nil {
		return err
	}
	if err := r.ManualOrderIndicator.Encode(_m, _w); err != nil {
		return err
	}
	if err := r.SplitMsg.Encode(_m, _w); err != nil {
		return err
	}
	if err := r.PossRetransFlag.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteUint64(_w, r.DelayToTime); err != nil {
		return err
	}
	return nil
}

func (r *RequestForQuoteAck546) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	if !r.SeqNumInActingVersion(actingVersion) {
		r.SeqNum = r.SeqNumNullValue()
	} else {
		if err := _m.ReadUint32(_r, &r.SeqNum); err != nil {
			return err
		}
	}
	if !r.UUIDInActingVersion(actingVersion) {
		r.UUID = r.UUIDNullValue()
	} else {
		if err := _m.ReadUint64(_r, &r.UUID); err != nil {
			return err
		}
	}
	if !r.TextInActingVersion(actingVersion) {
		for idx := 0; idx < 256; idx++ {
			r.Text[idx] = r.TextNullValue()
		}
	} else {
		if err := _m.ReadBytes(_r, r.Text[:]); err != nil {
			return err
		}
	}
	if !r.SenderIDInActingVersion(actingVersion) {
		for idx := 0; idx < 20; idx++ {
			r.SenderID[idx] = r.SenderIDNullValue()
		}
	} else {
		if err := _m.ReadBytes(_r, r.SenderID[:]); err != nil {
			return err
		}
	}
	if !r.ExchangeQuoteReqIDInActingVersion(actingVersion) {
		for idx := 0; idx < 17; idx++ {
			r.ExchangeQuoteReqID[idx] = r.ExchangeQuoteReqIDNullValue()
		}
	} else {
		if err := _m.ReadBytes(_r, r.ExchangeQuoteReqID[:]); err != nil {
			return err
		}
	}
	if !r.PartyDetailsListReqIDInActingVersion(actingVersion) {
		r.PartyDetailsListReqID = r.PartyDetailsListReqIDNullValue()
	} else {
		if err := _m.ReadUint64(_r, &r.PartyDetailsListReqID); err != nil {
			return err
		}
	}
	if !r.RequestTimeInActingVersion(actingVersion) {
		r.RequestTime = r.RequestTimeNullValue()
	} else {
		if err := _m.ReadUint64(_r, &r.RequestTime); err != nil {
			return err
		}
	}
	if !r.SendingTimeEpochInActingVersion(actingVersion) {
		r.SendingTimeEpoch = r.SendingTimeEpochNullValue()
	} else {
		if err := _m.ReadUint64(_r, &r.SendingTimeEpoch); err != nil {
			return err
		}
	}
	if !r.QuoteReqIDInActingVersion(actingVersion) {
		r.QuoteReqID = r.QuoteReqIDNullValue()
	} else {
		if err := _m.ReadUint64(_r, &r.QuoteReqID); err != nil {
			return err
		}
	}
	if !r.LocationInActingVersion(actingVersion) {
		for idx := 0; idx < 5; idx++ {
			r.Location[idx] = r.LocationNullValue()
		}
	} else {
		if err := _m.ReadBytes(_r, r.Location[:]); err != nil {
			return err
		}
	}
	if !r.QuoteRejectReasonInActingVersion(actingVersion) {
		r.QuoteRejectReason = r.QuoteRejectReasonNullValue()
	} else {
		if err := _m.ReadUint16(_r, &r.QuoteRejectReason); err != nil {
			return err
		}
	}
	if !r.DelayDurationInActingVersion(actingVersion) {
		r.DelayDuration = r.DelayDurationNullValue()
	} else {
		if err := _m.ReadUint16(_r, &r.DelayDuration); err != nil {
			return err
		}
	}
	if r.QuoteStatusInActingVersion(actingVersion) {
		if err := r.QuoteStatus.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if r.ManualOrderIndicatorInActingVersion(actingVersion) {
		if err := r.ManualOrderIndicator.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if r.SplitMsgInActingVersion(actingVersion) {
		if err := r.SplitMsg.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if r.PossRetransFlagInActingVersion(actingVersion) {
		if err := r.PossRetransFlag.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if !r.DelayToTimeInActingVersion(actingVersion) {
		r.DelayToTime = r.DelayToTimeNullValue()
	} else {
		if err := _m.ReadUint64(_r, &r.DelayToTime); err != nil {
			return err
		}
	}
	if actingVersion > r.SbeSchemaVersion() && blockLength > r.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-r.SbeBlockLength()))
	}
	if doRangeCheck {
		if err := r.RangeCheck(actingVersion, r.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	return nil
}

func (r *RequestForQuoteAck546) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if r.SeqNumInActingVersion(actingVersion) {
		if r.SeqNum < r.SeqNumMinValue() || r.SeqNum > r.SeqNumMaxValue() {
			return fmt.Errorf("Range check failed on r.SeqNum (%v < %v > %v)", r.SeqNumMinValue(), r.SeqNum, r.SeqNumMaxValue())
		}
	}
	if r.UUIDInActingVersion(actingVersion) {
		if r.UUID < r.UUIDMinValue() || r.UUID > r.UUIDMaxValue() {
			return fmt.Errorf("Range check failed on r.UUID (%v < %v > %v)", r.UUIDMinValue(), r.UUID, r.UUIDMaxValue())
		}
	}
	if r.TextInActingVersion(actingVersion) {
		for idx := 0; idx < 256; idx++ {
			if r.Text[idx] < r.TextMinValue() || r.Text[idx] > r.TextMaxValue() {
				return fmt.Errorf("Range check failed on r.Text[%d] (%v < %v > %v)", idx, r.TextMinValue(), r.Text[idx], r.TextMaxValue())
			}
		}
	}
	if r.SenderIDInActingVersion(actingVersion) {
		for idx := 0; idx < 20; idx++ {
			if r.SenderID[idx] < r.SenderIDMinValue() || r.SenderID[idx] > r.SenderIDMaxValue() {
				return fmt.Errorf("Range check failed on r.SenderID[%d] (%v < %v > %v)", idx, r.SenderIDMinValue(), r.SenderID[idx], r.SenderIDMaxValue())
			}
		}
	}
	if r.ExchangeQuoteReqIDInActingVersion(actingVersion) {
		for idx := 0; idx < 17; idx++ {
			if r.ExchangeQuoteReqID[idx] < r.ExchangeQuoteReqIDMinValue() || r.ExchangeQuoteReqID[idx] > r.ExchangeQuoteReqIDMaxValue() {
				return fmt.Errorf("Range check failed on r.ExchangeQuoteReqID[%d] (%v < %v > %v)", idx, r.ExchangeQuoteReqIDMinValue(), r.ExchangeQuoteReqID[idx], r.ExchangeQuoteReqIDMaxValue())
			}
		}
	}
	if r.PartyDetailsListReqIDInActingVersion(actingVersion) {
		if r.PartyDetailsListReqID < r.PartyDetailsListReqIDMinValue() || r.PartyDetailsListReqID > r.PartyDetailsListReqIDMaxValue() {
			return fmt.Errorf("Range check failed on r.PartyDetailsListReqID (%v < %v > %v)", r.PartyDetailsListReqIDMinValue(), r.PartyDetailsListReqID, r.PartyDetailsListReqIDMaxValue())
		}
	}
	if r.RequestTimeInActingVersion(actingVersion) {
		if r.RequestTime < r.RequestTimeMinValue() || r.RequestTime > r.RequestTimeMaxValue() {
			return fmt.Errorf("Range check failed on r.RequestTime (%v < %v > %v)", r.RequestTimeMinValue(), r.RequestTime, r.RequestTimeMaxValue())
		}
	}
	if r.SendingTimeEpochInActingVersion(actingVersion) {
		if r.SendingTimeEpoch < r.SendingTimeEpochMinValue() || r.SendingTimeEpoch > r.SendingTimeEpochMaxValue() {
			return fmt.Errorf("Range check failed on r.SendingTimeEpoch (%v < %v > %v)", r.SendingTimeEpochMinValue(), r.SendingTimeEpoch, r.SendingTimeEpochMaxValue())
		}
	}
	if r.QuoteReqIDInActingVersion(actingVersion) {
		if r.QuoteReqID < r.QuoteReqIDMinValue() || r.QuoteReqID > r.QuoteReqIDMaxValue() {
			return fmt.Errorf("Range check failed on r.QuoteReqID (%v < %v > %v)", r.QuoteReqIDMinValue(), r.QuoteReqID, r.QuoteReqIDMaxValue())
		}
	}
	if r.LocationInActingVersion(actingVersion) {
		for idx := 0; idx < 5; idx++ {
			if r.Location[idx] < r.LocationMinValue() || r.Location[idx] > r.LocationMaxValue() {
				return fmt.Errorf("Range check failed on r.Location[%d] (%v < %v > %v)", idx, r.LocationMinValue(), r.Location[idx], r.LocationMaxValue())
			}
		}
	}
	if r.QuoteRejectReasonInActingVersion(actingVersion) {
		if r.QuoteRejectReason != r.QuoteRejectReasonNullValue() && (r.QuoteRejectReason < r.QuoteRejectReasonMinValue() || r.QuoteRejectReason > r.QuoteRejectReasonMaxValue()) {
			return fmt.Errorf("Range check failed on r.QuoteRejectReason (%v < %v > %v)", r.QuoteRejectReasonMinValue(), r.QuoteRejectReason, r.QuoteRejectReasonMaxValue())
		}
	}
	if r.DelayDurationInActingVersion(actingVersion) {
		if r.DelayDuration != r.DelayDurationNullValue() && (r.DelayDuration < r.DelayDurationMinValue() || r.DelayDuration > r.DelayDurationMaxValue()) {
			return fmt.Errorf("Range check failed on r.DelayDuration (%v < %v > %v)", r.DelayDurationMinValue(), r.DelayDuration, r.DelayDurationMaxValue())
		}
	}
	if err := r.QuoteStatus.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if err := r.ManualOrderIndicator.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if err := r.SplitMsg.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if err := r.PossRetransFlag.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if r.DelayToTimeInActingVersion(actingVersion) {
		if r.DelayToTime != r.DelayToTimeNullValue() && (r.DelayToTime < r.DelayToTimeMinValue() || r.DelayToTime > r.DelayToTimeMaxValue()) {
			return fmt.Errorf("Range check failed on r.DelayToTime (%v < %v > %v)", r.DelayToTimeMinValue(), r.DelayToTime, r.DelayToTimeMaxValue())
		}
	}
	return nil
}

func RequestForQuoteAck546Init(r *RequestForQuoteAck546) {
	for idx := 0; idx < 256; idx++ {
		r.Text[idx] = 0
	}
	for idx := 0; idx < 17; idx++ {
		r.ExchangeQuoteReqID[idx] = 0
	}
	r.QuoteRejectReason = 65535
	r.DelayDuration = 65535
	r.DelayToTime = 18446744073709551615
	return
}

func (*RequestForQuoteAck546) SbeBlockLength() (blockLength uint16) {
	return 358
}

func (*RequestForQuoteAck546) SbeTemplateId() (templateId uint16) {
	return 546
}

func (*RequestForQuoteAck546) SbeSchemaId() (schemaId uint16) {
	return 8
}

func (*RequestForQuoteAck546) SbeSchemaVersion() (schemaVersion uint16) {
	return 5
}

func (*RequestForQuoteAck546) SbeSemanticType() (semanticType []byte) {
	return []byte("b")
}

func (*RequestForQuoteAck546) SeqNumId() uint16 {
	return 9726
}

func (*RequestForQuoteAck546) SeqNumSinceVersion() uint16 {
	return 0
}

func (r *RequestForQuoteAck546) SeqNumInActingVersion(actingVersion uint16) bool {
	return actingVersion >= r.SeqNumSinceVersion()
}

func (*RequestForQuoteAck546) SeqNumDeprecated() uint16 {
	return 0
}

func (*RequestForQuoteAck546) SeqNumMetaAttribute(meta int) string {
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

func (*RequestForQuoteAck546) SeqNumMinValue() uint32 {
	return 0
}

func (*RequestForQuoteAck546) SeqNumMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*RequestForQuoteAck546) SeqNumNullValue() uint32 {
	return math.MaxUint32
}

func (*RequestForQuoteAck546) UUIDId() uint16 {
	return 39001
}

func (*RequestForQuoteAck546) UUIDSinceVersion() uint16 {
	return 0
}

func (r *RequestForQuoteAck546) UUIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= r.UUIDSinceVersion()
}

func (*RequestForQuoteAck546) UUIDDeprecated() uint16 {
	return 0
}

func (*RequestForQuoteAck546) UUIDMetaAttribute(meta int) string {
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

func (*RequestForQuoteAck546) UUIDMinValue() uint64 {
	return 0
}

func (*RequestForQuoteAck546) UUIDMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*RequestForQuoteAck546) UUIDNullValue() uint64 {
	return math.MaxUint64
}

func (*RequestForQuoteAck546) TextId() uint16 {
	return 58
}

func (*RequestForQuoteAck546) TextSinceVersion() uint16 {
	return 0
}

func (r *RequestForQuoteAck546) TextInActingVersion(actingVersion uint16) bool {
	return actingVersion >= r.TextSinceVersion()
}

func (*RequestForQuoteAck546) TextDeprecated() uint16 {
	return 0
}

func (*RequestForQuoteAck546) TextMetaAttribute(meta int) string {
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

func (*RequestForQuoteAck546) TextMinValue() byte {
	return byte(32)
}

func (*RequestForQuoteAck546) TextMaxValue() byte {
	return byte(126)
}

func (*RequestForQuoteAck546) TextNullValue() byte {
	return 0
}

func (r *RequestForQuoteAck546) TextCharacterEncoding() string {
	return "US-ASCII"
}

func (*RequestForQuoteAck546) SenderIDId() uint16 {
	return 5392
}

func (*RequestForQuoteAck546) SenderIDSinceVersion() uint16 {
	return 0
}

func (r *RequestForQuoteAck546) SenderIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= r.SenderIDSinceVersion()
}

func (*RequestForQuoteAck546) SenderIDDeprecated() uint16 {
	return 0
}

func (*RequestForQuoteAck546) SenderIDMetaAttribute(meta int) string {
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

func (*RequestForQuoteAck546) SenderIDMinValue() byte {
	return byte(32)
}

func (*RequestForQuoteAck546) SenderIDMaxValue() byte {
	return byte(126)
}

func (*RequestForQuoteAck546) SenderIDNullValue() byte {
	return 0
}

func (r *RequestForQuoteAck546) SenderIDCharacterEncoding() string {
	return "US-ASCII"
}

func (*RequestForQuoteAck546) ExchangeQuoteReqIDId() uint16 {
	return 9770
}

func (*RequestForQuoteAck546) ExchangeQuoteReqIDSinceVersion() uint16 {
	return 0
}

func (r *RequestForQuoteAck546) ExchangeQuoteReqIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= r.ExchangeQuoteReqIDSinceVersion()
}

func (*RequestForQuoteAck546) ExchangeQuoteReqIDDeprecated() uint16 {
	return 0
}

func (*RequestForQuoteAck546) ExchangeQuoteReqIDMetaAttribute(meta int) string {
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

func (*RequestForQuoteAck546) ExchangeQuoteReqIDMinValue() byte {
	return byte(32)
}

func (*RequestForQuoteAck546) ExchangeQuoteReqIDMaxValue() byte {
	return byte(126)
}

func (*RequestForQuoteAck546) ExchangeQuoteReqIDNullValue() byte {
	return 0
}

func (r *RequestForQuoteAck546) ExchangeQuoteReqIDCharacterEncoding() string {
	return "US-ASCII"
}

func (*RequestForQuoteAck546) PartyDetailsListReqIDId() uint16 {
	return 1505
}

func (*RequestForQuoteAck546) PartyDetailsListReqIDSinceVersion() uint16 {
	return 0
}

func (r *RequestForQuoteAck546) PartyDetailsListReqIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= r.PartyDetailsListReqIDSinceVersion()
}

func (*RequestForQuoteAck546) PartyDetailsListReqIDDeprecated() uint16 {
	return 0
}

func (*RequestForQuoteAck546) PartyDetailsListReqIDMetaAttribute(meta int) string {
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

func (*RequestForQuoteAck546) PartyDetailsListReqIDMinValue() uint64 {
	return 0
}

func (*RequestForQuoteAck546) PartyDetailsListReqIDMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*RequestForQuoteAck546) PartyDetailsListReqIDNullValue() uint64 {
	return math.MaxUint64
}

func (*RequestForQuoteAck546) RequestTimeId() uint16 {
	return 5979
}

func (*RequestForQuoteAck546) RequestTimeSinceVersion() uint16 {
	return 0
}

func (r *RequestForQuoteAck546) RequestTimeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= r.RequestTimeSinceVersion()
}

func (*RequestForQuoteAck546) RequestTimeDeprecated() uint16 {
	return 0
}

func (*RequestForQuoteAck546) RequestTimeMetaAttribute(meta int) string {
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

func (*RequestForQuoteAck546) RequestTimeMinValue() uint64 {
	return 0
}

func (*RequestForQuoteAck546) RequestTimeMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*RequestForQuoteAck546) RequestTimeNullValue() uint64 {
	return math.MaxUint64
}

func (*RequestForQuoteAck546) SendingTimeEpochId() uint16 {
	return 5297
}

func (*RequestForQuoteAck546) SendingTimeEpochSinceVersion() uint16 {
	return 0
}

func (r *RequestForQuoteAck546) SendingTimeEpochInActingVersion(actingVersion uint16) bool {
	return actingVersion >= r.SendingTimeEpochSinceVersion()
}

func (*RequestForQuoteAck546) SendingTimeEpochDeprecated() uint16 {
	return 0
}

func (*RequestForQuoteAck546) SendingTimeEpochMetaAttribute(meta int) string {
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

func (*RequestForQuoteAck546) SendingTimeEpochMinValue() uint64 {
	return 0
}

func (*RequestForQuoteAck546) SendingTimeEpochMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*RequestForQuoteAck546) SendingTimeEpochNullValue() uint64 {
	return math.MaxUint64
}

func (*RequestForQuoteAck546) QuoteReqIDId() uint16 {
	return 131
}

func (*RequestForQuoteAck546) QuoteReqIDSinceVersion() uint16 {
	return 0
}

func (r *RequestForQuoteAck546) QuoteReqIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= r.QuoteReqIDSinceVersion()
}

func (*RequestForQuoteAck546) QuoteReqIDDeprecated() uint16 {
	return 0
}

func (*RequestForQuoteAck546) QuoteReqIDMetaAttribute(meta int) string {
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

func (*RequestForQuoteAck546) QuoteReqIDMinValue() uint64 {
	return 0
}

func (*RequestForQuoteAck546) QuoteReqIDMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*RequestForQuoteAck546) QuoteReqIDNullValue() uint64 {
	return math.MaxUint64
}

func (*RequestForQuoteAck546) LocationId() uint16 {
	return 9537
}

func (*RequestForQuoteAck546) LocationSinceVersion() uint16 {
	return 0
}

func (r *RequestForQuoteAck546) LocationInActingVersion(actingVersion uint16) bool {
	return actingVersion >= r.LocationSinceVersion()
}

func (*RequestForQuoteAck546) LocationDeprecated() uint16 {
	return 0
}

func (*RequestForQuoteAck546) LocationMetaAttribute(meta int) string {
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

func (*RequestForQuoteAck546) LocationMinValue() byte {
	return byte(32)
}

func (*RequestForQuoteAck546) LocationMaxValue() byte {
	return byte(126)
}

func (*RequestForQuoteAck546) LocationNullValue() byte {
	return 0
}

func (r *RequestForQuoteAck546) LocationCharacterEncoding() string {
	return "US-ASCII"
}

func (*RequestForQuoteAck546) QuoteRejectReasonId() uint16 {
	return 300
}

func (*RequestForQuoteAck546) QuoteRejectReasonSinceVersion() uint16 {
	return 0
}

func (r *RequestForQuoteAck546) QuoteRejectReasonInActingVersion(actingVersion uint16) bool {
	return actingVersion >= r.QuoteRejectReasonSinceVersion()
}

func (*RequestForQuoteAck546) QuoteRejectReasonDeprecated() uint16 {
	return 0
}

func (*RequestForQuoteAck546) QuoteRejectReasonMetaAttribute(meta int) string {
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

func (*RequestForQuoteAck546) QuoteRejectReasonMinValue() uint16 {
	return 0
}

func (*RequestForQuoteAck546) QuoteRejectReasonMaxValue() uint16 {
	return math.MaxUint16 - 1
}

func (*RequestForQuoteAck546) QuoteRejectReasonNullValue() uint16 {
	return 65535
}

func (*RequestForQuoteAck546) DelayDurationId() uint16 {
	return 5904
}

func (*RequestForQuoteAck546) DelayDurationSinceVersion() uint16 {
	return 0
}

func (r *RequestForQuoteAck546) DelayDurationInActingVersion(actingVersion uint16) bool {
	return actingVersion >= r.DelayDurationSinceVersion()
}

func (*RequestForQuoteAck546) DelayDurationDeprecated() uint16 {
	return 0
}

func (*RequestForQuoteAck546) DelayDurationMetaAttribute(meta int) string {
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

func (*RequestForQuoteAck546) DelayDurationMinValue() uint16 {
	return 0
}

func (*RequestForQuoteAck546) DelayDurationMaxValue() uint16 {
	return math.MaxUint16 - 1
}

func (*RequestForQuoteAck546) DelayDurationNullValue() uint16 {
	return 65535
}

func (*RequestForQuoteAck546) QuoteStatusId() uint16 {
	return 297
}

func (*RequestForQuoteAck546) QuoteStatusSinceVersion() uint16 {
	return 0
}

func (r *RequestForQuoteAck546) QuoteStatusInActingVersion(actingVersion uint16) bool {
	return actingVersion >= r.QuoteStatusSinceVersion()
}

func (*RequestForQuoteAck546) QuoteStatusDeprecated() uint16 {
	return 0
}

func (*RequestForQuoteAck546) QuoteStatusMetaAttribute(meta int) string {
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

func (*RequestForQuoteAck546) ManualOrderIndicatorId() uint16 {
	return 1028
}

func (*RequestForQuoteAck546) ManualOrderIndicatorSinceVersion() uint16 {
	return 0
}

func (r *RequestForQuoteAck546) ManualOrderIndicatorInActingVersion(actingVersion uint16) bool {
	return actingVersion >= r.ManualOrderIndicatorSinceVersion()
}

func (*RequestForQuoteAck546) ManualOrderIndicatorDeprecated() uint16 {
	return 0
}

func (*RequestForQuoteAck546) ManualOrderIndicatorMetaAttribute(meta int) string {
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

func (*RequestForQuoteAck546) SplitMsgId() uint16 {
	return 9553
}

func (*RequestForQuoteAck546) SplitMsgSinceVersion() uint16 {
	return 0
}

func (r *RequestForQuoteAck546) SplitMsgInActingVersion(actingVersion uint16) bool {
	return actingVersion >= r.SplitMsgSinceVersion()
}

func (*RequestForQuoteAck546) SplitMsgDeprecated() uint16 {
	return 0
}

func (*RequestForQuoteAck546) SplitMsgMetaAttribute(meta int) string {
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

func (*RequestForQuoteAck546) PossRetransFlagId() uint16 {
	return 9765
}

func (*RequestForQuoteAck546) PossRetransFlagSinceVersion() uint16 {
	return 2
}

func (r *RequestForQuoteAck546) PossRetransFlagInActingVersion(actingVersion uint16) bool {
	return actingVersion >= r.PossRetransFlagSinceVersion()
}

func (*RequestForQuoteAck546) PossRetransFlagDeprecated() uint16 {
	return 0
}

func (*RequestForQuoteAck546) PossRetransFlagMetaAttribute(meta int) string {
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

func (*RequestForQuoteAck546) DelayToTimeId() uint16 {
	return 7552
}

func (*RequestForQuoteAck546) DelayToTimeSinceVersion() uint16 {
	return 4
}

func (r *RequestForQuoteAck546) DelayToTimeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= r.DelayToTimeSinceVersion()
}

func (*RequestForQuoteAck546) DelayToTimeDeprecated() uint16 {
	return 0
}

func (*RequestForQuoteAck546) DelayToTimeMetaAttribute(meta int) string {
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

func (*RequestForQuoteAck546) DelayToTimeMinValue() uint64 {
	return 0
}

func (*RequestForQuoteAck546) DelayToTimeMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*RequestForQuoteAck546) DelayToTimeNullValue() uint64 {
	return 18446744073709551615
}
