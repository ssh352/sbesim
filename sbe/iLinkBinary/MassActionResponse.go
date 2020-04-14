// Generated SBE (Simple Binary Encoding) message codec

package ilinkbinary

import (
	"fmt"
	"io"
	"reflect"
)

type MassActionResponseEnum uint8
type MassActionResponseValues struct {
	Rejected  MassActionResponseEnum
	Accepted  MassActionResponseEnum
	NullValue MassActionResponseEnum
}

var MassActionResponse = MassActionResponseValues{0, 1, 255}

func (m MassActionResponseEnum) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint8(_w, uint8(m)); err != nil {
		return err
	}
	return nil
}

func (m *MassActionResponseEnum) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if err := _m.ReadUint8(_r, (*uint8)(m)); err != nil {
		return err
	}
	return nil
}

func (m MassActionResponseEnum) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if actingVersion > schemaVersion {
		return nil
	}
	value := reflect.ValueOf(MassActionResponse)
	for idx := 0; idx < value.NumField(); idx++ {
		if m == value.Field(idx).Interface() {
			return nil
		}
	}
	return fmt.Errorf("Range check failed on MassActionResponse, unknown enumeration value %d", m)
}

func (*MassActionResponseEnum) EncodedLength() int64 {
	return 1
}

func (*MassActionResponseEnum) RejectedSinceVersion() uint16 {
	return 0
}

func (m *MassActionResponseEnum) RejectedInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.RejectedSinceVersion()
}

func (*MassActionResponseEnum) RejectedDeprecated() uint16 {
	return 0
}

func (*MassActionResponseEnum) AcceptedSinceVersion() uint16 {
	return 0
}

func (m *MassActionResponseEnum) AcceptedInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.AcceptedSinceVersion()
}

func (*MassActionResponseEnum) AcceptedDeprecated() uint16 {
	return 0
}
