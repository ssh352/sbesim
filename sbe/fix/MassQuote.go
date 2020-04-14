// Generated SBE (Simple Binary Encoding) message codec

package fix

import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
)

type MassQuote struct {
	QuoteReqID            [23]byte
	QuoteID               [10]byte
	MMAccount             [12]byte
	ManualOrderIndicator  BooleanTypeEnum
	CustOrderHandlingInst CustOrderHandlingInstEnum
	CustomerOrFirm        CustomerOrFirmEnum
	SelfMatchPreventionID [12]byte
	CtiCode               CtiCodeEnum
	MMProtectionReset     MMProtectionResetEnum
	QuoteSets             []MassQuoteQuoteSets
}
type MassQuoteQuoteSets struct {
	QuoteSetID             [3]byte
	UnderlyingSecurityDesc [20]byte
	TotQuoteEntries        uint8
	QuoteEntries           []MassQuoteQuoteSetsQuoteEntries
}
type MassQuoteQuoteSetsQuoteEntries struct {
	QuoteEntryID     [10]byte
	Symbol           [6]byte
	SecurityDesc     [20]byte
	SecurityType     [3]byte
	SecurityID       int64
	SecurityIDSource SecurityIDSourceEnum
	TransactTime     uint64
	BidPx            OptionalPrice
	BidSize          int64
	OfferPx          OptionalPrice
	OfferSize        int64
}

func (m *MassQuote) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := m.RangeCheck(m.SbeSchemaVersion(), m.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	if err := _m.WriteBytes(_w, m.QuoteReqID[:]); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, m.QuoteID[:]); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, m.MMAccount[:]); err != nil {
		return err
	}
	if err := m.ManualOrderIndicator.Encode(_m, _w); err != nil {
		return err
	}
	if err := m.CustOrderHandlingInst.Encode(_m, _w); err != nil {
		return err
	}
	if err := m.CustomerOrFirm.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, m.SelfMatchPreventionID[:]); err != nil {
		return err
	}
	if err := m.CtiCode.Encode(_m, _w); err != nil {
		return err
	}
	if err := m.MMProtectionReset.Encode(_m, _w); err != nil {
		return err
	}
	var QuoteSetsBlockLength uint16 = 24
	if err := _m.WriteUint16(_w, QuoteSetsBlockLength); err != nil {
		return err
	}
	var QuoteSetsNumInGroup uint16 = uint16(len(m.QuoteSets))
	if err := _m.WriteUint16(_w, QuoteSetsNumInGroup); err != nil {
		return err
	}
	for _, prop := range m.QuoteSets {
		if err := prop.Encode(_m, _w); err != nil {
			return err
		}
	}
	return nil
}

