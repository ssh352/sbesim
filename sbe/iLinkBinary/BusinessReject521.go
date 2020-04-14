// Generated SBE (Simple Binary Encoding) message codec

package ilinkbinary

import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
)

type BusinessReject521 struct {
	SeqNum                uint32
	UUID                  uint64
	Text                  [256]byte
	SenderID              [20]byte
	PartyDetailsListReqID uint64
	SendingTimeEpoch      uint64
	BusinessRejectRefID   uint64
	Location              [5]byte
	RefSeqNum             uint32
	RefTagID              uint16
	BusinessRejectReason  uint16
	RefMsgType            [2]byte
	PossRetransFlag       BooleanFlagEnum
	ManualOrderIndicator  ManualOrdIndEnum
	SplitMsg              SplitMsgEnum
}

func (b *BusinessReject521) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := b.RangeCheck(b.SbeSchemaVersion(), b.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	if err := _m.WriteUint32(_w, b.SeqNum); err != nil {
		return err
	}
	if err := _m.WriteUint64(_w, b.UUID); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, b.Text[:]); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, b.SenderID[:]); err != nil {
		return err
	}
	if err := _m.WriteUint64(_w, b.PartyDetailsListReqID); err != nil {
		return err
	}
	if err := _m.WriteUint64(_w, b.SendingTimeEpoch); err != nil {
		return err
	}
	if err := _m.WriteUint64(_w, b.BusinessRejectRefID); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, b.Location[:]); err != nil {
		return err
	}
	if err := _m.WriteUint32(_w, b.RefSeqNum); err != nil {
		return err
	}
	if err := _m.WriteUint16(_w, b.RefTagID); err != nil {
		return err
	}
	if err := _m.WriteUint16(_w, b.BusinessRejectReason); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, b.RefMsgType[:]); err != nil {
		return err
	}
	if err := b.PossRetransFlag.Encode(_m, _w); err != nil {
		return err
	}
	if err := b.ManualOrderIndicator.Encode(_m, _w); err != nil {
		return err
	}
	if err := b.SplitMsg.Encode(_m, _w); err != nil {
		return err
	}
	return nil
}

