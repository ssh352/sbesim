// Generated SBE (Simple Binary Encoding) message codec

package ilinkbinary

import (
	"fmt"
	"io"
	"reflect"
)

type ClearingAcctTypeEnum uint8
type ClearingAcctTypeValues struct {
	Customer  ClearingAcctTypeEnum
	Firm      ClearingAcctTypeEnum
	NullValue ClearingAcctTypeEnum
}

var ClearingAcctType = ClearingAcctTypeValues{0, 1, 255}

func (c ClearingAcctTypeEnum) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint8(_w, uint8(c)); err != nil {
		return err
	}
	return nil
}

func (c *ClearingAcctTypeEnum) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if err := _m.ReadUint8(_r, (*uint8)(c)); err != nil {
		return err
	}
	return nil
}

func (c ClearingAcctTypeEnum) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if actingVersion > schemaVersion {
		return nil
	}
	value := reflect.ValueOf(ClearingAcctType)
	for idx := 0; idx < value.NumField(); idx++ {
		if c == value.Field(idx).Interface() {
			return nil
		}
	}
	return fmt.Errorf("Range check failed on ClearingAcctType, unknown enumeration value %d", c)
}

func (*ClearingAcctTypeEnum) EncodedLength() int64 {
	return 1
}

func (*ClearingAcctTypeEnum) CustomerSinceVersion() uint16 {
	return 0
}

func (c *ClearingAcctTypeEnum) CustomerInActingVersion(actingVersion uint16) bool {
	return actingVersion >= c.CustomerSinceVersion()
}

func (*ClearingAcctTypeEnum) CustomerDeprecated() uint16 {
	return 0
}

func (*ClearingAcctTypeEnum) FirmSinceVersion() uint16 {
	return 0
}

func (c *ClearingAcctTypeEnum) FirmInActingVersion(actingVersion uint16) bool {
	return actingVersion >= c.FirmSinceVersion()
}

func (*ClearingAcctTypeEnum) FirmDeprecated() uint16 {
	return 0
}
