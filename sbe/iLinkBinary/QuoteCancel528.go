// Generated SBE (Simple Binary Encoding) message codec

package ilinkbinary

import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
)

type QuoteCancel528 struct {
	PartyDetailsListReqID uint64
	SendingTimeEpoch      uint64
	ManualOrderIndicator  ManualOrdIndReqEnum
	SeqNum                uint32
	SenderID              [20]byte
	Location              [5]byte
	QuoteID               uint32
	QuoteCancelType       QuoteCxlTypEnum
	LiquidityFlag         BooleanNULLEnum
	NoQuoteEntries        []QuoteCancel528NoQuoteEntries
	NoQuoteSets           []QuoteCancel528NoQuoteSets
}
type QuoteCancel528NoQuoteEntries struct {
	SecurityGroup [6]byte
	SecurityID    int32
}
type QuoteCancel528NoQuoteSets struct {
	BidSize    uint32
	OfferSize  uint32
	QuoteSetID uint16
}

func (q *QuoteCancel528) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := q.RangeCheck(q.SbeSchemaVersion(), q.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	if err := _m.WriteUint64(_w, q.PartyDetailsListReqID); err != nil {
		return err
	}
	if err := _m.WriteUint64(_w, q.SendingTimeEpoch); err != nil {
		return err
	}
	if err := q.ManualOrderIndicator.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteUint32(_w, q.SeqNum); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, q.SenderID[:]); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, q.Location[:]); err != nil {
		return err
	}
	if err := _m.WriteUint32(_w, q.QuoteID); err != nil {
		return err
	}
	if err := q.QuoteCancelType.Encode(_m, _w); err != nil {
		return err
	}
	if err := q.LiquidityFlag.Encode(_m, _w); err != nil {
		return err
	}
	var NoQuoteEntriesBlockLength uint16 = 10
	if err := _m.WriteUint16(_w, NoQuoteEntriesBlockLength); err != nil {
		return err
	}
	var NoQuoteEntriesNumInGroup uint8 = uint8(len(q.NoQuoteEntries))
	if err := _m.WriteUint8(_w, NoQuoteEntriesNumInGroup); err != nil {
		return err
	}
	for _, prop := range q.NoQuoteEntries {
		if err := prop.Encode(_m, _w); err != nil {
			return err
		}
	}
	var NoQuoteSetsBlockLength uint16 = 10
	if err := _m.WriteUint16(_w, NoQuoteSetsBlockLength); err != nil {
		return err
	}
	var NoQuoteSetsNumInGroup uint8 = uint8(len(q.NoQuoteSets))
	if err := _m.WriteUint8(_w, NoQuoteSetsNumInGroup); err != nil {
		return err
	}
	for _, prop := range q.NoQuoteSets {
		if err := prop.Encode(_m, _w); err != nil {
			return err
		}
	}
	return nil
}

