// Generated SBE (Simple Binary Encoding) message codec

package ilinkbinary

import (
	"fmt"
	"io"
	"reflect"
)

type MassStatusReqTypEnum uint8
type MassStatusReqTypValues struct {
	Instrument      MassStatusReqTypEnum
	InstrumentGroup MassStatusReqTypEnum
	MarketSegment   MassStatusReqTypEnum
	NullValue       MassStatusReqTypEnum
}

var MassStatusReqTyp = MassStatusReqTypValues{1, 3, 100, 255}

func (m MassStatusReqTypEnum) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint8(_w, uint8(m)); err != nil {
		return err
	}
	return nil
}

func (m *MassStatusReqTypEnum) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if err := _m.ReadUint8(_r, (*uint8)(m)); err != nil {
		return err
	}
	return nil
}

func (m MassStatusReqTypEnum) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if actingVersion > schemaVersion {
		return nil
	}
	value := reflect.ValueOf(MassStatusReqTyp)
	for idx := 0; idx < value.NumField(); idx++ {
		if m == value.Field(idx).Interface() {
			return nil
		}
	}
	return fmt.Errorf("Range check failed on MassStatusReqTyp, unknown enumeration value %d", m)
}

func (*MassStatusReqTypEnum) EncodedLength() int64 {
	return 1
}

func (*MassStatusReqTypEnum) InstrumentSinceVersion() uint16 {
	return 0
}

func (m *MassStatusReqTypEnum) InstrumentInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.InstrumentSinceVersion()
}

func (*MassStatusReqTypEnum) InstrumentDeprecated() uint16 {
	return 0
}

func (*MassStatusReqTypEnum) InstrumentGroupSinceVersion() uint16 {
	return 0
}

func (m *MassStatusReqTypEnum) InstrumentGroupInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.InstrumentGroupSinceVersion()
}

func (*MassStatusReqTypEnum) InstrumentGroupDeprecated() uint16 {
	return 0
}

func (*MassStatusReqTypEnum) MarketSegmentSinceVersion() uint16 {
	return 0
}

func (m *MassStatusReqTypEnum) MarketSegmentInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.MarketSegmentSinceVersion()
}

func (*MassStatusReqTypEnum) MarketSegmentDeprecated() uint16 {
	return 0
}