func (b *BusinessReject521) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	if !b.SeqNumInActingVersion(actingVersion) {
		b.SeqNum = b.SeqNumNullValue()
	} else {
		if err := _m.ReadUint32(_r, &b.SeqNum); err != nil {
			return err
		}
	}
	if !b.UUIDInActingVersion(actingVersion) {
		b.UUID = b.UUIDNullValue()
	} else {
		if err := _m.ReadUint64(_r, &b.UUID); err != nil {
			return err
		}
	}
	if !b.TextInActingVersion(actingVersion) {
		for idx := 0; idx < 256; idx++ {
			b.Text[idx] = b.TextNullValue()
		}
	} else {
		if err := _m.ReadBytes(_r, b.Text[:]); err != nil {
			return err
		}
	}
	if !b.SenderIDInActingVersion(actingVersion) {
		for idx := 0; idx < 20; idx++ {
			b.SenderID[idx] = b.SenderIDNullValue()
		}
	} else {
		if err := _m.ReadBytes(_r, b.SenderID[:]); err != nil {
			return err
		}
	}
	if !b.PartyDetailsListReqIDInActingVersion(actingVersion) {
		b.PartyDetailsListReqID = b.PartyDetailsListReqIDNullValue()
	} else {
		if err := _m.ReadUint64(_r, &b.PartyDetailsListReqID); err != nil {
			return err
		}
	}
	if !b.SendingTimeEpochInActingVersion(actingVersion) {
		b.SendingTimeEpoch = b.SendingTimeEpochNullValue()
	} else {
		if err := _m.ReadUint64(_r, &b.SendingTimeEpoch); err != nil {
			return err
		}
	}
	if !b.BusinessRejectRefIDInActingVersion(actingVersion) {
		b.BusinessRejectRefID = b.BusinessRejectRefIDNullValue()
	} else {
		if err := _m.ReadUint64(_r, &b.BusinessRejectRefID); err != nil {
			return err
		}
	}
	if !b.LocationInActingVersion(actingVersion) {
		for idx := 0; idx < 5; idx++ {
			b.Location[idx] = b.LocationNullValue()
		}
	} else {
		if err := _m.ReadBytes(_r, b.Location[:]); err != nil {
			return err
		}
	}
	if !b.RefSeqNumInActingVersion(actingVersion) {
		b.RefSeqNum = b.RefSeqNumNullValue()
	} else {
		if err := _m.ReadUint32(_r, &b.RefSeqNum); err != nil {
			return err
		}
	}
	if !b.RefTagIDInActingVersion(actingVersion) {
		b.RefTagID = b.RefTagIDNullValue()
	} else {
		if err := _m.ReadUint16(_r, &b.RefTagID); err != nil {
			return err
		}
	}
	if !b.BusinessRejectReasonInActingVersion(actingVersion) {
		b.BusinessRejectReason = b.BusinessRejectReasonNullValue()
	} else {
		if err := _m.ReadUint16(_r, &b.BusinessRejectReason); err != nil {
			return err
		}
	}
	if !b.RefMsgTypeInActingVersion(actingVersion) {
		for idx := 0; idx < 2; idx++ {
			b.RefMsgType[idx] = b.RefMsgTypeNullValue()
		}
	} else {
		if err := _m.ReadBytes(_r, b.RefMsgType[:]); err != nil {
			return err
		}
	}
	if b.PossRetransFlagInActingVersion(actingVersion) {
		if err := b.PossRetransFlag.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if b.ManualOrderIndicatorInActingVersion(actingVersion) {
		if err := b.ManualOrderIndicator.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if b.SplitMsgInActingVersion(actingVersion) {
		if err := b.SplitMsg.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if actingVersion > b.SbeSchemaVersion() && blockLength > b.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-b.SbeBlockLength()))
	}
	if doRangeCheck {
		if err := b.RangeCheck(actingVersion, b.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	return nil
}

func (b *BusinessReject521) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if b.SeqNumInActingVersion(actingVersion) {
		if b.SeqNum < b.SeqNumMinValue() || b.SeqNum > b.SeqNumMaxValue() {
			return fmt.Errorf("Range check failed on b.SeqNum (%v < %v > %v)", b.SeqNumMinValue(), b.SeqNum, b.SeqNumMaxValue())
		}
	}
	if b.UUIDInActingVersion(actingVersion) {
		if b.UUID < b.UUIDMinValue() || b.UUID > b.UUIDMaxValue() {
			return fmt.Errorf("Range check failed on b.UUID (%v < %v > %v)", b.UUIDMinValue(), b.UUID, b.UUIDMaxValue())
		}
	}
	if b.TextInActingVersion(actingVersion) {
		for idx := 0; idx < 256; idx++ {
			if b.Text[idx] < b.TextMinValue() || b.Text[idx] > b.TextMaxValue() {
				return fmt.Errorf("Range check failed on b.Text[%d] (%v < %v > %v)", idx, b.TextMinValue(), b.Text[idx], b.TextMaxValue())
			}
		}
	}
	if b.SenderIDInActingVersion(actingVersion) {
		for idx := 0; idx < 20; idx++ {
			if b.SenderID[idx] < b.SenderIDMinValue() || b.SenderID[idx] > b.SenderIDMaxValue() {
				return fmt.Errorf("Range check failed on b.SenderID[%d] (%v < %v > %v)", idx, b.SenderIDMinValue(), b.SenderID[idx], b.SenderIDMaxValue())
			}
		}
	}
	if b.PartyDetailsListReqIDInActingVersion(actingVersion) {
		if b.PartyDetailsListReqID != b.PartyDetailsListReqIDNullValue() && (b.PartyDetailsListReqID < b.PartyDetailsListReqIDMinValue() || b.PartyDetailsListReqID > b.PartyDetailsListReqIDMaxValue()) {
			return fmt.Errorf("Range check failed on b.PartyDetailsListReqID (%v < %v > %v)", b.PartyDetailsListReqIDMinValue(), b.PartyDetailsListReqID, b.PartyDetailsListReqIDMaxValue())
		}
	}
	if b.SendingTimeEpochInActingVersion(actingVersion) {
		if b.SendingTimeEpoch < b.SendingTimeEpochMinValue() || b.SendingTimeEpoch > b.SendingTimeEpochMaxValue() {
			return fmt.Errorf("Range check failed on b.SendingTimeEpoch (%v < %v > %v)", b.SendingTimeEpochMinValue(), b.SendingTimeEpoch, b.SendingTimeEpochMaxValue())
		}
	}
	if b.BusinessRejectRefIDInActingVersion(actingVersion) {
		if b.BusinessRejectRefID != b.BusinessRejectRefIDNullValue() && (b.BusinessRejectRefID < b.BusinessRejectRefIDMinValue() || b.BusinessRejectRefID > b.BusinessRejectRefIDMaxValue()) {
			return fmt.Errorf("Range check failed on b.BusinessRejectRefID (%v < %v > %v)", b.BusinessRejectRefIDMinValue(), b.BusinessRejectRefID, b.BusinessRejectRefIDMaxValue())
		}
	}
	if b.LocationInActingVersion(actingVersion) {
		for idx := 0; idx < 5; idx++ {
			if b.Location[idx] < b.LocationMinValue() || b.Location[idx] > b.LocationMaxValue() {
				return fmt.Errorf("Range check failed on b.Location[%d] (%v < %v > %v)", idx, b.LocationMinValue(), b.Location[idx], b.LocationMaxValue())
			}
		}
	}
	if b.RefSeqNumInActingVersion(actingVersion) {
		if b.RefSeqNum != b.RefSeqNumNullValue() && (b.RefSeqNum < b.RefSeqNumMinValue() || b.RefSeqNum > b.RefSeqNumMaxValue()) {
			return fmt.Errorf("Range check failed on b.RefSeqNum (%v < %v > %v)", b.RefSeqNumMinValue(), b.RefSeqNum, b.RefSeqNumMaxValue())
		}
	}
	if b.RefTagIDInActingVersion(actingVersion) {
		if b.RefTagID != b.RefTagIDNullValue() && (b.RefTagID < b.RefTagIDMinValue() || b.RefTagID > b.RefTagIDMaxValue()) {
			return fmt.Errorf("Range check failed on b.RefTagID (%v < %v > %v)", b.RefTagIDMinValue(), b.RefTagID, b.RefTagIDMaxValue())
		}
	}
	if b.BusinessRejectReasonInActingVersion(actingVersion) {
		if b.BusinessRejectReason < b.BusinessRejectReasonMinValue() || b.BusinessRejectReason > b.BusinessRejectReasonMaxValue() {
			return fmt.Errorf("Range check failed on b.BusinessRejectReason (%v < %v > %v)", b.BusinessRejectReasonMinValue(), b.BusinessRejectReason, b.BusinessRejectReasonMaxValue())
		}
	}
	if b.RefMsgTypeInActingVersion(actingVersion) {
		for idx := 0; idx < 2; idx++ {
			if b.RefMsgType[idx] < b.RefMsgTypeMinValue() || b.RefMsgType[idx] > b.RefMsgTypeMaxValue() {
				return fmt.Errorf("Range check failed on b.RefMsgType[%d] (%v < %v > %v)", idx, b.RefMsgTypeMinValue(), b.RefMsgType[idx], b.RefMsgTypeMaxValue())
			}
		}
	}
	if err := b.PossRetransFlag.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if err := b.ManualOrderIndicator.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if err := b.SplitMsg.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	return nil
}

