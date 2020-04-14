// Generated SBE (Simple Binary Encoding) message codec

package ilinkbinary

import (
	"fmt"
	"io"
	"reflect"
)

type MassStatusTIFEnum uint8
type MassStatusTIFValues struct {
	Day       MassStatusTIFEnum
	GTC       MassStatusTIFEnum
	GTD       MassStatusTIFEnum
	NullValue MassStatusTIFEnum
}

var MassStatusTIF = MassStatusTIFValues{0, 1, 6, 255}

func (m MassStatusTIFEnum) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint8(_w, uint8(m)); err != nil {
		return err
	}
	return nil
}

func (m *MassStatusTIFEnum) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if err := _m.ReadUint8(_r, (*uint8)(m)); err != nil {
		return err
	}
	return nil
}

func (m MassStatusTIFEnum) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if actingVersion > schemaVersion {
		return nil
	}
	value := reflect.ValueOf(MassStatusTIF)
	for idx := 0; idx < value.NumField(); idx++ {
		if m == value.Field(idx).Interface() {
			return nil
		}
	}
	return fmt.Errorf("Range check failed on MassStatusTIF, unknown enumeration value %d", m)
}

func (*MassStatusTIFEnum) EncodedLength() int64 {
	return 1
}

func (*MassStatusTIFEnum) DaySinceVersion() uint16 {
	return 0
}

func (m *MassStatusTIFEnum) DayInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.DaySinceVersion()
}

func (*MassStatusTIFEnum) DayDeprecated() uint16 {
	return 0
}

func (*MassStatusTIFEnum) GTCSinceVersion() uint16 {
	return 0
}

func (m *MassStatusTIFEnum) GTCInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.GTCSinceVersion()
}

func (*MassStatusTIFEnum) GTCDeprecated() uint16 {
	return 0
}

func (*MassStatusTIFEnum) GTDSinceVersion() uint16 {
	return 0
}

func (m *MassStatusTIFEnum) GTDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.GTDSinceVersion()
}

func (*MassStatusTIFEnum) GTDDeprecated() uint16 {
	return 0
}
