// Generated SBE (Simple Binary Encoding) message codec

package fix

import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
)

type MarketDataIncrementalRefreshTrades struct {
	TransactTime        uint64
	EventTimeDelta      uint16
	MatchEventIndicator MatchEventIndicatorEnum
	MdIncGrp            []MarketDataIncrementalRefreshTradesMdIncGrp
}
type MarketDataIncrementalRefreshTradesMdIncGrp struct {
	TradeId        uint64
	SecurityId     uint64
	MdEntryPx      Decimal64
	MdEntrySize    IntQty32
	NumberOfOrders uint16
	MdUpdateAction MDUpdateActionEnum
	RptSeq         uint8
	AggressorSide  SideEnum
	MdEntryType    MDEntryTypeEnum
}

func (m *MarketDataIncrementalRefreshTrades) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := m.RangeCheck(m.SbeSchemaVersion(), m.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	if err := _m.WriteUint64(_w, m.TransactTime); err != nil {
		return err
	}
	if err := _m.WriteUint16(_w, m.EventTimeDelta); err != nil {
		return err
	}
	if err := m.MatchEventIndicator.Encode(_m, _w); err != nil {
		return err
	}
	var MdIncGrpBlockLength uint16 = 33
	if err := _m.WriteUint16(_w, MdIncGrpBlockLength); err != nil {
		return err
	}
	var MdIncGrpNumInGroup uint16 = uint16(len(m.MdIncGrp))
	if err := _m.WriteUint16(_w, MdIncGrpNumInGroup); err != nil {
		return err
	}
	for _, prop := range m.MdIncGrp {
		if err := prop.Encode(_m, _w); err != nil {
			return err
		}
	}
	return nil
}