func BusinessReject521Init(b *BusinessReject521) {
	for idx := 0; idx < 256; idx++ {
		b.Text[idx] = 0
	}
	for idx := 0; idx < 20; idx++ {
		b.SenderID[idx] = 0
	}
	b.PartyDetailsListReqID = 18446744073709551615
	b.BusinessRejectRefID = 18446744073709551615
	for idx := 0; idx < 5; idx++ {
		b.Location[idx] = 0
	}
	b.RefSeqNum = 4294967295
	b.RefTagID = 65535
	return
}

func (*BusinessReject521) SbeBlockLength() (blockLength uint16) {
	return 330
}

func (*BusinessReject521) SbeTemplateId() (templateId uint16) {
	return 521
}

func (*BusinessReject521) SbeSchemaId() (schemaId uint16) {
	return 8
}

func (*BusinessReject521) SbeSchemaVersion() (schemaVersion uint16) {
	return 5
}

func (*BusinessReject521) SbeSemanticType() (semanticType []byte) {
	return []byte("j")
}

func (*BusinessReject521) SeqNumId() uint16 {
	return 9726
}

func (*BusinessReject521) SeqNumSinceVersion() uint16 {
	return 0
}

func (b *BusinessReject521) SeqNumInActingVersion(actingVersion uint16) bool {
	return actingVersion >= b.SeqNumSinceVersion()
}

