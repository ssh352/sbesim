// Generated SBE (Simple Binary Encoding) message codec

package ilinkbinary

import (
	"fmt"
	"io"
	"reflect"
)

type MassActionOrdTypEnum byte
type MassActionOrdTypValues struct {
	Limit     MassActionOrdTypEnum
	StopLimit MassActionOrdTypEnum
	NullValue MassActionOrdTypEnum
}

var MassActionOrdTyp = MassActionOrdTypValues{50, 52, 0}

func (m MassActionOrdTypEnum) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint8(_w, byte(m)); err != nil {
		return err
	}
	return nil
}

func (m *MassActionOrdTypEnum) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if err := _m.ReadUint8(_r, (*byte)(m)); err != nil {
		return err
	}
	return nil
}

func (m MassActionOrdTypEnum) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if actingVersion > schemaVersion {
		return nil
	}
	value := reflect.ValueOf(MassActionOrdTyp)
	for idx := 0; idx < value.NumField(); idx++ {
		if m == value.Field(idx).Interface() {
			return nil
		}
	}
	return fmt.Errorf("Range check failed on MassActionOrdTyp, unknown enumeration value %d", m)
}

func (*MassActionOrdTypEnum) EncodedLength() int64 {
	return 1
}

func (*MassActionOrdTypEnum) LimitSinceVersion() uint16 {
	return 0
}

func (m *MassActionOrdTypEnum) LimitInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.LimitSinceVersion()
}

func (*MassActionOrdTypEnum) LimitDeprecated() uint16 {
	return 0
}

func (*MassActionOrdTypEnum) StopLimitSinceVersion() uint16 {
	return 0
}

func (m *MassActionOrdTypEnum) StopLimitInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.StopLimitSinceVersion()
}

func (*MassActionOrdTypEnum) StopLimitDeprecated() uint16 {
	return 0
}
