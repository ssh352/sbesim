// Generated SBE (Simple Binary Encoding) message codec

package ilinkbinary

import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
)

type MassQuote517 struct {
	PartyDetailsListReqID uint64
	SendingTimeEpoch      uint64
	ManualOrderIndicator  ManualOrdIndReqEnum
	SeqNum                uint32
	SenderID              [20]byte
	QuoteReqID            uint64
	Location              [5]byte
	QuoteID               uint32
	TotNoQuoteEntries     uint8
	MMProtectionReset     BooleanFlagEnum
	LiquidityFlag         BooleanNULLEnum
	ShortSaleType         ShortSaleTypeEnum
	Reserved              [30]byte
	NoQuoteEntries        []MassQuote517NoQuoteEntries
}
type MassQuote517NoQuoteEntries struct {
	BidPx                PRICENULL9
	OfferPx              PRICENULL9
	QuoteEntryID         uint32
	SecurityID           int32
	BidSize              uint32
	OfferSize            uint32
	UnderlyingSecurityID int32
	QuoteSetID           uint16
}

func (m *MassQuote517) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := m.RangeCheck(m.SbeSchemaVersion(), m.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	if err := _m.WriteUint64(_w, m.PartyDetailsListReqID); err != nil {
		return err
	}
	if err := _m.WriteUint64(_w, m.SendingTimeEpoch); err != nil {
		return err
	}
	if err := m.ManualOrderIndicator.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteUint32(_w, m.SeqNum); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, m.SenderID[:]); err != nil {
		return err
	}
	if err := _m.WriteUint64(_w, m.QuoteReqID); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, m.Location[:]); err != nil {
		return err
	}
	if err := _m.WriteUint32(_w, m.QuoteID); err != nil {
		return err
	}
	if err := _m.WriteUint8(_w, m.TotNoQuoteEntries); err != nil {
		return err
	}
	if err := m.MMProtectionReset.Encode(_m, _w); err != nil {
		return err
	}
	if err := m.LiquidityFlag.Encode(_m, _w); err != nil {
		return err
	}
	if err := m.ShortSaleType.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, m.Reserved[:]); err != nil {
		return err
	}
	var NoQuoteEntriesBlockLength uint16 = 38
	if err := _m.WriteUint16(_w, NoQuoteEntriesBlockLength); err != nil {
		return err
	}
	var NoQuoteEntriesNumInGroup uint8 = uint8(len(m.NoQuoteEntries))
	if err := _m.WriteUint8(_w, NoQuoteEntriesNumInGroup); err != nil {
		return err
	}
	for _, prop := range m.NoQuoteEntries {
		if err := prop.Encode(_m, _w); err != nil {
			return err
		}
	}
	return nil
}

