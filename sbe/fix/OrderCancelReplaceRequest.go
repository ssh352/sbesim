// Generated SBE (Simple Binary Encoding) message codec

package fix

import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
)

type OrderCancelReplaceRequest struct {
	Account               [12]byte
	ClOrdID               [20]byte
	OrderID               int64
	HandInst              HandInstEnum
	OrderQty              IntQty32
	CustOrderHandlingInst CustOrderHandlingInstEnum
	OrdType               OrdTypeEnum
	OrigClOrdID           [20]byte
	Price                 OptionalPrice
	Side                  SideEnum
	Symbol                [6]byte
	Test                  [18]byte
	TimeInForce           TimeInForceEnum
	ManualOrderIndicator  BooleanTypeEnum
	TransactTime          uint64
	NoAllocs              NoAllocsEnum
	AllocAccount          [10]byte
	StopPx                OptionalPrice
	SecurityDesc          [20]byte
	MinQty                IntQty32
	SecurityType          [3]byte
	CustomerOrFirm        CustomerOrFirmEnum
	MaxShow               IntQty32
	ExpireDate            uint16
	SelfMatchPreventionID [12]byte
	CtiCode               CtiCodeEnum
	GiveUpFirm            [3]byte
	CmtaGiveupCD          [2]byte
	CorrelationClOrdID    [20]byte
	OFMOverride           OFMOverrideEnum
}

func (o *OrderCancelReplaceRequest) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
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
	if err := o.HandInst.Encode(_m, _w); err != nil {
		return err
	}
	if err := o.OrderQty.Encode(_m, _w); err != nil {
		return err
	}
	if err := o.CustOrderHandlingInst.Encode(_m, _w); err != nil {
		return err
	}
	if err := o.OrdType.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, o.OrigClOrdID[:]); err != nil {
		return err
	}
	if err := o.Price.Encode(_m, _w); err != nil {
		return err
	}
	if err := o.Side.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, o.Symbol[:]); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, o.Test[:]); err != nil {
		return err
	}
	if err := o.TimeInForce.Encode(_m, _w); err != nil {
		return err
	}
	if err := o.ManualOrderIndicator.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteUint64(_w, o.TransactTime); err != nil {
		return err
	}
	if err := o.NoAllocs.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, o.AllocAccount[:]); err != nil {
		return err
	}
	if err := o.StopPx.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, o.SecurityDesc[:]); err != nil {
		return err
	}
	if err := o.MinQty.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, o.SecurityType[:]); err != nil {
		return err
	}
	if err := o.CustomerOrFirm.Encode(_m, _w); err != nil {
		return err
	}
	if err := o.MaxShow.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteUint16(_w, o.ExpireDate); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, o.SelfMatchPreventionID[:]); err != nil {
		return err
	}
	if err := o.CtiCode.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, o.GiveUpFirm[:]); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, o.CmtaGiveupCD[:]); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, o.CorrelationClOrdID[:]); err != nil {
		return err
	}
	if err := o.OFMOverride.Encode(_m, _w); err != nil {
		return err
	}
	return nil
}