func (m *MassQuote) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	if !m.QuoteReqIDInActingVersion(actingVersion) {
		for idx := 0; idx < 23; idx++ {
			m.QuoteReqID[idx] = m.QuoteReqIDNullValue()
		}
	} else {
		if err := _m.ReadBytes(_r, m.QuoteReqID[:]); err != nil {
			return err
		}
	}
	if !m.QuoteIDInActingVersion(actingVersion) {
		for idx := 0; idx < 10; idx++ {
			m.QuoteID[idx] = m.QuoteIDNullValue()
		}
	} else {
		if err := _m.ReadBytes(_r, m.QuoteID[:]); err != nil {
			return err
		}
	}
	if !m.MMAccountInActingVersion(actingVersion) {
		for idx := 0; idx < 12; idx++ {
			m.MMAccount[idx] = m.MMAccountNullValue()
		}
	} else {
		if err := _m.ReadBytes(_r, m.MMAccount[:]); err != nil {
			return err
		}
	}
	if m.ManualOrderIndicatorInActingVersion(actingVersion) {
		if err := m.ManualOrderIndicator.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if m.CustOrderHandlingInstInActingVersion(actingVersion) {
		if err := m.CustOrderHandlingInst.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if m.CustomerOrFirmInActingVersion(actingVersion) {
		if err := m.CustomerOrFirm.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if !m.SelfMatchPreventionIDInActingVersion(actingVersion) {
		for idx := 0; idx < 12; idx++ {
			m.SelfMatchPreventionID[idx] = m.SelfMatchPreventionIDNullValue()
		}
	} else {
		if err := _m.ReadBytes(_r, m.SelfMatchPreventionID[:]); err != nil {
			return err
		}
	}
	if m.CtiCodeInActingVersion(actingVersion) {
		if err := m.CtiCode.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if m.MMProtectionResetInActingVersion(actingVersion) {
		if err := m.MMProtectionReset.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if actingVersion > m.SbeSchemaVersion() && blockLength > m.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-m.SbeBlockLength()))
	}

	if m.QuoteSetsInActingVersion(actingVersion) {
		var QuoteSetsBlockLength uint16
		if err := _m.ReadUint16(_r, &QuoteSetsBlockLength); err != nil {
			return err
		}
		var QuoteSetsNumInGroup uint16
		if err := _m.ReadUint16(_r, &QuoteSetsNumInGroup); err != nil {
			return err
		}
		if cap(m.QuoteSets) < int(QuoteSetsNumInGroup) {
			m.QuoteSets = make([]MassQuoteQuoteSets, QuoteSetsNumInGroup)
		}
		for i, _ := range m.QuoteSets {
			if err := m.QuoteSets[i].Decode(_m, _r, actingVersion, uint(QuoteSetsBlockLength)); err != nil {
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

func (m *MassQuote) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if m.QuoteReqIDInActingVersion(actingVersion) {
		for idx := 0; idx < 23; idx++ {
			if m.QuoteReqID[idx] < m.QuoteReqIDMinValue() || m.QuoteReqID[idx] > m.QuoteReqIDMaxValue() {
				return fmt.Errorf("Range check failed on m.QuoteReqID[%d] (%v < %v > %v)", idx, m.QuoteReqIDMinValue(), m.QuoteReqID[idx], m.QuoteReqIDMaxValue())
			}
		}
	}
	if m.QuoteIDInActingVersion(actingVersion) {
		for idx := 0; idx < 10; idx++ {
			if m.QuoteID[idx] < m.QuoteIDMinValue() || m.QuoteID[idx] > m.QuoteIDMaxValue() {
				return fmt.Errorf("Range check failed on m.QuoteID[%d] (%v < %v > %v)", idx, m.QuoteIDMinValue(), m.QuoteID[idx], m.QuoteIDMaxValue())
			}
		}
	}
	if m.MMAccountInActingVersion(actingVersion) {
		for idx := 0; idx < 12; idx++ {
			if m.MMAccount[idx] < m.MMAccountMinValue() || m.MMAccount[idx] > m.MMAccountMaxValue() {
				return fmt.Errorf("Range check failed on m.MMAccount[%d] (%v < %v > %v)", idx, m.MMAccountMinValue(), m.MMAccount[idx], m.MMAccountMaxValue())
			}
		}
	}
	if err := m.ManualOrderIndicator.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if err := m.CustOrderHandlingInst.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if err := m.CustomerOrFirm.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if m.SelfMatchPreventionIDInActingVersion(actingVersion) {
		for idx := 0; idx < 12; idx++ {
			if m.SelfMatchPreventionID[idx] < m.SelfMatchPreventionIDMinValue() || m.SelfMatchPreventionID[idx] > m.SelfMatchPreventionIDMaxValue() {
				return fmt.Errorf("Range check failed on m.SelfMatchPreventionID[%d] (%v < %v > %v)", idx, m.SelfMatchPreventionIDMinValue(), m.SelfMatchPreventionID[idx], m.SelfMatchPreventionIDMaxValue())
			}
		}
	}
	if err := m.CtiCode.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if err := m.MMProtectionReset.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	for _, prop := range m.QuoteSets {
		if err := prop.RangeCheck(actingVersion, schemaVersion); err != nil {
			return err
		}
	}
	return nil
}

func MassQuoteInit(m *MassQuote) {
	return
}

func (m *MassQuoteQuoteSets) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteBytes(_w, m.QuoteSetID[:]); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, m.UnderlyingSecurityDesc[:]); err != nil {
		return err
	}
	if err := _m.WriteUint8(_w, m.TotQuoteEntries); err != nil {
		return err
	}
	var QuoteEntriesBlockLength uint16 = 90
	if err := _m.WriteUint16(_w, QuoteEntriesBlockLength); err != nil {
		return err
	}
	var QuoteEntriesNumInGroup uint16 = uint16(len(m.QuoteEntries))
	if err := _m.WriteUint16(_w, QuoteEntriesNumInGroup); err != nil {
		return err
	}
	for _, prop := range m.QuoteEntries {
		if err := prop.Encode(_m, _w); err != nil {
			return err
		}
	}
	return nil
}

func (m *MassQuoteQuoteSets) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint) error {
	if !m.QuoteSetIDInActingVersion(actingVersion) {
		for idx := 0; idx < 3; idx++ {
			m.QuoteSetID[idx] = m.QuoteSetIDNullValue()
		}
	} else {
		if err := _m.ReadBytes(_r, m.QuoteSetID[:]); err != nil {
			return err
		}
	}
	if !m.UnderlyingSecurityDescInActingVersion(actingVersion) {
		for idx := 0; idx < 20; idx++ {
			m.UnderlyingSecurityDesc[idx] = m.UnderlyingSecurityDescNullValue()
		}
	} else {
		if err := _m.ReadBytes(_r, m.UnderlyingSecurityDesc[:]); err != nil {
			return err
		}
	}
	if !m.TotQuoteEntriesInActingVersion(actingVersion) {
		m.TotQuoteEntries = m.TotQuoteEntriesNullValue()
	} else {
		if err := _m.ReadUint8(_r, &m.TotQuoteEntries); err != nil {
			return err
		}
	}
	if actingVersion > m.SbeSchemaVersion() && blockLength > m.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-m.SbeBlockLength()))
	}

	if m.QuoteEntriesInActingVersion(actingVersion) {
		var QuoteEntriesBlockLength uint16
		if err := _m.ReadUint16(_r, &QuoteEntriesBlockLength); err != nil {
			return err
		}
		var QuoteEntriesNumInGroup uint16
		if err := _m.ReadUint16(_r, &QuoteEntriesNumInGroup); err != nil {
			return err
		}
		if cap(m.QuoteEntries) < int(QuoteEntriesNumInGroup) {
			m.QuoteEntries = make([]MassQuoteQuoteSetsQuoteEntries, QuoteEntriesNumInGroup)
		}
		for i, _ := range m.QuoteEntries {
			if err := m.QuoteEntries[i].Decode(_m, _r, actingVersion, uint(QuoteEntriesBlockLength)); err != nil {
				return err
			}
		}
	}
	return nil
}

