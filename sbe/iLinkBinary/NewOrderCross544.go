// Generated SBE (Simple Binary Encoding) message codec

package ilinkbinary

import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
)

type NewOrderCross544 struct {
	CrossID              uint64
	OrderRequestID       uint64
	ManualOrderIndicator ManualOrdIndReqEnum
	SeqNum               uint32
	SenderID             [20]byte
	OrdType              [1]byte
	CrossType            [1]byte
	CrossPrioritization  [1]byte
	Price                PRICE9
	TransBkdTime         uint64
	SendingTimeEpoch     uint64
	Location             [5]byte
	SecurityID           int32
	NoSides              []NewOrderCross544NoSides
}
type NewOrderCross544NoSides struct {
	ClOrdID               [20]byte
	PartyDetailsListReqID uint64
	OrderQty              uint32
	Side                  SideReqEnum
	SideTimeInForce       SideTimeInForceEnum
}

func (n *NewOrderCross544) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := n.RangeCheck(n.SbeSchemaVersion(), n.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	if err := _m.WriteUint64(_w, n.CrossID); err != nil {
		return err
	}
	if err := _m.WriteUint64(_w, n.OrderRequestID); err != nil {
		return err
	}
	if err := n.ManualOrderIndicator.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteUint32(_w, n.SeqNum); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, n.SenderID[:]); err != nil {
		return err
	}
	if err := n.Price.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteUint64(_w, n.TransBkdTime); err != nil {
		return err
	}
	if err := _m.WriteUint64(_w, n.SendingTimeEpoch); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, n.Location[:]); err != nil {
		return err
	}
	if err := _m.WriteInt32(_w, n.SecurityID); err != nil {
		return err
	}
	var NoSidesBlockLength uint16 = 34
	if err := _m.WriteUint16(_w, NoSidesBlockLength); err != nil {
		return err
	}
	var NoSidesNumInGroup uint8 = uint8(len(n.NoSides))
	if err := _m.WriteUint8(_w, NoSidesNumInGroup); err != nil {
		return err
	}
	for _, prop := range n.NoSides {
		if err := prop.Encode(_m, _w); err != nil {
			return err
		}
	}
	return nil
}

