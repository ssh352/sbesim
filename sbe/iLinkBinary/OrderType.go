// Generated SBE (Simple Binary Encoding) message codec

package ilinkbinary

import (
	"fmt"
	"io"
	"reflect"
)

type OrderTypeEnum byte
type OrderTypeValues struct {
	MarketWithProtection      OrderTypeEnum
	Limit                     OrderTypeEnum
	StopLimit                 OrderTypeEnum
	MarketWithLeftoverAsLimit OrderTypeEnum
	NullValue                 OrderTypeEnum
}

var OrderType = OrderTypeValues{49, 50, 52, 75, 0}

func (o OrderTypeEnum) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint8(_w, byte(o)); err != nil {
		return err
	}
	return nil
}

func (o *OrderTypeEnum) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if err := _m.ReadUint8(_r, (*byte)(o)); err != nil {
		return err
	}
	return nil
}

func (o OrderTypeEnum) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if actingVersion > schemaVersion {
		return nil
	}
	value := reflect.ValueOf(OrderType)
	for idx := 0; idx < value.NumField(); idx++ {
		if o == value.Field(idx).Interface() {
			return nil
		}
	}
	return fmt.Errorf("Range check failed on OrderType, unknown enumeration value %d", o)
}

func (*OrderTypeEnum) EncodedLength() int64 {
	return 1
}

func (*OrderTypeEnum) MarketWithProtectionSinceVersion() uint16 {
	return 0
}

func (o *OrderTypeEnum) MarketWithProtectionInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.MarketWithProtectionSinceVersion()
}

func (*OrderTypeEnum) MarketWithProtectionDeprecated() uint16 {
	return 0
}

func (*OrderTypeEnum) LimitSinceVersion() uint16 {
	return 0
}

func (o *OrderTypeEnum) LimitInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.LimitSinceVersion()
}

func (*OrderTypeEnum) LimitDeprecated() uint16 {
	return 0
}

func (*OrderTypeEnum) StopLimitSinceVersion() uint16 {
	return 0
}

func (o *OrderTypeEnum) StopLimitInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.StopLimitSinceVersion()
}

func (*OrderTypeEnum) StopLimitDeprecated() uint16 {
	return 0
}

func (*OrderTypeEnum) MarketWithLeftoverAsLimitSinceVersion() uint16 {
	return 0
}

func (o *OrderTypeEnum) MarketWithLeftoverAsLimitInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.MarketWithLeftoverAsLimitSinceVersion()
}

func (*OrderTypeEnum) MarketWithLeftoverAsLimitDeprecated() uint16 {
	return 0
}
