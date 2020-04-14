// Generated SBE (Simple Binary Encoding) message codec

package ilinkbinary

import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
)

type OrderMassActionReport562 struct {
	SeqNum                 uint32
	UUID                   uint64
	SenderID               [20]byte
	PartyDetailsListReqID  uint64
	TransactTime           uint64
	SendingTimeEpoch       uint64
	OrderRequestID         uint64
	MassActionReportID     uint64
	MassActionType         [1]byte
	SecurityGroup          [6]byte
	Location               [5]byte
	SecurityID             int32
	DelayDuration          uint16
	MassActionResponse     MassActionResponseEnum
	ManualOrderIndicator   ManualOrdIndReqEnum
	MassActionScope        MassActionScopeEnum
	TotalAffectedOrders    uint32
	LastFragment           BooleanFlagEnum
	MassActionRejectReason uint8
	MarketSegmentID        uint8
	MassCancelRequestType  MassCxlReqTypEnum
	Side                   SideNULLEnum
	OrdType                MassActionOrdTypEnum
	TimeInForce            MassCancelTIFEnum
	SplitMsg               SplitMsgEnum
	LiquidityFlag          BooleanNULLEnum
	PossRetransFlag        BooleanFlagEnum
	DelayToTime            uint64
	NoAffectedOrders       []OrderMassActionReport562NoAffectedOrders
}
type OrderMassActionReport562NoAffectedOrders struct {
	OrigCIOrdID     [20]byte
	AffectedOrderID uint64
	CxlQuantity     uint32
}

