// Generated SBE (Simple Binary Encoding) message codec

package ilinkbinary

import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
)

type OrderMassStatusRequest530 struct {
	PartyDetailsListReqID uint64
	MassStatusReqID       uint64
	ManualOrderIndicator  ManualOrdIndReqEnum
	SeqNum                uint32
	SenderID              [20]byte
	SendingTimeEpoch      uint64
	SecurityGroup         [6]byte
	Location              [5]byte
	SecurityID            int32
	MassStatusReqType     MassStatusReqTypEnum
	OrdStatusReqType      MassStatusOrdTypEnum
	TimeInForce           MassStatusTIFEnum
	MarketSegmentID       uint8
}

func (o *OrderMassStatusRequest530) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := o.RangeCheck(o.SbeSchemaVersion(), o.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	if err := _m.WriteUint64(_w, o.PartyDetailsListReqID); err != nil {
		return err
	}
	if err := _m.WriteUint64(_w, o.MassStatusReqID); err != nil {
		return err
	}
	if err := o.ManualOrderIndicator.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteUint32(_w, o.SeqNum); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, o.SenderID[:]); err != nil {
		return err
	}
	if err := _m.WriteUint64(_w, o.SendingTimeEpoch); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, o.SecurityGroup[:]); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, o.Location[:]); err != nil {
		return err
	}
	if err := _m.WriteInt32(_w, o.SecurityID); err != nil {
		return err
	}
	if err := o.MassStatusReqType.Encode(_m, _w); err != nil {
		return err
	}
	if err := o.OrdStatusReqType.Encode(_m, _w); err != nil {
		return err
	}
	if err := o.TimeInForce.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteUint8(_w, o.MarketSegmentID); err != nil {
		return err
	}
	return nil
}