func (m *MassQuote517) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	if !m.PartyDetailsListReqIDInActingVersion(actingVersion) {
		m.PartyDetailsListReqID = m.PartyDetailsListReqIDNullValue()
	} else {
		if err := _m.ReadUint64(_r, &m.PartyDetailsListReqID); err != nil {
			return err
		}
	}
	if !m.SendingTimeEpochInActingVersion(actingVersion) {
		m.SendingTimeEpoch = m.SendingTimeEpochNullValue()
	} else {
		if err := _m.ReadUint64(_r, &m.SendingTimeEpoch); err != nil {
			return err
		}
	}
	if m.ManualOrderIndicatorInActingVersion(actingVersion) {
		if err := m.ManualOrderIndicator.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if !m.SeqNumInActingVersion(actingVersion) {
		m.SeqNum = m.SeqNumNullValue()
	} else {
		if err := _m.ReadUint32(_r, &m.SeqNum); err != nil {
			return err
		}
	}
	if !m.SenderIDInActingVersion(actingVersion) {
		for idx := 0; idx < 20; idx++ {
			m.SenderID[idx] = m.SenderIDNullValue()
		}
	} else {
		if err := _m.ReadBytes(_r, m.SenderID[:]); err != nil {
			return err
		}
	}
	if !m.QuoteReqIDInActingVersion(actingVersion) {
		m.QuoteReqID = m.QuoteReqIDNullValue()
	} else {
		if err := _m.ReadUint64(_r, &m.QuoteReqID); err != nil {
			return err
		}
	}
	if !m.LocationInActingVersion(actingVersion) {
		for idx := 0; idx < 5; idx++ {
			m.Location[idx] = m.LocationNullValue()
		}
	} else {
		if err := _m.ReadBytes(_r, m.Location[:]); err != nil {
			return err
		}
	}
	if !m.QuoteIDInActingVersion(actingVersion) {
		m.QuoteID = m.QuoteIDNullValue()
	} else {
		if err := _m.ReadUint32(_r, &m.QuoteID); err != nil {
			return err
		}
	}
	if !m.TotNoQuoteEntriesInActingVersion(actingVersion) {
		m.TotNoQuoteEntries = m.TotNoQuoteEntriesNullValue()
	} else {
		if err := _m.ReadUint8(_r, &m.TotNoQuoteEntries); err != nil {
			return err
		}
	}
	if m.MMProtectionResetInActingVersion(actingVersion) {
		if err := m.MMProtectionReset.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if m.LiquidityFlagInActingVersion(actingVersion) {
		if err := m.LiquidityFlag.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if m.ShortSaleTypeInActingVersion(actingVersion) {
		if err := m.ShortSaleType.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if !m.ReservedInActingVersion(actingVersion) {
		for idx := 0; idx < 30; idx++ {
			m.Reserved[idx] = m.ReservedNullValue()
		}
	} else {
		if err := _m.ReadBytes(_r, m.Reserved[:]); err != nil {
			return err
		}
	}
	if actingVersion > m.SbeSchemaVersion() && blockLength > m.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-m.SbeBlockLength()))
	}

	if m.NoQuoteEntriesInActingVersion(actingVersion) {
		var NoQuoteEntriesBlockLength uint16
		if err := _m.ReadUint16(_r, &NoQuoteEntriesBlockLength); err != nil {
			return err
		}
		var NoQuoteEntriesNumInGroup uint8
		if err := _m.ReadUint8(_r, &NoQuoteEntriesNumInGroup); err != nil {
			return err
		}
		if cap(m.NoQuoteEntries) < int(NoQuoteEntriesNumInGroup) {
			m.NoQuoteEntries = make([]MassQuote517NoQuoteEntries, NoQuoteEntriesNumInGroup)
		}
		for i, _ := range m.NoQuoteEntries {
			if err := m.NoQuoteEntries[i].Decode(_m, _r, actingVersion, uint(NoQuoteEntriesBlockLength)); err != nil {
				return err
			}
		}
	}
	if doRangeCheck {
		if err := m.RangeCheck(actingVersion, m.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	return nil
}

func (m *MassQuote517) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if m.PartyDetailsListReqIDInActingVersion(actingVersion) {
		if m.PartyDetailsListReqID < m.PartyDetailsListReqIDMinValue() || m.PartyDetailsListReqID > m.PartyDetailsListReqIDMaxValue() {
			return fmt.Errorf("Range check failed on m.PartyDetailsListReqID (%v < %v > %v)", m.PartyDetailsListReqIDMinValue(), m.PartyDetailsListReqID, m.PartyDetailsListReqIDMaxValue())
		}
	}
	if m.SendingTimeEpochInActingVersion(actingVersion) {
		if m.SendingTimeEpoch < m.SendingTimeEpochMinValue() || m.SendingTimeEpoch > m.SendingTimeEpochMaxValue() {
			return fmt.Errorf("Range check failed on m.SendingTimeEpoch (%v < %v > %v)", m.SendingTimeEpochMinValue(), m.SendingTimeEpoch, m.SendingTimeEpochMaxValue())
		}
	}
	if err := m.ManualOrderIndicator.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if m.SeqNumInActingVersion(actingVersion) {
		if m.SeqNum < m.SeqNumMinValue() || m.SeqNum > m.SeqNumMaxValue() {
			return fmt.Errorf("Range check failed on m.SeqNum (%v < %v > %v)", m.SeqNumMinValue(), m.SeqNum, m.SeqNumMaxValue())
		}
	}
	if m.SenderIDInActingVersion(actingVersion) {
		for idx := 0; idx < 20; idx++ {
			if m.SenderID[idx] < m.SenderIDMinValue() || m.SenderID[idx] > m.SenderIDMaxValue() {
				return fmt.Errorf("Range check failed on m.SenderID[%d] (%v < %v > %v)", idx, m.SenderIDMinValue(), m.SenderID[idx], m.SenderIDMaxValue())
			}
		}
	}
	if m.QuoteReqIDInActingVersion(actingVersion) {
		if m.QuoteReqID != m.QuoteReqIDNullValue() && (m.QuoteReqID < m.QuoteReqIDMinValue() || m.QuoteReqID > m.QuoteReqIDMaxValue()) {
			return fmt.Errorf("Range check failed on m.QuoteReqID (%v < %v > %v)", m.QuoteReqIDMinValue(), m.QuoteReqID, m.QuoteReqIDMaxValue())
		}
	}
	if m.LocationInActingVersion(actingVersion) {
		for idx := 0; idx < 5; idx++ {
			if m.Location[idx] < m.LocationMinValue() || m.Location[idx] > m.LocationMaxValue() {
				return fmt.Errorf("Range check failed on m.Location[%d] (%v < %v > %v)", idx, m.LocationMinValue(), m.Location[idx], m.LocationMaxValue())
			}
		}
	}
	if m.QuoteIDInActingVersion(actingVersion) {
		if m.QuoteID < m.QuoteIDMinValue() || m.QuoteID > m.QuoteIDMaxValue() {
			return fmt.Errorf("Range check failed on m.QuoteID (%v < %v > %v)", m.QuoteIDMinValue(), m.QuoteID, m.QuoteIDMaxValue())
		}
	}
	if m.TotNoQuoteEntriesInActingVersion(actingVersion) {
		if m.TotNoQuoteEntries < m.TotNoQuoteEntriesMinValue() || m.TotNoQuoteEntries > m.TotNoQuoteEntriesMaxValue() {
			return fmt.Errorf("Range check failed on m.TotNoQuoteEntries (%v < %v > %v)", m.TotNoQuoteEntriesMinValue(), m.TotNoQuoteEntries, m.TotNoQuoteEntriesMaxValue())
		}
	}
	if err := m.MMProtectionReset.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if err := m.LiquidityFlag.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if err := m.ShortSaleType.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if m.ReservedInActingVersion(actingVersion) {
		for idx := 0; idx < 30; idx++ {
			if m.Reserved[idx] < m.ReservedMinValue() || m.Reserved[idx] > m.ReservedMaxValue() {
				return fmt.Errorf("Range check failed on m.Reserved[%d] (%v < %v > %v)", idx, m.ReservedMinValue(), m.Reserved[idx], m.ReservedMaxValue())
			}
		}
	}
	for _, prop := range m.NoQuoteEntries {
		if err := prop.RangeCheck(actingVersion, schemaVersion); err != nil {
			return err
		}
	}
	return nil
}

func MassQuote517Init(m *MassQuote517) {
	m.QuoteReqID = 18446744073709551615
	for idx := 0; idx < 30; idx++ {
		m.Reserved[idx] = 0
	}
	return
}

func (m *MassQuote517NoQuoteEntries) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := m.BidPx.Encode(_m, _w); err != nil {
		return err
	}
	if err := m.OfferPx.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteUint32(_w, m.QuoteEntryID); err != nil {
		return err
	}
	if err := _m.WriteInt32(_w, m.SecurityID); err != nil {
		return err
	}
	if err := _m.WriteUint32(_w, m.BidSize); err != nil {
		return err
	}
	if err := _m.WriteUint32(_w, m.OfferSize); err != nil {
		return err
	}
	if err := _m.WriteInt32(_w, m.UnderlyingSecurityID); err != nil {
		return err
	}
	if err := _m.WriteUint16(_w, m.QuoteSetID); err != nil {
		return err
	}
	return nil
}