func (o *OrderMassActionReport562) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := o.RangeCheck(o.SbeSchemaVersion(), o.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	if err := _m.WriteUint32(_w, o.SeqNum); err != nil {
		return err
	}
	if err := _m.WriteUint64(_w, o.UUID); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, o.SenderID[:]); err != nil {
		return err
	}
	if err := _m.WriteUint64(_w, o.PartyDetailsListReqID); err != nil {
		return err
	}
	if err := _m.WriteUint64(_w, o.TransactTime); err != nil {
		return err
	}
	if err := _m.WriteUint64(_w, o.SendingTimeEpoch); err != nil {
		return err
	}
	if err := _m.WriteUint64(_w, o.OrderRequestID); err != nil {
		return err
	}
	if err := _m.WriteUint64(_w, o.MassActionReportID); err != nil {
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
	if err := _m.WriteUint16(_w, o.DelayDuration); err != nil {
		return err
	}
	if err := o.MassActionResponse.Encode(_m, _w); err != nil {
		return err
	}
	if err := o.ManualOrderIndicator.Encode(_m, _w); err != nil {
		return err
	}
	if err := o.MassActionScope.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteUint32(_w, o.TotalAffectedOrders); err != nil {
		return err
	}
	if err := o.LastFragment.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteUint8(_w, o.MassActionRejectReason); err != nil {
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
	if err := o.SplitMsg.Encode(_m, _w); err != nil {
		return err
	}
	if err := o.LiquidityFlag.Encode(_m, _w); err != nil {
		return err
	}
	if err := o.PossRetransFlag.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteUint64(_w, o.DelayToTime); err != nil {
		return err
	}
	var NoAffectedOrdersBlockLength uint16 = 32
	if err := _m.WriteUint16(_w, NoAffectedOrdersBlockLength); err != nil {
		return err
	}
	var NoAffectedOrdersNumInGroup uint8 = uint8(len(o.NoAffectedOrders))
	if err := _m.WriteUint8(_w, NoAffectedOrdersNumInGroup); err != nil {
		return err
	}
	for _, prop := range o.NoAffectedOrders {
		if err := prop.Encode(_m, _w); err != nil {
			return err
		}
	}
	return nil
}

func (o *OrderMassActionReport562) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	if !o.SeqNumInActingVersion(actingVersion) {
		o.SeqNum = o.SeqNumNullValue()
	} else {
		if err := _m.ReadUint32(_r, &o.SeqNum); err != nil {
			return err
		}
	}
	if !o.UUIDInActingVersion(actingVersion) {
		o.UUID = o.UUIDNullValue()
	} else {
		if err := _m.ReadUint64(_r, &o.UUID); err != nil {
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
	if !o.PartyDetailsListReqIDInActingVersion(actingVersion) {
		o.PartyDetailsListReqID = o.PartyDetailsListReqIDNullValue()
	} else {
		if err := _m.ReadUint64(_r, &o.PartyDetailsListReqID); err != nil {
			return err
		}
	}
	if !o.TransactTimeInActingVersion(actingVersion) {
		o.TransactTime = o.TransactTimeNullValue()
	} else {
		if err := _m.ReadUint64(_r, &o.TransactTime); err != nil {
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
	if !o.OrderRequestIDInActingVersion(actingVersion) {
		o.OrderRequestID = o.OrderRequestIDNullValue()
	} else {
		if err := _m.ReadUint64(_r, &o.OrderRequestID); err != nil {
			return err
		}
	}
	if !o.MassActionReportIDInActingVersion(actingVersion) {
		o.MassActionReportID = o.MassActionReportIDNullValue()
	} else {
		if err := _m.ReadUint64(_r, &o.MassActionReportID); err != nil {
			return err
		}
	}
	o.MassActionType[0] = 51
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
	if !o.DelayDurationInActingVersion(actingVersion) {
		o.DelayDuration = o.DelayDurationNullValue()
	} else {
		if err := _m.ReadUint16(_r, &o.DelayDuration); err != nil {
			return err
		}
	}
	if o.MassActionResponseInActingVersion(actingVersion) {
		if err := o.MassActionResponse.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if o.ManualOrderIndicatorInActingVersion(actingVersion) {
		if err := o.ManualOrderIndicator.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if o.MassActionScopeInActingVersion(actingVersion) {
		if err := o.MassActionScope.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if !o.TotalAffectedOrdersInActingVersion(actingVersion) {
		o.TotalAffectedOrders = o.TotalAffectedOrdersNullValue()
	} else {
		if err := _m.ReadUint32(_r, &o.TotalAffectedOrders); err != nil {
			return err
		}
	}
	if o.LastFragmentInActingVersion(actingVersion) {
		if err := o.LastFragment.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if !o.MassActionRejectReasonInActingVersion(actingVersion) {
		o.MassActionRejectReason = o.MassActionRejectReasonNullValue()
	} else {
		if err := _m.ReadUint8(_r, &o.MassActionRejectReason); err != nil {
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
	if o.SplitMsgInActingVersion(actingVersion) {
		if err := o.SplitMsg.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if o.LiquidityFlagInActingVersion(actingVersion) {
		if err := o.LiquidityFlag.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if o.PossRetransFlagInActingVersion(actingVersion) {
		if err := o.PossRetransFlag.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if !o.DelayToTimeInActingVersion(actingVersion) {
		o.DelayToTime = o.DelayToTimeNullValue()
	} else {
		if err := _m.ReadUint64(_r, &o.DelayToTime); err != nil {
			return err
		}
	}
	if actingVersion > o.SbeSchemaVersion() && blockLength > o.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-o.SbeBlockLength()))
	}

	if o.NoAffectedOrdersInActingVersion(actingVersion) {
		var NoAffectedOrdersBlockLength uint16
		if err := _m.ReadUint16(_r, &NoAffectedOrdersBlockLength); err != nil {
			return err
		}
		var NoAffectedOrdersNumInGroup uint8
		if err := _m.ReadUint8(_r, &NoAffectedOrdersNumInGroup); err != nil {
			return err
		}
		if cap(o.NoAffectedOrders) < int(NoAffectedOrdersNumInGroup) {
			o.NoAffectedOrders = make([]OrderMassActionReport562NoAffectedOrders, NoAffectedOrdersNumInGroup)
		}
		for i, _ := range o.NoAffectedOrders {
			if err := o.NoAffectedOrders[i].Decode(_m, _r, actingVersion, uint(NoAffectedOrdersBlockLength)); err != nil {
				return err
			}
		}
	}
	if doRangeCheck {
		if err := o.RangeCheck(actingVersion, o.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	return nil
}

func (o *OrderMassActionReport562) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if o.SeqNumInActingVersion(actingVersion) {
		if o.SeqNum < o.SeqNumMinValue() || o.SeqNum > o.SeqNumMaxValue() {
			return fmt.Errorf("Range check failed on o.SeqNum (%v < %v > %v)", o.SeqNumMinValue(), o.SeqNum, o.SeqNumMaxValue())
		}
	}
	if o.UUIDInActingVersion(actingVersion) {
		if o.UUID < o.UUIDMinValue() || o.UUID > o.UUIDMaxValue() {
			return fmt.Errorf("Range check failed on o.UUID (%v < %v > %v)", o.UUIDMinValue(), o.UUID, o.UUIDMaxValue())
		}
	}
	if o.SenderIDInActingVersion(actingVersion) {
		for idx := 0; idx < 20; idx++ {
			if o.SenderID[idx] < o.SenderIDMinValue() || o.SenderID[idx] > o.SenderIDMaxValue() {
				return fmt.Errorf("Range check failed on o.SenderID[%d] (%v < %v > %v)", idx, o.SenderIDMinValue(), o.SenderID[idx], o.SenderIDMaxValue())
			}
		}
	}
	if o.PartyDetailsListReqIDInActingVersion(actingVersion) {
		if o.PartyDetailsListReqID < o.PartyDetailsListReqIDMinValue() || o.PartyDetailsListReqID > o.PartyDetailsListReqIDMaxValue() {
			return fmt.Errorf("Range check failed on o.PartyDetailsListReqID (%v < %v > %v)", o.PartyDetailsListReqIDMinValue(), o.PartyDetailsListReqID, o.PartyDetailsListReqIDMaxValue())
		}
	}
	if o.TransactTimeInActingVersion(actingVersion) {
		if o.TransactTime < o.TransactTimeMinValue() || o.TransactTime > o.TransactTimeMaxValue() {
			return fmt.Errorf("Range check failed on o.TransactTime (%v < %v > %v)", o.TransactTimeMinValue(), o.TransactTime, o.TransactTimeMaxValue())
		}
	}
	if o.SendingTimeEpochInActingVersion(actingVersion) {
		if o.SendingTimeEpoch < o.SendingTimeEpochMinValue() || o.SendingTimeEpoch > o.SendingTimeEpochMaxValue() {
			return fmt.Errorf("Range check failed on o.SendingTimeEpoch (%v < %v > %v)", o.SendingTimeEpochMinValue(), o.SendingTimeEpoch, o.SendingTimeEpochMaxValue())
		}
	}
	if o.OrderRequestIDInActingVersion(actingVersion) {
		if o.OrderRequestID < o.OrderRequestIDMinValue() || o.OrderRequestID > o.OrderRequestIDMaxValue() {
			return fmt.Errorf("Range check failed on o.OrderRequestID (%v < %v > %v)", o.OrderRequestIDMinValue(), o.OrderRequestID, o.OrderRequestIDMaxValue())
		}
	}
	if o.MassActionReportIDInActingVersion(actingVersion) {
		if o.MassActionReportID < o.MassActionReportIDMinValue() || o.MassActionReportID > o.MassActionReportIDMaxValue() {
			return fmt.Errorf("Range check failed on o.MassActionReportID (%v < %v > %v)", o.MassActionReportIDMinValue(), o.MassActionReportID, o.MassActionReportIDMaxValue())
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
	if o.DelayDurationInActingVersion(actingVersion) {
		if o.DelayDuration != o.DelayDurationNullValue() && (o.DelayDuration < o.DelayDurationMinValue() || o.DelayDuration > o.DelayDurationMaxValue()) {
			return fmt.Errorf("Range check failed on o.DelayDuration (%v < %v > %v)", o.DelayDurationMinValue(), o.DelayDuration, o.DelayDurationMaxValue())
		}
	}
	if err := o.MassActionResponse.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if err := o.ManualOrderIndicator.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if err := o.MassActionScope.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if o.TotalAffectedOrdersInActingVersion(actingVersion) {
		if o.TotalAffectedOrders < o.TotalAffectedOrdersMinValue() || o.TotalAffectedOrders > o.TotalAffectedOrdersMaxValue() {
			return fmt.Errorf("Range check failed on o.TotalAffectedOrders (%v < %v > %v)", o.TotalAffectedOrdersMinValue(), o.TotalAffectedOrders, o.TotalAffectedOrdersMaxValue())
		}
	}
	if err := o.LastFragment.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if o.MassActionRejectReasonInActingVersion(actingVersion) {
		if o.MassActionRejectReason != o.MassActionRejectReasonNullValue() && (o.MassActionRejectReason < o.MassActionRejectReasonMinValue() || o.MassActionRejectReason > o.MassActionRejectReasonMaxValue()) {
			return fmt.Errorf("Range check failed on o.MassActionRejectReason (%v < %v > %v)", o.MassActionRejectReasonMinValue(), o.MassActionRejectReason, o.MassActionRejectReasonMaxValue())
		}
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
	if err := o.SplitMsg.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if err := o.LiquidityFlag.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if err := o.PossRetransFlag.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if o.DelayToTimeInActingVersion(actingVersion) {
		if o.DelayToTime != o.DelayToTimeNullValue() && (o.DelayToTime < o.DelayToTimeMinValue() || o.DelayToTime > o.DelayToTimeMaxValue()) {
			return fmt.Errorf("Range check failed on o.DelayToTime (%v < %v > %v)", o.DelayToTimeMinValue(), o.DelayToTime, o.DelayToTimeMaxValue())
		}
	}
	for _, prop := range o.NoAffectedOrders {
		if err := prop.RangeCheck(actingVersion, schemaVersion); err != nil {
			return err
		}
	}
	return nil
}

func OrderMassActionReport562Init(o *OrderMassActionReport562) {
	o.MassActionType[0] = 51
	for idx := 0; idx < 6; idx++ {
		o.SecurityGroup[idx] = 0
	}
	o.SecurityID = 2147483647
	o.DelayDuration = 65535
	o.MassActionRejectReason = 255
	o.MarketSegmentID = 255
	o.DelayToTime = 18446744073709551615
	return
}

func (o *OrderMassActionReport562NoAffectedOrders) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteBytes(_w, o.OrigCIOrdID[:]); err != nil {
		return err
	}
	if err := _m.WriteUint64(_w, o.AffectedOrderID); err != nil {
		return err
	}
	if err := _m.WriteUint32(_w, o.CxlQuantity); err != nil {
		return err
	}
	return nil
}

func (o *OrderMassActionReport562NoAffectedOrders) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint) error {
	if !o.OrigCIOrdIDInActingVersion(actingVersion) {
		for idx := 0; idx < 20; idx++ {
			o.OrigCIOrdID[idx] = o.OrigCIOrdIDNullValue()
		}
	} else {
		if err := _m.ReadBytes(_r, o.OrigCIOrdID[:]); err != nil {
			return err
		}
	}
	if !o.AffectedOrderIDInActingVersion(actingVersion) {
		o.AffectedOrderID = o.AffectedOrderIDNullValue()
	} else {
		if err := _m.ReadUint64(_r, &o.AffectedOrderID); err != nil {
			return err
		}
	}
	if !o.CxlQuantityInActingVersion(actingVersion) {
		o.CxlQuantity = o.CxlQuantityNullValue()
	} else {
		if err := _m.ReadUint32(_r, &o.CxlQuantity); err != nil {
			return err
		}
	}
	if actingVersion > o.SbeSchemaVersion() && blockLength > o.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-o.SbeBlockLength()))
	}
	return nil
}

func (o *OrderMassActionReport562NoAffectedOrders) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if o.OrigCIOrdIDInActingVersion(actingVersion) {
		for idx := 0; idx < 20; idx++ {
			if o.OrigCIOrdID[idx] < o.OrigCIOrdIDMinValue() || o.OrigCIOrdID[idx] > o.OrigCIOrdIDMaxValue() {
				return fmt.Errorf("Range check failed on o.OrigCIOrdID[%d] (%v < %v > %v)", idx, o.OrigCIOrdIDMinValue(), o.OrigCIOrdID[idx], o.OrigCIOrdIDMaxValue())
			}
		}
	}
	if o.AffectedOrderIDInActingVersion(actingVersion) {
		if o.AffectedOrderID < o.AffectedOrderIDMinValue() || o.AffectedOrderID > o.AffectedOrderIDMaxValue() {
			return fmt.Errorf("Range check failed on o.AffectedOrderID (%v < %v > %v)", o.AffectedOrderIDMinValue(), o.AffectedOrderID, o.AffectedOrderIDMaxValue())
		}
	}
	if o.CxlQuantityInActingVersion(actingVersion) {
		if o.CxlQuantity < o.CxlQuantityMinValue() || o.CxlQuantity > o.CxlQuantityMaxValue() {
			return fmt.Errorf("Range check failed on o.CxlQuantity (%v < %v > %v)", o.CxlQuantityMinValue(), o.CxlQuantity, o.CxlQuantityMaxValue())
		}
	}
	return nil
}

func OrderMassActionReport562NoAffectedOrdersInit(o *OrderMassActionReport562NoAffectedOrders) {
	return
}

func (*OrderMassActionReport562) SbeBlockLength() (blockLength uint16) {
	return 114
}

func (*OrderMassActionReport562) SbeTemplateId() (templateId uint16) {
	return 562
}

func (*OrderMassActionReport562) SbeSchemaId() (schemaId uint16) {
	return 8
}

func (*OrderMassActionReport562) SbeSchemaVersion() (schemaVersion uint16) {
	return 5
}

func (*OrderMassActionReport562) SbeSemanticType() (semanticType []byte) {
	return []byte("BZ")
}

func (*OrderMassActionReport562) SeqNumId() uint16 {
	return 9726
}

func (*OrderMassActionReport562) SeqNumSinceVersion() uint16 {
	return 0
}

func (o *OrderMassActionReport562) SeqNumInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.SeqNumSinceVersion()
}

func (*OrderMassActionReport562) SeqNumDeprecated() uint16 {
	return 0
}

func (*OrderMassActionReport562) SeqNumMetaAttribute(meta int) string {
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

func (*OrderMassActionReport562) SeqNumMinValue() uint32 {
	return 0
}

func (*OrderMassActionReport562) SeqNumMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*OrderMassActionReport562) SeqNumNullValue() uint32 {
	return math.MaxUint32
}

func (*OrderMassActionReport562) UUIDId() uint16 {
	return 39001
}

func (*OrderMassActionReport562) UUIDSinceVersion() uint16 {
	return 0
}

func (o *OrderMassActionReport562) UUIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.UUIDSinceVersion()
}

func (*OrderMassActionReport562) UUIDDeprecated() uint16 {
	return 0
}

func (*OrderMassActionReport562) UUIDMetaAttribute(meta int) string {
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

func (*OrderMassActionReport562) UUIDMinValue() uint64 {
	return 0
}

func (*OrderMassActionReport562) UUIDMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*OrderMassActionReport562) UUIDNullValue() uint64 {
	return math.MaxUint64
}

func (*OrderMassActionReport562) SenderIDId() uint16 {
	return 5392
}

func (*OrderMassActionReport562) SenderIDSinceVersion() uint16 {
	return 0
}

func (o *OrderMassActionReport562) SenderIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.SenderIDSinceVersion()
}

func (*OrderMassActionReport562) SenderIDDeprecated() uint16 {
	return 0
}

func (*OrderMassActionReport562) SenderIDMetaAttribute(meta int) string {
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

func (*OrderMassActionReport562) SenderIDMinValue() byte {
	return byte(32)
}

func (*OrderMassActionReport562) SenderIDMaxValue() byte {
	return byte(126)
}

func (*OrderMassActionReport562) SenderIDNullValue() byte {
	return 0
}

func (o *OrderMassActionReport562) SenderIDCharacterEncoding() string {
	return "US-ASCII"
}

func (*OrderMassActionReport562) PartyDetailsListReqIDId() uint16 {
	return 1505
}

func (*OrderMassActionReport562) PartyDetailsListReqIDSinceVersion() uint16 {
	return 0
}

func (o *OrderMassActionReport562) PartyDetailsListReqIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.PartyDetailsListReqIDSinceVersion()
}

func (*OrderMassActionReport562) PartyDetailsListReqIDDeprecated() uint16 {
	return 0
}

func (*OrderMassActionReport562) PartyDetailsListReqIDMetaAttribute(meta int) string {
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

func (*OrderMassActionReport562) PartyDetailsListReqIDMinValue() uint64 {
	return 0
}

func (*OrderMassActionReport562) PartyDetailsListReqIDMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*OrderMassActionReport562) PartyDetailsListReqIDNullValue() uint64 {
	return math.MaxUint64
}

func (*OrderMassActionReport562) TransactTimeId() uint16 {
	return 60
}

func (*OrderMassActionReport562) TransactTimeSinceVersion() uint16 {
	return 0
}

func (o *OrderMassActionReport562) TransactTimeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.TransactTimeSinceVersion()
}

func (*OrderMassActionReport562) TransactTimeDeprecated() uint16 {
	return 0
}

func (*OrderMassActionReport562) TransactTimeMetaAttribute(meta int) string {
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

func (*OrderMassActionReport562) TransactTimeMinValue() uint64 {
	return 0
}

func (*OrderMassActionReport562) TransactTimeMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*OrderMassActionReport562) TransactTimeNullValue() uint64 {
	return math.MaxUint64
}

func (*OrderMassActionReport562) SendingTimeEpochId() uint16 {
	return 5297
}

func (*OrderMassActionReport562) SendingTimeEpochSinceVersion() uint16 {
	return 0
}

func (o *OrderMassActionReport562) SendingTimeEpochInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.SendingTimeEpochSinceVersion()
}

func (*OrderMassActionReport562) SendingTimeEpochDeprecated() uint16 {
	return 0
}

func (*OrderMassActionReport562) SendingTimeEpochMetaAttribute(meta int) string {
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

func (*OrderMassActionReport562) SendingTimeEpochMinValue() uint64 {
	return 0
}

func (*OrderMassActionReport562) SendingTimeEpochMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*OrderMassActionReport562) SendingTimeEpochNullValue() uint64 {
	return math.MaxUint64
}

func (*OrderMassActionReport562) OrderRequestIDId() uint16 {
	return 2422
}

func (*OrderMassActionReport562) OrderRequestIDSinceVersion() uint16 {
	return 0
}

func (o *OrderMassActionReport562) OrderRequestIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.OrderRequestIDSinceVersion()
}

func (*OrderMassActionReport562) OrderRequestIDDeprecated() uint16 {
	return 0
}

func (*OrderMassActionReport562) OrderRequestIDMetaAttribute(meta int) string {
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

func (*OrderMassActionReport562) OrderRequestIDMinValue() uint64 {
	return 0
}

func (*OrderMassActionReport562) OrderRequestIDMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*OrderMassActionReport562) OrderRequestIDNullValue() uint64 {
	return math.MaxUint64
}

func (*OrderMassActionReport562) MassActionReportIDId() uint16 {
	return 1369
}

func (*OrderMassActionReport562) MassActionReportIDSinceVersion() uint16 {
	return 0
}

func (o *OrderMassActionReport562) MassActionReportIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.MassActionReportIDSinceVersion()
}

func (*OrderMassActionReport562) MassActionReportIDDeprecated() uint16 {
	return 0
}

func (*OrderMassActionReport562) MassActionReportIDMetaAttribute(meta int) string {
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

func (*OrderMassActionReport562) MassActionReportIDMinValue() uint64 {
	return 0
}

func (*OrderMassActionReport562) MassActionReportIDMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*OrderMassActionReport562) MassActionReportIDNullValue() uint64 {
	return math.MaxUint64
}

func (*OrderMassActionReport562) MassActionTypeId() uint16 {
	return 1373
}

func (*OrderMassActionReport562) MassActionTypeSinceVersion() uint16 {
	return 0
}

func (o *OrderMassActionReport562) MassActionTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.MassActionTypeSinceVersion()
}