func (m *MarketDataIncrementalRefreshTrades) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	if !m.TransactTimeInActingVersion(actingVersion) {
		m.TransactTime = m.TransactTimeNullValue()
	} else {
		if err := _m.ReadUint64(_r, &m.TransactTime); err != nil {
			return err
		}
	}
	if !m.EventTimeDeltaInActingVersion(actingVersion) {
		m.EventTimeDelta = m.EventTimeDeltaNullValue()
	} else {
		if err := _m.ReadUint16(_r, &m.EventTimeDelta); err != nil {
			return err
		}
	}
	if m.MatchEventIndicatorInActingVersion(actingVersion) {
		if err := m.MatchEventIndicator.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if actingVersion > m.SbeSchemaVersion() && blockLength > m.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-m.SbeBlockLength()))
	}

	if m.MdIncGrpInActingVersion(actingVersion) {
		var MdIncGrpBlockLength uint16
		if err := _m.ReadUint16(_r, &MdIncGrpBlockLength); err != nil {
			return err
		}
		var MdIncGrpNumInGroup uint16
		if err := _m.ReadUint16(_r, &MdIncGrpNumInGroup); err != nil {
			return err
		}
		if cap(m.MdIncGrp) < int(MdIncGrpNumInGroup) {
			m.MdIncGrp = make([]MarketDataIncrementalRefreshTradesMdIncGrp, MdIncGrpNumInGroup)
		}
		for i, _ := range m.MdIncGrp {
			if err := m.MdIncGrp[i].Decode(_m, _r, actingVersion, uint(MdIncGrpBlockLength)); err != nil {
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

func (m *MarketDataIncrementalRefreshTrades) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if m.TransactTimeInActingVersion(actingVersion) {
		if m.TransactTime < m.TransactTimeMinValue() || m.TransactTime > m.TransactTimeMaxValue() {
			return fmt.Errorf("Range check failed on m.TransactTime (%v < %v > %v)", m.TransactTimeMinValue(), m.TransactTime, m.TransactTimeMaxValue())
		}
	}
	if m.EventTimeDeltaInActingVersion(actingVersion) {
		if m.EventTimeDelta < m.EventTimeDeltaMinValue() || m.EventTimeDelta > m.EventTimeDeltaMaxValue() {
			return fmt.Errorf("Range check failed on m.EventTimeDelta (%v < %v > %v)", m.EventTimeDeltaMinValue(), m.EventTimeDelta, m.EventTimeDeltaMaxValue())
		}
	}
	if err := m.MatchEventIndicator.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	for _, prop := range m.MdIncGrp {
		if err := prop.RangeCheck(actingVersion, schemaVersion); err != nil {
			return err
		}
	}
	return nil
}

func MarketDataIncrementalRefreshTradesInit(m *MarketDataIncrementalRefreshTrades) {
	return
}

func (m *MarketDataIncrementalRefreshTradesMdIncGrp) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint64(_w, m.TradeId); err != nil {
		return err
	}
	if err := _m.WriteUint64(_w, m.SecurityId); err != nil {
		return err
	}
	if err := m.MdEntryPx.Encode(_m, _w); err != nil {
		return err
	}
	if err := m.MdEntrySize.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteUint16(_w, m.NumberOfOrders); err != nil {
		return err
	}
	if err := m.MdUpdateAction.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteUint8(_w, m.RptSeq); err != nil {
		return err
	}
	if err := m.AggressorSide.Encode(_m, _w); err != nil {
		return err
	}
	return nil
}

func (m *MarketDataIncrementalRefreshTradesMdIncGrp) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint) error {
	if !m.TradeIdInActingVersion(actingVersion) {
		m.TradeId = m.TradeIdNullValue()
	} else {
		if err := _m.ReadUint64(_r, &m.TradeId); err != nil {
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
	if m.MdEntryPxInActingVersion(actingVersion) {
		if err := m.MdEntryPx.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if m.MdEntrySizeInActingVersion(actingVersion) {
		if err := m.MdEntrySize.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if !m.NumberOfOrdersInActingVersion(actingVersion) {
		m.NumberOfOrders = m.NumberOfOrdersNullValue()
	} else {
		if err := _m.ReadUint16(_r, &m.NumberOfOrders); err != nil {
			return err
		}
	}
	if m.MdUpdateActionInActingVersion(actingVersion) {
		if err := m.MdUpdateAction.Decode(_m, _r, actingVersion); err != nil {
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
	if m.AggressorSideInActingVersion(actingVersion) {
		if err := m.AggressorSide.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	m.MdEntryType = MDEntryType.TRADE
	if actingVersion > m.SbeSchemaVersion() && blockLength > m.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-m.SbeBlockLength()))
	}
	return nil
}

func (m *MarketDataIncrementalRefreshTradesMdIncGrp) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if m.TradeIdInActingVersion(actingVersion) {
		if m.TradeId < m.TradeIdMinValue() || m.TradeId > m.TradeIdMaxValue() {
			return fmt.Errorf("Range check failed on m.TradeId (%v < %v > %v)", m.TradeIdMinValue(), m.TradeId, m.TradeIdMaxValue())
		}
	}
	if m.SecurityIdInActingVersion(actingVersion) {
		if m.SecurityId < m.SecurityIdMinValue() || m.SecurityId > m.SecurityIdMaxValue() {
			return fmt.Errorf("Range check failed on m.SecurityId (%v < %v > %v)", m.SecurityIdMinValue(), m.SecurityId, m.SecurityIdMaxValue())
		}
	}
	if m.NumberOfOrdersInActingVersion(actingVersion) {
		if m.NumberOfOrders < m.NumberOfOrdersMinValue() || m.NumberOfOrders > m.NumberOfOrdersMaxValue() {
			return fmt.Errorf("Range check failed on m.NumberOfOrders (%v < %v > %v)", m.NumberOfOrdersMinValue(), m.NumberOfOrders, m.NumberOfOrdersMaxValue())
		}
	}
	if err := m.MdUpdateAction.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if m.RptSeqInActingVersion(actingVersion) {
		if m.RptSeq < m.RptSeqMinValue() || m.RptSeq > m.RptSeqMaxValue() {
			return fmt.Errorf("Range check failed on m.RptSeq (%v < %v > %v)", m.RptSeqMinValue(), m.RptSeq, m.RptSeqMaxValue())
		}
	}
	if err := m.AggressorSide.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if err := m.MdEntryType.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	return nil
}

func MarketDataIncrementalRefreshTradesMdIncGrpInit(m *MarketDataIncrementalRefreshTradesMdIncGrp) {
	m.MdEntryType = MDEntryType.TRADE
	return
}

func (*MarketDataIncrementalRefreshTrades) SbeBlockLength() (blockLength uint16) {
	return 11
}

func (*MarketDataIncrementalRefreshTrades) SbeTemplateId() (templateId uint16) {
	return 2
}

func (*MarketDataIncrementalRefreshTrades) SbeSchemaId() (schemaId uint16) {
	return 1
}

func (*MarketDataIncrementalRefreshTrades) SbeSchemaVersion() (schemaVersion uint16) {
	return 1
}

func (*MarketDataIncrementalRefreshTrades) SbeSemanticType() (semanticType []byte) {
	return []byte("X")
}

func (*MarketDataIncrementalRefreshTrades) TransactTimeId() uint16 {
	return 60
}

func (*MarketDataIncrementalRefreshTrades) TransactTimeSinceVersion() uint16 {
	return 0
}

func (m *MarketDataIncrementalRefreshTrades) TransactTimeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.TransactTimeSinceVersion()
}

func (*MarketDataIncrementalRefreshTrades) TransactTimeDeprecated() uint16 {
	return 0
}

func (*MarketDataIncrementalRefreshTrades) TransactTimeMetaAttribute(meta int) string {
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

func (*MarketDataIncrementalRefreshTrades) TransactTimeMinValue() uint64 {
	return 0
}

func (*MarketDataIncrementalRefreshTrades) TransactTimeMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*MarketDataIncrementalRefreshTrades) TransactTimeNullValue() uint64 {
	return math.MaxUint64
}

func (*MarketDataIncrementalRefreshTrades) EventTimeDeltaId() uint16 {
	return 37704
}

func (*MarketDataIncrementalRefreshTrades) EventTimeDeltaSinceVersion() uint16 {
	return 0
}

func (m *MarketDataIncrementalRefreshTrades) EventTimeDeltaInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.EventTimeDeltaSinceVersion()
}

func (*MarketDataIncrementalRefreshTrades) EventTimeDeltaDeprecated() uint16 {
	return 0
}

func (*MarketDataIncrementalRefreshTrades) EventTimeDeltaMetaAttribute(meta int) string {
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

func (*MarketDataIncrementalRefreshTrades) EventTimeDeltaMinValue() uint16 {
	return 0
}

func (*MarketDataIncrementalRefreshTrades) EventTimeDeltaMaxValue() uint16 {
	return math.MaxUint16 - 1
}

func (*MarketDataIncrementalRefreshTrades) EventTimeDeltaNullValue() uint16 {
	return math.MaxUint16
}

func (*MarketDataIncrementalRefreshTrades) MatchEventIndicatorId() uint16 {
	return 5799
}

func (*MarketDataIncrementalRefreshTrades) MatchEventIndicatorSinceVersion() uint16 {
	return 0
}

func (m *MarketDataIncrementalRefreshTrades) MatchEventIndicatorInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.MatchEventIndicatorSinceVersion()
}

func (*MarketDataIncrementalRefreshTrades) MatchEventIndicatorDeprecated() uint16 {
	return 0
}

func (*MarketDataIncrementalRefreshTrades) MatchEventIndicatorMetaAttribute(meta int) string {
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

func (*MarketDataIncrementalRefreshTradesMdIncGrp) TradeIdId() uint16 {
	return 1003
}

func (*MarketDataIncrementalRefreshTradesMdIncGrp) TradeIdSinceVersion() uint16 {
	return 0
}

func (m *MarketDataIncrementalRefreshTradesMdIncGrp) TradeIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.TradeIdSinceVersion()
}

func (*MarketDataIncrementalRefreshTradesMdIncGrp) TradeIdDeprecated() uint16 {
	return 0
}

func (*MarketDataIncrementalRefreshTradesMdIncGrp) TradeIdMetaAttribute(meta int) string {
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

func (*MarketDataIncrementalRefreshTradesMdIncGrp) TradeIdMinValue() uint64 {
	return 0
}

func (*MarketDataIncrementalRefreshTradesMdIncGrp) TradeIdMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*MarketDataIncrementalRefreshTradesMdIncGrp) TradeIdNullValue() uint64 {
	return math.MaxUint64
}

func (*MarketDataIncrementalRefreshTradesMdIncGrp) SecurityIdId() uint16 {
	return 48
}

func (*MarketDataIncrementalRefreshTradesMdIncGrp) SecurityIdSinceVersion() uint16 {
	return 0
}

func (m *MarketDataIncrementalRefreshTradesMdIncGrp) SecurityIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.SecurityIdSinceVersion()
}