func (m *MassQuote517NoQuoteEntries) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint) error {
	if m.BidPxInActingVersion(actingVersion) {
		if err := m.BidPx.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if m.OfferPxInActingVersion(actingVersion) {
		if err := m.OfferPx.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if !m.QuoteEntryIDInActingVersion(actingVersion) {
		m.QuoteEntryID = m.QuoteEntryIDNullValue()
	} else {
		if err := _m.ReadUint32(_r, &m.QuoteEntryID); err != nil {
			return err
		}
	}
	if !m.SecurityIDInActingVersion(actingVersion) {
		m.SecurityID = m.SecurityIDNullValue()
	} else {
		if err := _m.ReadInt32(_r, &m.SecurityID); err != nil {
			return err
		}
	}
	if !m.BidSizeInActingVersion(actingVersion) {
		m.BidSize = m.BidSizeNullValue()
	} else {
		if err := _m.ReadUint32(_r, &m.BidSize); err != nil {
			return err
		}
	}
	if !m.OfferSizeInActingVersion(actingVersion) {
		m.OfferSize = m.OfferSizeNullValue()
	} else {
		if err := _m.ReadUint32(_r, &m.OfferSize); err != nil {
			return err
		}
	}
	if !m.UnderlyingSecurityIDInActingVersion(actingVersion) {
		m.UnderlyingSecurityID = m.UnderlyingSecurityIDNullValue()
	} else {
		if err := _m.ReadInt32(_r, &m.UnderlyingSecurityID); err != nil {
			return err
		}
	}
	if !m.QuoteSetIDInActingVersion(actingVersion) {
		m.QuoteSetID = m.QuoteSetIDNullValue()
	} else {
		if err := _m.ReadUint16(_r, &m.QuoteSetID); err != nil {
			return err
		}
	}
	if actingVersion > m.SbeSchemaVersion() && blockLength > m.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-m.SbeBlockLength()))
	}
	return nil
}