func (q *QuoteCancel528) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	if !q.PartyDetailsListReqIDInActingVersion(actingVersion) {
		q.PartyDetailsListReqID = q.PartyDetailsListReqIDNullValue()
	} else {
		if err := _m.ReadUint64(_r, &q.PartyDetailsListReqID); err != nil {
			return err
		}
	}
	if !q.SendingTimeEpochInActingVersion(actingVersion) {
		q.SendingTimeEpoch = q.SendingTimeEpochNullValue()
	} else {
		if err := _m.ReadUint64(_r, &q.SendingTimeEpoch); err != nil {
			return err
		}
	}
	if q.ManualOrderIndicatorInActingVersion(actingVersion) {
		if err := q.ManualOrderIndicator.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if !q.SeqNumInActingVersion(actingVersion) {
		q.SeqNum = q.SeqNumNullValue()
	} else {
		if err := _m.ReadUint32(_r, &q.SeqNum); err != nil {
			return err
		}
	}
	if !q.SenderIDInActingVersion(actingVersion) {
		for idx := 0; idx < 20; idx++ {
			q.SenderID[idx] = q.SenderIDNullValue()
		}
	} else {
		if err := _m.ReadBytes(_r, q.SenderID[:]); err != nil {
			return err
		}
	}
	if !q.LocationInActingVersion(actingVersion) {
		for idx := 0; idx < 5; idx++ {
			q.Location[idx] = q.LocationNullValue()
		}
	} else {
		if err := _m.ReadBytes(_r, q.Location[:]); err != nil {
			return err
		}
	}
	if !q.QuoteIDInActingVersion(actingVersion) {
		q.QuoteID = q.QuoteIDNullValue()
	} else {
		if err := _m.ReadUint32(_r, &q.QuoteID); err != nil {
			return err
		}
	}
	if q.QuoteCancelTypeInActingVersion(actingVersion) {
		if err := q.QuoteCancelType.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if q.LiquidityFlagInActingVersion(actingVersion) {
		if err := q.LiquidityFlag.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if actingVersion > q.SbeSchemaVersion() && blockLength > q.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-q.SbeBlockLength()))
	}

	if q.NoQuoteEntriesInActingVersion(actingVersion) {
		var NoQuoteEntriesBlockLength uint16
		if err := _m.ReadUint16(_r, &NoQuoteEntriesBlockLength); err != nil {
			return err
		}
		var NoQuoteEntriesNumInGroup uint8
		if err := _m.ReadUint8(_r, &NoQuoteEntriesNumInGroup); err != nil {
			return err
		}
		if cap(q.NoQuoteEntries) < int(NoQuoteEntriesNumInGroup) {
			q.NoQuoteEntries = make([]QuoteCancel528NoQuoteEntries, NoQuoteEntriesNumInGroup)
		}
		for i, _ := range q.NoQuoteEntries {
			if err := q.NoQuoteEntries[i].Decode(_m, _r, actingVersion, uint(NoQuoteEntriesBlockLength)); err != nil {
				return err
			}
		}
	}

	if q.NoQuoteSetsInActingVersion(actingVersion) {
		var NoQuoteSetsBlockLength uint16
		if err := _m.ReadUint16(_r, &NoQuoteSetsBlockLength); err != nil {
			return err
		}
		var NoQuoteSetsNumInGroup uint8
		if err := _m.ReadUint8(_r, &NoQuoteSetsNumInGroup); err != nil {
			return err
		}
		if cap(q.NoQuoteSets) < int(NoQuoteSetsNumInGroup) {
			q.NoQuoteSets = make([]QuoteCancel528NoQuoteSets, NoQuoteSetsNumInGroup)
		}
		for i, _ := range q.NoQuoteSets {
			if err := q.NoQuoteSets[i].Decode(_m, _r, actingVersion, uint(NoQuoteSetsBlockLength)); err != nil {
				return err
			}
		}
	}
	if doRangeCheck {
		if err := q.RangeCheck(actingVersion, q.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	return nil
}

func (q *QuoteCancel528) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if q.PartyDetailsListReqIDInActingVersion(actingVersion) {
		if q.PartyDetailsListReqID < q.PartyDetailsListReqIDMinValue() || q.PartyDetailsListReqID > q.PartyDetailsListReqIDMaxValue() {
			return fmt.Errorf("Range check failed on q.PartyDetailsListReqID (%v < %v > %v)", q.PartyDetailsListReqIDMinValue(), q.PartyDetailsListReqID, q.PartyDetailsListReqIDMaxValue())
		}
	}
	if q.SendingTimeEpochInActingVersion(actingVersion) {
		if q.SendingTimeEpoch < q.SendingTimeEpochMinValue() || q.SendingTimeEpoch > q.SendingTimeEpochMaxValue() {
			return fmt.Errorf("Range check failed on q.SendingTimeEpoch (%v < %v > %v)", q.SendingTimeEpochMinValue(), q.SendingTimeEpoch, q.SendingTimeEpochMaxValue())
		}
	}
	if err := q.ManualOrderIndicator.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if q.SeqNumInActingVersion(actingVersion) {
		if q.SeqNum < q.SeqNumMinValue() || q.SeqNum > q.SeqNumMaxValue() {
			return fmt.Errorf("Range check failed on q.SeqNum (%v < %v > %v)", q.SeqNumMinValue(), q.SeqNum, q.SeqNumMaxValue())
		}
	}
	if q.SenderIDInActingVersion(actingVersion) {
		for idx := 0; idx < 20; idx++ {
			if q.SenderID[idx] < q.SenderIDMinValue() || q.SenderID[idx] > q.SenderIDMaxValue() {
				return fmt.Errorf("Range check failed on q.SenderID[%d] (%v < %v > %v)", idx, q.SenderIDMinValue(), q.SenderID[idx], q.SenderIDMaxValue())
			}
		}
	}
	if q.LocationInActingVersion(actingVersion) {
		for idx := 0; idx < 5; idx++ {
			if q.Location[idx] < q.LocationMinValue() || q.Location[idx] > q.LocationMaxValue() {
				return fmt.Errorf("Range check failed on q.Location[%d] (%v < %v > %v)", idx, q.LocationMinValue(), q.Location[idx], q.LocationMaxValue())
			}
		}
	}
	if q.QuoteIDInActingVersion(actingVersion) {
		if q.QuoteID < q.QuoteIDMinValue() || q.QuoteID > q.QuoteIDMaxValue() {
			return fmt.Errorf("Range check failed on q.QuoteID (%v < %v > %v)", q.QuoteIDMinValue(), q.QuoteID, q.QuoteIDMaxValue())
		}
	}
	if err := q.QuoteCancelType.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if err := q.LiquidityFlag.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	for _, prop := range q.NoQuoteEntries {
		if err := prop.RangeCheck(actingVersion, schemaVersion); err != nil {
			return err
		}
	}
	for _, prop := range q.NoQuoteSets {
		if err := prop.RangeCheck(actingVersion, schemaVersion); err != nil {
			return err
		}
	}
	return nil
}

