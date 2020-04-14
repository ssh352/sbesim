// Generated SBE (Simple Binary Encoding) message codec

package fix

import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
)

type NewOrder struct {
	Account               [12]byte
	ClOrdID               [20]byte
	HandInst              HandInstEnum
	CustOrderHandlingInst CustOrderHandlingInstEnum
	OrderQty              IntQty32
	OrdType               OrdTypeEnum
	Price                 OptionalPrice
	Side                  SideEnum
	Symbol                [6]byte
	TimeInForce           TimeInForceEnum
	TransactTime          uint64
	ManualOrderIndicator  BooleanTypeEnum
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
}

func (n *NewOrder) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := n.RangeCheck(n.SbeSchemaVersion(), n.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	if err := _m.WriteBytes(_w, n.Account[:]); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, n.ClOrdID[:]); err != nil {
		return err
	}
	if err := n.HandInst.Encode(_m, _w); err != nil {
		return err
	}
	if err := n.CustOrderHandlingInst.Encode(_m, _w); err != nil {
		return err
	}
	if err := n.OrderQty.Encode(_m, _w); err != nil {
		return err
	}
	if err := n.OrdType.Encode(_m, _w); err != nil {
		return err
	}
	if err := n.Price.Encode(_m, _w); err != nil {
		return err
	}
	if err := n.Side.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, n.Symbol[:]); err != nil {
		return err
	}
	if err := n.TimeInForce.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteUint64(_w, n.TransactTime); err != nil {
		return err
	}
	if err := n.ManualOrderIndicator.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, n.AllocAccount[:]); err != nil {
		return err
	}
	if err := n.StopPx.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, n.SecurityDesc[:]); err != nil {
		return err
	}
	if err := n.MinQty.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, n.SecurityType[:]); err != nil {
		return err
	}
	if err := n.CustomerOrFirm.Encode(_m, _w); err != nil {
		return err
	}
	if err := n.MaxShow.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteUint16(_w, n.ExpireDate); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, n.SelfMatchPreventionID[:]); err != nil {
		return err
	}
	if err := n.CtiCode.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, n.GiveUpFirm[:]); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, n.CmtaGiveupCD[:]); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, n.CorrelationClOrdID[:]); err != nil {
		return err
	}
	return nil
}

