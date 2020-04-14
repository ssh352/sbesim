// Generated SBE (Simple Binary Encoding) message codec

package fix

import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
)

type OrderCancelRequest struct {
	Account              [12]byte
	ClOrdID              [20]byte
	OrderID              int64
	OrigClOrdID          [20]byte
	Side                 SideEnum
	Symbol               [6]byte
	TransactTime         uint64
	ManualOrderIndicator BooleanTypeEnum
	SecurityDesc         [20]byte
	SecurityType         [3]byte
	CorrelationClOrdID   [20]byte
}

func (o *OrderCancelRequest) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := o.RangeCheck(o.SbeSchemaVersion(), o.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	if err := _m.WriteBytes(_w, o.Account[:]); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, o.ClOrdID[:]); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, o.OrderID); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, o.OrigClOrdID[:]); err != nil {
		return err
	}
	if err := o.Side.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, o.Symbol[:]); err != nil {
		return err
	}
	if err := _m.WriteUint64(_w, o.TransactTime); err != nil {
		return err
	}
	if err := o.ManualOrderIndicator.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, o.SecurityDesc[:]); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, o.SecurityType[:]); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, o.CorrelationClOrdID[:]); err != nil {
		return err
	}
	return nil
}

func (o *OrderCancelRequest) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	if !o.AccountInActingVersion(actingVersion) {
		for idx := 0; idx < 12; idx++ {
			o.Account[idx] = o.AccountNullValue()
		}
	} else {
		if err := _m.ReadBytes(_r, o.Account[:]); err != nil {
			return err
		}
	}
	if !o.ClOrdIDInActingVersion(actingVersion) {
		for idx := 0; idx < 20; idx++ {
			o.ClOrdID[idx] = o.ClOrdIDNullValue()
		}
	} else {
		if err := _m.ReadBytes(_r, o.ClOrdID[:]); err != nil {
			return err
		}
	}
	if !o.OrderIDInActingVersion(actingVersion) {
		o.OrderID = o.OrderIDNullValue()
	} else {
		if err := _m.ReadInt64(_r, &o.OrderID); err != nil {
			return err
		}
	}
	if !o.OrigClOrdIDInActingVersion(actingVersion) {
		for idx := 0; idx < 20; idx++ {
			o.OrigClOrdID[idx] = o.OrigClOrdIDNullValue()
		}
	} else {
		if err := _m.ReadBytes(_r, o.OrigClOrdID[:]); err != nil {
			return err
		}
	}
	if o.SideInActingVersion(actingVersion) {
		if err := o.Side.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if !o.SymbolInActingVersion(actingVersion) {
		for idx := 0; idx < 6; idx++ {
			o.Symbol[idx] = o.SymbolNullValue()
		}
	} else {
		if err := _m.ReadBytes(_r, o.Symbol[:]); err != nil {
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
	if o.ManualOrderIndicatorInActingVersion(actingVersion) {
		if err := o.ManualOrderIndicator.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if !o.SecurityDescInActingVersion(actingVersion) {
		for idx := 0; idx < 20; idx++ {
			o.SecurityDesc[idx] = o.SecurityDescNullValue()
		}
	} else {
		if err := _m.ReadBytes(_r, o.SecurityDesc[:]); err != nil {
			return err
		}
	}
	if !o.SecurityTypeInActingVersion(actingVersion) {
		for idx := 0; idx < 3; idx++ {
			o.SecurityType[idx] = o.SecurityTypeNullValue()
		}
	} else {
		if err := _m.ReadBytes(_r, o.SecurityType[:]); err != nil {
			return err
		}
	}
	if !o.CorrelationClOrdIDInActingVersion(actingVersion) {
		for idx := 0; idx < 20; idx++ {
			o.CorrelationClOrdID[idx] = o.CorrelationClOrdIDNullValue()
		}
	} else {
		if err := _m.ReadBytes(_r, o.CorrelationClOrdID[:]); err != nil {
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

func (o *OrderCancelRequest) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if o.AccountInActingVersion(actingVersion) {
		for idx := 0; idx < 12; idx++ {
			if o.Account[idx] < o.AccountMinValue() || o.Account[idx] > o.AccountMaxValue() {
				return fmt.Errorf("Range check failed on o.Account[%d] (%v < %v > %v)", idx, o.AccountMinValue(), o.Account[idx], o.AccountMaxValue())
			}
		}
	}
	if o.ClOrdIDInActingVersion(actingVersion) {
		for idx := 0; idx < 20; idx++ {
			if o.ClOrdID[idx] < o.ClOrdIDMinValue() || o.ClOrdID[idx] > o.ClOrdIDMaxValue() {
				return fmt.Errorf("Range check failed on o.ClOrdID[%d] (%v < %v > %v)", idx, o.ClOrdIDMinValue(), o.ClOrdID[idx], o.ClOrdIDMaxValue())
			}
		}
	}
	if o.OrderIDInActingVersion(actingVersion) {
		if o.OrderID < o.OrderIDMinValue() || o.OrderID > o.OrderIDMaxValue() {
			return fmt.Errorf("Range check failed on o.OrderID (%v < %v > %v)", o.OrderIDMinValue(), o.OrderID, o.OrderIDMaxValue())
		}
	}
	if o.OrigClOrdIDInActingVersion(actingVersion) {
		for idx := 0; idx < 20; idx++ {
			if o.OrigClOrdID[idx] < o.OrigClOrdIDMinValue() || o.OrigClOrdID[idx] > o.OrigClOrdIDMaxValue() {
				return fmt.Errorf("Range check failed on o.OrigClOrdID[%d] (%v < %v > %v)", idx, o.OrigClOrdIDMinValue(), o.OrigClOrdID[idx], o.OrigClOrdIDMaxValue())
			}
		}
	}
	if err := o.Side.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if o.SymbolInActingVersion(actingVersion) {
		for idx := 0; idx < 6; idx++ {
			if o.Symbol[idx] < o.SymbolMinValue() || o.Symbol[idx] > o.SymbolMaxValue() {
				return fmt.Errorf("Range check failed on o.Symbol[%d] (%v < %v > %v)", idx, o.SymbolMinValue(), o.Symbol[idx], o.SymbolMaxValue())
			}
		}
	}
	if o.TransactTimeInActingVersion(actingVersion) {
		if o.TransactTime < o.TransactTimeMinValue() || o.TransactTime > o.TransactTimeMaxValue() {
			return fmt.Errorf("Range check failed on o.TransactTime (%v < %v > %v)", o.TransactTimeMinValue(), o.TransactTime, o.TransactTimeMaxValue())
		}
	}
	if err := o.ManualOrderIndicator.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if o.SecurityDescInActingVersion(actingVersion) {
		for idx := 0; idx < 20; idx++ {
			if o.SecurityDesc[idx] < o.SecurityDescMinValue() || o.SecurityDesc[idx] > o.SecurityDescMaxValue() {
				return fmt.Errorf("Range check failed on o.SecurityDesc[%d] (%v < %v > %v)", idx, o.SecurityDescMinValue(), o.SecurityDesc[idx], o.SecurityDescMaxValue())
			}
		}
	}
	if o.SecurityTypeInActingVersion(actingVersion) {
		for idx := 0; idx < 3; idx++ {
			if o.SecurityType[idx] < o.SecurityTypeMinValue() || o.SecurityType[idx] > o.SecurityTypeMaxValue() {
				return fmt.Errorf("Range check failed on o.SecurityType[%d] (%v < %v > %v)", idx, o.SecurityTypeMinValue(), o.SecurityType[idx], o.SecurityTypeMaxValue())
			}
		}
	}
	if o.CorrelationClOrdIDInActingVersion(actingVersion) {
		for idx := 0; idx < 20; idx++ {
			if o.CorrelationClOrdID[idx] < o.CorrelationClOrdIDMinValue() || o.CorrelationClOrdID[idx] > o.CorrelationClOrdIDMaxValue() {
				return fmt.Errorf("Range check failed on o.CorrelationClOrdID[%d] (%v < %v > %v)", idx, o.CorrelationClOrdIDMinValue(), o.CorrelationClOrdID[idx], o.CorrelationClOrdIDMaxValue())
			}
		}
	}
	return nil
}

func OrderCancelRequestInit(o *OrderCancelRequest) {
	return
}

func (*OrderCancelRequest) SbeBlockLength() (blockLength uint16) {
	return 119
}

func (*OrderCancelRequest) SbeTemplateId() (templateId uint16) {
	return 70
}

func (*OrderCancelRequest) SbeSchemaId() (schemaId uint16) {
	return 1
}

func (*OrderCancelRequest) SbeSchemaVersion() (schemaVersion uint16) {
	return 1
}

func (*OrderCancelRequest) SbeSemanticType() (semanticType []byte) {
	return []byte("F")
}

func (*OrderCancelRequest) AccountId() uint16 {
	return 1
}

func (*OrderCancelRequest) AccountSinceVersion() uint16 {
	return 0
}

func (o *OrderCancelRequest) AccountInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.AccountSinceVersion()
}

func (*OrderCancelRequest) AccountDeprecated() uint16 {
	return 0
}

func (*OrderCancelRequest) AccountMetaAttribute(meta int) string {
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

func (*OrderCancelRequest) AccountMinValue() byte {
	return byte(32)
}

func (*OrderCancelRequest) AccountMaxValue() byte {
	return byte(126)
}

func (*OrderCancelRequest) AccountNullValue() byte {
	return 0
}

func (o *OrderCancelRequest) AccountCharacterEncoding() string {
	return "US-ASCII"
}

func (*OrderCancelRequest) ClOrdIDId() uint16 {
	return 11
}

func (*OrderCancelRequest) ClOrdIDSinceVersion() uint16 {
	return 0
}

func (o *OrderCancelRequest) ClOrdIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.ClOrdIDSinceVersion()
}

func (*OrderCancelRequest) ClOrdIDDeprecated() uint16 {
	return 0
}

func (*OrderCancelRequest) ClOrdIDMetaAttribute(meta int) string {
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

func (*OrderCancelRequest) ClOrdIDMinValue() byte {
	return byte(32)
}

func (*OrderCancelRequest) ClOrdIDMaxValue() byte {
	return byte(126)
}

func (*OrderCancelRequest) ClOrdIDNullValue() byte {
	return 0
}

func (o *OrderCancelRequest) ClOrdIDCharacterEncoding() string {
	return "US-ASCII"
}

func (*OrderCancelRequest) OrderIDId() uint16 {
	return 37
}

func (*OrderCancelRequest) OrderIDSinceVersion() uint16 {
	return 0
}

func (o *OrderCancelRequest) OrderIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.OrderIDSinceVersion()
}

func (*OrderCancelRequest) OrderIDDeprecated() uint16 {
	return 0
}

func (*OrderCancelRequest) OrderIDMetaAttribute(meta int) string {
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

func (*OrderCancelRequest) OrderIDMinValue() int64 {
	return math.MinInt64 + 1
}

func (*OrderCancelRequest) OrderIDMaxValue() int64 {
	return math.MaxInt64
}

func (*OrderCancelRequest) OrderIDNullValue() int64 {
	return math.MinInt64
}

func (*OrderCancelRequest) OrigClOrdIDId() uint16 {
	return 41
}

func (*OrderCancelRequest) OrigClOrdIDSinceVersion() uint16 {
	return 0
}

func (o *OrderCancelRequest) OrigClOrdIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.OrigClOrdIDSinceVersion()
}

func (*OrderCancelRequest) OrigClOrdIDDeprecated() uint16 {
	return 0
}

func (*OrderCancelRequest) OrigClOrdIDMetaAttribute(meta int) string {
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

func (*OrderCancelRequest) OrigClOrdIDMinValue() byte {
	return byte(32)
}

func (*OrderCancelRequest) OrigClOrdIDMaxValue() byte {
	return byte(126)
}

func (*OrderCancelRequest) OrigClOrdIDNullValue() byte {
	return 0
}

func (o *OrderCancelRequest) OrigClOrdIDCharacterEncoding() string {
	return "US-ASCII"
}

func (*OrderCancelRequest) SideId() uint16 {
	return 54
}

func (*OrderCancelRequest) SideSinceVersion() uint16 {
	return 0
}

func (o *OrderCancelRequest) SideInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.SideSinceVersion()
}

func (*OrderCancelRequest) SideDeprecated() uint16 {
	return 0
}

func (*OrderCancelRequest) SideMetaAttribute(meta int) string {
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

func (*OrderCancelRequest) SymbolId() uint16 {
	return 55
}

func (*OrderCancelRequest) SymbolSinceVersion() uint16 {
	return 0
}

func (o *OrderCancelRequest) SymbolInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.SymbolSinceVersion()
}

func (*OrderCancelRequest) SymbolDeprecated() uint16 {
	return 0
}

func (*OrderCancelRequest) SymbolMetaAttribute(meta int) string {
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

func (*OrderCancelRequest) SymbolMinValue() byte {
	return byte(32)
}

func (*OrderCancelRequest) SymbolMaxValue() byte {
	return byte(126)
}

func (*OrderCancelRequest) SymbolNullValue() byte {
	return 0
}

func (o *OrderCancelRequest) SymbolCharacterEncoding() string {
	return "US-ASCII"
}

func (*OrderCancelRequest) TransactTimeId() uint16 {
	return 60
}

func (*OrderCancelRequest) TransactTimeSinceVersion() uint16 {
	return 0
}

func (o *OrderCancelRequest) TransactTimeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.TransactTimeSinceVersion()
}

func (*OrderCancelRequest) TransactTimeDeprecated() uint16 {
	return 0
}

func (*OrderCancelRequest) TransactTimeMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return "nanosecond"
	case 3:
		return "UTCTimestamp"
	case 4:
		return "required"
	}
	return ""
}

func (*OrderCancelRequest) TransactTimeMinValue() uint64 {
	return 0
}

func (*OrderCancelRequest) TransactTimeMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*OrderCancelRequest) TransactTimeNullValue() uint64 {
	return math.MaxUint64
}

func (*OrderCancelRequest) ManualOrderIndicatorId() uint16 {
	return 1028
}

func (*OrderCancelRequest) ManualOrderIndicatorSinceVersion() uint16 {
	return 0
}

func (o *OrderCancelRequest) ManualOrderIndicatorInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.ManualOrderIndicatorSinceVersion()
}

