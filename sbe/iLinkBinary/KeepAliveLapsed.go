// Generated SBE (Simple Binary Encoding) message codec

package ilinkbinary

import (
	"fmt"
	"io"
	"reflect"
)

type KeepAliveLapsedEnum uint8
type KeepAliveLapsedValues struct {
	NotLapsed KeepAliveLapsedEnum
	Lapsed    KeepAliveLapsedEnum
	NullValue KeepAliveLapsedEnum
}

var KeepAliveLapsed = KeepAliveLapsedValues{0, 1, 255}

func (k KeepAliveLapsedEnum) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint8(_w, uint8(k)); err != nil {
		return err
	}
	return nil
}

func (k *KeepAliveLapsedEnum) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if err := _m.ReadUint8(_r, (*uint8)(k)); err != nil {
		return err
	}
	return nil
}

func (k KeepAliveLapsedEnum) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if actingVersion > schemaVersion {
		return nil
	}
	value := reflect.ValueOf(KeepAliveLapsed)
	for idx := 0; idx < value.NumField(); idx++ {
		if k == value.Field(idx).Interface() {
			return nil
		}
	}
	return fmt.Errorf("Range check failed on KeepAliveLapsed, unknown enumeration value %d", k)
}

func (*KeepAliveLapsedEnum) EncodedLength() int64 {
	return 1
}

func (*KeepAliveLapsedEnum) NotLapsedSinceVersion() uint16 {
	return 0
}

func (k *KeepAliveLapsedEnum) NotLapsedInActingVersion(actingVersion uint16) bool {
	return actingVersion >= k.NotLapsedSinceVersion()
}

func (*KeepAliveLapsedEnum) NotLapsedDeprecated() uint16 {
	return 0
}

func (*KeepAliveLapsedEnum) LapsedSinceVersion() uint16 {
	return 0
}

func (k *KeepAliveLapsedEnum) LapsedInActingVersion(actingVersion uint16) bool {
	return actingVersion >= k.LapsedSinceVersion()
}

func (*KeepAliveLapsedEnum) LapsedDeprecated() uint16 {
	return 0
}
