// Generated SBE (Simple Binary Encoding) message codec

package fix

import (
	"fmt"
	"io"
	"reflect"
)

type SecurityIDSourceEnum byte
type SecurityIDSourceValues struct {
	EXCHANGE_SYMBOL SecurityIDSourceEnum
	NullValue       SecurityIDSourceEnum
}

var SecurityIDSource = SecurityIDSourceValues{56, 0}

func (s SecurityIDSourceEnum) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint8(_w, byte(s)); err != nil {
		return err
	}
	return nil
}

func (s *SecurityIDSourceEnum) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if err := _m.ReadUint8(_r, (*byte)(s)); err != nil {
		return err
	}
	return nil
}

func (s SecurityIDSourceEnum) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if actingVersion > schemaVersion {
		return nil
	}
	value := reflect.ValueOf(SecurityIDSource)
	for idx := 0; idx < value.NumField(); idx++ {
		if s == value.Field(idx).Interface() {
			return nil
		}
	}
	return fmt.Errorf("Range check failed on SecurityIDSource, unknown enumeration value %d", s)
}

func (*SecurityIDSourceEnum) EncodedLength() int64 {
	return 1
}

func (*SecurityIDSourceEnum) EXCHANGE_SYMBOLSinceVersion() uint16 {
	return 0
}

func (s *SecurityIDSourceEnum) EXCHANGE_SYMBOLInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.EXCHANGE_SYMBOLSinceVersion()
}

func (*SecurityIDSourceEnum) EXCHANGE_SYMBOLDeprecated() uint16 {
	return 0
}
