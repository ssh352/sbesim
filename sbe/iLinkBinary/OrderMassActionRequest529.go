// Generated SBE (Simple Binary Encoding) message codec

package ilinkbinary

import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
)

type OrderMassActionRequest529 struct {
	PartyDetailsListReqID uint64
	OrderRequestID        uint64
	ManualOrderIndicator  ManualOrdIndReqEnum
	SeqNum                uint32
	SenderID              [20]byte
	MassActionType        [1]byte
	SendingTimeEpoch      uint64
	SecurityGroup         [6]byte
	Location              [5]byte
	SecurityID            int32
	MassActionScope       MassActionScopeEnum
	MarketSegmentID       uint8
	MassCancelRequestType MassCxlReqTypEnum
	Side                  SideNULLEnum
	OrdType               MassActionOrdTypEnum
	TimeInForce           MassCancelTIFEnum
	LiquidityFlag         BooleanNULLEnum
}

func (o *OrderMassActionRequest529) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := o.RangeCheck(o.SbeSchemaVersion(), o.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	if err := _m.WriteUint64(_w, o.PartyDetailsListReqID); err != nil {
		return err
	}
	if err := _m.WriteUint64(_w, o.OrderRequestID); err != nil {
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
	if err := o.MassActionScope.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteUint8(_w, o.MarketSegmentID); err != nil {
		return err
	}
	if err := o.MassCancelRequestType.Encode(_m, _w); err != nil {
		return err
	}
	if err := o.Side.Encode(_m, _w); err != nil {
		return err
	}
	if err := o.OrdType.Encode(_m, _w); err != nil {
		return err
	}
	if err := o.TimeInForce.Encode(_m, _w); err != nil {
		return err
	}
	if err := o.LiquidityFlag.Encode(_m, _w); err != nil {
		return err
	}
	return nil
}

func (o *OrderMassActionRequest529) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	if !o.PartyDetailsListReqIDInActingVersion(actingVersion) {
		o.PartyDetailsListReqID = o.PartyDetailsListReqIDNullValue()
	} else {
		if err := _m.ReadUint64(_r, &o.PartyDetailsListReqID); err != nil {
			return err
		}
	}
	if !o.OrderRequestIDInActingVersion(actingVersion) {
		o.OrderRequestID = o.OrderRequestIDNullValue()
	} else {
		if err := _m.ReadUint64(_r, &o.OrderRequestID); err != nil {
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
	o.MassActionType[0] = 51
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
	if o.MassActionScopeInActingVersion(actingVersion) {
		if err := o.MassActionScope.Decode(_m, _r, actingVersion); err != nil {
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
	if o.MassCancelRequestTypeInActingVersion(actingVersion) {
		if err := o.MassCancelRequestType.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if o.SideInActingVersion(actingVersion) {
		if err := o.Side.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if o.OrdTypeInActingVersion(actingVersion) {
		if err := o.OrdType.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if o.TimeInForceInActingVersion(actingVersion) {
		if err := o.TimeInForce.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if o.LiquidityFlagInActingVersion(actingVersion) {
		if err := o.LiquidityFlag.Decode(_m, _r, actingVersion); err != nil {
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

func (o *OrderMassActionRequest529) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if o.PartyDetailsListReqIDInActingVersion(actingVersion) {
		if o.PartyDetailsListReqID < o.PartyDetailsListReqIDMinValue() || o.PartyDetailsListReqID > o.PartyDetailsListReqIDMaxValue() {
			return fmt.Errorf("Range check failed on o.PartyDetailsListReqID (%v < %v > %v)", o.PartyDetailsListReqIDMinValue(), o.PartyDetailsListReqID, o.PartyDetailsListReqIDMaxValue())
		}
	}
	if o.OrderRequestIDInActingVersion(actingVersion) {
		if o.OrderRequestID < o.OrderRequestIDMinValue() || o.OrderRequestID > o.OrderRequestIDMaxValue() {
			return fmt.Errorf("Range check failed on o.OrderRequestID (%v < %v > %v)", o.OrderRequestIDMinValue(), o.OrderRequestID, o.OrderRequestIDMaxValue())
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
	if err := o.MassActionScope.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if o.MarketSegmentIDInActingVersion(actingVersion) {
		if o.MarketSegmentID != o.MarketSegmentIDNullValue() && (o.MarketSegmentID < o.MarketSegmentIDMinValue() || o.MarketSegmentID > o.MarketSegmentIDMaxValue()) {
			return fmt.Errorf("Range check failed on o.MarketSegmentID (%v < %v > %v)", o.MarketSegmentIDMinValue(), o.MarketSegmentID, o.MarketSegmentIDMaxValue())
		}
	}
	if err := o.MassCancelRequestType.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if err := o.Side.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if err := o.OrdType.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if err := o.TimeInForce.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if err := o.LiquidityFlag.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	return nil
}

func OrderMassActionRequest529Init(o *OrderMassActionRequest529) {
	o.MassActionType[0] = 51
	for idx := 0; idx < 6; idx++ {
		o.SecurityGroup[idx] = 0
	}
	o.SecurityID = 2147483647
	o.MarketSegmentID = 255
	return
}

func (*OrderMassActionRequest529) SbeBlockLength() (blockLength uint16) {
	return 71
}

func (*OrderMassActionRequest529) SbeTemplateId() (templateId uint16) {
	return 529
}

func (*OrderMassActionRequest529) SbeSchemaId() (schemaId uint16) {
	return 8
}

func (*OrderMassActionRequest529) SbeSchemaVersion() (schemaVersion uint16) {
	return 5
}

func (*OrderMassActionRequest529) SbeSemanticType() (semanticType []byte) {
	return []byte("CA")
}

func (*OrderMassActionRequest529) PartyDetailsListReqIDId() uint16 {
	return 1505
}

func (*OrderMassActionRequest529) PartyDetailsListReqIDSinceVersion() uint16 {
	return 0
}

func (o *OrderMassActionRequest529) PartyDetailsListReqIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.PartyDetailsListReqIDSinceVersion()
}

func (*OrderMassActionRequest529) PartyDetailsListReqIDDeprecated() uint16 {
	return 0
}

func (*OrderMassActionRequest529) PartyDetailsListReqIDMetaAttribute(meta int) string {
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

func (*OrderMassActionRequest529) PartyDetailsListReqIDMinValue() uint64 {
	return 0
}

func (*OrderMassActionRequest529) PartyDetailsListReqIDMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*OrderMassActionRequest529) PartyDetailsListReqIDNullValue() uint64 {
	return math.MaxUint64
}

func (*OrderMassActionRequest529) OrderRequestIDId() uint16 {
	return 2422
}

func (*OrderMassActionRequest529) OrderRequestIDSinceVersion() uint16 {
	return 0
}

func (o *OrderMassActionRequest529) OrderRequestIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.OrderRequestIDSinceVersion()
}

func (*OrderMassActionRequest529) OrderRequestIDDeprecated() uint16 {
	return 0
}

func (*OrderMassActionRequest529) OrderRequestIDMetaAttribute(meta int) string {
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

func (*OrderMassActionRequest529) OrderRequestIDMinValue() uint64 {
	return 0
}

func (*OrderMassActionRequest529) OrderRequestIDMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*OrderMassActionRequest529) OrderRequestIDNullValue() uint64 {
	return math.MaxUint64
}

func (*OrderMassActionRequest529) ManualOrderIndicatorId() uint16 {
	return 1028
}

func (*OrderMassActionRequest529) ManualOrderIndicatorSinceVersion() uint16 {
	return 0
}

func (o *OrderMassActionRequest529) ManualOrderIndicatorInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.ManualOrderIndicatorSinceVersion()
}

func (*OrderMassActionRequest529) ManualOrderIndicatorDeprecated() uint16 {
	return 0
}

func (*OrderMassActionRequest529) ManualOrderIndicatorMetaAttribute(meta int) string {
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

func (*OrderMassActionRequest529) SeqNumId() uint16 {
	return 9726
}

func (*OrderMassActionRequest529) SeqNumSinceVersion() uint16 {
	return 0
}

func (o *OrderMassActionRequest529) SeqNumInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.SeqNumSinceVersion()
}

func (*OrderMassActionRequest529) SeqNumDeprecated() uint16 {
	return 0
}

func (*OrderMassActionRequest529) SeqNumMetaAttribute(meta int) string {
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

func (*OrderMassActionRequest529) SeqNumMinValue() uint32 {
	return 0
}

func (*OrderMassActionRequest529) SeqNumMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*OrderMassActionRequest529) SeqNumNullValue() uint32 {
	return math.MaxUint32
}

func (*OrderMassActionRequest529) SenderIDId() uint16 {
	return 5392
}

func (*OrderMassActionRequest529) SenderIDSinceVersion() uint16 {
	return 0
}

func (o *OrderMassActionRequest529) SenderIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.SenderIDSinceVersion()
}

func (*OrderMassActionRequest529) SenderIDDeprecated() uint16 {
	return 0
}

func (*OrderMassActionRequest529) SenderIDMetaAttribute(meta int) string {
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

func (*OrderMassActionRequest529) SenderIDMinValue() byte {
	return byte(32)
}

func (*OrderMassActionRequest529) SenderIDMaxValue() byte {
	return byte(126)
}

func (*OrderMassActionRequest529) SenderIDNullValue() byte {
	return 0
}

func (o *OrderMassActionRequest529) SenderIDCharacterEncoding() string {
	return "US-ASCII"
}

func (*OrderMassActionRequest529) MassActionTypeId() uint16 {
	return 1373
}

func (*OrderMassActionRequest529) MassActionTypeSinceVersion() uint16 {
	return 0
}

func (o *OrderMassActionRequest529) MassActionTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.MassActionTypeSinceVersion()
}

func (*OrderMassActionRequest529) MassActionTypeDeprecated() uint16 {
	return 0
}

func (*OrderMassActionRequest529) MassActionTypeMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return "char"
	case 4:
		return "constant"
	}
	return ""
}

func (*OrderMassActionRequest529) MassActionTypeMinValue() byte {
	return byte(32)
}

func (*OrderMassActionRequest529) MassActionTypeMaxValue() byte {
	return byte(126)
}

func (*OrderMassActionRequest529) MassActionTypeNullValue() byte {
	return 0
}

func (*OrderMassActionRequest529) SendingTimeEpochId() uint16 {
	return 5297
}

func (*OrderMassActionRequest529) SendingTimeEpochSinceVersion() uint16 {
	return 0
}

func (o *OrderMassActionRequest529) SendingTimeEpochInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.SendingTimeEpochSinceVersion()
}

func (*OrderMassActionRequest529) SendingTimeEpochDeprecated() uint16 {
	return 0
}

func (*OrderMassActionRequest529) SendingTimeEpochMetaAttribute(meta int) string {
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

func (*OrderMassActionRequest529) SendingTimeEpochMinValue() uint64 {
	return 0
}

func (*OrderMassActionRequest529) SendingTimeEpochMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*OrderMassActionRequest529) SendingTimeEpochNullValue() uint64 {
	return math.MaxUint64
}

func (*OrderMassActionRequest529) SecurityGroupId() uint16 {
	return 1151
}

func (*OrderMassActionRequest529) SecurityGroupSinceVersion() uint16 {
	return 0
}

func (o *OrderMassActionRequest529) SecurityGroupInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.SecurityGroupSinceVersion()
}

func (*OrderMassActionRequest529) SecurityGroupDeprecated() uint16 {
	return 0
}

func (*OrderMassActionRequest529) SecurityGroupMetaAttribute(meta int) string {
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

func (*OrderMassActionRequest529) SecurityGroupMinValue() byte {
	return byte(32)
}

func (*OrderMassActionRequest529) SecurityGroupMaxValue() byte {
	return byte(126)
}

func (*OrderMassActionRequest529) SecurityGroupNullValue() byte {
	return 0
}

func (o *OrderMassActionRequest529) SecurityGroupCharacterEncoding() string {
	return "US-ASCII"
}

func (*OrderMassActionRequest529) LocationId() uint16 {
	return 9537
}

func (*OrderMassActionRequest529) LocationSinceVersion() uint16 {
	return 0
}

func (o *OrderMassActionRequest529) LocationInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.LocationSinceVersion()
}

func (*OrderMassActionRequest529) LocationDeprecated() uint16 {
	return 0
}

func (*OrderMassActionRequest529) LocationMetaAttribute(meta int) string {
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

func (*OrderMassActionRequest529) LocationMinValue() byte {
	return byte(32)
}

func (*OrderMassActionRequest529) LocationMaxValue() byte {
	return byte(126)
}

func (*OrderMassActionRequest529) LocationNullValue() byte {
	return 0
}

func (o *OrderMassActionRequest529) LocationCharacterEncoding() string {
	return "US-ASCII"
}

func (*OrderMassActionRequest529) SecurityIDId() uint16 {
	return 48
}

func (*OrderMassActionRequest529) SecurityIDSinceVersion() uint16 {
	return 0
}

func (o *OrderMassActionRequest529) SecurityIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.SecurityIDSinceVersion()
}

func (*OrderMassActionRequest529) SecurityIDDeprecated() uint16 {
	return 0
}

func (*OrderMassActionRequest529) SecurityIDMetaAttribute(meta int) string {
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

func (*OrderMassActionRequest529) SecurityIDMinValue() int32 {
	return math.MinInt32 + 1
}

func (*OrderMassActionRequest529) SecurityIDMaxValue() int32 {
	return math.MaxInt32
}

func (*OrderMassActionRequest529) SecurityIDNullValue() int32 {
	return 2147483647
}

func (*OrderMassActionRequest529) MassActionScopeId() uint16 {
	return 1374
}

func (*OrderMassActionRequest529) MassActionScopeSinceVersion() uint16 {
	return 0
}

func (o *OrderMassActionRequest529) MassActionScopeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.MassActionScopeSinceVersion()
}

func (*OrderMassActionRequest529) MassActionScopeDeprecated() uint16 {
	return 0
}

func (*OrderMassActionRequest529) MassActionScopeMetaAttribute(meta int) string {
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

func (*OrderMassActionRequest529) MarketSegmentIDId() uint16 {
	return 1300
}

func (*OrderMassActionRequest529) MarketSegmentIDSinceVersion() uint16 {
	return 0
}

func (o *OrderMassActionRequest529) MarketSegmentIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.MarketSegmentIDSinceVersion()
}

func (*OrderMassActionRequest529) MarketSegmentIDDeprecated() uint16 {
	return 0
}

func (*OrderMassActionRequest529) MarketSegmentIDMetaAttribute(meta int) string {
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

func (*OrderMassActionRequest529) MarketSegmentIDMinValue() uint8 {
	return 0
}

func (*OrderMassActionRequest529) MarketSegmentIDMaxValue() uint8 {
	return math.MaxUint8 - 1
}

func (*OrderMassActionRequest529) MarketSegmentIDNullValue() uint8 {
	return 255
}

func (*OrderMassActionRequest529) MassCancelRequestTypeId() uint16 {
	return 6115
}

func (*OrderMassActionRequest529) MassCancelRequestTypeSinceVersion() uint16 {
	return 0
}

func (o *OrderMassActionRequest529) MassCancelRequestTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.MassCancelRequestTypeSinceVersion()
}

func (*OrderMassActionRequest529) MassCancelRequestTypeDeprecated() uint16 {
	return 0
}

func (*OrderMassActionRequest529) MassCancelRequestTypeMetaAttribute(meta int) string {
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

func (*OrderMassActionRequest529) SideId() uint16 {
	return 54
}

func (*OrderMassActionRequest529) SideSinceVersion() uint16 {
	return 0
}

func (o *OrderMassActionRequest529) SideInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.SideSinceVersion()
}

func (*OrderMassActionRequest529) SideDeprecated() uint16 {
	return 0
}

func (*OrderMassActionRequest529) SideMetaAttribute(meta int) string {
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

func (*OrderMassActionRequest529) OrdTypeId() uint16 {
	return 40
}

func (*OrderMassActionRequest529) OrdTypeSinceVersion() uint16 {
	return 0
}

func (o *OrderMassActionRequest529) OrdTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.OrdTypeSinceVersion()
}

func (*OrderMassActionRequest529) OrdTypeDeprecated() uint16 {
	return 0
}

func (*OrderMassActionRequest529) OrdTypeMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return "char"
	case 4:
		return "required"
	}
	return ""
}

func (*OrderMassActionRequest529) TimeInForceId() uint16 {
	return 59
}

func (*OrderMassActionRequest529) TimeInForceSinceVersion() uint16 {
	return 0
}

func (o *OrderMassActionRequest529) TimeInForceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.TimeInForceSinceVersion()
}

func (*OrderMassActionRequest529) TimeInForceDeprecated() uint16 {
	return 0
}

func (*OrderMassActionRequest529) TimeInForceMetaAttribute(meta int) string {
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

func (*OrderMassActionRequest529) LiquidityFlagId() uint16 {
	return 9373
}

func (*OrderMassActionRequest529) LiquidityFlagSinceVersion() uint16 {
	return 0
}

func (o *OrderMassActionRequest529) LiquidityFlagInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.LiquidityFlagSinceVersion()
}

func (*OrderMassActionRequest529) LiquidityFlagDeprecated() uint16 {
	return 0
}

func (*OrderMassActionRequest529) LiquidityFlagMetaAttribute(meta int) string {
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