func QuoteCancel528Init(q *QuoteCancel528) {
	return
}

func (q *QuoteCancel528NoQuoteEntries) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteBytes(_w, q.SecurityGroup[:]); err != nil {
		return err
	}
	if err := _m.WriteInt32(_w, q.SecurityID); err != nil {
		return err
	}
	return nil
}

func (q *QuoteCancel528NoQuoteEntries) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint) error {
	if !q.SecurityGroupInActingVersion(actingVersion) {
		for idx := 0; idx < 6; idx++ {
			q.SecurityGroup[idx] = q.SecurityGroupNullValue()
		}
	} else {
		if err := _m.ReadBytes(_r, q.SecurityGroup[:]); err != nil {
			return err
		}
	}
	if !q.SecurityIDInActingVersion(actingVersion) {
		q.SecurityID = q.SecurityIDNullValue()
	} else {
		if err := _m.ReadInt32(_r, &q.SecurityID); err != nil {
			return err
		}
	}
	if actingVersion > q.SbeSchemaVersion() && blockLength > q.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-q.SbeBlockLength()))
	}
	return nil
}

func (q *QuoteCancel528NoQuoteEntries) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if q.SecurityGroupInActingVersion(actingVersion) {
		for idx := 0; idx < 6; idx++ {
			if q.SecurityGroup[idx] < q.SecurityGroupMinValue() || q.SecurityGroup[idx] > q.SecurityGroupMaxValue() {
				return fmt.Errorf("Range check failed on q.SecurityGroup[%d] (%v < %v > %v)", idx, q.SecurityGroupMinValue(), q.SecurityGroup[idx], q.SecurityGroupMaxValue())
			}
		}
	}
	if q.SecurityIDInActingVersion(actingVersion) {
		if q.SecurityID != q.SecurityIDNullValue() && (q.SecurityID < q.SecurityIDMinValue() || q.SecurityID > q.SecurityIDMaxValue()) {
			return fmt.Errorf("Range check failed on q.SecurityID (%v < %v > %v)", q.SecurityIDMinValue(), q.SecurityID, q.SecurityIDMaxValue())
		}
	}
	return nil
}