func (*OrderMassActionReport562) MassActionTypeDeprecated() uint16 {
	return 0
}

func (*OrderMassActionReport562) MassActionTypeMetaAttribute(meta int) string {
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

func (*OrderMassActionReport562) MassActionTypeMinValue() byte {
	return byte(32)
}

func (*OrderMassActionReport562) MassActionTypeMaxValue() byte {
	return byte(126)
}

func (*OrderMassActionReport562) MassActionTypeNullValue() byte {
	return 0
}

func (*OrderMassActionReport562) SecurityGroupId() uint16 {
	return 1151
}

func (*OrderMassActionReport562) SecurityGroupSinceVersion() uint16 {
	return 0
}

func (o *OrderMassActionReport562) SecurityGroupInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.SecurityGroupSinceVersion()
}

func (*OrderMassActionReport562) SecurityGroupDeprecated() uint16 {
	return 0
}

func (*OrderMassActionReport562) SecurityGroupMetaAttribute(meta int) string {
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

func (*OrderMassActionReport562) SecurityGroupMinValue() byte {
	return byte(32)
}

func (*OrderMassActionReport562) SecurityGroupMaxValue() byte {
	return byte(126)
}

func (*OrderMassActionReport562) SecurityGroupNullValue() byte {
	return 0
}

func (o *OrderMassActionReport562) SecurityGroupCharacterEncoding() string {
	return "US-ASCII"
}

func (*OrderMassActionReport562) LocationId() uint16 {
	return 9537
}

func (*OrderMassActionReport562) LocationSinceVersion() uint16 {
	return 0
}

func (o *OrderMassActionReport562) LocationInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.LocationSinceVersion()
}