func (m *MassQuote517NoQuoteEntries) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if m.QuoteEntryIDInActingVersion(actingVersion) {
		if m.QuoteEntryID < m.QuoteEntryIDMinValue() || m.QuoteEntryID > m.QuoteEntryIDMaxValue() {
			return fmt.Errorf("Range check failed on m.QuoteEntryID (%v < %v > %v)", m.QuoteEntryIDMinValue(), m.QuoteEntryID, m.QuoteEntryIDMaxValue())
		}
	}
	if m.SecurityIDInActingVersion(actingVersion) {
		if m.SecurityID < m.SecurityIDMinValue() || m.SecurityID > m.SecurityIDMaxValue() {
			return fmt.Errorf("Range check failed on m.SecurityID (%v < %v > %v)", m.SecurityIDMinValue(), m.SecurityID, m.SecurityIDMaxValue())
		}
	}
	if m.BidSizeInActingVersion(actingVersion) {
		if m.BidSize != m.BidSizeNullValue() && (m.BidSize < m.BidSizeMinValue() || m.BidSize > m.BidSizeMaxValue()) {
			return fmt.Errorf("Range check failed on m.BidSize (%v < %v > %v)", m.BidSizeMinValue(), m.BidSize, m.BidSizeMaxValue())
		}
	}
	if m.OfferSizeInActingVersion(actingVersion) {
		if m.OfferSize != m.OfferSizeNullValue() && (m.OfferSize < m.OfferSizeMinValue() || m.OfferSize > m.OfferSizeMaxValue()) {
			return fmt.Errorf("Range check failed on m.OfferSize (%v < %v > %v)", m.OfferSizeMinValue(), m.OfferSize, m.OfferSizeMaxValue())
		}
	}
	if m.UnderlyingSecurityIDInActingVersion(actingVersion) {
		if m.UnderlyingSecurityID != m.UnderlyingSecurityIDNullValue() && (m.UnderlyingSecurityID < m.UnderlyingSecurityIDMinValue() || m.UnderlyingSecurityID > m.UnderlyingSecurityIDMaxValue()) {
			return fmt.Errorf("Range check failed on m.UnderlyingSecurityID (%v < %v > %v)", m.UnderlyingSecurityIDMinValue(), m.UnderlyingSecurityID, m.UnderlyingSecurityIDMaxValue())
		}
	}
	if m.QuoteSetIDInActingVersion(actingVersion) {
		if m.QuoteSetID < m.QuoteSetIDMinValue() || m.QuoteSetID > m.QuoteSetIDMaxValue() {
			return fmt.Errorf("Range check failed on m.QuoteSetID (%v < %v > %v)", m.QuoteSetIDMinValue(), m.QuoteSetID, m.QuoteSetIDMaxValue())
		}
	}
	return nil
}

