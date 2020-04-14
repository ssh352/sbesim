// Generated SBE (Simple Binary Encoding) message codec

package ilinkbinary

import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
)

type MassQuoteAck545 struct {
	SeqNum                uint32
	UUID                  uint64
	Text                  [256]byte
	SenderID              [20]byte
	PartyDetailsListReqID uint64
	RequestTime           uint64
	SendingTimeEpoch      uint64
	QuoteReqID            uint64
	Location              [5]byte
	QuoteID               uint32
	QuoteRejectReason     uint16
	DelayDuration         uint16
	QuoteStatus           QuoteAckStatusEnum
	ManualOrderIndicator  ManualOrdIndReqEnum
	NoProcessedEntries    uint8
	MMProtectionReset     BooleanFlagEnum
	SplitMsg              SplitMsgEnum
	LiquidityFlag         BooleanNULLEnum
	ShortSaleType         ShortSaleTypeEnum
	TotNoQuoteEntries     uint8
	PossRetransFlag       BooleanFlagEnum
	DelayToTime           uint64
	NoQuoteEntries        []MassQuoteAck545NoQuoteEntries
}
type MassQuoteAck545NoQuoteEntries struct {
	QuoteEntryID           uint32
	SecurityID             int32
	QuoteSetID             uint16
	QuoteEntryRejectReason uint8
}