func (*OrderMassActionReport562) LocationDeprecated() uint16 {
	return 0
}

func (*OrderMassActionReport562) LocationMetaAttribute(meta int) string {
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

func (*OrderMassActionReport562) LocationMinValue() byte {
	return byte(32)
}

func (*OrderMassActionReport562) LocationMaxValue() byte {
	return byte(126)
}

func (*OrderMassActionReport562) LocationNullValue() byte {
	return 0
}

func (o *OrderMassActionReport562) LocationCharacterEncoding() string {
	return "US-ASCII"
}

func (*OrderMassActionReport562) SecurityIDId() uint16 {
	return 48
}

func (*OrderMassActionReport562) SecurityIDSinceVersion() uint16 {
	return 0
}

func (o *OrderMassActionReport562) SecurityIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.SecurityIDSinceVersion()
}

func (*OrderMassActionReport562) SecurityIDDeprecated() uint16 {
	return 0
}

func (*OrderMassActionReport562) SecurityIDMetaAttribute(meta int) string {
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

func (*OrderMassActionReport562) SecurityIDMinValue() int32 {
	return math.MinInt32 + 1
}

func (*OrderMassActionReport562) SecurityIDMaxValue() int32 {
	return math.MaxInt32
}

func (*OrderMassActionReport562) SecurityIDNullValue() int32 {
	return 2147483647
}

func (*OrderMassActionReport562) DelayDurationId() uint16 {
	return 5904
}

func (*OrderMassActionReport562) DelayDurationSinceVersion() uint16 {
	return 0
}

func (o *OrderMassActionReport562) DelayDurationInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.DelayDurationSinceVersion()
}