func QuoteCancel528NoQuoteEntriesInit(q *QuoteCancel528NoQuoteEntries) {
	for idx := 0; idx < 6; idx++ {
		q.SecurityGroup[idx] = 0
	}
	q.SecurityID = 2147483647
	return
}

func (q *QuoteCancel528NoQuoteSets) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint32(_w, q.BidSize); err != nil {
		return err
	}
	if err := _m.WriteUint32(_w, q.OfferSize); err != nil {
		return err
	}
	if err := _m.WriteUint16(_w, q.QuoteSetID); err != nil {
		return err
	}
	return nil
}

func (q *QuoteCancel528NoQuoteSets) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint) error {
	if !q.BidSizeInActingVersion(actingVersion) {
		q.BidSize = q.BidSizeNullValue()
	} else {
		if err := _m.ReadUint32(_r, &q.BidSize); err != nil {
			return err
		}
	}
	if !q.OfferSizeInActingVersion(actingVersion) {
		q.OfferSize = q.OfferSizeNullValue()
	} else {
		if err := _m.ReadUint32(_r, &q.OfferSize); err != nil {
			return err
		}
	}
	if !q.QuoteSetIDInActingVersion(actingVersion) {
		q.QuoteSetID = q.QuoteSetIDNullValue()
	} else {
		if err := _m.ReadUint16(_r, &q.QuoteSetID); err != nil {
			return err
		}
	}
	if actingVersion > q.SbeSchemaVersion() && blockLength > q.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-q.SbeBlockLength()))
	}
	return nil
}

func (q *QuoteCancel528NoQuoteSets) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if q.BidSizeInActingVersion(actingVersion) {
		if q.BidSize != q.BidSizeNullValue() && (q.BidSize < q.BidSizeMinValue() || q.BidSize > q.BidSizeMaxValue()) {
			return fmt.Errorf("Range check failed on q.BidSize (%v < %v > %v)", q.BidSizeMinValue(), q.BidSize, q.BidSizeMaxValue())
		}
	}
	if q.OfferSizeInActingVersion(actingVersion) {
		if q.OfferSize != q.OfferSizeNullValue() && (q.OfferSize < q.OfferSizeMinValue() || q.OfferSize > q.OfferSizeMaxValue()) {
			return fmt.Errorf("Range check failed on q.OfferSize (%v < %v > %v)", q.OfferSizeMinValue(), q.OfferSize, q.OfferSizeMaxValue())
		}
	}
	if q.QuoteSetIDInActingVersion(actingVersion) {
		if q.QuoteSetID < q.QuoteSetIDMinValue() || q.QuoteSetID > q.QuoteSetIDMaxValue() {
			return fmt.Errorf("Range check failed on q.QuoteSetID (%v < %v > %v)", q.QuoteSetIDMinValue(), q.QuoteSetID, q.QuoteSetIDMaxValue())
		}
	}
	return nil
}

func QuoteCancel528NoQuoteSetsInit(q *QuoteCancel528NoQuoteSets) {
	q.BidSize = 4294967295
	q.OfferSize = 4294967295
	return
}

func (*QuoteCancel528) SbeBlockLength() (blockLength uint16) {
	return 52
}

func (*QuoteCancel528) SbeTemplateId() (templateId uint16) {
	return 528
}

func (*QuoteCancel528) SbeSchemaId() (schemaId uint16) {
	return 8
}

func (*QuoteCancel528) SbeSchemaVersion() (schemaVersion uint16) {
	return 5
}

func (*QuoteCancel528) SbeSemanticType() (semanticType []byte) {
	return []byte("Z")
}

func (*QuoteCancel528) PartyDetailsListReqIDId() uint16 {
	return 1505
}

func (*QuoteCancel528) PartyDetailsListReqIDSinceVersion() uint16 {
	return 0
}

func (q *QuoteCancel528) PartyDetailsListReqIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= q.PartyDetailsListReqIDSinceVersion()
}

func (*QuoteCancel528) PartyDetailsListReqIDDeprecated() uint16 {
	return 0
}