func (n *NewOrder) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	if !n.AccountInActingVersion(actingVersion) {
		for idx := 0; idx < 12; idx++ {
			n.Account[idx] = n.AccountNullValue()
		}
	} else {
		if err := _m.ReadBytes(_r, n.Account[:]); err != nil {
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
	if n.HandInstInActingVersion(actingVersion) {
		if err := n.HandInst.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if n.CustOrderHandlingInstInActingVersion(actingVersion) {
		if err := n.CustOrderHandlingInst.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if n.OrderQtyInActingVersion(actingVersion) {
		if err := n.OrderQty.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if n.OrdTypeInActingVersion(actingVersion) {
		if err := n.OrdType.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if n.PriceInActingVersion(actingVersion) {
		if err := n.Price.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if n.SideInActingVersion(actingVersion) {
		if err := n.Side.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if !n.SymbolInActingVersion(actingVersion) {
		for idx := 0; idx < 6; idx++ {
			n.Symbol[idx] = n.SymbolNullValue()
		}
	} else {
		if err := _m.ReadBytes(_r, n.Symbol[:]); err != nil {
			return err
		}
	}
	if n.TimeInForceInActingVersion(actingVersion) {
		if err := n.TimeInForce.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if !n.TransactTimeInActingVersion(actingVersion) {
		n.TransactTime = n.TransactTimeNullValue()
	} else {
		if err := _m.ReadUint64(_r, &n.TransactTime); err != nil {
			return err
		}
	}
	if n.ManualOrderIndicatorInActingVersion(actingVersion) {
		if err := n.ManualOrderIndicator.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if !n.AllocAccountInActingVersion(actingVersion) {
		for idx := 0; idx < 10; idx++ {
			n.AllocAccount[idx] = n.AllocAccountNullValue()
		}
	} else {
		if err := _m.ReadBytes(_r, n.AllocAccount[:]); err != nil {
			return err
		}
	}
	if n.StopPxInActingVersion(actingVersion) {
		if err := n.StopPx.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if !n.SecurityDescInActingVersion(actingVersion) {
		for idx := 0; idx < 20; idx++ {
			n.SecurityDesc[idx] = n.SecurityDescNullValue()
		}
	} else {
		if err := _m.ReadBytes(_r, n.SecurityDesc[:]); err != nil {
			return err
		}
	}
	if n.MinQtyInActingVersion(actingVersion) {
		if err := n.MinQty.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if !n.SecurityTypeInActingVersion(actingVersion) {
		for idx := 0; idx < 3; idx++ {
			n.SecurityType[idx] = n.SecurityTypeNullValue()
		}
	} else {
		if err := _m.ReadBytes(_r, n.SecurityType[:]); err != nil {
			return err
		}
	}
	if n.CustomerOrFirmInActingVersion(actingVersion) {
		if err := n.CustomerOrFirm.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if n.MaxShowInActingVersion(actingVersion) {
		if err := n.MaxShow.Decode(_m, _r, actingVersion); err != nil {
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
	if !n.SelfMatchPreventionIDInActingVersion(actingVersion) {
		for idx := 0; idx < 12; idx++ {
			n.SelfMatchPreventionID[idx] = n.SelfMatchPreventionIDNullValue()
		}
	} else {
		if err := _m.ReadBytes(_r, n.SelfMatchPreventionID[:]); err != nil {
			return err
		}
	}
	if n.CtiCodeInActingVersion(actingVersion) {
		if err := n.CtiCode.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if !n.GiveUpFirmInActingVersion(actingVersion) {
		for idx := 0; idx < 3; idx++ {
			n.GiveUpFirm[idx] = n.GiveUpFirmNullValue()
		}
	} else {
		if err := _m.ReadBytes(_r, n.GiveUpFirm[:]); err != nil {
			return err
		}
	}
	if !n.CmtaGiveupCDInActingVersion(actingVersion) {
		for idx := 0; idx < 2; idx++ {
			n.CmtaGiveupCD[idx] = n.CmtaGiveupCDNullValue()
		}
	} else {
		if err := _m.ReadBytes(_r, n.CmtaGiveupCD[:]); err != nil {
			return err
		}
	}
	if !n.CorrelationClOrdIDInActingVersion(actingVersion) {
		for idx := 0; idx < 20; idx++ {
			n.CorrelationClOrdID[idx] = n.CorrelationClOrdIDNullValue()
		}
	} else {
		if err := _m.ReadBytes(_r, n.CorrelationClOrdID[:]); err != nil {
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

func (n *NewOrder) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if n.AccountInActingVersion(actingVersion) {
		for idx := 0; idx < 12; idx++ {
			if n.Account[idx] < n.AccountMinValue() || n.Account[idx] > n.AccountMaxValue() {
				return fmt.Errorf("Range check failed on n.Account[%d] (%v < %v > %v)", idx, n.AccountMinValue(), n.Account[idx], n.AccountMaxValue())
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
	if err := n.HandInst.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if err := n.CustOrderHandlingInst.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if err := n.OrdType.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if err := n.Side.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if n.SymbolInActingVersion(actingVersion) {
		for idx := 0; idx < 6; idx++ {
			if n.Symbol[idx] < n.SymbolMinValue() || n.Symbol[idx] > n.SymbolMaxValue() {
				return fmt.Errorf("Range check failed on n.Symbol[%d] (%v < %v > %v)", idx, n.SymbolMinValue(), n.Symbol[idx], n.SymbolMaxValue())
			}
		}
	}
	if err := n.TimeInForce.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if n.TransactTimeInActingVersion(actingVersion) {
		if n.TransactTime < n.TransactTimeMinValue() || n.TransactTime > n.TransactTimeMaxValue() {
			return fmt.Errorf("Range check failed on n.TransactTime (%v < %v > %v)", n.TransactTimeMinValue(), n.TransactTime, n.TransactTimeMaxValue())
		}
	}
	if err := n.ManualOrderIndicator.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if n.AllocAccountInActingVersion(actingVersion) {
		for idx := 0; idx < 10; idx++ {
			if n.AllocAccount[idx] < n.AllocAccountMinValue() || n.AllocAccount[idx] > n.AllocAccountMaxValue() {
				return fmt.Errorf("Range check failed on n.AllocAccount[%d] (%v < %v > %v)", idx, n.AllocAccountMinValue(), n.AllocAccount[idx], n.AllocAccountMaxValue())
			}
		}
	}
	if n.SecurityDescInActingVersion(actingVersion) {
		for idx := 0; idx < 20; idx++ {
			if n.SecurityDesc[idx] < n.SecurityDescMinValue() || n.SecurityDesc[idx] > n.SecurityDescMaxValue() {
				return fmt.Errorf("Range check failed on n.SecurityDesc[%d] (%v < %v > %v)", idx, n.SecurityDescMinValue(), n.SecurityDesc[idx], n.SecurityDescMaxValue())
			}
		}
	}
	if n.SecurityTypeInActingVersion(actingVersion) {
		for idx := 0; idx < 3; idx++ {
			if n.SecurityType[idx] < n.SecurityTypeMinValue() || n.SecurityType[idx] > n.SecurityTypeMaxValue() {
				return fmt.Errorf("Range check failed on n.SecurityType[%d] (%v < %v > %v)", idx, n.SecurityTypeMinValue(), n.SecurityType[idx], n.SecurityTypeMaxValue())
			}
		}
	}
	if err := n.CustomerOrFirm.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if n.ExpireDateInActingVersion(actingVersion) {
		if n.ExpireDate < n.ExpireDateMinValue() || n.ExpireDate > n.ExpireDateMaxValue() {
			return fmt.Errorf("Range check failed on n.ExpireDate (%v < %v > %v)", n.ExpireDateMinValue(), n.ExpireDate, n.ExpireDateMaxValue())
		}
	}
	if n.SelfMatchPreventionIDInActingVersion(actingVersion) {
		for idx := 0; idx < 12; idx++ {
			if n.SelfMatchPreventionID[idx] < n.SelfMatchPreventionIDMinValue() || n.SelfMatchPreventionID[idx] > n.SelfMatchPreventionIDMaxValue() {
				return fmt.Errorf("Range check failed on n.SelfMatchPreventionID[%d] (%v < %v > %v)", idx, n.SelfMatchPreventionIDMinValue(), n.SelfMatchPreventionID[idx], n.SelfMatchPreventionIDMaxValue())
			}
		}
	}
	if err := n.CtiCode.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if n.GiveUpFirmInActingVersion(actingVersion) {
		for idx := 0; idx < 3; idx++ {
			if n.GiveUpFirm[idx] < n.GiveUpFirmMinValue() || n.GiveUpFirm[idx] > n.GiveUpFirmMaxValue() {
				return fmt.Errorf("Range check failed on n.GiveUpFirm[%d] (%v < %v > %v)", idx, n.GiveUpFirmMinValue(), n.GiveUpFirm[idx], n.GiveUpFirmMaxValue())
			}
		}
	}
	if n.CmtaGiveupCDInActingVersion(actingVersion) {
		for idx := 0; idx < 2; idx++ {
			if n.CmtaGiveupCD[idx] < n.CmtaGiveupCDMinValue() || n.CmtaGiveupCD[idx] > n.CmtaGiveupCDMaxValue() {
				return fmt.Errorf("Range check failed on n.CmtaGiveupCD[%d] (%v < %v > %v)", idx, n.CmtaGiveupCDMinValue(), n.CmtaGiveupCD[idx], n.CmtaGiveupCDMaxValue())
			}
		}
	}
	if n.CorrelationClOrdIDInActingVersion(actingVersion) {
		for idx := 0; idx < 20; idx++ {
			if n.CorrelationClOrdID[idx] < n.CorrelationClOrdIDMinValue() || n.CorrelationClOrdID[idx] > n.CorrelationClOrdIDMaxValue() {
				return fmt.Errorf("Range check failed on n.CorrelationClOrdID[%d] (%v < %v > %v)", idx, n.CorrelationClOrdIDMinValue(), n.CorrelationClOrdID[idx], n.CorrelationClOrdIDMaxValue())
			}
		}
	}
	return nil
}

func NewOrderInit(n *NewOrder) {
	return
}

func (*NewOrder) SbeBlockLength() (blockLength uint16) {
	return 156
}

func (*NewOrder) SbeTemplateId() (templateId uint16) {
	return 68
}

func (*NewOrder) SbeSchemaId() (schemaId uint16) {
	return 1
}

func (*NewOrder) SbeSchemaVersion() (schemaVersion uint16) {
	return 1
}

func (*NewOrder) SbeSemanticType() (semanticType []byte) {
	return []byte("D")
}

func (*NewOrder) AccountId() uint16 {
	return 1
}

func (*NewOrder) AccountSinceVersion() uint16 {
	return 0
}

func (n *NewOrder) AccountInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.AccountSinceVersion()
}

func (*NewOrder) AccountDeprecated() uint16 {
	return 0
}

func (*NewOrder) AccountMetaAttribute(meta int) string {
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

func (*NewOrder) AccountMinValue() byte {
	return byte(32)
}

func (*NewOrder) AccountMaxValue() byte {
	return byte(126)
}

func (*NewOrder) AccountNullValue() byte {
	return 0
}

func (n *NewOrder) AccountCharacterEncoding() string {
	return "US-ASCII"
}

func (*NewOrder) ClOrdIDId() uint16 {
	return 11
}

func (*NewOrder) ClOrdIDSinceVersion() uint16 {
	return 0
}

func (n *NewOrder) ClOrdIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.ClOrdIDSinceVersion()
}

func (*NewOrder) ClOrdIDDeprecated() uint16 {
	return 0
}

func (*NewOrder) ClOrdIDMetaAttribute(meta int) string {
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

func (*NewOrder) ClOrdIDMinValue() byte {
	return byte(32)
}

func (*NewOrder) ClOrdIDMaxValue() byte {
	return byte(126)
}

func (*NewOrder) ClOrdIDNullValue() byte {
	return 0
}

func (n *NewOrder) ClOrdIDCharacterEncoding() string {
	return "US-ASCII"
}

func (*NewOrder) HandInstId() uint16 {
	return 21
}

func (*NewOrder) HandInstSinceVersion() uint16 {
	return 0
}

func (n *NewOrder) HandInstInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.HandInstSinceVersion()
}

func (*NewOrder) HandInstDeprecated() uint16 {
	return 0
}

func (*NewOrder) HandInstMetaAttribute(meta int) string {
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

func (*NewOrder) CustOrderHandlingInstId() uint16 {
	return 1031
}

func (*NewOrder) CustOrderHandlingInstSinceVersion() uint16 {
	return 0
}

func (n *NewOrder) CustOrderHandlingInstInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.CustOrderHandlingInstSinceVersion()
}

func (*NewOrder) CustOrderHandlingInstDeprecated() uint16 {
	return 0
}

func (*NewOrder) CustOrderHandlingInstMetaAttribute(meta int) string {
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

func (*NewOrder) OrderQtyId() uint16 {
	return 38
}

func (*NewOrder) OrderQtySinceVersion() uint16 {
	return 0
}

func (n *NewOrder) OrderQtyInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.OrderQtySinceVersion()
}

func (*NewOrder) OrderQtyDeprecated() uint16 {
	return 0
}

func (*NewOrder) OrderQtyMetaAttribute(meta int) string {
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

func (*NewOrder) OrdTypeId() uint16 {
	return 40
}

func (*NewOrder) OrdTypeSinceVersion() uint16 {
	return 0
}

func (n *NewOrder) OrdTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.OrdTypeSinceVersion()
}

func (*NewOrder) OrdTypeDeprecated() uint16 {
	return 0
}

func (*NewOrder) OrdTypeMetaAttribute(meta int) string {
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

func (*NewOrder) PriceId() uint16 {
	return 44
}

func (*NewOrder) PriceSinceVersion() uint16 {
	return 0
}

func (n *NewOrder) PriceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.PriceSinceVersion()
}

func (*NewOrder) PriceDeprecated() uint16 {
	return 0
}

func (*NewOrder) PriceMetaAttribute(meta int) string {
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

func (*NewOrder) SideId() uint16 {
	return 54
}

func (*NewOrder) SideSinceVersion() uint16 {
	return 0
}

func (n *NewOrder) SideInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.SideSinceVersion()
}

func (*NewOrder) SideDeprecated() uint16 {
	return 0
}

func (*NewOrder) SideMetaAttribute(meta int) string {
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

func (*NewOrder) SymbolId() uint16 {
	return 55
}

func (*NewOrder) SymbolSinceVersion() uint16 {
	return 0
}

func (n *NewOrder) SymbolInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.SymbolSinceVersion()
}

func (*NewOrder) SymbolDeprecated() uint16 {
	return 0
}

func (*NewOrder) SymbolMetaAttribute(meta int) string {
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

func (*NewOrder) SymbolMinValue() byte {
	return byte(32)
}

func (*NewOrder) SymbolMaxValue() byte {
	return byte(126)
}

func (*NewOrder) SymbolNullValue() byte {
	return 0
}

func (n *NewOrder) SymbolCharacterEncoding() string {
	return "US-ASCII"
}

func (*NewOrder) TimeInForceId() uint16 {
	return 59
}

func (*NewOrder) TimeInForceSinceVersion() uint16 {
	return 0
}

func (n *NewOrder) TimeInForceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.TimeInForceSinceVersion()
}

func (*NewOrder) TimeInForceDeprecated() uint16 {
	return 0
}

func (*NewOrder) TimeInForceMetaAttribute(meta int) string {
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

func (*NewOrder) TransactTimeId() uint16 {
	return 60
}

func (*NewOrder) TransactTimeSinceVersion() uint16 {
	return 0
}

func (n *NewOrder) TransactTimeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.TransactTimeSinceVersion()
}

func (*NewOrder) TransactTimeDeprecated() uint16 {
	return 0
}

func (*NewOrder) TransactTimeMetaAttribute(meta int) string {
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

func (*NewOrder) TransactTimeMinValue() uint64 {
	return 0
}

func (*NewOrder) TransactTimeMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*NewOrder) TransactTimeNullValue() uint64 {
	return math.MaxUint64
}

func (*NewOrder) ManualOrderIndicatorId() uint16 {
	return 1028
}

func (*NewOrder) ManualOrderIndicatorSinceVersion() uint16 {
	return 0
}

func (n *NewOrder) ManualOrderIndicatorInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.ManualOrderIndicatorSinceVersion()
}

func (*NewOrder) ManualOrderIndicatorDeprecated() uint16 {
	return 0
}

func (*NewOrder) ManualOrderIndicatorMetaAttribute(meta int) string {
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

func (*NewOrder) AllocAccountId() uint16 {
	return 79
}

func (*NewOrder) AllocAccountSinceVersion() uint16 {
	return 0
}

func (n *NewOrder) AllocAccountInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.AllocAccountSinceVersion()
}

func (*NewOrder) AllocAccountDeprecated() uint16 {
	return 0
}

func (*NewOrder) AllocAccountMetaAttribute(meta int) string {
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

func (*NewOrder) AllocAccountMinValue() byte {
	return byte(32)
}

func (*NewOrder) AllocAccountMaxValue() byte {
	return byte(126)
}

func (*NewOrder) AllocAccountNullValue() byte {
	return 0
}

func (n *NewOrder) AllocAccountCharacterEncoding() string {
	return "US-ASCII"
}

func (*NewOrder) StopPxId() uint16 {
	return 99
}

func (*NewOrder) StopPxSinceVersion() uint16 {
	return 0
}

func (n *NewOrder) StopPxInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.StopPxSinceVersion()
}

func (*NewOrder) StopPxDeprecated() uint16 {
	return 0
}

func (*NewOrder) StopPxMetaAttribute(meta int) string {
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

func (*NewOrder) SecurityDescId() uint16 {
	return 107
}

func (*NewOrder) SecurityDescSinceVersion() uint16 {
	return 0
}

func (n *NewOrder) SecurityDescInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.SecurityDescSinceVersion()
}

func (*NewOrder) SecurityDescDeprecated() uint16 {
	return 0
}

func (*NewOrder) SecurityDescMetaAttribute(meta int) string {
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

func (*NewOrder) SecurityDescMinValue() byte {
	return byte(32)
}

func (*NewOrder) SecurityDescMaxValue() byte {
	return byte(126)
}

func (*NewOrder) SecurityDescNullValue() byte {
	return 0
}

func (n *NewOrder) SecurityDescCharacterEncoding() string {
	return "US-ASCII"
}

func (*NewOrder) MinQtyId() uint16 {
	return 110
}

func (*NewOrder) MinQtySinceVersion() uint16 {
	return 0
}

func (n *NewOrder) MinQtyInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.MinQtySinceVersion()
}

func (*NewOrder) MinQtyDeprecated() uint16 {
	return 0
}

func (*NewOrder) MinQtyMetaAttribute(meta int) string {
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

func (*NewOrder) SecurityTypeId() uint16 {
	return 167
}

func (*NewOrder) SecurityTypeSinceVersion() uint16 {
	return 0
}

func (n *NewOrder) SecurityTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.SecurityTypeSinceVersion()
}

func (*NewOrder) SecurityTypeDeprecated() uint16 {
	return 0
}

func (*NewOrder) SecurityTypeMetaAttribute(meta int) string {
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

func (*NewOrder) SecurityTypeMinValue() byte {
	return byte(32)
}

func (*NewOrder) SecurityTypeMaxValue() byte {
	return byte(126)
}

func (*NewOrder) SecurityTypeNullValue() byte {
	return 0
}

func (n *NewOrder) SecurityTypeCharacterEncoding() string {
	return "US-ASCII"
}

func (*NewOrder) CustomerOrFirmId() uint16 {
	return 204
}

func (*NewOrder) CustomerOrFirmSinceVersion() uint16 {
	return 0
}

func (n *NewOrder) CustomerOrFirmInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.CustomerOrFirmSinceVersion()
}

func (*NewOrder) CustomerOrFirmDeprecated() uint16 {
	return 0
}

func (*NewOrder) CustomerOrFirmMetaAttribute(meta int) string {
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

func (*NewOrder) MaxShowId() uint16 {
	return 210
}

func (*NewOrder) MaxShowSinceVersion() uint16 {
	return 0
}

func (n *NewOrder) MaxShowInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.MaxShowSinceVersion()
}

func (*NewOrder) MaxShowDeprecated() uint16 {
	return 0
}

func (*NewOrder) MaxShowMetaAttribute(meta int) string {
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

func (*NewOrder) ExpireDateId() uint16 {
	return 432
}

func (*NewOrder) ExpireDateSinceVersion() uint16 {
	return 0
}

func (n *NewOrder) ExpireDateInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.ExpireDateSinceVersion()
}

func (*NewOrder) ExpireDateDeprecated() uint16 {
	return 0
}

func (*NewOrder) ExpireDateMetaAttribute(meta int) string {
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

func (*NewOrder) ExpireDateMinValue() uint16 {
	return 0
}

func (*NewOrder) ExpireDateMaxValue() uint16 {
	return math.MaxUint16 - 1
}

func (*NewOrder) ExpireDateNullValue() uint16 {
	return math.MaxUint16
}

func (*NewOrder) SelfMatchPreventionIDId() uint16 {
	return 7928
}

func (*NewOrder) SelfMatchPreventionIDSinceVersion() uint16 {
	return 0
}

func (n *NewOrder) SelfMatchPreventionIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.SelfMatchPreventionIDSinceVersion()
}

func (*NewOrder) SelfMatchPreventionIDDeprecated() uint16 {
	return 0
}

func (*NewOrder) SelfMatchPreventionIDMetaAttribute(meta int) string {
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

func (*NewOrder) SelfMatchPreventionIDMinValue() byte {
	return byte(32)
}

func (*NewOrder) SelfMatchPreventionIDMaxValue() byte {
	return byte(126)
}

func (*NewOrder) SelfMatchPreventionIDNullValue() byte {
	return 0
}

func (n *NewOrder) SelfMatchPreventionIDCharacterEncoding() string {
	return "US-ASCII"
}

func (*NewOrder) CtiCodeId() uint16 {
	return 9702
}

func (*NewOrder) CtiCodeSinceVersion() uint16 {
	return 0
}

func (n *NewOrder) CtiCodeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.CtiCodeSinceVersion()
}

func (*NewOrder) CtiCodeDeprecated() uint16 {
	return 0
}

func (*NewOrder) CtiCodeMetaAttribute(meta int) string {
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

func (*NewOrder) GiveUpFirmId() uint16 {
	return 9707
}

func (*NewOrder) GiveUpFirmSinceVersion() uint16 {
	return 0
}

func (n *NewOrder) GiveUpFirmInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.GiveUpFirmSinceVersion()
}

func (*NewOrder) GiveUpFirmDeprecated() uint16 {
	return 0
}

func (*NewOrder) GiveUpFirmMetaAttribute(meta int) string {
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

func (*NewOrder) GiveUpFirmMinValue() byte {
	return byte(32)
}

func (*NewOrder) GiveUpFirmMaxValue() byte {
	return byte(126)
}

func (*NewOrder) GiveUpFirmNullValue() byte {
	return 0
}

func (n *NewOrder) GiveUpFirmCharacterEncoding() string {
	return "US-ASCII"
}

func (*NewOrder) CmtaGiveupCDId() uint16 {
	return 9708
}

func (*NewOrder) CmtaGiveupCDSinceVersion() uint16 {
	return 0
}

func (n *NewOrder) CmtaGiveupCDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.CmtaGiveupCDSinceVersion()
}

func (*NewOrder) CmtaGiveupCDDeprecated() uint16 {
	return 0
}

func (*NewOrder) CmtaGiveupCDMetaAttribute(meta int) string {
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

func (*NewOrder) CmtaGiveupCDMinValue() byte {
	return byte(32)
}

func (*NewOrder) CmtaGiveupCDMaxValue() byte {
	return byte(126)
}

func (*NewOrder) CmtaGiveupCDNullValue() byte {
	return 0
}

func (n *NewOrder) CmtaGiveupCDCharacterEncoding() string {
	return "US-ASCII"
}

func (*NewOrder) CorrelationClOrdIDId() uint16 {
	return 9717
}

func (*NewOrder) CorrelationClOrdIDSinceVersion() uint16 {
	return 0
}

func (n *NewOrder) CorrelationClOrdIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.CorrelationClOrdIDSinceVersion()
}

func (*NewOrder) CorrelationClOrdIDDeprecated() uint16 {
	return 0
}

func (*NewOrder) CorrelationClOrdIDMetaAttribute(meta int) string {
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

func (*NewOrder) CorrelationClOrdIDMinValue() byte {
	return byte(32)
}

func (*NewOrder) CorrelationClOrdIDMaxValue() byte {
	return byte(126)
}

func (*NewOrder) CorrelationClOrdIDNullValue() byte {
	return 0
}

func (n *NewOrder) CorrelationClOrdIDCharacterEncoding() string {
	return "US-ASCII"
}
