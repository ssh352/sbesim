// Generated SBE (Simple Binary Encoding) message codec

package ilinkbinary

import (
	"fmt"
	"io"
	"reflect"
)

type DKReasonEnum byte
type DKReasonValues struct {
	UnknownSecurity           DKReasonEnum
	WrongSide                 DKReasonEnum
	QuantityExceedsOrder      DKReasonEnum
	NoMatchingOrder           DKReasonEnum
	PriceExceedsLimit         DKReasonEnum
	CalculationDifference     DKReasonEnum
	NoMatchingExecutionReport DKReasonEnum
	Other                     DKReasonEnum
	NullValue                 DKReasonEnum
}

var DKReason = DKReasonValues{65, 66, 67, 68, 69, 70, 71, 90, 0}

func (d DKReasonEnum) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint8(_w, byte(d)); err != nil {
		return err
	}
	return nil
}

func (d *DKReasonEnum) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if err := _m.ReadUint8(_r, (*byte)(d)); err != nil {
		return err
	}
	return nil
}

func (d DKReasonEnum) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if actingVersion > schemaVersion {
		return nil
	}
	value := reflect.ValueOf(DKReason)
	for idx := 0; idx < value.NumField(); idx++ {
		if d == value.Field(idx).Interface() {
			return nil
		}
	}
	return fmt.Errorf("Range check failed on DKReason, unknown enumeration value %d", d)
}

func (*DKReasonEnum) EncodedLength() int64 {
	return 1
}

func (*DKReasonEnum) UnknownSecuritySinceVersion() uint16 {
	return 0
}

func (d *DKReasonEnum) UnknownSecurityInActingVersion(actingVersion uint16) bool {
	return actingVersion >= d.UnknownSecuritySinceVersion()
}

func (*DKReasonEnum) UnknownSecurityDeprecated() uint16 {
	return 0
}

func (*DKReasonEnum) WrongSideSinceVersion() uint16 {
	return 0
}

func (d *DKReasonEnum) WrongSideInActingVersion(actingVersion uint16) bool {
	return actingVersion >= d.WrongSideSinceVersion()
}

func (*DKReasonEnum) WrongSideDeprecated() uint16 {
	return 0
}

func (*DKReasonEnum) QuantityExceedsOrderSinceVersion() uint16 {
	return 0
}

func (d *DKReasonEnum) QuantityExceedsOrderInActingVersion(actingVersion uint16) bool {
	return actingVersion >= d.QuantityExceedsOrderSinceVersion()
}

func (*DKReasonEnum) QuantityExceedsOrderDeprecated() uint16 {
	return 0
}

func (*DKReasonEnum) NoMatchingOrderSinceVersion() uint16 {
	return 0
}

func (d *DKReasonEnum) NoMatchingOrderInActingVersion(actingVersion uint16) bool {
	return actingVersion >= d.NoMatchingOrderSinceVersion()
}

func (*DKReasonEnum) NoMatchingOrderDeprecated() uint16 {
	return 0
}

func (*DKReasonEnum) PriceExceedsLimitSinceVersion() uint16 {
	return 0
}

func (d *DKReasonEnum) PriceExceedsLimitInActingVersion(actingVersion uint16) bool {
	return actingVersion >= d.PriceExceedsLimitSinceVersion()
}

func (*DKReasonEnum) PriceExceedsLimitDeprecated() uint16 {
	return 0
}

func (*DKReasonEnum) CalculationDifferenceSinceVersion() uint16 {
	return 0
}

func (d *DKReasonEnum) CalculationDifferenceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= d.CalculationDifferenceSinceVersion()
}

func (*DKReasonEnum) CalculationDifferenceDeprecated() uint16 {
	return 0
}

func (*DKReasonEnum) NoMatchingExecutionReportSinceVersion() uint16 {
	return 0
}

func (d *DKReasonEnum) NoMatchingExecutionReportInActingVersion(actingVersion uint16) bool {
	return actingVersion >= d.NoMatchingExecutionReportSinceVersion()
}

func (*DKReasonEnum) NoMatchingExecutionReportDeprecated() uint16 {
	return 0
}

func (*DKReasonEnum) OtherSinceVersion() uint16 {
	return 0
}

func (d *DKReasonEnum) OtherInActingVersion(actingVersion uint16) bool {
	return actingVersion >= d.OtherSinceVersion()
}

func (*DKReasonEnum) OtherDeprecated() uint16 {
	return 0
}