func (*QuoteCancel528) PartyDetailsListReqIDMetaAttribute(meta int) string {
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

func (*QuoteCancel528) PartyDetailsListReqIDMinValue() uint64 {
	return 0
}

func (*QuoteCancel528) PartyDetailsListReqIDMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*QuoteCancel528) PartyDetailsListReqIDNullValue() uint64 {
	return math.MaxUint64
}

func (*QuoteCancel528) SendingTimeEpochId() uint16 {
	return 5297
}

func (*QuoteCancel528) SendingTimeEpochSinceVersion() uint16 {
	return 0
}

func (q *QuoteCancel528) SendingTimeEpochInActingVersion(actingVersion uint16) bool {
	return actingVersion >= q.SendingTimeEpochSinceVersion()
}

func (*QuoteCancel528) SendingTimeEpochDeprecated() uint16 {
	return 0
}

func (*QuoteCancel528) SendingTimeEpochMetaAttribute(meta int) string {
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

func (*QuoteCancel528) SendingTimeEpochMinValue() uint64 {
	return 0
}

func (*QuoteCancel528) SendingTimeEpochMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*QuoteCancel528) SendingTimeEpochNullValue() uint64 {
	return math.MaxUint64
}

func (*QuoteCancel528) ManualOrderIndicatorId() uint16 {
	return 1028
}

func (*QuoteCancel528) ManualOrderIndicatorSinceVersion() uint16 {
	return 0
}

func (q *QuoteCancel528) ManualOrderIndicatorInActingVersion(actingVersion uint16) bool {
	return actingVersion >= q.ManualOrderIndicatorSinceVersion()
}

func (*QuoteCancel528) ManualOrderIndicatorDeprecated() uint16 {
	return 0
}

func (*QuoteCancel528) ManualOrderIndicatorMetaAttribute(meta int) string {
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

func (*QuoteCancel528) SeqNumId() uint16 {
	return 9726
}

func (*QuoteCancel528) SeqNumSinceVersion() uint16 {
	return 0
}

func (q *QuoteCancel528) SeqNumInActingVersion(actingVersion uint16) bool {
	return actingVersion >= q.SeqNumSinceVersion()
}

func (*QuoteCancel528) SeqNumDeprecated() uint16 {
	return 0
}

func (*QuoteCancel528) SeqNumMetaAttribute(meta int) string {
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

func (*QuoteCancel528) SeqNumMinValue() uint32 {
	return 0
}

func (*QuoteCancel528) SeqNumMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*QuoteCancel528) SeqNumNullValue() uint32 {
	return math.MaxUint32
}

func (*QuoteCancel528) SenderIDId() uint16 {
	return 5392
}

func (*QuoteCancel528) SenderIDSinceVersion() uint16 {
	return 0
}

func (q *QuoteCancel528) SenderIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= q.SenderIDSinceVersion()
}

func (*QuoteCancel528) SenderIDDeprecated() uint16 {
	return 0
}

func (*QuoteCancel528) SenderIDMetaAttribute(meta int) string {
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

func (*QuoteCancel528) SenderIDMinValue() byte {
	return byte(32)
}

func (*QuoteCancel528) SenderIDMaxValue() byte {
	return byte(126)
}

func (*QuoteCancel528) SenderIDNullValue() byte {
	return 0
}

func (q *QuoteCancel528) SenderIDCharacterEncoding() string {
	return "US-ASCII"
}

func (*QuoteCancel528) LocationId() uint16 {
	return 9537
}

func (*QuoteCancel528) LocationSinceVersion() uint16 {
	return 0
}

func (q *QuoteCancel528) LocationInActingVersion(actingVersion uint16) bool {
	return actingVersion >= q.LocationSinceVersion()
}

func (*QuoteCancel528) LocationDeprecated() uint16 {
	return 0
}

func (*QuoteCancel528) LocationMetaAttribute(meta int) string {
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

func (*QuoteCancel528) LocationMinValue() byte {
	return byte(32)
}

func (*QuoteCancel528) LocationMaxValue() byte {
	return byte(126)
}

func (*QuoteCancel528) LocationNullValue() byte {
	return 0
}

func (q *QuoteCancel528) LocationCharacterEncoding() string {
	return "US-ASCII"
}

func (*QuoteCancel528) QuoteIDId() uint16 {
	return 117
}

func (*QuoteCancel528) QuoteIDSinceVersion() uint16 {
	return 0
}

func (q *QuoteCancel528) QuoteIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= q.QuoteIDSinceVersion()
}

func (*QuoteCancel528) QuoteIDDeprecated() uint16 {
	return 0
}

func (*QuoteCancel528) QuoteIDMetaAttribute(meta int) string {
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

func (*QuoteCancel528) QuoteIDMinValue() uint32 {
	return 0
}

func (*QuoteCancel528) QuoteIDMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*QuoteCancel528) QuoteIDNullValue() uint32 {
	return math.MaxUint32
}

func (*QuoteCancel528) QuoteCancelTypeId() uint16 {
	return 298
}

func (*QuoteCancel528) QuoteCancelTypeSinceVersion() uint16 {
	return 0
}

func (q *QuoteCancel528) QuoteCancelTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= q.QuoteCancelTypeSinceVersion()
}