func (m *MassQuoteAck545) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := m.RangeCheck(m.SbeSchemaVersion(), m.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	if err := _m.WriteUint32(_w, m.SeqNum); err != nil {
		return err
	}
	if err := _m.WriteUint64(_w, m.UUID); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, m.Text[:]); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, m.SenderID[:]); err != nil {
		return err
	}
	if err := _m.WriteUint64(_w, m.PartyDetailsListReqID); err != nil {
		return err
	}
	if err := _m.WriteUint64(_w, m.RequestTime); err != nil {
		return err
	}
	if err := _m.WriteUint64(_w, m.SendingTimeEpoch); err != nil {
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
	if err := _m.WriteUint16(_w, m.QuoteRejectReason); err != nil {
		return err
	}
	if err := _m.WriteUint16(_w, m.DelayDuration); err != nil {
		return err
	}
	if err := m.QuoteStatus.Encode(_m, _w); err != nil {
		return err
	}
	if err := m.ManualOrderIndicator.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteUint8(_w, m.NoProcessedEntries); err != nil {
		return err
	}
	if err := m.MMProtectionReset.Encode(_m, _w); err != nil {
		return err
	}
	if err := m.SplitMsg.Encode(_m, _w); err != nil {
		return err
	}
	if err := m.LiquidityFlag.Encode(_m, _w); err != nil {
		return err
	}
	if err := m.ShortSaleType.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteUint8(_w, m.TotNoQuoteEntries); err != nil {
		return err
	}
	if err := m.PossRetransFlag.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteUint64(_w, m.DelayToTime); err != nil {
		return err
	}
	var NoQuoteEntriesBlockLength uint16 = 11
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

func (m *MassQuoteAck545) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	if !m.SeqNumInActingVersion(actingVersion) {
		m.SeqNum = m.SeqNumNullValue()
	} else {
		if err := _m.ReadUint32(_r, &m.SeqNum); err != nil {
			return err
		}
	}
	if !m.UUIDInActingVersion(actingVersion) {
		m.UUID = m.UUIDNullValue()
	} else {
		if err := _m.ReadUint64(_r, &m.UUID); err != nil {
			return err
		}
	}
	if !m.TextInActingVersion(actingVersion) {
		for idx := 0; idx < 256; idx++ {
			m.Text[idx] = m.TextNullValue()
		}
	} else {
		if err := _m.ReadBytes(_r, m.Text[:]); err != nil {
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
	if !m.PartyDetailsListReqIDInActingVersion(actingVersion) {
		m.PartyDetailsListReqID = m.PartyDetailsListReqIDNullValue()
	} else {
		if err := _m.ReadUint64(_r, &m.PartyDetailsListReqID); err != nil {
			return err
		}
	}
	if !m.RequestTimeInActingVersion(actingVersion) {
		m.RequestTime = m.RequestTimeNullValue()
	} else {
		if err := _m.ReadUint64(_r, &m.RequestTime); err != nil {
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
	if !m.QuoteRejectReasonInActingVersion(actingVersion) {
		m.QuoteRejectReason = m.QuoteRejectReasonNullValue()
	} else {
		if err := _m.ReadUint16(_r, &m.QuoteRejectReason); err != nil {
			return err
		}
	}
	if !m.DelayDurationInActingVersion(actingVersion) {
		m.DelayDuration = m.DelayDurationNullValue()
	} else {
		if err := _m.ReadUint16(_r, &m.DelayDuration); err != nil {
			return err
		}
	}
	if m.QuoteStatusInActingVersion(actingVersion) {
		if err := m.QuoteStatus.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if m.ManualOrderIndicatorInActingVersion(actingVersion) {
		if err := m.ManualOrderIndicator.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if !m.NoProcessedEntriesInActingVersion(actingVersion) {
		m.NoProcessedEntries = m.NoProcessedEntriesNullValue()
	} else {
		if err := _m.ReadUint8(_r, &m.NoProcessedEntries); err != nil {
			return err
		}
	}
	if m.MMProtectionResetInActingVersion(actingVersion) {
		if err := m.MMProtectionReset.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if m.SplitMsgInActingVersion(actingVersion) {
		if err := m.SplitMsg.Decode(_m, _r, actingVersion); err != nil {
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
	if !m.TotNoQuoteEntriesInActingVersion(actingVersion) {
		m.TotNoQuoteEntries = m.TotNoQuoteEntriesNullValue()
	} else {
		if err := _m.ReadUint8(_r, &m.TotNoQuoteEntries); err != nil {
			return err
		}
	}
	if m.PossRetransFlagInActingVersion(actingVersion) {
		if err := m.PossRetransFlag.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if !m.DelayToTimeInActingVersion(actingVersion) {
		m.DelayToTime = m.DelayToTimeNullValue()
	} else {
		if err := _m.ReadUint64(_r, &m.DelayToTime); err != nil {
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
			m.NoQuoteEntries = make([]MassQuoteAck545NoQuoteEntries, NoQuoteEntriesNumInGroup)
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

func (m *MassQuoteAck545) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if m.SeqNumInActingVersion(actingVersion) {
		if m.SeqNum < m.SeqNumMinValue() || m.SeqNum > m.SeqNumMaxValue() {
			return fmt.Errorf("Range check failed on m.SeqNum (%v < %v > %v)", m.SeqNumMinValue(), m.SeqNum, m.SeqNumMaxValue())
		}
	}
	if m.UUIDInActingVersion(actingVersion) {
		if m.UUID < m.UUIDMinValue() || m.UUID > m.UUIDMaxValue() {
			return fmt.Errorf("Range check failed on m.UUID (%v < %v > %v)", m.UUIDMinValue(), m.UUID, m.UUIDMaxValue())
		}
	}
	if m.TextInActingVersion(actingVersion) {
		for idx := 0; idx < 256; idx++ {
			if m.Text[idx] < m.TextMinValue() || m.Text[idx] > m.TextMaxValue() {
				return fmt.Errorf("Range check failed on m.Text[%d] (%v < %v > %v)", idx, m.TextMinValue(), m.Text[idx], m.TextMaxValue())
			}
		}
	}
	if m.SenderIDInActingVersion(actingVersion) {
		for idx := 0; idx < 20; idx++ {
			if m.SenderID[idx] < m.SenderIDMinValue() || m.SenderID[idx] > m.SenderIDMaxValue() {
				return fmt.Errorf("Range check failed on m.SenderID[%d] (%v < %v > %v)", idx, m.SenderIDMinValue(), m.SenderID[idx], m.SenderIDMaxValue())
			}
		}
	}
	if m.PartyDetailsListReqIDInActingVersion(actingVersion) {
		if m.PartyDetailsListReqID < m.PartyDetailsListReqIDMinValue() || m.PartyDetailsListReqID > m.PartyDetailsListReqIDMaxValue() {
			return fmt.Errorf("Range check failed on m.PartyDetailsListReqID (%v < %v > %v)", m.PartyDetailsListReqIDMinValue(), m.PartyDetailsListReqID, m.PartyDetailsListReqIDMaxValue())
		}
	}
	if m.RequestTimeInActingVersion(actingVersion) {
		if m.RequestTime < m.RequestTimeMinValue() || m.RequestTime > m.RequestTimeMaxValue() {
			return fmt.Errorf("Range check failed on m.RequestTime (%v < %v > %v)", m.RequestTimeMinValue(), m.RequestTime, m.RequestTimeMaxValue())
		}
	}
	if m.SendingTimeEpochInActingVersion(actingVersion) {
		if m.SendingTimeEpoch < m.SendingTimeEpochMinValue() || m.SendingTimeEpoch > m.SendingTimeEpochMaxValue() {
			return fmt.Errorf("Range check failed on m.SendingTimeEpoch (%v < %v > %v)", m.SendingTimeEpochMinValue(), m.SendingTimeEpoch, m.SendingTimeEpochMaxValue())
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
	if m.QuoteRejectReasonInActingVersion(actingVersion) {
		if m.QuoteRejectReason != m.QuoteRejectReasonNullValue() && (m.QuoteRejectReason < m.QuoteRejectReasonMinValue() || m.QuoteRejectReason > m.QuoteRejectReasonMaxValue()) {
			return fmt.Errorf("Range check failed on m.QuoteRejectReason (%v < %v > %v)", m.QuoteRejectReasonMinValue(), m.QuoteRejectReason, m.QuoteRejectReasonMaxValue())
		}
	}
	if m.DelayDurationInActingVersion(actingVersion) {
		if m.DelayDuration != m.DelayDurationNullValue() && (m.DelayDuration < m.DelayDurationMinValue() || m.DelayDuration > m.DelayDurationMaxValue()) {
			return fmt.Errorf("Range check failed on m.DelayDuration (%v < %v > %v)", m.DelayDurationMinValue(), m.DelayDuration, m.DelayDurationMaxValue())
		}
	}
	if err := m.QuoteStatus.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if err := m.ManualOrderIndicator.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if m.NoProcessedEntriesInActingVersion(actingVersion) {
		if m.NoProcessedEntries < m.NoProcessedEntriesMinValue() || m.NoProcessedEntries > m.NoProcessedEntriesMaxValue() {
			return fmt.Errorf("Range check failed on m.NoProcessedEntries (%v < %v > %v)", m.NoProcessedEntriesMinValue(), m.NoProcessedEntries, m.NoProcessedEntriesMaxValue())
		}
	}
	if err := m.MMProtectionReset.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if err := m.SplitMsg.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if err := m.LiquidityFlag.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if err := m.ShortSaleType.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if m.TotNoQuoteEntriesInActingVersion(actingVersion) {
		if m.TotNoQuoteEntries != m.TotNoQuoteEntriesNullValue() && (m.TotNoQuoteEntries < m.TotNoQuoteEntriesMinValue() || m.TotNoQuoteEntries > m.TotNoQuoteEntriesMaxValue()) {
			return fmt.Errorf("Range check failed on m.TotNoQuoteEntries (%v < %v > %v)", m.TotNoQuoteEntriesMinValue(), m.TotNoQuoteEntries, m.TotNoQuoteEntriesMaxValue())
		}
	}
	if err := m.PossRetransFlag.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if m.DelayToTimeInActingVersion(actingVersion) {
		if m.DelayToTime != m.DelayToTimeNullValue() && (m.DelayToTime < m.DelayToTimeMinValue() || m.DelayToTime > m.DelayToTimeMaxValue()) {
			return fmt.Errorf("Range check failed on m.DelayToTime (%v < %v > %v)", m.DelayToTimeMinValue(), m.DelayToTime, m.DelayToTimeMaxValue())
		}
	}
	for _, prop := range m.NoQuoteEntries {
		if err := prop.RangeCheck(actingVersion, schemaVersion); err != nil {
			return err
		}
	}
	return nil
}

func MassQuoteAck545Init(m *MassQuoteAck545) {
	for idx := 0; idx < 256; idx++ {
		m.Text[idx] = 0
	}
	m.QuoteReqID = 18446744073709551615
	m.QuoteRejectReason = 65535
	m.DelayDuration = 65535
	m.TotNoQuoteEntries = 255
	m.DelayToTime = 18446744073709551615
	return
}

func (m *MassQuoteAck545NoQuoteEntries) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint32(_w, m.QuoteEntryID); err != nil {
		return err
	}
	if err := _m.WriteInt32(_w, m.SecurityID); err != nil {
		return err
	}
	if err := _m.WriteUint16(_w, m.QuoteSetID); err != nil {
		return err
	}
	if err := _m.WriteUint8(_w, m.QuoteEntryRejectReason); err != nil {
		return err
	}
	return nil
}

func (m *MassQuoteAck545NoQuoteEntries) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint) error {
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
	if !m.QuoteSetIDInActingVersion(actingVersion) {
		m.QuoteSetID = m.QuoteSetIDNullValue()
	} else {
		if err := _m.ReadUint16(_r, &m.QuoteSetID); err != nil {
			return err
		}
	}
	if !m.QuoteEntryRejectReasonInActingVersion(actingVersion) {
		m.QuoteEntryRejectReason = m.QuoteEntryRejectReasonNullValue()
	} else {
		if err := _m.ReadUint8(_r, &m.QuoteEntryRejectReason); err != nil {
			return err
		}
	}
	if actingVersion > m.SbeSchemaVersion() && blockLength > m.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-m.SbeBlockLength()))
	}
	return nil
}

func (m *MassQuoteAck545NoQuoteEntries) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
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
	if m.QuoteSetIDInActingVersion(actingVersion) {
		if m.QuoteSetID < m.QuoteSetIDMinValue() || m.QuoteSetID > m.QuoteSetIDMaxValue() {
			return fmt.Errorf("Range check failed on m.QuoteSetID (%v < %v > %v)", m.QuoteSetIDMinValue(), m.QuoteSetID, m.QuoteSetIDMaxValue())
		}
	}
	if m.QuoteEntryRejectReasonInActingVersion(actingVersion) {
		if m.QuoteEntryRejectReason < m.QuoteEntryRejectReasonMinValue() || m.QuoteEntryRejectReason > m.QuoteEntryRejectReasonMaxValue() {
			return fmt.Errorf("Range check failed on m.QuoteEntryRejectReason (%v < %v > %v)", m.QuoteEntryRejectReasonMinValue(), m.QuoteEntryRejectReason, m.QuoteEntryRejectReasonMaxValue())
		}
	}
	return nil
}

func MassQuoteAck545NoQuoteEntriesInit(m *MassQuoteAck545NoQuoteEntries) {
	return
}

func (*MassQuoteAck545) SbeBlockLength() (blockLength uint16) {
	return 350
}

func (*MassQuoteAck545) SbeTemplateId() (templateId uint16) {
	return 545
}

func (*MassQuoteAck545) SbeSchemaId() (schemaId uint16) {
	return 8
}

func (*MassQuoteAck545) SbeSchemaVersion() (schemaVersion uint16) {
	return 5
}

func (*MassQuoteAck545) SbeSemanticType() (semanticType []byte) {
	return []byte("b")
}

func (*MassQuoteAck545) SeqNumId() uint16 {
	return 9726
}

func (*MassQuoteAck545) SeqNumSinceVersion() uint16 {
	return 0
}

func (m *MassQuoteAck545) SeqNumInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.SeqNumSinceVersion()
}

func (*MassQuoteAck545) SeqNumDeprecated() uint16 {
	return 0
}

func (*MassQuoteAck545) SeqNumMetaAttribute(meta int) string {
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

func (*MassQuoteAck545) SeqNumMinValue() uint32 {
	return 0
}

func (*MassQuoteAck545) SeqNumMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*MassQuoteAck545) SeqNumNullValue() uint32 {
	return math.MaxUint32
}

func (*MassQuoteAck545) UUIDId() uint16 {
	return 39001
}

func (*MassQuoteAck545) UUIDSinceVersion() uint16 {
	return 0
}

func (m *MassQuoteAck545) UUIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.UUIDSinceVersion()
}

func (*MassQuoteAck545) UUIDDeprecated() uint16 {
	return 0
}

func (*MassQuoteAck545) UUIDMetaAttribute(meta int) string {
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

func (*MassQuoteAck545) UUIDMinValue() uint64 {
	return 0
}

func (*MassQuoteAck545) UUIDMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*MassQuoteAck545) UUIDNullValue() uint64 {
	return math.MaxUint64
}

func (*MassQuoteAck545) TextId() uint16 {
	return 58
}

func (*MassQuoteAck545) TextSinceVersion() uint16 {
	return 0
}

func (m *MassQuoteAck545) TextInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.TextSinceVersion()
}

func (*MassQuoteAck545) TextDeprecated() uint16 {
	return 0
}

func (*MassQuoteAck545) TextMetaAttribute(meta int) string {
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

func (*MassQuoteAck545) TextMinValue() byte {
	return byte(32)
}

func (*MassQuoteAck545) TextMaxValue() byte {
	return byte(126)
}

func (*MassQuoteAck545) TextNullValue() byte {
	return 0
}

func (m *MassQuoteAck545) TextCharacterEncoding() string {
	return "US-ASCII"
}

func (*MassQuoteAck545) SenderIDId() uint16 {
	return 5392
}

func (*MassQuoteAck545) SenderIDSinceVersion() uint16 {
	return 0
}

func (m *MassQuoteAck545) SenderIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.SenderIDSinceVersion()
}

func (*MassQuoteAck545) SenderIDDeprecated() uint16 {
	return 0
}

func (*MassQuoteAck545) SenderIDMetaAttribute(meta int) string {
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

func (*MassQuoteAck545) SenderIDMinValue() byte {
	return byte(32)
}

func (*MassQuoteAck545) SenderIDMaxValue() byte {
	return byte(126)
}

func (*MassQuoteAck545) SenderIDNullValue() byte {
	return 0
}

func (m *MassQuoteAck545) SenderIDCharacterEncoding() string {
	return "US-ASCII"
}

func (*MassQuoteAck545) PartyDetailsListReqIDId() uint16 {
	return 1505
}

func (*MassQuoteAck545) PartyDetailsListReqIDSinceVersion() uint16 {
	return 0
}

func (m *MassQuoteAck545) PartyDetailsListReqIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.PartyDetailsListReqIDSinceVersion()
}

func (*MassQuoteAck545) PartyDetailsListReqIDDeprecated() uint16 {
	return 0
}

func (*MassQuoteAck545) PartyDetailsListReqIDMetaAttribute(meta int) string {
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

func (*MassQuoteAck545) PartyDetailsListReqIDMinValue() uint64 {
	return 0
}

func (*MassQuoteAck545) PartyDetailsListReqIDMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*MassQuoteAck545) PartyDetailsListReqIDNullValue() uint64 {
	return math.MaxUint64
}

func (*MassQuoteAck545) RequestTimeId() uint16 {
	return 5979
}

func (*MassQuoteAck545) RequestTimeSinceVersion() uint16 {
	return 0
}

func (m *MassQuoteAck545) RequestTimeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.RequestTimeSinceVersion()
}

func (*MassQuoteAck545) RequestTimeDeprecated() uint16 {
	return 0
}

func (*MassQuoteAck545) RequestTimeMetaAttribute(meta int) string {
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

func (*MassQuoteAck545) RequestTimeMinValue() uint64 {
	return 0
}

func (*MassQuoteAck545) RequestTimeMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*MassQuoteAck545) RequestTimeNullValue() uint64 {
	return math.MaxUint64
}

func (*MassQuoteAck545) SendingTimeEpochId() uint16 {
	return 5297
}

func (*MassQuoteAck545) SendingTimeEpochSinceVersion() uint16 {
	return 0
}

func (m *MassQuoteAck545) SendingTimeEpochInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.SendingTimeEpochSinceVersion()
}

func (*MassQuoteAck545) SendingTimeEpochDeprecated() uint16 {
	return 0
}

func (*MassQuoteAck545) SendingTimeEpochMetaAttribute(meta int) string {
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

func (*MassQuoteAck545) SendingTimeEpochMinValue() uint64 {
	return 0
}

func (*MassQuoteAck545) SendingTimeEpochMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*MassQuoteAck545) SendingTimeEpochNullValue() uint64 {
	return math.MaxUint64
}

func (*MassQuoteAck545) QuoteReqIDId() uint16 {
	return 131
}

func (*MassQuoteAck545) QuoteReqIDSinceVersion() uint16 {
	return 0
}

func (m *MassQuoteAck545) QuoteReqIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.QuoteReqIDSinceVersion()
}

func (*MassQuoteAck545) QuoteReqIDDeprecated() uint16 {
	return 0
}

func (*MassQuoteAck545) QuoteReqIDMetaAttribute(meta int) string {
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

func (*MassQuoteAck545) QuoteReqIDMinValue() uint64 {
	return 0
}

func (*MassQuoteAck545) QuoteReqIDMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*MassQuoteAck545) QuoteReqIDNullValue() uint64 {
	return 18446744073709551615
}

func (*MassQuoteAck545) LocationId() uint16 {
	return 9537
}

func (*MassQuoteAck545) LocationSinceVersion() uint16 {
	return 0
}

func (m *MassQuoteAck545) LocationInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.LocationSinceVersion()
}

func (*MassQuoteAck545) LocationDeprecated() uint16 {
	return 0
}

func (*MassQuoteAck545) LocationMetaAttribute(meta int) string {
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

func (*MassQuoteAck545) LocationMinValue() byte {
	return byte(32)
}

func (*MassQuoteAck545) LocationMaxValue() byte {
	return byte(126)
}

func (*MassQuoteAck545) LocationNullValue() byte {
	return 0
}

func (m *MassQuoteAck545) LocationCharacterEncoding() string {
	return "US-ASCII"
}

func (*MassQuoteAck545) QuoteIDId() uint16 {
	return 117
}

func (*MassQuoteAck545) QuoteIDSinceVersion() uint16 {
	return 0
}

func (m *MassQuoteAck545) QuoteIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.QuoteIDSinceVersion()
}

func (*MassQuoteAck545) QuoteIDDeprecated() uint16 {
	return 0
}

func (*MassQuoteAck545) QuoteIDMetaAttribute(meta int) string {
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

func (*MassQuoteAck545) QuoteIDMinValue() uint32 {
	return 0
}

func (*MassQuoteAck545) QuoteIDMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*MassQuoteAck545) QuoteIDNullValue() uint32 {
	return math.MaxUint32
}

func (*MassQuoteAck545) QuoteRejectReasonId() uint16 {
	return 300
}

func (*MassQuoteAck545) QuoteRejectReasonSinceVersion() uint16 {
	return 0
}

func (m *MassQuoteAck545) QuoteRejectReasonInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.QuoteRejectReasonSinceVersion()
}

