// Generated SBE (Simple Binary Encoding) message codec

package ilinkbinary

import (
	"fmt"
	"io"
	"reflect"
)

type MassStatusOrdTypEnum uint8
type MassStatusOrdTypValues struct {
	SenderSubID MassStatusOrdTypEnum
	Account     MassStatusOrdTypEnum
	NullValue   MassStatusOrdTypEnum
}

var MassStatusOrdTyp = MassStatusOrdTypValues{100, 101, 255}

func (m MassStatusOrdTypEnum) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint8(_w, uint8(m)); err != nil {
		return err
	}
	return nil
}

func (m *MassStatusOrdTypEnum) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if err := _m.ReadUint8(_r, (*uint8)(m)); err != nil {
		return err
	}
	return nil
}

func (m MassStatusOrdTypEnum) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if actingVersion > schemaVersion {
		return nil
	}
	value := reflect.ValueOf(MassStatusOrdTyp)
	for idx := 0; idx < value.NumField(); idx++ {
		if m == value.Field(idx).Interface() {
			return nil
		}
	}
	return fmt.Errorf("Range check failed on MassStatusOrdTyp, unknown enumeration value %d", m)
}

func (*MassStatusOrdTypEnum) EncodedLength() int64 {
	return 1
}

func (*MassStatusOrdTypEnum) SenderSubIDSinceVersion() uint16 {
	return 0
}

func (m *MassStatusOrdTypEnum) SenderSubIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.SenderSubIDSinceVersion()
}

func (*MassStatusOrdTypEnum) SenderSubIDDeprecated() uint16 {
	return 0
}

func (*MassStatusOrdTypEnum) AccountSinceVersion() uint16 {
	return 0
}

func (m *MassStatusOrdTypEnum) AccountInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.AccountSinceVersion()
}

func (*MassStatusOrdTypEnum) AccountDeprecated() uint16 {
	return 0
}