func (*OrderMassActionReport562) DelayDurationDeprecated() uint16 {
	return 0
}

func (*OrderMassActionReport562) DelayDurationMetaAttribute(meta int) string {
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

func (*OrderMassActionReport562) DelayDurationMinValue() uint16 {
	return 0
}

func (*OrderMassActionReport562) DelayDurationMaxValue() uint16 {
	return math.MaxUint16 - 1
}

func (*OrderMassActionReport562) DelayDurationNullValue() uint16 {
	return 65535
}

func (*OrderMassActionReport562) MassActionResponseId() uint16 {
	return 1375
}

func (*OrderMassActionReport562) MassActionResponseSinceVersion() uint16 {
	return 0
}

func (o *OrderMassActionReport562) MassActionResponseInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.MassActionResponseSinceVersion()
}

func (*OrderMassActionReport562) MassActionResponseDeprecated() uint16 {
	return 0
}

func (*OrderMassActionReport562) MassActionResponseMetaAttribute(meta int) string {
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

func (*OrderMassActionReport562) ManualOrderIndicatorId() uint16 {
	return 1028
}

func (*OrderMassActionReport562) ManualOrderIndicatorSinceVersion() uint16 {
	return 0
}

func (o *OrderMassActionReport562) ManualOrderIndicatorInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.ManualOrderIndicatorSinceVersion()
}