func (*BusinessReject521) SeqNumDeprecated() uint16 {
	return 0
}

func (*BusinessReject521) SeqNumMetaAttribute(meta int) string {
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

func (*BusinessReject521) SeqNumMinValue() uint32 {
	return 0
}

func (*BusinessReject521) SeqNumMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*BusinessReject521) SeqNumNullValue() uint32 {
	return math.MaxUint32
}

func (*BusinessReject521) UUIDId() uint16 {
	return 39001
}

func (*BusinessReject521) UUIDSinceVersion() uint16 {
	return 0
}

func (b *BusinessReject521) UUIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= b.UUIDSinceVersion()
}

func (*BusinessReject521) UUIDDeprecated() uint16 {
	return 0
}

func (*BusinessReject521) UUIDMetaAttribute(meta int) string {
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

func (*BusinessReject521) UUIDMinValue() uint64 {
	return 0
}

func (*BusinessReject521) UUIDMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*BusinessReject521) UUIDNullValue() uint64 {
	return math.MaxUint64
}

func (*BusinessReject521) TextId() uint16 {
	return 58
}

func (*BusinessReject521) TextSinceVersion() uint16 {
	return 0
}

func (b *BusinessReject521) TextInActingVersion(actingVersion uint16) bool {
	return actingVersion >= b.TextSinceVersion()
}

func (*BusinessReject521) TextDeprecated() uint16 {
	return 0
}

func (*BusinessReject521) TextMetaAttribute(meta int) string {
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

func (*BusinessReject521) TextMinValue() byte {
	return byte(32)
}

func (*BusinessReject521) TextMaxValue() byte {
	return byte(126)
}

func (*BusinessReject521) TextNullValue() byte {
	return 0
}

func (b *BusinessReject521) TextCharacterEncoding() string {
	return "US-ASCII"
}

func (*BusinessReject521) SenderIDId() uint16 {
	return 5392
}

func (*BusinessReject521) SenderIDSinceVersion() uint16 {
	return 0
}

func (b *BusinessReject521) SenderIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= b.SenderIDSinceVersion()
}

func (*BusinessReject521) SenderIDDeprecated() uint16 {
	return 0
}

func (*BusinessReject521) SenderIDMetaAttribute(meta int) string {
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

func (*BusinessReject521) SenderIDMinValue() byte {
	return byte(32)
}

func (*BusinessReject521) SenderIDMaxValue() byte {
	return byte(126)
}

func (*BusinessReject521) SenderIDNullValue() byte {
	return 0
}

func (b *BusinessReject521) SenderIDCharacterEncoding() string {
	return "US-ASCII"
}

func (*BusinessReject521) PartyDetailsListReqIDId() uint16 {
	return 1505
}

func (*BusinessReject521) PartyDetailsListReqIDSinceVersion() uint16 {
	return 0
}

func (b *BusinessReject521) PartyDetailsListReqIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= b.PartyDetailsListReqIDSinceVersion()
}

