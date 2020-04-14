// Generated SBE (Simple Binary Encoding) message codec

package ilinkbinary

import (
	"fmt"
	"io"
	"reflect"
)

type ExecModeEnum byte
type ExecModeValues struct {
	Aggressive ExecModeEnum
	Passive    ExecModeEnum
	NullValue  ExecModeEnum
}

var ExecMode = ExecModeValues{65, 80, 0}

func (e ExecModeEnum) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint8(_w, byte(e)); err != nil {
		return err
	}
	return nil
}

func (e *ExecModeEnum) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if err := _m.ReadUint8(_r, (*byte)(e)); err != nil {
		return err
	}
	return nil
}

func (e ExecModeEnum) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if actingVersion > schemaVersion {
		return nil
	}
	value := reflect.ValueOf(ExecMode)
	for idx := 0; idx < value.NumField(); idx++ {
		if e == value.Field(idx).Interface() {
			return nil
		}
	}
	return fmt.Errorf("Range check failed on ExecMode, unknown enumeration value %d", e)
}

func (*ExecModeEnum) EncodedLength() int64 {
	return 1
}

func (*ExecModeEnum) AggressiveSinceVersion() uint16 {
	return 0
}

func (e *ExecModeEnum) AggressiveInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.AggressiveSinceVersion()
}

func (*ExecModeEnum) AggressiveDeprecated() uint16 {
	return 0
}

func (*ExecModeEnum) PassiveSinceVersion() uint16 {
	return 0
}

func (e *ExecModeEnum) PassiveInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.PassiveSinceVersion()
}

func (*ExecModeEnum) PassiveDeprecated() uint16 {
	return 0
}
