// Generated SBE (Simple Binary Encoding) message codec

package ilinkbinary

import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
)

type NewOrderSingle514 struct {
	Price                 PRICENULL9
	OrderQty              uint32
	SecurityID            int32
	Side                  SideReqEnum
	SeqNum                uint32
	SenderID              [20]byte
	ClOrdID               [20]byte
	PartyDetailsListReqID uint64
	OrderRequestID        uint64
	SendingTimeEpoch      uint64
	StopPx                PRICENULL9
	Location              [5]byte
	MinQty                uint32
	DisplayQty            uint32
	ExpireDate            uint16
	OrdType               OrderTypeReqEnum
	TimeInForce           TimeInForceEnum
	ManualOrderIndicator  ManualOrdIndReqEnum
	ExecInst              ExecInst
	ExecutionMode         ExecModeEnum
	LiquidityFlag         BooleanNULLEnum
	ManagedOrder          BooleanNULLEnum
	ShortSaleType         ShortSaleTypeEnum
}

func (n *NewOrderSingle514) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := n.RangeCheck(n.SbeSchemaVersion(), n.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	if err := n.Price.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteUint32(_w, n.OrderQty); err != nil {
		return err
	}
	if err := _m.WriteInt32(_w, n.SecurityID); err != nil {
		return err
	}
	if err := n.Side.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteUint32(_w, n.SeqNum); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, n.SenderID[:]); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, n.ClOrdID[:]); err != nil {
		return err
	}
	if err := _m.WriteUint64(_w, n.PartyDetailsListReqID); err != nil {
		return err
	}
	if err := _m.WriteUint64(_w, n.OrderRequestID); err != nil {
		return err
	}
	if err := _m.WriteUint64(_w, n.SendingTimeEpoch); err != nil {
		return err
	}
	if err := n.StopPx.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, n.Location[:]); err != nil {
		return err
	}
	if err := _m.WriteUint32(_w, n.MinQty); err != nil {
		return err
	}
	if err := _m.WriteUint32(_w, n.DisplayQty); err != nil {
		return err
	}
	if err := _m.WriteUint16(_w, n.ExpireDate); err != nil {
		return err
	}
	if err := n.OrdType.Encode(_m, _w); err != nil {
		return err
	}
	if err := n.TimeInForce.Encode(_m, _w); err != nil {
		return err
	}
	if err := n.ManualOrderIndicator.Encode(_m, _w); err != nil {
		return err
	}
	if err := n.ExecInst.Encode(_m, _w); err != nil {
		return err
	}
	if err := n.ExecutionMode.Encode(_m, _w); err != nil {
		return err
	}
	if err := n.LiquidityFlag.Encode(_m, _w); err != nil {
		return err
	}
	if err := n.ManagedOrder.Encode(_m, _w); err != nil {
		return err
	}
	if err := n.ShortSaleType.Encode(_m, _w); err != nil {
		return err
	}
	return nil
}

