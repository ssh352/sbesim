// Generated SBE (Simple Binary Encoding) message codec

package ilinkbinary

import (
	"fmt"
	"io"
	"reflect"
)

type MassCancelTIFEnum uint8
type MassCancelTIFValues struct {
	Day            MassCancelTIFEnum
	GoodTillCancel MassCancelTIFEnum
	GoodTillDate   MassCancelTIFEnum
	NullValue      MassCancelTIFEnum
}

var MassCancelTIF = MassCancelTIFValues{0, 1, 6, 255}

func (m MassCancelTIFEnum) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint8(_w, uint8(m)); err != nil {
		return err
	}
	return nil
}

func (m *MassCancelTIFEnum) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if err := _m.ReadUint8(_r, (*uint8)(m)); err != nil {
		return err
	}
	return nil
}

func (m MassCancelTIFEnum) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if actingVersion > schemaVersion {
		return nil
	}
	value := reflect.ValueOf(MassCancelTIF)
	for idx := 0; idx < value.NumField(); idx++ {
		if m == value.Field(idx).Interface() {
			return nil
		}
	}
	return fmt.Errorf("Range check failed on MassCancelTIF, unknown enumeration value %d", m)
}

func (*MassCancelTIFEnum) EncodedLength() int64 {
	return 1
}

func (*MassCancelTIFEnum) DaySinceVersion() uint16 {
	return 0
}

func (m *MassCancelTIFEnum) DayInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.DaySinceVersion()
}

func (*MassCancelTIFEnum) DayDeprecated() uint16 {
	return 0
}

func (*MassCancelTIFEnum) GoodTillCancelSinceVersion() uint16 {
	return 0
}

func (m *MassCancelTIFEnum) GoodTillCancelInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.GoodTillCancelSinceVersion()
}

func (*MassCancelTIFEnum) GoodTillCancelDeprecated() uint16 {
	return 0
}

func (*MassCancelTIFEnum) GoodTillDateSinceVersion() uint16 {
	return 0
}

func (m *MassCancelTIFEnum) GoodTillDateInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.GoodTillDateSinceVersion()
}

func (*MassCancelTIFEnum) GoodTillDateDeprecated() uint16 {
	return 0
}
