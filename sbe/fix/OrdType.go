// Generated SBE (Simple Binary Encoding) message codec

package fix

import (
	"fmt"
	"io"
	"reflect"
)

type OrdTypeEnum byte
type OrdTypeValues struct {
	MARKET_ORDER       OrdTypeEnum
	LIMIT_ORDER        OrdTypeEnum
	STOP_ORDER         OrdTypeEnum
	STOP_LIMIT_ORDER   OrdTypeEnum
	MARKET_LIMIT_ORDER OrdTypeEnum
	NullValue          OrdTypeEnum
}

var OrdType = OrdTypeValues{49, 50, 51, 52, 75, 0}

func (o OrdTypeEnum) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint8(_w, byte(o)); err != nil {
		return err
	}
	return nil
}

func (o *OrdTypeEnum) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if err := _m.ReadUint8(_r, (*byte)(o)); err != nil {
		return err
	}
	return nil
}

func (o OrdTypeEnum) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if actingVersion > schemaVersion {
		return nil
	}
	value := reflect.ValueOf(OrdType)
	for idx := 0; idx < value.NumField(); idx++ {
		if o == value.Field(idx).Interface() {
			return nil
		}
	}
	return fmt.Errorf("Range check failed on OrdType, unknown enumeration value %d", o)
}

func (*OrdTypeEnum) EncodedLength() int64 {
	return 1
}

func (*OrdTypeEnum) MARKET_ORDERSinceVersion() uint16 {
	return 0
}

func (o *OrdTypeEnum) MARKET_ORDERInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.MARKET_ORDERSinceVersion()
}

func (*OrdTypeEnum) MARKET_ORDERDeprecated() uint16 {
	return 0
}

func (*OrdTypeEnum) LIMIT_ORDERSinceVersion() uint16 {
	return 0
}

func (o *OrdTypeEnum) LIMIT_ORDERInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.LIMIT_ORDERSinceVersion()
}

func (*OrdTypeEnum) LIMIT_ORDERDeprecated() uint16 {
	return 0
}

func (*OrdTypeEnum) STOP_ORDERSinceVersion() uint16 {
	return 0
}

func (o *OrdTypeEnum) STOP_ORDERInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.STOP_ORDERSinceVersion()
}

func (*OrdTypeEnum) STOP_ORDERDeprecated() uint16 {
	return 0
}

func (*OrdTypeEnum) STOP_LIMIT_ORDERSinceVersion() uint16 {
	return 0
}

func (o *OrdTypeEnum) STOP_LIMIT_ORDERInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.STOP_LIMIT_ORDERSinceVersion()
}

func (*OrdTypeEnum) STOP_LIMIT_ORDERDeprecated() uint16 {
	return 0
}

func (*OrdTypeEnum) MARKET_LIMIT_ORDERSinceVersion() uint16 {
	return 0
}

func (o *OrdTypeEnum) MARKET_LIMIT_ORDERInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.MARKET_LIMIT_ORDERSinceVersion()
}

func (*OrdTypeEnum) MARKET_LIMIT_ORDERDeprecated() uint16 {
	return 0
}