func (m *MassQuoteQuoteSets) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if m.QuoteSetIDInActingVersion(actingVersion) {
		for idx := 0; idx < 3; idx++ {
			if m.QuoteSetID[idx] < m.QuoteSetIDMinValue() || m.QuoteSetID[idx] > m.QuoteSetIDMaxValue() {
				return fmt.Errorf("Range check failed on m.QuoteSetID[%d] (%v < %v > %v)", idx, m.QuoteSetIDMinValue(), m.QuoteSetID[idx], m.QuoteSetIDMaxValue())
			}
		}
	}
	if m.UnderlyingSecurityDescInActingVersion(actingVersion) {
		for idx := 0; idx < 20; idx++ {
			if m.UnderlyingSecurityDesc[idx] < m.UnderlyingSecurityDescMinValue() || m.UnderlyingSecurityDesc[idx] > m.UnderlyingSecurityDescMaxValue() {
				return fmt.Errorf("Range check failed on m.UnderlyingSecurityDesc[%d] (%v < %v > %v)", idx, m.UnderlyingSecurityDescMinValue(), m.UnderlyingSecurityDesc[idx], m.UnderlyingSecurityDescMaxValue())
			}
		}
	}
	if m.TotQuoteEntriesInActingVersion(actingVersion) {
		if m.TotQuoteEntries < m.TotQuoteEntriesMinValue() || m.TotQuoteEntries > m.TotQuoteEntriesMaxValue() {
			return fmt.Errorf("Range check failed on m.TotQuoteEntries (%v < %v > %v)", m.TotQuoteEntriesMinValue(), m.TotQuoteEntries, m.TotQuoteEntriesMaxValue())
		}
	}
	for _, prop := range m.QuoteEntries {
		if err := prop.RangeCheck(actingVersion, schemaVersion); err != nil {
			return err
		}
	}
	return nil
}

func MassQuoteQuoteSetsInit(m *MassQuoteQuoteSets) {
	return
}

func (m *MassQuoteQuoteSetsQuoteEntries) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteBytes(_w, m.QuoteEntryID[:]); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, m.Symbol[:]); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, m.SecurityDesc[:]); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, m.SecurityType[:]); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, m.SecurityID); err != nil {
		return err
	}
	if err := m.SecurityIDSource.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteUint64(_w, m.TransactTime); err != nil {
		return err
	}
	if err := m.BidPx.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, m.BidSize); err != nil {
		return err
	}
	if err := m.OfferPx.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, m.OfferSize); err != nil {
		return err
	}
	return nil
}

func (m *MassQuoteQuoteSetsQuoteEntries) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint) error {
	if !m.QuoteEntryIDInActingVersion(actingVersion) {
		for idx := 0; idx < 10; idx++ {
			m.QuoteEntryID[idx] = m.QuoteEntryIDNullValue()
		}
	} else {
		if err := _m.ReadBytes(_r, m.QuoteEntryID[:]); err != nil {
			return err
		}
	}
	if !m.SymbolInActingVersion(actingVersion) {
		for idx := 0; idx < 6; idx++ {
			m.Symbol[idx] = m.SymbolNullValue()
		}
	} else {
		if err := _m.ReadBytes(_r, m.Symbol[:]); err != nil {
			return err
		}
	}
	if !m.SecurityDescInActingVersion(actingVersion) {
		for idx := 0; idx < 20; idx++ {
			m.SecurityDesc[idx] = m.SecurityDescNullValue()
		}
	} else {
		if err := _m.ReadBytes(_r, m.SecurityDesc[:]); err != nil {
			return err
		}
	}
	if !m.SecurityTypeInActingVersion(actingVersion) {
		for idx := 0; idx < 3; idx++ {
			m.SecurityType[idx] = m.SecurityTypeNullValue()
		}
	} else {
		if err := _m.ReadBytes(_r, m.SecurityType[:]); err != nil {
			return err
		}
	}
	if !m.SecurityIDInActingVersion(actingVersion) {
		m.SecurityID = m.SecurityIDNullValue()
	} else {
		if err := _m.ReadInt64(_r, &m.SecurityID); err != nil {
			return err
		}
	}
	if m.SecurityIDSourceInActingVersion(actingVersion) {
		if err := m.SecurityIDSource.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if !m.TransactTimeInActingVersion(actingVersion) {
		m.TransactTime = m.TransactTimeNullValue()
	} else {
		if err := _m.ReadUint64(_r, &m.TransactTime); err != nil {
			return err
		}
	}
	if m.BidPxInActingVersion(actingVersion) {
		if err := m.BidPx.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if !m.BidSizeInActingVersion(actingVersion) {
		m.BidSize = m.BidSizeNullValue()
	} else {
		if err := _m.ReadInt64(_r, &m.BidSize); err != nil {
			return err
		}
	}
	if m.OfferPxInActingVersion(actingVersion) {
		if err := m.OfferPx.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if !m.OfferSizeInActingVersion(actingVersion) {
		m.OfferSize = m.OfferSizeNullValue()
	} else {
		if err := _m.ReadInt64(_r, &m.OfferSize); err != nil {
			return err
		}
	}
	if actingVersion > m.SbeSchemaVersion() && blockLength > m.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-m.SbeBlockLength()))
	}
	return nil
}