func (*QuoteCancel528) QuoteCancelTypeDeprecated() uint16 {
	return 0
}

func (*QuoteCancel528) QuoteCancelTypeMetaAttribute(meta int) string {
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

func (*QuoteCancel528) LiquidityFlagId() uint16 {
	return 9373
}

func (*QuoteCancel528) LiquidityFlagSinceVersion() uint16 {
	return 0
}

func (q *QuoteCancel528) LiquidityFlagInActingVersion(actingVersion uint16) bool {
	return actingVersion >= q.LiquidityFlagSinceVersion()
}

func (*QuoteCancel528) LiquidityFlagDeprecated() uint16 {
	return 0
}

func (*QuoteCancel528) LiquidityFlagMetaAttribute(meta int) string {
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

func (*QuoteCancel528NoQuoteEntries) SecurityGroupId() uint16 {
	return 1151
}

func (*QuoteCancel528NoQuoteEntries) SecurityGroupSinceVersion() uint16 {
	return 0
}

func (q *QuoteCancel528NoQuoteEntries) SecurityGroupInActingVersion(actingVersion uint16) bool {
	return actingVersion >= q.SecurityGroupSinceVersion()
}

func (*QuoteCancel528NoQuoteEntries) SecurityGroupDeprecated() uint16 {
	return 0
}

func (*QuoteCancel528NoQuoteEntries) SecurityGroupMetaAttribute(meta int) string {
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

func (*QuoteCancel528NoQuoteEntries) SecurityGroupMinValue() byte {
	return byte(32)
}

func (*QuoteCancel528NoQuoteEntries) SecurityGroupMaxValue() byte {
	return byte(126)
}

func (*QuoteCancel528NoQuoteEntries) SecurityGroupNullValue() byte {
	return 0
}

func (q *QuoteCancel528NoQuoteEntries) SecurityGroupCharacterEncoding() string {
	return "US-ASCII"
}

func (*QuoteCancel528NoQuoteEntries) SecurityIDId() uint16 {
	return 48
}

func (*QuoteCancel528NoQuoteEntries) SecurityIDSinceVersion() uint16 {
	return 0
}

func (q *QuoteCancel528NoQuoteEntries) SecurityIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= q.SecurityIDSinceVersion()
}

func (*QuoteCancel528NoQuoteEntries) SecurityIDDeprecated() uint16 {
	return 0
}

func (*QuoteCancel528NoQuoteEntries) SecurityIDMetaAttribute(meta int) string {
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

func (*QuoteCancel528NoQuoteEntries) SecurityIDMinValue() int32 {
	return math.MinInt32 + 1
}

func (*QuoteCancel528NoQuoteEntries) SecurityIDMaxValue() int32 {
	return math.MaxInt32
}

func (*QuoteCancel528NoQuoteEntries) SecurityIDNullValue() int32 {
	return 2147483647
}

func (*QuoteCancel528NoQuoteSets) BidSizeId() uint16 {
	return 134
}

func (*QuoteCancel528NoQuoteSets) BidSizeSinceVersion() uint16 {
	return 0
}

func (q *QuoteCancel528NoQuoteSets) BidSizeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= q.BidSizeSinceVersion()
}

func (*QuoteCancel528NoQuoteSets) BidSizeDeprecated() uint16 {
	return 0
}

func (*QuoteCancel528NoQuoteSets) BidSizeMetaAttribute(meta int) string {
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

func (*QuoteCancel528NoQuoteSets) BidSizeMinValue() uint32 {
	return 0
}

func (*QuoteCancel528NoQuoteSets) BidSizeMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*QuoteCancel528NoQuoteSets) BidSizeNullValue() uint32 {
	return 4294967295
}

func (*QuoteCancel528NoQuoteSets) OfferSizeId() uint16 {
	return 135
}

func (*QuoteCancel528NoQuoteSets) OfferSizeSinceVersion() uint16 {
	return 0
}

func (q *QuoteCancel528NoQuoteSets) OfferSizeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= q.OfferSizeSinceVersion()
}

func (*QuoteCancel528NoQuoteSets) OfferSizeDeprecated() uint16 {
	return 0
}

func (*QuoteCancel528NoQuoteSets) OfferSizeMetaAttribute(meta int) string {
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

func (*QuoteCancel528NoQuoteSets) OfferSizeMinValue() uint32 {
	return 0
}

func (*QuoteCancel528NoQuoteSets) OfferSizeMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*QuoteCancel528NoQuoteSets) OfferSizeNullValue() uint32 {
	return 4294967295
}

func (*QuoteCancel528NoQuoteSets) QuoteSetIDId() uint16 {
	return 302
}

func (*QuoteCancel528NoQuoteSets) QuoteSetIDSinceVersion() uint16 {
	return 0
}

func (q *QuoteCancel528NoQuoteSets) QuoteSetIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= q.QuoteSetIDSinceVersion()
}

