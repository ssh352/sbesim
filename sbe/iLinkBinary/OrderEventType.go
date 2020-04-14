// Generated SBE (Simple Binary Encoding) message codec

package ilinkbinary

import (
	"fmt"
	"io"
	"reflect"
)

type OrderEventTypeEnum uint8
type OrderEventTypeValues struct {
	PartiallyFilled OrderEventTypeEnum
	Filled          OrderEventTypeEnum
	NullValue       OrderEventTypeEnum
}

var OrderEventType = OrderEventTypeValues{4, 5, 255}

func (o OrderEventTypeEnum) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint8(_w, uint8(o)); err != nil {
		return err
	}
	return nil
}

func (o *OrderEventTypeEnum) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if err := _m.ReadUint8(_r, (*uint8)(o)); err != nil {
		return err
	}
	return nil
}

func (o OrderEventTypeEnum) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if actingVersion > schemaVersion {
		return nil
	}
	value := reflect.ValueOf(OrderEventType)
	for idx := 0; idx < value.NumField(); idx++ {
		if o == value.Field(idx).Interface() {
			return nil
		}
	}
	return fmt.Errorf("Range check failed on OrderEventType, unknown enumeration value %d", o)
}

func (*OrderEventTypeEnum) EncodedLength() int64 {
	return 1
}

func (*OrderEventTypeEnum) PartiallyFilledSinceVersion() uint16 {
	return 0
}

func (o *OrderEventTypeEnum) PartiallyFilledInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.PartiallyFilledSinceVersion()
}

func (*OrderEventTypeEnum) PartiallyFilledDeprecated() uint16 {
	return 0
}

func (*OrderEventTypeEnum) FilledSinceVersion() uint16 {
	return 0
}

func (o *OrderEventTypeEnum) FilledInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.FilledSinceVersion()
}

func (*OrderEventTypeEnum) FilledDeprecated() uint16 {
	return 0
}
