// Generated SBE (Simple Binary Encoding) message codec

package fix

import (
	"fmt"
	"io"
	"reflect"
)

type MMProtectionResetEnum byte
type MMProtectionResetValues struct {
	RESET        MMProtectionResetEnum
	DO_NOT_RESET MMProtectionResetEnum
	NullValue    MMProtectionResetEnum
}

var MMProtectionReset = MMProtectionResetValues{89, 78, 0}

func (m MMProtectionResetEnum) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint8(_w, byte(m)); err != nil {
		return err
	}
	return nil
}

func (m *MMProtectionResetEnum) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if err := _m.ReadUint8(_r, (*byte)(m)); err != nil {
		return err
	}
	return nil
}

func (m MMProtectionResetEnum) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if actingVersion > schemaVersion {
		return nil
	}
	value := reflect.ValueOf(MMProtectionReset)
	for idx := 0; idx < value.NumField(); idx++ {
		if m == value.Field(idx).Interface() {
			return nil
		}
	}
	return fmt.Errorf("Range check failed on MMProtectionReset, unknown enumeration value %d", m)
}

func (*MMProtectionResetEnum) EncodedLength() int64 {
	return 1
}

func (*MMProtectionResetEnum) RESETSinceVersion() uint16 {
	return 0
}

func (m *MMProtectionResetEnum) RESETInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.RESETSinceVersion()
}

func (*MMProtectionResetEnum) RESETDeprecated() uint16 {
	return 0
}

func (*MMProtectionResetEnum) DO_NOT_RESETSinceVersion() uint16 {
	return 0
}

func (m *MMProtectionResetEnum) DO_NOT_RESETInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.DO_NOT_RESETSinceVersion()
}

func (*MMProtectionResetEnum) DO_NOT_RESETDeprecated() uint16 {
	return 0
}
