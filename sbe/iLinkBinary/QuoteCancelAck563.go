// Generated SBE (Simple Binary Encoding) message codec

package ilinkbinary

import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
)

type QuoteCancelAck563 struct {
	SeqNum                uint32
	UUID                  uint64
	Text                  [256]byte
	SenderID              [20]byte
	PartyDetailsListReqID uint64
	RequestTime           uint64
	SendingTimeEpoch      uint64
	CancelledSymbol       [6]byte
	Location              [5]byte
	QuoteID               uint32
	QuoteRejectReason     uint16
	DelayDuration         uint16
	ManualOrderIndicator  ManualOrdIndReqEnum
	QuoteStatus           QuoteCxlStatusEnum
	NoProcessedEntries    uint32
	MMProtectionReset     BooleanFlagEnum
	UnsolicitedCancelType byte
	SplitMsg              SplitMsgEnum
	TotNoQuoteEntries     uint8
	LiquidityFlag         BooleanNULLEnum
	PossRetransFlag       BooleanFlagEnum
	DelayToTime           uint64
	NoQuoteEntries        []QuoteCancelAck563NoQuoteEntries
	NoQuoteSets           []QuoteCancelAck563NoQuoteSets
}
type QuoteCancelAck563NoQuoteEntries struct {
	QuoteEntryID           uint32
	SecurityID             int32
	QuoteEntryRejectReason uint8
}
type QuoteCancelAck563NoQuoteSets struct {
	QuoteSetID     uint16
	QuoteErrorCode uint16
}

func (q *QuoteCancelAck563) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := q.RangeCheck(q.SbeSchemaVersion(), q.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	if err := _m.WriteUint32(_w, q.SeqNum); err != nil {
		return err
	}
	if err := _m.WriteUint64(_w, q.UUID); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, q.Text[:]); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, q.SenderID[:]); err != nil {
		return err
	}
	if err := _m.WriteUint64(_w, q.PartyDetailsListReqID); err != nil {
		return err
	}
	if err := _m.WriteUint64(_w, q.RequestTime); err != nil {
		return err
	}
	if err := _m.WriteUint64(_w, q.SendingTimeEpoch); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, q.CancelledSymbol[:]); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, q.Location[:]); err != nil {
		return err
	}
	if err := _m.WriteUint32(_w, q.QuoteID); err != nil {
		return err
	}
	if err := _m.WriteUint16(_w, q.QuoteRejectReason); err != nil {
		return err
	}
	if err := _m.WriteUint16(_w, q.DelayDuration); err != nil {
		return err
	}
	if err := q.ManualOrderIndicator.Encode(_m, _w); err != nil {
		return err
	}
	if err := q.QuoteStatus.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteUint32(_w, q.NoProcessedEntries); err != nil {
		return err
	}
	if err := q.MMProtectionReset.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteUint8(_w, q.UnsolicitedCancelType); err != nil {
		return err
	}
	if err := q.SplitMsg.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteUint8(_w, q.TotNoQuoteEntries); err != nil {
		return err
	}
	if err := q.LiquidityFlag.Encode(_m, _w); err != nil {
		return err
	}
	if err := q.PossRetransFlag.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteUint64(_w, q.DelayToTime); err != nil {
		return err
	}
	var NoQuoteEntriesBlockLength uint16 = 9
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
	var NoQuoteSetsBlockLength uint16 = 4
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