func (*MassQuoteAck545) QuoteRejectReasonDeprecated() uint16 {
	return 0
}

func (*MassQuoteAck545) QuoteRejectReasonMetaAttribute(meta int) string {
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

func (*MassQuoteAck545) QuoteRejectReasonMinValue() uint16 {
	return 0
}

func (*MassQuoteAck545) QuoteRejectReasonMaxValue() uint16 {
	return math.MaxUint16 - 1
}

func (*MassQuoteAck545) QuoteRejectReasonNullValue() uint16 {
	return 65535
}

func (*MassQuoteAck545) DelayDurationId() uint16 {
	return 5904
}

func (*MassQuoteAck545) DelayDurationSinceVersion() uint16 {
	return 0
}

func (m *MassQuoteAck545) DelayDurationInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.DelayDurationSinceVersion()
}

func (*MassQuoteAck545) DelayDurationDeprecated() uint16 {
	return 0
}

func (*MassQuoteAck545) DelayDurationMetaAttribute(meta int) string {
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

func (*MassQuoteAck545) DelayDurationMinValue() uint16 {
	return 0
}

func (*MassQuoteAck545) DelayDurationMaxValue() uint16 {
	return math.MaxUint16 - 1
}

func (*MassQuoteAck545) DelayDurationNullValue() uint16 {
	return 65535
}

func (*MassQuoteAck545) QuoteStatusId() uint16 {
	return 297
}

func (*MassQuoteAck545) QuoteStatusSinceVersion() uint16 {
	return 0
}

func (m *MassQuoteAck545) QuoteStatusInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.QuoteStatusSinceVersion()
}

