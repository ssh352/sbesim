// Generated SBE (Simple Binary Encoding) message codec

package fix

import (
	"fmt"
	"io"
	"reflect"
)

type CustomerOrFirmEnum uint8
type CustomerOrFirmValues struct {
	CUSTOMER  CustomerOrFirmEnum
	FIRM      CustomerOrFirmEnum
	NullValue CustomerOrFirmEnum
}

var CustomerOrFirm = CustomerOrFirmValues{0, 1, 255}

func (c CustomerOrFirmEnum) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint8(_w, uint8(c)); err != nil {
		return err
	}
	return nil
}

func (c *CustomerOrFirmEnum) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if err := _m.ReadUint8(_r, (*uint8)(c)); err != nil {
		return err
	}
	return nil
}

func (c CustomerOrFirmEnum) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if actingVersion > schemaVersion {
		return nil
	}
	value := reflect.ValueOf(CustomerOrFirm)
	for idx := 0; idx < value.NumField(); idx++ {
		if c == value.Field(idx).Interface() {
			return nil
		}
	}
	return fmt.Errorf("Range check failed on CustomerOrFirm, unknown enumeration value %d", c)
}

func (*CustomerOrFirmEnum) EncodedLength() int64 {
	return 1
}

func (*CustomerOrFirmEnum) CUSTOMERSinceVersion() uint16 {
	return 0
}

func (c *CustomerOrFirmEnum) CUSTOMERInActingVersion(actingVersion uint16) bool {
	return actingVersion >= c.CUSTOMERSinceVersion()
}

func (*CustomerOrFirmEnum) CUSTOMERDeprecated() uint16 {
	return 0
}

func (*CustomerOrFirmEnum) FIRMSinceVersion() uint16 {
	return 0
}

func (c *CustomerOrFirmEnum) FIRMInActingVersion(actingVersion uint16) bool {
	return actingVersion >= c.FIRMSinceVersion()
}

func (*CustomerOrFirmEnum) FIRMDeprecated() uint16 {
	return 0
}
