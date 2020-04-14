// Generated SBE (Simple Binary Encoding) message codec

package fix

import (
	"fmt"
	"io"
	"reflect"
)

type OFMOverrideEnum byte
type OFMOverrideValues struct {
	ENABLED   OFMOverrideEnum
	DISABLED  OFMOverrideEnum
	NullValue OFMOverrideEnum
}

var OFMOverride = OFMOverrideValues{89, 78, 0}

func (o OFMOverrideEnum) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint8(_w, byte(o)); err != nil {
		return err
	}
	return nil
}

func (o *OFMOverrideEnum) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if err := _m.ReadUint8(_r, (*byte)(o)); err != nil {
		return err
	}
	return nil
}

func (o OFMOverrideEnum) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if actingVersion > schemaVersion {
		return nil
	}
	value := reflect.ValueOf(OFMOverride)
	for idx := 0; idx < value.NumField(); idx++ {
		if o == value.Field(idx).Interface() {
			return nil
		}
	}
	return fmt.Errorf("Range check failed on OFMOverride, unknown enumeration value %d", o)
}

func (*OFMOverrideEnum) EncodedLength() int64 {
	return 1
}

func (*OFMOverrideEnum) ENABLEDSinceVersion() uint16 {
	return 0
}

func (o *OFMOverrideEnum) ENABLEDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.ENABLEDSinceVersion()
}

func (*OFMOverrideEnum) ENABLEDDeprecated() uint16 {
	return 0
}

func (*OFMOverrideEnum) DISABLEDSinceVersion() uint16 {
	return 0
}

func (o *OFMOverrideEnum) DISABLEDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.DISABLEDSinceVersion()
}

func (*OFMOverrideEnum) DISABLEDDeprecated() uint16 {
	return 0
}
