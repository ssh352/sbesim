// Generated SBE (Simple Binary Encoding) message codec

package ilinkbinary

import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
)

type RequestForQuote543 struct {
	PartyDetailsListReqID uint64
	QuoteReqID            uint64
	ManualOrderIndicator  ManualOrdIndReqEnum
	SeqNum                uint32
	SenderID              [20]byte
	SendingTimeEpoch      uint64
	Location              [5]byte
	QuoteType             QuoteTypEnum
	NoRelatedSym          []RequestForQuote543NoRelatedSym
}
type RequestForQuote543NoRelatedSym struct {
	SecurityID int32
	OrderQty   uint32
	Side       RFQSideEnum
}

func (r *RequestForQuote543) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := r.RangeCheck(r.SbeSchemaVersion(), r.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	if err := _m.WriteUint64(_w, r.PartyDetailsListReqID); err != nil {
		return err
	}
	if err := _m.WriteUint64(_w, r.QuoteReqID); err != nil {
		return err
	}
	if err := r.ManualOrderIndicator.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteUint32(_w, r.SeqNum); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, r.SenderID[:]); err != nil {
		return err
	}
	if err := _m.WriteUint64(_w, r.SendingTimeEpoch); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, r.Location[:]); err != nil {
		return err
	}
	if err := r.QuoteType.Encode(_m, _w); err != nil {
		return err
	}
	var NoRelatedSymBlockLength uint16 = 9
	if err := _m.WriteUint16(_w, NoRelatedSymBlockLength); err != nil {
		return err
	}
	var NoRelatedSymNumInGroup uint8 = uint8(len(r.NoRelatedSym))
	if err := _m.WriteUint8(_w, NoRelatedSymNumInGroup); err != nil {
		return err
	}
	for _, prop := range r.NoRelatedSym {
		if err := prop.Encode(_m, _w); err != nil {
			return err
		}
	}
	return nil
}

