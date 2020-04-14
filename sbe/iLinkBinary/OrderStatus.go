// Generated SBE (Simple Binary Encoding) message codec

package ilinkbinary

import (
	"fmt"
	"io"
	"reflect"
)

type OrderStatusEnum byte
type OrderStatusValues struct {
	New             OrderStatusEnum
	PartiallyFilled OrderStatusEnum
	Filled          OrderStatusEnum
	Cancelled       OrderStatusEnum
	Replaced        OrderStatusEnum
	Rejected        OrderStatusEnum
	Expired         OrderStatusEnum
	Undefined       OrderStatusEnum
	NullValue       OrderStatusEnum
}

var OrderStatus = OrderStatusValues{48, 49, 50, 52, 53, 56, 67, 85, 0}

func (o OrderStatusEnum) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint8(_w, byte(o)); err != nil {
		return err
	}
	return nil
}

func (o *OrderStatusEnum) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if err := _m.ReadUint8(_r, (*byte)(o)); err != nil {
		return err
	}
	return nil
}

func (o OrderStatusEnum) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if actingVersion > schemaVersion {
		return nil
	}
	value := reflect.ValueOf(OrderStatus)
	for idx := 0; idx < value.NumField(); idx++ {
		if o == value.Field(idx).Interface() {
			return nil
		}
	}
	return fmt.Errorf("Range check failed on OrderStatus, unknown enumeration value %d", o)
}

func (*OrderStatusEnum) EncodedLength() int64 {
	return 1
}

func (*OrderStatusEnum) NewSinceVersion() uint16 {
	return 0
}

func (o *OrderStatusEnum) NewInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.NewSinceVersion()
}

func (*OrderStatusEnum) NewDeprecated() uint16 {
	return 0
}

func (*OrderStatusEnum) PartiallyFilledSinceVersion() uint16 {
	return 0
}

func (o *OrderStatusEnum) PartiallyFilledInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.PartiallyFilledSinceVersion()
}

func (*OrderStatusEnum) PartiallyFilledDeprecated() uint16 {
	return 0
}

func (*OrderStatusEnum) FilledSinceVersion() uint16 {
	return 0
}

func (o *OrderStatusEnum) FilledInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.FilledSinceVersion()
}

func (*OrderStatusEnum) FilledDeprecated() uint16 {
	return 0
}

func (*OrderStatusEnum) CancelledSinceVersion() uint16 {
	return 0
}

func (o *OrderStatusEnum) CancelledInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.CancelledSinceVersion()
}

func (*OrderStatusEnum) CancelledDeprecated() uint16 {
	return 0
}

func (*OrderStatusEnum) ReplacedSinceVersion() uint16 {
	return 0
}

func (o *OrderStatusEnum) ReplacedInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.ReplacedSinceVersion()
}

func (*OrderStatusEnum) ReplacedDeprecated() uint16 {
	return 0
}

func (*OrderStatusEnum) RejectedSinceVersion() uint16 {
	return 0
}

func (o *OrderStatusEnum) RejectedInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.RejectedSinceVersion()
}

func (*OrderStatusEnum) RejectedDeprecated() uint16 {
	return 0
}

func (*OrderStatusEnum) ExpiredSinceVersion() uint16 {
	return 0
}

func (o *OrderStatusEnum) ExpiredInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.ExpiredSinceVersion()
}

func (*OrderStatusEnum) ExpiredDeprecated() uint16 {
	return 0
}

func (*OrderStatusEnum) UndefinedSinceVersion() uint16 {
	return 0
}

func (o *OrderStatusEnum) UndefinedInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.UndefinedSinceVersion()
}

func (*OrderStatusEnum) UndefinedDeprecated() uint16 {
	return 0
}
