// Generated SBE (Simple Binary Encoding) message codec

package fix

import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
)

type MarketDataIncrementalRefresh struct {
	TradeDate uint16
	Entries   []MarketDataIncrementalRefreshEntries
}
type MarketDataIncrementalRefreshEntries struct {
	MdUpdateAction      MDUpdateActionEnum
	MdPriceLevel        uint8
	MdEntryType         MDEntryTypeEnum
	SecurityIdSource    byte
	SecurityId          uint64
	RptSeq              uint8
	QuoteCondition      QuoteCondition
	MdEntryPx           Decimal64
	NumberOfOrders      uint32
	MdEntryTime         uint64
	MdEntrySize         IntQty32
	TradingSessionId    MarketStateIdentifierEnum
	NetChgPrevDay       Decimal64
	TickDirection       TickDirectionEnum
	OpenCloseSettleFlag OpenCloseSettleFlagEnum
	SettleDate          uint64
	TradeCondition      TradeCondition
	TradeVolume         IntQty32
	MdQuoteType         MDQuoteTypeEnum
	FixingBracket       uint64
	AggressorSide       SideEnum
	MatchEventIndicator MatchEventIndicatorEnum
	TradeId             uint64
}

func (m *MarketDataIncrementalRefresh) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := m.RangeCheck(m.SbeSchemaVersion(), m.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	if err := _m.WriteUint16(_w, m.TradeDate); err != nil {
		return err
	}
	var EntriesBlockLength uint16 = 82
	if err := _m.WriteUint16(_w, EntriesBlockLength); err != nil {
		return err
	}
	var EntriesNumInGroup uint16 = uint16(len(m.Entries))
	if err := _m.WriteUint16(_w, EntriesNumInGroup); err != nil {
		return err
	}
	for _, prop := range m.Entries {
		if err := prop.Encode(_m, _w); err != nil {
			return err
		}
	}
	return nil
}

