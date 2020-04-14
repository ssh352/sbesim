// Generated SBE (Simple Binary Encoding) message codec

package fix

import (
	"fmt"
	"io"
	"reflect"
)

type NoAllocsEnum byte
type NoAllocsValues struct {
	ONE       NoAllocsEnum
	NullValue NoAllocsEnum
}

var NoAllocs = NoAllocsValues{49, 0}

func (n NoAllocsEnum) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint8(_w, byte(n)); err != nil {
		return err
	}
	return nil
}

func (n *NoAllocsEnum) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if err := _m.ReadUint8(_r, (*byte)(n)); err != nil {
		return err
	}
	return nil
}

func (n NoAllocsEnum) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if actingVersion > schemaVersion {
		return nil
	}
	value := reflect.ValueOf(NoAllocs)
	for idx := 0; idx < value.NumField(); idx++ {
		if n == value.Field(idx).Interface() {
			return nil
		}
	}
	return fmt.Errorf("Range check failed on NoAllocs, unknown enumeration value %d", n)
}

func (*NoAllocsEnum) EncodedLength() int64 {
	return 1
}

func (*NoAllocsEnum) ONESinceVersion() uint16 {
	return 0
}

func (n *NoAllocsEnum) ONEInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.ONESinceVersion()
}

func (*NoAllocsEnum) ONEDeprecated() uint16 {
	return 0
}