func (*OrderCancelRequest) ManualOrderIndicatorDeprecated() uint16 {
	return 0
}

func (*OrderCancelRequest) ManualOrderIndicatorMetaAttribute(meta int) string {
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

func (*OrderCancelRequest) SecurityDescId() uint16 {
	return 107
}

func (*OrderCancelRequest) SecurityDescSinceVersion() uint16 {
	return 0
}

func (o *OrderCancelRequest) SecurityDescInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.SecurityDescSinceVersion()
}

func (*OrderCancelRequest) SecurityDescDeprecated() uint16 {
	return 0
}

func (*OrderCancelRequest) SecurityDescMetaAttribute(meta int) string {
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

func (*OrderCancelRequest) SecurityDescMinValue() byte {
	return byte(32)
}

func (*OrderCancelRequest) SecurityDescMaxValue() byte {
	return byte(126)
}

func (*OrderCancelRequest) SecurityDescNullValue() byte {
	return 0
}

func (o *OrderCancelRequest) SecurityDescCharacterEncoding() string {
	return "US-ASCII"
}

func (*OrderCancelRequest) SecurityTypeId() uint16 {
	return 167
}

func (*OrderCancelRequest) SecurityTypeSinceVersion() uint16 {
	return 0
}

func (o *OrderCancelRequest) SecurityTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.SecurityTypeSinceVersion()
}