func (m *MassQuoteQuoteSetsQuoteEntries) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if m.QuoteEntryIDInActingVersion(actingVersion) {
		for idx := 0; idx < 10; idx++ {
			if m.QuoteEntryID[idx] < m.QuoteEntryIDMinValue() || m.QuoteEntryID[idx] > m.QuoteEntryIDMaxValue() {
				return fmt.Errorf("Range check failed on m.QuoteEntryID[%d] (%v < %v > %v)", idx, m.QuoteEntryIDMinValue(), m.QuoteEntryID[idx], m.QuoteEntryIDMaxValue())
			}
		}
	}
	if m.SymbolInActingVersion(actingVersion) {
		for idx := 0; idx < 6; idx++ {
			if m.Symbol[idx] < m.SymbolMinValue() || m.Symbol[idx] > m.SymbolMaxValue() {
				return fmt.Errorf("Range check failed on m.Symbol[%d] (%v < %v > %v)", idx, m.SymbolMinValue(), m.Symbol[idx], m.SymbolMaxValue())
			}
		}
	}
	if m.SecurityDescInActingVersion(actingVersion) {
		for idx := 0; idx < 20; idx++ {
			if m.SecurityDesc[idx] < m.SecurityDescMinValue() || m.SecurityDesc[idx] > m.SecurityDescMaxValue() {
				return fmt.Errorf("Range check failed on m.SecurityDesc[%d] (%v < %v > %v)", idx, m.SecurityDescMinValue(), m.SecurityDesc[idx], m.SecurityDescMaxValue())
			}
		}
	}
	if m.SecurityTypeInActingVersion(actingVersion) {
		for idx := 0; idx < 3; idx++ {
			if m.SecurityType[idx] < m.SecurityTypeMinValue() || m.SecurityType[idx] > m.SecurityTypeMaxValue() {
				return fmt.Errorf("Range check failed on m.SecurityType[%d] (%v < %v > %v)", idx, m.SecurityTypeMinValue(), m.SecurityType[idx], m.SecurityTypeMaxValue())
			}
		}
	}
	if m.SecurityIDInActingVersion(actingVersion) {
		if m.SecurityID != m.SecurityIDNullValue() && (m.SecurityID < m.SecurityIDMinValue() || m.SecurityID > m.SecurityIDMaxValue()) {
			return fmt.Errorf("Range check failed on m.SecurityID (%v < %v > %v)", m.SecurityIDMinValue(), m.SecurityID, m.SecurityIDMaxValue())
		}
	}
	if err := m.SecurityIDSource.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if m.TransactTimeInActingVersion(actingVersion) {
		if m.TransactTime < m.TransactTimeMinValue() || m.TransactTime > m.TransactTimeMaxValue() {
			return fmt.Errorf("Range check failed on m.TransactTime (%v < %v > %v)", m.TransactTimeMinValue(), m.TransactTime, m.TransactTimeMaxValue())
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
	return nil
}

func MassQuoteQuoteSetsQuoteEntriesInit(m *MassQuoteQuoteSetsQuoteEntries) {
	m.SecurityID = math.MinInt64
	m.BidSize = math.MinInt64
	m.OfferSize = math.MinInt64
	return
}

func (*MassQuote) SbeBlockLength() (blockLength uint16) {
	return 62
}

func (*MassQuote) SbeTemplateId() (templateId uint16) {
	return 105
}

func (*MassQuote) SbeSchemaId() (schemaId uint16) {
	return 1
}

func (*MassQuote) SbeSchemaVersion() (schemaVersion uint16) {
	return 1
}

func (*MassQuote) SbeSemanticType() (semanticType []byte) {
	return []byte("i")
}

func (*MassQuote) QuoteReqIDId() uint16 {
	return 131
}

func (*MassQuote) QuoteReqIDSinceVersion() uint16 {
	return 0
}

func (m *MassQuote) QuoteReqIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.QuoteReqIDSinceVersion()
}