func (n *NewOrderCross544) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	if !n.CrossIDInActingVersion(actingVersion) {
		n.CrossID = n.CrossIDNullValue()
	} else {
		if err := _m.ReadUint64(_r, &n.CrossID); err != nil {
			return err
		}
	}
	if !n.OrderRequestIDInActingVersion(actingVersion) {
		n.OrderRequestID = n.OrderRequestIDNullValue()
	} else {
		if err := _m.ReadUint64(_r, &n.OrderRequestID); err != nil {
			return err
		}
	}
	if n.ManualOrderIndicatorInActingVersion(actingVersion) {
		if err := n.ManualOrderIndicator.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if !n.SeqNumInActingVersion(actingVersion) {
		n.SeqNum = n.SeqNumNullValue()
	} else {
		if err := _m.ReadUint32(_r, &n.SeqNum); err != nil {
			return err
		}
	}
	if !n.SenderIDInActingVersion(actingVersion) {
		for idx := 0; idx < 20; idx++ {
			n.SenderID[idx] = n.SenderIDNullValue()
		}
	} else {
		if err := _m.ReadBytes(_r, n.SenderID[:]); err != nil {
			return err
		}
	}
	n.OrdType[0] = 50
	n.CrossType[0] = 51
	n.CrossPrioritization[0] = 48
	if n.PriceInActingVersion(actingVersion) {
		if err := n.Price.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if !n.TransBkdTimeInActingVersion(actingVersion) {
		n.TransBkdTime = n.TransBkdTimeNullValue()
	} else {
		if err := _m.ReadUint64(_r, &n.TransBkdTime); err != nil {
			return err
		}
	}
	if !n.SendingTimeEpochInActingVersion(actingVersion) {
		n.SendingTimeEpoch = n.SendingTimeEpochNullValue()
	} else {
		if err := _m.ReadUint64(_r, &n.SendingTimeEpoch); err != nil {
			return err
		}
	}
	if !n.LocationInActingVersion(actingVersion) {
		for idx := 0; idx < 5; idx++ {
			n.Location[idx] = n.LocationNullValue()
		}
	} else {
		if err := _m.ReadBytes(_r, n.Location[:]); err != nil {
			return err
		}
	}
	if !n.SecurityIDInActingVersion(actingVersion) {
		n.SecurityID = n.SecurityIDNullValue()
	} else {
		if err := _m.ReadInt32(_r, &n.SecurityID); err != nil {
			return err
		}
	}
	if actingVersion > n.SbeSchemaVersion() && blockLength > n.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-n.SbeBlockLength()))
	}

	if n.NoSidesInActingVersion(actingVersion) {
		var NoSidesBlockLength uint16
		if err := _m.ReadUint16(_r, &NoSidesBlockLength); err != nil {
			return err
		}
		var NoSidesNumInGroup uint8
		if err := _m.ReadUint8(_r, &NoSidesNumInGroup); err != nil {
			return err
		}
		if cap(n.NoSides) < int(NoSidesNumInGroup) {
			n.NoSides = make([]NewOrderCross544NoSides, NoSidesNumInGroup)
		}
		for i, _ := range n.NoSides {
			if err := n.NoSides[i].Decode(_m, _r, actingVersion, uint(NoSidesBlockLength)); err != nil {
				return err
			}
		}
	}
	if doRangeCheck {
		if err := n.RangeCheck(actingVersion, n.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	return nil
}

func (n *NewOrderCross544) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if n.CrossIDInActingVersion(actingVersion) {
		if n.CrossID < n.CrossIDMinValue() || n.CrossID > n.CrossIDMaxValue() {
			return fmt.Errorf("Range check failed on n.CrossID (%v < %v > %v)", n.CrossIDMinValue(), n.CrossID, n.CrossIDMaxValue())
		}
	}
	if n.OrderRequestIDInActingVersion(actingVersion) {
		if n.OrderRequestID < n.OrderRequestIDMinValue() || n.OrderRequestID > n.OrderRequestIDMaxValue() {
			return fmt.Errorf("Range check failed on n.OrderRequestID (%v < %v > %v)", n.OrderRequestIDMinValue(), n.OrderRequestID, n.OrderRequestIDMaxValue())
		}
	}
	if err := n.ManualOrderIndicator.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if n.SeqNumInActingVersion(actingVersion) {
		if n.SeqNum < n.SeqNumMinValue() || n.SeqNum > n.SeqNumMaxValue() {
			return fmt.Errorf("Range check failed on n.SeqNum (%v < %v > %v)", n.SeqNumMinValue(), n.SeqNum, n.SeqNumMaxValue())
		}
	}
	if n.SenderIDInActingVersion(actingVersion) {
		for idx := 0; idx < 20; idx++ {
			if n.SenderID[idx] < n.SenderIDMinValue() || n.SenderID[idx] > n.SenderIDMaxValue() {
				return fmt.Errorf("Range check failed on n.SenderID[%d] (%v < %v > %v)", idx, n.SenderIDMinValue(), n.SenderID[idx], n.SenderIDMaxValue())
			}
		}
	}
	if n.TransBkdTimeInActingVersion(actingVersion) {
		if n.TransBkdTime < n.TransBkdTimeMinValue() || n.TransBkdTime > n.TransBkdTimeMaxValue() {
			return fmt.Errorf("Range check failed on n.TransBkdTime (%v < %v > %v)", n.TransBkdTimeMinValue(), n.TransBkdTime, n.TransBkdTimeMaxValue())
		}
	}
	if n.SendingTimeEpochInActingVersion(actingVersion) {
		if n.SendingTimeEpoch < n.SendingTimeEpochMinValue() || n.SendingTimeEpoch > n.SendingTimeEpochMaxValue() {
			return fmt.Errorf("Range check failed on n.SendingTimeEpoch (%v < %v > %v)", n.SendingTimeEpochMinValue(), n.SendingTimeEpoch, n.SendingTimeEpochMaxValue())
		}
	}
	if n.LocationInActingVersion(actingVersion) {
		for idx := 0; idx < 5; idx++ {
			if n.Location[idx] < n.LocationMinValue() || n.Location[idx] > n.LocationMaxValue() {
				return fmt.Errorf("Range check failed on n.Location[%d] (%v < %v > %v)", idx, n.LocationMinValue(), n.Location[idx], n.LocationMaxValue())
			}
		}
	}
	if n.SecurityIDInActingVersion(actingVersion) {
		if n.SecurityID < n.SecurityIDMinValue() || n.SecurityID > n.SecurityIDMaxValue() {
			return fmt.Errorf("Range check failed on n.SecurityID (%v < %v > %v)", n.SecurityIDMinValue(), n.SecurityID, n.SecurityIDMaxValue())
		}
	}
	for _, prop := range n.NoSides {
		if err := prop.RangeCheck(actingVersion, schemaVersion); err != nil {
			return err
		}
	}
	return nil
}