func (q *QuoteCancelAck563) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	if !q.SeqNumInActingVersion(actingVersion) {
		q.SeqNum = q.SeqNumNullValue()
	} else {
		if err := _m.ReadUint32(_r, &q.SeqNum); err != nil {
			return err
		}
	}
	if !q.UUIDInActingVersion(actingVersion) {
		q.UUID = q.UUIDNullValue()
	} else {
		if err := _m.ReadUint64(_r, &q.UUID); err != nil {
			return err
		}
	}
	if !q.TextInActingVersion(actingVersion) {
		for idx := 0; idx < 256; idx++ {
			q.Text[idx] = q.TextNullValue()
		}
	} else {
		if err := _m.ReadBytes(_r, q.Text[:]); err != nil {
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
	if !q.PartyDetailsListReqIDInActingVersion(actingVersion) {
		q.PartyDetailsListReqID = q.PartyDetailsListReqIDNullValue()
	} else {
		if err := _m.ReadUint64(_r, &q.PartyDetailsListReqID); err != nil {
			return err
		}
	}
	if !q.RequestTimeInActingVersion(actingVersion) {
		q.RequestTime = q.RequestTimeNullValue()
	} else {
		if err := _m.ReadUint64(_r, &q.RequestTime); err != nil {
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
	if !q.CancelledSymbolInActingVersion(actingVersion) {
		for idx := 0; idx < 6; idx++ {
			q.CancelledSymbol[idx] = q.CancelledSymbolNullValue()
		}
	} else {
		if err := _m.ReadBytes(_r, q.CancelledSymbol[:]); err != nil {
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
	if !q.QuoteRejectReasonInActingVersion(actingVersion) {
		q.QuoteRejectReason = q.QuoteRejectReasonNullValue()
	} else {
		if err := _m.ReadUint16(_r, &q.QuoteRejectReason); err != nil {
			return err
		}
	}
	if !q.DelayDurationInActingVersion(actingVersion) {
		q.DelayDuration = q.DelayDurationNullValue()
	} else {
		if err := _m.ReadUint16(_r, &q.DelayDuration); err != nil {
			return err
		}
	}
	if q.ManualOrderIndicatorInActingVersion(actingVersion) {
		if err := q.ManualOrderIndicator.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if q.QuoteStatusInActingVersion(actingVersion) {
		if err := q.QuoteStatus.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if !q.NoProcessedEntriesInActingVersion(actingVersion) {
		q.NoProcessedEntries = q.NoProcessedEntriesNullValue()
	} else {
		if err := _m.ReadUint32(_r, &q.NoProcessedEntries); err != nil {
			return err
		}
	}
	if q.MMProtectionResetInActingVersion(actingVersion) {
		if err := q.MMProtectionReset.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if !q.UnsolicitedCancelTypeInActingVersion(actingVersion) {
		q.UnsolicitedCancelType = q.UnsolicitedCancelTypeNullValue()
	} else {
		if err := _m.ReadUint8(_r, &q.UnsolicitedCancelType); err != nil {
			return err
		}
	}
	if q.SplitMsgInActingVersion(actingVersion) {
		if err := q.SplitMsg.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if !q.TotNoQuoteEntriesInActingVersion(actingVersion) {
		q.TotNoQuoteEntries = q.TotNoQuoteEntriesNullValue()
	} else {
		if err := _m.ReadUint8(_r, &q.TotNoQuoteEntries); err != nil {
			return err
		}
	}
	if q.LiquidityFlagInActingVersion(actingVersion) {
		if err := q.LiquidityFlag.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if q.PossRetransFlagInActingVersion(actingVersion) {
		if err := q.PossRetransFlag.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if !q.DelayToTimeInActingVersion(actingVersion) {
		q.DelayToTime = q.DelayToTimeNullValue()
	} else {
		if err := _m.ReadUint64(_r, &q.DelayToTime); err != nil {
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
			q.NoQuoteEntries = make([]QuoteCancelAck563NoQuoteEntries, NoQuoteEntriesNumInGroup)
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
			q.NoQuoteSets = make([]QuoteCancelAck563NoQuoteSets, NoQuoteSetsNumInGroup)
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

func (q *QuoteCancelAck563) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if q.SeqNumInActingVersion(actingVersion) {
		if q.SeqNum < q.SeqNumMinValue() || q.SeqNum > q.SeqNumMaxValue() {
			return fmt.Errorf("Range check failed on q.SeqNum (%v < %v > %v)", q.SeqNumMinValue(), q.SeqNum, q.SeqNumMaxValue())
		}
	}
	if q.UUIDInActingVersion(actingVersion) {
		if q.UUID < q.UUIDMinValue() || q.UUID > q.UUIDMaxValue() {
			return fmt.Errorf("Range check failed on q.UUID (%v < %v > %v)", q.UUIDMinValue(), q.UUID, q.UUIDMaxValue())
		}
	}
	if q.TextInActingVersion(actingVersion) {
		for idx := 0; idx < 256; idx++ {
			if q.Text[idx] < q.TextMinValue() || q.Text[idx] > q.TextMaxValue() {
				return fmt.Errorf("Range check failed on q.Text[%d] (%v < %v > %v)", idx, q.TextMinValue(), q.Text[idx], q.TextMaxValue())
			}
		}
	}
	if q.SenderIDInActingVersion(actingVersion) {
		for idx := 0; idx < 20; idx++ {
			if q.SenderID[idx] < q.SenderIDMinValue() || q.SenderID[idx] > q.SenderIDMaxValue() {
				return fmt.Errorf("Range check failed on q.SenderID[%d] (%v < %v > %v)", idx, q.SenderIDMinValue(), q.SenderID[idx], q.SenderIDMaxValue())
			}
		}
	}
	if q.PartyDetailsListReqIDInActingVersion(actingVersion) {
		if q.PartyDetailsListReqID < q.PartyDetailsListReqIDMinValue() || q.PartyDetailsListReqID > q.PartyDetailsListReqIDMaxValue() {
			return fmt.Errorf("Range check failed on q.PartyDetailsListReqID (%v < %v > %v)", q.PartyDetailsListReqIDMinValue(), q.PartyDetailsListReqID, q.PartyDetailsListReqIDMaxValue())
		}
	}
	if q.RequestTimeInActingVersion(actingVersion) {
		if q.RequestTime < q.RequestTimeMinValue() || q.RequestTime > q.RequestTimeMaxValue() {
			return fmt.Errorf("Range check failed on q.RequestTime (%v < %v > %v)", q.RequestTimeMinValue(), q.RequestTime, q.RequestTimeMaxValue())
		}
	}
	if q.SendingTimeEpochInActingVersion(actingVersion) {
		if q.SendingTimeEpoch < q.SendingTimeEpochMinValue() || q.SendingTimeEpoch > q.SendingTimeEpochMaxValue() {
			return fmt.Errorf("Range check failed on q.SendingTimeEpoch (%v < %v > %v)", q.SendingTimeEpochMinValue(), q.SendingTimeEpoch, q.SendingTimeEpochMaxValue())
		}
	}
	if q.CancelledSymbolInActingVersion(actingVersion) {
		for idx := 0; idx < 6; idx++ {
			if q.CancelledSymbol[idx] < q.CancelledSymbolMinValue() || q.CancelledSymbol[idx] > q.CancelledSymbolMaxValue() {
				return fmt.Errorf("Range check failed on q.CancelledSymbol[%d] (%v < %v > %v)", idx, q.CancelledSymbolMinValue(), q.CancelledSymbol[idx], q.CancelledSymbolMaxValue())
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
	if q.QuoteRejectReasonInActingVersion(actingVersion) {
		if q.QuoteRejectReason != q.QuoteRejectReasonNullValue() && (q.QuoteRejectReason < q.QuoteRejectReasonMinValue() || q.QuoteRejectReason > q.QuoteRejectReasonMaxValue()) {
			return fmt.Errorf("Range check failed on q.QuoteRejectReason (%v < %v > %v)", q.QuoteRejectReasonMinValue(), q.QuoteRejectReason, q.QuoteRejectReasonMaxValue())
		}
	}
	if q.DelayDurationInActingVersion(actingVersion) {
		if q.DelayDuration != q.DelayDurationNullValue() && (q.DelayDuration < q.DelayDurationMinValue() || q.DelayDuration > q.DelayDurationMaxValue()) {
			return fmt.Errorf("Range check failed on q.DelayDuration (%v < %v > %v)", q.DelayDurationMinValue(), q.DelayDuration, q.DelayDurationMaxValue())
		}
	}
	if err := q.ManualOrderIndicator.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if err := q.QuoteStatus.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if q.NoProcessedEntriesInActingVersion(actingVersion) {
		if q.NoProcessedEntries < q.NoProcessedEntriesMinValue() || q.NoProcessedEntries > q.NoProcessedEntriesMaxValue() {
			return fmt.Errorf("Range check failed on q.NoProcessedEntries (%v < %v > %v)", q.NoProcessedEntriesMinValue(), q.NoProcessedEntries, q.NoProcessedEntriesMaxValue())
		}
	}
	if err := q.MMProtectionReset.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if q.UnsolicitedCancelTypeInActingVersion(actingVersion) {
		if q.UnsolicitedCancelType != q.UnsolicitedCancelTypeNullValue() && (q.UnsolicitedCancelType < q.UnsolicitedCancelTypeMinValue() || q.UnsolicitedCancelType > q.UnsolicitedCancelTypeMaxValue()) {
			return fmt.Errorf("Range check failed on q.UnsolicitedCancelType (%v < %v > %v)", q.UnsolicitedCancelTypeMinValue(), q.UnsolicitedCancelType, q.UnsolicitedCancelTypeMaxValue())
		}
	}
	if err := q.SplitMsg.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if q.TotNoQuoteEntriesInActingVersion(actingVersion) {
		if q.TotNoQuoteEntries != q.TotNoQuoteEntriesNullValue() && (q.TotNoQuoteEntries < q.TotNoQuoteEntriesMinValue() || q.TotNoQuoteEntries > q.TotNoQuoteEntriesMaxValue()) {
			return fmt.Errorf("Range check failed on q.TotNoQuoteEntries (%v < %v > %v)", q.TotNoQuoteEntriesMinValue(), q.TotNoQuoteEntries, q.TotNoQuoteEntriesMaxValue())
		}
	}
	if err := q.LiquidityFlag.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if err := q.PossRetransFlag.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if q.DelayToTimeInActingVersion(actingVersion) {
		if q.DelayToTime != q.DelayToTimeNullValue() && (q.DelayToTime < q.DelayToTimeMinValue() || q.DelayToTime > q.DelayToTimeMaxValue()) {
			return fmt.Errorf("Range check failed on q.DelayToTime (%v < %v > %v)", q.DelayToTimeMinValue(), q.DelayToTime, q.DelayToTimeMaxValue())
		}
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

func QuoteCancelAck563Init(q *QuoteCancelAck563) {
	for idx := 0; idx < 256; idx++ {
		q.Text[idx] = 0
	}
	for idx := 0; idx < 6; idx++ {
		q.CancelledSymbol[idx] = 0
	}
	q.QuoteRejectReason = 65535
	q.DelayDuration = 65535
	q.UnsolicitedCancelType = 48
	q.TotNoQuoteEntries = 255
	q.DelayToTime = 18446744073709551615
	return
}

func (q *QuoteCancelAck563NoQuoteEntries) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint32(_w, q.QuoteEntryID); err != nil {
		return err
	}
	if err := _m.WriteInt32(_w, q.SecurityID); err != nil {
		return err
	}
	if err := _m.WriteUint8(_w, q.QuoteEntryRejectReason); err != nil {
		return err
	}
	return nil
}

func (q *QuoteCancelAck563NoQuoteEntries) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint) error {
	if !q.QuoteEntryIDInActingVersion(actingVersion) {
		q.QuoteEntryID = q.QuoteEntryIDNullValue()
	} else {
		if err := _m.ReadUint32(_r, &q.QuoteEntryID); err != nil {
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
	if !q.QuoteEntryRejectReasonInActingVersion(actingVersion) {
		q.QuoteEntryRejectReason = q.QuoteEntryRejectReasonNullValue()
	} else {
		if err := _m.ReadUint8(_r, &q.QuoteEntryRejectReason); err != nil {
			return err
		}
	}
	if actingVersion > q.SbeSchemaVersion() && blockLength > q.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-q.SbeBlockLength()))
	}
	return nil
}

func (q *QuoteCancelAck563NoQuoteEntries) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if q.QuoteEntryIDInActingVersion(actingVersion) {
		if q.QuoteEntryID < q.QuoteEntryIDMinValue() || q.QuoteEntryID > q.QuoteEntryIDMaxValue() {
			return fmt.Errorf("Range check failed on q.QuoteEntryID (%v < %v > %v)", q.QuoteEntryIDMinValue(), q.QuoteEntryID, q.QuoteEntryIDMaxValue())
		}
	}
	if q.SecurityIDInActingVersion(actingVersion) {
		if q.SecurityID < q.SecurityIDMinValue() || q.SecurityID > q.SecurityIDMaxValue() {
			return fmt.Errorf("Range check failed on q.SecurityID (%v < %v > %v)", q.SecurityIDMinValue(), q.SecurityID, q.SecurityIDMaxValue())
		}
	}
	if q.QuoteEntryRejectReasonInActingVersion(actingVersion) {
		if q.QuoteEntryRejectReason < q.QuoteEntryRejectReasonMinValue() || q.QuoteEntryRejectReason > q.QuoteEntryRejectReasonMaxValue() {
			return fmt.Errorf("Range check failed on q.QuoteEntryRejectReason (%v < %v > %v)", q.QuoteEntryRejectReasonMinValue(), q.QuoteEntryRejectReason, q.QuoteEntryRejectReasonMaxValue())
		}
	}
	return nil
}

func QuoteCancelAck563NoQuoteEntriesInit(q *QuoteCancelAck563NoQuoteEntries) {
	return
}

func (q *QuoteCancelAck563NoQuoteSets) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint16(_w, q.QuoteSetID); err != nil {
		return err
	}
	if err := _m.WriteUint16(_w, q.QuoteErrorCode); err != nil {
		return err
	}
	return nil
}

func (q *QuoteCancelAck563NoQuoteSets) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint) error {
	if !q.QuoteSetIDInActingVersion(actingVersion) {
		q.QuoteSetID = q.QuoteSetIDNullValue()
	} else {
		if err := _m.ReadUint16(_r, &q.QuoteSetID); err != nil {
			return err
		}
	}
	if !q.QuoteErrorCodeInActingVersion(actingVersion) {
		q.QuoteErrorCode = q.QuoteErrorCodeNullValue()
	} else {
		if err := _m.ReadUint16(_r, &q.QuoteErrorCode); err != nil {
			return err
		}
	}
	if actingVersion > q.SbeSchemaVersion() && blockLength > q.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-q.SbeBlockLength()))
	}
	return nil
}

func (q *QuoteCancelAck563NoQuoteSets) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if q.QuoteSetIDInActingVersion(actingVersion) {
		if q.QuoteSetID < q.QuoteSetIDMinValue() || q.QuoteSetID > q.QuoteSetIDMaxValue() {
			return fmt.Errorf("Range check failed on q.QuoteSetID (%v < %v > %v)", q.QuoteSetIDMinValue(), q.QuoteSetID, q.QuoteSetIDMaxValue())
		}
	}
	if q.QuoteErrorCodeInActingVersion(actingVersion) {
		if q.QuoteErrorCode < q.QuoteErrorCodeMinValue() || q.QuoteErrorCode > q.QuoteErrorCodeMaxValue() {
			return fmt.Errorf("Range check failed on q.QuoteErrorCode (%v < %v > %v)", q.QuoteErrorCodeMinValue(), q.QuoteErrorCode, q.QuoteErrorCodeMaxValue())
		}
	}
	return nil
}

func QuoteCancelAck563NoQuoteSetsInit(q *QuoteCancelAck563NoQuoteSets) {
	return
}

func (*QuoteCancelAck563) SbeBlockLength() (blockLength uint16) {
	return 351
}

func (*QuoteCancelAck563) SbeTemplateId() (templateId uint16) {
	return 563
}

func (*QuoteCancelAck563) SbeSchemaId() (schemaId uint16) {
	return 8
}

func (*QuoteCancelAck563) SbeSchemaVersion() (schemaVersion uint16) {
	return 5
}

func (*QuoteCancelAck563) SbeSemanticType() (semanticType []byte) {
	return []byte("b")
}

func (*QuoteCancelAck563) SeqNumId() uint16 {
	return 9726
}

func (*QuoteCancelAck563) SeqNumSinceVersion() uint16 {
	return 0
}

func (q *QuoteCancelAck563) SeqNumInActingVersion(actingVersion uint16) bool {
	return actingVersion >= q.SeqNumSinceVersion()
}

func (*QuoteCancelAck563) SeqNumDeprecated() uint16 {
	return 0
}

func (*QuoteCancelAck563) SeqNumMetaAttribute(meta int) string {
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

func (*QuoteCancelAck563) SeqNumMinValue() uint32 {
	return 0
}

func (*QuoteCancelAck563) SeqNumMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*QuoteCancelAck563) SeqNumNullValue() uint32 {
	return math.MaxUint32
}

func (*QuoteCancelAck563) UUIDId() uint16 {
	return 39001
}

func (*QuoteCancelAck563) UUIDSinceVersion() uint16 {
	return 0
}

func (q *QuoteCancelAck563) UUIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= q.UUIDSinceVersion()
}

func (*QuoteCancelAck563) UUIDDeprecated() uint16 {
	return 0
}

func (*QuoteCancelAck563) UUIDMetaAttribute(meta int) string {
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

func (*QuoteCancelAck563) UUIDMinValue() uint64 {
	return 0
}

func (*QuoteCancelAck563) UUIDMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*QuoteCancelAck563) UUIDNullValue() uint64 {
	return math.MaxUint64
}

func (*QuoteCancelAck563) TextId() uint16 {
	return 58
}

func (*QuoteCancelAck563) TextSinceVersion() uint16 {
	return 0
}

func (q *QuoteCancelAck563) TextInActingVersion(actingVersion uint16) bool {
	return actingVersion >= q.TextSinceVersion()
}

func (*QuoteCancelAck563) TextDeprecated() uint16 {
	return 0
}

func (*QuoteCancelAck563) TextMetaAttribute(meta int) string {
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

func (*QuoteCancelAck563) TextMinValue() byte {
	return byte(32)
}

func (*QuoteCancelAck563) TextMaxValue() byte {
	return byte(126)
}

func (*QuoteCancelAck563) TextNullValue() byte {
	return 0
}

func (q *QuoteCancelAck563) TextCharacterEncoding() string {
	return "US-ASCII"
}

func (*QuoteCancelAck563) SenderIDId() uint16 {
	return 5392
}

func (*QuoteCancelAck563) SenderIDSinceVersion() uint16 {
	return 0
}

func (q *QuoteCancelAck563) SenderIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= q.SenderIDSinceVersion()
}

func (*QuoteCancelAck563) SenderIDDeprecated() uint16 {
	return 0
}

func (*QuoteCancelAck563) SenderIDMetaAttribute(meta int) string {
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

func (*QuoteCancelAck563) SenderIDMinValue() byte {
	return byte(32)
}

func (*QuoteCancelAck563) SenderIDMaxValue() byte {
	return byte(126)
}

func (*QuoteCancelAck563) SenderIDNullValue() byte {
	return 0
}

func (q *QuoteCancelAck563) SenderIDCharacterEncoding() string {
	return "US-ASCII"
}

func (*QuoteCancelAck563) PartyDetailsListReqIDId() uint16 {
	return 1505
}

func (*QuoteCancelAck563) PartyDetailsListReqIDSinceVersion() uint16 {
	return 0
}

func (q *QuoteCancelAck563) PartyDetailsListReqIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= q.PartyDetailsListReqIDSinceVersion()
}

func (*QuoteCancelAck563) PartyDetailsListReqIDDeprecated() uint16 {
	return 0
}

func (*QuoteCancelAck563) PartyDetailsListReqIDMetaAttribute(meta int) string {
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

func (*QuoteCancelAck563) PartyDetailsListReqIDMinValue() uint64 {
	return 0
}

func (*QuoteCancelAck563) PartyDetailsListReqIDMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*QuoteCancelAck563) PartyDetailsListReqIDNullValue() uint64 {
	return math.MaxUint64
}

func (*QuoteCancelAck563) RequestTimeId() uint16 {
	return 5979
}

func (*QuoteCancelAck563) RequestTimeSinceVersion() uint16 {
	return 0
}

func (q *QuoteCancelAck563) RequestTimeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= q.RequestTimeSinceVersion()
}

func (*QuoteCancelAck563) RequestTimeDeprecated() uint16 {
	return 0
}

func (*QuoteCancelAck563) RequestTimeMetaAttribute(meta int) string {
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

func (*QuoteCancelAck563) RequestTimeMinValue() uint64 {
	return 0
}

func (*QuoteCancelAck563) RequestTimeMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*QuoteCancelAck563) RequestTimeNullValue() uint64 {
	return math.MaxUint64
}

func (*QuoteCancelAck563) SendingTimeEpochId() uint16 {
	return 5297
}

func (*QuoteCancelAck563) SendingTimeEpochSinceVersion() uint16 {
	return 0
}

func (q *QuoteCancelAck563) SendingTimeEpochInActingVersion(actingVersion uint16) bool {
	return actingVersion >= q.SendingTimeEpochSinceVersion()
}

func (*QuoteCancelAck563) SendingTimeEpochDeprecated() uint16 {
	return 0
}

func (*QuoteCancelAck563) SendingTimeEpochMetaAttribute(meta int) string {
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

func (*QuoteCancelAck563) SendingTimeEpochMinValue() uint64 {
	return 0
}

func (*QuoteCancelAck563) SendingTimeEpochMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*QuoteCancelAck563) SendingTimeEpochNullValue() uint64 {
	return math.MaxUint64
}

func (*QuoteCancelAck563) CancelledSymbolId() uint16 {
	return 9774
}

func (*QuoteCancelAck563) CancelledSymbolSinceVersion() uint16 {
	return 0
}

func (q *QuoteCancelAck563) CancelledSymbolInActingVersion(actingVersion uint16) bool {
	return actingVersion >= q.CancelledSymbolSinceVersion()
}

func (*QuoteCancelAck563) CancelledSymbolDeprecated() uint16 {
	return 0
}

func (*QuoteCancelAck563) CancelledSymbolMetaAttribute(meta int) string {
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

func (*QuoteCancelAck563) CancelledSymbolMinValue() byte {
	return byte(32)
}

func (*QuoteCancelAck563) CancelledSymbolMaxValue() byte {
	return byte(126)
}

func (*QuoteCancelAck563) CancelledSymbolNullValue() byte {
	return 0
}

func (q *QuoteCancelAck563) CancelledSymbolCharacterEncoding() string {
	return "US-ASCII"
}

func (*QuoteCancelAck563) LocationId() uint16 {
	return 9537
}

func (*QuoteCancelAck563) LocationSinceVersion() uint16 {
	return 0
}

func (q *QuoteCancelAck563) LocationInActingVersion(actingVersion uint16) bool {
	return actingVersion >= q.LocationSinceVersion()
}

func (*QuoteCancelAck563) LocationDeprecated() uint16 {
	return 0
}

func (*QuoteCancelAck563) LocationMetaAttribute(meta int) string {
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

func (*QuoteCancelAck563) LocationMinValue() byte {
	return byte(32)
}

func (*QuoteCancelAck563) LocationMaxValue() byte {
	return byte(126)
}

func (*QuoteCancelAck563) LocationNullValue() byte {
	return 0
}

func (q *QuoteCancelAck563) LocationCharacterEncoding() string {
	return "US-ASCII"
}

func (*QuoteCancelAck563) QuoteIDId() uint16 {
	return 117
}

func (*QuoteCancelAck563) QuoteIDSinceVersion() uint16 {
	return 0
}

func (q *QuoteCancelAck563) QuoteIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= q.QuoteIDSinceVersion()
}

func (*QuoteCancelAck563) QuoteIDDeprecated() uint16 {
	return 0
}

func (*QuoteCancelAck563) QuoteIDMetaAttribute(meta int) string {
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

func (*QuoteCancelAck563) QuoteIDMinValue() uint32 {
	return 0
}

func (*QuoteCancelAck563) QuoteIDMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*QuoteCancelAck563) QuoteIDNullValue() uint32 {
	return math.MaxUint32
}

func (*QuoteCancelAck563) QuoteRejectReasonId() uint16 {
	return 300
}

func (*QuoteCancelAck563) QuoteRejectReasonSinceVersion() uint16 {
	return 0
}

func (q *QuoteCancelAck563) QuoteRejectReasonInActingVersion(actingVersion uint16) bool {
	return actingVersion >= q.QuoteRejectReasonSinceVersion()
}

func (*QuoteCancelAck563) QuoteRejectReasonDeprecated() uint16 {
	return 0
}

func (*QuoteCancelAck563) QuoteRejectReasonMetaAttribute(meta int) string {
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

func (*QuoteCancelAck563) QuoteRejectReasonMinValue() uint16 {
	return 0
}

func (*QuoteCancelAck563) QuoteRejectReasonMaxValue() uint16 {
	return math.MaxUint16 - 1
}

func (*QuoteCancelAck563) QuoteRejectReasonNullValue() uint16 {
	return 65535
}

func (*QuoteCancelAck563) DelayDurationId() uint16 {
	return 5904
}

func (*QuoteCancelAck563) DelayDurationSinceVersion() uint16 {
	return 0
}

func (q *QuoteCancelAck563) DelayDurationInActingVersion(actingVersion uint16) bool {
	return actingVersion >= q.DelayDurationSinceVersion()
}

func (*QuoteCancelAck563) DelayDurationDeprecated() uint16 {
	return 0
}

func (*QuoteCancelAck563) DelayDurationMetaAttribute(meta int) string {
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

func (*QuoteCancelAck563) DelayDurationMinValue() uint16 {
	return 0
}

func (*QuoteCancelAck563) DelayDurationMaxValue() uint16 {
	return math.MaxUint16 - 1
}

func (*QuoteCancelAck563) DelayDurationNullValue() uint16 {
	return 65535
}

func (*QuoteCancelAck563) ManualOrderIndicatorId() uint16 {
	return 1028
}

func (*QuoteCancelAck563) ManualOrderIndicatorSinceVersion() uint16 {
	return 0
}

func (q *QuoteCancelAck563) ManualOrderIndicatorInActingVersion(actingVersion uint16) bool {
	return actingVersion >= q.ManualOrderIndicatorSinceVersion()
}

func (*QuoteCancelAck563) ManualOrderIndicatorDeprecated() uint16 {
	return 0
}

func (*QuoteCancelAck563) ManualOrderIndicatorMetaAttribute(meta int) string {
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

func (*QuoteCancelAck563) QuoteStatusId() uint16 {
	return 297
}

func (*QuoteCancelAck563) QuoteStatusSinceVersion() uint16 {
	return 0
}

func (q *QuoteCancelAck563) QuoteStatusInActingVersion(actingVersion uint16) bool {
	return actingVersion >= q.QuoteStatusSinceVersion()
}

func (*QuoteCancelAck563) QuoteStatusDeprecated() uint16 {
	return 0
}

func (*QuoteCancelAck563) QuoteStatusMetaAttribute(meta int) string {
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

func (*QuoteCancelAck563) NoProcessedEntriesId() uint16 {
	return 9772
}

func (*QuoteCancelAck563) NoProcessedEntriesSinceVersion() uint16 {
	return 0
}

func (q *QuoteCancelAck563) NoProcessedEntriesInActingVersion(actingVersion uint16) bool {
	return actingVersion >= q.NoProcessedEntriesSinceVersion()
}

func (*QuoteCancelAck563) NoProcessedEntriesDeprecated() uint16 {
	return 0
}

func (*QuoteCancelAck563) NoProcessedEntriesMetaAttribute(meta int) string {
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

func (*QuoteCancelAck563) NoProcessedEntriesMinValue() uint32 {
	return 0
}

func (*QuoteCancelAck563) NoProcessedEntriesMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*QuoteCancelAck563) NoProcessedEntriesNullValue() uint32 {
	return math.MaxUint32
}

func (*QuoteCancelAck563) MMProtectionResetId() uint16 {
	return 9773
}

func (*QuoteCancelAck563) MMProtectionResetSinceVersion() uint16 {
	return 0
}

func (q *QuoteCancelAck563) MMProtectionResetInActingVersion(actingVersion uint16) bool {
	return actingVersion >= q.MMProtectionResetSinceVersion()
}

func (*QuoteCancelAck563) MMProtectionResetDeprecated() uint16 {
	return 0
}

func (*QuoteCancelAck563) MMProtectionResetMetaAttribute(meta int) string {
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

func (*QuoteCancelAck563) UnsolicitedCancelTypeId() uint16 {
	return 9775
}

func (*QuoteCancelAck563) UnsolicitedCancelTypeSinceVersion() uint16 {
	return 0
}

func (q *QuoteCancelAck563) UnsolicitedCancelTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= q.UnsolicitedCancelTypeSinceVersion()
}

func (*QuoteCancelAck563) UnsolicitedCancelTypeDeprecated() uint16 {
	return 0
}

func (*QuoteCancelAck563) UnsolicitedCancelTypeMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return "char"
	case 4:
		return "optional"
	}
	return ""
}

func (*QuoteCancelAck563) UnsolicitedCancelTypeMinValue() byte {
	return byte(32)
}

func (*QuoteCancelAck563) UnsolicitedCancelTypeMaxValue() byte {
	return byte(126)
}

func (*QuoteCancelAck563) UnsolicitedCancelTypeNullValue() byte {
	return 48
}

func (*QuoteCancelAck563) SplitMsgId() uint16 {
	return 9553
}

func (*QuoteCancelAck563) SplitMsgSinceVersion() uint16 {
	return 0
}

func (q *QuoteCancelAck563) SplitMsgInActingVersion(actingVersion uint16) bool {
	return actingVersion >= q.SplitMsgSinceVersion()
}

func (*QuoteCancelAck563) SplitMsgDeprecated() uint16 {
	return 0
}

func (*QuoteCancelAck563) SplitMsgMetaAttribute(meta int) string {
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

func (*QuoteCancelAck563) TotNoQuoteEntriesId() uint16 {
	return 304
}

func (*QuoteCancelAck563) TotNoQuoteEntriesSinceVersion() uint16 {
	return 0
}

func (q *QuoteCancelAck563) TotNoQuoteEntriesInActingVersion(actingVersion uint16) bool {
	return actingVersion >= q.TotNoQuoteEntriesSinceVersion()
}

func (*QuoteCancelAck563) TotNoQuoteEntriesDeprecated() uint16 {
	return 0
}

func (*QuoteCancelAck563) TotNoQuoteEntriesMetaAttribute(meta int) string {
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

func (*QuoteCancelAck563) TotNoQuoteEntriesMinValue() uint8 {
	return 0
}

func (*QuoteCancelAck563) TotNoQuoteEntriesMaxValue() uint8 {
	return math.MaxUint8 - 1
}

func (*QuoteCancelAck563) TotNoQuoteEntriesNullValue() uint8 {
	return 255
}

func (*QuoteCancelAck563) LiquidityFlagId() uint16 {
	return 9373
}

func (*QuoteCancelAck563) LiquidityFlagSinceVersion() uint16 {
	return 0
}

func (q *QuoteCancelAck563) LiquidityFlagInActingVersion(actingVersion uint16) bool {
	return actingVersion >= q.LiquidityFlagSinceVersion()
}

func (*QuoteCancelAck563) LiquidityFlagDeprecated() uint16 {
	return 0
}

func (*QuoteCancelAck563) LiquidityFlagMetaAttribute(meta int) string {
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

func (*QuoteCancelAck563) PossRetransFlagId() uint16 {
	return 9765
}

func (*QuoteCancelAck563) PossRetransFlagSinceVersion() uint16 {
	return 0
}

func (q *QuoteCancelAck563) PossRetransFlagInActingVersion(actingVersion uint16) bool {
	return actingVersion >= q.PossRetransFlagSinceVersion()
}

func (*QuoteCancelAck563) PossRetransFlagDeprecated() uint16 {
	return 0
}

func (*QuoteCancelAck563) PossRetransFlagMetaAttribute(meta int) string {
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

func (*QuoteCancelAck563) DelayToTimeId() uint16 {
	return 7552
}

func (*QuoteCancelAck563) DelayToTimeSinceVersion() uint16 {
	return 0
}

func (q *QuoteCancelAck563) DelayToTimeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= q.DelayToTimeSinceVersion()
}

func (*QuoteCancelAck563) DelayToTimeDeprecated() uint16 {
	return 0
}

func (*QuoteCancelAck563) DelayToTimeMetaAttribute(meta int) string {
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

func (*QuoteCancelAck563) DelayToTimeMinValue() uint64 {
	return 0
}

func (*QuoteCancelAck563) DelayToTimeMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*QuoteCancelAck563) DelayToTimeNullValue() uint64 {
	return 18446744073709551615
}

func (*QuoteCancelAck563NoQuoteEntries) QuoteEntryIDId() uint16 {
	return 299
}

func (*QuoteCancelAck563NoQuoteEntries) QuoteEntryIDSinceVersion() uint16 {
	return 0
}

func (q *QuoteCancelAck563NoQuoteEntries) QuoteEntryIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= q.QuoteEntryIDSinceVersion()
}

func (*QuoteCancelAck563NoQuoteEntries) QuoteEntryIDDeprecated() uint16 {
	return 0
}

func (*QuoteCancelAck563NoQuoteEntries) QuoteEntryIDMetaAttribute(meta int) string {
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

func (*QuoteCancelAck563NoQuoteEntries) QuoteEntryIDMinValue() uint32 {
	return 0
}

func (*QuoteCancelAck563NoQuoteEntries) QuoteEntryIDMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*QuoteCancelAck563NoQuoteEntries) QuoteEntryIDNullValue() uint32 {
	return math.MaxUint32
}

func (*QuoteCancelAck563NoQuoteEntries) SecurityIDId() uint16 {
	return 48
}

func (*QuoteCancelAck563NoQuoteEntries) SecurityIDSinceVersion() uint16 {
	return 0
}

func (q *QuoteCancelAck563NoQuoteEntries) SecurityIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= q.SecurityIDSinceVersion()
}

func (*QuoteCancelAck563NoQuoteEntries) SecurityIDDeprecated() uint16 {
	return 0
}

func (*QuoteCancelAck563NoQuoteEntries) SecurityIDMetaAttribute(meta int) string {
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

func (*QuoteCancelAck563NoQuoteEntries) SecurityIDMinValue() int32 {
	return math.MinInt32 + 1
}

func (*QuoteCancelAck563NoQuoteEntries) SecurityIDMaxValue() int32 {
	return math.MaxInt32
}

func (*QuoteCancelAck563NoQuoteEntries) SecurityIDNullValue() int32 {
	return math.MinInt32
}

func (*QuoteCancelAck563NoQuoteEntries) QuoteEntryRejectReasonId() uint16 {
	return 368
}

func (*QuoteCancelAck563NoQuoteEntries) QuoteEntryRejectReasonSinceVersion() uint16 {
	return 0
}

func (q *QuoteCancelAck563NoQuoteEntries) QuoteEntryRejectReasonInActingVersion(actingVersion uint16) bool {
	return actingVersion >= q.QuoteEntryRejectReasonSinceVersion()
}

func (*QuoteCancelAck563NoQuoteEntries) QuoteEntryRejectReasonDeprecated() uint16 {
	return 0
}

func (*QuoteCancelAck563NoQuoteEntries) QuoteEntryRejectReasonMetaAttribute(meta int) string {
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

func (*QuoteCancelAck563NoQuoteEntries) QuoteEntryRejectReasonMinValue() uint8 {
	return 0
}

func (*QuoteCancelAck563NoQuoteEntries) QuoteEntryRejectReasonMaxValue() uint8 {
	return math.MaxUint8 - 1
}

func (*QuoteCancelAck563NoQuoteEntries) QuoteEntryRejectReasonNullValue() uint8 {
	return math.MaxUint8
}

func (*QuoteCancelAck563NoQuoteSets) QuoteSetIDId() uint16 {
	return 302
}

func (*QuoteCancelAck563NoQuoteSets) QuoteSetIDSinceVersion() uint16 {
	return 0
}

func (q *QuoteCancelAck563NoQuoteSets) QuoteSetIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= q.QuoteSetIDSinceVersion()
}

func (*QuoteCancelAck563NoQuoteSets) QuoteSetIDDeprecated() uint16 {
	return 0
}

func (*QuoteCancelAck563NoQuoteSets) QuoteSetIDMetaAttribute(meta int) string {
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

func (*QuoteCancelAck563NoQuoteSets) QuoteSetIDMinValue() uint16 {
	return 0
}

func (*QuoteCancelAck563NoQuoteSets) QuoteSetIDMaxValue() uint16 {
	return math.MaxUint16 - 1
}

func (*QuoteCancelAck563NoQuoteSets) QuoteSetIDNullValue() uint16 {
	return math.MaxUint16
}

func (*QuoteCancelAck563NoQuoteSets) QuoteErrorCodeId() uint16 {
	return 9030
}

func (*QuoteCancelAck563NoQuoteSets) QuoteErrorCodeSinceVersion() uint16 {
	return 0
}

func (q *QuoteCancelAck563NoQuoteSets) QuoteErrorCodeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= q.QuoteErrorCodeSinceVersion()
}

func (*QuoteCancelAck563NoQuoteSets) QuoteErrorCodeDeprecated() uint16 {
	return 0
}

func (*QuoteCancelAck563NoQuoteSets) QuoteErrorCodeMetaAttribute(meta int) string {
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

func (*QuoteCancelAck563NoQuoteSets) QuoteErrorCodeMinValue() uint16 {
	return 0
}

func (*QuoteCancelAck563NoQuoteSets) QuoteErrorCodeMaxValue() uint16 {
	return math.MaxUint16 - 1
}

func (*QuoteCancelAck563NoQuoteSets) QuoteErrorCodeNullValue() uint16 {
	return math.MaxUint16
}

func (*QuoteCancelAck563) NoQuoteEntriesId() uint16 {
	return 295
}

func (*QuoteCancelAck563) NoQuoteEntriesSinceVersion() uint16 {
	return 0
}

func (q *QuoteCancelAck563) NoQuoteEntriesInActingVersion(actingVersion uint16) bool {
	return actingVersion >= q.NoQuoteEntriesSinceVersion()
}

func (*QuoteCancelAck563) NoQuoteEntriesDeprecated() uint16 {
	return 0
}

func (*QuoteCancelAck563NoQuoteEntries) SbeBlockLength() (blockLength uint) {
	return 9
}

func (*QuoteCancelAck563NoQuoteEntries) SbeSchemaVersion() (schemaVersion uint16) {
	return 5
}

func (*QuoteCancelAck563) NoQuoteSetsId() uint16 {
	return 296
}

func (*QuoteCancelAck563) NoQuoteSetsSinceVersion() uint16 {
	return 0
}

func (q *QuoteCancelAck563) NoQuoteSetsInActingVersion(actingVersion uint16) bool {
	return actingVersion >= q.NoQuoteSetsSinceVersion()
}

func (*QuoteCancelAck563) NoQuoteSetsDeprecated() uint16 {
	return 0
}

func (*QuoteCancelAck563NoQuoteSets) SbeBlockLength() (blockLength uint) {
	return 4
}

func (*QuoteCancelAck563NoQuoteSets) SbeSchemaVersion() (schemaVersion uint16) {
	return 5
}