func (*MassQuoteAck545) QuoteStatusDeprecated() uint16 {
	return 0
}

func (*MassQuoteAck545) QuoteStatusMetaAttribute(meta int) string {
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

func (*MassQuoteAck545) ManualOrderIndicatorId() uint16 {
	return 1028
}

func (*MassQuoteAck545) ManualOrderIndicatorSinceVersion() uint16 {
	return 0
}

func (m *MassQuoteAck545) ManualOrderIndicatorInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.ManualOrderIndicatorSinceVersion()
}

func (*MassQuoteAck545) ManualOrderIndicatorDeprecated() uint16 {
	return 0
}

func (*MassQuoteAck545) ManualOrderIndicatorMetaAttribute(meta int) string {
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

func (*MassQuoteAck545) NoProcessedEntriesId() uint16 {
	return 9772
}

func (*MassQuoteAck545) NoProcessedEntriesSinceVersion() uint16 {
	return 0
}

func (m *MassQuoteAck545) NoProcessedEntriesInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.NoProcessedEntriesSinceVersion()
}

func (*MassQuoteAck545) NoProcessedEntriesDeprecated() uint16 {
	return 0
}

func (*MassQuoteAck545) NoProcessedEntriesMetaAttribute(meta int) string {
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

func (*MassQuoteAck545) NoProcessedEntriesMinValue() uint8 {
	return 0
}

func (*MassQuoteAck545) NoProcessedEntriesMaxValue() uint8 {
	return math.MaxUint8 - 1
}

func (*MassQuoteAck545) NoProcessedEntriesNullValue() uint8 {
	return math.MaxUint8
}

func (*MassQuoteAck545) MMProtectionResetId() uint16 {
	return 9773
}

func (*MassQuoteAck545) MMProtectionResetSinceVersion() uint16 {
	return 0
}

func (m *MassQuoteAck545) MMProtectionResetInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.MMProtectionResetSinceVersion()
}