func NewOrderCross544Init(n *NewOrderCross544) {
	n.OrdType[0] = 50
	n.CrossType[0] = 51
	n.CrossPrioritization[0] = 48
	return
}

func (n *NewOrderCross544NoSides) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteBytes(_w, n.ClOrdID[:]); err != nil {
		return err
	}
	if err := _m.WriteUint64(_w, n.PartyDetailsListReqID); err != nil {
		return err
	}
	if err := _m.WriteUint32(_w, n.OrderQty); err != nil {
		return err
	}
	if err := n.Side.Encode(_m, _w); err != nil {
		return err
	}
	if err := n.SideTimeInForce.Encode(_m, _w); err != nil {
		return err
	}
	return nil
}

func (n *NewOrderCross544NoSides) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint) error {
	if !n.ClOrdIDInActingVersion(actingVersion) {
		for idx := 0; idx < 20; idx++ {
			n.ClOrdID[idx] = n.ClOrdIDNullValue()
		}
	} else {
		if err := _m.ReadBytes(_r, n.ClOrdID[:]); err != nil {
			return err
		}
	}
	if !n.PartyDetailsListReqIDInActingVersion(actingVersion) {
		n.PartyDetailsListReqID = n.PartyDetailsListReqIDNullValue()
	} else {
		if err := _m.ReadUint64(_r, &n.PartyDetailsListReqID); err != nil {
			return err
		}
	}
	if !n.OrderQtyInActingVersion(actingVersion) {
		n.OrderQty = n.OrderQtyNullValue()
	} else {
		if err := _m.ReadUint32(_r, &n.OrderQty); err != nil {
			return err
		}
	}
	if n.SideInActingVersion(actingVersion) {
		if err := n.Side.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if n.SideTimeInForceInActingVersion(actingVersion) {
		if err := n.SideTimeInForce.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if actingVersion > n.SbeSchemaVersion() && blockLength > n.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-n.SbeBlockLength()))
	}
	return nil
}

func (n *NewOrderCross544NoSides) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if n.ClOrdIDInActingVersion(actingVersion) {
		for idx := 0; idx < 20; idx++ {
			if n.ClOrdID[idx] < n.ClOrdIDMinValue() || n.ClOrdID[idx] > n.ClOrdIDMaxValue() {
				return fmt.Errorf("Range check failed on n.ClOrdID[%d] (%v < %v > %v)", idx, n.ClOrdIDMinValue(), n.ClOrdID[idx], n.ClOrdIDMaxValue())
			}
		}
	}
	if n.PartyDetailsListReqIDInActingVersion(actingVersion) {
		if n.PartyDetailsListReqID < n.PartyDetailsListReqIDMinValue() || n.PartyDetailsListReqID > n.PartyDetailsListReqIDMaxValue() {
			return fmt.Errorf("Range check failed on n.PartyDetailsListReqID (%v < %v > %v)", n.PartyDetailsListReqIDMinValue(), n.PartyDetailsListReqID, n.PartyDetailsListReqIDMaxValue())
		}
	}
	if n.OrderQtyInActingVersion(actingVersion) {
		if n.OrderQty < n.OrderQtyMinValue() || n.OrderQty > n.OrderQtyMaxValue() {
			return fmt.Errorf("Range check failed on n.OrderQty (%v < %v > %v)", n.OrderQtyMinValue(), n.OrderQty, n.OrderQtyMaxValue())
		}
	}
	if err := n.Side.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if err := n.SideTimeInForce.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	return nil
}

func NewOrderCross544NoSidesInit(n *NewOrderCross544NoSides) {
	return
}

func (*NewOrderCross544) SbeBlockLength() (blockLength uint16) {
	return 74
}

func (*NewOrderCross544) SbeTemplateId() (templateId uint16) {
	return 544
}

func (*NewOrderCross544) SbeSchemaId() (schemaId uint16) {
	return 8
}

func (*NewOrderCross544) SbeSchemaVersion() (schemaVersion uint16) {
	return 5
}

func (*NewOrderCross544) SbeSemanticType() (semanticType []byte) {
	return []byte("c")
}

func (*NewOrderCross544) CrossIDId() uint16 {
	return 548
}