func (m *MarketDataIncrementalRefresh) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	if !m.TradeDateInActingVersion(actingVersion) {
		m.TradeDate = m.TradeDateNullValue()
	} else {
		if err := _m.ReadUint16(_r, &m.TradeDate); err != nil {
			return err
		}
	}
	if actingVersion > m.SbeSchemaVersion() && blockLength > m.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-m.SbeBlockLength()))
	}

	if m.EntriesInActingVersion(actingVersion) {
		var EntriesBlockLength uint16
		if err := _m.ReadUint16(_r, &EntriesBlockLength); err != nil {
			return err
		}
		var EntriesNumInGroup uint16
		if err := _m.ReadUint16(_r, &EntriesNumInGroup); err != nil {
			return err
		}
		if cap(m.Entries) < int(EntriesNumInGroup) {
			m.Entries = make([]MarketDataIncrementalRefreshEntries, EntriesNumInGroup)
		}
		for i, _ := range m.Entries {
			if err := m.Entries[i].Decode(_m, _r, actingVersion, uint(EntriesBlockLength)); err != nil {
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

func (m *MarketDataIncrementalRefresh) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if m.TradeDateInActingVersion(actingVersion) {
		if m.TradeDate < m.TradeDateMinValue() || m.TradeDate > m.TradeDateMaxValue() {
			return fmt.Errorf("Range check failed on m.TradeDate (%v < %v > %v)", m.TradeDateMinValue(), m.TradeDate, m.TradeDateMaxValue())
		}
	}
	for _, prop := range m.Entries {
		if err := prop.RangeCheck(actingVersion, schemaVersion); err != nil {
			return err
		}
	}
	return nil
}

func MarketDataIncrementalRefreshInit(m *MarketDataIncrementalRefresh) {
	return
}

func (m *MarketDataIncrementalRefreshEntries) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := m.MdUpdateAction.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteUint8(_w, m.MdPriceLevel); err != nil {
		return err
	}
	if err := m.MdEntryType.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteUint8(_w, m.SecurityIdSource); err != nil {
		return err
	}
	if err := _m.WriteUint64(_w, m.SecurityId); err != nil {
		return err
	}
	if err := _m.WriteUint8(_w, m.RptSeq); err != nil {
		return err
	}
	if err := m.QuoteCondition.Encode(_m, _w); err != nil {
		return err
	}
	if err := m.MdEntryPx.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteUint32(_w, m.NumberOfOrders); err != nil {
		return err
	}
	if err := _m.WriteUint64(_w, m.MdEntryTime); err != nil {
		return err
	}
	if err := m.MdEntrySize.Encode(_m, _w); err != nil {
		return err
	}
	if err := m.TradingSessionId.Encode(_m, _w); err != nil {
		return err
	}
	if err := m.NetChgPrevDay.Encode(_m, _w); err != nil {
		return err
	}
	if err := m.TickDirection.Encode(_m, _w); err != nil {
		return err
	}
	if err := m.OpenCloseSettleFlag.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteUint64(_w, m.SettleDate); err != nil {
		return err
	}
	if err := m.TradeCondition.Encode(_m, _w); err != nil {
		return err
	}
	if err := m.TradeVolume.Encode(_m, _w); err != nil {
		return err
	}
	if err := m.MdQuoteType.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteUint64(_w, m.FixingBracket); err != nil {
		return err
	}
	if err := m.AggressorSide.Encode(_m, _w); err != nil {
		return err
	}
	if err := m.MatchEventIndicator.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteUint64(_w, m.TradeId); err != nil {
		return err
	}
	return nil
}

func (m *MarketDataIncrementalRefreshEntries) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint) error {
	if m.MdUpdateActionInActingVersion(actingVersion) {
		if err := m.MdUpdateAction.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if !m.MdPriceLevelInActingVersion(actingVersion) {
		m.MdPriceLevel = m.MdPriceLevelNullValue()
	} else {
		if err := _m.ReadUint8(_r, &m.MdPriceLevel); err != nil {
			return err
		}
	}
	if m.MdEntryTypeInActingVersion(actingVersion) {
		if err := m.MdEntryType.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if !m.SecurityIdSourceInActingVersion(actingVersion) {
		m.SecurityIdSource = m.SecurityIdSourceNullValue()
	} else {
		if err := _m.ReadUint8(_r, &m.SecurityIdSource); err != nil {
			return err
		}
	}
	if !m.SecurityIdInActingVersion(actingVersion) {
		m.SecurityId = m.SecurityIdNullValue()
	} else {
		if err := _m.ReadUint64(_r, &m.SecurityId); err != nil {
			return err
		}
	}
	if !m.RptSeqInActingVersion(actingVersion) {
		m.RptSeq = m.RptSeqNullValue()
	} else {
		if err := _m.ReadUint8(_r, &m.RptSeq); err != nil {
			return err
		}
	}
	if m.QuoteConditionInActingVersion(actingVersion) {
		if err := m.QuoteCondition.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if m.MdEntryPxInActingVersion(actingVersion) {
		if err := m.MdEntryPx.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if !m.NumberOfOrdersInActingVersion(actingVersion) {
		m.NumberOfOrders = m.NumberOfOrdersNullValue()
	} else {
		if err := _m.ReadUint32(_r, &m.NumberOfOrders); err != nil {
			return err
		}
	}
	if !m.MdEntryTimeInActingVersion(actingVersion) {
		m.MdEntryTime = m.MdEntryTimeNullValue()
	} else {
		if err := _m.ReadUint64(_r, &m.MdEntryTime); err != nil {
			return err
		}
	}
	if m.MdEntrySizeInActingVersion(actingVersion) {
		if err := m.MdEntrySize.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if m.TradingSessionIdInActingVersion(actingVersion) {
		if err := m.TradingSessionId.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if m.NetChgPrevDayInActingVersion(actingVersion) {
		if err := m.NetChgPrevDay.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if m.TickDirectionInActingVersion(actingVersion) {
		if err := m.TickDirection.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if m.OpenCloseSettleFlagInActingVersion(actingVersion) {
		if err := m.OpenCloseSettleFlag.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if !m.SettleDateInActingVersion(actingVersion) {
		m.SettleDate = m.SettleDateNullValue()
	} else {
		if err := _m.ReadUint64(_r, &m.SettleDate); err != nil {
			return err
		}
	}
	if m.TradeConditionInActingVersion(actingVersion) {
		if err := m.TradeCondition.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if m.TradeVolumeInActingVersion(actingVersion) {
		if err := m.TradeVolume.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if m.MdQuoteTypeInActingVersion(actingVersion) {
		if err := m.MdQuoteType.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if !m.FixingBracketInActingVersion(actingVersion) {
		m.FixingBracket = m.FixingBracketNullValue()
	} else {
		if err := _m.ReadUint64(_r, &m.FixingBracket); err != nil {
			return err
		}
	}
	if m.AggressorSideInActingVersion(actingVersion) {
		if err := m.AggressorSide.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if m.MatchEventIndicatorInActingVersion(actingVersion) {
		if err := m.MatchEventIndicator.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if !m.TradeIdInActingVersion(actingVersion) {
		m.TradeId = m.TradeIdNullValue()
	} else {
		if err := _m.ReadUint64(_r, &m.TradeId); err != nil {
			return err
		}
	}
	if actingVersion > m.SbeSchemaVersion() && blockLength > m.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-m.SbeBlockLength()))
	}
	return nil
}

func (m *MarketDataIncrementalRefreshEntries) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if err := m.MdUpdateAction.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if m.MdPriceLevelInActingVersion(actingVersion) {
		if m.MdPriceLevel < m.MdPriceLevelMinValue() || m.MdPriceLevel > m.MdPriceLevelMaxValue() {
			return fmt.Errorf("Range check failed on m.MdPriceLevel (%v < %v > %v)", m.MdPriceLevelMinValue(), m.MdPriceLevel, m.MdPriceLevelMaxValue())
		}
	}
	if err := m.MdEntryType.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if m.SecurityIdSourceInActingVersion(actingVersion) {
		if m.SecurityIdSource < m.SecurityIdSourceMinValue() || m.SecurityIdSource > m.SecurityIdSourceMaxValue() {
			return fmt.Errorf("Range check failed on m.SecurityIdSource (%v < %v > %v)", m.SecurityIdSourceMinValue(), m.SecurityIdSource, m.SecurityIdSourceMaxValue())
		}
	}
	if m.SecurityIdInActingVersion(actingVersion) {
		if m.SecurityId < m.SecurityIdMinValue() || m.SecurityId > m.SecurityIdMaxValue() {
			return fmt.Errorf("Range check failed on m.SecurityId (%v < %v > %v)", m.SecurityIdMinValue(), m.SecurityId, m.SecurityIdMaxValue())
		}
	}
	if m.RptSeqInActingVersion(actingVersion) {
		if m.RptSeq < m.RptSeqMinValue() || m.RptSeq > m.RptSeqMaxValue() {
			return fmt.Errorf("Range check failed on m.RptSeq (%v < %v > %v)", m.RptSeqMinValue(), m.RptSeq, m.RptSeqMaxValue())
		}
	}
	if m.NumberOfOrdersInActingVersion(actingVersion) {
		if m.NumberOfOrders < m.NumberOfOrdersMinValue() || m.NumberOfOrders > m.NumberOfOrdersMaxValue() {
			return fmt.Errorf("Range check failed on m.NumberOfOrders (%v < %v > %v)", m.NumberOfOrdersMinValue(), m.NumberOfOrders, m.NumberOfOrdersMaxValue())
		}
	}
	if m.MdEntryTimeInActingVersion(actingVersion) {
		if m.MdEntryTime < m.MdEntryTimeMinValue() || m.MdEntryTime > m.MdEntryTimeMaxValue() {
			return fmt.Errorf("Range check failed on m.MdEntryTime (%v < %v > %v)", m.MdEntryTimeMinValue(), m.MdEntryTime, m.MdEntryTimeMaxValue())
		}
	}
	if err := m.TradingSessionId.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if err := m.TickDirection.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if err := m.OpenCloseSettleFlag.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if m.SettleDateInActingVersion(actingVersion) {
		if m.SettleDate < m.SettleDateMinValue() || m.SettleDate > m.SettleDateMaxValue() {
			return fmt.Errorf("Range check failed on m.SettleDate (%v < %v > %v)", m.SettleDateMinValue(), m.SettleDate, m.SettleDateMaxValue())
		}
	}
	if err := m.MdQuoteType.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if m.FixingBracketInActingVersion(actingVersion) {
		if m.FixingBracket < m.FixingBracketMinValue() || m.FixingBracket > m.FixingBracketMaxValue() {
			return fmt.Errorf("Range check failed on m.FixingBracket (%v < %v > %v)", m.FixingBracketMinValue(), m.FixingBracket, m.FixingBracketMaxValue())
		}
	}
	if err := m.AggressorSide.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if err := m.MatchEventIndicator.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if m.TradeIdInActingVersion(actingVersion) {
		if m.TradeId < m.TradeIdMinValue() || m.TradeId > m.TradeIdMaxValue() {
			return fmt.Errorf("Range check failed on m.TradeId (%v < %v > %v)", m.TradeIdMinValue(), m.TradeId, m.TradeIdMaxValue())
		}
	}
	return nil
}

func MarketDataIncrementalRefreshEntriesInit(m *MarketDataIncrementalRefreshEntries) {
	return
}

func (*MarketDataIncrementalRefresh) SbeBlockLength() (blockLength uint16) {
	return 2
}

func (*MarketDataIncrementalRefresh) SbeTemplateId() (templateId uint16) {
	return 88
}

func (*MarketDataIncrementalRefresh) SbeSchemaId() (schemaId uint16) {
	return 1
}

func (*MarketDataIncrementalRefresh) SbeSchemaVersion() (schemaVersion uint16) {
	return 1
}

func (*MarketDataIncrementalRefresh) SbeSemanticType() (semanticType []byte) {
	return []byte("X")
}

func (*MarketDataIncrementalRefresh) TradeDateId() uint16 {
	return 75
}

func (*MarketDataIncrementalRefresh) TradeDateSinceVersion() uint16 {
	return 0
}

func (m *MarketDataIncrementalRefresh) TradeDateInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.TradeDateSinceVersion()
}

func (*MarketDataIncrementalRefresh) TradeDateDeprecated() uint16 {
	return 0
}

func (*MarketDataIncrementalRefresh) TradeDateMetaAttribute(meta int) string {
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

func (*MarketDataIncrementalRefresh) TradeDateMinValue() uint16 {
	return 0
}

func (*MarketDataIncrementalRefresh) TradeDateMaxValue() uint16 {
	return math.MaxUint16 - 1
}

func (*MarketDataIncrementalRefresh) TradeDateNullValue() uint16 {
	return math.MaxUint16
}

func (*MarketDataIncrementalRefreshEntries) MdUpdateActionId() uint16 {
	return 279
}

func (*MarketDataIncrementalRefreshEntries) MdUpdateActionSinceVersion() uint16 {
	return 0
}

func (m *MarketDataIncrementalRefreshEntries) MdUpdateActionInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.MdUpdateActionSinceVersion()
}

func (*MarketDataIncrementalRefreshEntries) MdUpdateActionDeprecated() uint16 {
	return 0
}

func (*MarketDataIncrementalRefreshEntries) MdUpdateActionMetaAttribute(meta int) string {
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

func (*MarketDataIncrementalRefreshEntries) MdPriceLevelId() uint16 {
	return 1023
}

func (*MarketDataIncrementalRefreshEntries) MdPriceLevelSinceVersion() uint16 {
	return 0
}

func (m *MarketDataIncrementalRefreshEntries) MdPriceLevelInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.MdPriceLevelSinceVersion()
}

func (*MarketDataIncrementalRefreshEntries) MdPriceLevelDeprecated() uint16 {
	return 0
}

func (*MarketDataIncrementalRefreshEntries) MdPriceLevelMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return "MDPriceLevel"
	case 4:
		return "required"
	}
	return ""
}

func (*MarketDataIncrementalRefreshEntries) MdPriceLevelMinValue() uint8 {
	return 0
}

func (*MarketDataIncrementalRefreshEntries) MdPriceLevelMaxValue() uint8 {
	return math.MaxUint8 - 1
}

func (*MarketDataIncrementalRefreshEntries) MdPriceLevelNullValue() uint8 {
	return math.MaxUint8
}

func (*MarketDataIncrementalRefreshEntries) MdEntryTypeId() uint16 {
	return 269
}

func (*MarketDataIncrementalRefreshEntries) MdEntryTypeSinceVersion() uint16 {
	return 0
}

func (m *MarketDataIncrementalRefreshEntries) MdEntryTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.MdEntryTypeSinceVersion()
}

func (*MarketDataIncrementalRefreshEntries) MdEntryTypeDeprecated() uint16 {
	return 0
}

func (*MarketDataIncrementalRefreshEntries) MdEntryTypeMetaAttribute(meta int) string {
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

func (*MarketDataIncrementalRefreshEntries) SecurityIdSourceId() uint16 {
	return 22
}

func (*MarketDataIncrementalRefreshEntries) SecurityIdSourceSinceVersion() uint16 {
	return 0
}

func (m *MarketDataIncrementalRefreshEntries) SecurityIdSourceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.SecurityIdSourceSinceVersion()
}

func (*MarketDataIncrementalRefreshEntries) SecurityIdSourceDeprecated() uint16 {
	return 0
}

func (*MarketDataIncrementalRefreshEntries) SecurityIdSourceMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return "SecurityID"
	case 4:
		return "required"
	}
	return ""
}

func (*MarketDataIncrementalRefreshEntries) SecurityIdSourceMinValue() byte {
	return byte(32)
}

func (*MarketDataIncrementalRefreshEntries) SecurityIdSourceMaxValue() byte {
	return byte(126)
}

func (*MarketDataIncrementalRefreshEntries) SecurityIdSourceNullValue() byte {
	return 0
}

func (*MarketDataIncrementalRefreshEntries) SecurityIdId() uint16 {
	return 48
}

func (*MarketDataIncrementalRefreshEntries) SecurityIdSinceVersion() uint16 {
	return 0
}

func (m *MarketDataIncrementalRefreshEntries) SecurityIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.SecurityIdSinceVersion()
}

func (*MarketDataIncrementalRefreshEntries) SecurityIdDeprecated() uint16 {
	return 0
}

func (*MarketDataIncrementalRefreshEntries) SecurityIdMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return "InstrumentID"
	case 4:
		return "required"
	}
	return ""
}

func (*MarketDataIncrementalRefreshEntries) SecurityIdMinValue() uint64 {
	return 0
}

func (*MarketDataIncrementalRefreshEntries) SecurityIdMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*MarketDataIncrementalRefreshEntries) SecurityIdNullValue() uint64 {
	return math.MaxUint64
}