func (*MassQuoteAck545) MMProtectionResetDeprecated() uint16 {
	return 0
}

func (*MassQuoteAck545) MMProtectionResetMetaAttribute(meta int) string {
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

func (*MassQuoteAck545) SplitMsgId() uint16 {
	return 9553
}

func (*MassQuoteAck545) SplitMsgSinceVersion() uint16 {
	return 0
}

func (m *MassQuoteAck545) SplitMsgInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.SplitMsgSinceVersion()
}

func (*MassQuoteAck545) SplitMsgDeprecated() uint16 {
	return 0
}

func (*MassQuoteAck545) SplitMsgMetaAttribute(meta int) string {
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

func (*MassQuoteAck545) LiquidityFlagId() uint16 {
	return 9373
}

func (*MassQuoteAck545) LiquidityFlagSinceVersion() uint16 {
	return 0
}

func (m *MassQuoteAck545) LiquidityFlagInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.LiquidityFlagSinceVersion()
}

func (*MassQuoteAck545) LiquidityFlagDeprecated() uint16 {
	return 0
}

func (*MassQuoteAck545) LiquidityFlagMetaAttribute(meta int) string {
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

func (*MassQuoteAck545) ShortSaleTypeId() uint16 {
	return 5409
}

func (*MassQuoteAck545) ShortSaleTypeSinceVersion() uint16 {
	return 0
}

func (m *MassQuoteAck545) ShortSaleTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.ShortSaleTypeSinceVersion()
}

