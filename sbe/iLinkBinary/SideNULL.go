// Generated SBE (Simple Binary Encoding) message codec

package ilinkbinary

import (
	"fmt"
	"io"
	"reflect"
)

type SideNULLEnum uint8
type SideNULLValues struct {
	Buy       SideNULLEnum
	Sell      SideNULLEnum
	NullValue SideNULLEnum
}

var SideNULL = SideNULLValues{1, 2, 255}

func (s SideNULLEnum) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint8(_w, uint8(s)); err != nil {
		return err
	}
	return nil
}

func (s *SideNULLEnum) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if err := _m.ReadUint8(_r, (*uint8)(s)); err != nil {
		return err
	}
	return nil
}

func (s SideNULLEnum) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if actingVersion > schemaVersion {
		return nil
	}
	value := reflect.ValueOf(SideNULL)
	for idx := 0; idx < value.NumField(); idx++ {
		if s == value.Field(idx).Interface() {
			return nil
		}
	}
	return fmt.Errorf("Range check failed on SideNULL, unknown enumeration value %d", s)
}

func (*SideNULLEnum) EncodedLength() int64 {
	return 1
}

func (*SideNULLEnum) BuySinceVersion() uint16 {
	return 0
}

func (s *SideNULLEnum) BuyInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.BuySinceVersion()
}

func (*SideNULLEnum) BuyDeprecated() uint16 {
	return 0
}

func (*SideNULLEnum) SellSinceVersion() uint16 {
	return 0
}

func (s *SideNULLEnum) SellInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.SellSinceVersion()
}

func (*SideNULLEnum) SellDeprecated() uint16 {
	return 0
}
