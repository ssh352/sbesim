// Generated SBE (Simple Binary Encoding) message codec

package ilinkbinary

import (
	"fmt"
	"io"
	"reflect"
)

type SLEDSEnum uint8
type SLEDSValues struct {
	TradeClearingatExecutionPrice         SLEDSEnum
	TradeClearingatAlternateClearingPrice SLEDSEnum
	NullValue                             SLEDSEnum
}

var SLEDS = SLEDSValues{0, 1, 255}

func (s SLEDSEnum) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint8(_w, uint8(s)); err != nil {
		return err
	}
	return nil
}

func (s *SLEDSEnum) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if err := _m.ReadUint8(_r, (*uint8)(s)); err != nil {
		return err
	}
	return nil
}

func (s SLEDSEnum) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if actingVersion > schemaVersion {
		return nil
	}
	value := reflect.ValueOf(SLEDS)
	for idx := 0; idx < value.NumField(); idx++ {
		if s == value.Field(idx).Interface() {
			return nil
		}
	}
	return fmt.Errorf("Range check failed on SLEDS, unknown enumeration value %d", s)
}

func (*SLEDSEnum) EncodedLength() int64 {
	return 1
}

func (*SLEDSEnum) TradeClearingatExecutionPriceSinceVersion() uint16 {
	return 0
}

func (s *SLEDSEnum) TradeClearingatExecutionPriceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.TradeClearingatExecutionPriceSinceVersion()
}

func (*SLEDSEnum) TradeClearingatExecutionPriceDeprecated() uint16 {
	return 0
}

func (*SLEDSEnum) TradeClearingatAlternateClearingPriceSinceVersion() uint16 {
	return 0
}

func (s *SLEDSEnum) TradeClearingatAlternateClearingPriceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.TradeClearingatAlternateClearingPriceSinceVersion()
}

func (*SLEDSEnum) TradeClearingatAlternateClearingPriceDeprecated() uint16 {
	return 0
}