func (*MassQuoteAck545) ShortSaleTypeDeprecated() uint16 {
	return 0
}

func (*MassQuoteAck545) ShortSaleTypeMetaAttribute(meta int) string {
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

func (*MassQuoteAck545) TotNoQuoteEntriesId() uint16 {
	return 304
}

func (*MassQuoteAck545) TotNoQuoteEntriesSinceVersion() uint16 {
	return 0
}

func (m *MassQuoteAck545) TotNoQuoteEntriesInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.TotNoQuoteEntriesSinceVersion()
}

func (*MassQuoteAck545) TotNoQuoteEntriesDeprecated() uint16 {
	return 0
}

func (*MassQuoteAck545) TotNoQuoteEntriesMetaAttribute(meta int) string {
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

func (*MassQuoteAck545) TotNoQuoteEntriesMinValue() uint8 {
	return 0
}

func (*MassQuoteAck545) TotNoQuoteEntriesMaxValue() uint8 {
	return math.MaxUint8 - 1
}

func (*MassQuoteAck545) TotNoQuoteEntriesNullValue() uint8 {
	return 255
}

func (*MassQuoteAck545) PossRetransFlagId() uint16 {
	return 9765
}

func (*MassQuoteAck545) PossRetransFlagSinceVersion() uint16 {
	return 2
}

func (m *MassQuoteAck545) PossRetransFlagInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.PossRetransFlagSinceVersion()
}