func (o *OrderCancelReplaceRequest) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
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
	if o.HandInstInActingVersion(actingVersion) {
		if err := o.HandInst.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if o.OrderQtyInActingVersion(actingVersion) {
		if err := o.OrderQty.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if o.CustOrderHandlingInstInActingVersion(actingVersion) {
		if err := o.CustOrderHandlingInst.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if o.OrdTypeInActingVersion(actingVersion) {
		if err := o.OrdType.Decode(_m, _r, actingVersion); err != nil {
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
	if o.PriceInActingVersion(actingVersion) {
		if err := o.Price.Decode(_m, _r, actingVersion); err != nil {
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
	if !o.TestInActingVersion(actingVersion) {
		for idx := 0; idx < 18; idx++ {
			o.Test[idx] = o.TestNullValue()
		}
	} else {
		if err := _m.ReadBytes(_r, o.Test[:]); err != nil {
			return err
		}
	}
	if o.TimeInForceInActingVersion(actingVersion) {
		if err := o.TimeInForce.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if o.ManualOrderIndicatorInActingVersion(actingVersion) {
		if err := o.ManualOrderIndicator.Decode(_m, _r, actingVersion); err != nil {
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
	if o.NoAllocsInActingVersion(actingVersion) {
		if err := o.NoAllocs.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if !o.AllocAccountInActingVersion(actingVersion) {
		for idx := 0; idx < 10; idx++ {
			o.AllocAccount[idx] = o.AllocAccountNullValue()
		}
	} else {
		if err := _m.ReadBytes(_r, o.AllocAccount[:]); err != nil {
			return err
		}
	}
	if o.StopPxInActingVersion(actingVersion) {
		if err := o.StopPx.Decode(_m, _r, actingVersion); err != nil {
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
	if o.MinQtyInActingVersion(actingVersion) {
		if err := o.MinQty.Decode(_m, _r, actingVersion); err != nil {
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
	if o.CustomerOrFirmInActingVersion(actingVersion) {
		if err := o.CustomerOrFirm.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if o.MaxShowInActingVersion(actingVersion) {
		if err := o.MaxShow.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if !o.ExpireDateInActingVersion(actingVersion) {
		o.ExpireDate = o.ExpireDateNullValue()
	} else {
		if err := _m.ReadUint16(_r, &o.ExpireDate); err != nil {
			return err
		}
	}
	if !o.SelfMatchPreventionIDInActingVersion(actingVersion) {
		for idx := 0; idx < 12; idx++ {
			o.SelfMatchPreventionID[idx] = o.SelfMatchPreventionIDNullValue()
		}
	} else {
		if err := _m.ReadBytes(_r, o.SelfMatchPreventionID[:]); err != nil {
			return err
		}
	}
	if o.CtiCodeInActingVersion(actingVersion) {
		if err := o.CtiCode.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if !o.GiveUpFirmInActingVersion(actingVersion) {
		for idx := 0; idx < 3; idx++ {
			o.GiveUpFirm[idx] = o.GiveUpFirmNullValue()
		}
	} else {
		if err := _m.ReadBytes(_r, o.GiveUpFirm[:]); err != nil {
			return err
		}
	}
	if !o.CmtaGiveupCDInActingVersion(actingVersion) {
		for idx := 0; idx < 2; idx++ {
			o.CmtaGiveupCD[idx] = o.CmtaGiveupCDNullValue()
		}
	} else {
		if err := _m.ReadBytes(_r, o.CmtaGiveupCD[:]); err != nil {
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
	if o.OFMOverrideInActingVersion(actingVersion) {
		if err := o.OFMOverride.Decode(_m, _r, actingVersion); err != nil {
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

func (o *OrderCancelReplaceRequest) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
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
	if err := o.HandInst.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if err := o.CustOrderHandlingInst.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if err := o.OrdType.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
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
	if o.TestInActingVersion(actingVersion) {
		for idx := 0; idx < 18; idx++ {
			if o.Test[idx] < o.TestMinValue() || o.Test[idx] > o.TestMaxValue() {
				return fmt.Errorf("Range check failed on o.Test[%d] (%v < %v > %v)", idx, o.TestMinValue(), o.Test[idx], o.TestMaxValue())
			}
		}
	}
	if err := o.TimeInForce.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if err := o.ManualOrderIndicator.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if o.TransactTimeInActingVersion(actingVersion) {
		if o.TransactTime < o.TransactTimeMinValue() || o.TransactTime > o.TransactTimeMaxValue() {
			return fmt.Errorf("Range check failed on o.TransactTime (%v < %v > %v)", o.TransactTimeMinValue(), o.TransactTime, o.TransactTimeMaxValue())
		}
	}
	if err := o.NoAllocs.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if o.AllocAccountInActingVersion(actingVersion) {
		for idx := 0; idx < 10; idx++ {
			if o.AllocAccount[idx] < o.AllocAccountMinValue() || o.AllocAccount[idx] > o.AllocAccountMaxValue() {
				return fmt.Errorf("Range check failed on o.AllocAccount[%d] (%v < %v > %v)", idx, o.AllocAccountMinValue(), o.AllocAccount[idx], o.AllocAccountMaxValue())
			}
		}
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
	if err := o.CustomerOrFirm.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if o.ExpireDateInActingVersion(actingVersion) {
		if o.ExpireDate < o.ExpireDateMinValue() || o.ExpireDate > o.ExpireDateMaxValue() {
			return fmt.Errorf("Range check failed on o.ExpireDate (%v < %v > %v)", o.ExpireDateMinValue(), o.ExpireDate, o.ExpireDateMaxValue())
		}
	}
	if o.SelfMatchPreventionIDInActingVersion(actingVersion) {
		for idx := 0; idx < 12; idx++ {
			if o.SelfMatchPreventionID[idx] < o.SelfMatchPreventionIDMinValue() || o.SelfMatchPreventionID[idx] > o.SelfMatchPreventionIDMaxValue() {
				return fmt.Errorf("Range check failed on o.SelfMatchPreventionID[%d] (%v < %v > %v)", idx, o.SelfMatchPreventionIDMinValue(), o.SelfMatchPreventionID[idx], o.SelfMatchPreventionIDMaxValue())
			}
		}
	}
	if err := o.CtiCode.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if o.GiveUpFirmInActingVersion(actingVersion) {
		for idx := 0; idx < 3; idx++ {
			if o.GiveUpFirm[idx] < o.GiveUpFirmMinValue() || o.GiveUpFirm[idx] > o.GiveUpFirmMaxValue() {
				return fmt.Errorf("Range check failed on o.GiveUpFirm[%d] (%v < %v > %v)", idx, o.GiveUpFirmMinValue(), o.GiveUpFirm[idx], o.GiveUpFirmMaxValue())
			}
		}
	}
	if o.CmtaGiveupCDInActingVersion(actingVersion) {
		for idx := 0; idx < 2; idx++ {
			if o.CmtaGiveupCD[idx] < o.CmtaGiveupCDMinValue() || o.CmtaGiveupCD[idx] > o.CmtaGiveupCDMaxValue() {
				return fmt.Errorf("Range check failed on o.CmtaGiveupCD[%d] (%v < %v > %v)", idx, o.CmtaGiveupCDMinValue(), o.CmtaGiveupCD[idx], o.CmtaGiveupCDMaxValue())
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
	if err := o.OFMOverride.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	return nil
}

func OrderCancelReplaceRequestInit(o *OrderCancelReplaceRequest) {
	return
}

func (*OrderCancelReplaceRequest) SbeBlockLength() (blockLength uint16) {
	return 204
}

func (*OrderCancelReplaceRequest) SbeTemplateId() (templateId uint16) {
	return 71
}

func (*OrderCancelReplaceRequest) SbeSchemaId() (schemaId uint16) {
	return 1
}

func (*OrderCancelReplaceRequest) SbeSchemaVersion() (schemaVersion uint16) {
	return 1
}

func (*OrderCancelReplaceRequest) SbeSemanticType() (semanticType []byte) {
	return []byte("G")
}

func (*OrderCancelReplaceRequest) AccountId() uint16 {
	return 1
}

func (*OrderCancelReplaceRequest) AccountSinceVersion() uint16 {
	return 0
}

func (o *OrderCancelReplaceRequest) AccountInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.AccountSinceVersion()
}

func (*OrderCancelReplaceRequest) AccountDeprecated() uint16 {
	return 0
}

func (*OrderCancelReplaceRequest) AccountMetaAttribute(meta int) string {
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

func (*OrderCancelReplaceRequest) AccountMinValue() byte {
	return byte(32)
}

func (*OrderCancelReplaceRequest) AccountMaxValue() byte {
	return byte(126)
}

func (*OrderCancelReplaceRequest) AccountNullValue() byte {
	return 0
}

func (o *OrderCancelReplaceRequest) AccountCharacterEncoding() string {
	return "US-ASCII"
}

func (*OrderCancelReplaceRequest) ClOrdIDId() uint16 {
	return 11
}

func (*OrderCancelReplaceRequest) ClOrdIDSinceVersion() uint16 {
	return 0
}

func (o *OrderCancelReplaceRequest) ClOrdIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.ClOrdIDSinceVersion()
}

func (*OrderCancelReplaceRequest) ClOrdIDDeprecated() uint16 {
	return 0
}

func (*OrderCancelReplaceRequest) ClOrdIDMetaAttribute(meta int) string {
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

func (*OrderCancelReplaceRequest) ClOrdIDMinValue() byte {
	return byte(32)
}

func (*OrderCancelReplaceRequest) ClOrdIDMaxValue() byte {
	return byte(126)
}

func (*OrderCancelReplaceRequest) ClOrdIDNullValue() byte {
	return 0
}

func (o *OrderCancelReplaceRequest) ClOrdIDCharacterEncoding() string {
	return "US-ASCII"
}

func (*OrderCancelReplaceRequest) OrderIDId() uint16 {
	return 37
}

func (*OrderCancelReplaceRequest) OrderIDSinceVersion() uint16 {
	return 0
}

func (o *OrderCancelReplaceRequest) OrderIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.OrderIDSinceVersion()
}

func (*OrderCancelReplaceRequest) OrderIDDeprecated() uint16 {
	return 0
}

func (*OrderCancelReplaceRequest) OrderIDMetaAttribute(meta int) string {
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

func (*OrderCancelReplaceRequest) OrderIDMinValue() int64 {
	return math.MinInt64 + 1
}

func (*OrderCancelReplaceRequest) OrderIDMaxValue() int64 {
	return math.MaxInt64
}

func (*OrderCancelReplaceRequest) OrderIDNullValue() int64 {
	return math.MinInt64
}

func (*OrderCancelReplaceRequest) HandInstId() uint16 {
	return 21
}

func (*OrderCancelReplaceRequest) HandInstSinceVersion() uint16 {
	return 0
}

func (o *OrderCancelReplaceRequest) HandInstInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.HandInstSinceVersion()
}

func (*OrderCancelReplaceRequest) HandInstDeprecated() uint16 {
	return 0
}

func (*OrderCancelReplaceRequest) HandInstMetaAttribute(meta int) string {
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

func (*OrderCancelReplaceRequest) OrderQtyId() uint16 {
	return 38
}

func (*OrderCancelReplaceRequest) OrderQtySinceVersion() uint16 {
	return 0
}

func (o *OrderCancelReplaceRequest) OrderQtyInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.OrderQtySinceVersion()
}

func (*OrderCancelReplaceRequest) OrderQtyDeprecated() uint16 {
	return 0
}

func (*OrderCancelReplaceRequest) OrderQtyMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return "Qty"
	case 4:
		return "required"
	}
	return ""
}

func (*OrderCancelReplaceRequest) CustOrderHandlingInstId() uint16 {
	return 1031
}

func (*OrderCancelReplaceRequest) CustOrderHandlingInstSinceVersion() uint16 {
	return 0
}

func (o *OrderCancelReplaceRequest) CustOrderHandlingInstInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.CustOrderHandlingInstSinceVersion()
}

func (*OrderCancelReplaceRequest) CustOrderHandlingInstDeprecated() uint16 {
	return 0
}

func (*OrderCancelReplaceRequest) CustOrderHandlingInstMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return "char"
	case 4:
		return "optional"
	}
	return ""
}

func (*OrderCancelReplaceRequest) OrdTypeId() uint16 {
	return 40
}

func (*OrderCancelReplaceRequest) OrdTypeSinceVersion() uint16 {
	return 0
}

func (o *OrderCancelReplaceRequest) OrdTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.OrdTypeSinceVersion()
}

func (*OrderCancelReplaceRequest) OrdTypeDeprecated() uint16 {
	return 0
}

func (*OrderCancelReplaceRequest) OrdTypeMetaAttribute(meta int) string {
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

func (*OrderCancelReplaceRequest) OrigClOrdIDId() uint16 {
	return 41
}

func (*OrderCancelReplaceRequest) OrigClOrdIDSinceVersion() uint16 {
	return 0
}

func (o *OrderCancelReplaceRequest) OrigClOrdIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.OrigClOrdIDSinceVersion()
}

func (*OrderCancelReplaceRequest) OrigClOrdIDDeprecated() uint16 {
	return 0
}

func (*OrderCancelReplaceRequest) OrigClOrdIDMetaAttribute(meta int) string {
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

func (*OrderCancelReplaceRequest) OrigClOrdIDMinValue() byte {
	return byte(32)
}

func (*OrderCancelReplaceRequest) OrigClOrdIDMaxValue() byte {
	return byte(126)
}

func (*OrderCancelReplaceRequest) OrigClOrdIDNullValue() byte {
	return 0
}

func (o *OrderCancelReplaceRequest) OrigClOrdIDCharacterEncoding() string {
	return "US-ASCII"
}

func (*OrderCancelReplaceRequest) PriceId() uint16 {
	return 44
}

func (*OrderCancelReplaceRequest) PriceSinceVersion() uint16 {
	return 0
}

func (o *OrderCancelReplaceRequest) PriceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.PriceSinceVersion()
}

func (*OrderCancelReplaceRequest) PriceDeprecated() uint16 {
	return 0
}

func (*OrderCancelReplaceRequest) PriceMetaAttribute(meta int) string {
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

func (*OrderCancelReplaceRequest) SideId() uint16 {
	return 54
}

func (*OrderCancelReplaceRequest) SideSinceVersion() uint16 {
	return 0
}

func (o *OrderCancelReplaceRequest) SideInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.SideSinceVersion()
}

func (*OrderCancelReplaceRequest) SideDeprecated() uint16 {
	return 0
}

func (*OrderCancelReplaceRequest) SideMetaAttribute(meta int) string {
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

func (*OrderCancelReplaceRequest) SymbolId() uint16 {
	return 55
}

func (*OrderCancelReplaceRequest) SymbolSinceVersion() uint16 {
	return 0
}

func (o *OrderCancelReplaceRequest) SymbolInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.SymbolSinceVersion()
}

func (*OrderCancelReplaceRequest) SymbolDeprecated() uint16 {
	return 0
}

func (*OrderCancelReplaceRequest) SymbolMetaAttribute(meta int) string {
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

func (*OrderCancelReplaceRequest) SymbolMinValue() byte {
	return byte(32)
}

func (*OrderCancelReplaceRequest) SymbolMaxValue() byte {
	return byte(126)
}

func (*OrderCancelReplaceRequest) SymbolNullValue() byte {
	return 0
}

func (o *OrderCancelReplaceRequest) SymbolCharacterEncoding() string {
	return "US-ASCII"
}

func (*OrderCancelReplaceRequest) TestId() uint16 {
	return 58
}

func (*OrderCancelReplaceRequest) TestSinceVersion() uint16 {
	return 0
}

func (o *OrderCancelReplaceRequest) TestInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.TestSinceVersion()
}

func (*OrderCancelReplaceRequest) TestDeprecated() uint16 {
	return 0
}

func (*OrderCancelReplaceRequest) TestMetaAttribute(meta int) string {
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

func (*OrderCancelReplaceRequest) TestMinValue() byte {
	return byte(32)
}

func (*OrderCancelReplaceRequest) TestMaxValue() byte {
	return byte(126)
}

func (*OrderCancelReplaceRequest) TestNullValue() byte {
	return 0
}

func (o *OrderCancelReplaceRequest) TestCharacterEncoding() string {
	return "US-ASCII"
}

func (*OrderCancelReplaceRequest) TimeInForceId() uint16 {
	return 59
}

func (*OrderCancelReplaceRequest) TimeInForceSinceVersion() uint16 {
	return 0
}

func (o *OrderCancelReplaceRequest) TimeInForceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.TimeInForceSinceVersion()
}

func (*OrderCancelReplaceRequest) TimeInForceDeprecated() uint16 {
	return 0
}

func (*OrderCancelReplaceRequest) TimeInForceMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return "char"
	case 4:
		return "optional"
	}
	return ""
}

func (*OrderCancelReplaceRequest) ManualOrderIndicatorId() uint16 {
	return 1028
}

func (*OrderCancelReplaceRequest) ManualOrderIndicatorSinceVersion() uint16 {
	return 0
}

func (o *OrderCancelReplaceRequest) ManualOrderIndicatorInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.ManualOrderIndicatorSinceVersion()
}

func (*OrderCancelReplaceRequest) ManualOrderIndicatorDeprecated() uint16 {
	return 0
}

func (*OrderCancelReplaceRequest) ManualOrderIndicatorMetaAttribute(meta int) string {
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

func (*OrderCancelReplaceRequest) TransactTimeId() uint16 {
	return 60
}

func (*OrderCancelReplaceRequest) TransactTimeSinceVersion() uint16 {
	return 0
}

func (o *OrderCancelReplaceRequest) TransactTimeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.TransactTimeSinceVersion()
}

func (*OrderCancelReplaceRequest) TransactTimeDeprecated() uint16 {
	return 0
}

func (*OrderCancelReplaceRequest) TransactTimeMetaAttribute(meta int) string {
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

func (*OrderCancelReplaceRequest) TransactTimeMinValue() uint64 {
	return 0
}

func (*OrderCancelReplaceRequest) TransactTimeMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*OrderCancelReplaceRequest) TransactTimeNullValue() uint64 {
	return math.MaxUint64
}

func (*OrderCancelReplaceRequest) NoAllocsId() uint16 {
	return 78
}

func (*OrderCancelReplaceRequest) NoAllocsSinceVersion() uint16 {
	return 0
}

func (o *OrderCancelReplaceRequest) NoAllocsInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.NoAllocsSinceVersion()
}

func (*OrderCancelReplaceRequest) NoAllocsDeprecated() uint16 {
	return 0
}

func (*OrderCancelReplaceRequest) NoAllocsMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "optional"
	}
	return ""
}

func (*OrderCancelReplaceRequest) AllocAccountId() uint16 {
	return 79
}

func (*OrderCancelReplaceRequest) AllocAccountSinceVersion() uint16 {
	return 0
}

func (o *OrderCancelReplaceRequest) AllocAccountInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.AllocAccountSinceVersion()
}

func (*OrderCancelReplaceRequest) AllocAccountDeprecated() uint16 {
	return 0
}

func (*OrderCancelReplaceRequest) AllocAccountMetaAttribute(meta int) string {
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

func (*OrderCancelReplaceRequest) AllocAccountMinValue() byte {
	return byte(32)
}

func (*OrderCancelReplaceRequest) AllocAccountMaxValue() byte {
	return byte(126)
}

func (*OrderCancelReplaceRequest) AllocAccountNullValue() byte {
	return 0
}

func (o *OrderCancelReplaceRequest) AllocAccountCharacterEncoding() string {
	return "US-ASCII"
}

func (*OrderCancelReplaceRequest) StopPxId() uint16 {
	return 99
}

func (*OrderCancelReplaceRequest) StopPxSinceVersion() uint16 {
	return 0
}

func (o *OrderCancelReplaceRequest) StopPxInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.StopPxSinceVersion()
}

func (*OrderCancelReplaceRequest) StopPxDeprecated() uint16 {
	return 0
}

func (*OrderCancelReplaceRequest) StopPxMetaAttribute(meta int) string {
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

func (*OrderCancelReplaceRequest) SecurityDescId() uint16 {
	return 107
}

func (*OrderCancelReplaceRequest) SecurityDescSinceVersion() uint16 {
	return 0
}

func (o *OrderCancelReplaceRequest) SecurityDescInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.SecurityDescSinceVersion()
}

func (*OrderCancelReplaceRequest) SecurityDescDeprecated() uint16 {
	return 0
}

func (*OrderCancelReplaceRequest) SecurityDescMetaAttribute(meta int) string {
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

func (*OrderCancelReplaceRequest) SecurityDescMinValue() byte {
	return byte(32)
}

func (*OrderCancelReplaceRequest) SecurityDescMaxValue() byte {
	return byte(126)
}

func (*OrderCancelReplaceRequest) SecurityDescNullValue() byte {
	return 0
}

func (o *OrderCancelReplaceRequest) SecurityDescCharacterEncoding() string {
	return "US-ASCII"
}

func (*OrderCancelReplaceRequest) MinQtyId() uint16 {
	return 110
}

func (*OrderCancelReplaceRequest) MinQtySinceVersion() uint16 {
	return 0
}

func (o *OrderCancelReplaceRequest) MinQtyInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.MinQtySinceVersion()
}

func (*OrderCancelReplaceRequest) MinQtyDeprecated() uint16 {
	return 0
}

func (*OrderCancelReplaceRequest) MinQtyMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return "Qty"
	case 4:
		return "required"
	}
	return ""
}

func (*OrderCancelReplaceRequest) SecurityTypeId() uint16 {
	return 167
}

func (*OrderCancelReplaceRequest) SecurityTypeSinceVersion() uint16 {
	return 0
}

func (o *OrderCancelReplaceRequest) SecurityTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.SecurityTypeSinceVersion()
}

func (*OrderCancelReplaceRequest) SecurityTypeDeprecated() uint16 {
	return 0
}

func (*OrderCancelReplaceRequest) SecurityTypeMetaAttribute(meta int) string {
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

func (*OrderCancelReplaceRequest) SecurityTypeMinValue() byte {
	return byte(32)
}

func (*OrderCancelReplaceRequest) SecurityTypeMaxValue() byte {
	return byte(126)
}

func (*OrderCancelReplaceRequest) SecurityTypeNullValue() byte {
	return 0
}

func (o *OrderCancelReplaceRequest) SecurityTypeCharacterEncoding() string {
	return "US-ASCII"
}

func (*OrderCancelReplaceRequest) CustomerOrFirmId() uint16 {
	return 204
}

func (*OrderCancelReplaceRequest) CustomerOrFirmSinceVersion() uint16 {
	return 0
}

func (o *OrderCancelReplaceRequest) CustomerOrFirmInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.CustomerOrFirmSinceVersion()
}

func (*OrderCancelReplaceRequest) CustomerOrFirmDeprecated() uint16 {
	return 0
}

func (*OrderCancelReplaceRequest) CustomerOrFirmMetaAttribute(meta int) string {
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

func (*OrderCancelReplaceRequest) MaxShowId() uint16 {
	return 210
}

func (*OrderCancelReplaceRequest) MaxShowSinceVersion() uint16 {
	return 0
}

func (o *OrderCancelReplaceRequest) MaxShowInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.MaxShowSinceVersion()
}

func (*OrderCancelReplaceRequest) MaxShowDeprecated() uint16 {
	return 0
}

func (*OrderCancelReplaceRequest) MaxShowMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return "Qty"
	case 4:
		return "required"
	}
	return ""
}

func (*OrderCancelReplaceRequest) ExpireDateId() uint16 {
	return 432
}

func (*OrderCancelReplaceRequest) ExpireDateSinceVersion() uint16 {
	return 0
}

func (o *OrderCancelReplaceRequest) ExpireDateInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.ExpireDateSinceVersion()
}

func (*OrderCancelReplaceRequest) ExpireDateDeprecated() uint16 {
	return 0
}

func (*OrderCancelReplaceRequest) ExpireDateMetaAttribute(meta int) string {
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

func (*OrderCancelReplaceRequest) ExpireDateMinValue() uint16 {
	return 0
}

func (*OrderCancelReplaceRequest) ExpireDateMaxValue() uint16 {
	return math.MaxUint16 - 1
}

func (*OrderCancelReplaceRequest) ExpireDateNullValue() uint16 {
	return math.MaxUint16
}

func (*OrderCancelReplaceRequest) SelfMatchPreventionIDId() uint16 {
	return 7928
}

func (*OrderCancelReplaceRequest) SelfMatchPreventionIDSinceVersion() uint16 {
	return 0
}

func (o *OrderCancelReplaceRequest) SelfMatchPreventionIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.SelfMatchPreventionIDSinceVersion()
}

func (*OrderCancelReplaceRequest) SelfMatchPreventionIDDeprecated() uint16 {
	return 0
}

func (*OrderCancelReplaceRequest) SelfMatchPreventionIDMetaAttribute(meta int) string {
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

func (*OrderCancelReplaceRequest) SelfMatchPreventionIDMinValue() byte {
	return byte(32)
}

func (*OrderCancelReplaceRequest) SelfMatchPreventionIDMaxValue() byte {
	return byte(126)
}

func (*OrderCancelReplaceRequest) SelfMatchPreventionIDNullValue() byte {
	return 0
}

func (o *OrderCancelReplaceRequest) SelfMatchPreventionIDCharacterEncoding() string {
	return "US-ASCII"
}

func (*OrderCancelReplaceRequest) CtiCodeId() uint16 {
	return 9702
}

func (*OrderCancelReplaceRequest) CtiCodeSinceVersion() uint16 {
	return 0
}

func (o *OrderCancelReplaceRequest) CtiCodeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.CtiCodeSinceVersion()
}

func (*OrderCancelReplaceRequest) CtiCodeDeprecated() uint16 {
	return 0
}

func (*OrderCancelReplaceRequest) CtiCodeMetaAttribute(meta int) string {
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

func (*OrderCancelReplaceRequest) GiveUpFirmId() uint16 {
	return 9707
}

func (*OrderCancelReplaceRequest) GiveUpFirmSinceVersion() uint16 {
	return 0
}

func (o *OrderCancelReplaceRequest) GiveUpFirmInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.GiveUpFirmSinceVersion()
}

func (*OrderCancelReplaceRequest) GiveUpFirmDeprecated() uint16 {
	return 0
}

func (*OrderCancelReplaceRequest) GiveUpFirmMetaAttribute(meta int) string {
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

func (*OrderCancelReplaceRequest) GiveUpFirmMinValue() byte {
	return byte(32)
}

func (*OrderCancelReplaceRequest) GiveUpFirmMaxValue() byte {
	return byte(126)
}

func (*OrderCancelReplaceRequest) GiveUpFirmNullValue() byte {
	return 0
}

func (o *OrderCancelReplaceRequest) GiveUpFirmCharacterEncoding() string {
	return "US-ASCII"
}

func (*OrderCancelReplaceRequest) CmtaGiveupCDId() uint16 {
	return 9708
}

func (*OrderCancelReplaceRequest) CmtaGiveupCDSinceVersion() uint16 {
	return 0
}

func (o *OrderCancelReplaceRequest) CmtaGiveupCDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.CmtaGiveupCDSinceVersion()
}

func (*OrderCancelReplaceRequest) CmtaGiveupCDDeprecated() uint16 {
	return 0
}

func (*OrderCancelReplaceRequest) CmtaGiveupCDMetaAttribute(meta int) string {
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

func (*OrderCancelReplaceRequest) CmtaGiveupCDMinValue() byte {
	return byte(32)
}

func (*OrderCancelReplaceRequest) CmtaGiveupCDMaxValue() byte {
	return byte(126)
}

func (*OrderCancelReplaceRequest) CmtaGiveupCDNullValue() byte {
	return 0
}

func (o *OrderCancelReplaceRequest) CmtaGiveupCDCharacterEncoding() string {
	return "US-ASCII"
}

func (*OrderCancelReplaceRequest) CorrelationClOrdIDId() uint16 {
	return 9717
}

func (*OrderCancelReplaceRequest) CorrelationClOrdIDSinceVersion() uint16 {
	return 0
}

func (o *OrderCancelReplaceRequest) CorrelationClOrdIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.CorrelationClOrdIDSinceVersion()
}

func (*OrderCancelReplaceRequest) CorrelationClOrdIDDeprecated() uint16 {
	return 0
}

func (*OrderCancelReplaceRequest) CorrelationClOrdIDMetaAttribute(meta int) string {
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

func (*OrderCancelReplaceRequest) CorrelationClOrdIDMinValue() byte {
	return byte(32)
}

func (*OrderCancelReplaceRequest) CorrelationClOrdIDMaxValue() byte {
	return byte(126)
}

func (*OrderCancelReplaceRequest) CorrelationClOrdIDNullValue() byte {
	return 0
}

func (o *OrderCancelReplaceRequest) CorrelationClOrdIDCharacterEncoding() string {
	return "US-ASCII"
}

func (*OrderCancelReplaceRequest) OFMOverrideId() uint16 {
	return 9768
}

func (*OrderCancelReplaceRequest) OFMOverrideSinceVersion() uint16 {
	return 0
}

func (o *OrderCancelReplaceRequest) OFMOverrideInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.OFMOverrideSinceVersion()
}

func (*OrderCancelReplaceRequest) OFMOverrideDeprecated() uint16 {
	return 0
}

func (*OrderCancelReplaceRequest) OFMOverrideMetaAttribute(meta int) string {
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