func (*OrderCancelRequest) SecurityTypeDeprecated() uint16 {
	return 0
}

func (*OrderCancelRequest) SecurityTypeMetaAttribute(meta int) string {
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

func (*OrderCancelRequest) SecurityTypeMinValue() byte {
	return byte(32)
}

func (*OrderCancelRequest) SecurityTypeMaxValue() byte {
	return byte(126)
}

func (*OrderCancelRequest) SecurityTypeNullValue() byte {
	return 0
}

func (o *OrderCancelRequest) SecurityTypeCharacterEncoding() string {
	return "US-ASCII"
}

func (*OrderCancelRequest) CorrelationClOrdIDId() uint16 {
	return 9717
}

func (*OrderCancelRequest) CorrelationClOrdIDSinceVersion() uint16 {
	return 0
}

func (o *OrderCancelRequest) CorrelationClOrdIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.CorrelationClOrdIDSinceVersion()
}

func (*OrderCancelRequest) CorrelationClOrdIDDeprecated() uint16 {
	return 0
}

func (*OrderCancelRequest) CorrelationClOrdIDMetaAttribute(meta int) string {
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

func (*OrderCancelRequest) CorrelationClOrdIDMinValue() byte {
	return byte(32)
}

func (*OrderCancelRequest) CorrelationClOrdIDMaxValue() byte {
	return byte(126)
}

func (*OrderCancelRequest) CorrelationClOrdIDNullValue() byte {
	return 0
}

func (o *OrderCancelRequest) CorrelationClOrdIDCharacterEncoding() string {
	return "US-ASCII"
}
