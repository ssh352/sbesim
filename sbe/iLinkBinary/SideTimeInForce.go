// Generated SBE (Simple Binary Encoding) message codec

package ilinkbinary

import (
	"fmt"
	"io"
	"reflect"
)

type SideTimeInForceEnum uint8
type SideTimeInForceValues struct {
	Day       SideTimeInForceEnum
	FAK       SideTimeInForceEnum
	NullValue SideTimeInForceEnum
}

var SideTimeInForce = SideTimeInForceValues{0, 3, 255}

func (s SideTimeInForceEnum) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint8(_w, uint8(s)); err != nil {
		return err
	}
	return nil
}

func (s *SideTimeInForceEnum) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if err := _m.ReadUint8(_r, (*uint8)(s)); err != nil {
		return err
	}
	return nil
}

func (s SideTimeInForceEnum) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if actingVersion > schemaVersion {
		return nil
	}
	value := reflect.ValueOf(SideTimeInForce)
	for idx := 0; idx < value.NumField(); idx++ {
		if s == value.Field(idx).Interface() {
			return nil
		}
	}
	return fmt.Errorf("Range check failed on SideTimeInForce, unknown enumeration value %d", s)
}

func (*SideTimeInForceEnum) EncodedLength() int64 {
	return 1
}

func (*SideTimeInForceEnum) DaySinceVersion() uint16 {
	return 0
}

func (s *SideTimeInForceEnum) DayInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.DaySinceVersion()
}

func (*SideTimeInForceEnum) DayDeprecated() uint16 {
	return 0
}

func (*SideTimeInForceEnum) FAKSinceVersion() uint16 {
	return 0
}

func (s *SideTimeInForceEnum) FAKInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.FAKSinceVersion()
}

func (*SideTimeInForceEnum) FAKDeprecated() uint16 {
	return 0
}