func (o *OrderMassStatusRequest530) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	if !o.PartyDetailsListReqIDInActingVersion(actingVersion) {
		o.PartyDetailsListReqID = o.PartyDetailsListReqIDNullValue()
	} else {
		if err := _m.ReadUint64(_r, &o.PartyDetailsListReqID); err != nil {
			return err
		}
	}
	if !o.MassStatusReqIDInActingVersion(actingVersion) {
		o.MassStatusReqID = o.MassStatusReqIDNullValue()
	} else {
		if err := _m.ReadUint64(_r, &o.MassStatusReqID); err != nil {
			return err
		}
	}
	if o.ManualOrderIndicatorInActingVersion(actingVersion) {
		if err := o.ManualOrderIndicator.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if !o.SeqNumInActingVersion(actingVersion) {
		o.SeqNum = o.SeqNumNullValue()
	} else {
		if err := _m.ReadUint32(_r, &o.SeqNum); err != nil {
			return err
		}
	}
	if !o.SenderIDInActingVersion(actingVersion) {
		for idx := 0; idx < 20; idx++ {
			o.SenderID[idx] = o.SenderIDNullValue()
		}
	} else {
		if err := _m.ReadBytes(_r, o.SenderID[:]); err != nil {
			return err
		}
	}
	if !o.SendingTimeEpochInActingVersion(actingVersion) {
		o.SendingTimeEpoch = o.SendingTimeEpochNullValue()
	} else {
		if err := _m.ReadUint64(_r, &o.SendingTimeEpoch); err != nil {
			return err
		}
	}
	if !o.SecurityGroupInActingVersion(actingVersion) {
		for idx := 0; idx < 6; idx++ {
			o.SecurityGroup[idx] = o.SecurityGroupNullValue()
		}
	} else {
		if err := _m.ReadBytes(_r, o.SecurityGroup[:]); err != nil {
			return err
		}
	}
	if !o.LocationInActingVersion(actingVersion) {
		for idx := 0; idx < 5; idx++ {
			o.Location[idx] = o.LocationNullValue()
		}
	} else {
		if err := _m.ReadBytes(_r, o.Location[:]); err != nil {
			return err
		}
	}
	if !o.SecurityIDInActingVersion(actingVersion) {
		o.SecurityID = o.SecurityIDNullValue()
	} else {
		if err := _m.ReadInt32(_r, &o.SecurityID); err != nil {
			return err
		}
	}
	if o.MassStatusReqTypeInActingVersion(actingVersion) {
		if err := o.MassStatusReqType.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if o.OrdStatusReqTypeInActingVersion(actingVersion) {
		if err := o.OrdStatusReqType.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if o.TimeInForceInActingVersion(actingVersion) {
		if err := o.TimeInForce.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if !o.MarketSegmentIDInActingVersion(actingVersion) {
		o.MarketSegmentID = o.MarketSegmentIDNullValue()
	} else {
		if err := _m.ReadUint8(_r, &o.MarketSegmentID); err != nil {
			return err
		}
	}
	if actingVersion > o.SbeSchemaVersion() && blockLength > o.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-o.SbeBlockLength()))
	}
	if doRangeCheck {
		if err := o.RangeCheck(actingVersion, o.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	return nil
}

func (o *OrderMassStatusRequest530) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if o.PartyDetailsListReqIDInActingVersion(actingVersion) {
		if o.PartyDetailsListReqID < o.PartyDetailsListReqIDMinValue() || o.PartyDetailsListReqID > o.PartyDetailsListReqIDMaxValue() {
			return fmt.Errorf("Range check failed on o.PartyDetailsListReqID (%v < %v > %v)", o.PartyDetailsListReqIDMinValue(), o.PartyDetailsListReqID, o.PartyDetailsListReqIDMaxValue())
		}
	}
	if o.MassStatusReqIDInActingVersion(actingVersion) {
		if o.MassStatusReqID < o.MassStatusReqIDMinValue() || o.MassStatusReqID > o.MassStatusReqIDMaxValue() {
			return fmt.Errorf("Range check failed on o.MassStatusReqID (%v < %v > %v)", o.MassStatusReqIDMinValue(), o.MassStatusReqID, o.MassStatusReqIDMaxValue())
		}
	}
	if err := o.ManualOrderIndicator.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if o.SeqNumInActingVersion(actingVersion) {
		if o.SeqNum < o.SeqNumMinValue() || o.SeqNum > o.SeqNumMaxValue() {
			return fmt.Errorf("Range check failed on o.SeqNum (%v < %v > %v)", o.SeqNumMinValue(), o.SeqNum, o.SeqNumMaxValue())
		}
	}
	if o.SenderIDInActingVersion(actingVersion) {
		for idx := 0; idx < 20; idx++ {
			if o.SenderID[idx] < o.SenderIDMinValue() || o.SenderID[idx] > o.SenderIDMaxValue() {
				return fmt.Errorf("Range check failed on o.SenderID[%d] (%v < %v > %v)", idx, o.SenderIDMinValue(), o.SenderID[idx], o.SenderIDMaxValue())
			}
		}
	}
	if o.SendingTimeEpochInActingVersion(actingVersion) {
		if o.SendingTimeEpoch < o.SendingTimeEpochMinValue() || o.SendingTimeEpoch > o.SendingTimeEpochMaxValue() {
			return fmt.Errorf("Range check failed on o.SendingTimeEpoch (%v < %v > %v)", o.SendingTimeEpochMinValue(), o.SendingTimeEpoch, o.SendingTimeEpochMaxValue())
		}
	}
	if o.SecurityGroupInActingVersion(actingVersion) {
		for idx := 0; idx < 6; idx++ {
			if o.SecurityGroup[idx] < o.SecurityGroupMinValue() || o.SecurityGroup[idx] > o.SecurityGroupMaxValue() {
				return fmt.Errorf("Range check failed on o.SecurityGroup[%d] (%v < %v > %v)", idx, o.SecurityGroupMinValue(), o.SecurityGroup[idx], o.SecurityGroupMaxValue())
			}
		}
	}
	if o.LocationInActingVersion(actingVersion) {
		for idx := 0; idx < 5; idx++ {
			if o.Location[idx] < o.LocationMinValue() || o.Location[idx] > o.LocationMaxValue() {
				return fmt.Errorf("Range check failed on o.Location[%d] (%v < %v > %v)", idx, o.LocationMinValue(), o.Location[idx], o.LocationMaxValue())
			}
		}
	}
	if o.SecurityIDInActingVersion(actingVersion) {
		if o.SecurityID != o.SecurityIDNullValue() && (o.SecurityID < o.SecurityIDMinValue() || o.SecurityID > o.SecurityIDMaxValue()) {
			return fmt.Errorf("Range check failed on o.SecurityID (%v < %v > %v)", o.SecurityIDMinValue(), o.SecurityID, o.SecurityIDMaxValue())
		}
	}
	if err := o.MassStatusReqType.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if err := o.OrdStatusReqType.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if err := o.TimeInForce.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if o.MarketSegmentIDInActingVersion(actingVersion) {
		if o.MarketSegmentID != o.MarketSegmentIDNullValue() && (o.MarketSegmentID < o.MarketSegmentIDMinValue() || o.MarketSegmentID > o.MarketSegmentIDMaxValue()) {
			return fmt.Errorf("Range check failed on o.MarketSegmentID (%v < %v > %v)", o.MarketSegmentIDMinValue(), o.MarketSegmentID, o.MarketSegmentIDMaxValue())
		}
	}
	return nil
}

func OrderMassStatusRequest530Init(o *OrderMassStatusRequest530) {
	for idx := 0; idx < 6; idx++ {
		o.SecurityGroup[idx] = 0
	}
	o.SecurityID = 2147483647
	o.MarketSegmentID = 255
	return
}

func (*OrderMassStatusRequest530) SbeBlockLength() (blockLength uint16) {
	return 68
}

func (*OrderMassStatusRequest530) SbeTemplateId() (templateId uint16) {
	return 530
}

func (*OrderMassStatusRequest530) SbeSchemaId() (schemaId uint16) {
	return 8
}

func (*OrderMassStatusRequest530) SbeSchemaVersion() (schemaVersion uint16) {
	return 5
}

func (*OrderMassStatusRequest530) SbeSemanticType() (semanticType []byte) {
	return []byte("AF")
}

func (*OrderMassStatusRequest530) PartyDetailsListReqIDId() uint16 {
	return 1505
}

func (*OrderMassStatusRequest530) PartyDetailsListReqIDSinceVersion() uint16 {
	return 0
}

func (o *OrderMassStatusRequest530) PartyDetailsListReqIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.PartyDetailsListReqIDSinceVersion()
}

