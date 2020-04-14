// Generated SBE (Simple Binary Encoding) message codec

package fix

import (
	"fmt"
	"io"
	"reflect"
)

type MatchEventIndicatorEnum byte
type MatchEventIndicatorValues struct {
	MID_EVENT               MatchEventIndicatorEnum
	BEGINNING_EVENT         MatchEventIndicatorEnum
	END_EVENT               MatchEventIndicatorEnum
	BEGINNING_AND_END_EVENT MatchEventIndicatorEnum
	NullValue               MatchEventIndicatorEnum
}

var MatchEventIndicator = MatchEventIndicatorValues{48, 49, 50, 51, 0}

func (m MatchEventIndicatorEnum) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint8(_w, byte(m)); err != nil {
		return err
	}
	return nil
}

func (m *MatchEventIndicatorEnum) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if err := _m.ReadUint8(_r, (*byte)(m)); err != nil {
		return err
	}
	return nil
}

func (m MatchEventIndicatorEnum) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if actingVersion > schemaVersion {
		return nil
	}
	value := reflect.ValueOf(MatchEventIndicator)
	for idx := 0; idx < value.NumField(); idx++ {
		if m == value.Field(idx).Interface() {
			return nil
		}
	}
	return fmt.Errorf("Range check failed on MatchEventIndicator, unknown enumeration value %d", m)
}

func (*MatchEventIndicatorEnum) EncodedLength() int64 {
	return 1
}

func (*MatchEventIndicatorEnum) MID_EVENTSinceVersion() uint16 {
	return 0
}

func (m *MatchEventIndicatorEnum) MID_EVENTInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.MID_EVENTSinceVersion()
}

func (*MatchEventIndicatorEnum) MID_EVENTDeprecated() uint16 {
	return 0
}

func (*MatchEventIndicatorEnum) BEGINNING_EVENTSinceVersion() uint16 {
	return 0
}

func (m *MatchEventIndicatorEnum) BEGINNING_EVENTInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.BEGINNING_EVENTSinceVersion()
}

func (*MatchEventIndicatorEnum) BEGINNING_EVENTDeprecated() uint16 {
	return 0
}

func (*MatchEventIndicatorEnum) END_EVENTSinceVersion() uint16 {
	return 0
}

func (m *MatchEventIndicatorEnum) END_EVENTInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.END_EVENTSinceVersion()
}

func (*MatchEventIndicatorEnum) END_EVENTDeprecated() uint16 {
	return 0
}

func (*MatchEventIndicatorEnum) BEGINNING_AND_END_EVENTSinceVersion() uint16 {
	return 0
}

func (m *MatchEventIndicatorEnum) BEGINNING_AND_END_EVENTInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.BEGINNING_AND_END_EVENTSinceVersion()
}

func (*MatchEventIndicatorEnum) BEGINNING_AND_END_EVENTDeprecated() uint16 {
	return 0
}