func (*NewOrderCross544) CrossIDSinceVersion() uint16 {
	return 0
}

func (n *NewOrderCross544) CrossIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.CrossIDSinceVersion()
}

func (*NewOrderCross544) CrossIDDeprecated() uint16 {
	return 0
}

func (*NewOrderCross544) CrossIDMetaAttribute(meta int) string {
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

func (*NewOrderCross544) CrossIDMinValue() uint64 {
	return 0
}

func (*NewOrderCross544) CrossIDMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*NewOrderCross544) CrossIDNullValue() uint64 {
	return math.MaxUint64
}

func (*NewOrderCross544) OrderRequestIDId() uint16 {
	return 2422
}

func (*NewOrderCross544) OrderRequestIDSinceVersion() uint16 {
	return 0
}

func (n *NewOrderCross544) OrderRequestIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.OrderRequestIDSinceVersion()
}

func (*NewOrderCross544) OrderRequestIDDeprecated() uint16 {
	return 0
}

func (*NewOrderCross544) OrderRequestIDMetaAttribute(meta int) string {
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

func (*NewOrderCross544) OrderRequestIDMinValue() uint64 {
	return 0
}

func (*NewOrderCross544) OrderRequestIDMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*NewOrderCross544) OrderRequestIDNullValue() uint64 {
	return math.MaxUint64
}

func (*NewOrderCross544) ManualOrderIndicatorId() uint16 {
	return 1028
}

func (*NewOrderCross544) ManualOrderIndicatorSinceVersion() uint16 {
	return 0
}

func (n *NewOrderCross544) ManualOrderIndicatorInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.ManualOrderIndicatorSinceVersion()
}

func (*NewOrderCross544) ManualOrderIndicatorDeprecated() uint16 {
	return 0
}

func (*NewOrderCross544) ManualOrderIndicatorMetaAttribute(meta int) string {
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

func (*NewOrderCross544) SeqNumId() uint16 {
	return 9726
}

func (*NewOrderCross544) SeqNumSinceVersion() uint16 {
	return 0
}

func (n *NewOrderCross544) SeqNumInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.SeqNumSinceVersion()
}

func (*NewOrderCross544) SeqNumDeprecated() uint16 {
	return 0
}

func (*NewOrderCross544) SeqNumMetaAttribute(meta int) string {
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

func (*NewOrderCross544) SeqNumMinValue() uint32 {
	return 0
}

func (*NewOrderCross544) SeqNumMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*NewOrderCross544) SeqNumNullValue() uint32 {
	return math.MaxUint32
}

func (*NewOrderCross544) SenderIDId() uint16 {
	return 5392
}

func (*NewOrderCross544) SenderIDSinceVersion() uint16 {
	return 0
}

func (n *NewOrderCross544) SenderIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.SenderIDSinceVersion()
}

func (*NewOrderCross544) SenderIDDeprecated() uint16 {
	return 0
}

func (*NewOrderCross544) SenderIDMetaAttribute(meta int) string {
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

func (*NewOrderCross544) SenderIDMinValue() byte {
	return byte(32)
}

func (*NewOrderCross544) SenderIDMaxValue() byte {
	return byte(126)
}

func (*NewOrderCross544) SenderIDNullValue() byte {
	return 0
}

func (n *NewOrderCross544) SenderIDCharacterEncoding() string {
	return "US-ASCII"
}

func (*NewOrderCross544) OrdTypeId() uint16 {
	return 40
}

func (*NewOrderCross544) OrdTypeSinceVersion() uint16 {
	return 0
}

func (n *NewOrderCross544) OrdTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.OrdTypeSinceVersion()
}

func (*NewOrderCross544) OrdTypeDeprecated() uint16 {
	return 0
}

func (*NewOrderCross544) OrdTypeMetaAttribute(meta int) string {
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

func (*NewOrderCross544) OrdTypeMinValue() byte {
	return byte(32)
}

func (*NewOrderCross544) OrdTypeMaxValue() byte {
	return byte(126)
}

func (*NewOrderCross544) OrdTypeNullValue() byte {
	return 0
}

func (*NewOrderCross544) CrossTypeId() uint16 {
	return 549
}

func (*NewOrderCross544) CrossTypeSinceVersion() uint16 {
	return 0
}

func (n *NewOrderCross544) CrossTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.CrossTypeSinceVersion()
}

func (*NewOrderCross544) CrossTypeDeprecated() uint16 {
	return 0
}