func (r *RequestForQuote543) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	if !r.PartyDetailsListReqIDInActingVersion(actingVersion) {
		r.PartyDetailsListReqID = r.PartyDetailsListReqIDNullValue()
	} else {
		if err := _m.ReadUint64(_r, &r.PartyDetailsListReqID); err != nil {
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
	if r.ManualOrderIndicatorInActingVersion(actingVersion) {
		if err := r.ManualOrderIndicator.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if !r.SeqNumInActingVersion(actingVersion) {
		r.SeqNum = r.SeqNumNullValue()
	} else {
		if err := _m.ReadUint32(_r, &r.SeqNum); err != nil {
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
	if !r.SendingTimeEpochInActingVersion(actingVersion) {
		r.SendingTimeEpoch = r.SendingTimeEpochNullValue()
	} else {
		if err := _m.ReadUint64(_r, &r.SendingTimeEpoch); err != nil {
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
	if r.QuoteTypeInActingVersion(actingVersion) {
		if err := r.QuoteType.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if actingVersion > r.SbeSchemaVersion() && blockLength > r.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-r.SbeBlockLength()))
	}

	if r.NoRelatedSymInActingVersion(actingVersion) {
		var NoRelatedSymBlockLength uint16
		if err := _m.ReadUint16(_r, &NoRelatedSymBlockLength); err != nil {
			return err
		}
		var NoRelatedSymNumInGroup uint8
		if err := _m.ReadUint8(_r, &NoRelatedSymNumInGroup); err != nil {
			return err
		}
		if cap(r.NoRelatedSym) < int(NoRelatedSymNumInGroup) {
			r.NoRelatedSym = make([]RequestForQuote543NoRelatedSym, NoRelatedSymNumInGroup)
		}
		for i, _ := range r.NoRelatedSym {
			if err := r.NoRelatedSym[i].Decode(_m, _r, actingVersion, uint(NoRelatedSymBlockLength)); err != nil {
				return err
			}
		}
	}
	if doRangeCheck {
		if err := r.RangeCheck(actingVersion, r.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	return nil
}

func (r *RequestForQuote543) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if r.PartyDetailsListReqIDInActingVersion(actingVersion) {
		if r.PartyDetailsListReqID < r.PartyDetailsListReqIDMinValue() || r.PartyDetailsListReqID > r.PartyDetailsListReqIDMaxValue() {
			return fmt.Errorf("Range check failed on r.PartyDetailsListReqID (%v < %v > %v)", r.PartyDetailsListReqIDMinValue(), r.PartyDetailsListReqID, r.PartyDetailsListReqIDMaxValue())
		}
	}
	if r.QuoteReqIDInActingVersion(actingVersion) {
		if r.QuoteReqID < r.QuoteReqIDMinValue() || r.QuoteReqID > r.QuoteReqIDMaxValue() {
			return fmt.Errorf("Range check failed on r.QuoteReqID (%v < %v > %v)", r.QuoteReqIDMinValue(), r.QuoteReqID, r.QuoteReqIDMaxValue())
		}
	}
	if err := r.ManualOrderIndicator.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if r.SeqNumInActingVersion(actingVersion) {
		if r.SeqNum < r.SeqNumMinValue() || r.SeqNum > r.SeqNumMaxValue() {
			return fmt.Errorf("Range check failed on r.SeqNum (%v < %v > %v)", r.SeqNumMinValue(), r.SeqNum, r.SeqNumMaxValue())
		}
	}
	if r.SenderIDInActingVersion(actingVersion) {
		for idx := 0; idx < 20; idx++ {
			if r.SenderID[idx] < r.SenderIDMinValue() || r.SenderID[idx] > r.SenderIDMaxValue() {
				return fmt.Errorf("Range check failed on r.SenderID[%d] (%v < %v > %v)", idx, r.SenderIDMinValue(), r.SenderID[idx], r.SenderIDMaxValue())
			}
		}
	}
	if r.SendingTimeEpochInActingVersion(actingVersion) {
		if r.SendingTimeEpoch < r.SendingTimeEpochMinValue() || r.SendingTimeEpoch > r.SendingTimeEpochMaxValue() {
			return fmt.Errorf("Range check failed on r.SendingTimeEpoch (%v < %v > %v)", r.SendingTimeEpochMinValue(), r.SendingTimeEpoch, r.SendingTimeEpochMaxValue())
		}
	}
	if r.LocationInActingVersion(actingVersion) {
		for idx := 0; idx < 5; idx++ {
			if r.Location[idx] < r.LocationMinValue() || r.Location[idx] > r.LocationMaxValue() {
				return fmt.Errorf("Range check failed on r.Location[%d] (%v < %v > %v)", idx, r.LocationMinValue(), r.Location[idx], r.LocationMaxValue())
			}
		}
	}
	if err := r.QuoteType.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	for _, prop := range r.NoRelatedSym {
		if err := prop.RangeCheck(actingVersion, schemaVersion); err != nil {
			return err
		}
	}
	return nil
}

func RequestForQuote543Init(r *RequestForQuote543) {
	return
}

func (r *RequestForQuote543NoRelatedSym) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteInt32(_w, r.SecurityID); err != nil {
		return err
	}
	if err := _m.WriteUint32(_w, r.OrderQty); err != nil {
		return err
	}
	if err := r.Side.Encode(_m, _w); err != nil {
		return err
	}
	return nil
}

func (r *RequestForQuote543NoRelatedSym) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint) error {
	if !r.SecurityIDInActingVersion(actingVersion) {
		r.SecurityID = r.SecurityIDNullValue()
	} else {
		if err := _m.ReadInt32(_r, &r.SecurityID); err != nil {
			return err
		}
	}
	if !r.OrderQtyInActingVersion(actingVersion) {
		r.OrderQty = r.OrderQtyNullValue()
	} else {
		if err := _m.ReadUint32(_r, &r.OrderQty); err != nil {
			return err
		}
	}
	if r.SideInActingVersion(actingVersion) {
		if err := r.Side.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if actingVersion > r.SbeSchemaVersion() && blockLength > r.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-r.SbeBlockLength()))
	}
	return nil
}

func (r *RequestForQuote543NoRelatedSym) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if r.SecurityIDInActingVersion(actingVersion) {
		if r.SecurityID < r.SecurityIDMinValue() || r.SecurityID > r.SecurityIDMaxValue() {
			return fmt.Errorf("Range check failed on r.SecurityID (%v < %v > %v)", r.SecurityIDMinValue(), r.SecurityID, r.SecurityIDMaxValue())
		}
	}
	if r.OrderQtyInActingVersion(actingVersion) {
		if r.OrderQty != r.OrderQtyNullValue() && (r.OrderQty < r.OrderQtyMinValue() || r.OrderQty > r.OrderQtyMaxValue()) {
			return fmt.Errorf("Range check failed on r.OrderQty (%v < %v > %v)", r.OrderQtyMinValue(), r.OrderQty, r.OrderQtyMaxValue())
		}
	}
	if err := r.Side.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	return nil
}

func RequestForQuote543NoRelatedSymInit(r *RequestForQuote543NoRelatedSym) {
	r.OrderQty = 4294967295
	return
}

func (*RequestForQuote543) SbeBlockLength() (blockLength uint16) {
	return 55
}

func (*RequestForQuote543) SbeTemplateId() (templateId uint16) {
	return 543
}