func (*MassQuoteAck545) PossRetransFlagDeprecated() uint16 {
	return 0
}

func (*MassQuoteAck545) PossRetransFlagMetaAttribute(meta int) string {
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

func (*MassQuoteAck545) DelayToTimeId() uint16 {
	return 7552
}

func (*MassQuoteAck545) DelayToTimeSinceVersion() uint16 {
	return 4
}

func (m *MassQuoteAck545) DelayToTimeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.DelayToTimeSinceVersion()
}

func (*MassQuoteAck545) DelayToTimeDeprecated() uint16 {
	return 0
}

func (*MassQuoteAck545) DelayToTimeMetaAttribute(meta int) string {
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

func (*MassQuoteAck545) DelayToTimeMinValue() uint64 {
	return 0
}

func (*MassQuoteAck545) DelayToTimeMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*MassQuoteAck545) DelayToTimeNullValue() uint64 {
	return 18446744073709551615
}

func (*MassQuoteAck545NoQuoteEntries) QuoteEntryIDId() uint16 {
	return 299
}

func (*MassQuoteAck545NoQuoteEntries) QuoteEntryIDSinceVersion() uint16 {
	return 0
}

func (m *MassQuoteAck545NoQuoteEntries) QuoteEntryIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.QuoteEntryIDSinceVersion()
}

