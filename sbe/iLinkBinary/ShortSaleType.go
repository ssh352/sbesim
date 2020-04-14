// Generated SBE (Simple Binary Encoding) message codec

package ilinkbinary

import (
	"fmt"
	"io"
	"reflect"
)

type ShortSaleTypeEnum uint8
type ShortSaleTypeValues struct {
	LongSell                                   ShortSaleTypeEnum
	ShortSaleWithNoExemptionSESH               ShortSaleTypeEnum
	ShortSaleWithExemptionSSEX                 ShortSaleTypeEnum
	UndisclosedSellInformationNotAvailableUNDI ShortSaleTypeEnum
	NullValue                                  ShortSaleTypeEnum
}

var ShortSaleType = ShortSaleTypeValues{0, 1, 2, 3, 255}

func (s ShortSaleTypeEnum) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint8(_w, uint8(s)); err != nil {
		return err
	}
	return nil
}

func (s *ShortSaleTypeEnum) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if err := _m.ReadUint8(_r, (*uint8)(s)); err != nil {
		return err
	}
	return nil
}

func (s ShortSaleTypeEnum) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if actingVersion > schemaVersion {
		return nil
	}
	value := reflect.ValueOf(ShortSaleType)
	for idx := 0; idx < value.NumField(); idx++ {
		if s == value.Field(idx).Interface() {
			return nil
		}
	}
	return fmt.Errorf("Range check failed on ShortSaleType, unknown enumeration value %d", s)
}

func (*ShortSaleTypeEnum) EncodedLength() int64 {
	return 1
}

func (*ShortSaleTypeEnum) LongSellSinceVersion() uint16 {
	return 0
}

func (s *ShortSaleTypeEnum) LongSellInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.LongSellSinceVersion()
}

func (*ShortSaleTypeEnum) LongSellDeprecated() uint16 {
	return 0
}

func (*ShortSaleTypeEnum) ShortSaleWithNoExemptionSESHSinceVersion() uint16 {
	return 0
}

func (s *ShortSaleTypeEnum) ShortSaleWithNoExemptionSESHInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.ShortSaleWithNoExemptionSESHSinceVersion()
}

func (*ShortSaleTypeEnum) ShortSaleWithNoExemptionSESHDeprecated() uint16 {
	return 0
}

func (*ShortSaleTypeEnum) ShortSaleWithExemptionSSEXSinceVersion() uint16 {
	return 0
}

func (s *ShortSaleTypeEnum) ShortSaleWithExemptionSSEXInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.ShortSaleWithExemptionSSEXSinceVersion()
}

func (*ShortSaleTypeEnum) ShortSaleWithExemptionSSEXDeprecated() uint16 {
	return 0
}

func (*ShortSaleTypeEnum) UndisclosedSellInformationNotAvailableUNDISinceVersion() uint16 {
	return 0
}

func (s *ShortSaleTypeEnum) UndisclosedSellInformationNotAvailableUNDIInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.UndisclosedSellInformationNotAvailableUNDISinceVersion()
}

func (*ShortSaleTypeEnum) UndisclosedSellInformationNotAvailableUNDIDeprecated() uint16 {
	return 0
}