func (*RequestForQuote543) SbeSchemaId() (schemaId uint16) {
	return 8
}

func (*RequestForQuote543) SbeSchemaVersion() (schemaVersion uint16) {
	return 5
}

func (*RequestForQuote543) SbeSemanticType() (semanticType []byte) {
	return []byte("R")
}

func (*RequestForQuote543) PartyDetailsListReqIDId() uint16 {
	return 1505
}

func (*RequestForQuote543) PartyDetailsListReqIDSinceVersion() uint16 {
	return 0
}

func (r *RequestForQuote543) PartyDetailsListReqIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= r.PartyDetailsListReqIDSinceVersion()
}

func (*RequestForQuote543) PartyDetailsListReqIDDeprecated() uint16 {
	return 0
}

func (*RequestForQuote543) PartyDetailsListReqIDMetaAttribute(meta int) string {
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

func (*RequestForQuote543) PartyDetailsListReqIDMinValue() uint64 {
	return 0
}

func (*RequestForQuote543) PartyDetailsListReqIDMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*RequestForQuote543) PartyDetailsListReqIDNullValue() uint64 {
	return math.MaxUint64
}

func (*RequestForQuote543) QuoteReqIDId() uint16 {
	return 131
}

func (*RequestForQuote543) QuoteReqIDSinceVersion() uint16 {
	return 0
}

func (r *RequestForQuote543) QuoteReqIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= r.QuoteReqIDSinceVersion()
}

func (*RequestForQuote543) QuoteReqIDDeprecated() uint16 {
	return 0
}

func (*RequestForQuote543) QuoteReqIDMetaAttribute(meta int) string {
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

func (*RequestForQuote543) QuoteReqIDMinValue() uint64 {
	return 0
}

func (*RequestForQuote543) QuoteReqIDMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*RequestForQuote543) QuoteReqIDNullValue() uint64 {
	return math.MaxUint64
}

func (*RequestForQuote543) ManualOrderIndicatorId() uint16 {
	return 1028
}

func (*RequestForQuote543) ManualOrderIndicatorSinceVersion() uint16 {
	return 0
}

func (r *RequestForQuote543) ManualOrderIndicatorInActingVersion(actingVersion uint16) bool {
	return actingVersion >= r.ManualOrderIndicatorSinceVersion()
}

func (*RequestForQuote543) ManualOrderIndicatorDeprecated() uint16 {
	return 0
}

func (*RequestForQuote543) ManualOrderIndicatorMetaAttribute(meta int) string {
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

func (*RequestForQuote543) SeqNumId() uint16 {
	return 9726
}

func (*RequestForQuote543) SeqNumSinceVersion() uint16 {
	return 0
}

func (r *RequestForQuote543) SeqNumInActingVersion(actingVersion uint16) bool {
	return actingVersion >= r.SeqNumSinceVersion()
}

func (*RequestForQuote543) SeqNumDeprecated() uint16 {
	return 0
}

func (*RequestForQuote543) SeqNumMetaAttribute(meta int) string {
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

func (*RequestForQuote543) SeqNumMinValue() uint32 {
	return 0
}

func (*RequestForQuote543) SeqNumMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*RequestForQuote543) SeqNumNullValue() uint32 {
	return math.MaxUint32
}

func (*RequestForQuote543) SenderIDId() uint16 {
	return 5392
}

func (*RequestForQuote543) SenderIDSinceVersion() uint16 {
	return 0
}

func (r *RequestForQuote543) SenderIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= r.SenderIDSinceVersion()
}

func (*RequestForQuote543) SenderIDDeprecated() uint16 {
	return 0
}

