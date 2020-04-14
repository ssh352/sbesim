// Generated SBE (Simple Binary Encoding) message codec

package fix

import (
	"fmt"
	"io"
	"reflect"
)

type OpenCloseSettleFlagEnum uint16
type OpenCloseSettleFlagValues struct {
	THEORETICAL_PRICE_VALUE        OpenCloseSettleFlagEnum
	ACTUAL_PRELIMINARY_NOT_ROUNDED OpenCloseSettleFlagEnum
	ACTUAL_PRELIMINARY_ROUNDED     OpenCloseSettleFlagEnum
	NullValue                      OpenCloseSettleFlagEnum
}

var OpenCloseSettleFlag = OpenCloseSettleFlagValues{5, 100, 101, 65535}

func (o OpenCloseSettleFlagEnum) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint16(_w, uint16(o)); err != nil {
		return err
	}
	return nil
}

func (o *OpenCloseSettleFlagEnum) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if err := _m.ReadUint16(_r, (*uint16)(o)); err != nil {
		return err
	}
	return nil
}

func (o OpenCloseSettleFlagEnum) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if actingVersion > schemaVersion {
		return nil
	}
	value := reflect.ValueOf(OpenCloseSettleFlag)
	for idx := 0; idx < value.NumField(); idx++ {
		if o == value.Field(idx).Interface() {
			return nil
		}
	}
	return fmt.Errorf("Range check failed on OpenCloseSettleFlag, unknown enumeration value %d", o)
}

func (*OpenCloseSettleFlagEnum) EncodedLength() int64 {
	return 2
}

func (*OpenCloseSettleFlagEnum) THEORETICAL_PRICE_VALUESinceVersion() uint16 {
	return 0
}

func (o *OpenCloseSettleFlagEnum) THEORETICAL_PRICE_VALUEInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.THEORETICAL_PRICE_VALUESinceVersion()
}

func (*OpenCloseSettleFlagEnum) THEORETICAL_PRICE_VALUEDeprecated() uint16 {
	return 0
}

func (*OpenCloseSettleFlagEnum) ACTUAL_PRELIMINARY_NOT_ROUNDEDSinceVersion() uint16 {
	return 0
}

func (o *OpenCloseSettleFlagEnum) ACTUAL_PRELIMINARY_NOT_ROUNDEDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.ACTUAL_PRELIMINARY_NOT_ROUNDEDSinceVersion()
}

func (*OpenCloseSettleFlagEnum) ACTUAL_PRELIMINARY_NOT_ROUNDEDDeprecated() uint16 {
	return 0
}

func (*OpenCloseSettleFlagEnum) ACTUAL_PRELIMINARY_ROUNDEDSinceVersion() uint16 {
	return 0
}

func (o *OpenCloseSettleFlagEnum) ACTUAL_PRELIMINARY_ROUNDEDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.ACTUAL_PRELIMINARY_ROUNDEDSinceVersion()
}

func (*OpenCloseSettleFlagEnum) ACTUAL_PRELIMINARY_ROUNDEDDeprecated() uint16 {
	return 0
}