func (*NewOrderCross544) CrossTypeMetaAttribute(meta int) string {
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

func (*NewOrderCross544) CrossTypeMinValue() byte {
	return byte(32)
}

func (*NewOrderCross544) CrossTypeMaxValue() byte {
	return byte(126)
}

func (*NewOrderCross544) CrossTypeNullValue() byte {
	return 0
}

func (*NewOrderCross544) CrossPrioritizationId() uint16 {
	return 550
}

func (*NewOrderCross544) CrossPrioritizationSinceVersion() uint16 {
	return 0
}

func (n *NewOrderCross544) CrossPrioritizationInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.CrossPrioritizationSinceVersion()
}

func (*NewOrderCross544) CrossPrioritizationDeprecated() uint16 {
	return 0
}

func (*NewOrderCross544) CrossPrioritizationMetaAttribute(meta int) string {
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

func (*NewOrderCross544) CrossPrioritizationMinValue() byte {
	return byte(32)
}

func (*NewOrderCross544) CrossPrioritizationMaxValue() byte {
	return byte(126)
}

func (*NewOrderCross544) CrossPrioritizationNullValue() byte {
	return 0
}

func (*NewOrderCross544) PriceId() uint16 {
	return 44
}

func (*NewOrderCross544) PriceSinceVersion() uint16 {
	return 0
}

func (n *NewOrderCross544) PriceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.PriceSinceVersion()
}

func (*NewOrderCross544) PriceDeprecated() uint16 {
	return 0
}

func (*NewOrderCross544) PriceMetaAttribute(meta int) string {
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

func (*NewOrderCross544) TransBkdTimeId() uint16 {
	return 483
}

func (*NewOrderCross544) TransBkdTimeSinceVersion() uint16 {
	return 0
}

func (n *NewOrderCross544) TransBkdTimeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.TransBkdTimeSinceVersion()
}

func (*NewOrderCross544) TransBkdTimeDeprecated() uint16 {
	return 0
}

func (*NewOrderCross544) TransBkdTimeMetaAttribute(meta int) string {
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

func (*NewOrderCross544) TransBkdTimeMinValue() uint64 {
	return 0
}

func (*NewOrderCross544) TransBkdTimeMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*NewOrderCross544) TransBkdTimeNullValue() uint64 {
	return math.MaxUint64
}

func (*NewOrderCross544) SendingTimeEpochId() uint16 {
	return 5297
}

func (*NewOrderCross544) SendingTimeEpochSinceVersion() uint16 {
	return 0
}

func (n *NewOrderCross544) SendingTimeEpochInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.SendingTimeEpochSinceVersion()
}

func (*NewOrderCross544) SendingTimeEpochDeprecated() uint16 {
	return 0
}

func (*NewOrderCross544) SendingTimeEpochMetaAttribute(meta int) string {
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

func (*NewOrderCross544) SendingTimeEpochMinValue() uint64 {
	return 0
}

func (*NewOrderCross544) SendingTimeEpochMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*NewOrderCross544) SendingTimeEpochNullValue() uint64 {
	return math.MaxUint64
}

func (*NewOrderCross544) LocationId() uint16 {
	return 9537
}

func (*NewOrderCross544) LocationSinceVersion() uint16 {
	return 0
}

func (n *NewOrderCross544) LocationInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.LocationSinceVersion()
}

func (*NewOrderCross544) LocationDeprecated() uint16 {
	return 0
}

func (*NewOrderCross544) LocationMetaAttribute(meta int) string {
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

func (*NewOrderCross544) LocationMinValue() byte {
	return byte(32)
}

func (*NewOrderCross544) LocationMaxValue() byte {
	return byte(126)
}

func (*NewOrderCross544) LocationNullValue() byte {
	return 0
}

func (n *NewOrderCross544) LocationCharacterEncoding() string {
	return "US-ASCII"
}

func (*NewOrderCross544) SecurityIDId() uint16 {
	return 48
}

func (*NewOrderCross544) SecurityIDSinceVersion() uint16 {
	return 0
}

func (n *NewOrderCross544) SecurityIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.SecurityIDSinceVersion()
}

func (*NewOrderCross544) SecurityIDDeprecated() uint16 {
	return 0
}

