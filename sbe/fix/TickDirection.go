// Generated SBE (Simple Binary Encoding) message codec

package fix

import (
	"fmt"
	"io"
	"reflect"
)

type TickDirectionEnum uint8
type TickDirectionValues struct {
	PLUS_TICK  TickDirectionEnum
	MINUS_TICK TickDirectionEnum
	NullValue  TickDirectionEnum
}

var TickDirection = TickDirectionValues{0, 1, 255}

func (t TickDirectionEnum) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint8(_w, uint8(t)); err != nil {
		return err
	}
	return nil
}

func (t *TickDirectionEnum) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if err := _m.ReadUint8(_r, (*uint8)(t)); err != nil {
		return err
	}
	return nil
}

func (t TickDirectionEnum) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if actingVersion > schemaVersion {
		return nil
	}
	value := reflect.ValueOf(TickDirection)
	for idx := 0; idx < value.NumField(); idx++ {
		if t == value.Field(idx).Interface() {
			return nil
		}
	}
	return fmt.Errorf("Range check failed on TickDirection, unknown enumeration value %d", t)
}

func (*TickDirectionEnum) EncodedLength() int64 {
	return 1
}

func (*TickDirectionEnum) PLUS_TICKSinceVersion() uint16 {
	return 0
}

func (t *TickDirectionEnum) PLUS_TICKInActingVersion(actingVersion uint16) bool {
	return actingVersion >= t.PLUS_TICKSinceVersion()
}

func (*TickDirectionEnum) PLUS_TICKDeprecated() uint16 {
	return 0
}

func (*TickDirectionEnum) MINUS_TICKSinceVersion() uint16 {
	return 0
}

func (t *TickDirectionEnum) MINUS_TICKInActingVersion(actingVersion uint16) bool {
	return actingVersion >= t.MINUS_TICKSinceVersion()
}

func (*TickDirectionEnum) MINUS_TICKDeprecated() uint16 {
	return 0
}
