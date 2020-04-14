// Generated SBE (Simple Binary Encoding) message codec

package ilinkbinary

import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
)

type OrderStatusRequest533 struct {
	PartyDetailsListReqID uint64
	OrdStatusReqID        uint64
	ManualOrderIndicator  ManualOrdIndReqEnum
	SeqNum                uint32
	SenderID              [20]byte
	OrderID               uint64
	SendingTimeEpoch      uint64
	Location              [5]byte
}

func (o *OrderStatusRequest533) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := o.RangeCheck(o.SbeSchemaVersion(), o.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	if err := _m.WriteUint64(_w, o.PartyDetailsListReqID); err != nil {
		return err
	}
	if err := _m.WriteUint64(_w, o.OrdStatusReqID); err != nil {
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
	if err := _m.WriteUint64(_w, o.OrderID); err != nil {
		return err
	}
	if err := _m.WriteUint64(_w, o.SendingTimeEpoch); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, o.Location[:]); err != nil {
		return err
	}
	return nil
}

func (o *OrderStatusRequest533) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	if !o.PartyDetailsListReqIDInActingVersion(actingVersion) {
		o.PartyDetailsListReqID = o.PartyDetailsListReqIDNullValue()
	} else {
		if err := _m.ReadUint64(_r, &o.PartyDetailsListReqID); err != nil {
			return err
		}
	}
	if !o.OrdStatusReqIDInActingVersion(actingVersion) {
		o.OrdStatusReqID = o.OrdStatusReqIDNullValue()
	} else {
		if err := _m.ReadUint64(_r, &o.OrdStatusReqID); err != nil {
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
	if !o.OrderIDInActingVersion(actingVersion) {
		o.OrderID = o.OrderIDNullValue()
	} else {
		if err := _m.ReadUint64(_r, &o.OrderID); err != nil {
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
	if !o.LocationInActingVersion(actingVersion) {
		for idx := 0; idx < 5; idx++ {
			o.Location[idx] = o.LocationNullValue()
		}
	} else {
		if err := _m.ReadBytes(_r, o.Location[:]); err != nil {
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

func (o *OrderStatusRequest533) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if o.PartyDetailsListReqIDInActingVersion(actingVersion) {
		if o.PartyDetailsListReqID < o.PartyDetailsListReqIDMinValue() || o.PartyDetailsListReqID > o.PartyDetailsListReqIDMaxValue() {
			return fmt.Errorf("Range check failed on o.PartyDetailsListReqID (%v < %v > %v)", o.PartyDetailsListReqIDMinValue(), o.PartyDetailsListReqID, o.PartyDetailsListReqIDMaxValue())
		}
	}
	if o.OrdStatusReqIDInActingVersion(actingVersion) {
		if o.OrdStatusReqID < o.OrdStatusReqIDMinValue() || o.OrdStatusReqID > o.OrdStatusReqIDMaxValue() {
			return fmt.Errorf("Range check failed on o.OrdStatusReqID (%v < %v > %v)", o.OrdStatusReqIDMinValue(), o.OrdStatusReqID, o.OrdStatusReqIDMaxValue())
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
	if o.OrderIDInActingVersion(actingVersion) {
		if o.OrderID < o.OrderIDMinValue() || o.OrderID > o.OrderIDMaxValue() {
			return fmt.Errorf("Range check failed on o.OrderID (%v < %v > %v)", o.OrderIDMinValue(), o.OrderID, o.OrderIDMaxValue())
		}
	}
	if o.SendingTimeEpochInActingVersion(actingVersion) {
		if o.SendingTimeEpoch < o.SendingTimeEpochMinValue() || o.SendingTimeEpoch > o.SendingTimeEpochMaxValue() {
			return fmt.Errorf("Range check failed on o.SendingTimeEpoch (%v < %v > %v)", o.SendingTimeEpochMinValue(), o.SendingTimeEpoch, o.SendingTimeEpochMaxValue())
		}
	}
	if o.LocationInActingVersion(actingVersion) {
		for idx := 0; idx < 5; idx++ {
			if o.Location[idx] < o.LocationMinValue() || o.Location[idx] > o.LocationMaxValue() {
				return fmt.Errorf("Range check failed on o.Location[%d] (%v < %v > %v)", idx, o.LocationMinValue(), o.Location[idx], o.LocationMaxValue())
			}
		}
	}
	return nil
}

func OrderStatusRequest533Init(o *OrderStatusRequest533) {
	return
}

func (*OrderStatusRequest533) SbeBlockLength() (blockLength uint16) {
	return 62
}

func (*OrderStatusRequest533) SbeTemplateId() (templateId uint16) {
	return 533
}

func (*OrderStatusRequest533) SbeSchemaId() (schemaId uint16) {
	return 8
}

func (*OrderStatusRequest533) SbeSchemaVersion() (schemaVersion uint16) {
	return 5
}

func (*OrderStatusRequest533) SbeSemanticType() (semanticType []byte) {
	return []byte("H")
}

func (*OrderStatusRequest533) PartyDetailsListReqIDId() uint16 {
	return 1505
}

func (*OrderStatusRequest533) PartyDetailsListReqIDSinceVersion() uint16 {
	return 0
}

func (o *OrderStatusRequest533) PartyDetailsListReqIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.PartyDetailsListReqIDSinceVersion()
}

func (*OrderStatusRequest533) PartyDetailsListReqIDDeprecated() uint16 {
	return 0
}

func (*OrderStatusRequest533) PartyDetailsListReqIDMetaAttribute(meta int) string {
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

func (*OrderStatusRequest533) PartyDetailsListReqIDMinValue() uint64 {
	return 0
}

func (*OrderStatusRequest533) PartyDetailsListReqIDMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*OrderStatusRequest533) PartyDetailsListReqIDNullValue() uint64 {
	return math.MaxUint64
}

func (*OrderStatusRequest533) OrdStatusReqIDId() uint16 {
	return 790
}

func (*OrderStatusRequest533) OrdStatusReqIDSinceVersion() uint16 {
	return 0
}

func (o *OrderStatusRequest533) OrdStatusReqIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.OrdStatusReqIDSinceVersion()
}

func (*OrderStatusRequest533) OrdStatusReqIDDeprecated() uint16 {
	return 0
}

func (*OrderStatusRequest533) OrdStatusReqIDMetaAttribute(meta int) string {
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

func (*OrderStatusRequest533) OrdStatusReqIDMinValue() uint64 {
	return 0
}

func (*OrderStatusRequest533) OrdStatusReqIDMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*OrderStatusRequest533) OrdStatusReqIDNullValue() uint64 {
	return math.MaxUint64
}

func (*OrderStatusRequest533) ManualOrderIndicatorId() uint16 {
	return 1028
}

func (*OrderStatusRequest533) ManualOrderIndicatorSinceVersion() uint16 {
	return 0
}

func (o *OrderStatusRequest533) ManualOrderIndicatorInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.ManualOrderIndicatorSinceVersion()
}

func (*OrderStatusRequest533) ManualOrderIndicatorDeprecated() uint16 {
	return 0
}

func (*OrderStatusRequest533) ManualOrderIndicatorMetaAttribute(meta int) string {
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

func (*OrderStatusRequest533) SeqNumId() uint16 {
	return 9726
}

func (*OrderStatusRequest533) SeqNumSinceVersion() uint16 {
	return 0
}

func (o *OrderStatusRequest533) SeqNumInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.SeqNumSinceVersion()
}

func (*OrderStatusRequest533) SeqNumDeprecated() uint16 {
	return 0
}

func (*OrderStatusRequest533) SeqNumMetaAttribute(meta int) string {
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

func (*OrderStatusRequest533) SeqNumMinValue() uint32 {
	return 0
}

func (*OrderStatusRequest533) SeqNumMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*OrderStatusRequest533) SeqNumNullValue() uint32 {
	return math.MaxUint32
}

func (*OrderStatusRequest533) SenderIDId() uint16 {
	return 5392
}

func (*OrderStatusRequest533) SenderIDSinceVersion() uint16 {
	return 0
}

func (o *OrderStatusRequest533) SenderIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.SenderIDSinceVersion()
}

func (*OrderStatusRequest533) SenderIDDeprecated() uint16 {
	return 0
}

func (*OrderStatusRequest533) SenderIDMetaAttribute(meta int) string {
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

func (*OrderStatusRequest533) SenderIDMinValue() byte {
	return byte(32)
}

func (*OrderStatusRequest533) SenderIDMaxValue() byte {
	return byte(126)
}

func (*OrderStatusRequest533) SenderIDNullValue() byte {
	return 0
}

func (o *OrderStatusRequest533) SenderIDCharacterEncoding() string {
	return "US-ASCII"
}

func (*OrderStatusRequest533) OrderIDId() uint16 {
	return 37
}

func (*OrderStatusRequest533) OrderIDSinceVersion() uint16 {
	return 0
}

func (o *OrderStatusRequest533) OrderIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.OrderIDSinceVersion()
}

func (*OrderStatusRequest533) OrderIDDeprecated() uint16 {
	return 0
}

func (*OrderStatusRequest533) OrderIDMetaAttribute(meta int) string {
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

func (*OrderStatusRequest533) OrderIDMinValue() uint64 {
	return 0
}

func (*OrderStatusRequest533) OrderIDMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*OrderStatusRequest533) OrderIDNullValue() uint64 {
	return math.MaxUint64
}

func (*OrderStatusRequest533) SendingTimeEpochId() uint16 {
	return 5297
}

func (*OrderStatusRequest533) SendingTimeEpochSinceVersion() uint16 {
	return 0
}

func (o *OrderStatusRequest533) SendingTimeEpochInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.SendingTimeEpochSinceVersion()
}

func (*OrderStatusRequest533) SendingTimeEpochDeprecated() uint16 {
	return 0
}

func (*OrderStatusRequest533) SendingTimeEpochMetaAttribute(meta int) string {
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

func (*OrderStatusRequest533) SendingTimeEpochMinValue() uint64 {
	return 0
}

func (*OrderStatusRequest533) SendingTimeEpochMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*OrderStatusRequest533) SendingTimeEpochNullValue() uint64 {
	return math.MaxUint64
}

func (*OrderStatusRequest533) LocationId() uint16 {
	return 9537
}

func (*OrderStatusRequest533) LocationSinceVersion() uint16 {
	return 0
}

func (o *OrderStatusRequest533) LocationInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.LocationSinceVersion()
}

func (*OrderStatusRequest533) LocationDeprecated() uint16 {
	return 0
}

func (*OrderStatusRequest533) LocationMetaAttribute(meta int) string {
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

func (*OrderStatusRequest533) LocationMinValue() byte {
	return byte(32)
}

func (*OrderStatusRequest533) LocationMaxValue() byte {
	return byte(126)
}

func (*OrderStatusRequest533) LocationNullValue() byte {
	return 0
}

func (o *OrderStatusRequest533) LocationCharacterEncoding() string {
	return "US-ASCII"
}
