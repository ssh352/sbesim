// Generated SBE (Simple Binary Encoding) message codec

package ilinkbinary

import (
	"fmt"
	"io"
	"reflect"
)

type CmtaGiveUpCDEnum byte
type CmtaGiveUpCDValues struct {
	GiveUp    CmtaGiveUpCDEnum
	SGXoffset CmtaGiveUpCDEnum
	NullValue CmtaGiveUpCDEnum
}

var CmtaGiveUpCD = CmtaGiveUpCDValues{71, 83, 0}

func (c CmtaGiveUpCDEnum) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint8(_w, byte(c)); err != nil {
		return err
	}
	return nil
}

func (c *CmtaGiveUpCDEnum) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if err := _m.ReadUint8(_r, (*byte)(c)); err != nil {
		return err
	}
	return nil
}

func (c CmtaGiveUpCDEnum) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if actingVersion > schemaVersion {
		return nil
	}
	value := reflect.ValueOf(CmtaGiveUpCD)
	for idx := 0; idx < value.NumField(); idx++ {
		if c == value.Field(idx).Interface() {
			return nil
		}
	}
	return fmt.Errorf("Range check failed on CmtaGiveUpCD, unknown enumeration value %d", c)
}

func (*CmtaGiveUpCDEnum) EncodedLength() int64 {
	return 1
}

func (*CmtaGiveUpCDEnum) GiveUpSinceVersion() uint16 {
	return 0
}

func (c *CmtaGiveUpCDEnum) GiveUpInActingVersion(actingVersion uint16) bool {
	return actingVersion >= c.GiveUpSinceVersion()
}

func (*CmtaGiveUpCDEnum) GiveUpDeprecated() uint16 {
	return 0
}

func (*CmtaGiveUpCDEnum) SGXoffsetSinceVersion() uint16 {
	return 0
}

func (c *CmtaGiveUpCDEnum) SGXoffsetInActingVersion(actingVersion uint16) bool {
	return actingVersion >= c.SGXoffsetSinceVersion()
}

func (*CmtaGiveUpCDEnum) SGXoffsetDeprecated() uint16 {
	return 0
}