func (*MassQuoteAck545NoQuoteEntries) QuoteEntryIDDeprecated() uint16 {
	return 0
}

func (*MassQuoteAck545NoQuoteEntries) QuoteEntryIDMetaAttribute(meta int) string {
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

func (*MassQuoteAck545NoQuoteEntries) QuoteEntryIDMinValue() uint32 {
	return 0
}

func (*MassQuoteAck545NoQuoteEntries) QuoteEntryIDMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*MassQuoteAck545NoQuoteEntries) QuoteEntryIDNullValue() uint32 {
	return math.MaxUint32
}

func (*MassQuoteAck545NoQuoteEntries) SecurityIDId() uint16 {
	return 48
}

func (*MassQuoteAck545NoQuoteEntries) SecurityIDSinceVersion() uint16 {
	return 0
}

func (m *MassQuoteAck545NoQuoteEntries) SecurityIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.SecurityIDSinceVersion()
}

func (*MassQuoteAck545NoQuoteEntries) SecurityIDDeprecated() uint16 {
	return 0
}

func (*MassQuoteAck545NoQuoteEntries) SecurityIDMetaAttribute(meta int) string {
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

func (*MassQuoteAck545NoQuoteEntries) SecurityIDMinValue() int32 {
	return math.MinInt32 + 1
}

func (*MassQuoteAck545NoQuoteEntries) SecurityIDMaxValue() int32 {
	return math.MaxInt32
}

func (*MassQuoteAck545NoQuoteEntries) SecurityIDNullValue() int32 {
	return math.MinInt32
}

func (*MassQuoteAck545NoQuoteEntries) QuoteSetIDId() uint16 {
	return 302
}

func (*MassQuoteAck545NoQuoteEntries) QuoteSetIDSinceVersion() uint16 {
	return 0
}

func (m *MassQuoteAck545NoQuoteEntries) QuoteSetIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.QuoteSetIDSinceVersion()
}

func (*MassQuoteAck545NoQuoteEntries) QuoteSetIDDeprecated() uint16 {
	return 0
}

func (*MassQuoteAck545NoQuoteEntries) QuoteSetIDMetaAttribute(meta int) string {
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

func (*MassQuoteAck545NoQuoteEntries) QuoteSetIDMinValue() uint16 {
	return 0
}

func (*MassQuoteAck545NoQuoteEntries) QuoteSetIDMaxValue() uint16 {
	return math.MaxUint16 - 1
}

func (*MassQuoteAck545NoQuoteEntries) QuoteSetIDNullValue() uint16 {
	return math.MaxUint16
}

func (*MassQuoteAck545NoQuoteEntries) QuoteEntryRejectReasonId() uint16 {
	return 368
}

func (*MassQuoteAck545NoQuoteEntries) QuoteEntryRejectReasonSinceVersion() uint16 {
	return 0
}

func (m *MassQuoteAck545NoQuoteEntries) QuoteEntryRejectReasonInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.QuoteEntryRejectReasonSinceVersion()
}

func (*MassQuoteAck545NoQuoteEntries) QuoteEntryRejectReasonDeprecated() uint16 {
	return 0
}

func (*MassQuoteAck545NoQuoteEntries) QuoteEntryRejectReasonMetaAttribute(meta int) string {
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

func (*MassQuoteAck545NoQuoteEntries) QuoteEntryRejectReasonMinValue() uint8 {
	return 0
}

func (*MassQuoteAck545NoQuoteEntries) QuoteEntryRejectReasonMaxValue() uint8 {
	return math.MaxUint8 - 1
}

func (*MassQuoteAck545NoQuoteEntries) QuoteEntryRejectReasonNullValue() uint8 {
	return math.MaxUint8
}

func (*MassQuoteAck545) NoQuoteEntriesId() uint16 {
	return 295
}

func (*MassQuoteAck545) NoQuoteEntriesSinceVersion() uint16 {
	return 0
}

func (m *MassQuoteAck545) NoQuoteEntriesInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.NoQuoteEntriesSinceVersion()
}

func (*MassQuoteAck545) NoQuoteEntriesDeprecated() uint16 {
	return 0
}

func (*MassQuoteAck545NoQuoteEntries) SbeBlockLength() (blockLength uint) {
	return 11
}

func (*MassQuoteAck545NoQuoteEntries) SbeSchemaVersion() (schemaVersion uint16) {
	return 5
}