func MassQuote517NoQuoteEntriesInit(m *MassQuote517NoQuoteEntries) {
	m.BidSize = 4294967295
	m.OfferSize = 4294967295
	m.UnderlyingSecurityID = 2147483647
	return
}

func (*MassQuote517) SbeBlockLength() (blockLength uint16) {
	return 92
}

func (*MassQuote517) SbeTemplateId() (templateId uint16) {
	return 517
}

func (*MassQuote517) SbeSchemaId() (schemaId uint16) {
	return 8
}

func (*MassQuote517) SbeSchemaVersion() (schemaVersion uint16) {
	return 5
}

func (*MassQuote517) SbeSemanticType() (semanticType []byte) {
	return []byte("i")
}

func (*MassQuote517) PartyDetailsListReqIDId() uint16 {
	return 1505
}

func (*MassQuote517) PartyDetailsListReqIDSinceVersion() uint16 {
	return 0
}

func (m *MassQuote517) PartyDetailsListReqIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.PartyDetailsListReqIDSinceVersion()
}

func (*MassQuote517) PartyDetailsListReqIDDeprecated() uint16 {
	return 0
}

func (*MassQuote517) PartyDetailsListReqIDMetaAttribute(meta int) string {
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

func (*MassQuote517) PartyDetailsListReqIDMinValue() uint64 {
	return 0
}

func (*MassQuote517) PartyDetailsListReqIDMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*MassQuote517) PartyDetailsListReqIDNullValue() uint64 {
	return math.MaxUint64
}

func (*MassQuote517) SendingTimeEpochId() uint16 {
	return 5297
}

func (*MassQuote517) SendingTimeEpochSinceVersion() uint16 {
	return 0
}

func (m *MassQuote517) SendingTimeEpochInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.SendingTimeEpochSinceVersion()
}

func (*MassQuote517) SendingTimeEpochDeprecated() uint16 {
	return 0
}

func (*MassQuote517) SendingTimeEpochMetaAttribute(meta int) string {
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

func (*MassQuote517) SendingTimeEpochMinValue() uint64 {
	return 0
}

func (*MassQuote517) SendingTimeEpochMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*MassQuote517) SendingTimeEpochNullValue() uint64 {
	return math.MaxUint64
}

func (*MassQuote517) ManualOrderIndicatorId() uint16 {
	return 1028
}

func (*MassQuote517) ManualOrderIndicatorSinceVersion() uint16 {
	return 0
}

func (m *MassQuote517) ManualOrderIndicatorInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.ManualOrderIndicatorSinceVersion()
}

func (*MassQuote517) ManualOrderIndicatorDeprecated() uint16 {
	return 0
}

func (*MassQuote517) ManualOrderIndicatorMetaAttribute(meta int) string {
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

func (*MassQuote517) SeqNumId() uint16 {
	return 9726
}

func (*MassQuote517) SeqNumSinceVersion() uint16 {
	return 0
}

func (m *MassQuote517) SeqNumInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.SeqNumSinceVersion()
}

func (*MassQuote517) SeqNumDeprecated() uint16 {
	return 0
}

func (*MassQuote517) SeqNumMetaAttribute(meta int) string {
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

func (*MassQuote517) SeqNumMinValue() uint32 {
	return 0
}

func (*MassQuote517) SeqNumMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*MassQuote517) SeqNumNullValue() uint32 {
	return math.MaxUint32
}

func (*MassQuote517) SenderIDId() uint16 {
	return 5392
}

func (*MassQuote517) SenderIDSinceVersion() uint16 {
	return 0
}

func (m *MassQuote517) SenderIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.SenderIDSinceVersion()
}

func (*MassQuote517) SenderIDDeprecated() uint16 {
	return 0
}