func (*BusinessReject521) PartyDetailsListReqIDDeprecated() uint16 {
	return 0
}

func (*BusinessReject521) PartyDetailsListReqIDMetaAttribute(meta int) string {
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

func (*BusinessReject521) PartyDetailsListReqIDMinValue() uint64 {
	return 0
}

func (*BusinessReject521) PartyDetailsListReqIDMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*BusinessReject521) PartyDetailsListReqIDNullValue() uint64 {
	return 18446744073709551615
}

func (*BusinessReject521) SendingTimeEpochId() uint16 {
	return 5297
}

func (*BusinessReject521) SendingTimeEpochSinceVersion() uint16 {
	return 0
}

func (b *BusinessReject521) SendingTimeEpochInActingVersion(actingVersion uint16) bool {
	return actingVersion >= b.SendingTimeEpochSinceVersion()
}

func (*BusinessReject521) SendingTimeEpochDeprecated() uint16 {
	return 0
}

func (*BusinessReject521) SendingTimeEpochMetaAttribute(meta int) string {
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

func (*BusinessReject521) SendingTimeEpochMinValue() uint64 {
	return 0
}

func (*BusinessReject521) SendingTimeEpochMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*BusinessReject521) SendingTimeEpochNullValue() uint64 {
	return math.MaxUint64
}

func (*BusinessReject521) BusinessRejectRefIDId() uint16 {
	return 379
}

func (*BusinessReject521) BusinessRejectRefIDSinceVersion() uint16 {
	return 0
}

func (b *BusinessReject521) BusinessRejectRefIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= b.BusinessRejectRefIDSinceVersion()
}

func (*BusinessReject521) BusinessRejectRefIDDeprecated() uint16 {
	return 0
}

func (*BusinessReject521) BusinessRejectRefIDMetaAttribute(meta int) string {
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

func (*BusinessReject521) BusinessRejectRefIDMinValue() uint64 {
	return 0
}

func (*BusinessReject521) BusinessRejectRefIDMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*BusinessReject521) BusinessRejectRefIDNullValue() uint64 {
	return 18446744073709551615
}

func (*BusinessReject521) LocationId() uint16 {
	return 9537
}

func (*BusinessReject521) LocationSinceVersion() uint16 {
	return 0
}

func (b *BusinessReject521) LocationInActingVersion(actingVersion uint16) bool {
	return actingVersion >= b.LocationSinceVersion()
}

func (*BusinessReject521) LocationDeprecated() uint16 {
	return 0
}

func (*BusinessReject521) LocationMetaAttribute(meta int) string {
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

func (*BusinessReject521) LocationMinValue() byte {
	return byte(32)
}

func (*BusinessReject521) LocationMaxValue() byte {
	return byte(126)
}

func (*BusinessReject521) LocationNullValue() byte {
	return 0
}

func (b *BusinessReject521) LocationCharacterEncoding() string {
	return "US-ASCII"
}

func (*BusinessReject521) RefSeqNumId() uint16 {
	return 45
}

func (*BusinessReject521) RefSeqNumSinceVersion() uint16 {
	return 0
}

func (b *BusinessReject521) RefSeqNumInActingVersion(actingVersion uint16) bool {
	return actingVersion >= b.RefSeqNumSinceVersion()
}

func (*BusinessReject521) RefSeqNumDeprecated() uint16 {
	return 0
}

func (*BusinessReject521) RefSeqNumMetaAttribute(meta int) string {
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

func (*BusinessReject521) RefSeqNumMinValue() uint32 {
	return 0
}

func (*BusinessReject521) RefSeqNumMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*BusinessReject521) RefSeqNumNullValue() uint32 {
	return 4294967295
}

func (*BusinessReject521) RefTagIDId() uint16 {
	return 371
}

func (*BusinessReject521) RefTagIDSinceVersion() uint16 {
	return 0
}

func (b *BusinessReject521) RefTagIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= b.RefTagIDSinceVersion()
}