func (*MarketDataIncrementalRefreshEntries) RptSeqId() uint16 {
	return 83
}

func (*MarketDataIncrementalRefreshEntries) RptSeqSinceVersion() uint16 {
	return 0
}

func (m *MarketDataIncrementalRefreshEntries) RptSeqInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.RptSeqSinceVersion()
}

func (*MarketDataIncrementalRefreshEntries) RptSeqDeprecated() uint16 {
	return 0
}

func (*MarketDataIncrementalRefreshEntries) RptSeqMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return "SequenceNumber"
	case 4:
		return "required"
	}
	return ""
}

func (*MarketDataIncrementalRefreshEntries) RptSeqMinValue() uint8 {
	return 0
}

func (*MarketDataIncrementalRefreshEntries) RptSeqMaxValue() uint8 {
	return math.MaxUint8 - 1
}

func (*MarketDataIncrementalRefreshEntries) RptSeqNullValue() uint8 {
	return math.MaxUint8
}

func (*MarketDataIncrementalRefreshEntries) QuoteConditionId() uint16 {
	return 276
}

func (*MarketDataIncrementalRefreshEntries) QuoteConditionSinceVersion() uint16 {
	return 0
}

func (m *MarketDataIncrementalRefreshEntries) QuoteConditionInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.QuoteConditionSinceVersion()
}