func (*MassQuote) QuoteReqIDDeprecated() uint16 {
	return 0
}

func (*MassQuote) QuoteReqIDMetaAttribute(meta int) string {
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

func (*MassQuote) QuoteReqIDMinValue() byte {
	return byte(32)
}

func (*MassQuote) QuoteReqIDMaxValue() byte {
	return byte(126)
}

func (*MassQuote) QuoteReqIDNullValue() byte {
	return 0
}

func (m *MassQuote) QuoteReqIDCharacterEncoding() string {
	return "US-ASCII"
}

func (*MassQuote) QuoteIDId() uint16 {
	return 117
}

func (*MassQuote) QuoteIDSinceVersion() uint16 {
	return 0
}

func (m *MassQuote) QuoteIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.QuoteIDSinceVersion()
}

func (*MassQuote) QuoteIDDeprecated() uint16 {
	return 0
}

func (*MassQuote) QuoteIDMetaAttribute(meta int) string {
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

func (*MassQuote) QuoteIDMinValue() byte {
	return byte(32)
}

func (*MassQuote) QuoteIDMaxValue() byte {
	return byte(126)
}

func (*MassQuote) QuoteIDNullValue() byte {
	return 0
}

func (m *MassQuote) QuoteIDCharacterEncoding() string {
	return "US-ASCII"
}

func (*MassQuote) MMAccountId() uint16 {
	return 9771
}

func (*MassQuote) MMAccountSinceVersion() uint16 {
	return 0
}

func (m *MassQuote) MMAccountInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.MMAccountSinceVersion()
}

func (*MassQuote) MMAccountDeprecated() uint16 {
	return 0
}

func (*MassQuote) MMAccountMetaAttribute(meta int) string {
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

func (*MassQuote) MMAccountMinValue() byte {
	return byte(32)
}

func (*MassQuote) MMAccountMaxValue() byte {
	return byte(126)
}

func (*MassQuote) MMAccountNullValue() byte {
	return 0
}

func (m *MassQuote) MMAccountCharacterEncoding() string {
	return "US-ASCII"
}

func (*MassQuote) ManualOrderIndicatorId() uint16 {
	return 1028
}

func (*MassQuote) ManualOrderIndicatorSinceVersion() uint16 {
	return 0
}

func (m *MassQuote) ManualOrderIndicatorInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.ManualOrderIndicatorSinceVersion()
}

func (*MassQuote) ManualOrderIndicatorDeprecated() uint16 {
	return 0
}

func (*MassQuote) ManualOrderIndicatorMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "required"
	}
	return ""
}

func (*MassQuote) CustOrderHandlingInstId() uint16 {
	return 1031
}

func (*MassQuote) CustOrderHandlingInstSinceVersion() uint16 {
	return 0
}

func (m *MassQuote) CustOrderHandlingInstInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.CustOrderHandlingInstSinceVersion()
}

func (*MassQuote) CustOrderHandlingInstDeprecated() uint16 {
	return 0
}

func (*MassQuote) CustOrderHandlingInstMetaAttribute(meta int) string {
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

func (*MassQuote) CustomerOrFirmId() uint16 {
	return 204
}

func (*MassQuote) CustomerOrFirmSinceVersion() uint16 {
	return 0
}

func (m *MassQuote) CustomerOrFirmInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.CustomerOrFirmSinceVersion()
}

func (*MassQuote) CustomerOrFirmDeprecated() uint16 {
	return 0
}

func (*MassQuote) CustomerOrFirmMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "required"
	}
	return ""
}

func (*MassQuote) SelfMatchPreventionIDId() uint16 {
	return 7928
}

func (*MassQuote) SelfMatchPreventionIDSinceVersion() uint16 {
	return 0
}

func (m *MassQuote) SelfMatchPreventionIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.SelfMatchPreventionIDSinceVersion()
}

func (*MassQuote) SelfMatchPreventionIDDeprecated() uint16 {
	return 0
}

func (*MassQuote) SelfMatchPreventionIDMetaAttribute(meta int) string {
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

func (*MassQuote) SelfMatchPreventionIDMinValue() byte {
	return byte(32)
}

func (*MassQuote) SelfMatchPreventionIDMaxValue() byte {
	return byte(126)
}

func (*MassQuote) SelfMatchPreventionIDNullValue() byte {
	return 0
}

func (m *MassQuote) SelfMatchPreventionIDCharacterEncoding() string {
	return "US-ASCII"
}

func (*MassQuote) CtiCodeId() uint16 {
	return 9702
}

func (*MassQuote) CtiCodeSinceVersion() uint16 {
	return 0
}

func (m *MassQuote) CtiCodeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.CtiCodeSinceVersion()
}

func (*MassQuote) CtiCodeDeprecated() uint16 {
	return 0
}

func (*MassQuote) CtiCodeMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "required"
	}
	return ""
}

func (*MassQuote) MMProtectionResetId() uint16 {
	return 9773
}

