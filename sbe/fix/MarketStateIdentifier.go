// Generated SBE (Simple Binary Encoding) message codec

package fix

import (
	"fmt"
	"io"
	"reflect"
)

type MarketStateIdentifierEnum uint8
type MarketStateIdentifierValues struct {
	PRE_OPENING             MarketStateIdentifierEnum
	OPENING_MODE            MarketStateIdentifierEnum
	CONTINUOUS_TRADING_MODE MarketStateIdentifierEnum
	NullValue               MarketStateIdentifierEnum
}

var MarketStateIdentifier = MarketStateIdentifierValues{0, 1, 2, 255}

func (m MarketStateIdentifierEnum) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint8(_w, uint8(m)); err != nil {
		return err
	}
	return nil
}

func (m *MarketStateIdentifierEnum) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if err := _m.ReadUint8(_r, (*uint8)(m)); err != nil {
		return err
	}
	return nil
}

func (m MarketStateIdentifierEnum) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if actingVersion > schemaVersion {
		return nil
	}
	value := reflect.ValueOf(MarketStateIdentifier)
	for idx := 0; idx < value.NumField(); idx++ {
		if m == value.Field(idx).Interface() {
			return nil
		}
	}
	return fmt.Errorf("Range check failed on MarketStateIdentifier, unknown enumeration value %d", m)
}

func (*MarketStateIdentifierEnum) EncodedLength() int64 {
	return 1
}

func (*MarketStateIdentifierEnum) PRE_OPENINGSinceVersion() uint16 {
	return 0
}

func (m *MarketStateIdentifierEnum) PRE_OPENINGInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.PRE_OPENINGSinceVersion()
}

func (*MarketStateIdentifierEnum) PRE_OPENINGDeprecated() uint16 {
	return 0
}

func (*MarketStateIdentifierEnum) OPENING_MODESinceVersion() uint16 {
	return 0
}

func (m *MarketStateIdentifierEnum) OPENING_MODEInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.OPENING_MODESinceVersion()
}

func (*MarketStateIdentifierEnum) OPENING_MODEDeprecated() uint16 {
	return 0
}

func (*MarketStateIdentifierEnum) CONTINUOUS_TRADING_MODESinceVersion() uint16 {
	return 0
}

func (m *MarketStateIdentifierEnum) CONTINUOUS_TRADING_MODEInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.CONTINUOUS_TRADING_MODESinceVersion()
}

func (*MarketStateIdentifierEnum) CONTINUOUS_TRADING_MODEDeprecated() uint16 {
	return 0
}
