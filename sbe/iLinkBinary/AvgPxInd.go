// Generated SBE (Simple Binary Encoding) message codec

package ilinkbinary

import (
	"fmt"
	"io"
	"reflect"
)

type AvgPxIndEnum uint8
type AvgPxIndValues struct {
	NoAveragePricing                                          AvgPxIndEnum
	TradeispartofanAveragePriceGroupIdentifiedbytheAvgPxGrpID AvgPxIndEnum
	TradeispartofaNotionalValueAveragePriceGroup              AvgPxIndEnum
	NullValue                                                 AvgPxIndEnum
}

var AvgPxInd = AvgPxIndValues{0, 1, 3, 255}

func (a AvgPxIndEnum) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint8(_w, uint8(a)); err != nil {
		return err
	}
	return nil
}

func (a *AvgPxIndEnum) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if err := _m.ReadUint8(_r, (*uint8)(a)); err != nil {
		return err
	}
	return nil
}

func (a AvgPxIndEnum) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if actingVersion > schemaVersion {
		return nil
	}
	value := reflect.ValueOf(AvgPxInd)
	for idx := 0; idx < value.NumField(); idx++ {
		if a == value.Field(idx).Interface() {
			return nil
		}
	}
	return fmt.Errorf("Range check failed on AvgPxInd, unknown enumeration value %d", a)
}

func (*AvgPxIndEnum) EncodedLength() int64 {
	return 1
}

func (*AvgPxIndEnum) NoAveragePricingSinceVersion() uint16 {
	return 0
}

func (a *AvgPxIndEnum) NoAveragePricingInActingVersion(actingVersion uint16) bool {
	return actingVersion >= a.NoAveragePricingSinceVersion()
}

func (*AvgPxIndEnum) NoAveragePricingDeprecated() uint16 {
	return 0
}

func (*AvgPxIndEnum) TradeispartofanAveragePriceGroupIdentifiedbytheAvgPxGrpIDSinceVersion() uint16 {
	return 0
}

func (a *AvgPxIndEnum) TradeispartofanAveragePriceGroupIdentifiedbytheAvgPxGrpIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= a.TradeispartofanAveragePriceGroupIdentifiedbytheAvgPxGrpIDSinceVersion()
}

func (*AvgPxIndEnum) TradeispartofanAveragePriceGroupIdentifiedbytheAvgPxGrpIDDeprecated() uint16 {
	return 0
}

func (*AvgPxIndEnum) TradeispartofaNotionalValueAveragePriceGroupSinceVersion() uint16 {
	return 0
}

func (a *AvgPxIndEnum) TradeispartofaNotionalValueAveragePriceGroupInActingVersion(actingVersion uint16) bool {
	return actingVersion >= a.TradeispartofaNotionalValueAveragePriceGroupSinceVersion()
}

func (*AvgPxIndEnum) TradeispartofaNotionalValueAveragePriceGroupDeprecated() uint16 {
	return 0
}