func (*MarketDataIncrementalRefreshEntries) QuoteConditionDeprecated() uint16 {
	return 0
}

func (*MarketDataIncrementalRefreshEntries) QuoteConditionMetaAttribute(meta int) string {
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

func (*MarketDataIncrementalRefreshEntries) MdEntryPxId() uint16 {
	return 270
}

func (*MarketDataIncrementalRefreshEntries) MdEntryPxSinceVersion() uint16 {
	return 0
}

func (m *MarketDataIncrementalRefreshEntries) MdEntryPxInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.MdEntryPxSinceVersion()
}

func (*MarketDataIncrementalRefreshEntries) MdEntryPxDeprecated() uint16 {
	return 0
}

func (*MarketDataIncrementalRefreshEntries) MdEntryPxMetaAttribute(meta int) string {
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

func (*MarketDataIncrementalRefreshEntries) NumberOfOrdersId() uint16 {
	return 346
}

func (*MarketDataIncrementalRefreshEntries) NumberOfOrdersSinceVersion() uint16 {
	return 0
}

func (m *MarketDataIncrementalRefreshEntries) NumberOfOrdersInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.NumberOfOrdersSinceVersion()
}

func (*MarketDataIncrementalRefreshEntries) NumberOfOrdersDeprecated() uint16 {
	return 0
}

