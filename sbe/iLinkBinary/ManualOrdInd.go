// Generated SBE (Simple Binary Encoding) message codec

package ilinkbinary

import (
	"fmt"
	"io"
	"reflect"
)

type ManualOrdIndEnum uint8
type ManualOrdIndValues struct {
	Automated ManualOrdIndEnum
	Manual    ManualOrdIndEnum
	NullValue ManualOrdIndEnum
}

var ManualOrdInd = ManualOrdIndValues{0, 1, 255}

func (m ManualOrdIndEnum) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint8(_w, uint8(m)); err != nil {
		return err
	}
	return nil
}

func (m *ManualOrdIndEnum) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if err := _m.ReadUint8(_r, (*uint8)(m)); err != nil {
		return err
	}
	return nil
}

func (m ManualOrdIndEnum) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if actingVersion > schemaVersion {
		return nil
	}
	value := reflect.ValueOf(ManualOrdInd)
	for idx := 0; idx < value.NumField(); idx++ {
		if m == value.Field(idx).Interface() {
			return nil
		}
	}
	return fmt.Errorf("Range check failed on ManualOrdInd, unknown enumeration value %d", m)
}

func (*ManualOrdIndEnum) EncodedLength() int64 {
	return 1
}

func (*ManualOrdIndEnum) AutomatedSinceVersion() uint16 {
	return 0
}

func (m *ManualOrdIndEnum) AutomatedInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.AutomatedSinceVersion()
}

func (*ManualOrdIndEnum) AutomatedDeprecated() uint16 {
	return 0
}

func (*ManualOrdIndEnum) ManualSinceVersion() uint16 {
	return 0
}

func (m *ManualOrdIndEnum) ManualInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.ManualSinceVersion()
}

func (*ManualOrdIndEnum) ManualDeprecated() uint16 {
	return 0
}