func (*OrderMassStatusRequest530) PartyDetailsListReqIDDeprecated() uint16 {
	return 0
}

func (*OrderMassStatusRequest530) PartyDetailsListReqIDMetaAttribute(meta int) string {
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

func (*OrderMassStatusRequest530) PartyDetailsListReqIDMinValue() uint64 {
	return 0
}

func (*OrderMassStatusRequest530) PartyDetailsListReqIDMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*OrderMassStatusRequest530) PartyDetailsListReqIDNullValue() uint64 {
	return math.MaxUint64
}

func (*OrderMassStatusRequest530) MassStatusReqIDId() uint16 {
	return 584
}

func (*OrderMassStatusRequest530) MassStatusReqIDSinceVersion() uint16 {
	return 0
}

func (o *OrderMassStatusRequest530) MassStatusReqIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.MassStatusReqIDSinceVersion()
}

func (*OrderMassStatusRequest530) MassStatusReqIDDeprecated() uint16 {
	return 0
}

func (*OrderMassStatusRequest530) MassStatusReqIDMetaAttribute(meta int) string {
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

func (*OrderMassStatusRequest530) MassStatusReqIDMinValue() uint64 {
	return 0
}

func (*OrderMassStatusRequest530) MassStatusReqIDMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*OrderMassStatusRequest530) MassStatusReqIDNullValue() uint64 {
	return math.MaxUint64
}

func (*OrderMassStatusRequest530) ManualOrderIndicatorId() uint16 {
	return 1028
}

func (*OrderMassStatusRequest530) ManualOrderIndicatorSinceVersion() uint16 {
	return 0
}