func (*MarketDataIncrementalRefreshEntries) NumberOfOrdersMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return "NumberOfOrders"
	case 4:
		return "required"
	}
	return ""
}

func (*MarketDataIncrementalRefreshEntries) NumberOfOrdersMinValue() uint32 {
	return 0
}

func (*MarketDataIncrementalRefreshEntries) NumberOfOrdersMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*MarketDataIncrementalRefreshEntries) NumberOfOrdersNullValue() uint32 {
	return math.MaxUint32
}

func (*MarketDataIncrementalRefreshEntries) MdEntryTimeId() uint16 {
	return 273
}

func (*MarketDataIncrementalRefreshEntries) MdEntryTimeSinceVersion() uint16 {
	return 0
}

func (m *MarketDataIncrementalRefreshEntries) MdEntryTimeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.MdEntryTimeSinceVersion()
}

func (*MarketDataIncrementalRefreshEntries) MdEntryTimeDeprecated() uint16 {
	return 0
}

func (*MarketDataIncrementalRefreshEntries) MdEntryTimeMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return "nanosecond"
	case 3:
		return ""
	case 4:
		return "required"
	}
	return ""
}

func (*MarketDataIncrementalRefreshEntries) MdEntryTimeMinValue() uint64 {
	return 0
}

func (*MarketDataIncrementalRefreshEntries) MdEntryTimeMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*MarketDataIncrementalRefreshEntries) MdEntryTimeNullValue() uint64 {
	return math.MaxUint64
}