func (*MarketDataIncrementalRefreshTradesMdIncGrp) SecurityIdDeprecated() uint16 {
	return 0
}

func (*MarketDataIncrementalRefreshTradesMdIncGrp) SecurityIdMetaAttribute(meta int) string {
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

func (*MarketDataIncrementalRefreshTradesMdIncGrp) SecurityIdMinValue() uint64 {
	return 0
}

func (*MarketDataIncrementalRefreshTradesMdIncGrp) SecurityIdMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*MarketDataIncrementalRefreshTradesMdIncGrp) SecurityIdNullValue() uint64 {
	return math.MaxUint64
}

func (*MarketDataIncrementalRefreshTradesMdIncGrp) MdEntryPxId() uint16 {
	return 270
}

func (*MarketDataIncrementalRefreshTradesMdIncGrp) MdEntryPxSinceVersion() uint16 {
	return 0
}

func (m *MarketDataIncrementalRefreshTradesMdIncGrp) MdEntryPxInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.MdEntryPxSinceVersion()
}

func (*MarketDataIncrementalRefreshTradesMdIncGrp) MdEntryPxDeprecated() uint16 {
	return 0
}

func (*MarketDataIncrementalRefreshTradesMdIncGrp) MdEntryPxMetaAttribute(meta int) string {
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

func (*MarketDataIncrementalRefreshTradesMdIncGrp) MdEntrySizeId() uint16 {
	return 271
}

func (*MarketDataIncrementalRefreshTradesMdIncGrp) MdEntrySizeSinceVersion() uint16 {
	return 0
}

func (m *MarketDataIncrementalRefreshTradesMdIncGrp) MdEntrySizeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.MdEntrySizeSinceVersion()
}