func (*MassQuote) MMProtectionResetSinceVersion() uint16 {
	return 0
}

func (m *MassQuote) MMProtectionResetInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.MMProtectionResetSinceVersion()
}

func (*MassQuote) MMProtectionResetDeprecated() uint16 {
	return 0
}

func (*MassQuote) MMProtectionResetMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "optional"
	}
	return ""
}

func (*MassQuoteQuoteSets) QuoteSetIDId() uint16 {
	return 302
}

func (*MassQuoteQuoteSets) QuoteSetIDSinceVersion() uint16 {
	return 0
}

func (m *MassQuoteQuoteSets) QuoteSetIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.QuoteSetIDSinceVersion()
}

func (*MassQuoteQuoteSets) QuoteSetIDDeprecated() uint16 {
	return 0
}

func (*MassQuoteQuoteSets) QuoteSetIDMetaAttribute(meta int) string {
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

func (*MassQuoteQuoteSets) QuoteSetIDMinValue() byte {
	return byte(32)
}

func (*MassQuoteQuoteSets) QuoteSetIDMaxValue() byte {
	return byte(126)
}

func (*MassQuoteQuoteSets) QuoteSetIDNullValue() byte {
	return 0
}

func (m *MassQuoteQuoteSets) QuoteSetIDCharacterEncoding() string {
	return "US-ASCII"
}

func (*MassQuoteQuoteSets) UnderlyingSecurityDescId() uint16 {
	return 307
}

func (*MassQuoteQuoteSets) UnderlyingSecurityDescSinceVersion() uint16 {
	return 0
}

func (m *MassQuoteQuoteSets) UnderlyingSecurityDescInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.UnderlyingSecurityDescSinceVersion()
}

func (*MassQuoteQuoteSets) UnderlyingSecurityDescDeprecated() uint16 {
	return 0
}

func (*MassQuoteQuoteSets) UnderlyingSecurityDescMetaAttribute(meta int) string {
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

func (*MassQuoteQuoteSets) UnderlyingSecurityDescMinValue() byte {
	return byte(32)
}

func (*MassQuoteQuoteSets) UnderlyingSecurityDescMaxValue() byte {
	return byte(126)
}

func (*MassQuoteQuoteSets) UnderlyingSecurityDescNullValue() byte {
	return 0
}

func (m *MassQuoteQuoteSets) UnderlyingSecurityDescCharacterEncoding() string {
	return "US-ASCII"
}

func (*MassQuoteQuoteSets) TotQuoteEntriesId() uint16 {
	return 304
}

func (*MassQuoteQuoteSets) TotQuoteEntriesSinceVersion() uint16 {
	return 0
}

func (m *MassQuoteQuoteSets) TotQuoteEntriesInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.TotQuoteEntriesSinceVersion()
}

func (*MassQuoteQuoteSets) TotQuoteEntriesDeprecated() uint16 {
	return 0
}

func (*MassQuoteQuoteSets) TotQuoteEntriesMetaAttribute(meta int) string {
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

func (*MassQuoteQuoteSets) TotQuoteEntriesMinValue() uint8 {
	return 0
}

func (*MassQuoteQuoteSets) TotQuoteEntriesMaxValue() uint8 {
	return math.MaxUint8 - 1
}

func (*MassQuoteQuoteSets) TotQuoteEntriesNullValue() uint8 {
	return math.MaxUint8
}

func (*MassQuoteQuoteSetsQuoteEntries) QuoteEntryIDId() uint16 {
	return 299
}

func (*MassQuoteQuoteSetsQuoteEntries) QuoteEntryIDSinceVersion() uint16 {
	return 0
}

func (m *MassQuoteQuoteSetsQuoteEntries) QuoteEntryIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.QuoteEntryIDSinceVersion()
}

func (*MassQuoteQuoteSetsQuoteEntries) QuoteEntryIDDeprecated() uint16 {
	return 0
}

func (*MassQuoteQuoteSetsQuoteEntries) QuoteEntryIDMetaAttribute(meta int) string {
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

func (*MassQuoteQuoteSetsQuoteEntries) QuoteEntryIDMinValue() byte {
	return byte(32)
}

func (*MassQuoteQuoteSetsQuoteEntries) QuoteEntryIDMaxValue() byte {
	return byte(126)
}

func (*MassQuoteQuoteSetsQuoteEntries) QuoteEntryIDNullValue() byte {
	return 0
}

func (m *MassQuoteQuoteSetsQuoteEntries) QuoteEntryIDCharacterEncoding() string {
	return "US-ASCII"
}

func (*MassQuoteQuoteSetsQuoteEntries) SymbolId() uint16 {
	return 55
}

func (*MassQuoteQuoteSetsQuoteEntries) SymbolSinceVersion() uint16 {
	return 0
}

func (m *MassQuoteQuoteSetsQuoteEntries) SymbolInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.SymbolSinceVersion()
}

func (*MassQuoteQuoteSetsQuoteEntries) SymbolDeprecated() uint16 {
	return 0
}