func (*MarketDataIncrementalRefreshEntries) MdEntrySizeId() uint16 {
	return 271
}

func (*MarketDataIncrementalRefreshEntries) MdEntrySizeSinceVersion() uint16 {
	return 0
}

func (m *MarketDataIncrementalRefreshEntries) MdEntrySizeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.MdEntrySizeSinceVersion()
}

func (*MarketDataIncrementalRefreshEntries) MdEntrySizeDeprecated() uint16 {
	return 0
}

func (*MarketDataIncrementalRefreshEntries) MdEntrySizeMetaAttribute(meta int) string {
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

func (*MarketDataIncrementalRefreshEntries) TradingSessionIdId() uint16 {
	return 336
}

func (*MarketDataIncrementalRefreshEntries) TradingSessionIdSinceVersion() uint16 {
	return 0
}

func (m *MarketDataIncrementalRefreshEntries) TradingSessionIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.TradingSessionIdSinceVersion()
}

func (*MarketDataIncrementalRefreshEntries) TradingSessionIdDeprecated() uint16 {
	return 0
}

func (*MarketDataIncrementalRefreshEntries) TradingSessionIdMetaAttribute(meta int) string {
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

func (*MarketDataIncrementalRefreshEntries) NetChgPrevDayId() uint16 {
	return 451
}

func (*MarketDataIncrementalRefreshEntries) NetChgPrevDaySinceVersion() uint16 {
	return 0
}

func (m *MarketDataIncrementalRefreshEntries) NetChgPrevDayInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.NetChgPrevDaySinceVersion()
}

