// Generated SBE (Simple Binary Encoding) message codec

package ilinkbinary

import (
	"fmt"
	"io"
	"reflect"
)

type MassActionScopeEnum uint8
type MassActionScopeValues struct {
	Instrument      MassActionScopeEnum
	All             MassActionScopeEnum
	MarketSegmentID MassActionScopeEnum
	InstrumentGroup MassActionScopeEnum
	QuoteSetID      MassActionScopeEnum
	NullValue       MassActionScopeEnum
}

var MassActionScope = MassActionScopeValues{1, 7, 9, 10, 100, 255}

func (m MassActionScopeEnum) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint8(_w, uint8(m)); err != nil {
		return err
	}
	return nil
}

func (m *MassActionScopeEnum) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if err := _m.ReadUint8(_r, (*uint8)(m)); err != nil {
		return err
	}
	return nil
}

func (m MassActionScopeEnum) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if actingVersion > schemaVersion {
		return nil
	}
	value := reflect.ValueOf(MassActionScope)
	for idx := 0; idx < value.NumField(); idx++ {
		if m == value.Field(idx).Interface() {
			return nil
		}
	}
	return fmt.Errorf("Range check failed on MassActionScope, unknown enumeration value %d", m)
}

func (*MassActionScopeEnum) EncodedLength() int64 {
	return 1
}

func (*MassActionScopeEnum) InstrumentSinceVersion() uint16 {
	return 0
}

func (m *MassActionScopeEnum) InstrumentInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.InstrumentSinceVersion()
}

func (*MassActionScopeEnum) InstrumentDeprecated() uint16 {
	return 0
}

func (*MassActionScopeEnum) AllSinceVersion() uint16 {
	return 0
}

func (m *MassActionScopeEnum) AllInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.AllSinceVersion()
}

func (*MassActionScopeEnum) AllDeprecated() uint16 {
	return 0
}

func (*MassActionScopeEnum) MarketSegmentIDSinceVersion() uint16 {
	return 0
}

func (m *MassActionScopeEnum) MarketSegmentIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.MarketSegmentIDSinceVersion()
}

func (*MassActionScopeEnum) MarketSegmentIDDeprecated() uint16 {
	return 0
}

func (*MassActionScopeEnum) InstrumentGroupSinceVersion() uint16 {
	return 0
}

func (m *MassActionScopeEnum) InstrumentGroupInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.InstrumentGroupSinceVersion()
}

func (*MassActionScopeEnum) InstrumentGroupDeprecated() uint16 {
	return 0
}

func (*MassActionScopeEnum) QuoteSetIDSinceVersion() uint16 {
	return 0
}

func (m *MassActionScopeEnum) QuoteSetIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.QuoteSetIDSinceVersion()
}

func (*MassActionScopeEnum) QuoteSetIDDeprecated() uint16 {
	return 0
}