func (*OrderMassActionReport562) ManualOrderIndicatorDeprecated() uint16 {
	return 0
}

func (*OrderMassActionReport562) ManualOrderIndicatorMetaAttribute(meta int) string {
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

func (*OrderMassActionReport562) MassActionScopeId() uint16 {
	return 1374
}

func (*OrderMassActionReport562) MassActionScopeSinceVersion() uint16 {
	return 0
}

func (o *OrderMassActionReport562) MassActionScopeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.MassActionScopeSinceVersion()
}

func (*OrderMassActionReport562) MassActionScopeDeprecated() uint16 {
	return 0
}

func (*OrderMassActionReport562) MassActionScopeMetaAttribute(meta int) string {
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

func (*OrderMassActionReport562) TotalAffectedOrdersId() uint16 {
	return 533
}

func (*OrderMassActionReport562) TotalAffectedOrdersSinceVersion() uint16 {
	return 0
}

func (o *OrderMassActionReport562) TotalAffectedOrdersInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.TotalAffectedOrdersSinceVersion()
}

func (*OrderMassActionReport562) TotalAffectedOrdersDeprecated() uint16 {
	return 0
}

func (*OrderMassActionReport562) TotalAffectedOrdersMetaAttribute(meta int) string {
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

func (*OrderMassActionReport562) TotalAffectedOrdersMinValue() uint32 {
	return 0
}

func (*OrderMassActionReport562) TotalAffectedOrdersMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*OrderMassActionReport562) TotalAffectedOrdersNullValue() uint32 {
	return math.MaxUint32
}

