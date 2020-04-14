// Generated SBE (Simple Binary Encoding) message codec

package ilinkbinary

import (
	"fmt"
	"io"
	"reflect"
)

type SMPIEnum byte
type SMPIValues struct {
	CancelNewest SMPIEnum
	CancelOldest SMPIEnum
	NullValue    SMPIEnum
}

var SMPI = SMPIValues{78, 79, 0}

func (s SMPIEnum) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint8(_w, byte(s)); err != nil {
		return err
	}
	return nil
}

func (s *SMPIEnum) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if err := _m.ReadUint8(_r, (*byte)(s)); err != nil {
		return err
	}
	return nil
}

func (s SMPIEnum) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if actingVersion > schemaVersion {
		return nil
	}
	value := reflect.ValueOf(SMPI)
	for idx := 0; idx < value.NumField(); idx++ {
		if s == value.Field(idx).Interface() {
			return nil
		}
	}
	return fmt.Errorf("Range check failed on SMPI, unknown enumeration value %d", s)
}

func (*SMPIEnum) EncodedLength() int64 {
	return 1
}

func (*SMPIEnum) CancelNewestSinceVersion() uint16 {
	return 0
}

func (s *SMPIEnum) CancelNewestInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.CancelNewestSinceVersion()
}

func (*SMPIEnum) CancelNewestDeprecated() uint16 {
	return 0
}

func (*SMPIEnum) CancelOldestSinceVersion() uint16 {
	return 0
}

func (s *SMPIEnum) CancelOldestInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.CancelOldestSinceVersion()
}

func (*SMPIEnum) CancelOldestDeprecated() uint16 {
	return 0
}