func (*MassQuote517) SenderIDMetaAttribute(meta int) string {
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

func (*MassQuote517) SenderIDMinValue() byte {
	return byte(32)
}

func (*MassQuote517) SenderIDMaxValue() byte {
	return byte(126)
}

func (*MassQuote517) SenderIDNullValue() byte {
	return 0
}

func (m *MassQuote517) SenderIDCharacterEncoding() string {
	return "US-ASCII"
}

func (*MassQuote517) QuoteReqIDId() uint16 {
	return 131
}

func (*MassQuote517) QuoteReqIDSinceVersion() uint16 {
	return 0
}

func (m *MassQuote517) QuoteReqIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.QuoteReqIDSinceVersion()
}

func (*MassQuote517) QuoteReqIDDeprecated() uint16 {
	return 0
}

func (*MassQuote517) QuoteReqIDMetaAttribute(meta int) string {
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

func (*MassQuote517) QuoteReqIDMinValue() uint64 {
	return 0
}

func (*MassQuote517) QuoteReqIDMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*MassQuote517) QuoteReqIDNullValue() uint64 {
	return 18446744073709551615
}

func (*MassQuote517) LocationId() uint16 {
	return 9537
}

func (*MassQuote517) LocationSinceVersion() uint16 {
	return 0
}

func (m *MassQuote517) LocationInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.LocationSinceVersion()
}

func (*MassQuote517) LocationDeprecated() uint16 {
	return 0
}

func (*MassQuote517) LocationMetaAttribute(meta int) string {
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

func (*MassQuote517) LocationMinValue() byte {
	return byte(32)
}

func (*MassQuote517) LocationMaxValue() byte {
	return byte(126)
}

func (*MassQuote517) LocationNullValue() byte {
	return 0
}

func (m *MassQuote517) LocationCharacterEncoding() string {
	return "US-ASCII"
}

func (*MassQuote517) QuoteIDId() uint16 {
	return 117
}

func (*MassQuote517) QuoteIDSinceVersion() uint16 {
	return 0
}

func (m *MassQuote517) QuoteIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.QuoteIDSinceVersion()
}

func (*MassQuote517) QuoteIDDeprecated() uint16 {
	return 0
}

func (*MassQuote517) QuoteIDMetaAttribute(meta int) string {
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

func (*MassQuote517) QuoteIDMinValue() uint32 {
	return 0
}

func (*MassQuote517) QuoteIDMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*MassQuote517) QuoteIDNullValue() uint32 {
	return math.MaxUint32
}

func (*MassQuote517) TotNoQuoteEntriesId() uint16 {
	return 304
}

func (*MassQuote517) TotNoQuoteEntriesSinceVersion() uint16 {
	return 0
}

func (m *MassQuote517) TotNoQuoteEntriesInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.TotNoQuoteEntriesSinceVersion()
}

func (*MassQuote517) TotNoQuoteEntriesDeprecated() uint16 {
	return 0
}

func (*MassQuote517) TotNoQuoteEntriesMetaAttribute(meta int) string {
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

func (*MassQuote517) TotNoQuoteEntriesMinValue() uint8 {
	return 0
}

func (*MassQuote517) TotNoQuoteEntriesMaxValue() uint8 {
	return math.MaxUint8 - 1
}

func (*MassQuote517) TotNoQuoteEntriesNullValue() uint8 {
	return math.MaxUint8
}

func (*MassQuote517) MMProtectionResetId() uint16 {
	return 9773
}

func (*MassQuote517) MMProtectionResetSinceVersion() uint16 {
	return 0
}

func (m *MassQuote517) MMProtectionResetInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.MMProtectionResetSinceVersion()
}

func (*MassQuote517) MMProtectionResetDeprecated() uint16 {
	return 0
}

func (*MassQuote517) MMProtectionResetMetaAttribute(meta int) string {
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

func (*MassQuote517) LiquidityFlagId() uint16 {
	return 9373
}

func (*MassQuote517) LiquidityFlagSinceVersion() uint16 {
	return 0
}

func (m *MassQuote517) LiquidityFlagInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.LiquidityFlagSinceVersion()
}

func (*MassQuote517) LiquidityFlagDeprecated() uint16 {
	return 0
}