func (o *OrderMassStatusRequest530) ManualOrderIndicatorInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.ManualOrderIndicatorSinceVersion()
}

func (*OrderMassStatusRequest530) ManualOrderIndicatorDeprecated() uint16 {
	return 0
}

func (*OrderMassStatusRequest530) ManualOrderIndicatorMetaAttribute(meta int) string {
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

func (*OrderMassStatusRequest530) SeqNumId() uint16 {
	return 9726
}

func (*OrderMassStatusRequest530) SeqNumSinceVersion() uint16 {
	return 0
}

func (o *OrderMassStatusRequest530) SeqNumInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.SeqNumSinceVersion()
}

func (*OrderMassStatusRequest530) SeqNumDeprecated() uint16 {
	return 0
}

func (*OrderMassStatusRequest530) SeqNumMetaAttribute(meta int) string {
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

func (*OrderMassStatusRequest530) SeqNumMinValue() uint32 {
	return 0
}

func (*OrderMassStatusRequest530) SeqNumMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*OrderMassStatusRequest530) SeqNumNullValue() uint32 {
	return math.MaxUint32
}

func (*OrderMassStatusRequest530) SenderIDId() uint16 {
	return 5392
}

func (*OrderMassStatusRequest530) SenderIDSinceVersion() uint16 {
	return 0
}

func (o *OrderMassStatusRequest530) SenderIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.SenderIDSinceVersion()
}

func (*OrderMassStatusRequest530) SenderIDDeprecated() uint16 {
	return 0
}

func (*OrderMassStatusRequest530) SenderIDMetaAttribute(meta int) string {
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

func (*OrderMassStatusRequest530) SenderIDMinValue() byte {
	return byte(32)
}

func (*OrderMassStatusRequest530) SenderIDMaxValue() byte {
	return byte(126)
}

func (*OrderMassStatusRequest530) SenderIDNullValue() byte {
	return 0
}

func (o *OrderMassStatusRequest530) SenderIDCharacterEncoding() string {
	return "US-ASCII"
}

func (*OrderMassStatusRequest530) SendingTimeEpochId() uint16 {
	return 5297
}

func (*OrderMassStatusRequest530) SendingTimeEpochSinceVersion() uint16 {
	return 0
}

func (o *OrderMassStatusRequest530) SendingTimeEpochInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.SendingTimeEpochSinceVersion()
}

func (*OrderMassStatusRequest530) SendingTimeEpochDeprecated() uint16 {
	return 0
}

func (*OrderMassStatusRequest530) SendingTimeEpochMetaAttribute(meta int) string {
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

func (*OrderMassStatusRequest530) SendingTimeEpochMinValue() uint64 {
	return 0
}

func (*OrderMassStatusRequest530) SendingTimeEpochMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*OrderMassStatusRequest530) SendingTimeEpochNullValue() uint64 {
	return math.MaxUint64
}

func (*OrderMassStatusRequest530) SecurityGroupId() uint16 {
	return 1151
}

func (*OrderMassStatusRequest530) SecurityGroupSinceVersion() uint16 {
	return 0
}

func (o *OrderMassStatusRequest530) SecurityGroupInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.SecurityGroupSinceVersion()
}

func (*OrderMassStatusRequest530) SecurityGroupDeprecated() uint16 {
	return 0
}

func (*OrderMassStatusRequest530) SecurityGroupMetaAttribute(meta int) string {
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

func (*OrderMassStatusRequest530) SecurityGroupMinValue() byte {
	return byte(32)
}

func (*OrderMassStatusRequest530) SecurityGroupMaxValue() byte {
	return byte(126)
}

func (*OrderMassStatusRequest530) SecurityGroupNullValue() byte {
	return 0
}

func (o *OrderMassStatusRequest530) SecurityGroupCharacterEncoding() string {
	return "US-ASCII"
}

func (*OrderMassStatusRequest530) LocationId() uint16 {
	return 9537
}

func (*OrderMassStatusRequest530) LocationSinceVersion() uint16 {
	return 0
}

func (o *OrderMassStatusRequest530) LocationInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.LocationSinceVersion()
}

