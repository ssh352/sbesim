// Generated SBE (Simple Binary Encoding) message codec

package fix

import (
	"fmt"
	"io"
	"reflect"
)

type MDQuoteTypeEnum uint8
type MDQuoteTypeValues struct {
	TRADABLE  MDQuoteTypeEnum
	NullValue MDQuoteTypeEnum
}

var MDQuoteType = MDQuoteTypeValues{1, 255}

func (m MDQuoteTypeEnum) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint8(_w, uint8(m)); err != nil {
		return err
	}
	return nil
}

func (m *MDQuoteTypeEnum) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if err := _m.ReadUint8(_r, (*uint8)(m)); err != nil {
		return err
	}
	return nil
}

func (m MDQuoteTypeEnum) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if actingVersion > schemaVersion {
		return nil
	}
	value := reflect.ValueOf(MDQuoteType)
	for idx := 0; idx < value.NumField(); idx++ {
		if m == value.Field(idx).Interface() {
			return nil
		}
	}
	return fmt.Errorf("Range check failed on MDQuoteType, unknown enumeration value %d", m)
}

func (*MDQuoteTypeEnum) EncodedLength() int64 {
	return 1
}

func (*MDQuoteTypeEnum) TRADABLESinceVersion() uint16 {
	return 0
}

func (m *MDQuoteTypeEnum) TRADABLEInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.TRADABLESinceVersion()
}

func (*MDQuoteTypeEnum) TRADABLEDeprecated() uint16 {
	return 0
}
