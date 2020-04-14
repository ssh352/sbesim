// Generated SBE (Simple Binary Encoding) message codec

package ilinkbinary

import (
	"fmt"
	"io"
	"reflect"
)

type SideReqEnum uint8
type SideReqValues struct {
	Buy       SideReqEnum
	Sell      SideReqEnum
	NullValue SideReqEnum
}

var SideReq = SideReqValues{1, 2, 255}

func (s SideReqEnum) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint8(_w, uint8(s)); err != nil {
		return err
	}
	return nil
}

func (s *SideReqEnum) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if err := _m.ReadUint8(_r, (*uint8)(s)); err != nil {
		return err
	}
	return nil
}

func (s SideReqEnum) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if actingVersion > schemaVersion {
		return nil
	}
	value := reflect.ValueOf(SideReq)
	for idx := 0; idx < value.NumField(); idx++ {
		if s == value.Field(idx).Interface() {
			return nil
		}
	}
	return fmt.Errorf("Range check failed on SideReq, unknown enumeration value %d", s)
}

func (*SideReqEnum) EncodedLength() int64 {
	return 1
}

func (*SideReqEnum) BuySinceVersion() uint16 {
	return 0
}

func (s *SideReqEnum) BuyInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.BuySinceVersion()
}

func (*SideReqEnum) BuyDeprecated() uint16 {
	return 0
}

func (*SideReqEnum) SellSinceVersion() uint16 {
	return 0
}

func (s *SideReqEnum) SellInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.SellSinceVersion()
}

func (*SideReqEnum) SellDeprecated() uint16 {
	return 0
}