func (*MassQuote517) LiquidityFlagMetaAttribute(meta int) string {
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

func (*MassQuote517) ShortSaleTypeId() uint16 {
	return 5409
}

func (*MassQuote517) ShortSaleTypeSinceVersion() uint16 {
	return 0
}

func (m *MassQuote517) ShortSaleTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.ShortSaleTypeSinceVersion()
}

func (*MassQuote517) ShortSaleTypeDeprecated() uint16 {
	return 0
}

func (*MassQuote517) ShortSaleTypeMetaAttribute(meta int) string {
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

func (*MassQuote517) ReservedId() uint16 {
	return 5187
}

func (*MassQuote517) ReservedSinceVersion() uint16 {
	return 5
}

func (m *MassQuote517) ReservedInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.ReservedSinceVersion()
}

func (*MassQuote517) ReservedDeprecated() uint16 {
	return 0
}

func (*MassQuote517) ReservedMetaAttribute(meta int) string {
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

func (*MassQuote517) ReservedMinValue() byte {
	return byte(32)
}

func (*MassQuote517) ReservedMaxValue() byte {
	return byte(126)
}

func (*MassQuote517) ReservedNullValue() byte {
	return 0
}

func (m *MassQuote517) ReservedCharacterEncoding() string {
	return "US-ASCII"
}

func (*MassQuote517NoQuoteEntries) BidPxId() uint16 {
	return 132
}

func (*MassQuote517NoQuoteEntries) BidPxSinceVersion() uint16 {
	return 0
}

func (m *MassQuote517NoQuoteEntries) BidPxInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.BidPxSinceVersion()
}

func (*MassQuote517NoQuoteEntries) BidPxDeprecated() uint16 {
	return 0
}

func (*MassQuote517NoQuoteEntries) BidPxMetaAttribute(meta int) string {
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

func (*MassQuote517NoQuoteEntries) OfferPxId() uint16 {
	return 133
}

func (*MassQuote517NoQuoteEntries) OfferPxSinceVersion() uint16 {
	return 0
}

func (m *MassQuote517NoQuoteEntries) OfferPxInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.OfferPxSinceVersion()
}

func (*MassQuote517NoQuoteEntries) OfferPxDeprecated() uint16 {
	return 0
}

func (*MassQuote517NoQuoteEntries) OfferPxMetaAttribute(meta int) string {
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

func (*MassQuote517NoQuoteEntries) QuoteEntryIDId() uint16 {
	return 299
}

func (*MassQuote517NoQuoteEntries) QuoteEntryIDSinceVersion() uint16 {
	return 0
}

func (m *MassQuote517NoQuoteEntries) QuoteEntryIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.QuoteEntryIDSinceVersion()
}

func (*MassQuote517NoQuoteEntries) QuoteEntryIDDeprecated() uint16 {
	return 0
}

func (*MassQuote517NoQuoteEntries) QuoteEntryIDMetaAttribute(meta int) string {
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

func (*MassQuote517NoQuoteEntries) QuoteEntryIDMinValue() uint32 {
	return 0
}

func (*MassQuote517NoQuoteEntries) QuoteEntryIDMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*MassQuote517NoQuoteEntries) QuoteEntryIDNullValue() uint32 {
	return math.MaxUint32
}

func (*MassQuote517NoQuoteEntries) SecurityIDId() uint16 {
	return 48
}

func (*MassQuote517NoQuoteEntries) SecurityIDSinceVersion() uint16 {
	return 0
}

func (m *MassQuote517NoQuoteEntries) SecurityIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.SecurityIDSinceVersion()
}

func (*MassQuote517NoQuoteEntries) SecurityIDDeprecated() uint16 {
	return 0
}

func (*MassQuote517NoQuoteEntries) SecurityIDMetaAttribute(meta int) string {
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

func (*MassQuote517NoQuoteEntries) SecurityIDMinValue() int32 {
	return math.MinInt32 + 1
}

func (*MassQuote517NoQuoteEntries) SecurityIDMaxValue() int32 {
	return math.MaxInt32
}

func (*MassQuote517NoQuoteEntries) SecurityIDNullValue() int32 {
	return math.MinInt32
}

func (*MassQuote517NoQuoteEntries) BidSizeId() uint16 {
	return 134
}

func (*MassQuote517NoQuoteEntries) BidSizeSinceVersion() uint16 {
	return 0
}

func (m *MassQuote517NoQuoteEntries) BidSizeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.BidSizeSinceVersion()
}