func (*MarketDataIncrementalRefreshEntries) NetChgPrevDayDeprecated() uint16 {
	return 0
}

func (*MarketDataIncrementalRefreshEntries) NetChgPrevDayMetaAttribute(meta int) string {
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

func (*MarketDataIncrementalRefreshEntries) TickDirectionId() uint16 {
	return 274
}

func (*MarketDataIncrementalRefreshEntries) TickDirectionSinceVersion() uint16 {
	return 0
}

func (m *MarketDataIncrementalRefreshEntries) TickDirectionInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.TickDirectionSinceVersion()
}

func (*MarketDataIncrementalRefreshEntries) TickDirectionDeprecated() uint16 {
	return 0
}

func (*MarketDataIncrementalRefreshEntries) TickDirectionMetaAttribute(meta int) string {
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

func (*MarketDataIncrementalRefreshEntries) OpenCloseSettleFlagId() uint16 {
	return 286
}

func (*MarketDataIncrementalRefreshEntries) OpenCloseSettleFlagSinceVersion() uint16 {
	return 0
}

func (m *MarketDataIncrementalRefreshEntries) OpenCloseSettleFlagInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.OpenCloseSettleFlagSinceVersion()
}

func (*MarketDataIncrementalRefreshEntries) OpenCloseSettleFlagDeprecated() uint16 {
	return 0
}

func (*MarketDataIncrementalRefreshEntries) OpenCloseSettleFlagMetaAttribute(meta int) string {
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

func (*MarketDataIncrementalRefreshEntries) SettleDateId() uint16 {
	return 64
}

func (*MarketDataIncrementalRefreshEntries) SettleDateSinceVersion() uint16 {
	return 0
}

func (m *MarketDataIncrementalRefreshEntries) SettleDateInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.SettleDateSinceVersion()
}

func (*MarketDataIncrementalRefreshEntries) SettleDateDeprecated() uint16 {
	return 0
}

func (*MarketDataIncrementalRefreshEntries) SettleDateMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return "nanosecond"
	case 3:
		return ""
	case 4:
		return "required"
	}
	return ""
}

func (*MarketDataIncrementalRefreshEntries) SettleDateMinValue() uint64 {
	return 0
}

func (*MarketDataIncrementalRefreshEntries) SettleDateMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*MarketDataIncrementalRefreshEntries) SettleDateNullValue() uint64 {
	return math.MaxUint64
}

func (*MarketDataIncrementalRefreshEntries) TradeConditionId() uint16 {
	return 277
}

func (*MarketDataIncrementalRefreshEntries) TradeConditionSinceVersion() uint16 {
	return 0
}

func (m *MarketDataIncrementalRefreshEntries) TradeConditionInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.TradeConditionSinceVersion()
}

func (*MarketDataIncrementalRefreshEntries) TradeConditionDeprecated() uint16 {
	return 0
}

func (*MarketDataIncrementalRefreshEntries) TradeConditionMetaAttribute(meta int) string {
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

func (*MarketDataIncrementalRefreshEntries) TradeVolumeId() uint16 {
	return 1020
}

func (*MarketDataIncrementalRefreshEntries) TradeVolumeSinceVersion() uint16 {
	return 0
}

func (m *MarketDataIncrementalRefreshEntries) TradeVolumeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.TradeVolumeSinceVersion()
}

func (*MarketDataIncrementalRefreshEntries) TradeVolumeDeprecated() uint16 {
	return 0
}

func (*MarketDataIncrementalRefreshEntries) TradeVolumeMetaAttribute(meta int) string {
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

func (*MarketDataIncrementalRefreshEntries) MdQuoteTypeId() uint16 {
	return 1070
}

func (*MarketDataIncrementalRefreshEntries) MdQuoteTypeSinceVersion() uint16 {
	return 0
}

func (m *MarketDataIncrementalRefreshEntries) MdQuoteTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.MdQuoteTypeSinceVersion()
}