func (*BusinessReject521) RefTagIDDeprecated() uint16 {
	return 0
}

func (*BusinessReject521) RefTagIDMetaAttribute(meta int) string {
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

func (*BusinessReject521) RefTagIDMinValue() uint16 {
	return 0
}

func (*BusinessReject521) RefTagIDMaxValue() uint16 {
	return math.MaxUint16 - 1
}

func (*BusinessReject521) RefTagIDNullValue() uint16 {
	return 65535
}

func (*BusinessReject521) BusinessRejectReasonId() uint16 {
	return 380
}

func (*BusinessReject521) BusinessRejectReasonSinceVersion() uint16 {
	return 0
}

func (b *BusinessReject521) BusinessRejectReasonInActingVersion(actingVersion uint16) bool {
	return actingVersion >= b.BusinessRejectReasonSinceVersion()
}

func (*BusinessReject521) BusinessRejectReasonDeprecated() uint16 {
	return 0
}

func (*BusinessReject521) BusinessRejectReasonMetaAttribute(meta int) string {
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

func (*BusinessReject521) BusinessRejectReasonMinValue() uint16 {
	return 0
}

func (*BusinessReject521) BusinessRejectReasonMaxValue() uint16 {
	return math.MaxUint16 - 1
}

func (*BusinessReject521) BusinessRejectReasonNullValue() uint16 {
	return math.MaxUint16
}

func (*BusinessReject521) RefMsgTypeId() uint16 {
	return 372
}

func (*BusinessReject521) RefMsgTypeSinceVersion() uint16 {
	return 0
}

func (b *BusinessReject521) RefMsgTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= b.RefMsgTypeSinceVersion()
}

func (*BusinessReject521) RefMsgTypeDeprecated() uint16 {
	return 0
}

func (*BusinessReject521) RefMsgTypeMetaAttribute(meta int) string {
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

func (*BusinessReject521) RefMsgTypeMinValue() byte {
	return byte(32)
}

func (*BusinessReject521) RefMsgTypeMaxValue() byte {
	return byte(126)
}

func (*BusinessReject521) RefMsgTypeNullValue() byte {
	return 0
}

func (b *BusinessReject521) RefMsgTypeCharacterEncoding() string {
	return "US-ASCII"
}

func (*BusinessReject521) PossRetransFlagId() uint16 {
	return 9765
}

func (*BusinessReject521) PossRetransFlagSinceVersion() uint16 {
	return 0
}

func (b *BusinessReject521) PossRetransFlagInActingVersion(actingVersion uint16) bool {
	return actingVersion >= b.PossRetransFlagSinceVersion()
}

func (*BusinessReject521) PossRetransFlagDeprecated() uint16 {
	return 0
}

func (*BusinessReject521) PossRetransFlagMetaAttribute(meta int) string {
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

func (*BusinessReject521) ManualOrderIndicatorId() uint16 {
	return 1028
}

func (*BusinessReject521) ManualOrderIndicatorSinceVersion() uint16 {
	return 0
}

func (b *BusinessReject521) ManualOrderIndicatorInActingVersion(actingVersion uint16) bool {
	return actingVersion >= b.ManualOrderIndicatorSinceVersion()
}

func (*BusinessReject521) ManualOrderIndicatorDeprecated() uint16 {
	return 0
}

func (*BusinessReject521) ManualOrderIndicatorMetaAttribute(meta int) string {
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

func (*BusinessReject521) SplitMsgId() uint16 {
	return 9553
}

func (*BusinessReject521) SplitMsgSinceVersion() uint16 {
	return 0
}

func (b *BusinessReject521) SplitMsgInActingVersion(actingVersion uint16) bool {
	return actingVersion >= b.SplitMsgSinceVersion()
}

func (*BusinessReject521) SplitMsgDeprecated() uint16 {
	return 0
}

func (*BusinessReject521) SplitMsgMetaAttribute(meta int) string {
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