func (*OrderMassStatusRequest530) LocationDeprecated() uint16 {
	return 0
}

func (*OrderMassStatusRequest530) LocationMetaAttribute(meta int) string {
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

func (*OrderMassStatusRequest530) LocationMinValue() byte {
	return byte(32)
}

func (*OrderMassStatusRequest530) LocationMaxValue() byte {
	return byte(126)
}

func (*OrderMassStatusRequest530) LocationNullValue() byte {
	return 0
}

func (o *OrderMassStatusRequest530) LocationCharacterEncoding() string {
	return "US-ASCII"
}

func (*OrderMassStatusRequest530) SecurityIDId() uint16 {
	return 48
}

func (*OrderMassStatusRequest530) SecurityIDSinceVersion() uint16 {
	return 0
}

func (o *OrderMassStatusRequest530) SecurityIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.SecurityIDSinceVersion()
}

func (*OrderMassStatusRequest530) SecurityIDDeprecated() uint16 {
	return 0
}

func (*OrderMassStatusRequest530) SecurityIDMetaAttribute(meta int) string {
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

func (*OrderMassStatusRequest530) SecurityIDMinValue() int32 {
	return math.MinInt32 + 1
}

func (*OrderMassStatusRequest530) SecurityIDMaxValue() int32 {
	return math.MaxInt32
}

func (*OrderMassStatusRequest530) SecurityIDNullValue() int32 {
	return 2147483647
}

func (*OrderMassStatusRequest530) MassStatusReqTypeId() uint16 {
	return 585
}

func (*OrderMassStatusRequest530) MassStatusReqTypeSinceVersion() uint16 {
	return 0
}

func (o *OrderMassStatusRequest530) MassStatusReqTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.MassStatusReqTypeSinceVersion()
}

func (*OrderMassStatusRequest530) MassStatusReqTypeDeprecated() uint16 {
	return 0
}

func (*OrderMassStatusRequest530) MassStatusReqTypeMetaAttribute(meta int) string {
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

func (*OrderMassStatusRequest530) OrdStatusReqTypeId() uint16 {
	return 5000
}

func (*OrderMassStatusRequest530) OrdStatusReqTypeSinceVersion() uint16 {
	return 0
}

func (o *OrderMassStatusRequest530) OrdStatusReqTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.OrdStatusReqTypeSinceVersion()
}

func (*OrderMassStatusRequest530) OrdStatusReqTypeDeprecated() uint16 {
	return 0
}

func (*OrderMassStatusRequest530) OrdStatusReqTypeMetaAttribute(meta int) string {
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

func (*OrderMassStatusRequest530) TimeInForceId() uint16 {
	return 59
}

func (*OrderMassStatusRequest530) TimeInForceSinceVersion() uint16 {
	return 0
}

func (o *OrderMassStatusRequest530) TimeInForceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.TimeInForceSinceVersion()
}

func (*OrderMassStatusRequest530) TimeInForceDeprecated() uint16 {
	return 0
}

func (*OrderMassStatusRequest530) TimeInForceMetaAttribute(meta int) string {
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

func (*OrderMassStatusRequest530) MarketSegmentIDId() uint16 {
	return 1300
}

func (*OrderMassStatusRequest530) MarketSegmentIDSinceVersion() uint16 {
	return 0
}

func (o *OrderMassStatusRequest530) MarketSegmentIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.MarketSegmentIDSinceVersion()
}

func (*OrderMassStatusRequest530) MarketSegmentIDDeprecated() uint16 {
	return 0
}

func (*OrderMassStatusRequest530) MarketSegmentIDMetaAttribute(meta int) string {
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

func (*OrderMassStatusRequest530) MarketSegmentIDMinValue() uint8 {
	return 0
}

func (*OrderMassStatusRequest530) MarketSegmentIDMaxValue() uint8 {
	return math.MaxUint8 - 1
}

func (*OrderMassStatusRequest530) MarketSegmentIDNullValue() uint8 {
	return 255
}