func (n *NewOrderSingle514) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	if n.PriceInActingVersion(actingVersion) {
		if err := n.Price.Decode(_m, _r, actingVersion); err != nil {
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
	if !n.SecurityIDInActingVersion(actingVersion) {
		n.SecurityID = n.SecurityIDNullValue()
	} else {
		if err := _m.ReadInt32(_r, &n.SecurityID); err != nil {
			return err
		}
	}
	if n.SideInActingVersion(actingVersion) {
		if err := n.Side.Decode(_m, _r, actingVersion); err != nil {
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
	if !n.OrderRequestIDInActingVersion(actingVersion) {
		n.OrderRequestID = n.OrderRequestIDNullValue()
	} else {
		if err := _m.ReadUint64(_r, &n.OrderRequestID); err != nil {
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
	if n.StopPxInActingVersion(actingVersion) {
		if err := n.StopPx.Decode(_m, _r, actingVersion); err != nil {
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
	if !n.MinQtyInActingVersion(actingVersion) {
		n.MinQty = n.MinQtyNullValue()
	} else {
		if err := _m.ReadUint32(_r, &n.MinQty); err != nil {
			return err
		}
	}
	if !n.DisplayQtyInActingVersion(actingVersion) {
		n.DisplayQty = n.DisplayQtyNullValue()
	} else {
		if err := _m.ReadUint32(_r, &n.DisplayQty); err != nil {
			return err
		}
	}
	if !n.ExpireDateInActingVersion(actingVersion) {
		n.ExpireDate = n.ExpireDateNullValue()
	} else {
		if err := _m.ReadUint16(_r, &n.ExpireDate); err != nil {
			return err
		}
	}
	if n.OrdTypeInActingVersion(actingVersion) {
		if err := n.OrdType.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if n.TimeInForceInActingVersion(actingVersion) {
		if err := n.TimeInForce.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if n.ManualOrderIndicatorInActingVersion(actingVersion) {
		if err := n.ManualOrderIndicator.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if n.ExecInstInActingVersion(actingVersion) {
		if err := n.ExecInst.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if n.ExecutionModeInActingVersion(actingVersion) {
		if err := n.ExecutionMode.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if n.LiquidityFlagInActingVersion(actingVersion) {
		if err := n.LiquidityFlag.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if n.ManagedOrderInActingVersion(actingVersion) {
		if err := n.ManagedOrder.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if n.ShortSaleTypeInActingVersion(actingVersion) {
		if err := n.ShortSaleType.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if actingVersion > n.SbeSchemaVersion() && blockLength > n.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-n.SbeBlockLength()))
	}
	if doRangeCheck {
		if err := n.RangeCheck(actingVersion, n.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	return nil
}

func (n *NewOrderSingle514) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if n.OrderQtyInActingVersion(actingVersion) {
		if n.OrderQty < n.OrderQtyMinValue() || n.OrderQty > n.OrderQtyMaxValue() {
			return fmt.Errorf("Range check failed on n.OrderQty (%v < %v > %v)", n.OrderQtyMinValue(), n.OrderQty, n.OrderQtyMaxValue())
		}
	}
	if n.SecurityIDInActingVersion(actingVersion) {
		if n.SecurityID < n.SecurityIDMinValue() || n.SecurityID > n.SecurityIDMaxValue() {
			return fmt.Errorf("Range check failed on n.SecurityID (%v < %v > %v)", n.SecurityIDMinValue(), n.SecurityID, n.SecurityIDMaxValue())
		}
	}
	if err := n.Side.RangeCheck(actingVersion, schemaVersion); err != nil {
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
	if n.OrderRequestIDInActingVersion(actingVersion) {
		if n.OrderRequestID < n.OrderRequestIDMinValue() || n.OrderRequestID > n.OrderRequestIDMaxValue() {
			return fmt.Errorf("Range check failed on n.OrderRequestID (%v < %v > %v)", n.OrderRequestIDMinValue(), n.OrderRequestID, n.OrderRequestIDMaxValue())
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
	if n.MinQtyInActingVersion(actingVersion) {
		if n.MinQty != n.MinQtyNullValue() && (n.MinQty < n.MinQtyMinValue() || n.MinQty > n.MinQtyMaxValue()) {
			return fmt.Errorf("Range check failed on n.MinQty (%v < %v > %v)", n.MinQtyMinValue(), n.MinQty, n.MinQtyMaxValue())
		}
	}
	if n.DisplayQtyInActingVersion(actingVersion) {
		if n.DisplayQty != n.DisplayQtyNullValue() && (n.DisplayQty < n.DisplayQtyMinValue() || n.DisplayQty > n.DisplayQtyMaxValue()) {
			return fmt.Errorf("Range check failed on n.DisplayQty (%v < %v > %v)", n.DisplayQtyMinValue(), n.DisplayQty, n.DisplayQtyMaxValue())
		}
	}
	if n.ExpireDateInActingVersion(actingVersion) {
		if n.ExpireDate != n.ExpireDateNullValue() && (n.ExpireDate < n.ExpireDateMinValue() || n.ExpireDate > n.ExpireDateMaxValue()) {
			return fmt.Errorf("Range check failed on n.ExpireDate (%v < %v > %v)", n.ExpireDateMinValue(), n.ExpireDate, n.ExpireDateMaxValue())
		}
	}
	if err := n.OrdType.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if err := n.TimeInForce.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if err := n.ManualOrderIndicator.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if err := n.ExecutionMode.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if err := n.LiquidityFlag.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if err := n.ManagedOrder.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if err := n.ShortSaleType.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	return nil
}

func NewOrderSingle514Init(n *NewOrderSingle514) {
	n.MinQty = 4294967295
	n.DisplayQty = 4294967295
	n.ExpireDate = 65535
	return
}

func (*NewOrderSingle514) SbeBlockLength() (blockLength uint16) {
	return 116
}

func (*NewOrderSingle514) SbeTemplateId() (templateId uint16) {
	return 514
}

func (*NewOrderSingle514) SbeSchemaId() (schemaId uint16) {
	return 8
}

func (*NewOrderSingle514) SbeSchemaVersion() (schemaVersion uint16) {
	return 5
}

func (*NewOrderSingle514) SbeSemanticType() (semanticType []byte) {
	return []byte("D")
}

func (*NewOrderSingle514) PriceId() uint16 {
	return 44
}

func (*NewOrderSingle514) PriceSinceVersion() uint16 {
	return 0
}

func (n *NewOrderSingle514) PriceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.PriceSinceVersion()
}

func (*NewOrderSingle514) PriceDeprecated() uint16 {
	return 0
}

func (*NewOrderSingle514) PriceMetaAttribute(meta int) string {
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

func (*NewOrderSingle514) OrderQtyId() uint16 {
	return 38
}

func (*NewOrderSingle514) OrderQtySinceVersion() uint16 {
	return 0
}

func (n *NewOrderSingle514) OrderQtyInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.OrderQtySinceVersion()
}

func (*NewOrderSingle514) OrderQtyDeprecated() uint16 {
	return 0
}

func (*NewOrderSingle514) OrderQtyMetaAttribute(meta int) string {
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

func (*NewOrderSingle514) OrderQtyMinValue() uint32 {
	return 0
}

func (*NewOrderSingle514) OrderQtyMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*NewOrderSingle514) OrderQtyNullValue() uint32 {
	return math.MaxUint32
}

func (*NewOrderSingle514) SecurityIDId() uint16 {
	return 48
}

func (*NewOrderSingle514) SecurityIDSinceVersion() uint16 {
	return 0
}

func (n *NewOrderSingle514) SecurityIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.SecurityIDSinceVersion()
}

func (*NewOrderSingle514) SecurityIDDeprecated() uint16 {
	return 0
}

func (*NewOrderSingle514) SecurityIDMetaAttribute(meta int) string {
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

func (*NewOrderSingle514) SecurityIDMinValue() int32 {
	return math.MinInt32 + 1
}

func (*NewOrderSingle514) SecurityIDMaxValue() int32 {
	return math.MaxInt32
}

func (*NewOrderSingle514) SecurityIDNullValue() int32 {
	return math.MinInt32
}

func (*NewOrderSingle514) SideId() uint16 {
	return 54
}

func (*NewOrderSingle514) SideSinceVersion() uint16 {
	return 0
}

func (n *NewOrderSingle514) SideInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.SideSinceVersion()
}

func (*NewOrderSingle514) SideDeprecated() uint16 {
	return 0
}

func (*NewOrderSingle514) SideMetaAttribute(meta int) string {
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

func (*NewOrderSingle514) SeqNumId() uint16 {
	return 9726
}

func (*NewOrderSingle514) SeqNumSinceVersion() uint16 {
	return 0
}

func (n *NewOrderSingle514) SeqNumInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.SeqNumSinceVersion()
}

func (*NewOrderSingle514) SeqNumDeprecated() uint16 {
	return 0
}

func (*NewOrderSingle514) SeqNumMetaAttribute(meta int) string {
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

func (*NewOrderSingle514) SeqNumMinValue() uint32 {
	return 0
}

func (*NewOrderSingle514) SeqNumMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*NewOrderSingle514) SeqNumNullValue() uint32 {
	return math.MaxUint32
}

func (*NewOrderSingle514) SenderIDId() uint16 {
	return 5392
}

func (*NewOrderSingle514) SenderIDSinceVersion() uint16 {
	return 0
}

func (n *NewOrderSingle514) SenderIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.SenderIDSinceVersion()
}

func (*NewOrderSingle514) SenderIDDeprecated() uint16 {
	return 0
}

func (*NewOrderSingle514) SenderIDMetaAttribute(meta int) string {
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

func (*NewOrderSingle514) SenderIDMinValue() byte {
	return byte(32)
}

func (*NewOrderSingle514) SenderIDMaxValue() byte {
	return byte(126)
}

func (*NewOrderSingle514) SenderIDNullValue() byte {
	return 0
}

func (n *NewOrderSingle514) SenderIDCharacterEncoding() string {
	return "US-ASCII"
}

func (*NewOrderSingle514) ClOrdIDId() uint16 {
	return 11
}

func (*NewOrderSingle514) ClOrdIDSinceVersion() uint16 {
	return 0
}

func (n *NewOrderSingle514) ClOrdIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.ClOrdIDSinceVersion()
}

func (*NewOrderSingle514) ClOrdIDDeprecated() uint16 {
	return 0
}

func (*NewOrderSingle514) ClOrdIDMetaAttribute(meta int) string {
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

func (*NewOrderSingle514) ClOrdIDMinValue() byte {
	return byte(32)
}

func (*NewOrderSingle514) ClOrdIDMaxValue() byte {
	return byte(126)
}

func (*NewOrderSingle514) ClOrdIDNullValue() byte {
	return 0
}

func (n *NewOrderSingle514) ClOrdIDCharacterEncoding() string {
	return "US-ASCII"
}

func (*NewOrderSingle514) PartyDetailsListReqIDId() uint16 {
	return 1505
}

func (*NewOrderSingle514) PartyDetailsListReqIDSinceVersion() uint16 {
	return 0
}

func (n *NewOrderSingle514) PartyDetailsListReqIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.PartyDetailsListReqIDSinceVersion()
}

func (*NewOrderSingle514) PartyDetailsListReqIDDeprecated() uint16 {
	return 0
}

func (*NewOrderSingle514) PartyDetailsListReqIDMetaAttribute(meta int) string {
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

func (*NewOrderSingle514) PartyDetailsListReqIDMinValue() uint64 {
	return 0
}

func (*NewOrderSingle514) PartyDetailsListReqIDMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*NewOrderSingle514) PartyDetailsListReqIDNullValue() uint64 {
	return math.MaxUint64
}

func (*NewOrderSingle514) OrderRequestIDId() uint16 {
	return 2422
}

func (*NewOrderSingle514) OrderRequestIDSinceVersion() uint16 {
	return 0
}

func (n *NewOrderSingle514) OrderRequestIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.OrderRequestIDSinceVersion()
}

func (*NewOrderSingle514) OrderRequestIDDeprecated() uint16 {
	return 0
}

func (*NewOrderSingle514) OrderRequestIDMetaAttribute(meta int) string {
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

func (*NewOrderSingle514) OrderRequestIDMinValue() uint64 {
	return 0
}

func (*NewOrderSingle514) OrderRequestIDMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*NewOrderSingle514) OrderRequestIDNullValue() uint64 {
	return math.MaxUint64
}

func (*NewOrderSingle514) SendingTimeEpochId() uint16 {
	return 5297
}

func (*NewOrderSingle514) SendingTimeEpochSinceVersion() uint16 {
	return 0
}

func (n *NewOrderSingle514) SendingTimeEpochInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.SendingTimeEpochSinceVersion()
}

func (*NewOrderSingle514) SendingTimeEpochDeprecated() uint16 {
	return 0
}

func (*NewOrderSingle514) SendingTimeEpochMetaAttribute(meta int) string {
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

func (*NewOrderSingle514) SendingTimeEpochMinValue() uint64 {
	return 0
}

func (*NewOrderSingle514) SendingTimeEpochMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*NewOrderSingle514) SendingTimeEpochNullValue() uint64 {
	return math.MaxUint64
}

func (*NewOrderSingle514) StopPxId() uint16 {
	return 99
}

func (*NewOrderSingle514) StopPxSinceVersion() uint16 {
	return 0
}

func (n *NewOrderSingle514) StopPxInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.StopPxSinceVersion()
}

func (*NewOrderSingle514) StopPxDeprecated() uint16 {
	return 0
}

func (*NewOrderSingle514) StopPxMetaAttribute(meta int) string {
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

func (*NewOrderSingle514) LocationId() uint16 {
	return 9537
}

func (*NewOrderSingle514) LocationSinceVersion() uint16 {
	return 0
}

func (n *NewOrderSingle514) LocationInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.LocationSinceVersion()
}

func (*NewOrderSingle514) LocationDeprecated() uint16 {
	return 0
}

func (*NewOrderSingle514) LocationMetaAttribute(meta int) string {
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

func (*NewOrderSingle514) LocationMinValue() byte {
	return byte(32)
}

func (*NewOrderSingle514) LocationMaxValue() byte {
	return byte(126)
}

func (*NewOrderSingle514) LocationNullValue() byte {
	return 0
}

func (n *NewOrderSingle514) LocationCharacterEncoding() string {
	return "US-ASCII"
}

func (*NewOrderSingle514) MinQtyId() uint16 {
	return 110
}

func (*NewOrderSingle514) MinQtySinceVersion() uint16 {
	return 0
}

func (n *NewOrderSingle514) MinQtyInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.MinQtySinceVersion()
}

func (*NewOrderSingle514) MinQtyDeprecated() uint16 {
	return 0
}

func (*NewOrderSingle514) MinQtyMetaAttribute(meta int) string {
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

func (*NewOrderSingle514) MinQtyMinValue() uint32 {
	return 0
}

func (*NewOrderSingle514) MinQtyMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*NewOrderSingle514) MinQtyNullValue() uint32 {
	return 4294967295
}

func (*NewOrderSingle514) DisplayQtyId() uint16 {
	return 1138
}

func (*NewOrderSingle514) DisplayQtySinceVersion() uint16 {
	return 0
}

func (n *NewOrderSingle514) DisplayQtyInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.DisplayQtySinceVersion()
}

func (*NewOrderSingle514) DisplayQtyDeprecated() uint16 {
	return 0
}

func (*NewOrderSingle514) DisplayQtyMetaAttribute(meta int) string {
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

func (*NewOrderSingle514) DisplayQtyMinValue() uint32 {
	return 0
}

func (*NewOrderSingle514) DisplayQtyMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*NewOrderSingle514) DisplayQtyNullValue() uint32 {
	return 4294967295
}

func (*NewOrderSingle514) ExpireDateId() uint16 {
	return 432
}

func (*NewOrderSingle514) ExpireDateSinceVersion() uint16 {
	return 0
}

func (n *NewOrderSingle514) ExpireDateInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.ExpireDateSinceVersion()
}

func (*NewOrderSingle514) ExpireDateDeprecated() uint16 {
	return 0
}

func (*NewOrderSingle514) ExpireDateMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return "LocalMktDate"
	case 4:
		return "optional"
	}
	return ""
}

func (*NewOrderSingle514) ExpireDateMinValue() uint16 {
	return 0
}

func (*NewOrderSingle514) ExpireDateMaxValue() uint16 {
	return math.MaxUint16 - 1
}

func (*NewOrderSingle514) ExpireDateNullValue() uint16 {
	return 65535
}

func (*NewOrderSingle514) OrdTypeId() uint16 {
	return 40
}

func (*NewOrderSingle514) OrdTypeSinceVersion() uint16 {
	return 0
}

func (n *NewOrderSingle514) OrdTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.OrdTypeSinceVersion()
}

func (*NewOrderSingle514) OrdTypeDeprecated() uint16 {
	return 0
}

func (*NewOrderSingle514) OrdTypeMetaAttribute(meta int) string {
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

func (*NewOrderSingle514) TimeInForceId() uint16 {
	return 59
}

func (*NewOrderSingle514) TimeInForceSinceVersion() uint16 {
	return 0
}

func (n *NewOrderSingle514) TimeInForceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.TimeInForceSinceVersion()
}

func (*NewOrderSingle514) TimeInForceDeprecated() uint16 {
	return 0
}

func (*NewOrderSingle514) TimeInForceMetaAttribute(meta int) string {
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

func (*NewOrderSingle514) ManualOrderIndicatorId() uint16 {
	return 1028
}

func (*NewOrderSingle514) ManualOrderIndicatorSinceVersion() uint16 {
	return 0
}

func (n *NewOrderSingle514) ManualOrderIndicatorInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.ManualOrderIndicatorSinceVersion()
}

func (*NewOrderSingle514) ManualOrderIndicatorDeprecated() uint16 {
	return 0
}

func (*NewOrderSingle514) ManualOrderIndicatorMetaAttribute(meta int) string {
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

func (*NewOrderSingle514) ExecInstId() uint16 {
	return 18
}

func (*NewOrderSingle514) ExecInstSinceVersion() uint16 {
	return 0
}

func (n *NewOrderSingle514) ExecInstInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.ExecInstSinceVersion()
}

func (*NewOrderSingle514) ExecInstDeprecated() uint16 {
	return 0
}

func (*NewOrderSingle514) ExecInstMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return "MultipleCharValue"
	case 4:
		return "required"
	}
	return ""
}

func (*NewOrderSingle514) ExecutionModeId() uint16 {
	return 5906
}

func (*NewOrderSingle514) ExecutionModeSinceVersion() uint16 {
	return 0
}

func (n *NewOrderSingle514) ExecutionModeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.ExecutionModeSinceVersion()
}

func (*NewOrderSingle514) ExecutionModeDeprecated() uint16 {
	return 0
}

func (*NewOrderSingle514) ExecutionModeMetaAttribute(meta int) string {
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

func (*NewOrderSingle514) LiquidityFlagId() uint16 {
	return 9373
}

func (*NewOrderSingle514) LiquidityFlagSinceVersion() uint16 {
	return 0
}

func (n *NewOrderSingle514) LiquidityFlagInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.LiquidityFlagSinceVersion()
}

func (*NewOrderSingle514) LiquidityFlagDeprecated() uint16 {
	return 0
}

func (*NewOrderSingle514) LiquidityFlagMetaAttribute(meta int) string {
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

func (*NewOrderSingle514) ManagedOrderId() uint16 {
	return 6881
}

func (*NewOrderSingle514) ManagedOrderSinceVersion() uint16 {
	return 0
}

func (n *NewOrderSingle514) ManagedOrderInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.ManagedOrderSinceVersion()
}

func (*NewOrderSingle514) ManagedOrderDeprecated() uint16 {
	return 0
}

func (*NewOrderSingle514) ManagedOrderMetaAttribute(meta int) string {
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

func (*NewOrderSingle514) ShortSaleTypeId() uint16 {
	return 5409
}

func (*NewOrderSingle514) ShortSaleTypeSinceVersion() uint16 {
	return 0
}

func (n *NewOrderSingle514) ShortSaleTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.ShortSaleTypeSinceVersion()
}

func (*NewOrderSingle514) ShortSaleTypeDeprecated() uint16 {
	return 0
}

func (*NewOrderSingle514) ShortSaleTypeMetaAttribute(meta int) string {
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