func (*MarketDataIncrementalRefreshTradesMdIncGrp) MdEntrySizeDeprecated() uint16 {
	return 0
}

func (*MarketDataIncrementalRefreshTradesMdIncGrp) MdEntrySizeMetaAttribute(meta int) string {
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

func (*MarketDataIncrementalRefreshTradesMdIncGrp) NumberOfOrdersId() uint16 {
	return 346
}

func (*MarketDataIncrementalRefreshTradesMdIncGrp) NumberOfOrdersSinceVersion() uint16 {
	return 0
}

func (m *MarketDataIncrementalRefreshTradesMdIncGrp) NumberOfOrdersInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.NumberOfOrdersSinceVersion()
}

func (*MarketDataIncrementalRefreshTradesMdIncGrp) NumberOfOrdersDeprecated() uint16 {
	return 0
}

func (*MarketDataIncrementalRefreshTradesMdIncGrp) NumberOfOrdersMetaAttribute(meta int) string {
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

func (*MarketDataIncrementalRefreshTradesMdIncGrp) NumberOfOrdersMinValue() uint16 {
	return 0
}

func (*MarketDataIncrementalRefreshTradesMdIncGrp) NumberOfOrdersMaxValue() uint16 {
	return math.MaxUint16 - 1
}

func (*MarketDataIncrementalRefreshTradesMdIncGrp) NumberOfOrdersNullValue() uint16 {
	return math.MaxUint16
}

func (*MarketDataIncrementalRefreshTradesMdIncGrp) MdUpdateActionId() uint16 {
	return 279
}

func (*MarketDataIncrementalRefreshTradesMdIncGrp) MdUpdateActionSinceVersion() uint16 {
	return 0
}

func (m *MarketDataIncrementalRefreshTradesMdIncGrp) MdUpdateActionInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.MdUpdateActionSinceVersion()
}

