// Generated SBE (Simple Binary Encoding) message codec

package ilinkbinary

import (
	"fmt"
	"io"
	"reflect"
)

type MassCxlReqTypEnum uint8
type MassCxlReqTypValues struct {
	SenderSubID MassCxlReqTypEnum
	Account     MassCxlReqTypEnum
	NullValue   MassCxlReqTypEnum
}

var MassCxlReqTyp = MassCxlReqTypValues{100, 101, 255}

func (m MassCxlReqTypEnum) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint8(_w, uint8(m)); err != nil {
		return err
	}
	return nil
}

func (m *MassCxlReqTypEnum) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if err := _m.ReadUint8(_r, (*uint8)(m)); err != nil {
		return err
	}
	return nil
}

func (m MassCxlReqTypEnum) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if actingVersion > schemaVersion {
		return nil
	}
	value := reflect.ValueOf(MassCxlReqTyp)
	for idx := 0; idx < value.NumField(); idx++ {
		if m == value.Field(idx).Interface() {
			return nil
		}
	}
	return fmt.Errorf("Range check failed on MassCxlReqTyp, unknown enumeration value %d", m)
}

func (*MassCxlReqTypEnum) EncodedLength() int64 {
	return 1
}

func (*MassCxlReqTypEnum) SenderSubIDSinceVersion() uint16 {
	return 0
}

func (m *MassCxlReqTypEnum) SenderSubIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.SenderSubIDSinceVersion()
}

func (*MassCxlReqTypEnum) SenderSubIDDeprecated() uint16 {
	return 0
}

func (*MassCxlReqTypEnum) AccountSinceVersion() uint16 {
	return 0
}

func (m *MassCxlReqTypEnum) AccountInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.AccountSinceVersion()
}

func (*MassCxlReqTypEnum) AccountDeprecated() uint16 {
	return 0
}
