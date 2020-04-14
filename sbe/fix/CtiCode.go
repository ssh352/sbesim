// Generated SBE (Simple Binary Encoding) message codec

package fix

import (
	"fmt"
	"io"
	"reflect"
)

type CtiCodeEnum byte
type CtiCodeValues struct {
	OWN          CtiCodeEnum
	HOUSE        CtiCodeEnum
	ON_FLOOR     CtiCodeEnum
	NOT_ON_FLOOR CtiCodeEnum
	NullValue    CtiCodeEnum
}

var CtiCode = CtiCodeValues{49, 50, 51, 52, 0}

func (c CtiCodeEnum) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint8(_w, byte(c)); err != nil {
		return err
	}
	return nil
}

func (c *CtiCodeEnum) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if err := _m.ReadUint8(_r, (*byte)(c)); err != nil {
		return err
	}
	return nil
}

func (c CtiCodeEnum) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if actingVersion > schemaVersion {
		return nil
	}
	value := reflect.ValueOf(CtiCode)
	for idx := 0; idx < value.NumField(); idx++ {
		if c == value.Field(idx).Interface() {
			return nil
		}
	}
	return fmt.Errorf("Range check failed on CtiCode, unknown enumeration value %d", c)
}

func (*CtiCodeEnum) EncodedLength() int64 {
	return 1
}

func (*CtiCodeEnum) OWNSinceVersion() uint16 {
	return 0
}

func (c *CtiCodeEnum) OWNInActingVersion(actingVersion uint16) bool {
	return actingVersion >= c.OWNSinceVersion()
}

func (*CtiCodeEnum) OWNDeprecated() uint16 {
	return 0
}

func (*CtiCodeEnum) HOUSESinceVersion() uint16 {
	return 0
}

func (c *CtiCodeEnum) HOUSEInActingVersion(actingVersion uint16) bool {
	return actingVersion >= c.HOUSESinceVersion()
}

func (*CtiCodeEnum) HOUSEDeprecated() uint16 {
	return 0
}

func (*CtiCodeEnum) ON_FLOORSinceVersion() uint16 {
	return 0
}

func (c *CtiCodeEnum) ON_FLOORInActingVersion(actingVersion uint16) bool {
	return actingVersion >= c.ON_FLOORSinceVersion()
}

func (*CtiCodeEnum) ON_FLOORDeprecated() uint16 {
	return 0
}

func (*CtiCodeEnum) NOT_ON_FLOORSinceVersion() uint16 {
	return 0
}

func (c *CtiCodeEnum) NOT_ON_FLOORInActingVersion(actingVersion uint16) bool {
	return actingVersion >= c.NOT_ON_FLOORSinceVersion()
}

func (*CtiCodeEnum) NOT_ON_FLOORDeprecated() uint16 {
	return 0
}