func (*MarketDataIncrementalRefreshTradesMdIncGrp) MdUpdateActionDeprecated() uint16 {
	return 0
}

func (*MarketDataIncrementalRefreshTradesMdIncGrp) MdUpdateActionMetaAttribute(meta int) string {
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

func (*MarketDataIncrementalRefreshTradesMdIncGrp) RptSeqId() uint16 {
	return 83
}

func (*MarketDataIncrementalRefreshTradesMdIncGrp) RptSeqSinceVersion() uint16 {
	return 0
}

func (m *MarketDataIncrementalRefreshTradesMdIncGrp) RptSeqInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.RptSeqSinceVersion()
}

func (*MarketDataIncrementalRefreshTradesMdIncGrp) RptSeqDeprecated() uint16 {
	return 0
}

func (*MarketDataIncrementalRefreshTradesMdIncGrp) RptSeqMetaAttribute(meta int) string {
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

func (*MarketDataIncrementalRefreshTradesMdIncGrp) RptSeqMinValue() uint8 {
	return 0
}

func (*MarketDataIncrementalRefreshTradesMdIncGrp) RptSeqMaxValue() uint8 {
	return math.MaxUint8 - 1
}

func (*MarketDataIncrementalRefreshTradesMdIncGrp) RptSeqNullValue() uint8 {
	return math.MaxUint8
}

func (*MarketDataIncrementalRefreshTradesMdIncGrp) AggressorSideId() uint16 {
	return 5797
}

func (*MarketDataIncrementalRefreshTradesMdIncGrp) AggressorSideSinceVersion() uint16 {
	return 0
}

func (m *MarketDataIncrementalRefreshTradesMdIncGrp) AggressorSideInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.AggressorSideSinceVersion()
}

func (*MarketDataIncrementalRefreshTradesMdIncGrp) AggressorSideDeprecated() uint16 {
	return 0
}

func (*MarketDataIncrementalRefreshTradesMdIncGrp) AggressorSideMetaAttribute(meta int) string {
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

func (*MarketDataIncrementalRefreshTradesMdIncGrp) MdEntryTypeId() uint16 {
	return 269
}

func (*MarketDataIncrementalRefreshTradesMdIncGrp) MdEntryTypeSinceVersion() uint16 {
	return 0
}

func (m *MarketDataIncrementalRefreshTradesMdIncGrp) MdEntryTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.MdEntryTypeSinceVersion()
}

func (*MarketDataIncrementalRefreshTradesMdIncGrp) MdEntryTypeDeprecated() uint16 {
	return 0
}

func (*MarketDataIncrementalRefreshTradesMdIncGrp) MdEntryTypeMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "constant"
	}
	return ""
}

func (*MarketDataIncrementalRefreshTrades) MdIncGrpId() uint16 {
	return 268
}

func (*MarketDataIncrementalRefreshTrades) MdIncGrpSinceVersion() uint16 {
	return 0
}

func (m *MarketDataIncrementalRefreshTrades) MdIncGrpInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.MdIncGrpSinceVersion()
}

func (*MarketDataIncrementalRefreshTrades) MdIncGrpDeprecated() uint16 {
	return 0
}

func (*MarketDataIncrementalRefreshTradesMdIncGrp) SbeBlockLength() (blockLength uint) {
	return 33
}

func (*MarketDataIncrementalRefreshTradesMdIncGrp) SbeSchemaVersion() (schemaVersion uint16) {
	return 1
}