func (*QuoteCancel528NoQuoteSets) QuoteSetIDDeprecated() uint16 {
	return 0
}

func (*QuoteCancel528NoQuoteSets) QuoteSetIDMetaAttribute(meta int) string {
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

func (*QuoteCancel528NoQuoteSets) QuoteSetIDMinValue() uint16 {
	return 0
}

func (*QuoteCancel528NoQuoteSets) QuoteSetIDMaxValue() uint16 {
	return math.MaxUint16 - 1
}

func (*QuoteCancel528NoQuoteSets) QuoteSetIDNullValue() uint16 {
	return math.MaxUint16
}

func (*QuoteCancel528) NoQuoteEntriesId() uint16 {
	return 295
}

func (*QuoteCancel528) NoQuoteEntriesSinceVersion() uint16 {
	return 0
}

func (q *QuoteCancel528) NoQuoteEntriesInActingVersion(actingVersion uint16) bool {
	return actingVersion >= q.NoQuoteEntriesSinceVersion()
}

func (*QuoteCancel528) NoQuoteEntriesDeprecated() uint16 {
	return 0
}

func (*QuoteCancel528NoQuoteEntries) SbeBlockLength() (blockLength uint) {
	return 10
}

func (*QuoteCancel528NoQuoteEntries) SbeSchemaVersion() (schemaVersion uint16) {
	return 5
}

func (*QuoteCancel528) NoQuoteSetsId() uint16 {
	return 296
}

func (*QuoteCancel528) NoQuoteSetsSinceVersion() uint16 {
	return 0
}

func (q *QuoteCancel528) NoQuoteSetsInActingVersion(actingVersion uint16) bool {
	return actingVersion >= q.NoQuoteSetsSinceVersion()
}

func (*QuoteCancel528) NoQuoteSetsDeprecated() uint16 {
	return 0
}

func (*QuoteCancel528NoQuoteSets) SbeBlockLength() (blockLength uint) {
	return 10
}

func (*QuoteCancel528NoQuoteSets) SbeSchemaVersion() (schemaVersion uint16) {
	return 5
}