func (*RequestForQuote543) SenderIDMetaAttribute(meta int) string {
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

func (*RequestForQuote543) SenderIDMinValue() byte {
	return byte(32)
}

func (*RequestForQuote543) SenderIDMaxValue() byte {
	return byte(126)
}

func (*RequestForQuote543) SenderIDNullValue() byte {
	return 0
}

func (r *RequestForQuote543) SenderIDCharacterEncoding() string {
	return "US-ASCII"
}

func (*RequestForQuote543) SendingTimeEpochId() uint16 {
	return 5297
}

func (*RequestForQuote543) SendingTimeEpochSinceVersion() uint16 {
	return 0
}

func (r *RequestForQuote543) SendingTimeEpochInActingVersion(actingVersion uint16) bool {
	return actingVersion >= r.SendingTimeEpochSinceVersion()
}

func (*RequestForQuote543) SendingTimeEpochDeprecated() uint16 {
	return 0
}

func (*RequestForQuote543) SendingTimeEpochMetaAttribute(meta int) string {
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

func (*RequestForQuote543) SendingTimeEpochMinValue() uint64 {
	return 0
}

func (*RequestForQuote543) SendingTimeEpochMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*RequestForQuote543) SendingTimeEpochNullValue() uint64 {
	return math.MaxUint64
}

func (*RequestForQuote543) LocationId() uint16 {
	return 9537
}

func (*RequestForQuote543) LocationSinceVersion() uint16 {
	return 0
}

func (r *RequestForQuote543) LocationInActingVersion(actingVersion uint16) bool {
	return actingVersion >= r.LocationSinceVersion()
}

func (*RequestForQuote543) LocationDeprecated() uint16 {
	return 0
}

func (*RequestForQuote543) LocationMetaAttribute(meta int) string {
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

func (*RequestForQuote543) LocationMinValue() byte {
	return byte(32)
}

func (*RequestForQuote543) LocationMaxValue() byte {
	return byte(126)
}

func (*RequestForQuote543) LocationNullValue() byte {
	return 0
}

func (r *RequestForQuote543) LocationCharacterEncoding() string {
	return "US-ASCII"
}

func (*RequestForQuote543) QuoteTypeId() uint16 {
	return 537
}

func (*RequestForQuote543) QuoteTypeSinceVersion() uint16 {
	return 0
}

func (r *RequestForQuote543) QuoteTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= r.QuoteTypeSinceVersion()
}

func (*RequestForQuote543) QuoteTypeDeprecated() uint16 {
	return 0
}

func (*RequestForQuote543) QuoteTypeMetaAttribute(meta int) string {
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

func (*RequestForQuote543NoRelatedSym) SecurityIDId() uint16 {
	return 48
}

func (*RequestForQuote543NoRelatedSym) SecurityIDSinceVersion() uint16 {
	return 0
}

func (r *RequestForQuote543NoRelatedSym) SecurityIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= r.SecurityIDSinceVersion()
}

func (*RequestForQuote543NoRelatedSym) SecurityIDDeprecated() uint16 {
	return 0
}

func (*RequestForQuote543NoRelatedSym) SecurityIDMetaAttribute(meta int) string {
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

func (*RequestForQuote543NoRelatedSym) SecurityIDMinValue() int32 {
	return math.MinInt32 + 1
}

func (*RequestForQuote543NoRelatedSym) SecurityIDMaxValue() int32 {
	return math.MaxInt32
}

func (*RequestForQuote543NoRelatedSym) SecurityIDNullValue() int32 {
	return math.MinInt32
}

func (*RequestForQuote543NoRelatedSym) OrderQtyId() uint16 {
	return 38
}

func (*RequestForQuote543NoRelatedSym) OrderQtySinceVersion() uint16 {
	return 0
}

func (r *RequestForQuote543NoRelatedSym) OrderQtyInActingVersion(actingVersion uint16) bool {
	return actingVersion >= r.OrderQtySinceVersion()
}

func (*RequestForQuote543NoRelatedSym) OrderQtyDeprecated() uint16 {
	return 0
}

func (*RequestForQuote543NoRelatedSym) OrderQtyMetaAttribute(meta int) string {
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

func (*RequestForQuote543NoRelatedSym) OrderQtyMinValue() uint32 {
	return 0
}

func (*RequestForQuote543NoRelatedSym) OrderQtyMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*RequestForQuote543NoRelatedSym) OrderQtyNullValue() uint32 {
	return 4294967295
}

func (*RequestForQuote543NoRelatedSym) SideId() uint16 {
	return 54
}

func (*RequestForQuote543NoRelatedSym) SideSinceVersion() uint16 {
	return 0
}

func (r *RequestForQuote543NoRelatedSym) SideInActingVersion(actingVersion uint16) bool {
	return actingVersion >= r.SideSinceVersion()
}

func (*RequestForQuote543NoRelatedSym) SideDeprecated() uint16 {
	return 0
}

func (*RequestForQuote543NoRelatedSym) SideMetaAttribute(meta int) string {
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

func (*RequestForQuote543) NoRelatedSymId() uint16 {
	return 146
}

func (*RequestForQuote543) NoRelatedSymSinceVersion() uint16 {
	return 0
}

func (r *RequestForQuote543) NoRelatedSymInActingVersion(actingVersion uint16) bool {
	return actingVersion >= r.NoRelatedSymSinceVersion()
}

func (*RequestForQuote543) NoRelatedSymDeprecated() uint16 {
	return 0
}

func (*RequestForQuote543NoRelatedSym) SbeBlockLength() (blockLength uint) {
	return 9
}

func (*RequestForQuote543NoRelatedSym) SbeSchemaVersion() (schemaVersion uint16) {
	return 5
}
