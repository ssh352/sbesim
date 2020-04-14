// Generated SBE (Simple Binary Encoding) message codec

package ilinkbinary

import (
	"fmt"
	"io"
	"reflect"
)

type OrderTypeReqEnum byte
type OrderTypeReqValues struct {
	MarketwithProtection      OrderTypeReqEnum
	Limit                     OrderTypeReqEnum
	StopwithProtection        OrderTypeReqEnum
	StopLimit                 OrderTypeReqEnum
	MarketWithLeftoverAsLimit OrderTypeReqEnum
	NullValue                 OrderTypeReqEnum
}

var OrderTypeReq = OrderTypeReqValues{49, 50, 51, 52, 75, 0}

func (o OrderTypeReqEnum) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint8(_w, byte(o)); err != nil {
		return err
	}
	return nil
}

func (o *OrderTypeReqEnum) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if err := _m.ReadUint8(_r, (*byte)(o)); err != nil {
		return err
	}
	return nil
}

func (o OrderTypeReqEnum) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if actingVersion > schemaVersion {
		return nil
	}
	value := reflect.ValueOf(OrderTypeReq)
	for idx := 0; idx < value.NumField(); idx++ {
		if o == value.Field(idx).Interface() {
			return nil
		}
	}
	return fmt.Errorf("Range check failed on OrderTypeReq, unknown enumeration value %d", o)
}

func (*OrderTypeReqEnum) EncodedLength() int64 {
	return 1
}

func (*OrderTypeReqEnum) MarketwithProtectionSinceVersion() uint16 {
	return 0
}

func (o *OrderTypeReqEnum) MarketwithProtectionInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.MarketwithProtectionSinceVersion()
}

func (*OrderTypeReqEnum) MarketwithProtectionDeprecated() uint16 {
	return 0
}

func (*OrderTypeReqEnum) LimitSinceVersion() uint16 {
	return 0
}

func (o *OrderTypeReqEnum) LimitInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.LimitSinceVersion()
}

func (*OrderTypeReqEnum) LimitDeprecated() uint16 {
	return 0
}

func (*OrderTypeReqEnum) StopwithProtectionSinceVersion() uint16 {
	return 0
}

func (o *OrderTypeReqEnum) StopwithProtectionInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.StopwithProtectionSinceVersion()
}

func (*OrderTypeReqEnum) StopwithProtectionDeprecated() uint16 {
	return 0
}

func (*OrderTypeReqEnum) StopLimitSinceVersion() uint16 {
	return 0
}

func (o *OrderTypeReqEnum) StopLimitInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.StopLimitSinceVersion()
}

func (*OrderTypeReqEnum) StopLimitDeprecated() uint16 {
	return 0
}

func (*OrderTypeReqEnum) MarketWithLeftoverAsLimitSinceVersion() uint16 {
	return 0
}

func (o *OrderTypeReqEnum) MarketWithLeftoverAsLimitInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.MarketWithLeftoverAsLimitSinceVersion()
}

func (*OrderTypeReqEnum) MarketWithLeftoverAsLimitDeprecated() uint16 {
	return 0
}