func (*OrderMassActionReport562) LastFragmentId() uint16 {
	return 893
}

func (*OrderMassActionReport562) LastFragmentSinceVersion() uint16 {
	return 0
}

func (o *OrderMassActionReport562) LastFragmentInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.LastFragmentSinceVersion()
}

func (*OrderMassActionReport562) LastFragmentDeprecated() uint16 {
	return 0
}

func (*OrderMassActionReport562) LastFragmentMetaAttribute(meta int) string {
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

func (*OrderMassActionReport562) MassActionRejectReasonId() uint16 {
	return 1376
}

func (*OrderMassActionReport562) MassActionRejectReasonSinceVersion() uint16 {
	return 0
}

func (o *OrderMassActionReport562) MassActionRejectReasonInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.MassActionRejectReasonSinceVersion()
}

func (*OrderMassActionReport562) MassActionRejectReasonDeprecated() uint16 {
	return 0
}

func (*OrderMassActionReport562) MassActionRejectReasonMetaAttribute(meta int) string {
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

func (*OrderMassActionReport562) MassActionRejectReasonMinValue() uint8 {
	return 0
}

func (*OrderMassActionReport562) MassActionRejectReasonMaxValue() uint8 {
	return math.MaxUint8 - 1
}

func (*OrderMassActionReport562) MassActionRejectReasonNullValue() uint8 {
	return 255
}

func (*OrderMassActionReport562) MarketSegmentIDId() uint16 {
	return 1300
}

func (*OrderMassActionReport562) MarketSegmentIDSinceVersion() uint16 {
	return 0
}

func (o *OrderMassActionReport562) MarketSegmentIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.MarketSegmentIDSinceVersion()
}

func (*OrderMassActionReport562) MarketSegmentIDDeprecated() uint16 {
	return 0
}

func (*OrderMassActionReport562) MarketSegmentIDMetaAttribute(meta int) string {
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

func (*OrderMassActionReport562) MarketSegmentIDMinValue() uint8 {
	return 0
}

func (*OrderMassActionReport562) MarketSegmentIDMaxValue() uint8 {
	return math.MaxUint8 - 1
}

func (*OrderMassActionReport562) MarketSegmentIDNullValue() uint8 {
	return 255
}

func (*OrderMassActionReport562) MassCancelRequestTypeId() uint16 {
	return 6115
}

func (*OrderMassActionReport562) MassCancelRequestTypeSinceVersion() uint16 {
	return 0
}

func (o *OrderMassActionReport562) MassCancelRequestTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.MassCancelRequestTypeSinceVersion()
}

func (*OrderMassActionReport562) MassCancelRequestTypeDeprecated() uint16 {
	return 0
}

func (*OrderMassActionReport562) MassCancelRequestTypeMetaAttribute(meta int) string {
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

func (*OrderMassActionReport562) SideId() uint16 {
	return 54
}

func (*OrderMassActionReport562) SideSinceVersion() uint16 {
	return 0
}

func (o *OrderMassActionReport562) SideInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.SideSinceVersion()
}

func (*OrderMassActionReport562) SideDeprecated() uint16 {
	return 0
}

func (*OrderMassActionReport562) SideMetaAttribute(meta int) string {
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

func (*OrderMassActionReport562) OrdTypeId() uint16 {
	return 40
}

func (*OrderMassActionReport562) OrdTypeSinceVersion() uint16 {
	return 0
}

func (o *OrderMassActionReport562) OrdTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.OrdTypeSinceVersion()
}

func (*OrderMassActionReport562) OrdTypeDeprecated() uint16 {
	return 0
}

func (*OrderMassActionReport562) OrdTypeMetaAttribute(meta int) string {
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

func (*OrderMassActionReport562) TimeInForceId() uint16 {
	return 59
}

func (*OrderMassActionReport562) TimeInForceSinceVersion() uint16 {
	return 0
}

func (o *OrderMassActionReport562) TimeInForceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.TimeInForceSinceVersion()
}

func (*OrderMassActionReport562) TimeInForceDeprecated() uint16 {
	return 0
}

func (*OrderMassActionReport562) TimeInForceMetaAttribute(meta int) string {
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

func (*OrderMassActionReport562) SplitMsgId() uint16 {
	return 9553
}

func (*OrderMassActionReport562) SplitMsgSinceVersion() uint16 {
	return 0
}

func (o *OrderMassActionReport562) SplitMsgInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.SplitMsgSinceVersion()
}

func (*OrderMassActionReport562) SplitMsgDeprecated() uint16 {
	return 0
}

func (*OrderMassActionReport562) SplitMsgMetaAttribute(meta int) string {
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

func (*OrderMassActionReport562) LiquidityFlagId() uint16 {
	return 9373
}

func (*OrderMassActionReport562) LiquidityFlagSinceVersion() uint16 {
	return 0
}

func (o *OrderMassActionReport562) LiquidityFlagInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.LiquidityFlagSinceVersion()
}

func (*OrderMassActionReport562) LiquidityFlagDeprecated() uint16 {
	return 0
}