func (*MarketDataIncrementalRefreshEntries) MdQuoteTypeDeprecated() uint16 {
	return 0
}

func (*MarketDataIncrementalRefreshEntries) MdQuoteTypeMetaAttribute(meta int) string {
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

func (*MarketDataIncrementalRefreshEntries) FixingBracketId() uint16 {
	return 5790
}

func (*MarketDataIncrementalRefreshEntries) FixingBracketSinceVersion() uint16 {
	return 0
}

func (m *MarketDataIncrementalRefreshEntries) FixingBracketInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.FixingBracketSinceVersion()
}

func (*MarketDataIncrementalRefreshEntries) FixingBracketDeprecated() uint16 {
	return 0
}

func (*MarketDataIncrementalRefreshEntries) FixingBracketMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return "nanosecond"
	case 3:
		return ""
	case 4:
		return "required"
	}
	return ""
}

func (*MarketDataIncrementalRefreshEntries) FixingBracketMinValue() uint64 {
	return 0
}

func (*MarketDataIncrementalRefreshEntries) FixingBracketMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*MarketDataIncrementalRefreshEntries) FixingBracketNullValue() uint64 {
	return math.MaxUint64
}

func (*MarketDataIncrementalRefreshEntries) AggressorSideId() uint16 {
	return 5797
}

func (*MarketDataIncrementalRefreshEntries) AggressorSideSinceVersion() uint16 {
	return 0
}

func (m *MarketDataIncrementalRefreshEntries) AggressorSideInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.AggressorSideSinceVersion()
}

func (*MarketDataIncrementalRefreshEntries) AggressorSideDeprecated() uint16 {
	return 0
}

func (*MarketDataIncrementalRefreshEntries) AggressorSideMetaAttribute(meta int) string {
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

func (*MarketDataIncrementalRefreshEntries) MatchEventIndicatorId() uint16 {
	return 5799
}

func (*MarketDataIncrementalRefreshEntries) MatchEventIndicatorSinceVersion() uint16 {
	return 0
}

func (m *MarketDataIncrementalRefreshEntries) MatchEventIndicatorInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.MatchEventIndicatorSinceVersion()
}

func (*MarketDataIncrementalRefreshEntries) MatchEventIndicatorDeprecated() uint16 {
	return 0
}

func (*MarketDataIncrementalRefreshEntries) MatchEventIndicatorMetaAttribute(meta int) string {
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

func (*MarketDataIncrementalRefreshEntries) TradeIdId() uint16 {
	return 1003
}

func (*MarketDataIncrementalRefreshEntries) TradeIdSinceVersion() uint16 {
	return 0
}

func (m *MarketDataIncrementalRefreshEntries) TradeIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.TradeIdSinceVersion()
}

func (*MarketDataIncrementalRefreshEntries) TradeIdDeprecated() uint16 {
	return 0
}

func (*MarketDataIncrementalRefreshEntries) TradeIdMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return "ExecID"
	case 4:
		return "required"
	}
	return ""
}

func (*MarketDataIncrementalRefreshEntries) TradeIdMinValue() uint64 {
	return 0
}

func (*MarketDataIncrementalRefreshEntries) TradeIdMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*MarketDataIncrementalRefreshEntries) TradeIdNullValue() uint64 {
	return math.MaxUint64
}

func (*MarketDataIncrementalRefresh) EntriesId() uint16 {
	return 268
}

func (*MarketDataIncrementalRefresh) EntriesSinceVersion() uint16 {
	return 0
}

func (m *MarketDataIncrementalRefresh) EntriesInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.EntriesSinceVersion()
}

func (*MarketDataIncrementalRefresh) EntriesDeprecated() uint16 {
	return 0
}

func (*MarketDataIncrementalRefreshEntries) SbeBlockLength() (blockLength uint) {
	return 82
}

func (*MarketDataIncrementalRefreshEntries) SbeSchemaVersion() (schemaVersion uint16) {
	return 1
}
