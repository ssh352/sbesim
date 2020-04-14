// Generated SBE (Simple Binary Encoding) message codec

package fix

import (
	"fmt"
	"io"
	"reflect"
)

type MDUpdateActionEnum uint8
type MDUpdateActionValues struct {
	NEW       MDUpdateActionEnum
	CHANGE    MDUpdateActionEnum
	DELETE    MDUpdateActionEnum
	OVERLAY   MDUpdateActionEnum
	NullValue MDUpdateActionEnum
}

var MDUpdateAction = MDUpdateActionValues{0, 1, 2, 5, 255}

func (m MDUpdateActionEnum) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint8(_w, uint8(m)); err != nil {
		return err
	}
	return nil
}

func (m *MDUpdateActionEnum) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if err := _m.ReadUint8(_r, (*uint8)(m)); err != nil {
		return err
	}
	return nil
}

func (m MDUpdateActionEnum) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if actingVersion > schemaVersion {
		return nil
	}
	value := reflect.ValueOf(MDUpdateAction)
	for idx := 0; idx < value.NumField(); idx++ {
		if m == value.Field(idx).Interface() {
			return nil
		}
	}
	return fmt.Errorf("Range check failed on MDUpdateAction, unknown enumeration value %d", m)
}

func (*MDUpdateActionEnum) EncodedLength() int64 {
	return 1
}

func (*MDUpdateActionEnum) NEWSinceVersion() uint16 {
	return 0
}

func (m *MDUpdateActionEnum) NEWInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.NEWSinceVersion()
}

func (*MDUpdateActionEnum) NEWDeprecated() uint16 {
	return 0
}

func (*MDUpdateActionEnum) CHANGESinceVersion() uint16 {
	return 0
}

func (m *MDUpdateActionEnum) CHANGEInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.CHANGESinceVersion()
}

func (*MDUpdateActionEnum) CHANGEDeprecated() uint16 {
	return 0
}

func (*MDUpdateActionEnum) DELETESinceVersion() uint16 {
	return 0
}

func (m *MDUpdateActionEnum) DELETEInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.DELETESinceVersion()
}

func (*MDUpdateActionEnum) DELETEDeprecated() uint16 {
	return 0
}

func (*MDUpdateActionEnum) OVERLAYSinceVersion() uint16 {
	return 0
}

func (m *MDUpdateActionEnum) OVERLAYInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.OVERLAYSinceVersion()
}

func (*MDUpdateActionEnum) OVERLAYDeprecated() uint16 {
	return 0
}