func (*OrderMassActionReport562) LiquidityFlagMetaAttribute(meta int) string {
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

func (*OrderMassActionReport562) PossRetransFlagId() uint16 {
	return 9765
}

func (*OrderMassActionReport562) PossRetransFlagSinceVersion() uint16 {
	return 0
}

func (o *OrderMassActionReport562) PossRetransFlagInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.PossRetransFlagSinceVersion()
}

func (*OrderMassActionReport562) PossRetransFlagDeprecated() uint16 {
	return 0
}

func (*OrderMassActionReport562) PossRetransFlagMetaAttribute(meta int) string {
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

func (*OrderMassActionReport562) DelayToTimeId() uint16 {
	return 7552
}

func (*OrderMassActionReport562) DelayToTimeSinceVersion() uint16 {
	return 0
}

func (o *OrderMassActionReport562) DelayToTimeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.DelayToTimeSinceVersion()
}

func (*OrderMassActionReport562) DelayToTimeDeprecated() uint16 {
	return 0
}

func (*OrderMassActionReport562) DelayToTimeMetaAttribute(meta int) string {
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

func (*OrderMassActionReport562) DelayToTimeMinValue() uint64 {
	return 0
}

func (*OrderMassActionReport562) DelayToTimeMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*OrderMassActionReport562) DelayToTimeNullValue() uint64 {
	return 18446744073709551615
}

func (*OrderMassActionReport562NoAffectedOrders) OrigCIOrdIDId() uint16 {
	return 41
}

func (*OrderMassActionReport562NoAffectedOrders) OrigCIOrdIDSinceVersion() uint16 {
	return 0
}

func (o *OrderMassActionReport562NoAffectedOrders) OrigCIOrdIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.OrigCIOrdIDSinceVersion()
}

func (*OrderMassActionReport562NoAffectedOrders) OrigCIOrdIDDeprecated() uint16 {
	return 0
}

func (*OrderMassActionReport562NoAffectedOrders) OrigCIOrdIDMetaAttribute(meta int) string {
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

func (*OrderMassActionReport562NoAffectedOrders) OrigCIOrdIDMinValue() byte {
	return byte(32)
}

func (*OrderMassActionReport562NoAffectedOrders) OrigCIOrdIDMaxValue() byte {
	return byte(126)
}

func (*OrderMassActionReport562NoAffectedOrders) OrigCIOrdIDNullValue() byte {
	return 0
}

func (o *OrderMassActionReport562NoAffectedOrders) OrigCIOrdIDCharacterEncoding() string {
	return "US-ASCII"
}

func (*OrderMassActionReport562NoAffectedOrders) AffectedOrderIDId() uint16 {
	return 535
}

func (*OrderMassActionReport562NoAffectedOrders) AffectedOrderIDSinceVersion() uint16 {
	return 0
}

func (o *OrderMassActionReport562NoAffectedOrders) AffectedOrderIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.AffectedOrderIDSinceVersion()
}

func (*OrderMassActionReport562NoAffectedOrders) AffectedOrderIDDeprecated() uint16 {
	return 0
}

func (*OrderMassActionReport562NoAffectedOrders) AffectedOrderIDMetaAttribute(meta int) string {
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

func (*OrderMassActionReport562NoAffectedOrders) AffectedOrderIDMinValue() uint64 {
	return 0
}

func (*OrderMassActionReport562NoAffectedOrders) AffectedOrderIDMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*OrderMassActionReport562NoAffectedOrders) AffectedOrderIDNullValue() uint64 {
	return math.MaxUint64
}

func (*OrderMassActionReport562NoAffectedOrders) CxlQuantityId() uint16 {
	return 84
}

func (*OrderMassActionReport562NoAffectedOrders) CxlQuantitySinceVersion() uint16 {
	return 0
}

func (o *OrderMassActionReport562NoAffectedOrders) CxlQuantityInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.CxlQuantitySinceVersion()
}

func (*OrderMassActionReport562NoAffectedOrders) CxlQuantityDeprecated() uint16 {
	return 0
}

func (*OrderMassActionReport562NoAffectedOrders) CxlQuantityMetaAttribute(meta int) string {
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

func (*OrderMassActionReport562NoAffectedOrders) CxlQuantityMinValue() uint32 {
	return 0
}

func (*OrderMassActionReport562NoAffectedOrders) CxlQuantityMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*OrderMassActionReport562NoAffectedOrders) CxlQuantityNullValue() uint32 {
	return math.MaxUint32
}

func (*OrderMassActionReport562) NoAffectedOrdersId() uint16 {
	return 534
}

func (*OrderMassActionReport562) NoAffectedOrdersSinceVersion() uint16 {
	return 0
}

func (o *OrderMassActionReport562) NoAffectedOrdersInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.NoAffectedOrdersSinceVersion()
}

func (*OrderMassActionReport562) NoAffectedOrdersDeprecated() uint16 {
	return 0
}

func (*OrderMassActionReport562NoAffectedOrders) SbeBlockLength() (blockLength uint) {
	return 32
}

func (*OrderMassActionReport562NoAffectedOrders) SbeSchemaVersion() (schemaVersion uint16) {
	return 5
}