func (*MassQuoteQuoteSetsQuoteEntries) SymbolMetaAttribute(meta int) string {
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

func (*MassQuoteQuoteSetsQuoteEntries) SymbolMinValue() byte {
	return byte(32)
}

func (*MassQuoteQuoteSetsQuoteEntries) SymbolMaxValue() byte {
	return byte(126)
}

func (*MassQuoteQuoteSetsQuoteEntries) SymbolNullValue() byte {
	return 0
}

func (m *MassQuoteQuoteSetsQuoteEntries) SymbolCharacterEncoding() string {
	return "US-ASCII"
}

func (*MassQuoteQuoteSetsQuoteEntries) SecurityDescId() uint16 {
	return 107
}

func (*MassQuoteQuoteSetsQuoteEntries) SecurityDescSinceVersion() uint16 {
	return 0
}

func (m *MassQuoteQuoteSetsQuoteEntries) SecurityDescInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.SecurityDescSinceVersion()
}

func (*MassQuoteQuoteSetsQuoteEntries) SecurityDescDeprecated() uint16 {
	return 0
}

func (*MassQuoteQuoteSetsQuoteEntries) SecurityDescMetaAttribute(meta int) string {
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

func (*MassQuoteQuoteSetsQuoteEntries) SecurityDescMinValue() byte {
	return byte(32)
}

func (*MassQuoteQuoteSetsQuoteEntries) SecurityDescMaxValue() byte {
	return byte(126)
}

func (*MassQuoteQuoteSetsQuoteEntries) SecurityDescNullValue() byte {
	return 0
}

func (m *MassQuoteQuoteSetsQuoteEntries) SecurityDescCharacterEncoding() string {
	return "US-ASCII"
}

func (*MassQuoteQuoteSetsQuoteEntries) SecurityTypeId() uint16 {
	return 167
}

func (*MassQuoteQuoteSetsQuoteEntries) SecurityTypeSinceVersion() uint16 {
	return 0
}

func (m *MassQuoteQuoteSetsQuoteEntries) SecurityTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.SecurityTypeSinceVersion()
}

func (*MassQuoteQuoteSetsQuoteEntries) SecurityTypeDeprecated() uint16 {
	return 0
}

func (*MassQuoteQuoteSetsQuoteEntries) SecurityTypeMetaAttribute(meta int) string {
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

func (*MassQuoteQuoteSetsQuoteEntries) SecurityTypeMinValue() byte {
	return byte(32)
}

func (*MassQuoteQuoteSetsQuoteEntries) SecurityTypeMaxValue() byte {
	return byte(126)
}

func (*MassQuoteQuoteSetsQuoteEntries) SecurityTypeNullValue() byte {
	return 0
}

func (m *MassQuoteQuoteSetsQuoteEntries) SecurityTypeCharacterEncoding() string {
	return "US-ASCII"
}

func (*MassQuoteQuoteSetsQuoteEntries) SecurityIDId() uint16 {
	return 48
}

func (*MassQuoteQuoteSetsQuoteEntries) SecurityIDSinceVersion() uint16 {
	return 0
}

func (m *MassQuoteQuoteSetsQuoteEntries) SecurityIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.SecurityIDSinceVersion()
}

func (*MassQuoteQuoteSetsQuoteEntries) SecurityIDDeprecated() uint16 {
	return 0
}

func (*MassQuoteQuoteSetsQuoteEntries) SecurityIDMetaAttribute(meta int) string {
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

func (*MassQuoteQuoteSetsQuoteEntries) SecurityIDMinValue() int64 {
	return math.MinInt64 + 1
}

func (*MassQuoteQuoteSetsQuoteEntries) SecurityIDMaxValue() int64 {
	return math.MaxInt64
}

func (*MassQuoteQuoteSetsQuoteEntries) SecurityIDNullValue() int64 {
	return math.MinInt64
}

func (*MassQuoteQuoteSetsQuoteEntries) SecurityIDSourceId() uint16 {
	return 22
}

func (*MassQuoteQuoteSetsQuoteEntries) SecurityIDSourceSinceVersion() uint16 {
	return 0
}

func (m *MassQuoteQuoteSetsQuoteEntries) SecurityIDSourceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.SecurityIDSourceSinceVersion()
}

func (*MassQuoteQuoteSetsQuoteEntries) SecurityIDSourceDeprecated() uint16 {
	return 0
}

func (*MassQuoteQuoteSetsQuoteEntries) SecurityIDSourceMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "optional"
	}
	return ""
}

func (*MassQuoteQuoteSetsQuoteEntries) TransactTimeId() uint16 {
	return 60
}

func (*MassQuoteQuoteSetsQuoteEntries) TransactTimeSinceVersion() uint16 {
	return 0
}

func (m *MassQuoteQuoteSetsQuoteEntries) TransactTimeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.TransactTimeSinceVersion()
}

func (*MassQuoteQuoteSetsQuoteEntries) TransactTimeDeprecated() uint16 {
	return 0
}