func (*MassQuote517NoQuoteEntries) BidSizeDeprecated() uint16 {
	return 0
}

func (*MassQuote517NoQuoteEntries) BidSizeMetaAttribute(meta int) string {
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

func (*MassQuote517NoQuoteEntries) BidSizeMinValue() uint32 {
	return 0
}

func (*MassQuote517NoQuoteEntries) BidSizeMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*MassQuote517NoQuoteEntries) BidSizeNullValue() uint32 {
	return 4294967295
}

func (*MassQuote517NoQuoteEntries) OfferSizeId() uint16 {
	return 135
}

func (*MassQuote517NoQuoteEntries) OfferSizeSinceVersion() uint16 {
	return 0
}

func (m *MassQuote517NoQuoteEntries) OfferSizeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.OfferSizeSinceVersion()
}

func (*MassQuote517NoQuoteEntries) OfferSizeDeprecated() uint16 {
	return 0
}

func (*MassQuote517NoQuoteEntries) OfferSizeMetaAttribute(meta int) string {
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

func (*MassQuote517NoQuoteEntries) OfferSizeMinValue() uint32 {
	return 0
}

func (*MassQuote517NoQuoteEntries) OfferSizeMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*MassQuote517NoQuoteEntries) OfferSizeNullValue() uint32 {
	return 4294967295
}

func (*MassQuote517NoQuoteEntries) UnderlyingSecurityIDId() uint16 {
	return 309
}

func (*MassQuote517NoQuoteEntries) UnderlyingSecurityIDSinceVersion() uint16 {
	return 0
}

func (m *MassQuote517NoQuoteEntries) UnderlyingSecurityIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.UnderlyingSecurityIDSinceVersion()
}

func (*MassQuote517NoQuoteEntries) UnderlyingSecurityIDDeprecated() uint16 {
	return 0
}

func (*MassQuote517NoQuoteEntries) UnderlyingSecurityIDMetaAttribute(meta int) string {
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

func (*MassQuote517NoQuoteEntries) UnderlyingSecurityIDMinValue() int32 {
	return math.MinInt32 + 1
}

func (*MassQuote517NoQuoteEntries) UnderlyingSecurityIDMaxValue() int32 {
	return math.MaxInt32
}

func (*MassQuote517NoQuoteEntries) UnderlyingSecurityIDNullValue() int32 {
	return 2147483647
}

func (*MassQuote517NoQuoteEntries) QuoteSetIDId() uint16 {
	return 302
}

func (*MassQuote517NoQuoteEntries) QuoteSetIDSinceVersion() uint16 {
	return 0
}

func (m *MassQuote517NoQuoteEntries) QuoteSetIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.QuoteSetIDSinceVersion()
}

func (*MassQuote517NoQuoteEntries) QuoteSetIDDeprecated() uint16 {
	return 0
}

func (*MassQuote517NoQuoteEntries) QuoteSetIDMetaAttribute(meta int) string {
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

func (*MassQuote517NoQuoteEntries) QuoteSetIDMinValue() uint16 {
	return 0
}

func (*MassQuote517NoQuoteEntries) QuoteSetIDMaxValue() uint16 {
	return math.MaxUint16 - 1
}

func (*MassQuote517NoQuoteEntries) QuoteSetIDNullValue() uint16 {
	return math.MaxUint16
}

func (*MassQuote517) NoQuoteEntriesId() uint16 {
	return 295
}

func (*MassQuote517) NoQuoteEntriesSinceVersion() uint16 {
	return 0
}

func (m *MassQuote517) NoQuoteEntriesInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.NoQuoteEntriesSinceVersion()
}

func (*MassQuote517) NoQuoteEntriesDeprecated() uint16 {
	return 0
}

func (*MassQuote517NoQuoteEntries) SbeBlockLength() (blockLength uint) {
	return 38
}

func (*MassQuote517NoQuoteEntries) SbeSchemaVersion() (schemaVersion uint16) {
	return 5
}