func (*NewOrderCross544) SecurityIDMetaAttribute(meta int) string {
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

func (*NewOrderCross544) SecurityIDMinValue() int32 {
	return math.MinInt32 + 1
}

func (*NewOrderCross544) SecurityIDMaxValue() int32 {
	return math.MaxInt32
}

func (*NewOrderCross544) SecurityIDNullValue() int32 {
	return math.MinInt32
}

func (*NewOrderCross544NoSides) ClOrdIDId() uint16 {
	return 11
}

func (*NewOrderCross544NoSides) ClOrdIDSinceVersion() uint16 {
	return 0
}

func (n *NewOrderCross544NoSides) ClOrdIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.ClOrdIDSinceVersion()
}

func (*NewOrderCross544NoSides) ClOrdIDDeprecated() uint16 {
	return 0
}

func (*NewOrderCross544NoSides) ClOrdIDMetaAttribute(meta int) string {
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

func (*NewOrderCross544NoSides) ClOrdIDMinValue() byte {
	return byte(32)
}

func (*NewOrderCross544NoSides) ClOrdIDMaxValue() byte {
	return byte(126)
}

func (*NewOrderCross544NoSides) ClOrdIDNullValue() byte {
	return 0
}

func (n *NewOrderCross544NoSides) ClOrdIDCharacterEncoding() string {
	return "US-ASCII"
}

func (*NewOrderCross544NoSides) PartyDetailsListReqIDId() uint16 {
	return 1505
}

func (*NewOrderCross544NoSides) PartyDetailsListReqIDSinceVersion() uint16 {
	return 0
}

func (n *NewOrderCross544NoSides) PartyDetailsListReqIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.PartyDetailsListReqIDSinceVersion()
}

func (*NewOrderCross544NoSides) PartyDetailsListReqIDDeprecated() uint16 {
	return 0
}

func (*NewOrderCross544NoSides) PartyDetailsListReqIDMetaAttribute(meta int) string {
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

func (*NewOrderCross544NoSides) PartyDetailsListReqIDMinValue() uint64 {
	return 0
}

func (*NewOrderCross544NoSides) PartyDetailsListReqIDMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*NewOrderCross544NoSides) PartyDetailsListReqIDNullValue() uint64 {
	return math.MaxUint64
}

func (*NewOrderCross544NoSides) OrderQtyId() uint16 {
	return 38
}

func (*NewOrderCross544NoSides) OrderQtySinceVersion() uint16 {
	return 0
}

func (n *NewOrderCross544NoSides) OrderQtyInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.OrderQtySinceVersion()
}

func (*NewOrderCross544NoSides) OrderQtyDeprecated() uint16 {
	return 0
}

func (*NewOrderCross544NoSides) OrderQtyMetaAttribute(meta int) string {
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

func (*NewOrderCross544NoSides) OrderQtyMinValue() uint32 {
	return 0
}

func (*NewOrderCross544NoSides) OrderQtyMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*NewOrderCross544NoSides) OrderQtyNullValue() uint32 {
	return math.MaxUint32
}

func (*NewOrderCross544NoSides) SideId() uint16 {
	return 54
}

func (*NewOrderCross544NoSides) SideSinceVersion() uint16 {
	return 0
}

func (n *NewOrderCross544NoSides) SideInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.SideSinceVersion()
}

func (*NewOrderCross544NoSides) SideDeprecated() uint16 {
	return 0
}

func (*NewOrderCross544NoSides) SideMetaAttribute(meta int) string {
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

func (*NewOrderCross544NoSides) SideTimeInForceId() uint16 {
	return 962
}

func (*NewOrderCross544NoSides) SideTimeInForceSinceVersion() uint16 {
	return 0
}

func (n *NewOrderCross544NoSides) SideTimeInForceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.SideTimeInForceSinceVersion()
}

func (*NewOrderCross544NoSides) SideTimeInForceDeprecated() uint16 {
	return 0
}

func (*NewOrderCross544NoSides) SideTimeInForceMetaAttribute(meta int) string {
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

func (*NewOrderCross544) NoSidesId() uint16 {
	return 552
}

func (*NewOrderCross544) NoSidesSinceVersion() uint16 {
	return 0
}

func (n *NewOrderCross544) NoSidesInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.NoSidesSinceVersion()
}

func (*NewOrderCross544) NoSidesDeprecated() uint16 {
	return 0
}

func (*NewOrderCross544NoSides) SbeBlockLength() (blockLength uint) {
	return 34
}

func (*NewOrderCross544NoSides) SbeSchemaVersion() (schemaVersion uint16) {
	return 5
}