func (*MassQuoteQuoteSetsQuoteEntries) TransactTimeMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return "nanosecond"
	case 3:
		return "UTCTimestamp"
	case 4:
		return "required"
	}
	return ""
}

func (*MassQuoteQuoteSetsQuoteEntries) TransactTimeMinValue() uint64 {
	return 0
}

func (*MassQuoteQuoteSetsQuoteEntries) TransactTimeMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*MassQuoteQuoteSetsQuoteEntries) TransactTimeNullValue() uint64 {
	return math.MaxUint64
}

func (*MassQuoteQuoteSetsQuoteEntries) BidPxId() uint16 {
	return 132
}

func (*MassQuoteQuoteSetsQuoteEntries) BidPxSinceVersion() uint16 {
	return 0
}

func (m *MassQuoteQuoteSetsQuoteEntries) BidPxInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.BidPxSinceVersion()
}

func (*MassQuoteQuoteSetsQuoteEntries) BidPxDeprecated() uint16 {
	return 0
}

func (*MassQuoteQuoteSetsQuoteEntries) BidPxMetaAttribute(meta int) string {
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

func (*MassQuoteQuoteSetsQuoteEntries) BidSizeId() uint16 {
	return 134
}

func (*MassQuoteQuoteSetsQuoteEntries) BidSizeSinceVersion() uint16 {
	return 0
}

func (m *MassQuoteQuoteSetsQuoteEntries) BidSizeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.BidSizeSinceVersion()
}

func (*MassQuoteQuoteSetsQuoteEntries) BidSizeDeprecated() uint16 {
	return 0
}

func (*MassQuoteQuoteSetsQuoteEntries) BidSizeMetaAttribute(meta int) string {
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

func (*MassQuoteQuoteSetsQuoteEntries) BidSizeMinValue() int64 {
	return math.MinInt64 + 1
}

func (*MassQuoteQuoteSetsQuoteEntries) BidSizeMaxValue() int64 {
	return math.MaxInt64
}

func (*MassQuoteQuoteSetsQuoteEntries) BidSizeNullValue() int64 {
	return math.MinInt64
}

func (*MassQuoteQuoteSetsQuoteEntries) OfferPxId() uint16 {
	return 133
}

func (*MassQuoteQuoteSetsQuoteEntries) OfferPxSinceVersion() uint16 {
	return 0
}

func (m *MassQuoteQuoteSetsQuoteEntries) OfferPxInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.OfferPxSinceVersion()
}

func (*MassQuoteQuoteSetsQuoteEntries) OfferPxDeprecated() uint16 {
	return 0
}

func (*MassQuoteQuoteSetsQuoteEntries) OfferPxMetaAttribute(meta int) string {
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

func (*MassQuoteQuoteSetsQuoteEntries) OfferSizeId() uint16 {
	return 135
}

func (*MassQuoteQuoteSetsQuoteEntries) OfferSizeSinceVersion() uint16 {
	return 0
}

func (m *MassQuoteQuoteSetsQuoteEntries) OfferSizeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.OfferSizeSinceVersion()
}

func (*MassQuoteQuoteSetsQuoteEntries) OfferSizeDeprecated() uint16 {
	return 0
}

func (*MassQuoteQuoteSetsQuoteEntries) OfferSizeMetaAttribute(meta int) string {
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

func (*MassQuoteQuoteSetsQuoteEntries) OfferSizeMinValue() int64 {
	return math.MinInt64 + 1
}

func (*MassQuoteQuoteSetsQuoteEntries) OfferSizeMaxValue() int64 {
	return math.MaxInt64
}

func (*MassQuoteQuoteSetsQuoteEntries) OfferSizeNullValue() int64 {
	return math.MinInt64
}

func (*MassQuote) QuoteSetsId() uint16 {
	return 296
}

func (*MassQuote) QuoteSetsSinceVersion() uint16 {
	return 0
}

func (m *MassQuote) QuoteSetsInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.QuoteSetsSinceVersion()
}

func (*MassQuote) QuoteSetsDeprecated() uint16 {
	return 0
}

func (*MassQuoteQuoteSets) SbeBlockLength() (blockLength uint) {
	return 24
}

func (*MassQuoteQuoteSets) SbeSchemaVersion() (schemaVersion uint16) {
	return 1
}

func (*MassQuoteQuoteSets) QuoteEntriesId() uint16 {
	return 295
}

func (*MassQuoteQuoteSets) QuoteEntriesSinceVersion() uint16 {
	return 0
}

func (m *MassQuoteQuoteSets) QuoteEntriesInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.QuoteEntriesSinceVersion()
}

func (*MassQuoteQuoteSets) QuoteEntriesDeprecated() uint16 {
	return 0
}

func (*MassQuoteQuoteSetsQuoteEntries) SbeBlockLength() (blockLength uint) {
	return 90
}

func (*MassQuoteQuoteSetsQuoteEntries) SbeSchemaVersion() (schemaVersion uint16) {
	return 1
}
